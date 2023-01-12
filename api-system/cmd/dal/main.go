package main

import (
	"github.com/axiangcoding/ax-web/data"
	"github.com/axiangcoding/ax-web/logging"
	"github.com/axiangcoding/ax-web/settings"
)

func main() {
	settings.Setup()
	logging.Setup()
	data.Setup()
	db := data.GetDB()
	data.GenCode(db)
}
