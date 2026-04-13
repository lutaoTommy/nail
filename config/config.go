package config

import (
	"github.com/Unknwon/goconfig"
)

/*读取配置文件*/
var localConfig Config

func LoadConfig() error {
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil {
		return err
	}
	/*服务器IP或域名*/
	localConfig.Domain, err = cfg.GetValue("server", "domain")
	if err != nil {
		return err
	}
	/*http端口*/
	localConfig.HttpPort, err = cfg.Int("server", "http_port")
	if err != nil {
		return err
	}
	/*mysql链接*/
	localConfig.MysqlUrl, err = cfg.GetValue("database", "mysql_url")
	if err != nil {
		return err
	}
	/*OSS*/
	localConfig.OssEndpoint, _ = cfg.GetValue("oss", "oss_endpoint")
	localConfig.OssAccessKeyId, _ = cfg.GetValue("oss", "oss_access_key_id")
	localConfig.OssAccessKeySecret, _ = cfg.GetValue("oss", "oss_access_key_secret")
	localConfig.OssBucket, _ = cfg.GetValue("oss", "oss_bucket")
	localConfig.OssDomain, _ = cfg.GetValue("oss", "oss_domain")
	/*邮件*/
	localConfig.MailEndpoint, _ = cfg.GetValue("mail", "mail_endpoint")
	localConfig.MailAccessKeyId, _ = cfg.GetValue("mail", "mail_access_key_id")
	localConfig.MailAccessKeySecret, _ = cfg.GetValue("mail", "mail_access_key_secret")
	localConfig.MailAccountName, _ = cfg.GetValue("mail", "mail_account_name")
	localConfig.MailFromAlias, _ = cfg.GetValue("mail", "mail_from_alias")
	/*阿里云图像识别*/
	localConfig.AliAccessKeyId, _ = cfg.GetValue("ali", "ali_access_key_id")
	localConfig.AliAccessKeySecret, _ = cfg.GetValue("ali", "ali_access_key_secret")
	localConfig.AliImagerecogEndpoint, _ = cfg.GetValue("ali", "ali_imagerecog_endpoint")
	/*微信小程序*/
	localConfig.WxAppId, _ = cfg.GetValue("wx", "wx_app_id")
	localConfig.WxAppSecret, _ = cfg.GetValue("wx", "wx_app_secret")
	return nil
}

/*获取配置*/
func GetDomain() string {
	return localConfig.Domain
}

/*获取配置*/
func GetHttpPort() int {
	return localConfig.HttpPort
}

/*获取配置*/
func GetLanguage() string {
	return localConfig.Language
}

/*获取配置*/
func GetMysqlUrl() string {
	return localConfig.MysqlUrl
}

/*获取 OSS 配置*/
func GetOssEndpoint() string        { return localConfig.OssEndpoint }
func GetOssAccessKeyId() string     { return localConfig.OssAccessKeyId }
func GetOssAccessKeySecret() string { return localConfig.OssAccessKeySecret }
func GetOssBucket() string          { return localConfig.OssBucket }
func GetOssDomain() string          { return localConfig.OssDomain }

/*获取邮件配置*/
func GetMailEndpoint() string        { return localConfig.MailEndpoint }
func GetMailAccessKeyId() string     { return localConfig.MailAccessKeyId }
func GetMailAccessKeySecret() string { return localConfig.MailAccessKeySecret }
func GetMailAccountName() string     { return localConfig.MailAccountName }
func GetMailFromAlias() string       { return localConfig.MailFromAlias }

/*获取阿里云图像识别配置*/
func GetAliAccessKeyId() string        { return localConfig.AliAccessKeyId }
func GetAliAccessKeySecret() string    { return localConfig.AliAccessKeySecret }
func GetAliImagerecogEndpoint() string { return localConfig.AliImagerecogEndpoint }

/*获取微信小程序配置*/
func GetWxAppId() string     { return localConfig.WxAppId }
func GetWxAppSecret() string { return localConfig.WxAppSecret }
