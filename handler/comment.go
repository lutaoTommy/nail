package handler

import (
	"time"
	"gorm.io/gorm"
	"github.com/kataras/iris/v12"
)

/*评论管理*/
func CommentHandler(comment iris.Party) {
	/*发布评论*/
	comment.Post("/create", createCommentHandler)
	/*查询评论*/
	comment.Get("/list", listCommentHandler)
	/*删除评论*/
	comment.Post("/remove", removeCommentHandler)
}

/*发布评论*/
func createCommentHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	if err = ctx.ReadJSON(&params); err != nil {
	} else if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} else if params.Id == "" {
		err = newError(400, "E_NO_ID")
	} else if err = params.checkContent(); err != nil {
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
		return
	}
	err = createComment(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success"})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*发布评论*/
func createComment(params *Params) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	var circlePost CirclePost
	if err := db.Where("id = ? AND status = ?", params.Id, 1).First(&circlePost).Error; err != nil {
		return newError(404, "E_NO_CIRCLE_POST")
	}
	comment := Comment{
		Id:      RandStringBytes(3),
		Content: params.Content,
		PostId:  circlePost.Id,
		UserId:  userInfo.UserId,
		Time:    time.Now().Format("2006-01-02 15:04:05"),
	}
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&comment).Error; err != nil {
			return err
		}
		expr := gorm.Expr("COALESCE(comment_count, 0) + ?", 1)
		return tx.Table("circle_posts").Where("id = ?", circlePost.Id).Update("comment_count", expr).Error
	})
}


/*查询评论*/
func listCommentHandler(ctx iris.Context) {
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
	data, err := listComment(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "total": params.Total, "data": data})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": err.Error()})
	}
}

/*查询评论*/
func listComment(params *Params) ([]Comment, error) {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return nil, newError(401, "E_NO_TOKEN")
	}
	var circlePost CirclePost
	if err := db.Where("id = ? AND status = ?", params.Id, 1).First(&circlePost).Error; err != nil {
		return nil, newError(404, "E_NO_CIRCLE_POST")
	}
	db = db.Table("comments")
	db = db.Where("post_id = ?", params.Id)
	var err error
	err = db.Count(&params.Total).Error
	if err != nil {
		return nil, err
	}
    data := []Comment{}
    if params.Page > 0 {
    	db = db.Offset((params.Page - 1) * params.Limit).Limit(params.Limit)
    }
    err = db.Preload("User").Order("time desc").Find(&data).Error
    return data, err
}


/*删除评论*/
func removeCommentHandler(ctx iris.Context) {
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
		err = removeComment(&params)
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

/*删除评论*/
func removeComment(params *Params) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	var comment Comment
	if err := db.Where("id = ?", params.Id).First(&comment).Error; err != nil {
		return newError(404, "E_NO_COMMENT")
	}
	if comment.UserId != userInfo.UserId {
		return newError(403, "E_NO_AUTH")
	}
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&comment).Error; err != nil {
			return err
		}
		expr := gorm.Expr("GREATEST(COALESCE(comment_count, 0) - ?, 0)", 1)
		return tx.Table("circle_posts").Where("id = ?", comment.PostId).Update("comment_count", expr).Error
	})
}

