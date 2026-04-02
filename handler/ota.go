package handler

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/kataras/iris/v12"
)

/*OTA管理*/
func OtaHandler(ota iris.Party) {
	/*上传hex (V1)*/
	ota.Post("/upload", func(ctx iris.Context) { handleUploadOta(ctx, "1") })
	/*上传hex (V2)*/
	ota.Post("/upload/v2", func(ctx iris.Context) { handleUploadOta(ctx, "2") })

	/*OTA版本 (V1)*/
	ota.Get("/version", func(ctx iris.Context) { handleOtaVersion(ctx, "1") })
	/*OTA版本 (V2)*/
	ota.Get("/version/v2", func(ctx iris.Context) { handleOtaVersion(ctx, "2") })
}

/*通用上传处理*/
func handleUploadOta(ctx iris.Context, versionID string) {
	var params Params
	params.Token = ctx.GetHeader("token")
	// 不同的接口可能传递不同的参数名？原代码都叫 "code"，都用 URLParam 获取
	params.Index = AtoUI(ctx.URLParam("code"), 0)

	returnData := iris.Map{"result_code": 500}

	// 统一逻辑检查
	file, info, err := ctx.FormFile("image") // 参数名虽然叫 image，实际是 hex 文件
	if err != nil {
		err = newError(400, "E_INVALID_FILE")
	} else if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} else if params.Index <= 0 {
		err = newError(400, "E_NO_VERSION")
	}

	if err != nil {
		returnData["result_code"] = getErrCode(err)
		returnData["result_msg"] = err.Error()
	} else {
		defer file.Close()
		params.Name = info.Filename
		err = uploadOta(file, &params, versionID)
		if err == nil {
			returnData["result_code"] = 200
			returnData["result_msg"] = "success"
		} else {
			returnData["result_code"] = getErrCode(err)
			returnData["result_msg"] = err.Error()
		}
	}
	ctx.JSON(returnData)
}

/*通用上传逻辑*/
func uploadOta(file multipart.File, params *Params, versionID string) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	} else if userInfo.UserId != ADMIN {
		return newError(403, "E_NO_AUTH")
	}

	// 简单清理文件名，防止路径穿越
	cleanName := filepath.Base(params.Name)
	// 如果文件名为空或只有点，处理一下安全隐患
	if cleanName == "" || cleanName == "." || cleanName == "/" {
		return newError(400, "E_INVALID_FILE")
	}

	// 确保目录存在
	dir := "public/ota"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		_ = os.MkdirAll(dir, 0755)
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

	var ota OtaVersion
	ota.Id = versionID
	ota.Name = cleanName
	ota.Code = params.Index
	ota.Time = time.Now().Format("2006-01-02 15:04:05")

	return db.Save(&ota).Error
}

/*通用版本查询处理*/
func handleOtaVersion(ctx iris.Context, versionID string) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	}

	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}

	data, err := otaVersion(&params, versionID)
	if err == nil {
		data.Url = BuildPublicURL(ctx, "/ota/"+data.Name)
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "data": data})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*通用版本查询逻辑*/
func otaVersion(params *Params, versionID string) (OtaOut, error) {
	var data OtaOut
	db := getMysqlConn()

	// 鉴权
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return data, newError(401, "E_NO_TOKEN")
	}

	// 查询对应的版本 ID；URL 由 handler 层用 BuildPublicURL 设置
	err := db.Table("ota_versons").Where("id = ?", versionID).First(&data).Error
	return data, err
}
