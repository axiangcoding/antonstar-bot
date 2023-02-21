package main

import (
	"github.com/axiangcoding/antonstar-bot/internal/data"
	"github.com/axiangcoding/antonstar-bot/pkg/logging"
	"github.com/axiangcoding/antonstar-bot/setting"
)

func main() {
	setting.InitConf()
	logging.InitLogger()
	data.InitData()
	db := data.Db()
	data.GenCode(db)
}
