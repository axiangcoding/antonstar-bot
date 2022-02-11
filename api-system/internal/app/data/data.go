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
			DisableForeignKeyConstraintWhenMigrating: true,
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
		logging.Fatalf("Can't connect to MySql: %s", err)
	}
	logging.Info("Database mysql connected success")
	setProperties(db)
	autoMigrate(db)
	return db
}

func GetDB() *gorm.DB {
	return db
}

// 自动更新表结构
func autoMigrate(db *gorm.DB) {
	if err := db.Set("gorm:table_options",
		"ENGINE=InnoDB CHARSET=utf8mb4 COLLATE utf8mb4_bin").AutoMigrate(
		&schema.User{},
		&schema.UserCert{},
		&schema.CrawlerData{},
		&schema.Visit{},
		&schema.BugReport{},
		&schema.SystemConfig{},
	); err != nil {
		logging.Fatal(err)
	} else {
		logging.Info("Auto migrate database table success")
	}
}

func setProperties(db *gorm.DB) {
	s, err := db.DB()
	if err != nil {
		logging.Fatal(err)
	}
	s.SetMaxOpenConns(conf.Config.App.Data.Database.MaxOpenConn)
	s.SetMaxIdleConns(conf.Config.App.Data.Database.MaxIdleConn)
}
