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
	if config.GetLanguage() == "en-US" {
		return errors.New(languageMapEn[id])
	} else {
		return errors.New(languageMapZhcn[id])
	}
}

/*字符串返回，若 key 无对应翻译则返回 key 本身*/
func GetRawMessage(id string) string {
	var msg string
	if config.GetLanguage() == "en-US" {
		msg = languageMapEn[id]
	} else {
		msg = languageMapZhcn[id]
	}
	if msg == "" {
		return id
	}
	return msg
}
