package handler

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

/*建议管理*/
func SuggestHandler(suggest iris.Party) {
	/*上传建议*/
	suggest.Post("/upload", uploadSuggestHandler)
	/*上传建议*/
	suggest.Post("/image/upload", uploadSuggestImageHandler)
	/*查询建议*/
	suggest.Get("/list", listSuggestHandler)
	/*查询建议*/
	suggest.Get("/detail", suggestDetailHandler)
	/*删除意见*/
	suggest.Post("/remove", removeSuggestHandler)
}

/*上传建议*/
func uploadSuggestHandler(ctx iris.Context) {
	var suggest Suggest
	suggest.Token = ctx.GetHeader("token")
	returnData := iris.Map{"result_code": 500}
	err := ctx.ReadJSON(&suggest)
	if err != nil {
	} else if suggest.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} else if err = suggest.checkContent(); err != nil {
	}
	if err != nil {
		returnData["result_code"] = getErrCode(err)
		returnData["result_msg"] = err.Error()
	} else {
		err = uploadSuggest(&suggest)
		if err == nil {
			returnData["result_code"] = 200
			returnData["result_msg"] = "success"
			returnData["id"] = suggest.Id
		} else {
			returnData["result_code"] = getErrCode(err)
			returnData["result_msg"] = err.Error()
		}
	}
	ctx.JSON(returnData)
}

/*上传建议逻辑*/
func uploadSuggest(suggest *Suggest) error {
	db := getMysqlConn()
	var userInfo User
	err := db.Where("token = ?", suggest.Token).First(&userInfo).Error
	if err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	suggest.Id = RandStringBytes(3)
	suggest.Avatar = userInfo.AvatarObjectKey
	suggest.Phone = userInfo.Phone
	suggest.Nickname = userInfo.Nickname
	suggest.UserId = userInfo.UserId
	suggest.Time = time.Now().Format("2006-01-02 15:04:05")
	return db.Create(suggest).Error
}

/*上传建议图片*/
func uploadSuggestImageHandler(ctx iris.Context) {
	var picture SuggestImage
	picture.Suggest_id = ctx.URLParam("id")
	picture.Token = ctx.GetHeader("token")
	returnData := iris.Map{"result_code": 500}

	var err error
	if picture.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} else if picture.Suggest_id == "" {
		err = newError(400, "E_NO_ID")
	}

	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}

	// 复用 avatar.go 中的 helper
	fileBytes, format, cfg, size, err := readAndValidateImage(ctx)
	if err != nil {
		returnData["result_code"] = getErrCode(err)
		returnData["result_msg"] = err.Error()
		ctx.JSON(returnData)
		return
	}

	picture.Width = cfg.Width
	picture.Height = cfg.Height
	picture.Size = strconv.FormatInt(size, 10)
	picture.Type = format

	err = uploadSuggestImage(fileBytes, &picture)
	if err == nil {
		returnData["result_code"] = 200
		returnData["result_msg"] = "success"
	} else {
		returnData["result_code"] = getErrCode(err)
		returnData["result_msg"] = err.Error()
	}
	ctx.JSON(returnData)
}

/*上传建议图片逻辑*/
func uploadSuggestImage(imgData []byte, picture *SuggestImage) error {
	db := getMysqlConn()
	var userInfo User
	err := db.Where("token = ?", picture.Token).First(&userInfo).Error
	if err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	var suggest Suggest
	err = db.Where("id = ?", picture.Suggest_id).First(&suggest).Error
	if err != nil {
		return newError(404, "E_NO_SUGGEST")
	}

	/*解析图片*/
	img, err := parseImageData(bytes.NewBuffer(imgData))
	if err != nil {
		return err
	}

	/*保存图片 (统一转为 PNG)*/
	picture.Id = RandStringBytes(3)
	name := picture.Id + ".png"

	// 确保目录存在
	dir := "public/suggest"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		_ = os.MkdirAll(dir, 0755)
	}
	path := filepath.Join(dir, name)

	// 使用更可靠的保存方式（或复用 savePngImage，但 savePngImage 本身比较简单）
	err = savePngImage(img, path)
	if err != nil {
		return err
	}

	picture.UserId = userInfo.UserId
	picture.Time = time.Now().Format("2006-01-02 15:04:05")
	return db.Create(picture).Error
}

/*查询建议列表*/
func listSuggestHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	params.Page = AtoUI(ctx.URLParam("page"), 1)
	params.Limit = AtoUI(ctx.URLParam("limit"), 10)
	if params.Limit > 100 {
		params.Limit = 100
	}

	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	data, err := listSuggest(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "total": params.Total, "data": data})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*查询建议列表逻辑*/
func listSuggest(params *Params) ([]SuggestOut, error) {
	db := getMysqlConn()
	var userInfo User
	err := db.Where("token = ?", params.Token).First(&userInfo).Error
	if err != nil {
		return nil, newError(401, "E_NO_TOKEN")
	}
	db = db.Table("suggests")
	if userInfo.UserId != ADMIN {
		db = db.Where("user_id = ?", userInfo.UserId)
	}
	err = db.Count(&params.Total).Error
	if err != nil {
		return nil, err
	}
	data := []SuggestOut{}
	db = db.Offset((params.Page - 1) * params.Limit).Limit(params.Limit)
	if err := db.Order("time desc").Find(&data).Error; err != nil {
		return nil, err
	}
	for i := range data {
		if data[i].Avatar == "" {
			continue
		}
		if u, err := signAvatarURL(data[i].Avatar); err == nil {
			data[i].Avatar = u
		}
	}
	return data, nil
}

/*建议详情*/
func suggestDetailHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	params.Id = ctx.URLParam("id")
	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	data, err := suggestDetail(ctx, &params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "data": data})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*建议详情逻辑*/
func suggestDetail(ctx iris.Context, params *Params) (SuggestDetail, error) {
	db := getMysqlConn()
	var userInfo User
	var detail SuggestDetail
	err := db.Where("token = ?", params.Token).First(&userInfo).Error
	if err != nil {
		return detail, newError(401, "E_NO_TOKEN")
	}

	var suggest Suggest
	err = db.Where("id = ?", params.Id).First(&suggest).Error
	if err != nil {
		return detail, newError(404, "E_NO_SUGGEST")
	}

	// 鉴权：如果是普通用户，只能看自己的建议（可选，保持现状或增强安全）
	if userInfo.UserId != ADMIN && suggest.UserId != userInfo.UserId {
		return detail, newError(403, "E_NO_AUTH")
	}

	detail.Id = suggest.Id
	detail.Phone = suggest.Phone
	detail.Nickname = suggest.Nickname
	if suggest.Avatar != "" {
		if u, err := signAvatarURL(suggest.Avatar); err == nil {
			detail.Avatar = u
		}
	}
	detail.Time = suggest.Time
	detail.Content = suggest.Content

	/*图片*/
	detail.Attachs = []Attach{}
	err = db.Table("suggest_images").Where("suggest_id = ?", suggest.Id).Find(&detail.Attachs).Error
	if err != nil {
		return detail, err
	}

	for i := range detail.Attachs {
		detail.Attachs[i].Url = BuildPublicURL(ctx, "/suggest/"+detail.Attachs[i].Id+".png")
	}
	return detail, nil
}

/*建议删除*/
func removeSuggestHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	_ = ctx.ReadJSON(&params)

	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} else if params.Id == "" {
		err = newError(400, "E_NO_ID")
	}

	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}

	err = removeSuggest(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success"})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*建议删除逻辑*/
func removeSuggest(params *Params) error {
	db := getMysqlConn()
	var userInfo User
	err := db.Where("token = ?", params.Token).First(&userInfo).Error
	if err != nil {
		return newError(401, "E_NO_TOKEN")
	}

	// 开启事务
	return db.Transaction(func(tx *gorm.DB) error {
		var suggest Suggest
		err := tx.Where("id = ?", params.Id).First(&suggest).Error
		if err != nil {
			return newError(404, "E_NO_SUGGEST")
		} else if userInfo.UserId != suggest.UserId && userInfo.UserId != ADMIN {
			return newError(403, "E_NO_AUTH")
		}

		// 1. 查找相关图片
		var pictures []SuggestImage
		if err := tx.Where("suggest_id = ?", suggest.Id).Find(&pictures).Error; err != nil {
			return err
		}

		// 2. 删除建议记录
		if err := tx.Delete(&suggest).Error; err != nil {
			return err
		}

		// 3. 删除图片记录
		if len(pictures) > 0 {
			if err := tx.Where("suggest_id = ?", suggest.Id).Delete(&SuggestImage{}).Error; err != nil {
				return err
			}
		}

		// 4. 删除物理文件 (事务提交后执行不可回滚，但通常为了保持一致性放在这里也可以，
		for _, picture := range pictures {
			file := fmt.Sprintf("public/suggest/%s.png", picture.Id)
			if isExist(file) {
				os.Remove(file)
			}
		}

		return nil
	})
}
