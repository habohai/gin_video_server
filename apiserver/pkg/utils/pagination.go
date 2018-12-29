package utils

import (
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"

	"github.com/haibeichina/gin_video_server/apiserver/pkg/setting"
)

// GetPage 分页获取方法
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}

	return result
}
