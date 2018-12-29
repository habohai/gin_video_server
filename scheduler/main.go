package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/haibeichina/gin_video_server/scheduler/models"
	"github.com/haibeichina/gin_video_server/scheduler/pkg/logging"
	"github.com/haibeichina/gin_video_server/scheduler/pkg/setting"
	"github.com/haibeichina/gin_video_server/scheduler/routers"
	"github.com/haibeichina/gin_video_server/scheduler/taskrunner"
)

func main() {
	setting.SetUp()
	logging.SetUp()
	models.SetUp()

	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go taskrunner.Start()

	go func() {
		if err := s.ListenAndServe(); err != nil {
			logging.Infof("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	logging.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		logging.Fatal("Server Shutdown:", err)
	}

	logging.Info("Server exiting")
}
