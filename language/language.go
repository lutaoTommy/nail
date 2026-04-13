package language

import (
	"errors"
	"nail/config"
)

/*初始化语言*/
func InitLanguage() {
	initLanguageZhcn()
	initLanguageEnus()
}

/*错误信息返回*/
func GetLanguageMessage(id string) error {
	if config.GetLanguage() == "zh-CN" {
		return errors.New(languageMapZhcn[id])
	} else {
		return errors.New(languageMapEn[id])
	}
}

/*字符串返回，若 key 无对应翻译则返回 key 本身*/
func GetRawMessage(id string) string {
	var msg string
	if config.GetLanguage() == "zh-CN" {
		msg = languageMapZhcn[id]
	} else {
		msg = languageMapEn[id]
	}
	if msg == "" {
		return id
	}
	return msg
}
