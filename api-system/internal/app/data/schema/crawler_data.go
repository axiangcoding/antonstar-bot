package schema

import "gorm.io/gorm"

type CrawlerData struct {
	gorm.Model
	HttpStatus int
	//下面三种信息在库中都是不可重复的
	Source  string `gorm:"size:255"`
	Nick    string `gorm:"size:255"`
	Content string
}
