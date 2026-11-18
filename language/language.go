package language

import (
	"errors"
	"nail/config"
	"strings"
)

const defaultLocale = "en-US"

/*所有语言文案注册表：locale -> (key -> 文案)*/
var languageMaps = map[string]map[string]string{}

/*支持的 locale 列表（与 language.md 对应）*/
var supportedLocales = []string{
	"zh-CN", "en-US", "zh-Hant", "ja-JP", "ko-KR", "es-ES", "de-DE", "fr-FR",
	"it-IT", "pt-PT", "pl-PL", "ru-RU", "vi-VN", "th-TH", "ms-MY", "ar",
}

/*注册语言包*/
func registerLocale(locale string, m map[string]string) {
	languageMaps[locale] = m
}

/*SupportedLocales 返回支持的 locale 列表（与 language.md 对应）*/
func SupportedLocales() []string {
	return supportedLocales
}

/*初始化语言*/
func InitLanguage() {
	initLanguageZhcn()
	initLanguageEnus()
	initLanguageZhhant()
	initLanguageJajp()
	initLanguageKokr()
	initLanguageEses()
	initLanguageDede()
	initLanguageFrf()
	initLanguageItit()
	initLanguagePtpt()
	initLanguagePlpl()
	initLanguageRuru()
	initLanguageVivn()
	initLanguageThth()
	initLanguageMsmy()
	initLanguageAr()
}

// NormalizeLocale 将客户端语言归一为受支持的 locale, 无法识别时返回 en-US。
func NormalizeLocale(lang string) string {
	lang = strings.ToLower(strings.TrimSpace(lang))
	if lang == "" {
		return defaultLocale
	}
	/*精确匹配（大小写不敏感）*/
	for _, loc := range supportedLocales {
		if strings.ToLower(loc) == lang {
			return loc
		}
	}
	/*繁体中文别名*/
	if strings.HasPrefix(lang, "zh-hant") || strings.HasPrefix(lang, "zh-tw") ||
		strings.HasPrefix(lang, "zh-hk") || strings.HasPrefix(lang, "zh-mo") {
		return "zh-Hant"
	}
	/*按主语言子标签匹配：ja -> ja-JP、pt -> pt-PT 等*/
	primary := strings.Split(lang, "-")[0]
	for _, loc := range supportedLocales {
		if strings.Split(strings.ToLower(loc), "-")[0] == primary {
			return loc
		}
	}
	return defaultLocale
}

// GetRawMessageFor 按 locale 返回文案；无对应 key 时回退默认语言，仍无则返回 key 本身。
func GetRawMessageFor(locale, id string) string {
	locale = NormalizeLocale(locale)
	if m, ok := languageMaps[locale]; ok {
		if msg := m[id]; msg != "" {
			return msg
		}
	}
	if m, ok := languageMaps[defaultLocale]; ok {
		if msg := m[id]; msg != "" {
			return msg
		}
	}
	return id
}

/*错误信息返回*/
func GetLanguageMessage(id string) error {
	return errors.New(GetRawMessageFor(config.GetLanguage(), id))
}

/*字符串返回，若 key 无对应翻译则返回 key 本身*/
func GetRawMessage(id string) string {
	return GetRawMessageFor(config.GetLanguage(), id)
}
