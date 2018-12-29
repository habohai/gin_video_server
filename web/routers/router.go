package routers

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/haibeichina/gin_video_server/web/pkg/setting"
	"github.com/haibeichina/gin_video_server/web/pkg/tmp"
	"github.com/haibeichina/gin_video_server/web/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.LoadHTMLGlob(setting.AppSetting.RuntimeRootPath + setting.AppSetting.TemplatePath + "*.html")

	gin.SetMode(setting.ServerSetting.RunMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Static("/static", tmp.GetRootFullPath())

	apiv1 := r.Group("/api/v1")

	{
		apiv1.GET("/", v1.HomeHandler)
		apiv1.POST("/", v1.HomeHandler)
		apiv1.GET("/userhome", v1.UserHomeHandler)
		apiv1.POST("/userhome", v1.UserHomeHandler)
		apiv1.POST("/api", v1.APIHandler)
		// apiv1.GET("/videos/:vid-id", v1.ReverseProxy())
		// apiv1.POST("/upload/:vid-id", v1.ReverseProxy())
		apiv1.GET("/videos/:vid-id", v1.ProxyVideoHandler)
		apiv1.POST("/upload/:vid-id", v1.ProxyUploadHandler)
		fmt.Println(tmp.GetRootFullPath())
		//apiv1.Static("/static", tmp.GetRootFullPath())
	}

	return r
}
