package commentservice

import (
	//"github.com/haibeichina/gin_video_server/0commonpkg/comdefs"
	"github.com/haibeichina/gin_video_server/apiserver/models"
	"github.com/haibeichina/gin_video_server/apiserver/pkg/defs"
)

// Comment 评论类
type Comment struct {
	ID       string `json:"id"`
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	VideoID  string `json:"video_id"`
	Content  string `json:"content"`
	CreateOn int64  `json:"created_on"`

	PageNum  int
	PageSize int
}

// AddComment 对视频新增评论
func (c *Comment) AddComment() (*defs.Comment, error) {
	comment := map[string]interface{}{
		"user_id":   c.UserID,
		"user_name": c.UserName,
		"video_id":  c.VideoID,
		"content":   c.Content,
	}

	co, err := models.AddComment(comment)
	if err != nil {
		return nil, err
	}

	return co, nil
}

// GetComments 获取视频对应的所有评论
func (c *Comment) GetComments() ([]defs.Comment, error) {
	comments, err := models.GetCommentsByPage(c.PageNum, c.PageSize, c.getMaps())
	if err != nil {
		return nil, err
	}

	return comments, nil
}

// Count 获取视频的总评论数
func (c *Comment) Count() (int, error) {
	count, err := models.GetCommentTotal(c.getMaps())
	if err != nil {
		return 0, err
	}

	return count, nil
}

// getMaps 获取某视频所属所有评论的条件
func (c *Comment) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	if c.VideoID != "" {
		maps["video_id"] = c.VideoID
	}

	return maps
}
