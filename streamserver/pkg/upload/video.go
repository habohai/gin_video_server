package upload

import (
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"strings"

	"github.com/haibeichina/gin_video_server/0commonpkg/file"
	"github.com/haibeichina/gin_video_server/streamserver/pkg/logging"
	"github.com/haibeichina/gin_video_server/streamserver/pkg/setting"
)

// GetVideoFullPath 获取视频文件的路径
func GetVideoFullPath() string {
	return setting.AppSetting.RuntimeRootPath + setting.AppSetting.VideoPath
}

// CheckVideoExist 检查视频文件是否存在
func CheckVideoExist(fileName string) (bool, error) {
	dir, err := os.Getwd()
	if err != nil {
		return false, err
	}

	filesrc := dir + "/" + fileName
	return file.CheckExist(filesrc)
}

// CheckVideoExt 检查上传文件格式
func CheckVideoExt(fileName string) bool {
	ext := file.GetExt(fileName)
	for _, allowExt := range setting.AppSetting.ImageAllowExts {
		if strings.ToUpper(allowExt) == strings.ToUpper(ext) {
			return true
		}
	}

	return false
}

// CheckVideoSize 检查上传文件的大小
func CheckVideoSize(f multipart.File) bool {
	size, err := file.GetSize(f)
	if err != nil {
		log.Println(err)
		logging.Warn(err)
		return false
	}

	return int64(size) <= setting.LimiterSetting.MaxUploadSize
}

// CheckVideo 检查视频
func CheckVideo(src string) error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("os.Getwd err: %v", err)
	}

	err = file.IsNotExistMkDir(dir + "/" + src)
	if err != nil {
		return fmt.Errorf("file.IsNotExistMkDir err: %v", err)
	}

	perm := file.CheckPermission(src)
	if perm == true {
		return fmt.Errorf("file.CheckPermission Permission denied src: %s", src)
	}

	return nil
}
