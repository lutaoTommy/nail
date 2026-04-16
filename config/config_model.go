package config

/*配置文件*/
type Config struct {
	Domain   string `json:"domain"`
	HttpPort int    `json:"http_port"`
	MysqlUrl string `json:"mysql_url"`
	Language string `json:"language"`
	/*OSS 阿里云对象存储*/
	OssEndpoint        string `json:"oss_endpoint"`
	OssAccessKeyId     string `json:"oss_access_key_id"`
	OssAccessKeySecret string `json:"oss_access_key_secret"`
	OssBucket          string `json:"oss_bucket"`
	/*邮件（阿里云 DirectMail）*/
	MailEndpoint        string `json:"mail_endpoint"` /*如 dm.aliyuncs.com */
	MailAccessKeyId     string `json:"mail_access_key_id"`
	MailAccessKeySecret string `json:"mail_access_key_secret"`
	MailAccountName     string `json:"mail_account_name"` /*发信地址，如 register@mail.smartepapersystem.com */
	MailFromAlias       string `json:"mail_from_alias"`   /*发信人显示名（FromAlias）*/
}
