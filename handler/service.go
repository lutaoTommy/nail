package handler

import (
	"nail/config"
	"github.com/kataras/iris/v12"
)

/*服务器管理*/
func ServiceHandler(service iris.Party) {
	/*服务器版本*/
	service.Get("/version", serviceVersionHander)
}


/*服务器版本*/
func serviceVersionHander(ctx iris.Context) {
	ctx.JSON(iris.Map{
		"result_code":  200,
		"result_msg":   "success",
		"area":         config.GetArea(),
		"version":      config.GetVersion(),
	})
}
