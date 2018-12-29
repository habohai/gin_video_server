package limiter

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/haibeichina/gin_video_server/0commonpkg/e"
	"github.com/haibeichina/gin_video_server/streamserver/pkg/app"
	"github.com/haibeichina/gin_video_server/streamserver/pkg/logging"
	"github.com/haibeichina/gin_video_server/streamserver/pkg/setting"
)

// Limiter 中间件
func Limiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}

		l := NewConnLimiter(setting.LimiterSetting.MaxLinkNum)
		if ok := l.GetConn(); !ok {
			logging.Warn("Too many requests")
			appG.Response(http.StatusTooManyRequests, e.ERROR_AUTH_INVALID_SESSION, "Too many requests")
			c.Abort()
			return
		}

		c.Next()

		l.ReleaseConn()
	}
}
