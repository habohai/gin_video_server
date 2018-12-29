package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/haibeichina/gin_video_server/0commonpkg/e"

	"github.com/haibeichina/gin_video_server/web/service/apihandlerservice"

	"github.com/haibeichina/gin_video_server/web/pkg/app"
	"github.com/haibeichina/gin_video_server/web/pkg/defs"
	"github.com/haibeichina/gin_video_server/web/pkg/logging"
	"github.com/haibeichina/gin_video_server/web/pkg/setting"
)

// HomeHandler 登陆界面
// @Summary sign in home
// @Accept json
// @Produce json
// @Param username cookie string true "username"
// @Param session cookie string true "session"
// @Success 200 {html} html "home.html"
// @Router / [post, get]
func HomeHandler(c *gin.Context) {
	cname, err1 := c.Cookie("username")
	sid, err2 := c.Cookie("session")

	if err1 != nil || err2 != nil {

		c.HTML(http.StatusOK, "home.html", gin.H{
			"Name": "youke",
		})
		return
	}

	// 不仅仅要判断是否存在还要判断是否是正确的
	if len(cname) != 0 && len(sid) != 0 {
		fmt.Println("redirect")
		c.Redirect(http.StatusFound, setting.AppSetting.BashPath+"/userhome")
		return
	}
}

// UserHomeHandler 用户页面
// @Summary user home
// @Accept json
// @Produce json
// @Param username cookie string true "username"
// @Param session cookie string true "session"
// @Success 200 {html} html "userhome.html"
// @Router /userhome [post, get]
func UserHomeHandler(c *gin.Context) {
	cname, err1 := c.Cookie("username")
	_, err2 := c.Cookie("session")
	if err1 != nil || err2 != nil {
		c.Redirect(http.StatusFound, setting.AppSetting.BashPath+"/")
		return
	}

	fname := c.PostForm("username")

	var name string

	if len(cname) != 0 {
		name = cname
	} else if len(fname) != 0 {
		name = fname
	}

	c.HTML(http.StatusOK, "userhome.html", gin.H{
		"Name": name,
	})
	return
}

// APIHandler api转发
// @Summary API Handler
// @Accept json
// @Produce json
// @Param resbody body defs.APIBody true "a api body"
// @Success 200 {html} html "userhome.html"
// @Failure 400 {string} json "{"code":400,"msg":"","data":""}"
// @Router /api [post]
func APIHandler(c *gin.Context) {
	appG := app.Gin{C: c}
	if c.Request.Method != http.MethodPost {
		appG.ResponseErr(http.StatusBadRequest, e.ERROR_REQUEST_NOT_RECOGNIZED)
		return
	}

	resbody, _ := ioutil.ReadAll(c.Request.Body)
	apibody := &defs.APIBody{}
	if err := json.Unmarshal(resbody, apibody); err != nil {
		logging.Info(err)
		appG.ResponseErr(http.StatusBadRequest, e.ERROR_REQUEST_BODY_PARSE_FAILED)
		return
	}

	apihandlers := apihandlerservice.APIHandler{

		URL:     apibody.URL,
		Method:  apibody.Method,
		ReqBody: apibody.ReqBody,

		HTTPRequest: c.Request,
		HTTPClient:  &http.Client{},
	}
	apiResp, err := apihandlers.APIHandlerRequest()
	if err != nil {
		logging.Info(err)
		appG.ResponseErr(http.StatusInternalServerError, e.ERROR)
		return
	}

	apiResBody, err := ioutil.ReadAll(apiResp.Body)
	if err != nil {
		logging.Info(err)
		appG.ResponseErr(http.StatusInternalServerError, e.ERROR)
		return
	}

	c.String(apiResp.StatusCode, string(apiResBody))
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

func ReverseProxy() gin.HandlerFunc {
	return func(c *gin.Context) {
		target, err := url.Parse(setting.AppSetting.ProxyURL)
		if err != nil {
			logging.Info(err)
			c.String(http.StatusInternalServerError, "internal server error")
			return
		}
		targetQuery := target.RawQuery
		director := func(req *http.Request) {
			req.URL.Scheme = target.Scheme
			req.URL.Host = target.Host
			req.Host = target.Host // REWRITE req.host
			req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
			if targetQuery == "" || req.URL.RawQuery == "" {
				req.URL.RawQuery = targetQuery + req.URL.RawQuery
			} else {
				req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
			}
			if _, ok := req.Header["User-Agent"]; !ok {
				// explicitly disable User-Agent so it's not set to default value
				req.Header.Set("User-Agent", "")
			}
		}
		proxy := &httputil.ReverseProxy{Director: director}
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

// ProxyVideoHandler proxy转发
// @Summary Proxy Handler
// @Router /upload/{vid-id} [get]
func ProxyVideoHandler(c *gin.Context) {
	u, err := url.Parse(setting.AppSetting.ProxyURL)
	if err != nil {
		logging.Info(err)
		c.String(http.StatusInternalServerError, "internal server error")
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(u)
	c.Request.Host = u.Host
	proxy.ServeHTTP(c.Writer, c.Request)
}

// ProxyUploadHandler proxy转发
// @Summary Proxy Handler
// @Router /upload/{vid-id} [post]
func ProxyUploadHandler(c *gin.Context) {
	u, err := url.Parse(setting.AppSetting.ProxyURL)
	if err != nil {
		logging.Info(err)
		c.String(http.StatusInternalServerError, "internal server error")
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(u)
	c.Request.Host = u.Host
	proxy.ServeHTTP(c.Writer, c.Request)
}
