package handler

import "github.com/kataras/iris/v12"

// AuthHandler 认证相关接口（挑战码、签名等）。
func AuthHandler(auth iris.Party) {
	auth.Post("/challenge", socialChallengeHandler)
}
