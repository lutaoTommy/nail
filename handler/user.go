package handler

import (
	"fmt"
	"strings"
	"time"

	"nail/language"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

/*用户管理*/
func UserHandler(user iris.Party) {
	/*注册获取邮箱验证码*/
	user.Post("/register/verification/mail", mailRegisterVerificationHandler)
	/*忘记获取邮箱验证码*/
	user.Post("/forget/verification/mail", mailForgetVerificationHandler)
	/*邮箱注册*/
	user.Post("/register/mail", mailRegisterHandler)
	/*邮箱忘记密码*/
	user.Post("/forget/mail", mailForgetHandler)
	/*创建用户*/
	user.Post("/create", createUserHandler)
	/*网页登陆*/
	user.Post("/login", userLoginHandler)
	/*邮箱登陆*/
	user.Post("/login/mail", mailLoginHandler)
	/*人员信息*/
	user.Post("/info/update", updateUserInfoHandler)
	/*人员信息*/
	user.Get("/info", getUserInfoHandler)
	/*用户修改自己密码*/
	user.Post("/passwd/update", changePasswdHandler)
	/*销毁用户*/
	user.Post("/destroy", destroyUserHandler)
	/*用户列表*/
	user.Get("/list", userListHandler)
	/*关注用户*/
	user.Post("/follow", followUserHandler)
	/*取关用户*/
	user.Post("/follow/cancel", cancelFollowHandler)
	/*关注列表*/
	user.Get("/follow/list", userFollowListHandler)
}

/*注册获取邮箱验证码*/
func mailRegisterVerificationHandler(ctx iris.Context) {
	var user User
	var err error
	returnData := iris.Map{"result_code": 500}
	ip := GetClientIP(ctx)
	if err = ctx.ReadJSON(&user); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
	} else if err = user.checkMail(); err != nil {
	}
	if err != nil {
		returnData["result_code"] = getErrCode(err)
		returnData["result_msg"] = err.Error()
		ctx.JSON(returnData)
		return
	}
	// IP + 邮箱维度限流：同一 IP 15 分钟内最多 10 次，同一邮箱最多 5 次
	if ok, retrySec := AllowVerificationRequest(ip, user.Email); !ok {
		returnData["result_code"] = 429
		returnData["result_msg"] = language.GetRawMessage("E_VERIFICATION_LIMIT")
		if retrySec > 0 {
			returnData["retry_after"] = retrySec
		}
		ctx.JSON(returnData)
		return
	}
	RecordVerificationRequest(ip, user.Email)
	err = mailRegisterVerification(&user)
	if err == nil {
		returnData["result_code"] = 200
		returnData["result_msg"] = "success"
	} else {
		returnData["result_code"] = getErrCode(err)
		returnData["result_msg"] = err.Error()
	}
	ctx.JSON(returnData)
}

/*注册获取邮箱验证码*/
func mailRegisterVerification(user *User) error {
	code := randInt()
	db := getMysqlConn()
	var userInfo User
	err := db.Where("email = ?", user.Email).First(&userInfo).Error
	/*新注册用户*/
	if err != nil {
		user.Status = -1
		user.Cert = code
		user.CertTime = time.Now().Format("2006-01-02 15:04:05")
		user.UserId = RandStringBytes(5)
		err = db.Create(user).Error
		/*已存在该邮箱*/
	} else if userInfo.Status == 1 {
		return newError(400, "E_USER_EXIST")
		/*再次获取*/
	} else {
		now := time.Now().Add(-1 * time.Minute).Format("2006-01-02 15:04:05")
		if now < userInfo.CertTime {
			return newError(400, "E_VERIFICATION_QUICKLY")
		}
		userInfo.Cert = code
		userInfo.CertTime = time.Now().Format("2006-01-02 15:04:05")
		err = db.Updates(userInfo).Error
	}
	user.Cert = code
	return sendMail(user)
}

/*忘记密码获取邮箱验证码*/
func mailForgetVerificationHandler(ctx iris.Context) {
	var user User
	var err error
	returnData := iris.Map{"result_code": 500}
	ip := GetClientIP(ctx)
	if err = ctx.ReadJSON(&user); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
	} else if err = user.checkMail(); err != nil {
	}
	if err != nil {
		returnData["result_code"] = getErrCode(err)
		returnData["result_msg"] = err.Error()
		ctx.JSON(returnData)
		return
	}
	if ok, retrySec := AllowVerificationRequest(ip, user.Email); !ok {
		returnData["result_code"] = 429
		returnData["result_msg"] = language.GetRawMessage("E_VERIFICATION_LIMIT")
		if retrySec > 0 {
			returnData["retry_after"] = retrySec
		}
		ctx.JSON(returnData)
		return
	}
	RecordVerificationRequest(ip, user.Email)
	err = mailForgetVerificatio(&user)
	if err == nil {
		returnData["result_code"] = 200
		returnData["result_msg"] = "success"
	} else {
		returnData["result_code"] = getErrCode(err)
		returnData["result_msg"] = err.Error()
	}
	ctx.JSON(returnData)
}

/*忘记密码获取邮箱验证码*/
func mailForgetVerificatio(user *User) error {
	code := randInt()
	db := getMysqlConn()
	var userInfo User
	err := db.Where("email = ?", user.Email).First(&userInfo).Error
	if err != nil {
		return newError(400, "E_INVALID_USER")
		/*已存在该邮箱*/
	} else if userInfo.Status == -1 {
		return newError(400, "E_INVALID_USER")
		/*再次获取*/
	} else {
		now := time.Now().Add(-1 * time.Minute).Format("2006-01-02 15:04:05")
		if now < userInfo.CertTime {
			return newError(400, "E_VERIFICATION_QUICKLY")
		}
		userInfo.Cert = code
		userInfo.CertTime = time.Now().Format("2006-01-02 15:04:05")
		err = db.Updates(userInfo).Error
	}
	user.Cert = code
	return sendMail(user)
}

/*邮箱注册*/
func mailRegisterHandler(ctx iris.Context) {
	var user User
	var err error
	returnData := iris.Map{"result_code": 500}
	if err = ctx.ReadJSON(&user); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
	} else if err = user.checkMail(); err != nil {
	} else if err = user.checkPasswd(); err != nil {
	} else if err = user.checkCert(); err != nil {
	}
	if err != nil {
		returnData["result_code"] = getErrCode(err)
		returnData["result_msg"] = err.Error()
	} else {
		err = mailRegister(&user)
		if err == nil {
			returnData["result_code"] = 200
			returnData["result_msg"] = "success"
			returnData["token"] = user.Token
		} else {
			returnData["result_code"] = getErrCode(err)
			returnData["result_msg"] = err.Error()
		}
	}
	ctx.JSON(returnData)
}

/*邮箱注册*/
func mailRegister(user *User) error {
	db := getMysqlConn()
	var userInfo User
	err := db.Where("email = ?", user.Email).First(&userInfo).Error
	if err != nil {
		return newError(400, "E_CERT_FIRST")
		/*已存在该邮箱*/
	} else if userInfo.Status == 1 {
		return newError(400, "E_USER_EXIST")
	} else if userInfo.Cert != user.Cert {
		return newError(400, "E_CERT_ERR")
	}
	hashed, err := HashPassword(user.Passwd)
	if err != nil {
		return err
	}
	tk, err := newToken()
	if err != nil {
		return err
	}
	user.Status = 1
	user.UserId = userInfo.UserId
	user.CertTime = userInfo.CertTime
	user.Token = tk
	user.Passwd = hashed
	user.RegisterTime = time.Now().Format("2006-01-02 15:04:05")
	return db.Updates(user).Error
}

/*邮箱忘记密码*/
func mailForgetHandler(ctx iris.Context) {
	var user User
	var err error
	returnData := iris.Map{"result_code": 500}
	if err = ctx.ReadJSON(&user); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
	} else if err = user.checkMail(); err != nil {
	} else if err = user.checkPasswd(); err != nil {
	} else if err = user.checkCert(); err != nil {
	}
	if err != nil {
		returnData["result_code"] = getErrCode(err)
		returnData["result_msg"] = err.Error()
	} else {
		err = mailForget(&user)
		if err == nil {
			returnData["result_code"] = 200
			returnData["result_msg"] = "success"
			returnData["token"] = user.Token
		} else {
			returnData["result_code"] = getErrCode(err)
			returnData["result_msg"] = err.Error()
		}
	}
	ctx.JSON(returnData)
}

/*邮箱忘记密码*/
func mailForget(user *User) error {
	db := getMysqlConn()
	var userInfo User
	err := db.Where("email = ?", user.Email).First(&userInfo).Error
	if err != nil {
		return newError(400, "E_INVALID_USER")
		/*已存在该邮箱*/
	} else if userInfo.Status == -1 {
		return newError(400, "E_INVALID_USER")
	} else if userInfo.Cert != user.Cert {
		return newError(400, "E_CERT_ERR")
	}
	hashed, err := HashPassword(user.Passwd)
	if err != nil {
		return err
	}
	tk, err := newToken()
	if err != nil {
		return err
	}
	// 找回密码成功后轮换 token，避免旧 token 长期可用
	if err := db.Model(&User{}).Where("user_id = ?", userInfo.UserId).Updates(map[string]interface{}{
		"passwd": hashed,
		"token":  tk,
	}).Error; err != nil {
		return err
	}
	user.Token = tk
	return nil
}

/*用户创建(内部用)*/
func createUserHandler(ctx iris.Context) {
	var err error
	var user User
	if err = ctx.ReadJSON(&user); err != nil {
	} else if err = user.checkPhone(); err != nil {
	} else if err = user.checkPasswd(); err != nil {
	}
	// token 只允许从 header 获取，避免 body 覆盖鉴权 token
	user.Token = ctx.GetHeader("token")
	if err == nil {
		err = user.checkToken()
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	err = createUser(&user)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "token": user.Token})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*用户创建(内部用)*/
func createUser(user *User) error {
	db := getMysqlConn()
	var userInfo User
	err := db.Where("token = ?", user.Token).First(&userInfo).Error
	if err != nil {
		return newError(401, "E_NO_TOKEN")
	} else if userInfo.UserId != ADMIN {
		return newError(403, "E_NO_AUTH")
	}
	var tempUser User
	err = db.Where("phone = ?", user.Phone).First(&tempUser).Error
	if err == nil {
		return newError(400, "E_USER_EXIST")
	}
	tk, err := newToken()
	if err != nil {
		return err
	}
	hashed, err := HashPassword(user.Passwd)
	if err != nil {
		return err
	}
	user.Token = tk
	user.UserId = RandStringBytes(5)
	user.Passwd = hashed
	return db.Create(user).Error
}

/*网页登录*/
func userLoginHandler(ctx iris.Context) {
	var err error
	var user User
	user.Token = ctx.GetHeader("token")
	if err = ctx.ReadJSON(&user); err != nil {
	} else if err = user.checkPasswd(); err != nil {
	} else if user.Phone == "" {
		err = newError(400, "E_NO_PHONE")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	ip := GetClientIP(ctx)
	accKey := "phone:" + strings.TrimSpace(user.Phone)
	if ok, retrySec := AllowLogin(ip, accKey); !ok {
		res := iris.Map{"result_code": 429, "result_msg": language.GetRawMessage("E_ACCOUNT_LOCKED")}
		if retrySec > 0 {
			res["retry_after"] = retrySec
		}
		ctx.JSON(res)
		return
	}
	err = userLogin(&user)
	if err != nil {
		RecordLoginFailure(ip, accKey)
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	RecordLoginSuccess(ip, accKey)
	ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "token": user.Token})
}

/*网页登录*/
func userLogin(user *User) error {
	db := getMysqlConn()
	var userInfo User
	err := db.Where("phone = ?", user.Phone).First(&userInfo).Error
	if err != nil {
		return newError(400, "E_INVALID_USER")
	} else if userInfo.Status == -1 {
		return newError(400, "E_INVALID_USER")
	}
	ok, err := VerifyPassword(userInfo.Passwd, user.Passwd)
	if err != nil {
		return err
	}
	if !ok {
		return newError(400, "E_INVALID_PWD")
	}
	user.Token = userInfo.Token
	return nil
}

/*邮箱登录*/
func mailLoginHandler(ctx iris.Context) {
	var err error
	var user User
	user.Token = ctx.GetHeader("token")
	if err = ctx.ReadJSON(&user); err != nil {
	} else if err = user.checkPasswd(); err != nil {
	} else if err = user.checkMail(); err != nil {
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	ip := GetClientIP(ctx)
	accKey := "email:" + strings.TrimSpace(strings.ToLower(user.Email))
	if ok, retrySec := AllowLogin(ip, accKey); !ok {
		res := iris.Map{"result_code": 429, "result_msg": language.GetRawMessage("E_ACCOUNT_LOCKED")}
		if retrySec > 0 {
			res["retry_after"] = retrySec
		}
		ctx.JSON(res)
		return
	}
	err = mailLogin(&user)
	if err != nil {
		RecordLoginFailure(ip, accKey)
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	RecordLoginSuccess(ip, accKey)
	ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "token": user.Token})
}

/*邮箱登录*/
func mailLogin(user *User) error {
	db := getMysqlConn()
	var userInfo User
	err := db.Where("email = ?", user.Email).First(&userInfo).Error
	if err != nil {
		return newError(400, "E_INVALID_USER")
	} else if userInfo.Status == -1 {
		return newError(400, "E_INVALID_USER")
	}
	ok, err := VerifyPassword(userInfo.Passwd, user.Passwd)
	if err != nil {
		return err
	}
	if !ok {
		return newError(400, "E_INVALID_PWD")
	}
	// 邮箱登录成功后，生成新 token，避免旧 token 继续长期可用
	newTk, err := newToken()
	if err != nil {
		return err
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	if err := db.Model(&User{}).Where("user_id = ?", userInfo.UserId).Updates(map[string]interface{}{
		"token":      newTk,
		"login_time": now,
	}).Error; err != nil {
		return err
	}
	user.Token = newTk
	return nil
}

/*更新用户信息*/
func updateUserInfoHandler(ctx iris.Context) {
	var err error
	var user User
	if err = ctx.ReadJSON(&user); err != nil {
	}
	// token 只允许从 header 获取，避免 body 覆盖鉴权 token
	user.Token = ctx.GetHeader("token")
	if err == nil {
		err = user.checkToken()
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	err = updateUserInfo(&user)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success"})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*更新用户信息*/
func updateUserInfo(user *User) error {
	db := getMysqlConn()
	var userInfo User
	err := db.Where("token = ?", user.Token).First(&userInfo).Error
	if err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	data := make(map[string]interface{})
	if user.Nickname != "" {
		if len(user.Nickname) > 100 {
			return newError(400, "E_TOO_LONG")
		}
		data["nickname"] = user.Nickname
	}
	if user.Language != "" {
		if len(user.Language) > 5 {
			return newError(400, "E_TOO_LONG")
		}
		data["language"] = user.Language
	}
	if len(user.Biography) > 200 {
		return newError(400, "E_TOO_LONG")
	}
	data["biography"] = user.Biography
	return db.Model(&User{}).Where("user_id = ?", userInfo.UserId).Updates(data).Error
}

/*用户修改自己密码*/
func changePasswdHandler(ctx iris.Context) {
	var err error
	var req ChangePasswdReq
	token := ctx.GetHeader("token")
	if err = ctx.ReadJSON(&req); err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	if token == "" {
		err = newError(401, "E_NO_TOKEN")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	newToken, err := changePasswd(token, &req)
	if err == nil {
		// 兼容：额外返回 token 字段，不影响老客户端（忽略即可）
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "token": newToken})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*用户修改自己密码*/
func changePasswd(token string, req *ChangePasswdReq) (string, error) {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", token).First(&userInfo).Error; err != nil {
		return "", newError(401, "E_NO_TOKEN")
	}
	if req.OldPasswd == "" || req.NewPasswd == "" {
		return "", newError(400, "E_NO_PWD")
	}
	if len(req.NewPasswd) > 72 {
		return "", newError(400, "E_TOO_LONG")
	}
	ok, err := VerifyPassword(userInfo.Passwd, req.OldPasswd)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", newError(400, "E_INVALID_PWD")
	}
	hashed, err := HashPassword(req.NewPasswd)
	if err != nil {
		return "", err
	}
	newTk, err := newToken()
	if err != nil {
		return "", err
	}
	// 改密成功后轮换 token，避免旧 token 继续可用
	if err := db.Model(&User{}).Where("user_id = ?", userInfo.UserId).Updates(map[string]interface{}{
		"passwd": hashed,
		"token":  newTk,
	}).Error; err != nil {
		return "", err
	}
	return newTk, nil
}

/*查询用户信息*/
func getUserInfoHandler(ctx iris.Context) {
	var err error
	var user User
	user.Token = ctx.GetHeader("token")
	if user.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	err = getUserInfo(&user)
	if err == nil {
		ctx.JSON(iris.Map{
			"result_code":  200,
			"result_msg":   "success",
			"phone":        user.Phone,
			"email":        user.Email,
			"user_id":      user.UserId,
			"avatar":       user.Avatar,
			"nickname":     user.Nickname,
			"language":     user.Language,
			"biography":    user.Biography,
			"follow_count": user.FollowCount,
			"fans_count":   user.FansCount,
			"post_count":   user.PostCount,
		})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*查询用户信息*/
func getUserInfo(user *User) error {
	db := getMysqlConn()
	err := db.Where("token = ?", user.Token).First(user).Error
	if err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	return nil
}

/*销毁账号*/
func destroyUserHandler(ctx iris.Context) {
	var err error
	type DestroyUserReq struct {
		DeleteInfo bool `json:"delete_info"`
	}
	var req DestroyUserReq
	_ = ctx.ReadJSON(&req)
	var user UserDestory
	user.Token = ctx.GetHeader("token")
	if err = user.checkToken(); err != nil {
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	err = destroyUser(&user, req.DeleteInfo)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success"})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*销毁账号*/
func destroyUser(user *UserDestory, deleteInfo bool) error {
	db := getMysqlConn()
	var userInfo User
	err := db.Where("token = ?", user.Token).First(&userInfo).Error
	if err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	return db.Transaction(func(tx *gorm.DB) error {
		if deleteInfo {
			var postInfoIDs []string
			if err := tx.Model(&PostInfo{}).Select("id").Where("user_id = ?", userInfo.UserId).Find(&postInfoIDs).Error; err != nil {
				return err
			}
			if len(postInfoIDs) > 0 {
				if err := tx.Where("post_id IN ?", postInfoIDs).Delete(&PostColor{}).Error; err != nil {
					return err
				}
			}
			if err := tx.Where("user_id = ?", userInfo.UserId).Delete(&PostInfo{}).Error; err != nil {
				return err
			}
			if err := tx.Where("user_id = ?", userInfo.UserId).Delete(&ColorFavorite{}).Error; err != nil {
				return err
			}
		}
		if err := tx.Model(&User{}).Where("user_id = ?", userInfo.UserId).Update("token", "").Error; err != nil {
			return err
		}
		return tx.Delete(&userInfo).Error
	})
}

/*查询用户列表*/
func userListHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Phone = ctx.URLParam("phone")
	params.Email = ctx.URLParam("email")
	params.Token = ctx.GetHeader("token")
	params.Name = ctx.URLParam("nickname")
	params.Page = AtoUI(ctx.URLParam("page"), 1)
	params.Limit = AtoUI(ctx.URLParam("limit"), 10)
	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	data, err := userList(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "total": params.Total, "data": data})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*查询用户列表*/
func userList(params *Params) ([]UserOut, error) {
	db := getMysqlConn()
	var userInfo User
	err := db.Where("token = ?", params.Token).First(&userInfo).Error
	if err != nil {
		return nil, newError(401, "E_NO_TOKEN")
	} else if userInfo.UserId != ADMIN {
		return nil, newError(403, "E_NO_AUTH")
	}
	db = db.Table("users")
	db = db.Where("user_id != ?", ADMIN)
	if params.Phone != "" {
		db = db.Where("phone like ?", fmt.Sprintf("%%%s%%", params.Phone))
	}
	if params.Email != "" {
		db = db.Where("email like ?", fmt.Sprintf("%%%s%%", params.Email))
	}
	if params.Name != "" {
		db = db.Where("nickname like ?", fmt.Sprintf("%%%s%%", params.Name))
	}
	err = db.Count(&params.Total).Error
	if err != nil {
		return nil, err
	}
	data := []UserOut{}
	db = db.Offset((params.Page - 1) * params.Limit).Limit(params.Limit)
	err = db.Order("register_time desc").Find(&data).Error
	return data, err
}
