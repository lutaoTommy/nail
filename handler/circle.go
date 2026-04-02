package handler

import (
	"bytes"
	"image"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

/*圈子管理*/
func CircleHandler(circle iris.Party) {
	/*发布动态*/
	circle.Post("/post", postCircleHandler)
	/*上传图片*/
	circle.Post("/image/upload", uploadCircleImageHandler)
	/*动态列表*/
	circle.Get("/list", listCirclePostHandler)
	/*删除动态*/
	circle.Post("/remove", removeCircleHandler)
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
		returnData["result_msg"] = err.Error()
	} else {
		err = postCircle(&circlePost)
		if err == nil {
			returnData["result_code"] = 200
			returnData["result_msg"] = "success"
			returnData["id"] = circlePost.Id
		} else {
			returnData["result_code"] = getErrCode(err)
			returnData["result_msg"] = err.Error()
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
		err = e
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
		returnData["result_msg"] = err.Error()
	} else {
		picture.Size = strconv.FormatInt(info.Size, 10)
		picture.Type = strings.Split(cType, "/")[1]
		err = uploadCircleImage(ctx, fileBytes, &picture)
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
	var img image.Image
	var err error
	img, err = parseImageData(bytes.NewBuffer(imgData))
	if err != nil {
		return err
	}
	/*保存图片*/
	picture.Id = RandStringBytes(3)
	name := picture.Id + ".png"
	path := "public/circle/" + name
	err = savePngImage(img, path)
	if err != nil {
		return err
	}
	picture.Image = BuildPublicURL(ctx, "/circle/"+name)
	/*保存缩略图*/
	var size, scale int
	size = min(picture.Width, picture.Height)
	scale = size / 200
	if scale > 1 {
		img, err = resizeImg(img, uint(picture.Width/scale), uint(picture.Height/scale))
		if err != nil {
			return err
		}
	}
	path = "public/thumbnail/" + name
	err = savePngImage(img, path)
	if err != nil {
		return err
	} 
	picture.Thumbnail = BuildPublicURL(ctx, "/thumbnail/"+name)
	picture.Time = time.Now().Format("2006-01-02 15:04:05")
	return db.Create(&picture).Error
}

/*查询动态*/
func listCirclePostHandler(ctx iris.Context) {
	var err error
	var params Params
	params.UserId = ctx.URLParam("user_id")
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
	data, err := listCirclePost(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "total": params.Total, "data": data})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
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
	db = db.Select("users.nickname, users.avatar, users.biography, users.follow_count, users.fans_count, users.post_count, " +
		"circle_posts.id, circle_posts.user_id, circle_posts.content, circle_posts.create_time, circle_posts.like_count, " + 
		"circle_posts.collect_count, circle_posts.comment_count")
	db = db.Joins("LEFT JOIN users ON circle_posts.user_id = users.user_id")
	db = db.Where("circle_posts.status = ?", 1)
	if params.UserId !=  "" {
		db = db.Where("circle_posts.user_id = ?", params.UserId)
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
    return data, err
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
		returnData["result_msg"] = err.Error()
	} else {
		err = removeCircle(&params)
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
	if userInfo.UserId != circlePost.UserId {
		return newError(403, "E_NO_AUTH")
	}
	now := time.Now().Format("2006-01-02 15:04:05")
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&circlePost).Updates(map[string]interface{}{
			"status":       -1,
			"update_time":  now,
		}).Error; err != nil {
			return err
		}
		expr := gorm.Expr("GREATEST(COALESCE(post_count, 0) - ?, 0)", 1)
		return tx.Table("users").Where("user_id = ?", userInfo.UserId).Update("post_count", expr).Error
	})
}


