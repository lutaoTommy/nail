package handler

import (
	"time"

	"github.com/kataras/iris/v12"
)

/*话题管理*/
func TopicHandler(topic iris.Party) {
	/*查询话题列表*/
	topic.Get("/list", listTopicHandler)
	/*添加话题*/
	topic.Post("/add", addTopicHandler)
	/*删除话题*/
	topic.Post("/remove", removeTopicHandler)
}

/*查询话题列表*/
func listTopicHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	data, err := listTopic()
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success", "data": data})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
	}
}

/*查询话题列表逻辑*/
func listTopic() ([]Topic, error) {
	db := getMysqlConn()
	var topics []Topic
	err := db.Where("status = ?", 1).Order("id asc").Find(&topics).Error
	return topics, err
}

/*添加话题*/
func addTopicHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	params.Name = ctx.URLParam("name")
	if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} else if params.Name == "" {
		err = newError(400, "E_NO_NAME")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	err = addTopic(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success"})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
	}
}

/*添加话题逻辑*/
func addTopic(params *Params) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	if userInfo.UserId != ADMIN {
		return newError(403, "E_NO_AUTH")
	}
	var exist Topic
	if err := db.Where("name = ?", params.Name).First(&exist).Error; err == nil {
		return newError(400, "E_TOPIC_EXIST")
	}
	topic := Topic{
		Id:         RandStringBytes(3),
		Name:       params.Name,
		Status:     1,
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	return db.Create(&topic).Error
}

/*删除话题*/
func removeTopicHandler(ctx iris.Context) {
	var err error
	var params Params
	params.Token = ctx.GetHeader("token")
	err = ctx.ReadJSON(&params)
	if err != nil {
	} else if params.Token == "" {
		err = newError(401, "E_NO_TOKEN")
	} else if params.Id == "" {
		err = newError(400, "E_NO_ID")
	}
	if err != nil {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	err = removeTopic(&params)
	if err == nil {
		ctx.JSON(iris.Map{"result_code": 200, "result_msg": "success"})
	} else {
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
	}
}

/*删除话题逻辑*/
func removeTopic(params *Params) error {
	db := getMysqlConn()
	var userInfo User
	if err := db.Where("token = ?", params.Token).First(&userInfo).Error; err != nil {
		return newError(401, "E_NO_TOKEN")
	}
	if userInfo.UserId != ADMIN {
		return newError(403, "E_NO_AUTH")
	}
	var topic Topic
	if err := db.Where("id = ?", params.Id).First(&topic).Error; err != nil {
		return newError(404, "E_NO_TOPIC")
	}
	return db.Delete(&topic).Error
}
