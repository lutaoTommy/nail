package handler

import (
	"nail/nlu"

	log "github.com/jeanphorn/log4go"
	"github.com/kataras/iris/v12"
)

func VoiceHandler(voice iris.Party) {
	voice.Post("/command", voiceCommandHandler)
}

type voiceCommandReq struct {
	Token string `json:"token"`
	Text  string `json:"text"`
}

type voiceCmdResp struct {
	Intent string      `json:"intent"`
	Detail interface{} `json:"detail,omitempty"`
}

func voiceCommandHandler(ctx iris.Context) {
	var req voiceCommandReq
	req.Token = ctx.GetHeader("token")
	if err := ctx.ReadJSON(&req); err != nil {
		log.LOGGER("app").Warn("voice/command invalid param: %v", err)
		ctx.JSON(iris.Map{"result_code": 400, "result_msg": Msg(ctx, "E_INVALID_PARAM")})
		return
	}
	if req.Token == "" {
		log.LOGGER("app").Warn("voice/command no token")
		ctx.JSON(iris.Map{"result_code": 401, "result_msg": Msg(ctx, "E_NO_TOKEN")})
		return
	}
	if req.Text == "" {
		log.LOGGER("app").Warn("voice/command no text: token=%s", req.Token[:min(len(req.Token), 8)]+"...")
		ctx.JSON(iris.Map{"result_code": 400, "result_msg": Msg(ctx, "E_NO_VOICE_TEXT")})
		return
	}

	result := nlu.Classify(req.Text)
	log.LOGGER("app").Info("voice/command input: token=%s text=%q intent=%s", req.Token[:min(len(req.Token), 8)]+"...", req.Text, result.Intent)

	switch result.Intent {
	case nlu.IntentChangeColor:
		execChangeColor(ctx, req.Token, req.Text)
	case nlu.IntentStartDelivery:
		execStartDelivery(ctx, req.Token, req.Text)
	case nlu.IntentCollectColor:
		execCollectColor(ctx, req.Token, req.Text)
	default:
		log.LOGGER("app").Warn("voice/command unknown intent: text=%q", req.Text)
		ctx.JSON(iris.Map{
			"result_code": 400,
			"result_msg":  Msg(ctx, "V_UNKNOWN_COMMAND"),
			"data": voiceCmdResp{
				Intent: "unknown",
			},
		})
	}
}

func execChangeColor(ctx iris.Context, token, text string) {
	items, err := parseColor(&Params{Token: token, Name: text})
	if err != nil {
		log.LOGGER("app").Warn("voice/command change_color error: text=%q err=%v", text, err)
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	log.LOGGER("app").Info("voice/command change_color success: text=%q items=%v", text, items)
	ctx.JSON(iris.Map{
		"result_code": 200,
		"result_msg":  "success",
		"data": voiceCmdResp{
			Intent: "change_color",
			Detail: items,
		},
	})
}

func execStartDelivery(ctx iris.Context, token, text string) {
	log.LOGGER("app").Info("voice/command start_delivery: text=%q", text)
	ctx.JSON(iris.Map{
		"result_code": 200,
		"result_msg":  "success",
		"data": voiceCmdResp{
			Intent: "start_delivery",
		},
	})
}

func execCollectColor(ctx iris.Context, token, text string) {
	items, err := parseColor(&Params{Token: token, Name: text})
	if err != nil {
		log.LOGGER("app").Warn("voice/command collect_color error: text=%q err=%v", text, err)
		ctx.JSON(iris.Map{"result_code": getErrCode(err), "result_msg": ErrMsg(ctx, err)})
		return
	}
	log.LOGGER("app").Info("voice/command collect_color success: text=%q items=%v", text, items)
	ctx.JSON(iris.Map{
		"result_code": 200,
		"result_msg":  "success",
		"data": voiceCmdResp{
			Intent: "collect_color",
			Detail: items,
		},
	})
}
