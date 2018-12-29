package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

// App 设置对象
type App struct {
	RuntimeRootPath string
	VideoPath       string
	TemplatePath    string
	LogSavePath     string
	LogSaveName     string
	LogFileExt      string
	TimeFormat      string
	ImageAllowExts  []string
}

// Server 服务端设置对象
type Server struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Limiter 视频播放限制对象
type Limiter struct {
	MaxUploadSize int64
	MaxLinkNum    int
}

// 统一实例化设置对象
var (
	AppSetting     = &App{}
	ServerSetting  = &Server{}
	LimiterSetting = &Limiter{}
)

var cfg *ini.File

// SetUp 设置初始化函数
func SetUp() {
	var err error
	cfg, err = ini.Load("conf/streamserver.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/streamserver.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("limiter", LimiterSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
