package handler

import (
	"nail/config"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	imagerecog20190930 "github.com/alibabacloud-go/imagerecog-20190930/v2/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

/*使用凭据初始化账号 Client，账号参数从 config.ini [ali] 读取*/
func CreateClient() (_result *imagerecog20190930.Client, _err error) {
	cfg := &openapi.Config{
		AccessKeyId:     tea.String(config.GetAliAccessKeyId()),
		AccessKeySecret: tea.String(config.GetAliAccessKeySecret()),
	}
	cfg.Endpoint = tea.String(config.GetAliImagerecogEndpoint())
	_result, _err = imagerecog20190930.NewClient(cfg)
	return _result, _err
}

func imageRecog(url string) (bodyData *imagerecog20190930.RecognizeImageColorResponseBodyData, _err error) {
	client, _err := CreateClient()
	if _err != nil {
		return bodyData, _err
	}

	recognizeImageColorRequest := &imagerecog20190930.RecognizeImageColorRequest{
		Url:        tea.String(url),
		ColorCount: tea.Int32(6),
	}
	response := &imagerecog20190930.RecognizeImageColorResponse{}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		response, _err = client.RecognizeImageColorWithOptions(recognizeImageColorRequest, runtime)
		if _err != nil {
			return _err
		}
		return nil
	}()

	if tryErr != nil {
		return bodyData, tryErr
	}

	return response.Body.Data, _err
}
