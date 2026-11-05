package handler

import (
	"encoding/json"
	"strings"
	"time"

	"nail/config"

	log "github.com/jeanphorn/log4go"
	"github.com/kataras/iris/v12"
)

/*默认替换图（需存在于 OSS bucket 根目录）*/
func defaultReplaceImage() string {
	return config.GetOssDomain() + "/default.jpeg"
}

/*回调接口管理*/
func CallbackHandler(callback iris.Party) {
	/*阿里云 OSS 内容审核回调*/
	callback.Post("/oss", ossCallbackHandler)
}

/*
阿里云 OSS 内容审核回调

	Content-Type: application/x-www-form-urlencoded
	Form 参数: checksum, content
*/
func ossCallbackHandler(ctx iris.Context) {
	content := ctx.FormValue("content")
	log.Info("OSS callback received, content: %s", content)
	if content == "" {
		ctx.StatusCode(200)
		return
	}
	var callbackContent OssCallbackContent
	if err := json.Unmarshal([]byte(content), &callbackContent); err != nil {
		log.Error("OSS callback parse content error: %v", err)
		ctx.StatusCode(200)
		return
	}
	if callbackContent.Data == nil || len(callbackContent.Data.Results) == 0 {
		ctx.StatusCode(200)
		return
	}
	saveAuditLogs(&callbackContent)
	if isViolated(&callbackContent) {
		replaceViolatedContent(callbackContent.Data.OssObjectName)
	}
	ctx.StatusCode(200)
}

/*检查是否有违规（high/medium 均视为违规）*/
func isViolated(content *OssCallbackContent) bool {
	if content.Data.RiskLevel == "high" || content.Data.RiskLevel == "medium" {
		return true
	}
	for _, sr := range content.Data.Results {
		if sr.RiskLevel == "high" || sr.RiskLevel == "medium" {
			return true
		}
		for _, r := range sr.Result {
			if r.RiskLevel == "high" || r.RiskLevel == "medium" {
				return true
			}
		}
	}
	return false
}

/*替换违规内容为默认图片*/
func replaceViolatedContent(objectName string) {
	db := getMysqlConn()
	domain := config.GetOssDomain()
	fullUrl := domain + "/" + objectName

	/*圈子帖子图片/视频*/
	if strings.HasPrefix(objectName, "circle/") {
		var images []PostImage
		db.Where("image = ? OR thumbnail = ?", fullUrl, fullUrl).Find(&images)
		for _, img := range images {
			db.Model(&PostImage{}).Where("id = ?", img.Id).Updates(map[string]interface{}{
				"image":     defaultReplaceImage(),
				"thumbnail": defaultReplaceImage(),
			})
			log.Info("Replaced violated circle content, object: %s, image id: %s", objectName, img.Id)
		}
		deleteOssObject(objectName)
	}
}

/*保存审核日志*/
func saveAuditLogs(content *OssCallbackContent) {
	db := getMysqlConn()
	now := time.Now().Format("2006-01-02 15:04:05")
	scanResultBytes, _ := json.Marshal(content.Data)
	scanResultStr := string(scanResultBytes)
	for _, serviceResult := range content.Data.Results {
		for _, item := range serviceResult.Result {
			auditLog := AuditLog{
				Object:     content.Data.OssObjectName,
				Label:      item.Label,
				Service:    serviceResult.Service,
				RiskLevel:  item.RiskLevel,
				Bucket:     content.Data.OssBucketName,
				Region:     content.Data.OssRegionId,
				RequestId:  content.RequestId,
				ScanResult: scanResultStr,
				CreateTime: now,
			}
			if err := db.Create(&auditLog).Error; err != nil {
				log.Error("OSS callback save audit log error: %v, object: %s", err, content.Data.OssObjectName)
			}
		}
	}
}
