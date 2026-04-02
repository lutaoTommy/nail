package handler

import (
	"os"
	"path/filepath"
	"sync"
	"testing"

	"nail/config"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
)

// 这些测试只覆盖不依赖数据库的分支，用于快速回归校验 user.go 的基本校验逻辑。

func newUserApp() *iris.Application {
	app := iris.New()
	app.PartyFunc("/user", UserHandler)
	return app
}

var (
	configOnce sync.Once
	configErr  error
)

// 确保当前工作目录在项目根（能找到 config.ini），兼容 go test ./handler 的相对路径问题。
func ensureProjectRoot(t *testing.T) {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd failed: %v", err)
	}
	dir := wd
	for i := 0; i < 6; i++ {
		if _, err := os.Stat(filepath.Join(dir, "config.ini")); err == nil {
			if err := os.Chdir(dir); err != nil {
				t.Fatalf("Chdir(%s) failed: %v", dir, err)
			}
			return
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	t.Fatalf("config.ini not found (start from %s)", wd)
}

func loadConfigOnce(t *testing.T) {
	t.Helper()
	configOnce.Do(func() {
		ensureProjectRoot(t)
		configErr = config.LoadConfig()
	})
	if configErr != nil {
		t.Fatalf("LoadConfig failed: %v", configErr)
	}
}

// 准备一个 phone=\"test\" 的用户并返回其明文密码，依赖真实数据库。
func prepareTestUser(t *testing.T) (User, string) {
	loadConfigOnce(t)
	db := getMysqlConn()

	const phone = "test"
	const email = "test@example.com"
	const userId = "test"
	const token = "test"
	const nickname = "测试用户"
	const lang = "zh-CN"
	const plainPwd = "123456"

	hashed, err := HashPassword(plainPwd)
	if err != nil {
		t.Fatalf("HashPassword failed: %v", err)
	}

	// 每次测试前重置 test 用户：先删后建，保证干净初始状态
	if err := db.Where("user_id = ?", userId).Delete(&User{}).Error; err != nil {
		t.Fatalf("reset test user failed: %v", err)
	}

	u := User{
		Phone:    phone,
		Email:    email,
		Token:    token,
		Passwd:   hashed,
		UserId:   userId,
		Language: lang,
		Nickname: nickname,
		Status:   1,
	}
	if err := db.Create(&u).Error; err != nil {
		t.Fatalf("prepare user failed: %v", err)
	}
	return u, plainPwd
}

func prepareAdminUser(t *testing.T) {
	t.Helper()
	loadConfigOnce(t)
	db := getMysqlConn()
	pwd, err := HashPassword("fndroid")
	if err != nil {
		t.Fatalf("HashPassword admin failed: %v", err)
	}
	admin := User{
		Phone:    "admin",
		Token:    "admin",
		Passwd:   pwd,
		UserId:   "admin",
		Nickname: "管理员",
		Language: "zh-CN",
		Status:   1,
	}
	if err := db.Save(&admin).Error; err != nil {
		t.Fatalf("prepare admin failed: %v", err)
	}
}

func resetRateLimiterForTest() {
	loginMu.Lock()
	loginLimitByIP = make(map[string]*limitRec)
	loginLimitByAcc = make(map[string]*limitRec)
	loginMu.Unlock()

	vcodeMu.Lock()
	vcodeLimitByIP = make(map[string]*limitRec)
	vcodeLimitByEmail = make(map[string]*limitRec)
	vcodeMu.Unlock()
}

// /user/login：缺少手机号时应返回 E_NO_PHONE（400）
func TestUserLoginHandler_NoPhone(t *testing.T) {
	app := newUserApp()
	e := httptest.New(t, app)

	resp := e.POST("/user/login").
		WithJSON(iris.Map{
			"passwd": "123456",
		}).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	resp.Value("result_code").Number().Equal(400)
	resp.Value("result_msg").String().NotEmpty()
}

// /user/login/mail：缺少邮箱时应返回 E_NO_EMAIL（400）
func TestMailLoginHandler_NoEmail(t *testing.T) {
	app := newUserApp()
	e := httptest.New(t, app)

	resp := e.POST("/user/login/mail").
		WithJSON(iris.Map{
			"passwd": "123456",
		}).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	resp.Value("result_code").Number().Equal(400)
	resp.Value("result_msg").String().NotEmpty()
}

// /user/passwd/update：缺少 token 时应返回 E_NO_TOKEN（401）
func TestChangePasswdHandler_NoToken(t *testing.T) {
	app := newUserApp()
	e := httptest.New(t, app)

	resp := e.POST("/user/passwd/update").
		WithJSON(iris.Map{
			"old_passwd": "old",
			"new_passwd": "new",
		}).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	resp.Value("result_code").Number().Equal(401)
	resp.Value("result_msg").String().NotEmpty()
}

// /user/register/verification/mail：邮箱格式非法时应返回 E_INVALID_EMAIL（400）
func TestMailRegisterVerificationHandler_InvalidEmail(t *testing.T) {
	app := newUserApp()
	e := httptest.New(t, app)

	resp := e.POST("/user/register/verification/mail").
		WithJSON(iris.Map{
			"email": "not-an-email",
		}).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	resp.Value("result_code").Number().Equal(400)
	resp.Value("result_msg").String().NotEmpty()
}

// 使用 phone=\"test\" 的账号做一次真实 DB 登录测试
// 依赖 config.ini 中的 MySQL 配置以及可写的 users 表。
func TestUserLoginHandler_WithDBPhoneTest(t *testing.T) {
	u, plainPwd := prepareTestUser(t)

	app := newUserApp()
	e := httptest.New(t, app)

	resp := e.POST("/user/login").
		WithJSON(iris.Map{
			"phone":  u.Phone,
			"passwd": plainPwd,
		}).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	resp.Value("result_code").Number().Equal(200)
	resp.Value("token").String().NotEmpty()
}

// 先修改用户信息和密码，再通过 /user/info 验证修改结果
func TestUpdateUserInfoAndPasswd_ThenGetUserInfo(t *testing.T) {
	u, plainPwd := prepareTestUser(t)

	app := newUserApp()
	e := httptest.New(t, app)

	// 先登录获取初始 token
	login := e.POST("/user/login").
		WithJSON(iris.Map{
			"phone":  u.Phone,
			"passwd": plainPwd,
		}).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	token := login.Value("token").String().Raw()

	// 修改用户资料
	newNickname := "test_nick"
	newBio := "hello from combined test"

	updateResp := e.POST("/user/info/update").
		WithHeader("token", token).
		WithJSON(iris.Map{
			"nickname":  newNickname,
			"biography": newBio,
		}).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	updateResp.Value("result_code").Number().Equal(200)

	// 修改密码（会返回新的 token）
	newPwd := "newpassword123"

	passwdResp := e.POST("/user/passwd/update").
		WithHeader("token", token).
		WithJSON(iris.Map{
			"old_passwd": plainPwd,
			"new_passwd": newPwd,
		}).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	passwdResp.Value("result_code").Number().Equal(200)
	newToken := passwdResp.Value("token").String().Raw()

	// 使用新 token 获取用户信息，并校验前面两次修改是否生效
	info := e.GET("/user/info").
		WithHeader("token", newToken).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	info.Value("result_code").Number().Equal(200)
	info.Value("phone").String().Equal(u.Phone)
	info.Value("email").String().Equal(u.Email)
	info.Value("nickname").String().Equal(newNickname)
	info.Value("biography").String().Equal(newBio)
}

// 使用 phone="test" 的账号测试 /user/login/mail（邮箱登录）
func TestMailLoginHandler_WithDBEmailTest(t *testing.T) {
	u, plainPwd := prepareTestUser(t)

	app := newUserApp()
	e := httptest.New(t, app)

	resp := e.POST("/user/login/mail").
		WithJSON(iris.Map{
			"email":  u.Email,
			"passwd": plainPwd,
		}).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	resp.Value("result_code").Number().Equal(200)
	resp.Value("token").String().NotEmpty()
}

// /user/list：需要 admin token（非注册功能）
func TestUserListHandler_WithDBAdmin(t *testing.T) {
	prepareAdminUser(t)

	app := newUserApp()
	e := httptest.New(t, app)

	resp := e.GET("/user/list").
		WithHeader("token", "admin").
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	resp.Value("result_code").Number().Equal(200)
	resp.Value("data").Array().NotEmpty()
}

// 登录失败 5 次后应被限流（429）
func TestUserLoginHandler_RateLimitAfterFailures(t *testing.T) {
	resetRateLimiterForTest()
	u, _ := prepareTestUser(t)

	app := newUserApp()
	e := httptest.New(t, app)

	for i := 0; i < 5; i++ {
		r := e.POST("/user/login").
			WithJSON(iris.Map{
				"phone":  u.Phone,
				"passwd": "wrong_password",
			}).
			Expect().
			Status(iris.StatusOK).
			JSON().Object()
		r.Value("result_code").Number().Equal(400) // E_INVALID_PWD
	}

	locked := e.POST("/user/login").
		WithJSON(iris.Map{
			"phone":  u.Phone,
			"passwd": "wrong_password",
		}).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	locked.Value("result_code").Number().Equal(429)
	locked.Value("retry_after").Number().Gt(0)
}

