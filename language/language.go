package language

import "strings"

const defaultLocale = "en-US"

/*初始化语言*/
func InitLanguage() {
	initLanguageZhcn()
	initLanguageEnus()
}

// NormalizeLocale 将客户端语言归一为 zh-CN 或 en-US。
func NormalizeLocale(lang string) string {
	lang = strings.ToLower(strings.TrimSpace(lang))
	if strings.HasPrefix(lang, "zh") {
		return "zh-CN"
	}
	return "en-US"
}

/*字符串返回*/
func GetRawMessage(id string) string {
	return GetRawMessageFor(defaultLocale, id)
}

// GetRawMessageFor 按 locale 返回文案。
func GetRawMessageFor(locale, id string) string {
	var msg string
	if NormalizeLocale(locale) == "zh-CN" {
		msg = languageMapZhcn[id]
	} else {
		msg = languageMapEn[id]
	}
	if msg == "" {
		return id
	}
	return msg
}
