package models

import (
	//"github.com/haibeichina/gin_video_server/0commonpkg/comdefs"
	"github.com/haibeichina/gin_video_server/scheduler/pkg/defs"
	"github.com/jinzhu/gorm"
)

// AddDelRec 将软删除的视频信息添加到删除表中，等待定时硬删除
func AddDelRec(vid string) error {
	session := defs.Delrec{
		VideoID: vid,
	}

	if err := db.Create(&session).Error; err != nil {
		return err
	}

	return nil
}

// GetDelRecs 获取一定数量的视频删除记录用于硬删除
func GetDelRecs(count int) ([]defs.Delrec, error) {
	var delrecs []defs.Delrec
	err := db.Limit(count).Find(&delrecs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return delrecs, nil
}

// DeleteDelRecs 硬删除成功后删除视频删除表中的多个记录
func DeleteDelRecs(vids []string) error {
	if err := db.Where("video_id in (?)", vids).Delete(&defs.Delrec{}).Error; err != nil {
		return err
	}

	return nil
}

// DeleteDelRec 硬删除成功后删除视频删除表中的记录
func DeleteDelRec(vid string) error {
	if err := db.Where("video_id = ?", vid).Delete(&defs.Delrec{}).Error; err != nil {
		return err
	}

	return nil
}
