package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/haibeichina/gin_video_server/scheduler/docs"

	"github.com/haibeichina/gin_video_server/scheduler/pkg/setting"
	"github.com/haibeichina/gin_video_server/scheduler/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/video-delete-record/:vid-id", v1.AddDelRec)
	}

	return r
}
