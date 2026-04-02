package handler

import (
	"fmt"
	"strings"

	"github.com/kataras/iris/v12"
)

/*BuildPublicURL 根据请求生成对外可访问的完整 URL。
 * 用于反向代理场景（如外部 443 -> 本项目 8008），优先使用 X-Forwarded-Proto / X-Forwarded-Host，
 * 使返回的链接为外部可直连的 https://域名/... 形式。
 * path 为资源路径，如 "/apk/app.apk"、"/circle/xxx.png"，需以 / 开头。
 */
func BuildPublicURL(ctx iris.Context, path string) string {
	scheme := strings.TrimSpace(ctx.Request().Header.Get("X-Forwarded-Proto"))
	if scheme == "" {
		if p := strings.TrimSpace(ctx.Request().Header.Get("X-Forwarded-Port")); p == "443" {
			scheme = "https"
		} else {
			scheme = "http"
		}
	}
	host := strings.TrimSpace(ctx.Request().Header.Get("X-Forwarded-Host"))
	if host == "" {
		host = strings.TrimSpace(ctx.Request().Host)
	}
	host = strings.TrimSuffix(host, ":443")
	host = strings.TrimSuffix(host, ":80")
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return fmt.Sprintf("%s://%s%s", scheme, host, path)
}
