package models

import (
	"errors"

	//"github.com/haibeichina/gin_video_server/0commonpkg/comdefs"
	"github.com/haibeichina/gin_video_server/apiserver/pkg/defs"
	"github.com/jinzhu/gorm"
)

// CheckUser 查看用户是否存在
func CheckUser(uname string) (bool, error) {
	var user defs.User
	err := db.Where("user_name = ?", uname).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

// AddUser 新增用户
func AddUser(data map[string]interface{}) error {
	uname := data["user_name"].(string)
	pwd := data["pwd"].(string)
	if !(len(uname) > 0 && len(pwd) > 0) {
		return errors.New("username and pwd is required")
	}

	user := defs.User{
		UserName: uname,
		Pwd:      pwd,
	}
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

// GetUserInfo 获取用户信息
func GetUserInfo(uname string) (*defs.User, error) {
	var user defs.User
	err := db.Where("user_name = ?", uname).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}

// DelUserInfo 删除用户
func DelUserInfo(data map[string]interface{}) error {
	uname := data["user_name"].(string)
	pwd := data["pwd"].(string)
	if !(len(uname) > 0 && len(pwd) > 0) {
		return errors.New("username and pwd is required")
	}
	user := defs.User{
		UserName: uname,
		Pwd:      pwd,
	}

	err := db.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
