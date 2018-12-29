package models

import (
	"github.com/jinzhu/gorm"

	//"github.com/haibeichina/gin_video_server/0commonpkg/comdefs"
	"github.com/haibeichina/gin_video_server/apiserver/pkg/defs"
)

// AddSession 添加session
func AddSession(data map[string]interface{}) error {
	session := defs.Session{
		ID:       data["id"].(string),
		UserName: data["user_name"].(string),
		TTL:      data["ttl"].(int64),
	}

	if err := db.Create(&session).Error; err != nil {
		return err
	}

	return nil
}

// GetSession 根据sid获取session的信息
func GetSession(sid string) (*defs.Session, error) {
	var session defs.Session
	err := db.Where("id = ?", sid).First(&session).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &session, nil
}

// GetSessions 获取所有session信息，用于程序初始化
func GetSessions() ([]defs.Session, error) {
	var sessions []defs.Session
	err := db.Find(&sessions).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return sessions, nil
}

// DeleteSession 删除session信息
func DeleteSession(sid string) error {
	if err := db.Where("id = ?", sid).Delete(&defs.Session{}).Error; err != nil {
		return err
	}

	return nil
}
