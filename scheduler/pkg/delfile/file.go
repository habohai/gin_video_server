package delfile

import (
	"fmt"

	"github.com/haibeichina/gin_video_server/scheduler/pkg/setting"
)

func GetVideoFilePath() string {
	return fmt.Sprintf("%s%s", setting.AppSetting.RuntimeRootPath, setting.AppSetting.VideoPath)
}
