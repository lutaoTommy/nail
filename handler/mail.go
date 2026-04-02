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
	name := language.GetRawMessage("VERIFICATION_CODE")
	params := &dm20151123.SingleSendMailRequest{
		AccountName:    tea.String(config.GetMailAccountName()),
		AddressType:    tea.Int32(1),
		ReplyToAddress: tea.Bool(false),
		ToAddress:      tea.String(user.Email),
		Subject:        tea.String(name),
		TextBody:       tea.String(fmt.Sprintf("%s: %s %s", name, user.Cert, language.GetRawMessage("DO_NOT_TELL"))),
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
