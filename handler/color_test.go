package handler

import (
	"testing"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/httptest"
)

func newColorApp() *iris.Application {
	app := iris.New()
	app.PartyFunc("/color", ColorHandler)
	return app
}

// 同时挂载 /user 和 /color，用于需要先登录再请求 color 的用例
func newUserAndColorApp() *iris.Application {
	app := iris.New()
	app.PartyFunc("/user", UserHandler)
	app.PartyFunc("/color", ColorHandler)
	return app
}

func getAnyColorID(t *testing.T) string {
	t.Helper()
	loadConfigOnce(t)
	db := getMysqlConn()

	var c Color
	// 只取一个 id 即可，不依赖具体种子数据
	if err := db.Table("colors").Select("id").First(&c).Error; err != nil {
		t.Skipf("colors table unavailable or empty: %v", err)
	}
	if c.Id == "" {
		t.Skip("colors table returned empty id")
	}
	return c.Id
}

// /color/list：缺少 token 时应返回 401
func TestColorListHandler_NoToken(t *testing.T) {
	app := newColorApp()
	e := httptest.New(t, app)

	resp := e.GET("/color/list").
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	resp.Value("result_code").Number().Equal(401)
	resp.Value("result_msg").String().NotEmpty()
}

// 使用 phone="test" 的账号测试 /color/list
func TestColorListHandler_WithDBPhoneTest(t *testing.T) {
	u, plainPwd := prepareTestUser(t)

	app := newUserAndColorApp()
	e := httptest.New(t, app)

	// 先登录拿到 token
	login := e.POST("/user/login").
		WithJSON(iris.Map{
			"phone":  u.Phone,
			"passwd": plainPwd,
		}).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	token := login.Value("token").String().Raw()

	resp := e.GET("/color/list").
		WithHeader("token", token).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	resp.Value("result_code").Number().Equal(200)
	resp.Value("data").Array().NotEmpty()
}

// /color/list：一个颜色归属多个色系时，按任一行色系都能查到
func TestColorListHandler_MultiGroup_WithDBPhoneTest(t *testing.T) {
	u, plainPwd := prepareTestUser(t)
	loadConfigOnce(t)
	db := getMysqlConn()

	const colorID = "T9001"
	_ = db.Where("id = ?", colorID).Delete(&Color{}).Error
	color := Color{
		Id:      colorID,
		Name:    "多色系测试",
		Color:   "TEST",
		Desc:    "#000000",
		Count:   61,
		GroupId: []int{1, 7},
	}
	if err := db.Save(&color).Error; err != nil {
		t.Fatalf("insert test color failed: %v", err)
	}
	defer db.Where("id = ?", colorID).Delete(&Color{})

	app := newUserAndColorApp()
	e := httptest.New(t, app)

	login := e.POST("/user/login").
		WithJSON(iris.Map{
			"phone":  u.Phone,
			"passwd": plainPwd,
		}).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	token := login.Value("token").String().Raw()

	for _, index := range []int{1, 7} {
		resp := e.GET("/color/list").
			WithHeader("token", token).
			WithQuery("count", 61).
			WithQuery("index", index).
			Expect().
			Status(iris.StatusOK).
			JSON().Object()

		resp.Value("result_code").Number().Equal(200)
		arr := resp.Value("data").Array()
		found := false
		for i := 0; i < len(arr.Raw()); i++ {
			obj := arr.Element(i).Object()
			if obj.Value("id").String().Raw() == colorID {
				obj.Value("group_id").Array().Length().Equal(2)
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("expected color id %s in /color/list?index=%d", colorID, index)
		}
	}

	resp := e.GET("/color/list").
		WithHeader("token", token).
		WithQuery("count", 61).
		WithQuery("index", 2).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	arr := resp.Value("data").Array()
	for i := 0; i < len(arr.Raw()); i++ {
		if arr.Element(i).Object().Value("id").String().Raw() == colorID {
			t.Fatalf("color id %s should not be in /color/list?index=2", colorID)
		}
	}
}

// 收藏：添加 -> 列表 -> 取消
func TestColorFavorite_Flow_WithDBPhoneTest(t *testing.T) {
	u, plainPwd := prepareTestUser(t)
	colorID := getAnyColorID(t)

	loadConfigOnce(t)
	db := getMysqlConn()
	// 清理该用户的历史收藏，避免顺序/重复影响断言
	_ = db.Where("user_id = ?", u.UserId).Delete(&ColorFavorite{}).Error

	app := newUserAndColorApp()
	e := httptest.New(t, app)

	login := e.POST("/user/login").
		WithJSON(iris.Map{
			"phone":  u.Phone,
			"passwd": plainPwd,
		}).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	token := login.Value("token").String().Raw()

	// 添加收藏（重复 id 只应新增 1 条）
	add := e.POST("/color/favorite").
		WithHeader("token", token).
		WithJSON(iris.Map{
			"id": []string{colorID, colorID},
		}).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	add.Value("result_code").Number().Equal(200)
	add.Value("added").Number().Equal(1)

	// 收藏列表应包含该颜色
	list := e.GET("/color/favorite/list").
		WithHeader("token", token).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	list.Value("result_code").Number().Equal(200)
	arr := list.Value("data").Array()
	arr.NotEmpty()

	// 最新收藏在前；且我们已经清空收藏，所以第一个应是本次添加的 colorID
	arr.Element(0).Object().Value("id").String().Equal(colorID)

	// 取消收藏
	rm := e.POST("/color/favorite/remove").
		WithHeader("token", token).
		WithJSON(iris.Map{
			"id": colorID,
		}).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	rm.Value("result_code").Number().Equal(200)
}

// /color/lut/max：需要 token 且返回 max_id
func TestLutMaxIdHandler_WithDBPhoneTest(t *testing.T) {
	u, plainPwd := prepareTestUser(t)

	app := newUserAndColorApp()
	e := httptest.New(t, app)

	login := e.POST("/user/login").
		WithJSON(iris.Map{
			"phone":  u.Phone,
			"passwd": plainPwd,
		}).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	token := login.Value("token").String().Raw()

	resp := e.GET("/color/lut/max").
		WithHeader("token", token).
		Expect().
		Status(iris.StatusOK).
		JSON().Object()

	resp.Value("result_code").Number().Equal(200)
	resp.Value("max_id").Number().Ge(0)
}
