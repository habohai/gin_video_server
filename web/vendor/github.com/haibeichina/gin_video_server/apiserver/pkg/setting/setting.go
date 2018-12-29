package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

// App 设置对象
type App struct {
	PageSize        int
	RuntimeRootPath string
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

// HTTPHeadParam http前端填充的头部参数
type HTTPHeadParam struct {
	HeadweFieldSession string
	HeadweFieldUname   string
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

// Scheduler 定时任务服务配置
type Scheduler struct {
	Host  string
	Param string
}

// 统一实例化设置对象
var (
	AppSetting           = &App{}
	ServerSetting        = &Server{}
	HTTPHeadParamSetting = &HTTPHeadParam{}
	DatabaseSetting      = &Database{}
	SchedulerSetting     = &Scheduler{}
)

var cfg *ini.File

// SetUp 设置初始化函数
func SetUp() {
	var err error
	cfg, err = ini.Load("conf/apiserver.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/apiserver.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("httpheadparam", HTTPHeadParamSetting)
	mapTo("database", DatabaseSetting)
	mapTo("scheduler", SchedulerSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
