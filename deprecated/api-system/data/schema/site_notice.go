package schema

import "gorm.io/gorm"

type SiteNotice struct {
	gorm.Model
	// 公告标题
	Title string `gorm:"size:255"`
	// 公告内容
	Content string
	// 编辑者的用户ID
	EditorUserId int64
}
