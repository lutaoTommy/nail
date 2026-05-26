package config

/*配置文件*/
type Config struct {
	Area     string `json:"area"`
	Domain   string `json:"domain"`
	Version  string `json:"version"`
	HttpPort int    `json:"http_port"`
	MysqlUrl string `json:"mysql_url"`
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
	/*自定义社交登录签名密钥（HMAC）*/
	SocialAppSecret string `json:"-"`
	/*Google Sign-In：access_token 对应 OAuth Client ID 列表（tokeninfo aud/azp），逗号分隔多个*/
	GoogleOAuthClientIDs []string `json:"-"`
	/*Sign in with Apple*/
	AppleTeamID      string   `json:"-"`
	AppleKeyID       string   `json:"-"`
	AppleAuthKeyPath string   `json:"-"`
	AppleClientIDs      []string `json:"-"` // 原生 App Bundle ID（identity token 的 aud）
	AppleServicesIDs    []string `json:"-"` // Web OAuth Services ID（浏览器回调 id_token 的 aud）
	AppleDeepLinkScheme string   `json:"-"` // OAuth 回调后跳回 App，如 tintashift
	AppleDeepLinkPath   string   `json:"-"` // Deep Link 路径段，如 apple-login → tintashift://apple-login
}
