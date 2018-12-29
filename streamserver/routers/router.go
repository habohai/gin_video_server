package routers

import (
	"github.com/gin-gonic/gin"

	_ "github.com/haibeichina/gin_video_server/streamserver/docs"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/haibeichina/gin_video_server/streamserver/middleware/limiter"
	"github.com/haibeichina/gin_video_server/streamserver/pkg/setting"
	"github.com/haibeichina/gin_video_server/streamserver/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.LoadHTMLFiles(setting.AppSetting.RuntimeRootPath + setting.AppSetting.TemplatePath + "upload.html")

	gin.SetMode(setting.ServerSetting.RunMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")

	{
		apiv1.GET("/videos/:vid-id", limiter.Limiter(), v1.Stream)

		apiv1.POST("/upload/:vid-id", v1.Upload)

		apiv1.GET("/testpage", v1.TestPage)
	}

	return r
}
