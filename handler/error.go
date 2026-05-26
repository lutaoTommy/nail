package handler

import (
	"fmt"
	"strings"

	"github.com/kataras/iris/v12"
)

/*自定义错误*/
type HttpError struct {
	Code    int
	Temp    string
	Message string // 仅兼容旧数据；新错误请用 Temp + Suffix，由 ErrMsg 翻译
	Suffix  string // 追加在翻译文案后，如敏感词列表
}

/*自定义错误*/
func (e HttpError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	if e.Temp != "" {
		return e.Temp
	}
	return "error"
}

/*自定义错误*/
func newErr(temp string) HttpError {
	return HttpError{
		Code: 500,
		Temp: temp,
	}
}

/*自定义错误*/
func newError(code int, temp string) HttpError {
	return HttpError{
		Code: code,
		Temp: temp,
	}
}

func newSensitiveWordError(words []string) HttpError {
	return HttpError{
		Code:   400,
		Temp:   "E_SENSITIVEWORD",
		Suffix: fmt.Sprintf("(%s)", strings.Join(words, ",")),
	}
}

// ErrMsg 按请求语言输出错误文案。
func ErrMsg(ctx iris.Context, err error) string {
	if e, ok := err.(HttpError); ok {
		if e.Temp != "" {
			msg := Msg(ctx, e.Temp)
			if e.Suffix != "" {
				msg += e.Suffix
			}
			return msg
		}
		if e.Message != "" {
			return e.Message
		}
	}
	return err.Error()
}

/*自定义错误*/
func getErrCode(err error) int {
	e, ok := err.(HttpError)
	if ok {
		return e.Code
	}
	return 500
}
