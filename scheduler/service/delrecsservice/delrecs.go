package delrecsservice

import (
	"github.com/haibeichina/gin_video_server/scheduler/models"

	//"github.com/haibeichina/gin_video_server/0commonpkg/comdefs"
	"github.com/haibeichina/gin_video_server/scheduler/pkg/defs"
)

// DelRecs 视频应删除表结构
type DelRecs struct {
	VideoID string `json:"video_id"`

	Count int
}

// AddDelRec 将软删除的视频信息添加到删除表中，等待定时硬删除
func (d *DelRecs) AddDelRec() error {
	if err := models.AddDelRec(d.VideoID); err != nil {
		return err
	}

	return nil
}

// GetDelRecs 获取一定数量的视频删除记录用于硬删除
func (d *DelRecs) GetDelRecs() ([]defs.Delrec, error) {
	delrecs, err := models.GetDelRecs(d.Count)
	if err != nil {
		return nil, err
	}

	return delrecs, nil
}
