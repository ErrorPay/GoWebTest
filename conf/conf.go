package conf

import (
	"go-web-server/cache"
	"go-web-server/model"
	"go-web-server/util"
	"gorm.io/gorm/logger"
	"os"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	_ = godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 读取翻译文件
	if err := LoadLocales("conf/locales/zh-cn.yaml"); err != nil {
		util.Log().Panic("翻译文件加载失败", err)
	}

	var logLevel logger.LogLevel
	switch os.Getenv("DB_LOG_LEVEL") {
	case "silent":
		logLevel = logger.Silent
	case "warn":
		logLevel = logger.Warn
	case "error":
		logLevel = logger.Error
	case "info":
		logLevel = logger.Info
	}

	// 连接数据库
	model.Database(os.Getenv("MYSQL_DSN"), logLevel)
	cache.Redis()
}
