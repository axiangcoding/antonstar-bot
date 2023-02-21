package main

import (
	"github.com/axiangcoding/antonstar-bot/internal/data"
	"github.com/axiangcoding/antonstar-bot/pkg/logging"
	"github.com/axiangcoding/antonstar-bot/setting"
)

func main() {
	setting.InitConf()
	cfg := setting.C()
	logging.InitLogger(cfg.App.Log.Level, cfg.Server.RunMode)
	data.InitData(cfg.App.Data.Db.Source, cfg.App.Data.Db.MaxOpenConn, cfg.App.Data.Db.MaxIdleConn)
	db := data.Db()
	data.GenCode(db)
}
