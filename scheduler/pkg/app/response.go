package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haibeichina/gin_video_server/0commonpkg/e"
)

// Gin gin.Context 封装
type Gin struct {
	C *gin.Context
}

// // Response Gin http的返回方法
// func (g *Gin) Response(httpCode, errCode int, data interface{}) {
// 	g.C.JSON(httpCode, gin.H{
// 		"code": httpCode,
// 		"msg":  e.GetMsg(errCode),
// 		"data": data,
// 	})

// 	return
// }

// ResponseErr Gin http的返回错误的方法
func (g *Gin) ResponseErr(httpCode, errCode int) {
	g.C.String(httpCode, "error code:%d, error msg:%s", errCode, e.GetMsg(errCode))

	return
}

// ResponseOk gin http的返回成功的方法
func (g *Gin) ResponseOk(obj interface{}) {
	g.C.JSON(http.StatusOK, obj)
	return
}
