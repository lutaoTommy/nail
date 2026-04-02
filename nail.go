package main

import (
	"fmt"
	"nail/config"
	"nail/handler"
	"nail/language"
	"nail/logger"

	"github.com/kataras/iris/v12"
)

func main() {
	/*加载端口配置*/
	err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	app := iris.New()
	app.OnErrorCode(iris.StatusNotFound, notFoundHandler)
	app.Use(func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		if ctx.Request().Method == "OPTIONS" {
			ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
			ctx.Header("Access-Control-Allow-Headers", "Content-Type, x-requested-with, Token")
			ctx.StatusCode(204)
			return
		}
		ctx.Next()
	})
	app.HandleDir("/", "./public")
	app.AllowMethods(iris.MethodOptions)
	app.PartyFunc("/ota", handler.OtaHandler)
	app.PartyFunc("/user", handler.UserHandler)
	app.PartyFunc("/word", handler.WordHandler)
	app.PartyFunc("/like", handler.LikeHandler)
	app.PartyFunc("/collect", handler.CollectHandler)
	app.PartyFunc("/color", handler.ColorHandler)
	app.PartyFunc("/circle", handler.CircleHandler)
	app.PartyFunc("/device", handler.DeviceHandler)
	app.PartyFunc("/avatar", handler.AvatarHandler)
	app.PartyFunc("/comment", handler.CommentHandler)
	app.PartyFunc("/suggest", handler.SuggestHandler)
	app.PartyFunc("/apk", handler.ApkHandler)
	/*日志初始化*/
	logger.InitLogger()
	/*多语言初始化*/
	language.InitLanguage()
	/*初始化Tire树*/
	handler.InitTrie()
	app.Run(iris.Addr(fmt.Sprintf(":%d", config.GetHttpPort())), iris.WithConfiguration(iris.Configuration{
		DisablePathCorrection: true,
	}))
}

func notFoundHandler(ctx iris.Context) {
	ctx.HTML("404: no route for " + ctx.Path())
}
