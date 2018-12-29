package validatesession

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/haibeichina/gin_video_server/0commonpkg/e"
	"github.com/haibeichina/gin_video_server/apiserver/pkg/app"
	"github.com/haibeichina/gin_video_server/apiserver/pkg/logging"
	"github.com/haibeichina/gin_video_server/apiserver/pkg/setting"
	"github.com/haibeichina/gin_video_server/apiserver/routers/api"
)

// ValidateSession 认证session
func ValidateSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := app.Gin{C: c}

		sid := c.Request.Header.Get(setting.HTTPHeadParamSetting.HeadweFieldSession)

		valid := validation.Validation{}
		valid.Required(sid, setting.HTTPHeadParamSetting.HeadweFieldSession).Message("session不能为空")

		if valid.HasErrors() {
			//app.MarkErrors(valid.Errors)
			logging.Info(valid.Errors)
			appG.ResponseErr(http.StatusBadRequest, e.INVALID_PARAMS)
			c.Abort()
			return
		}

		uname, ok := api.IsSessionExpired(sid)
		if ok {
			appG.ResponseErr(http.StatusUnauthorized, e.ERROR_AUTH_INVALID_SESSION)
			c.Abort()
			return
		}

		c.Request.Header.Add(setting.HTTPHeadParamSetting.HeadweFieldUname, uname)
		c.Next()
	}
}
