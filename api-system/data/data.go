package data

import (
	"github.com/axiangcoding/antonstar-bot/data/dal"
	"github.com/axiangcoding/antonstar-bot/data/table"
	"github.com/axiangcoding/antonstar-bot/logging"
	"github.com/axiangcoding/antonstar-bot/settings"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var db *gorm.DB

func Setup() {
	db = initDB()
}

func selectDb() gorm.Dialector {
	driver := settings.Config.App.Data.Database.Driver
	source := settings.Config.App.Data.Database.Source
	var dial gorm.Dialector
	switch driver {
	case "postgres":
		dial = postgres.Open(source)
		break
	default:
		logging.Fatalf("no such driver: %s", driver)
	}
	return dial
}

func initDB() *gorm.DB {
	dial := selectDb()
	db, err := gorm.Open(dial,
		&gorm.Config{
			NamingStrategy: &schema.NamingStrategy{SingularTable: true},
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
		logging.Fatalf("Can't connect to db: %s", err)
	}
	logging.Info("Database mysql connected success")
	setProperties(db)
	autoMigrate(db)
	dal.SetDefault(db)
	return db
}

func GetDB() *gorm.DB {
	return db
}

// 自动更新表结构
func autoMigrate(db *gorm.DB) {
	if err := db.AutoMigrate(
		&table.Mission{},
		&table.GameUser{},
		&table.QQGroupConfig{},
		&table.QQUserConfig{},
		&table.GlobalConfig{},
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
	s.SetMaxOpenConns(settings.Config.App.Data.Database.MaxOpenConn)
	s.SetMaxIdleConns(settings.Config.App.Data.Database.MaxIdleConn)
}

func GenCode(db *gorm.DB) {
	g := gen.NewGenerator(gen.Config{
		OutPath:       "./data/dal",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
		FieldNullable: true,
	})

	// Use the above `*gorm.DB` instance to initialize the generator,
	// which is required to generate structs from db when using `GenerateModel/GenerateModelAs`
	g.UseDB(db)

	// Generate default DAO interface for those specified structs
	g.ApplyBasic(table.Mission{},
		table.GameUser{},
		table.QQGroupConfig{},
		table.QQUserConfig{},
		table.GlobalConfig{})

	// Execute the generator
	g.Execute()
}
