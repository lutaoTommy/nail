package handler

import (
	"strings"

	"nail/language"

	"github.com/kataras/iris/v12"
)

const ctxLocaleKey = "locale"

// HeaderLanguage 客户端统一语言头（推荐 zh-CN / en-US）。
const HeaderLanguage = "X-Language"

// LocaleMiddleware 从 X-Language 解析语言并写入请求上下文。
func LocaleMiddleware(ctx iris.Context) {
	locale := language.NormalizeLocale(strings.TrimSpace(ctx.GetHeader(HeaderLanguage)))
	ctx.Values().Set(ctxLocaleKey, locale)
	ctx.Next()
}

// LocaleFromCtx 返回当前请求语言（zh-CN 或 en-US）。
func LocaleFromCtx(ctx iris.Context) string {
	if v, ok := ctx.Values().Get(ctxLocaleKey).(string); ok && v != "" {
		return language.NormalizeLocale(v)
	}
	return language.NormalizeLocale("")
}

// Msg 按请求语言返回文案。
func Msg(ctx iris.Context, key string) string {
	return language.GetRawMessageFor(LocaleFromCtx(ctx), key)
}
