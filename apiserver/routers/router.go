package routers

import (
	"github.com/gin-gonic/gin"

	_ "github.com/haibeichina/gin_video_server/apiserver/docs"

	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/haibeichina/gin_video_server/apiserver/middleware/validatesession"
	"github.com/haibeichina/gin_video_server/apiserver/pkg/setting"
	"github.com/haibeichina/gin_video_server/apiserver/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.ServerSetting.RunMode)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiv1 := r.Group("/api/v1")

	// 认证中间件
	// apiv1.Use(validatesession.ValidateSession())

	{
		// 新建用户
		apiv1.POST("/user", v1.CreateUser)

		// 用户登陆
		apiv1.POST("/user/:username", v1.UserLogin)

		// 获取用户信息
		apiv1.GET("/user/:username", validatesession.ValidateSession(), v1.GetUserInfo)

		// 上传新视频
		apiv1.POST("/user/:username/videos", validatesession.ValidateSession(), v1.AddVideo)
		// 获取视频列表
		apiv1.GET("/user/:username/videos", validatesession.ValidateSession(), v1.GetUserVideos)
		// 删除视频
		apiv1.DELETE("/user/:username/videos/:vid-id", validatesession.ValidateSession(), v1.DeleteVideo)

		//新建评论
		apiv1.POST("/videos/:vid-id/comments", validatesession.ValidateSession(), v1.AddComment)
		//获取评论列表
		apiv1.GET("/videos/:vid-id/comments", validatesession.ValidateSession(), v1.GetComments)
	}

	return r
}
