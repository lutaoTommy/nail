package handler

import (
	"strings"

	"nail/language"

	"github.com/kataras/iris/v12"
)

// appleExchangeHandler App 凭回调页下发的 ticket 一次性换取 OAuth 参数（避免 Deep Link URL 过长）。
func appleExchangeHandler(ctx iris.Context) {
	ticket := strings.TrimSpace(ctx.URLParam("ticket"))
	if ticket == "" {
		ctx.JSON(iris.Map{"result_code": 400, "result_msg": language.GetRawMessage("E_INVALID_PARAM")})
		return
	}
	payload, ok := consumeAppleOAuthTicket(ticket)
	if !ok {
		ctx.JSON(iris.Map{"result_code": 404, "result_msg": language.GetRawMessage("E_APPLE_TICKET_INVALID")})
		return
	}
	res := iris.Map{
		"result_code": 200,
		"result_msg":  "success",
		"state":       payload.State,
	}
	if payload.Code != "" {
		res["code"] = payload.Code
	}
	if payload.IDToken != "" {
		res["id_token"] = payload.IDToken
	}
	ctx.JSON(res)
}
