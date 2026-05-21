package handler

import (
	"bytes"
	"fmt"
	"html/template"
	"net/url"
	"strings"

	"nail/config"
	"nail/logger"

	"github.com/kataras/iris/v12"
)

const appleCallbackLogTag = "[Apple Callback]"

// appleCallbackHandler 处理 Apple Sign in with Apple（Web）form_post 回调。
// 仅返回 HTML 并通过 Deep Link 把授权参数交还 App，不调用用户登录、不写库、不限流。
func appleCallbackHandler(ctx iris.Context) {
	ip := GetClientIP(ctx)
	logger.Info("%s request ip=%s method=%s", appleCallbackLogTag, ip, ctx.Method())

	if err := ctx.Request().ParseForm(); err != nil {
		logger.Error("%s parse form failed ip=%s err=%v", appleCallbackLogTag, ip, err)
		ctx.StatusCode(iris.StatusBadRequest)
		writeAppleCallbackHTML(ctx, appleCallbackErrorPage("Invalid request"))
		return
	}

	oauthErr := strings.TrimSpace(ctx.FormValue("error"))
	code := strings.TrimSpace(ctx.FormValue("code"))
	idToken := strings.TrimSpace(ctx.FormValue("id_token"))
	state := strings.TrimSpace(ctx.FormValue("state"))
	hasUser := strings.TrimSpace(ctx.FormValue("user")) != ""

	logger.Info("%s received ip=%s code=%t id_token=%t state=%q user=%t error=%q",
		appleCallbackLogTag, ip, code != "", idToken != "", state, hasUser, oauthErr)

	if oauthErr != "" {
		logger.Warn("%s apple returned error ip=%s error=%q state=%q", appleCallbackLogTag, ip, oauthErr, state)
		writeAppleCallbackHTML(ctx, appleCallbackErrorPage("Authorization failed: "+oauthErr))
		return
	}

	if code == "" && idToken == "" {
		logger.Warn("%s missing code and id_token ip=%s state=%q", appleCallbackLogTag, ip, state)
		writeAppleCallbackHTML(ctx, appleCallbackErrorPage("Missing authorization data"))
		return
	}

	if idToken != "" && config.AppleSignInEnabled() {
		reqCtx := ctx.Request().Context()
		sub, email, vErr := validateAppleIdentityToken(reqCtx, idToken)
		if vErr != nil {
			detail := DiagnoseAppleIdentityToken(reqCtx, idToken)
			logger.Warn("%s identity token verify failed ip=%s state=%q detail=%s client_id=%v services_id=%v (still redirect to app)",
				appleCallbackLogTag, ip, state, detail, config.GetAppleClientIDs(), config.GetAppleServicesIDs())
		} else {
			logger.Info("%s identity token ok ip=%s sub=%s email=%t state=%q",
				appleCallbackLogTag, ip, sub, email != "", state)
		}
	} else if idToken != "" {
		logger.Info("%s skip token verify (apple client_id not configured) ip=%s state=%q",
			appleCallbackLogTag, ip, state)
	}

	deepLink, err := buildAppleDeepLink(code, idToken, state)
	if err != nil {
		logger.Error("%s deep link not configured ip=%s scheme=%q path=%q err=%v",
			appleCallbackLogTag, ip, config.GetAppleDeepLinkScheme(), config.GetAppleDeepLinkPath(), err)
		writeAppleCallbackHTML(ctx, appleCallbackErrorPage("Server redirect is not configured"))
		return
	}

	logger.Info("%s redirect ip=%s target=%s state=%q", appleCallbackLogTag, ip, appleDeepLinkLogSafe(deepLink), state)
	writeAppleCallbackHTML(ctx, appleCallbackSuccessPage(deepLink))
}

// appleDeepLinkLogSafe 日志用：仅保留 scheme/path/query 键名，不输出 token 取值。
func appleDeepLinkLogSafe(raw string) string {
	u, err := url.Parse(raw)
	if err != nil {
		return fmt.Sprintf("%s?…", strings.Split(raw, "?")[0])
	}
	keys := make([]string, 0, len(u.Query()))
	for k := range u.Query() {
		keys = append(keys, k)
	}
	return fmt.Sprintf("%s://%s?%s", u.Scheme, u.Host, strings.Join(keys, "&"))
}

func buildAppleDeepLink(code, idToken, state string) (string, error) {
	scheme := config.GetAppleDeepLinkScheme()
	path := config.GetAppleDeepLinkPath()
	if scheme == "" || path == "" {
		return "", deepLinkConfigError{}
	}

	q := url.Values{}
	if code != "" {
		q.Set("code", code)
	}
	if idToken != "" {
		q.Set("id_token", idToken)
	}
	if state != "" {
		q.Set("state", state)
	}

	u := &url.URL{
		Scheme:   scheme,
		Host:     path,
		RawQuery: q.Encode(),
	}
	return u.String(), nil
}

type deepLinkConfigError struct{}

func (deepLinkConfigError) Error() string { return "apple deep link not configured" }

type appleCallbackPage struct {
	DeepLink   template.URL
	DeepLinkJS template.JS
	Message    string
	IsError    bool
}

var appleCallbackTmpl = template.Must(template.New("apple_callback").Parse(`<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>{{if .IsError}}Authorization failed{{else}}Authorization success{{end}}</title>
  <style>
    body {
      font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
      display: flex;
      justify-content: center;
      align-items: center;
      min-height: 100vh;
      margin: 0;
      background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
      color: white;
    }
    .container { text-align: center; padding: 40px; max-width: 320px; }
    h1 { font-size: 22px; margin-bottom: 10px; }
    p { font-size: 14px; opacity: 0.9; line-height: 1.5; }
    .spinner {
      border: 3px solid rgba(255,255,255,0.3);
      border-top: 3px solid white;
      border-radius: 50%;
      width: 40px;
      height: 40px;
      animation: spin 1s linear infinite;
      margin: 20px auto;
    }
    @keyframes spin { 0% { transform: rotate(0deg); } 100% { transform: rotate(360deg); } }
  </style>
</head>
<body>
  <div class="container">
    {{if .IsError}}
    <h1>Authorization failed</h1>
    <p>{{.Message}}</p>
    <p>Please return to the app and try again.</p>
    {{else}}
    <div class="spinner"></div>
    <h1>Authorization success</h1>
    <p>Returning to the app...</p>
    <meta http-equiv="refresh" content="0;url={{.DeepLink}}">
    <script>window.location.replace({{.DeepLinkJS}});</script>
    {{end}}
  </div>
  {{if not .IsError}}
  <script>
    setTimeout(function() {
      document.querySelector('h1').textContent = 'Please open the app';
      document.querySelector('p').textContent = 'Authorization is complete. Return to the app.';
    }, 5000);
  </script>
  {{end}}
</body>
</html>`))

func appleCallbackSuccessPage(deepLink string) []byte {
	return renderAppleCallbackPage(appleCallbackPage{
		DeepLink:   template.URL(deepLink),
		DeepLinkJS: template.JS(deepLink),
		IsError:    false,
	})
}

func appleCallbackErrorPage(message string) []byte {
	return renderAppleCallbackPage(appleCallbackPage{
		Message: message,
		IsError: true,
	})
}

func renderAppleCallbackPage(data appleCallbackPage) []byte {
	var buf bytes.Buffer
	if err := appleCallbackTmpl.Execute(&buf, data); err != nil {
		logger.Error("%s render html failed err=%v", appleCallbackLogTag, err)
		return []byte("<!DOCTYPE html><html><body><p>Server error</p></body></html>")
	}
	return buf.Bytes()
}

func writeAppleCallbackHTML(ctx iris.Context, body []byte) {
	ctx.StatusCode(iris.StatusOK)
	ctx.Header("Content-Type", "text/html; charset=utf-8")
	_, _ = ctx.Write(body)
}
