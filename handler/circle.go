package handler

import (
	"bytes"
	"image"
	"image/png"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"nail/config"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

/*圈子管理*/
func CircleHandler(circle iris.Party) {
	/*发布动态*/
	circle.Post("/post", postCircleHandler)
	/*上传图片*/
	circle.Post("/image/upload", uploadCircleImageHandler)
	/*批量上传图片*/
	circle.Post("/image/batch/upload", batchUploadCircleImageHandler)
	/*上传视频*/
	circle.Post("/video/upload", uploadCircleVideoHandler)
	/*动态列表*/
	circle.Get("/list", listCirclePostHandler)
	/*删除动态*/
	circle.Post("/remove", removeCircleHandler)
	/*标记审核动态*/
	circle.Post("/review", reviewCircleHandler)
}

/*发布动态*/
func postCircleHandler(ctx iris.Context) {
	var circlePost CirclePost
	circlePost.Token = ctx.GetHeader("token")
	returnData := iris.Map{"result_code": 500}
	err := ctx.ReadJSON(&circlePost)
	if err != nil {
	} else if circlePost.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} else if err = circlePost.checkContent(); err != nil {
	}
	if err != nil {
		returnData["result_code"] = getErrCode(err)
		returnData["result_msg"] = ErrMsg(ctx, err)
	} else {
		err = postCircle(&circlePost)
		if err == nil {
			returnData["result_code"] = 200
			returnData["result_msg"] = "success"
			returnData["id"] = circlePost.Id
		} else {
			returnData["result_code"] = getErrCode(err)
			returnData["result_msg"] = ErrMsg(ctx, err)
		}
	}
	ctx.JSON(returnData)
}

/*发布动态*/
func postCircle(circlePost *CirclePost) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", circlePost.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	circlePost.Status = 1
	circlePost.Id = RandStringBytes(3)
	circlePost.UserId = userInfo.UserId
	circlePost.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(circlePost).Error; err != nil {
			return err
		}
		expr := gorm.Expr("COALESCE(post_count, 0) + ?", 1)
		return tx.Table("users").Where("user_id = ?", userInfo.UserId).Update("post_count", expr).Error
	})
}

/*上传动态图片*/
func uploadCircleImageHandler(ctx iris.Context) {
	var cType string
	var fileBytes []byte
	var picture PostImage
	picture.PostId = ctx.URLParam("id")
	picture.Token = ctx.GetHeader("token")
	returnData := iris.Map{"result_code": 500}
	file, info, err := ctx.FormFile("image")
	if err != nil {
		err = newError(400, "E_INVALID_PICTURE")
	} else if c, _, e := image.DecodeConfig(file); e != nil {
		err = newError(400, "E_INVALID_PICTURE")
	} else {
		/*再次读取*/
		file.Seek(0, 0)
		defer file.Close()
		fileBytes, err = io.ReadAll(file)
		if err != nil {
			/*do nothing*/
		} else if picture.Token == "" {
			err = newError(401, "E_NO_TOKEN")
		} else if picture.PostId == "" {
			err = newError(400, "E_NO_ID")
		} else if cType = http.DetectContentType(fileBytes); cType != JPEG && cType != PNG {
			err = newError(400, "E_INVALID_PICTURE")
		}
		picture.Width = c.Width
		picture.Height = c.Height
	}
	if err != nil {
		returnData["result_code"] = getErrCode(err)
		returnData["result_msg"] = ErrMsg(ctx, err)
	} else {
		picture.Size = strconv.FormatInt(info.Size, 10)
		picture.Type = strings.Split(cType, "/")[1]
		picture.MediaType = "image"
		err = uploadCircleImage(ctx, fileBytes, &picture)
		if err == nil {
			returnData["result_code"] = 200
			returnData["result_msg"] = "success"
		} else {
			returnData["result_code"] = getErrCode(err)
			returnData["result_msg"] = ErrMsg(ctx, err)
		}
	}
	ctx.JSON(returnData)
}

/*批量上传动态图片*/
func batchUploadCircleImageHandler(ctx iris.Context) {
	postId := ctx.URLParam("id")
	token := ctx.GetHeader("token")
	returnData := iris.Map{"result_code": 500}
	if token == "" {
		err := newError(401, "E_NO_TOKEN")
		returnData["result_code"] = getErrCode(err)
		returnData["result_msg"] = ErrMsg(ctx, err)
		ctx.JSON(returnData)
		return
	}
	if postId == "" {
		err := newError(400, "E_NO_ID")
		returnData["result_code"] = getErrCode(err)
		returnData["result_msg"] = ErrMsg(ctx, err)
		ctx.JSON(returnData)
		return
	}
	files, fileHeaders, err := ctx.FormFiles("images")
	if err != nil || len(files) == 0 {
		err := newError(400, "E_INVALID_PICTURE")
		returnData["result_code"] = getErrCode(err)
		returnData["result_msg"] = ErrMsg(ctx, err)
		ctx.JSON(returnData)
		return
	}
	var results []iris.Map
	for i, file := range files {
		fileBytes, readErr := io.ReadAll(file)
		file.Close()
		if readErr != nil {
			results = append(results, iris.Map{"result_code": getErrCode(newError(400, "E_INVALID_PICTURE")), "result_msg": "read file failed"})
			continue
		}
		cType := http.DetectContentType(fileBytes)
		if cType != JPEG && cType != PNG {
			results = append(results, iris.Map{"result_code": getErrCode(newError(400, "E_INVALID_PICTURE")), "result_msg": "invalid picture type"})
			continue
		}
		if len(fileBytes) == 0 {
			results = append(results, iris.Map{"result_code": getErrCode(newError(400, "E_INVALID_PICTURE")), "result_msg": "empty file"})
			continue
		}
		cfg, _, cfgErr := image.DecodeConfig(bytes.NewReader(fileBytes))
		if cfgErr != nil {
			results = append(results, iris.Map{"result_code": getErrCode(newError(400, "E_INVALID_PICTURE")), "result_msg": "decode image config failed"})
			continue
		}
		var picture PostImage
		picture.PostId = postId
		picture.Token = token
		picture.Size = strconv.FormatInt(fileHeaders[i].Size, 10)
		picture.Type = strings.Split(cType, "/")[1]
		picture.MediaType = "image"
		picture.Width = cfg.Width
		picture.Height = cfg.Height
		if uploadErr := uploadCircleImage(ctx, fileBytes, &picture); uploadErr != nil {
			results = append(results, iris.Map{"result_code": getErrCode(uploadErr), "result_msg": ErrMsg(ctx, uploadErr)})
			continue
		}
		results = append(results, iris.Map{"id": picture.Id, "image": picture.Image, "thumbnail": picture.Thumbnail})
	}
	returnData["result_code"] = 200
	returnData["result_msg"] = "success"
	returnData["images"] = results
	ctx.JSON(returnData)
}

/*上传动态图片*/
func uploadCircleImage(ctx iris.Context, imgData []byte, picture *PostImage) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", picture.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	var circlePost CirclePost
	if err := db.Where("id = ? AND status = ?", picture.PostId, 1).First(&circlePost).Error; err != nil {
		return newError(404, "E_NO_CIRCLE_POST")
	}
	if circlePost.UserId != userInfo.UserId {
		return newError(403, "E_NO_AUTH")
	}
	/*解析图片*/
	img, err := parseImageData(bytes.NewBuffer(imgData))
	if err != nil {
		return err
	}
	picture.Id = RandStringBytes(3)
	/*编码为 PNG 存入内存*/
	var pngBuf bytes.Buffer
	if err = png.Encode(&pngBuf, img); err != nil {
		return err
	}
	/*上传原图到 OSS*/
	ossKey := "circle/" + picture.Id + ".png"
	if err = putOssObjectFromBytes(ossKey, pngBuf.Bytes()); err != nil {
		return err
	}
	picture.Image = config.GetOssDomain() + "/" + ossKey
	/*生成缩略图*/
	var size, scale int
	size = min(picture.Width, picture.Height)
	scale = size / 200
	if scale > 1 {
		img, err = resizeImg(img, uint(picture.Width/scale), uint(picture.Height/scale))
		if err != nil {
			return err
		}
	}
	var thumbBuf bytes.Buffer
	if err = png.Encode(&thumbBuf, img); err != nil {
		return err
	}
	/*上传缩略图到 OSS*/
	thumbKey := "circle/thumbnail/" + picture.Id + ".png"
	if err = putOssObjectFromBytes(thumbKey, thumbBuf.Bytes()); err != nil {
		return err
	}
	picture.Thumbnail = config.GetOssDomain() + "/" + thumbKey
	picture.Time = time.Now().Format("2006-01-02 15:04:05")
	return db.Create(&picture).Error
}

/*上传动态视频*/
func uploadCircleVideoHandler(ctx iris.Context) {
	var cType string
	var fileBytes []byte
	var picture PostImage
	picture.PostId = ctx.URLParam("id")
	picture.Token = ctx.GetHeader("token")
	picture.MediaType = "video"
	returnData := iris.Map{"result_code": 500}
	file, info, err := ctx.FormFile("video")
	if err != nil {
		err = newError(400, "E_INVALID_VIDEO")
	} else {
		defer file.Close()
		fileBytes, err = io.ReadAll(file)
		if err != nil {
			/*do nothing*/
		} else if picture.Token == "" {
			err = newError(401, "E_NO_TOKEN")
		} else if picture.PostId == "" {
			err = newError(400, "E_NO_ID")
		} else {
			cType = http.DetectContentType(fileBytes)
			if cType != MP4 && cType != MOV {
				err = newError(400, "E_INVALID_VIDEO")
			}
		}
	}
	if err != nil {
		returnData["result_code"] = getErrCode(err)
		returnData["result_msg"] = ErrMsg(ctx, err)
	} else {
		picture.Size = strconv.FormatInt(info.Size, 10)
		picture.Type = strings.Split(cType, "/")[1]
		err = uploadCircleVideo(ctx, fileBytes, &picture)
		if err == nil {
			returnData["result_code"] = 200
			returnData["result_msg"] = "success"
		} else {
			returnData["result_code"] = getErrCode(err)
			returnData["result_msg"] = ErrMsg(ctx, err)
		}
	}
	ctx.JSON(returnData)
}

/*上传动态视频*/
func uploadCircleVideo(ctx iris.Context, videoData []byte, picture *PostImage) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", picture.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	var circlePost CirclePost
	if err := db.Where("id = ? AND status = ?", picture.PostId, 1).First(&circlePost).Error; err != nil {
		return newError(404, "E_NO_CIRCLE_POST")
	}
	if circlePost.UserId != userInfo.UserId {
		return newError(403, "E_NO_AUTH")
	}
	picture.Id = RandStringBytes(3)
	/*保存视频到临时文件（ffprobe/ffmpeg 需要文件路径）*/
	videoName := picture.Id + "." + picture.Type
	tmpVideoPath := "public/circle/" + videoName
	if err := os.MkdirAll("public/circle", 0755); err != nil {
		return err
	}
	if err := os.WriteFile(tmpVideoPath, videoData, 0644); err != nil {
		return err
	}
	defer os.Remove(tmpVideoPath)
	/*上传视频到 OSS*/
	ossVideoKey := "circle/" + videoName
	if err := putOssObjectFromBytes(ossVideoKey, videoData); err != nil {
		return err
	}
	picture.Image = config.GetOssDomain() + "/" + ossVideoKey
	/*用 ffprobe 获取视频宽高和时长*/
	probeVideoInfo(tmpVideoPath, picture)
	/*用 ffmpeg 截取首帧作为缩略图*/
	thumbName := picture.Id + ".png"
	tmpThumbPath := "public/circle/" + thumbName
	if err := extractVideoThumbnail(tmpVideoPath, tmpThumbPath, picture); err == nil {
		defer os.Remove(tmpThumbPath)
		thumbData, err := os.ReadFile(tmpThumbPath)
		if err == nil {
			ossThumbKey := "circle/thumbnail/" + thumbName
			if err = putOssObjectFromBytes(ossThumbKey, thumbData); err == nil {
				picture.Thumbnail = config.GetOssDomain() + "/" + ossThumbKey
			}
		}
	}
	picture.Time = time.Now().Format("2006-01-02 15:04:05")
	return db.Create(&picture).Error
}

/*用 ffprobe 获取视频元数据*/
func probeVideoInfo(videoPath string, picture *PostImage) error {
	cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json",
		"-show_streams", "-show_format", videoPath)
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	output := string(out)
	/*简单解析 width、height、duration（避免引入额外 JSON 解析库）*/
	picture.Width = parseProbeFieldInt(output, `"width"`)
	picture.Height = parseProbeFieldInt(output, `"height"`)
	picture.Duration = parseProbeFieldFloat(output, `"duration"`)
	return nil
}

/*从 ffprobe JSON 输出中解析整型字段*/
func parseProbeFieldInt(jsonStr string, field string) int {
	idx := strings.Index(jsonStr, field)
	if idx < 0 {
		return 0
	}
	sub := jsonStr[idx+len(field):]
	sub = strings.TrimLeft(sub, `": `)
	end := strings.IndexAny(sub, ",\n}")
	if end < 0 {
		end = len(sub)
	}
	val, _ := strconv.Atoi(strings.TrimSpace(sub[:end]))
	return val
}

/*从 ffprobe JSON 输出中解析浮点字段*/
func parseProbeFieldFloat(jsonStr string, field string) float64 {
	idx := strings.Index(jsonStr, field)
	if idx < 0 {
		return 0
	}
	sub := jsonStr[idx+len(field):]
	sub = strings.TrimLeft(sub, `": `)
	end := strings.IndexAny(sub, ",\n}")
	if end < 0 {
		end = len(sub)
	}
	val, _ := strconv.ParseFloat(strings.TrimSpace(sub[:end]), 64)
	return val
}

/*用 ffmpeg 截取视频首帧作为缩略图*/
func extractVideoThumbnail(videoPath string, thumbPath string, picture *PostImage) error {
	args := []string{"-y", "-i", videoPath, "-ss", "1", "-vframes", "1", "-vf", "scale=400:-1"}
	if picture.Duration > 0 && picture.Duration < 1 {
		args[3] = "0.5"
	}
	args = append(args, thumbPath)
	cmd := exec.Command("ffmpeg", args...)
	return cmd.Run()
}

/*查询动态*/
func listCirclePostHandler(ctx iris.Context) {
	var err error
	var params Params
	params.UserId = ctx.URLParam("user_id")
	params.Keyword = ctx.URLParam("keyword")
	params.Topic = ctx.URLParam("topic")
	params.ReviewStatus = ctx.URLParam("review_status")
	params.Token = ctx.GetHeader("token")
	params.Page = AtoUI(ctx.URLParam("page"), 1)
	params.Limit = AtoUI(ctx.URLParam("limit"), 10)
	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	data, err := listCirclePost(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "total": params.Total, "data": data})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
	}
}

/*查询动态*/
func listCirclePost(params *Params) ([]CirclePostOut, error) {
	db := getMysqlConn()
	var userInfo User
	err := db.Where("token = ?", params.Token).First(&userInfo).Error
	if err != nil {
		return nil, newError(401, "E_NO_TOKEN")
	}
	db = db.Table("circle_posts")
	db = db.Select("users.nickname, users.avatar_object_key as avatar, users.biography, users.follow_count, users.fans_count, users.post_count, "+
		"circle_posts.id, circle_posts.user_id, circle_posts.content, circle_posts.topic, circle_posts.create_time, circle_posts.like_count, "+
		"circle_posts.collect_count, circle_posts.comment_count, circle_posts.review_status, circle_posts.review_time, circle_posts.reviewer, "+
		"CASE WHEN EXISTS (SELECT 1 FROM follows WHERE follows.follower_id = ? AND follows.following_id = circle_posts.user_id AND follows.status = 1) THEN 1 ELSE 0 END as is_followed", userInfo.UserId)
	db = db.Joins("LEFT JOIN users ON circle_posts.user_id = users.user_id")
	db = db.Where("circle_posts.status = ?", 1)
	if params.UserId != "" {
		db = db.Where("circle_posts.user_id = ?", params.UserId)
	}
	if params.Keyword != "" {
		keyword := "%" + params.Keyword + "%"
		db = db.Where("(circle_posts.content LIKE ? OR circle_posts.topic LIKE ?)", keyword, keyword)
	}
	if params.Topic != "" {
		db = db.Where("circle_posts.topic = ?", params.Topic)
	}
	if params.ReviewStatus == "0" || params.ReviewStatus == "1" {
		db = db.Where("circle_posts.review_status = ?", params.ReviewStatus)
	}
	err = db.Count(&params.Total).Error
	if err != nil {
		return nil, err
	}
	data := []CirclePostOut{}
	db = db.Offset((params.Page - 1) * params.Limit).Limit(params.Limit)
	err = db.Preload("Images").Preload("Likes").Preload("Likes.User").
		Preload("Comments").Preload("Comments.User").
		Order("create_time desc").Find(&data).Error
	if err != nil {
		return nil, err
	}
	for i := range data {
		if data[i].Avatar != "" {
			if u, err := signAvatarURL(data[i].Avatar); err == nil {
				data[i].Avatar = u
			}
		}
		for j := range data[i].Likes {
			if data[i].Likes[j].User.Avatar != "" {
				if u, err := signAvatarURL(data[i].Likes[j].User.Avatar); err == nil {
					data[i].Likes[j].User.Avatar = u
				}
			}
		}
		for j := range data[i].Comments {
			if data[i].Comments[j].User.Avatar != "" {
				if u, err := signAvatarURL(data[i].Comments[j].User.Avatar); err == nil {
					data[i].Comments[j].User.Avatar = u
				}
			}
		}
	}
	return data, nil
}

/*删除动态*/
func removeCircleHandler(ctx iris.Context) {
	var params Params
	params.Token = ctx.GetHeader("token")
	returnData := iris.Map{"result_code": 500}
	err := ctx.ReadJSON(&params)
	if err != nil {
	} else if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} else if params.Id == "" {
		err = newError(400, "E_NO_ID")
	}
	if err != nil {
		returnData["result_code"] = getErrCode(err)
		returnData["result_msg"] = ErrMsg(ctx, err)
	} else {
		err = removeCircle(&params)
		if err == nil {
			returnData["result_code"] = 200
			returnData["result_msg"] = "success"
		} else {
			returnData["result_code"] = getErrCode(err)
			returnData["result_msg"] = ErrMsg(ctx, err)
		}
	}
	ctx.JSON(returnData)
}

/*标记审核动态*/
func reviewCircleHandler(ctx iris.Context) {
	var review CircleReview
	review.Token = ctx.GetHeader("token")
	returnData := iris.Map{"result_code": 500}
	err := ctx.ReadJSON(&review)
	if err != nil {
	} else if review.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} else if review.Id == "" {
		err = newError(400, "E_NO_ID")
	} else if review.ReviewStatus != 0 && review.ReviewStatus != 1 {
		err = newError(400, "E_INVALID_REVIEW_STATUS")
	}
	if err != nil {
		returnData["result_code"] = getErrCode(err)
		returnData["result_msg"] = ErrMsg(ctx, err)
	} else {
		err = reviewCircle(&review)
		if err == nil {
			returnData["result_code"] = 200
			returnData["result_msg"] = "success"
		} else {
			returnData["result_code"] = getErrCode(err)
			returnData["result_msg"] = ErrMsg(ctx, err)
		}
	}
	ctx.JSON(returnData)
}

/*标记审核动态*/
func reviewCircle(review *CircleReview) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", review.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	if userInfo.UserId != ADMIN {
		return newError(403, "E_NO_AUTH")
	}
	var circlePost CirclePost
	if err := db.Where("id = ? AND status = ?", review.Id, 1).First(&circlePost).Error; err != nil {
		return newError(404, "E_NO_CIRCLE_POST")
	}
	updates := map[string]interface{}{
		"review_status": review.ReviewStatus,
	}
	if review.ReviewStatus == 1 {
		updates["review_time"] = time.Now().Format("2006-01-02 15:04:05")
		updates["reviewer"] = userInfo.UserId
	} else {
		updates["review_time"] = ""
		updates["reviewer"] = ""
	}
	return db.Model(&CirclePost{}).Where("id = ?", circlePost.Id).Updates(updates).Error
}

/*删除动态*/
func removeCircle(params *Params) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	var circlePost CirclePost
	if err := db.Where("id = ? AND status = ?", params.Id, 1).First(&circlePost).Error; err != nil {
		return newError(404, "E_NO_CIRCLE_POST")
	}
	if userInfo.UserId != circlePost.UserId && userInfo.UserId != ADMIN {
		return newError(403, "E_NO_AUTH")
	}
	/*删除关联的 OSS 文件*/
	var images []PostImage
	db.Where("post_id = ?", circlePost.Id).Find(&images)
	for _, img := range images {
		deleteCircleOssFile(img.Image)
		deleteCircleOssFile(img.Thumbnail)
	}
	return db.Transaction(func(tx *gorm.DB) error {
		/*删除关联的图片记录*/
		if err := tx.Exec("DELETE FROM post_images WHERE post_id = ?", circlePost.Id).Error; err != nil {
			return err
		}
		/*删除关联的评论记录*/
		if err := tx.Exec("DELETE FROM comments WHERE post_id = ?", circlePost.Id).Error; err != nil {
			return err
		}
		/*删除关联的点赞记录*/
		if err := tx.Exec("DELETE FROM likes WHERE post_id = ?", circlePost.Id).Error; err != nil {
			return err
		}
		/*删除关联的收藏记录*/
		if err := tx.Exec("DELETE FROM collects WHERE post_id = ?", circlePost.Id).Error; err != nil {
			return err
		}
		/*完全删除动态记录*/
		if err := tx.Exec("DELETE FROM circle_posts WHERE id = ?", circlePost.Id).Error; err != nil {
			return err
		}
		expr := gorm.Expr("GREATEST(COALESCE(post_count, 0) - ?, 0)", 1)
		return tx.Table("users").Where("user_id = ?", circlePost.UserId).Update("post_count", expr).Error
	})
}

/*删除 OSS 文件（忽略错误）*/
func deleteCircleOssFile(url string) {
	if url == "" {
		return
	}
	domain := config.GetOssDomain()
	if domain == "" || !strings.HasPrefix(url, domain) {
		return
	}
	key := strings.TrimPrefix(url, domain)
	key = strings.TrimPrefix(key, "/")
	deleteOssObject(key)
}
