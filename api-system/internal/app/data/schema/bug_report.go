package schema

import "gorm.io/gorm"

type BugReport struct {
	gorm.Model
	// bug类型
	Type string `gorm:"size:255"`
	// bug标题
	Title string `gorm:"size:255"`
	// bug内容
	Content string
	// 用户ID
	UserId int64
}
