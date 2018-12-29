package v1

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/haibeichina/gin_video_server/0commonpkg/e"
	"github.com/haibeichina/gin_video_server/scheduler/pkg/app"
	"github.com/haibeichina/gin_video_server/scheduler/pkg/logging"
	"github.com/haibeichina/gin_video_server/scheduler/service/delrecsservice"
)

// AddDelRec 新增删除记录
// @Summary add a new record for delete video
// @Accept json
// @Produce json
// @Param vid-id path string true "video id"
// @Success 200 {string} json "{"code":200,"msg":"ok","data":nil}"
// @Failure 400 {string} json "{"code":400,"msg":"","data":nil}"
// @Failure 500 {string} json "{"code":500,"msg":"","data":nil}"
// @Router /api/v1/user/{username} [get]
func AddDelRec(c *gin.Context) {
	appG := app.Gin{C: c}

	vid := c.Param("vid-id")

	valid := validation.Validation{}
	valid.Required(vid, "vid-id").Message("视频id不能为空")

	if valid.HasErrors() {
		// app.MarkErrors(valid.Errors)
		logging.Info(valid.Errors)
		appG.ResponseErr(http.StatusBadRequest, e.INVALID_PARAMS)
		return
	}

	commentSer := delrecsservice.DelRecs{
		VideoID: vid,
	}

	if err := commentSer.AddDelRec(); err != nil {
		logging.Info(err)
		appG.ResponseErr(http.StatusInternalServerError, e.ERROR_ADD_DELRECS_FAIL)
		return
	}

	appG.ResponseOk(e.SUCCESS)
}
