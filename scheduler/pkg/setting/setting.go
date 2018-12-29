package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

// App 设置对象
type App struct {
	ReadItemNum     int
	IntervalTime    time.Duration
	RuntimeRootPath string
	VideoPath       string
	LogSavePath     string
	LogSaveName     string
	LogFileExt      string
	TimeFormat      string
}

// Server 服务端设置对象
type Server struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Database 数据库设置对象
type Database struct {
	Type        string
	User        string
	PassWord    string
	Host        string
	DBName      string
	TablePrefix string
}

// 统一实例化设置对象
var (
	AppSetting      = &App{}
	ServerSetting   = &Server{}
	DatabaseSetting = &Database{}
)

var cfg *ini.File

// SetUp 设置初始化函数
func SetUp() {
	var err error
	cfg, err = ini.Load("conf/scheduler.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/scheduler.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)

	AppSetting.IntervalTime = AppSetting.IntervalTime * time.Second
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
