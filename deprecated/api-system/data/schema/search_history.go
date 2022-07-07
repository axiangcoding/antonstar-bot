package schema

import "gorm.io/gorm"

type SearchHistory struct {
	gorm.Model
	// 搜索的类型
	Type string `gorm:"size:255"`
	// 用户
	UserId int64
	// 搜索的内容
	Context string `gorm:"size:255"`
}

const (
	SearchHistoryTypeCrawlerQuery = "crawler_data"
)
