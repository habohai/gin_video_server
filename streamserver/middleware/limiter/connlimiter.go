package limiter

import (
	"fmt"
)

// ConnLimiter 设置连接限制类
type ConnLimiter struct {
	concurrentConn int
	bucket         chan int
}

// NewConnLimiter 生产链接限制类对象
func NewConnLimiter(cc int) *ConnLimiter {
	return &ConnLimiter{
		concurrentConn: cc,
		bucket:         make(chan int, cc),
	}
}

// CheckConn 查看链接数是否超过设置
func (cl *ConnLimiter) checkConn() bool {
	if len(cl.bucket) >= cl.concurrentConn {
		fmt.Printf("Reached the rate limitation")
		return false
	}
	return true
}

// GetConn 新增链接
func (cl *ConnLimiter) GetConn() bool {
	if ok := cl.checkConn(); !ok {
		return false
	}

	cl.bucket <- 1
	return true
}

// ReleaseConn 释放链接
func (cl *ConnLimiter) ReleaseConn() {
	c := <-cl.bucket
	fmt.Printf("New connction coming: %d", c)
}
