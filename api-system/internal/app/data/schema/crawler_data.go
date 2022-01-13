package schema

import "gorm.io/gorm"

type CrawlerData struct {
	gorm.Model
	HttpStatus int
	Found      bool
	QueryID    string `gorm:"size:255"`
	Source     string `gorm:"size:255"`
	Nick       string `gorm:"size:255"`
	Content    string
}
