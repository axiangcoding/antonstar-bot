package display

import (
	"time"
)

type GameNew struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Link      string `gorm:"uniqueIndex;size:255"`
	PosterUrl string
	Title     string
	Comment   string
	DateStr   string
}

const templateGameNewStr = `
大家好，我为各位带来了官网的最新新闻
== {{.Title}} ==
-------
{{.Comment}}
-------
发布于 {{.DateStr}}
如果感兴趣，请到官网查看详情哦~
`

func (n GameNew) ToFriendlyString() string {
	return parseTemplate(templateGameNewStr, n)
}
