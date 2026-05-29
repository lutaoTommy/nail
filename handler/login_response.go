package handler

import (
	"fmt"

	"github.com/kataras/iris/v12"
)

func msgWithRetries(ctx iris.Context, base string, remaining int) string {
	return base + fmt.Sprintf(Msg(ctx, "E_RETRIES_REMAINING"), remaining)
}

func loginFailureResponse(ctx iris.Context, err error, ip, accKey string) iris.Map {
	remaining := RecordLoginFailure(ip, accKey)
	return iris.Map{
		"result_code":        getErrCode(err),
		"result_msg":         msgWithRetries(ctx, ErrMsg(ctx, err), remaining),
		"remaining_attempts": remaining,
	}
}

func loginLockedResponse(ctx iris.Context, retrySec int) iris.Map {
	res := iris.Map{
		"result_code":        429,
		"remaining_attempts": 0,
	}
	if retrySec > 0 {
		res["result_msg"] = fmt.Sprintf(Msg(ctx, "E_ACCOUNT_LOCKED_RETRY"), retrySec)
		res["retry_after"] = retrySec
	} else {
		res["result_msg"] = Msg(ctx, "E_ACCOUNT_LOCKED")
	}
	return res
}
