package taskrunner

import (
	"time"

	"github.com/haibeichina/gin_video_server/scheduler/pkg/setting"
)

// Worker 定时任务类
type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

// NewWorker Worker的构造函数
func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker{
		ticker: time.NewTicker(interval),
		runner: r,
	}
}

func (w *Worker) startWorker() {
	for {
		select {
		case <-w.ticker.C:
			go w.runner.StartAll()
		}
	}
}

// Start 开始任务
func Start() {
	r := NewRunner(setting.AppSetting.ReadItemNum, true, VideoClearDispatcher, VideoClearExecutor)
	w := NewWorker(setting.AppSetting.IntervalTime, r)
	go w.startWorker()
}
