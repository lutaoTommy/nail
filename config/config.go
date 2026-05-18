package config

import (
	"strings"

	"github.com/Unknwon/goconfig"
)

/*读取配置文件*/
var localConfig Config

func LoadConfig() error {
	cfg, err := goconfig.LoadConfigFile("config.ini")
	if err != nil {
		return err
	}
	/*服务器区域*/
	localConfig.Area, err = cfg.GetValue("server", "area")
	if err != nil {
		return err
	}
	/*服务器IP或域名*/
	localConfig.Domain, err = cfg.GetValue("server", "domain")
	if err != nil {
		return err
	}
	/*服务器版本*/
	localConfig.Version, err = cfg.GetValue("server", "version")
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
	/*邮件*/
	localConfig.MailEndpoint, _ = cfg.GetValue("mail", "mail_endpoint")
	localConfig.MailAccessKeyId, _ = cfg.GetValue("mail", "mail_access_key_id")
	localConfig.MailAccessKeySecret, _ = cfg.GetValue("mail", "mail_access_key_secret")
	localConfig.MailAccountName, _ = cfg.GetValue("mail", "mail_account_name")
	localConfig.MailFromAlias, _ = cfg.GetValue("mail", "mail_from_alias")
	/*Google OAuth Client ID（可配置多个，逗号分隔）*/
	localConfig.GoogleOAuthClientIDs = nil
	if raw, err := cfg.GetValue("google", "oauth_client_ids"); err == nil && raw != "" {
		for _, part := range strings.Split(raw, ",") {
			if s := strings.TrimSpace(part); s != "" {
				localConfig.GoogleOAuthClientIDs = append(localConfig.GoogleOAuthClientIDs, s)
			}
		}
	}
	return nil
}

/*获取配置*/
func GetArea() string {
	return localConfig.Area
}

/*获取配置*/
func GetDomain() string {
	return localConfig.Domain
}

/*获取配置*/
func GetVersion() string {
	return localConfig.Version
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

/*获取邮件配置*/
func GetMailEndpoint() string        { return localConfig.MailEndpoint }
func GetMailAccessKeyId() string     { return localConfig.MailAccessKeyId }
func GetMailAccessKeySecret() string { return localConfig.MailAccessKeySecret }
func GetMailAccountName() string     { return localConfig.MailAccountName }
func GetMailFromAlias() string       { return localConfig.MailFromAlias }

// GetGoogleOAuthClientIDs 返回允许的 Google ID Token aud（Client ID）列表；为空表示未启用服务端校验配置。
func GetGoogleOAuthClientIDs() []string {
	return localConfig.GoogleOAuthClientIDs
}
