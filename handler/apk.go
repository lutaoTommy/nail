package handler

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

/*APK 管理*/
func ApkHandler(apk iris.Party) {
	/*上传 APK*/
	apk.Post("/upload", uploadApkHandler)
	/*获取最新 APK 版本*/
	apk.Get("/latest", latestApkHandler)
}

type ApkVersion struct {
	Id      string `gorm:"primaryKey;column:id;size:20" json:"id"`
	Version string `gorm:"column:version;size:30" json:"version"`
	Name    string `gorm:"column:name;size:100" json:"name"`
	Time    string `gorm:"column:time;size:20" json:"time"`
}

func (ApkVersion) TableName() string {
	return "apk_versions"
}

/*上传 APK*/
func uploadApkHandler(ctx iris.Context) {
	var params Params
	params.Token = ctx.GetHeader("token")
	params.Name = ctx.URLParam("version")

	returnData := iris.Map{"result_code": 500}

	file, info, err := ctx.FormFile("file")
	if err != nil {
		err = newError(400, "E_INVALID_FILE")
	} else if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} else if params.Name == "" {
		err = newError(400, "E_NO_VERSION")
	}

	if err != nil {
		returnData["result_code"] = getErrCode(err)
		returnData["result_msg"] = err.Error()
		ctx.JSON(returnData)
		return
	}
	defer file.Close()

	params.Id = info.Filename
	err = uploadApk(file, &params)
	if err == nil {
		returnData["result_code"] = 200
		returnData["result_msg"] = "success"
	} else {
		returnData["result_code"] = getErrCode(err)
		returnData["result_msg"] = err.Error()
	}
	ctx.JSON(returnData)
}

/*上传 APK 逻辑*/
func uploadApk(file multipart.File, params *Params) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	} else if userInfo.UserId != ADMIN {
		return newError(403, "E_NO_AUTH")
	}

	cleanName := filepath.Base(params.Id)
	if cleanName == "" || cleanName == "." || cleanName == "/" {
		return newError(400, "E_INVALID_FILE")
	}

	dir := "public/apk"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if mkErr := os.MkdirAll(dir, 0755); mkErr != nil {
			return mkErr
		}
	}

	path := filepath.Join(dir, cleanName)
	dstFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	if _, err = io.Copy(dstFile, file); err != nil {
		return err
	}

	now := time.Now().Format("2006-01-02 15:04:05")

	apk := ApkVersion{
		Id:      "latest",
		Version: params.Name,
		Name:    cleanName,
		Time:    now,
	}

	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", apk.Id).Delete(&ApkVersion{}).Error; err != nil {
			return err
		}
		return tx.Create(&apk).Error
	})
}

/*获取最新 APK 版本*/
func latestApkHandler(ctx iris.Context) {
	var params Params
	params.Token = ctx.GetHeader("token")
	if params.Token == "" {
		ctx.JSON(iris.Map{"result_code": 401, "result_msg": "E_NO_TOKEN"})
		return
	}

	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		ctx.JSON(iris.Map{"result_code": 401, "result_msg": "E_NO_TOKEN"})
		return
	}

	var apk ApkVersion
	if err := db.Where("id = ?", "latest").First(&apk).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(iris.Map{"result_code": 404, "result_msg": "E_NO_APK"})
		} else {
			ctx.JSON(iris.Map{"result_code": 500, "result_msg": err.Error()})
		}
		return
	}

	url := BuildPublicURL(ctx, "/apk/"+apk.Name)

	ctx.JSON(iris.Map{
		"result_code": 200,
		"result_msg":  "success",
		"data": iris.Map{
			"version": apk.Version,
			"url":     url,
			"time":    apk.Time,
		},
	})
}

