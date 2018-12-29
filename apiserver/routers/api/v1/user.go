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
	"github.com/haibeichina/gin_video_server/apiserver/routers/api"
	"github.com/haibeichina/gin_video_server/apiserver/service/userservice"
)

// CreateUser 创建用户
// @Summary add a new user
// @Accept json
// @Produce json
// @Param ubody body defs.ReqUserCredential true "user credential"
// @Success 200 {string} json "{"code":200,"msg":"ok","data":nil}"
// @Failure 400 {string} json "{"code":400,"msg":"bad request","data":nil}"
// @Failure 500 {string} json "{"code":500,"msg":"internal error","data":nil}"
// @Router /api/v1/user [post]
func CreateUser(c *gin.Context) {
	appG := app.Gin{C: c}

	resbody, _ := ioutil.ReadAll(c.Request.Body)
	ubody := &defs.ReqUserCredential{}
	if err := json.Unmarshal(resbody, ubody); err != nil {
		logging.Info(err)
		appG.ResponseErr(http.StatusBadRequest, e.INVALID_PARAMS)
		return
	}

	uname := ubody.Username
	pwd := ubody.Pwd

	valid := validation.Validation{}
	valid.Required(uname, "user_name").Message("用户名不能为空")
	valid.Required(pwd, "pwd").Message("密码不能为空")
	valid.MaxSize(pwd, 24, "pwd").Message("用户名最长24个字符")
	valid.MinSize(pwd, 6, "pwd").Message("用户名最短6个字符")

	if valid.HasErrors() {
		//app.MarkErrors(valid.Errors)
		logging.Info(valid.Errors)
		appG.ResponseErr(http.StatusBadRequest, e.INVALID_PARAMS)
		return
	}

	userSer := userservice.User{
		UserName: uname,
		Pwd:      pwd,
	}

	if ie, err := userSer.CreateUser(); err != nil {
		logging.Info(err)
		appG.ResponseErr(http.StatusInternalServerError, ie)
		return
	}

	id := api.GenerateNewSessionID(ubody.Username)
	resdata := defs.ResSignedIn{
		Success:   true,
		SessionID: id,
	}

	appG.ResponseOk(resdata)
}

// GetUserInfo 获取用户信息
// @Summary get user info
// @Accept json
// @Produce json
// @Param username path string true "username"
// @Success 200 {string} json "{"code":200,"msg":"ok","data":defs.ResUserInfo}"
// @Failure 400 {string} json "{"code":400,"msg":"请求参数错误","data":nil}"
// @Failure 500 {string} json "{"code":500,"msg":"数据库错误","data":nil}"
// @Router /api/v1/user/{username} [get]
func GetUserInfo(c *gin.Context) {
	appG := app.Gin{C: c}
	uname := c.Param("username")

	valid := validation.Validation{}
	valid.Required(uname, "username").Message("用户名不能为空")

	if valid.HasErrors() {
		// app.MarkErrors(valid.Errors)
		logging.Info(valid.Errors)
		appG.ResponseErr(http.StatusBadRequest, e.INVALID_PARAMS)
		return
	}

	userServer := userservice.User{
		UserName: uname,
	}

	user, err := userServer.GetUserInfo()
	if err != nil {
		logging.Info(err)
		appG.ResponseErr(http.StatusInternalServerError, e.ERROR_DB_ERROR_INFO)
		return
	}

	ui := defs.ResUserInfo{
		ID: user.ID,
	}
	appG.ResponseOk(ui)
}

// UserLogin 用户登陆请求
// @Summary user login
// @Accept json
// @Produce json
// @Param username path string true "username"
// @Param ubody body defs.ReqUserCredential true "user credential"
// @Success 200 {string} json "{"code":200,"msg":"ok","data":defs.ResSignedIn}"
// @Failure 400 {string} json "{"code":400,"msg":"请求参数错误","data":nil}"
// @Failure 500 {string} json "{"code":500,"msg":"","data":nil}"
// @Router /api/v1/user/{username} [post]
func UserLogin(c *gin.Context) {
	appG := app.Gin{C: c}

	resbody, _ := ioutil.ReadAll(c.Request.Body)
	ubody := &defs.ReqUserCredential{}
	if err := json.Unmarshal(resbody, ubody); err != nil {
		logging.Info(err)
		appG.ResponseErr(http.StatusBadRequest, e.ERROR_REQUEST_BODY_UNMARSHAL_JSON_FAIL)
		return
	}

	uname := c.Param("username")

	if uname != ubody.Username {
		logging.Infof("username Mismatch，in url is %s, in http body is %s", uname, ubody.Username)
		appG.ResponseErr(http.StatusBadRequest, e.ERROR_REQUEST_MISMATCH_USERNAME_URL_BODY)
		return
	}

	userser := userservice.User{
		UserName: ubody.Username,
		Pwd:      ubody.Pwd,
	}

	if ie, err := userser.UserLogin(); err != nil {
		logging.Info(err)
		appG.ResponseErr(http.StatusInternalServerError, ie)
		return
	}

	id := api.GenerateNewSessionID(ubody.Username)
	resdata := defs.ResSignedIn{
		Success:   true,
		SessionID: id,
	}

	appG.ResponseOk(resdata)
}
