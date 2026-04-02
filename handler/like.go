package handler

import (
	"time"
	"gorm.io/gorm"
	"github.com/kataras/iris/v12"
)

/*点赞管理*/
func LikeHandler(like iris.Party) {
	/*发布点赞*/
	like.Post("/create", createLikeHandler)
	/*查询点赞*/
	like.Get("/list", listLikeHandler)
	/*删除点赞*/
	like.Post("/remove", removeLikeHandler)
}

/*发布点赞*/
func createLikeHandler(ctx iris.Context) {
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
	err = createLike(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success"})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*发布点赞*/
func createLike(params *Params) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	var circlePost CirclePost
	if err := db.Where("id = ? AND status = ?", params.Id, 1).First(&circlePost).Error; err != nil {
		return newError(404, "E_NO_CIRCLE_POST")
	}
	var exist Like
	if err := db.Where("user_id = ? AND post_id = ?", userInfo.UserId, circlePost.Id).First(&exist).Error; err == nil {
		return nil // 已点赞，幂等返回成功
	}
	like := Like{
		Id:     RandStringBytes(3),
		PostId: circlePost.Id,
		UserId: userInfo.UserId,
		Time:   time.Now().Format("2006-01-02 15:04:05"),
	}
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&like).Error; err != nil {
			return err
		}
		expr := gorm.Expr("COALESCE(like_count, 0) + ?", 1)
		return tx.Table("circle_posts").Where("id = ?", circlePost.Id).Update("like_count", expr).Error
	})
}


/*查询点赞*/
func listLikeHandler(ctx iris.Context) {
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
	data, err := listLike(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "total": params.Total, "data": data})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*查询点赞*/
func listLike(params *Params) ([]Like, error) {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return nil, newError(401, "E_NO_TOKEN")
	}
	var circlePost CirclePost
	if err := db.Where("id = ? AND status = ?", params.Id, 1).First(&circlePost).Error; err != nil {
		return nil, newError(404, "E_NO_CIRCLE_POST")
	}
	db = db.Table("likes")
	db = db.Where("post_id = ?", params.Id)
	var err error
	err = db.Count(&params.Total).Error
	if err != nil {
		return nil, err
	}
    data := []Like{}
    if params.Page > 0 {
    	db = db.Offset((params.Page - 1) * params.Limit).Limit(params.Limit)
    }
    err = db.Preload("User").Order("time desc").Find(&data).Error
    return data, err
}


/*删除点赞*/
func removeLikeHandler(ctx iris.Context) {
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
		err = removeLike(&params)
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

/*删除点赞（传入动态 id，与点赞接口一致）*/
func removeLike(params *Params) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	var circlePost CirclePost
	if err := db.Where("id = ? AND status = ?", params.Id, 1).First(&circlePost).Error; err != nil {
		return newError(404, "E_NO_CIRCLE_POST")
	}
	var like Like
	if err := db.Where("user_id = ? AND post_id = ?", userInfo.UserId, params.Id).First(&like).Error; err != nil {
		return newError(400, "E_NO_LIKE") // 未点赞过，取消失败
	}
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&like).Error; err != nil {
			return err
		}
		expr := gorm.Expr("GREATEST(COALESCE(like_count, 0) - ?, 0)", 1)
		return tx.Table("circle_posts").Where("id = ?", params.Id).Update("like_count", expr).Error
	})
}

