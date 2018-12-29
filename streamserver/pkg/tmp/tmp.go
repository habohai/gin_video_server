package tmp

import (
	"os"

	"github.com/haibeichina/gin_video_server/streamserver/pkg/setting"
)

// GetTmpFullPath 获取视频文件的路径
func GetTmpFullPath() string {
	return setting.AppSetting.RuntimeRootPath + setting.AppSetting.TemplatePath
}

// GetRootFullPath 获取视频文件的全路径
func GetRootFullPath() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	sf := GetTmpFullPath()
	sf = sf[:len(sf)-1]

	src := dir + "/" + sf

	return src
}
