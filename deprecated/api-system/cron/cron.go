package cron

import (
	"axiangcoding/antonstar/api-system/logging"
	"github.com/robfig/cron/v3"
)

func Setup() {
	c := cron.New()
	// addJob(c)
	c.Start()
}

func addJob(c *cron.Cron) {
	_, err := c.AddFunc("@every 5s", doSomething)
	if err != nil {
		logging.Fatalf("Add cron job failed. %s", err)
	}
}

func doSomething() {
	logging.Info("do something")
}
