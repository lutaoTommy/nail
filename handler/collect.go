package handler

import (
	"time"

	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

/*收藏管理*/
func CollectHandler(collect iris.Party) {
	/*收藏动态*/
	collect.Post("/create", createCollectHandler)
	/*查询某动态的收藏列表*/
	collect.Get("/list", listCollectHandler)
	/*取消收藏*/
	collect.Post("/remove", removeCollectHandler)
	/*我的收藏列表*/
	collect.Get("/my", listMyCollectsHandler)
}

/*收藏动态*/
func createCollectHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	if err = ctx.ReadJSON(&params); err != nil {
	} else if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} else if params.Id == "" {
		err = newError(400, "E_NO_ID")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	err = createCollect(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success"})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*收藏动态*/
func createCollect(params *Params) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	var circlePost CirclePost
	if err := db.Where("id = ? AND status = ?", params.Id, 1).First(&circlePost).Error; err != nil {
		return newError(404, "E_NO_CIRCLE_POST")
	}
	var exist Collect
	if err := db.Where("user_id = ? AND post_id = ?", userInfo.UserId, circlePost.Id).First(&exist).Error; err == nil {
		return nil // 已收藏，幂等返回成功
	}
	collect := Collect{
		Id:     RandStringBytes(3),
		PostId: circlePost.Id,
		UserId: userInfo.UserId,
		Time:   time.Now().Format("2006-01-02 15:04:05"),
	}
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&collect).Error; err != nil {
			return err
		}
		expr := gorm.Expr("COALESCE(collect_count, 0) + ?", 1)
		return tx.Table("circle_posts").Where("id = ?", circlePost.Id).Update("collect_count", expr).Error
	})
}

/*查询某动态的收藏列表*/
func listCollectHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Id = ctx.URLParam("id")
	params.Token = ctx.GetHeader("token")
	params.Page = AtoUI(ctx.URLParam("page"), 0)
	params.Limit = AtoUI(ctx.URLParam("limit"), 10)
	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} else if params.Id == "" {
		err = newError(400, "E_NO_ID")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	data, err := listCollect(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "total": params.Total, "data": data})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*查询某动态的收藏列表*/
func listCollect(params *Params) ([]Collect, error) {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return nil, newError(401, "E_NO_TOKEN")
	}
	var circlePost CirclePost
	if err := db.Where("id = ? AND status = ?", params.Id, 1).First(&circlePost).Error; err != nil {
		return nil, newError(404, "E_NO_CIRCLE_POST")
	}
	db = db.Table("collects")
	db = db.Where("post_id = ?", params.Id)
	var err error
	err = db.Count(&params.Total).Error
	if err != nil {
		return nil, err
	}
	data := []Collect{}
	if params.Page > 0 {
		db = db.Offset((params.Page - 1) * params.Limit).Limit(params.Limit)
	}
	err = db.Preload("User").Order("time desc").Find(&data).Error
	if err != nil {
		return nil, err
	}
	for i := range data {
		if data[i].User.Avatar == "" {
			continue
		}
		if u, err := signAvatarURL(data[i].User.Avatar); err == nil {
			data[i].User.Avatar = u
		}
	}
	return data, nil
}

/*取消收藏*/
func removeCollectHandler(ctx iris.Context) {
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
		err = removeCollect(&params)
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

/*取消收藏（传入动态 id，与收藏接口一致）*/
func removeCollect(params *Params) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	var circlePost CirclePost
	if err := db.Where("id = ? AND status = ?", params.Id, 1).First(&circlePost).Error; err != nil {
		return newError(404, "E_NO_CIRCLE_POST")
	}
	var collect Collect
	if err := db.Where("user_id = ? AND post_id = ?", userInfo.UserId, params.Id).First(&collect).Error; err != nil {
		return newError(400, "E_NO_COLLECT") // 未收藏过，取消失败
	}
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&collect).Error; err != nil {
			return err
		}
		expr := gorm.Expr("GREATEST(COALESCE(collect_count, 0) - ?, 0)", 1)
		return tx.Table("circle_posts").Where("id = ?", params.Id).Update("collect_count", expr).Error
	})
}

/*我的收藏列表*/
func listMyCollectsHandler(ctx iris.Context) {
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
	data, err := listMyCollects(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "total": params.Total, "data": data})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*我的收藏列表*/
func listMyCollects(params *Params) ([]CirclePostOut, error) {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return nil, newError(401, "E_NO_TOKEN")
	}
	db = db.Table("circle_posts")
	db = db.Select("users.nickname, users.avatar_object_key as avatar, users.biography, users.follow_count, users.fans_count, users.post_count, " +
		"circle_posts.id, circle_posts.user_id, circle_posts.content, circle_posts.create_time, circle_posts.like_count, " +
		"circle_posts.collect_count, circle_posts.comment_count")
	db = db.Joins("LEFT JOIN users ON circle_posts.user_id = users.user_id")
	db = db.Joins("INNER JOIN collects ON collects.post_id = circle_posts.id AND collects.user_id = ?", userInfo.UserId)
	db = db.Where("circle_posts.status = ?", 1)
	err := db.Count(&params.Total).Error
	if err != nil {
		return nil, err
	}
	data := []CirclePostOut{}
	db = db.Offset((params.Page - 1) * params.Limit).Limit(params.Limit)
	err = db.Preload("Images").Preload("Likes").Preload("Likes.User").
		Preload("Comments").Preload("Comments.User").
		Order("collects.time desc").Find(&data).Error
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
