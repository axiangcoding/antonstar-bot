package main

import (
	"github.com/axiangcoding/antonstar-bot/data"
	"github.com/axiangcoding/antonstar-bot/logging"
	"github.com/axiangcoding/antonstar-bot/settings"
)

func main() {
	settings.InitConf()
	logging.InitLogger()
	data.InitData()
	db := data.Db()
	data.GenCode(db)
}
