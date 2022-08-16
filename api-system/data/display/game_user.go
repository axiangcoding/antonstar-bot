package display

import (
	"bytes"
	"github.com/axiangcoding/ax-web/logging"
	"text/template"
)

type GameUser struct {
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	Nick         string `json:"nick"`
	Clan         string `json:"clan"`
	ClanUrl      string `json:"clan_url"`
	RegisterDate string `json:"register_date"`
	Level        int    `json:"level"`
	Title        string `json:"title"`
	StatAb       UserStat
	StatRb       UserStat
	StatSb       UserStat
	TsABRate     float64 `json:"ts_ab_rate"`
	TsRBRate     float64 `json:"ts_rb_rate"`
	TsSBRate     float64 `json:"ts_sb_rate"`
	AsABRate     float64 `json:"as_ab_rate"`
	AsRBRate     float64 `json:"as_rb_rate"`
	AsSBRate     float64 `json:"as_sb_rate"`
	Banned       string  `json:"banned"`
}

type UserStat struct {
	TotalMission         int    `json:"total_mission,omitempty"`
	WinRate              string `json:"win_rate,omitempty"`
	GroundDestroyCount   int    `json:"ground_destroy_count,omitempty"`
	FleetDestroyCount    int    `json:"fleet_destroy_count,omitempty"`
	GameTime             string `json:"game_time,omitempty"`
	AviationDestroyCount int    `json:"aviation_destroy_count,omitempty"`
	WinCount             int    `json:"win_count,omitempty"`
	SliverEagleEarned    int64  `json:"sliver_eagle_earned,omitempty"`
	DeadCount            int    `json:"dead_count,omitempty"`
	Kd                   string `json:"kd,omitempty"`
}

const templateStr = `
是否被封禁: {{.Banned}}
游戏昵称: {{.Nick}}
联队: {{.Clan}}
注册时间: {{.RegisterDate}}
等级: {{.Level}}
头衔: {{.Title}}

街机任务数: {{.StatAb.TotalMission}}
街机胜率: {{.StatAb.WinRate}}
街机KD: {{.StatAb.Kd}}
历史任务数: {{.StatRb.TotalMission}}
历史胜率: {{.StatRb.WinRate}}
陆历KD: {{.StatRb.Kd}}
全真任务数: {{.StatSb.TotalMission}}
全真胜率: {{.StatSb.WinRate}}
全真KD: {{.StatSb.Kd}}

ThunderSkill街机效率值: {{.TsABRate}}
ThunderSkill历史效率值: {{.TsRBRate}}
ThunderSkill全真效率值: {{.TsSBRate}}
（初次使用需要到TS官网生成，否则均为0）

数据最后刷新时间: {{.UpdatedAt}}

Tips: 输入”.cqbot 刷新 {{.Nick}}“可以刷新游戏数据
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
