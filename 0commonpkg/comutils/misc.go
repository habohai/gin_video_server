package comutils

import (
	"strconv"
	"time"
)

// GetCurrentTimestampSec 获取当前秒数
func GetCurrentTimestampSec() int {
	ts, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))
	return ts
}
