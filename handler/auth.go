package handler

import "github.com/kataras/iris/v12"

// AuthHandler 认证相关接口（挑战码、Apple OAuth 浏览器回调等；不含用户登录发 token）。
func AuthHandler(auth iris.Party) {
	auth.Post("/challenge", socialChallengeHandler)
	auth.Post("/apple/callback", appleCallbackHandler)
}
