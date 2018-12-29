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
	"github.com/haibeichina/gin_video_server/apiserver/service/commentservice"
)

// GetComments 获取多个视频评论
// @Summary get comments by video id
// @Accept json
// @Produce json
// @Param vid-id path string true "video id"
// @Success 200 {string} json "{"code":200,"msg":"ok","data":defs.ResComments}"
// @Failure 400 {string} json "{"code":400,"msg":"请求参数错误","data":nil}"
// @Failure 500 {string} json "{"code":500,"msg":"","data":nil}"
// @Router /api/v1/videos/{vid-id}/comments [get]
func GetComments(c *gin.Context) {
	appG := app.Gin{C: c}

	vid := c.Param("vid-id")

	valid := validation.Validation{}
	valid.Required(vid, "vid-id").Message("视频id不能为空")

	if valid.HasErrors() {
		//app.MarkErrors(valid.Errors)
		logging.Info(valid.Errors)
		appG.ResponseErr(http.StatusBadRequest, e.INVALID_PARAMS)
		return
	}

	commentSer := commentservice.Comment{
		VideoID: vid,
		//PageNum:  utils.GetPage(c),
		PageNum:  0,
		PageSize: setting.AppSetting.PageSize,
	}
	logging.Info(1)
	comments, err := commentSer.GetComments()
	if err != nil {
		logging.Info(2)
		appG.ResponseErr(http.StatusInternalServerError, e.ERROR_GET_COMMENTS_FAIL)
		return
	}

	count, err := commentSer.Count()
	if err != nil {
		logging.Info(4)
		appG.ResponseErr(http.StatusInternalServerError, e.ERROR_COUNT_COMMENT_FAIL)
		return
	}

	appG.ResponseOk(defs.ResComments{
		Comments: comments,
		Count:    count,
	})
}

// AddComment 新增视频评论
// @Summary get add comment for video
// @Accept json
// @Produce json
// @Param vid-id path string true "video id"
// @Param cbody body defs.ReqNewComment true "user credential"
// @Success 200 {string} json "{"code":200,"msg":"ok","data":defs.ResAddComment}"
// @Failure 400 {string} json "{"code":400,"msg":"请求参数错误","data":nil}"
// @Failure 500 {string} json "{"code":400,"msg":"添加评论失败","data":nil}"
// @Router /api/v1/videos/{vid-id}/comments [post]
func AddComment(c *gin.Context) {
	appG := app.Gin{C: c}

	// var cxbody defs.ReqNewComment
	// if err := c.ShouldBindJSON(&cxbody); err != nil {
	// 	appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, err)
	// 	return
	// }

	resbody, _ := ioutil.ReadAll(c.Request.Body)
	cbody := defs.ReqNewComment{}
	if err := json.Unmarshal(resbody, &cbody); err != nil {
		logging.Info(err)
		appG.ResponseErr(http.StatusBadRequest, e.ERROR_REQUEST_BODY_UNMARSHAL_JSON_FAIL)
		return
	}

	uid := cbody.UserID
	uname := c.Request.Header.Get(setting.HTTPHeadParamSetting.HeadweFieldUname)
	vid := c.Param("vid-id")
	content := cbody.Content

	valid := validation.Validation{}
	valid.Min(uid, 1, "user_id").Message("用户ID必须大于0")
	valid.Required(vid, "video_id").Message("视频id不能为空")
	valid.Required(uname, "user_name").Message("用户名不能为空")
	valid.Required(content, "content").Message("评论内容不能为空")

	if valid.HasErrors() {
		// app.MarkErrors(valid.Errors)
		logging.Info(valid.Errors)
		appG.ResponseErr(http.StatusBadRequest, e.INVALID_PARAMS)
		return
	}

	commentSer := commentservice.Comment{
		UserID:   uid,
		UserName: uname,
		VideoID:  vid,
		Content:  content,
	}

	_, err := commentSer.AddComment()
	if err != nil {
		logging.Info(err)
		appG.ResponseErr(http.StatusInternalServerError, e.ERROR_ADD_COMMENT_FAIL)
		return
	}

	appG.ResponseOk(defs.ResAddComment{
		Status: "ok",
	})
}
