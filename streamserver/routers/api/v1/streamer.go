package v1

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"github.com/haibeichina/gin_video_server/0commonpkg/e"
	"github.com/haibeichina/gin_video_server/streamserver/pkg/app"
	"github.com/haibeichina/gin_video_server/streamserver/pkg/logging"
	"github.com/haibeichina/gin_video_server/streamserver/pkg/upload"
)

// TestPage 测试
func TestPage(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", nil)
}

// Stream 视频流服务
// @Summary a video stream server
// @Accept json
// @Produce json
// @Param vid-id path string true "video id"
// @Success 200 {stream} json "stream"
// @Failure 400 {string} json ""
// @Failure 500 {string} json ""
// @Router /api/v1/videos/{vid-id} [get]
func Stream(c *gin.Context) {
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

	filePath := upload.GetVideoFullPath()

	if err := upload.CheckVideo(filePath); err != nil {
		logging.Info(err)
		appG.ResponseErr(http.StatusInternalServerError, e.ERROR_NOT_EXIST_VIDEO)
		return
	}

	ok, err := upload.CheckVideoExist(filePath + vid)
	if err != nil {
		logging.Info(err)
		appG.ResponseErr(http.StatusInternalServerError, e.ERROR)
	}
	if !ok {
		logging.Info(ok)
		appG.ResponseErr(http.StatusNotFound, e.ERROR_NOT_EXIST_VIDEO)
		return
	}
	c.File(filePath + vid)

	// F, err := file.IsToOpen(vid, filePath)
	// if err != nil {
	// 	logging.Info(err)
	// 	appG.ResponseErr(http.StatusInternalServerError, e.ERROR_OPEN_FILE_FAIL)
	// 	return
	// }
	// defer F.Close()
	// c.Header("Content-Type", "video/mp4")
	// http.ServeContent(c.Writer, c.Request, "", time.Now(), F)
}

// Upload 视频上传服务
// @Summary upload video file server
// @Accept json
// @Produce json
// @Param vid-id path string true "video id"
// @Param file formData file true "video file"
// @Success 200 {stream} json "stream"
// @Failure 400 {string} json ""
// @Failure 500 {string} json ""
// @Router /api/v1/upload/{vid-id} [post]
func Upload(c *gin.Context) {
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

	file, video, err := c.Request.FormFile("file")
	if err != nil {
		logging.Errorf("Error when try to get file: %v", err)
		appG.ResponseErr(http.StatusInternalServerError, e.ERROR)
		return
	}

	if video == nil {
		appG.ResponseErr(http.StatusInternalServerError, e.ERROR)
	}

	videoName := video.Filename
	fullPath := upload.GetVideoFullPath()
	src := fullPath + vid

	if ok := upload.CheckVideoExt(videoName); !ok {
		appG.ResponseErr(http.StatusBadRequest, e.ERROR_UPLOAD_VIDEO_FORMAT_UNSUPPORTED)
		return
	}

	if ok := upload.CheckVideoSize(file); !ok {
		logging.Warn(err)
		appG.ResponseErr(http.StatusBadRequest, e.ERROR_UPLOAD_VIDEO_FILE_TOO_BIG)
		return
	}

	if err := upload.CheckVideo(fullPath); err != nil {
		logging.Warn(err)
		appG.ResponseErr(http.StatusInternalServerError, e.ERROR)
		return
	}

	if err := c.SaveUploadedFile(video, src); err != nil {
		logging.Warn(err)
		appG.ResponseErr(http.StatusInternalServerError, e.ERROR_UPLOAD_VIDEO_SAVE_FAIL)
		return
	}

	appG.ResponseOk(e.SUCCESS)
}
