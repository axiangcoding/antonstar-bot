package display

import (
	"bytes"
	"github.com/axiangcoding/ax-web/logging"
	"text/template"
)

type GameUser struct {
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
	Nick         string  `json:"nick"`
	Clan         string  `json:"clan"`
	ClanUrl      string  `json:"clan_url"`
	RegisterDate string  `json:"register_date"`
	Level        int     `json:"level"`
	Title        string  `json:"title"`
	TsABRate     float64 `json:"ts_ab_rate"`
	TsRBRate     float64 `json:"ts_rb_rate"`
	TsSBRate     float64 `json:"ts_sb_rate"`
	AsABRate     float64 `json:"as_ab_rate"`
	AsRBRate     float64 `json:"as_rb_rate"`
	AsSBRate     float64 `json:"as_sb_rate"`
	Banned       string  `json:"banned"`
}

const templateStr = `
游戏昵称: {{.Nick}}
联队: {{.Clan}}
注册时间: {{.RegisterDate}}
等级: {{.Level}}
头衔: {{.Title}}
是否被封禁: {{.Banned}}
ThunderSkill街机效率值: {{.TsABRate}}
ThunderSkill历史效率值: {{.TsRBRate}}
ThunderSkill全真效率值: {{.TsSBRate}}
数据最后刷新时间: {{.UpdatedAt}}

Tips: 输入”刷新 {{.Nick}}“可以刷新游戏数据，但距离上次刷新不能小于24小时
`

func (u GameUser) ToFriendlyString() string {
	var buf bytes.Buffer

	t, err := template.New("display").Parse(templateStr)
	if err != nil {
		logging.Warn(err)
		return "Error"
	}
	if err := t.Execute(&buf, u); err != nil {
		logging.Error(err)
		return "Error"
	}
	return buf.String()
}
