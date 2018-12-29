package utils

import (
	"errors"
	"net/http"

	"github.com/haibeichina/gin_video_server/apiserver/pkg/setting"

	"github.com/haibeichina/gin_video_server/apiserver/pkg/logging"
)

// SendDeleteVideoRequest 发送删除视频的请求
func SendDeleteVideoRequest(id string) error {
	host := setting.SchedulerSetting.Host
	param := setting.SchedulerSetting.Param
	url := "http://" + host + "/" + param + "/" + id
	res, err := http.Get(url)
	if err != nil {
		logging.Error("Sending deleting video request error:", err)
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("error")
	}

	return nil
}
