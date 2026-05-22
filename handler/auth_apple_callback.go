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

	logger.Info("%s redirect ip=%s target=%s state=%q",
		appleCallbackLogTag, ip, appleDeepLinkLogSafe(deepLink), state)
	writeAppleCallbackHTML(ctx, appleCallbackSuccessPage(deepLink))
}

func buildAppleDeepLink(code, idToken, state string) (string, error) {
	scheme := config.GetAppleDeepLinkScheme()
	path := config.GetAppleDeepLinkPath()
	if scheme == "" || path == "" {
		return "", fmt.Errorf("apple deep link not configured")
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

func appleDeepLinkLogSafe(raw string) string {
	u, err := url.Parse(raw)
	if err != nil {
		return strings.Split(raw, "?")[0] + "?…"
	}
	keys := make([]string, 0, len(u.Query()))
	for k := range u.Query() {
		keys = append(keys, k)
	}
	return fmt.Sprintf("%s://%s?%s", u.Scheme, u.Host, strings.Join(keys, "&"))
}

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
  <title>{{if .IsError}}授权失败{{else}}授权成功{{end}}</title>
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
    .container { text-align: center; padding: 40px 24px; max-width: 360px; }
    h1 { font-size: 22px; margin-bottom: 12px; }
    p { font-size: 14px; opacity: 0.9; line-height: 1.6; margin: 8px 0; }
    .btn {
      display: inline-block;
      margin-top: 20px;
      padding: 14px 28px;
      background: #fff;
      color: #5b6ee1;
      font-size: 16px;
      font-weight: 600;
      text-decoration: none;
      border-radius: 999px;
      box-shadow: 0 4px 14px rgba(0,0,0,0.15);
    }
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
    <h1>授权失败</h1>
    <p>{{.Message}}</p>
    <p>请返回 App 重试。</p>
    {{else}}
    <div class="spinner" id="spin"></div>
    <h1 id="title">授权成功</h1>
    <p id="hint">正在打开 App…</p>
    <p>若未自动跳转，请点击下方按钮</p>
    <a class="btn" id="openApp" href="{{.DeepLink}}">打开 TintaShift App</a>
    {{end}}
  </div>
  {{if not .IsError}}
  <script>
    (function() {
      var link = {{.DeepLinkJS}};
      function openApp() {
        window.location.href = link;
      }
      openApp();
      setTimeout(openApp, 400);
      setTimeout(openApp, 1200);
      setTimeout(function() {
        var spin = document.getElementById('spin');
        if (spin) spin.style.display = 'none';
        document.getElementById('title').textContent = '请打开 App';
        document.getElementById('hint').textContent = '授权已完成，点击按钮返回 App';
      }, 3500);
    })();
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
