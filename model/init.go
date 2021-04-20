package model

import (
	"go-web-server/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// DB 数据库链接单例
var DB *gorm.DB

// Database 在中间件中初始化mysql链接
func Database(connString string, logLevel logger.LogLevel) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             0,
			Colorful:                  true,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  logLevel,
		},
	)

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: newLogger,
	})

	// Error
	if err != nil {
		util.Log().Panic("连接数据库不成功", err)
	}
	//设置连接池
	sqlDB, err1 := db.DB()
	if err1 != nil {
		util.Log().Panic("连接池设置失败", err)
	} else {
		//空闲
		sqlDB.SetMaxIdleConns(50)
		//打开
		sqlDB.SetMaxOpenConns(100)
		//超时
		sqlDB.SetConnMaxLifetime(time.Second * 30)
	}

	DB = db

	migration()
}
