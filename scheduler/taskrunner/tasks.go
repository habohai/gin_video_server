package taskrunner

import (
	"errors"
	"sync"

	"github.com/haibeichina/gin_video_server/0commonpkg/file"
	"github.com/haibeichina/gin_video_server/scheduler/models"
	"github.com/haibeichina/gin_video_server/scheduler/pkg/delfile"
	"github.com/haibeichina/gin_video_server/scheduler/pkg/logging"
	"github.com/haibeichina/gin_video_server/scheduler/pkg/setting"
)

// deleteVideo 删除实际的视频文件
func deleteVideo(vid string) error {
	path := delfile.GetVideoFilePath()
	logging.Info(path + vid)
	err := file.DeleteFile(path, vid)
	if err != nil {
		logging.Errorf("Deleting video error: %v", err)
		return err
	}
	return nil
}

// VideoClearDispatcher 删除视频的任务分发方法
func VideoClearDispatcher(dc dataChan) error {
	res, err := models.GetDelRecs(setting.AppSetting.ReadItemNum)
	logging.Info(res)
	if err != nil {
		logging.Infof("Video clear dispatcher error: %v", err)
		return err
	}
	if len(res) == 0 {
		return errors.New("All tasks finished")
	}

	for _, delres := range res {
		dc <- delres.VideoID
	}
	return nil
}

// VideoClearExecutor 执行删除视频的方法
func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error
	var wg sync.WaitGroup

forloop:
	for {
		select {
		case vid := <-dc:
			wg.Add(1)
			go func(id interface{}) {
				defer wg.Done()
				if err := deleteVideo(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
				if err := models.DeleteDelRec(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
			}(vid)
		default:
			break forloop
		}
	}

	wg.Wait()

	errMap.Range(func(k, v interface{}) bool {
		err = v.(error)
		if err != nil {
			return false
		}
		return true
	})

	return err
}
