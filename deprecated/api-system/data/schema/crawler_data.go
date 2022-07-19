package schema

import (
	"gorm.io/gorm"
)

type CrawlerData struct {
	gorm.Model
	HttpStatus    int
	Found         bool
	QueryID       string `gorm:"size:255"`
	Source        string `gorm:"size:255"`
	Nick          string `gorm:"size:255"`
	Status        string `gorm:"size:255"`
	TriggerUserId int64
	Content       string
}

type ShortCrawlerData struct {
	gorm.Model
	Found   bool
	QueryID string `gorm:"size:255"`
	Source  string `gorm:"size:255"`
	Status  string `gorm:"size:255"`
}

const (
	CrawlerStatusRunning = "running"
	CrawlerStatusDone    = "done"
)

const (
	SourceGaijin       = "gaijin"
	SourceThunderskill = "thunder_skill"
)
