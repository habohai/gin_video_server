package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/haibeichina/gin_video_server/0commonpkg/comutils"
	"github.com/haibeichina/gin_video_server/apiserver/pkg/logging"
	"github.com/haibeichina/gin_video_server/apiserver/pkg/setting"
)

var db *gorm.DB

func SetUp() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.PassWord,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.DBName))

	if err != nil {
		logging.Info(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	comDBSetup()
}

func ForTestSetUp() {
	var err error
	dbType := "mysql"
	dbName := "ginvideoserver"
	user := "haibei"
	password := "h9420x"
	host := "120.79.57.219:3306"
	tablePrefix := "video_"

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		logging.Info(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	comDBSetup()

}

func comDBSetup() {
	db.SingularTable(true)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.DB().SetMaxIdleConns(10)  // 设置空闲链接数
	db.DB().SetMaxOpenConns(100) // 设置数据库最大打开链接数
}

func CloseDB() {
	defer db.Close()
}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		if createTimeField, ok := scope.FieldByName("ID"); ok {
			if createTimeField.IsBlank {
				id, err := comutils.NewUUID()
				if err != nil {
					logging.Info(err)
				}
				createTimeField.Set(id)
			}
		}

		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

	}
}
