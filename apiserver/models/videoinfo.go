package models

import (
	//"github.com/haibeichina/gin_video_server/0commonpkg/comdefs"
	"github.com/haibeichina/gin_video_server/apiserver/pkg/defs"
	"github.com/jinzhu/gorm"
)

// AddVideoInfo 新增视频信息
func AddVideoInfo(data map[string]interface{}) (*defs.Videoinfo, error) {
	videoInfo := defs.Videoinfo{
		UserID:    data["user_id"].(int64),
		UserName:  data["user_name"].(string),
		VideoName: data["video_name"].(string),
	}

	if err := db.Create(&videoInfo).Error; err != nil {
		return nil, err
	}

	return &videoInfo, nil
}

// GetVideoInfo 获取视频信息
func GetVideoInfo(vid string) (*defs.Videoinfo, error) {
	var videoInfo defs.Videoinfo
	err := db.Where("id = ?", vid).First(&videoInfo).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &videoInfo, nil
}

// GetVideoInfos 获取某段时间内用户的视频
func GetVideoInfos(uname string, from, to int) ([]defs.Videoinfo, error) {
	//db.Joins("INNER JOIN video_users ON video_viedoinfos.user_id = video_users.id ").Where("video_users.login_name = ? AND video_videoinfos.created_on > ? AND video_videoinfos.create_on <= ?", uname, from, to).Order("video_videoinfos.create_on DESC").Find(&videoInfos)
	var videoInfos []defs.Videoinfo
	err := db.Where("login_name = ? AND created_on > ? AND created_on <= ?", uname, from, to).Order("created_on DESC").Find(&videoInfos).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return videoInfos, nil
}

// GetVideoInfosByPage 按照分页获取视频信息
func GetVideoInfosByPage(pageNum int, pageSize int, maps interface{}) ([]defs.Videoinfo, error) {
	var videoInfos []defs.Videoinfo
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Order("created_on DESC").Find(&videoInfos).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return videoInfos, nil
}

// GetVideoInfoTotal 获取此用户的所有视频数
func GetVideoInfoTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&defs.Videoinfo{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// ExistVideoInfoByVID 查看视频信息是否存在
func ExistVideoInfoByVID(maps interface{}) (bool, error) {
	// rows, err := db.Table("video_videoinfos").Where(maps).Rows()
	// if err != nil {
	// 	return false, err
	// }
	// defer rows.Close()

	// if !rows.Next() {
	// 	return false, nil
	// }
	var videoInfo defs.Videoinfo
	err := db.Where(maps).First(&videoInfo).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if videoInfo.ID != "" {
		return true, nil
	}
	return false, nil
}

// DeleteVideoInfo 删除视频
func DeleteVideoInfo(vid string) error {
	if err := db.Where("id = ?", vid).Delete(&defs.Videoinfo{}).Error; err != nil {
		return err
	}

	return nil
}
