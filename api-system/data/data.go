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

var _db *gorm.DB

func InitData() {
	_db = initDB()
}

func initDB() *gorm.DB {
	source := settings.C().App.Data.Db.Source
	dial := postgres.Open(source)
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
		logging.L().Fatal("can't connect to db",
			logging.Error(err),
			logging.Any("source", source))
	}
	logging.L().Info("database mysql connected success")
	setProperties(db)
	autoMigrate(db)
	dal.SetDefault(db)
	return db
}

func GetDB() *gorm.DB {
	return _db
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
		logging.L().Fatal("auto migrate error", logging.Error(err))
	} else {
		logging.L().Info("auto migrate db tables success")
	}
}

func setProperties(db *gorm.DB) {
	s, err := db.DB()
	if err != nil {
		logging.L().Fatal("open db failed while set properties",
			logging.Error(err))
	}
	s.SetMaxOpenConns(settings.C().App.Data.Db.MaxOpenConn)
	s.SetMaxIdleConns(settings.C().App.Data.Db.MaxIdleConn)
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
