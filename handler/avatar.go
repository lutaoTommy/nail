package handler

import (
	"bytes"
	"image"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
)

/*头像管理*/
func AvatarHandler(avatar iris.Party) {
	/*上传头像*/
	avatar.Post("/upload", uploadAvatarHandler)
	/*用户上传头像（存 OSS，更新当前用户 avatar）*/
	avatar.Post("/user/upload", userUploadAvatarHandler)
	/*查询头像*/
	avatar.Get("/list", listAvatarHandler)
	/*删除头像*/
	avatar.Post("/remove", removeAvatarHandler)
	/*切换头像*/
	avatar.Post("/change", changeAvatarHandler)
}

/*辅助函数：读取并验证图片文件*/
func readAndValidateImage(ctx iris.Context) ([]byte, string, *image.Config, int64, error) {
	file, info, err := ctx.FormFile("image")
	if err != nil {
		return nil, "", nil, 0, newError(400, "E_INVALID_PICTURE")
	}
	defer file.Close()

	// 读取配置以获取宽高和格式
	cfg, format, err := image.DecodeConfig(file)
	if err != nil {
		return nil, "", nil, 0, newError(400, "E_INVALID_PICTURE")
	}
	if format != "jpeg" && format != "png" {
		return nil, "", nil, 0, newError(400, "E_INVALID_PICTURE")
	}

	// 重置文件指针并读取所有内容
	if _, err := file.Seek(0, 0); err != nil {
		return nil, "", nil, 0, err
	}
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, "", nil, 0, err
	}

	// 再次检查 Content-Type
	cType := http.DetectContentType(fileBytes)
	if cType != "image/jpeg" && cType != "image/png" {
		return nil, "", nil, 0, newError(400, "E_INVALID_PICTURE")
	}

	return fileBytes, format, &cfg, info.Size, nil
}

/*上传头像（管理员）*/
func uploadAvatarHandler(ctx iris.Context) {
	var avatar Avatar
	avatar.Token = ctx.GetHeader("token")
	if avatar.Token == "" {
		err := newError(401, "E_NO_TOKEN")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}

	fileBytes, format, cfg, size, err := readAndValidateImage(ctx)
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}

	avatar.Width = cfg.Width
	avatar.Height = cfg.Height
	avatar.Size = strconv.FormatInt(size, 10)
	avatar.Type = format

	err = uploadAvatar(fileBytes, &avatar)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success"})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*管理员上传头像逻辑*/
func uploadAvatar(imgData []byte, avatar *Avatar) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", avatar.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	} else if userInfo.UserId != ADMIN {
		return newError(403, "E_NO_AUTH")
	}

	/*解析图片 (校验是否有效图片)*/
	img, err := parseImageData(bytes.NewBuffer(imgData))
	if err != nil {
		return err
	}

	/*转为 PNG 统一存储*/
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return err
	}

	avatar.Id = RandStringBytes(3)
	name := avatar.Id + ".png"
	ossPath := "avatar/" + name
	if err := putOssObjectFromBytes(ossPath, buf.Bytes()); err != nil {
		return err
	}

	avatar.ObjectKey = ossPath
	avatar.Time = time.Now().Format("2006-01-02 15:04:05")
	return db.Create(avatar).Error
}

/*用户上传头像（存 OSS，并更新当前用户 avatar）*/
func userUploadAvatarHandler(ctx iris.Context) {
	token := ctx.GetHeader("token")
	if token == "" {
		err := newError(401, "E_NO_TOKEN")
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}

	fileBytes, _, _, _, err := readAndValidateImage(ctx)
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}

	avatarURL, err := userUploadAvatar(token, fileBytes)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "avatar": avatarURL})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*用户上传头像逻辑*/
func userUploadAvatar(token string, fileBytes []byte) (string, error) {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", token).First(&userInfo).Error; err != nil {
		return "", newError(401, "E_NO_TOKEN")
	}

	img, err := parseImageData(bytes.NewBuffer(fileBytes))
	if err != nil {
		return "", err
	}

	var pngBuf bytes.Buffer
	if err := png.Encode(&pngBuf, img); err != nil {
		return "", err
	}

	name := userInfo.UserId + "_" + RandStringBytes(3) + ".png"
	ossPath := "user/avatar/" + name
	if err := putOssObjectFromBytes(ossPath, pngBuf.Bytes()); err != nil {
		return "", err
	}

	// 只更新 avatar_object_key 字段
	if err := db.Model(&User{}).Where("user_id = ?", userInfo.UserId).Update("avatar_object_key", ossPath).Error; err != nil {
		return "", err
	}
	return signAvatarURL(ossPath)
}

/*查询头像列表*/
func listAvatarHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	params.Page = AtoUI(ctx.URLParam("page"), 1)
	params.Limit = AtoUI(ctx.URLParam("limit"), 10)
	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	data, err := listAvatar(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "total": params.Total, "data": data}, iris.JSON{UnescapeHTML: true})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*查询头像列表逻辑*/
func listAvatar(params *Params) ([]AvatarOut, error) {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return nil, newError(401, "E_NO_TOKEN")
	}

	db = db.Model(&Avatar{})
	if err := db.Count(&params.Total).Error; err != nil {
		return nil, err
	}

	var list []Avatar
	db = db.Offset((params.Page - 1) * params.Limit).Limit(params.Limit)
	if err := db.Order("time desc").Find(&list).Error; err != nil {
		return nil, err
	}
	out := make([]AvatarOut, 0, len(list))
	for _, a := range list {
		url, err := signAvatarURL(a.ObjectKey)
		if err != nil {
			return nil, err
		}
		out = append(out, AvatarOut{Id: a.Id, Url: url, Time: a.Time})
	}
	return out, nil
}

/*删除头像*/
func removeAvatarHandler(ctx iris.Context) {
	var params Params
	params.Token = ctx.GetHeader("token")

	// 尝试解析 JSON，忽略错误继续校验字段
	_ = ctx.ReadJSON(&params)

	var err error
	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} else if params.Id == "" {
		err = newError(400, "E_NO_ID")
	}

	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}

	err = removeAvatar(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success"})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*删除头像逻辑*/
func removeAvatar(params *Params) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	} else if userInfo.UserId != ADMIN {
		return newError(403, "E_NO_AUTH")
	}

	var avatar Avatar
	if err := db.Where("id = ?", params.Id).First(&avatar).Error; err != nil {
		return newError(404, "E_NO_AVATAR")
	}

	return db.Delete(&avatar).Error
}

/*文件是否存在*/
func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

/*切换头像*/
func changeAvatarHandler(ctx iris.Context) {
	var params Params
	params.Token = ctx.GetHeader("token")
	_ = ctx.ReadJSON(&params)

	var err error
	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} else if params.Id == "" {
		err = newError(400, "E_NO_ID")
	}

	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}

	err = changeAvatar(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success"})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*切换头像逻辑*/
func changeAvatar(params *Params) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	}

	var avatar Avatar
	if err := db.Where("id = ?", params.Id).First(&avatar).Error; err != nil {
		return newError(404, "E_NO_AVATAR")
	}

	// 只更新 avatar_object_key 字段，避免覆盖其他字段
	return db.Model(&User{}).Where("user_id = ?", userInfo.UserId).Update("avatar_object_key", avatar.ObjectKey).Error
}
