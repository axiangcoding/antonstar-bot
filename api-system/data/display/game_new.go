package display

import (
	"time"
)

type GameNew struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Link      string    `json:"link,omitempty"`
	PosterUrl string    `json:"poster_url,omitempty"`
	Title     string    `json:"title,omitempty"`
	Comment   string    `json:"comment,omitempty"`
	DateStr   string    `json:"date_str,omitempty"`
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
