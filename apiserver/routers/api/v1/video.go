package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	// "github.com/haibeichina/gin_video_server/0commonpkg/comdefs"
	"github.com/haibeichina/gin_video_server/0commonpkg/e"

	"github.com/haibeichina/gin_video_server/apiserver/pkg/app"
	"github.com/haibeichina/gin_video_server/apiserver/pkg/defs"
	"github.com/haibeichina/gin_video_server/apiserver/pkg/logging"
	"github.com/haibeichina/gin_video_server/apiserver/pkg/setting"
	"github.com/haibeichina/gin_video_server/apiserver/pkg/utils"
	"github.com/haibeichina/gin_video_server/apiserver/service/videoservice"
)

// GetUserVideos 获取多个视频信息
// @Summary get user videoinfos
// @Accept json
// @Produce json
// @Param username path string true "username"
// @Success 200 {string} json "{"code":200,"msg":"ok","data":defs.ResVideoInfos}"
// @Failure 400 {string} json "{"code":400,"msg":"请求参数错误","data":nil}"
// @Failure 500 {string} json "{"code":500,"msg":"数据库错误","data":nil}"
// @Router /api/v1/user/{username}/videos [get]
func GetUserVideos(c *gin.Context) {
	appG := app.Gin{C: c}

	uname := c.Param("username")
	valid := validation.Validation{}
	valid.Required(uname, "user_name").Message("用户名不能为空")

	if valid.HasErrors() {
		// app.MarkErrors(valid.Errors)
		logging.Info(valid.Errors)
		appG.ResponseErr(http.StatusBadRequest, e.INVALID_PARAMS)
		return
	}

	videoInfoSer := videoservice.VideoInfo{
		UserName: uname,
		PageNum:  0,
		PageSize: setting.AppSetting.PageSize,
	}

	videoInfos, err := videoInfoSer.GetUserVideos()
	if err != nil {
		logging.Info(err)
		appG.ResponseErr(http.StatusInternalServerError, e.ERROR_GET_VIDEOINFOS_FAIL)
		return
	}

	count, err := videoInfoSer.Count()
	if err != nil {
		appG.ResponseErr(http.StatusInternalServerError, e.ERROR_COUNT_VIDEOINFO_FAIL)
		return
	}

	appG.ResponseOk(defs.ResVideoInfos{
		VideoInfos: videoInfos,
		Count:      count,
	})
}

// AddVideo 新增视频
// @Summary add a new video
// @Accept json
// @Produce json
// @Param username path string true "username"
// @Param nvbody body defs.ReqNewVideo true "new video info"
// @Success 200 {string} json "{"code":200,"msg":"ok","data":defs.ResAddVideoInfo}"
// @Failure 400 {string} json "{"code":400,"msg":"","data":nil}"
// @Failure 500 {string} json "{"code":500,"msg":"","data":nil}"
// @Router /api/v1/user/{username}/videos [post]
func AddVideo(c *gin.Context) {
	appG := app.Gin{C: c}

	res, _ := ioutil.ReadAll(c.Request.Body)
	nvbody := &defs.ReqNewVideo{}
	if err := json.Unmarshal(res, nvbody); err != nil {
		logging.Info(err)
		appG.ResponseErr(http.StatusBadRequest, e.ERROR_REQUEST_BODY_UNMARSHAL_JSON_FAIL)
		return
	}

	uid := nvbody.UserID
	uname := c.Param("username")
	vname := nvbody.VideoName

	valid := validation.Validation{}
	valid.Min(uid, 1, "user_id").Message("用户ID必须大于0")
	valid.Required(uname, "user_name").Message("用户名不能为空")
	valid.Required(vname, "video_name").Message("视频名不能为空")

	if valid.HasErrors() {
		// app.MarkErrors(valid.Errors)
		logging.Info(valid.Errors)
		appG.ResponseErr(http.StatusBadRequest, e.INVALID_PARAMS)
		return
	}
	videoInfoSer := videoservice.VideoInfo{
		UserID:    uid,
		UserName:  uname,
		VideoName: vname,
	}

	video, err := videoInfoSer.AddVideo()
	if err != nil {
		logging.Info(err)
		appG.ResponseErr(http.StatusInternalServerError, e.ERROR_ADD_VIDEOINFO_FAIL)
		return
	}

	appG.ResponseOk(defs.ResAddVideoInfo{
		ID: video.ID,
	})
}

// DeleteVideo 删除指定的视频
// @Summary delete video by video_id
// @Accept json
// @Produce json
// @Param username path string true "username"
// @Param vid-id path string true "video id"
// @Success 200 {string} json "{"code":200,"msg":"ok","data":{}}"
// @Failure 400 {string} json "{"code":400,"msg":"请求参数错误","data":{}}"
// @Router /api/v1/user/{username}/videos/{vid-id} [post]
func DeleteVideo(c *gin.Context) {
	appG := app.Gin{C: c}

	uname := c.Param("username")
	vid := c.Param("vid-id")

	valid := validation.Validation{}
	valid.Required(uname, "username").Message("用户名不能为空")
	valid.Required(vid, "vid-id").Message("视频ID不能为空")

	if valid.HasErrors() {
		// app.MarkErrors(valid.Errors)
		logging.Info(valid.Errors)
		appG.ResponseErr(http.StatusBadRequest, e.INVALID_PARAMS)
		return
	}

	if err := utils.SendDeleteVideoRequest(vid); err != nil {
		appG.ResponseErr(http.StatusInternalServerError, e.ERROR_DELETE_VIDEOINFO_FAIL)
		return
	}

	videoInfoSer := videoservice.VideoInfo{
		ID:       vid,
		UserName: uname,
	}

	ei, err := videoInfoSer.DeleteVideo()
	if err != nil {
		logging.Info(err)
		appG.ResponseErr(http.StatusInternalServerError, ei)
		return
	}

	appG.ResponseOk(nil)
}
