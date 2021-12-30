package data

import (
	"axiangcoding/antonstar/api-system/internal/app/conf"
	"axiangcoding/antonstar/api-system/internal/app/data/schema"
	"axiangcoding/antonstar/api-system/pkg/logging"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var db *gorm.DB

func Setup() {
	db = initDB()
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(conf.Config.App.Data.Database.Source),
		&gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
				logger.Config{
					SlowThreshold:             time.Second,   // 慢 SQL 阈值
					LogLevel:                  logger.Silent, // 日志级别
					IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
					Colorful:                  false,         // 禁用彩色打印
				},
			)})
	if err != nil {
		logging.Fatal(err)
	}
	// auto migrate should not be used in production mode
	if err := db.AutoMigrate(
		&schema.User{},
	); err != nil {
		logging.Fatal(err)
	}
	logging.Info("database mysql connected success")
	s, err := db.DB()
	if err != nil {
		logging.Fatal(err)
	}
	s.SetMaxOpenConns(conf.Config.App.Data.Database.MaxOpenConn)
	s.SetMaxIdleConns(conf.Config.App.Data.Database.MaxIdleConn)
	return db
}

func GetDB() *gorm.DB {
	return db
}
