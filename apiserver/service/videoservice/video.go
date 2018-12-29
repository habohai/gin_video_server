package videoservice

import (
	//"github.com/haibeichina/gin_video_server/0commonpkg/comdefs"
	"github.com/haibeichina/gin_video_server/0commonpkg/e"
	"github.com/haibeichina/gin_video_server/apiserver/models"
	"github.com/haibeichina/gin_video_server/apiserver/pkg/defs"
)

// VideoInfo 视频信息表结构
type VideoInfo struct {
	ID        string `json:"id"`
	UserID    int64  `json:"user_id"`
	UserName  string `json:"user_name"`
	VideoName string `json:"video_name"`
	CreateOn  int64  `json:"created_on"`

	PageNum  int
	PageSize int
}

// GetUserVideos 获取给定用户的视频信息
func (v *VideoInfo) GetUserVideos() ([]defs.Videoinfo, error) {
	videoInfos, err := models.GetVideoInfosByPage(v.PageNum, v.PageSize, v.getMapForGet())
	if err != nil {
		return nil, err
	}

	return videoInfos, nil
}

// Count 获取视频的总评论数
func (v *VideoInfo) Count() (int, error) {
	count, err := models.GetVideoInfoTotal(v.getMapForGet())
	if err != nil {
		return 0, err
	}

	return count, nil
}

// AddVideo 添加视频
func (v *VideoInfo) AddVideo() (*defs.Videoinfo, error) {
	video, err := models.AddVideoInfo(v.getMapForAdd())
	if err != nil {
		return nil, err
	}
	return video, nil
}

// DeleteVideo 删除视频
func (v *VideoInfo) DeleteVideo() (int, error) {
	maps := make(map[string]interface{})

	if v.UserID > 0 {
		maps["id"] = v.ID
	}

	ok, err := models.ExistVideoInfoByVID(maps)
	if err != nil {
		return e.ERROR_DB_ERROR_INFO, err
	}

	if !ok {
		return e.SUCCESS, nil
	}

	if err := models.DeleteVideoInfo(v.ID); err != nil {
		return e.ERROR_DB_ERROR_INFO, err
	}

	return e.SUCCESS, nil
}

func (v *VideoInfo) getMapForGet() map[string]interface{} {
	maps := make(map[string]interface{})
	if v.UserName != "" {
		maps["user_name"] = v.UserName
	}

	return maps
}

func (v *VideoInfo) getMapForAdd() map[string]interface{} {
	maps := make(map[string]interface{})

	if v.UserID > 0 {
		maps["user_id"] = v.UserID
	}

	if v.UserName != "" {
		maps["user_name"] = v.UserName
	}

	if v.VideoName != "" {
		maps["video_name"] = v.VideoName
	}

	return maps
}
