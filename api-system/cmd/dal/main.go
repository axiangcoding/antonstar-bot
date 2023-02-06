package main

import (
	"github.com/axiangcoding/antonstar-bot/data"
	"github.com/axiangcoding/antonstar-bot/logging"
	"github.com/axiangcoding/antonstar-bot/settings"
)

func main() {
	settings.Setup()
	logging.Setup()
	data.Setup()
	db := data.GetDB()
	data.GenCode(db)
}
