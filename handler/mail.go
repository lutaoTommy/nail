package handler

import (
	"fmt"

	"nail/config"
	"nail/language"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dm20151123 "github.com/alibabacloud-go/dm-20151123/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

/*使用 AK&SK 初始化账号 Client，账号参数从 config.ini [mail] 读取*/
func createClient() (result *dm20151123.Client, err error) {
	cfg := &openapi.Config{
		AccessKeyId:     tea.String(config.GetMailAccessKeyId()),
		AccessKeySecret: tea.String(config.GetMailAccessKeySecret()),
	}
	cfg.Endpoint = tea.String(config.GetMailEndpoint())
	result, err = dm20151123.NewClient(cfg)
	return result, err
}

/*发送单个邮件*/
func sendMail(user *User) (err error) {
	client, err := createClient()
	if err != nil {
		return err
	}
	subject := language.GetRawMessage("MAIL_VERIFY_SUBJECT")
	htmlBody := fmt.Sprintf(`<!DOCTYPE html>
<html>
<body style="font-family: Arial, sans-serif;">
    <h2>%s</h2>
    <p>%s</p>
    <p style="font-size: 32px; font-weight: bold; color: #1a73e8; text-align: center; letter-spacing: 5px;">%s</p>
    <p><strong>%s</strong>%s</p>
    <hr>
    <p style="color: #999; font-size: 12px;">%s</p>
    <p style="color: #999; font-size: 12px;">%s</p>
</body>
</html>`,
		language.GetRawMessage("MAIL_VERIFY_TITLE"),
		language.GetRawMessage("MAIL_VERIFY_DESC"),
		user.Cert,
		language.GetRawMessage("MAIL_VERIFY_SECURITY_TITLE"),
		language.GetRawMessage("MAIL_VERIFY_SECURITY_DESC"),
		language.GetRawMessage("MAIL_VERIFY_IGNORE"),
		language.GetRawMessage("MAIL_COPYRIGHT"),
	)
	params := &dm20151123.SingleSendMailRequest{
		AccountName:    tea.String(config.GetMailAccountName()),
		FromAlias:      tea.String(config.GetMailFromAlias()),
		AddressType:    tea.Int32(1),
		ReplyToAddress: tea.Bool(false),
		ToAddress:      tea.String(user.Email),
		Subject:        tea.String(subject),
		HtmlBody:       tea.String(htmlBody),
	}
	tryErr := func() (e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				e = r
			}
		}()
		_, err = client.SingleSendMail(params)
		return err
	}()
	return tryErr
}
