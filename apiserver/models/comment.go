package models

import (
	"github.com/jinzhu/gorm"

	//"github.com/haibeichina/gin_video_server/0commonpkg/comdefs"
	"github.com/haibeichina/gin_video_server/apiserver/pkg/defs"
)

// AddComment 添加评论
func AddComment(data map[string]interface{}) (*defs.Comment, error) {
	comment := defs.Comment{
		UserID:   data["user_id"].(int64),
		UserName: data["user_name"].(string),
		VideoID:  data["video_id"].(string),
		Content:  data["content"].(string),
	}

	if err := db.Create(&comment).Error; err != nil {
		return nil, err
	}

	return &comment, nil
}

// GetComment 获取某个评论
func GetComment(vid string) (*defs.Comment, error) {
	var comment defs.Comment
	err := db.Where("video_id = ?", vid).Order("created_on DESC").Find(&comment).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &comment, nil
}

// GetComments 获取指定时间段的评论
func GetComments(vid string, from, to int64) ([]defs.Comment, error) {
	var comments []defs.Comment
	err := db.Where("video_id = ? AND (created_on > ? AND created_on <= ?)", vid, from, to).Order("created_on DESC").Find(&comments).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return comments, nil
}

// GetCommentsByPage 获取指定页的评论
func GetCommentsByPage(pageNum int, pageSize int, maps interface{}) ([]defs.Comment, error) {
	var comments []defs.Comment
	err := db.Where(maps).Offset(pageNum).Limit(pageSize).Order("created_on DESC").Find(&comments).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return comments, nil
}

// GetCommentTotal 获取此用户的所有视频数
func GetCommentTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&defs.Comment{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}
