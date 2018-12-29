package userservice

import (
	"errors"

	//"github.com/haibeichina/gin_video_server/0commonpkg/comdefs"
	"github.com/haibeichina/gin_video_server/0commonpkg/e"

	"github.com/haibeichina/gin_video_server/apiserver/models"
	"github.com/haibeichina/gin_video_server/apiserver/pkg/defs"
)

// User 认证表结构
type User struct {
	UserID   int64  `json:"id"`
	UserName string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

// CreateUser 创建用户
func (u *User) CreateUser() (int, error) {
	user := map[string]interface{}{
		"user_name": u.UserName,
		"pwd":       u.Pwd,
	}

	ok, err := models.CheckUser(u.UserName)
	if err != nil {
		return e.ERROR_DB_ERROR_INFO, err
	}

	if ok {
		return e.ERROR_EXIST_USER_NAME, errors.New("username existed")
	}

	if err := models.AddUser(user); err != nil {
		return e.ERROR_DB_ERROR_INFO, err
	}

	return e.SUCCESS, nil
}

// GetUserInfo 获取用户信息
func (u *User) GetUserInfo() (*defs.User, error) {
	user, err := models.GetUserInfo(u.UserName)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UserLogin 验证用户登陆信息
func (u *User) UserLogin() (int, error) {
	user, err := models.GetUserInfo(u.UserName)
	if err != nil {
		return e.ERROR_DB_ERROR_INFO, err
	}

	if u.Pwd != user.Pwd {
		return e.ERROR_AUTH_PWD_UNMATCH, errors.New("Mismatch pwd")
	}

	return e.SUCCESS, nil
}
