package display

import (
	"bytes"
	"text/template"
)

type GameUser struct {
	CreatedAt    string  `json:"CreatedAt"`
	UpdatedAt    string  `json:"UpdatedAt"`
	Nick         string  `json:"Nick"`
	Clan         string  `json:"Clan"`
	ClanUrl      string  `json:"ClanUrl"`
	RegisterDate string  `json:"RegisterDate"`
	Level        int     `json:"Level"`
	Title        string  `json:"Title"`
	TsABRate     float64 `json:"TsABRate"`
	TsRBRate     float64 `json:"TsRBRate"`
	TsSBRate     float64 `json:"TsSBRate"`
	AsABRate     float64 `json:"AsABRate"`
	AsRBRate     float64 `json:"AsRBRate"`
	AsSBRate     float64 `json:"AsSBRate"`
	Banned       string  `json:"Banned"`
}

const templateStr = `
游戏昵称: {{.Nick}}
联队: {{.Clan}}
注册时间: {{.RegisterDate}}
等级: {{.Level}}
头衔: {{.Title}}
是否被封禁: {{.Banned}}
`

func (u GameUser) ToFriendlyString() string {
	var buf bytes.Buffer

	t, err := template.New("display").Parse(templateStr)
	if err != nil {
		return "Error"
	}
	if err := t.Execute(&buf, u); err != nil {
		return "Error"
	}
	return buf.String()
}
