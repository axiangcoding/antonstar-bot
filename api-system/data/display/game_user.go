package display

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

	GroundRateAb UserRate
	GroundRateRb UserRate
	GroundRateSb UserRate

	AviationRateAb UserRate
	AviationRateRb UserRate
	AviationRateSb UserRate

	FleetRateAb UserRate
	FleetRateRb UserRate
	FleetRateSb UserRate

	TsABRate float64 `json:"ts_ab_rate"`
	TsRBRate float64 `json:"ts_rb_rate"`
	TsSBRate float64 `json:"ts_sb_rate"`
	AsABRate float64 `json:"as_ab_rate"`
	AsRBRate float64 `json:"as_rb_rate"`
	AsSBRate float64 `json:"as_sb_rate"`
	Banned   bool    `json:"banned"`
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

type UserRate struct {
	Ka string `json:"ka,omitempty"`
}

const templateShortStr = `
{{if .Banned}}==== 已被封禁 ===={{end}}
游戏昵称: {{.Nick}}
联队: {{.Clan}}
注册时间: {{.RegisterDate}}
等级: {{.Level}}
头衔: {{.Title}}
{{if .Banned}}==== 已被封禁 ===={{end}}

街机任务数: {{.StatAb.TotalMission}}
街机胜率: {{.StatAb.WinRate}}
街机KD: {{.StatAb.Kd}}
历史任务数: {{.StatRb.TotalMission}}
历史胜率: {{.StatRb.WinRate}}
历史KD: {{.StatRb.Kd}}
全真任务数: {{.StatSb.TotalMission}}
全真胜率: {{.StatSb.WinRate}}
全真KD: {{.StatSb.Kd}}

（ThunderSkill效率值需要到其网站更新）
TS街机效率: {{.TsABRate}}%
TS历史效率: {{.TsRBRate}}%
TS全真效率: {{.TsSBRate}}%

数据最后刷新时间: {{.UpdatedAt}}
`

const templateFullStr = `
{{if .Banned}}==== 已被封禁 ===={{end}}
游戏昵称: {{.Nick}}
联队: {{.Clan}}
注册时间: {{.RegisterDate}}
等级: {{.Level}}
头衔: {{.Title}}
{{if .Banned}}==== 已被封禁 ===={{end}}

街机任务数: {{.StatAb.TotalMission}}
街机胜率: {{.StatAb.WinRate}}
街机KD: {{.StatAb.Kd}}
街机游戏时间: {{.StatAb.GameTime}}
历史任务数: {{.StatRb.TotalMission}}
历史胜率: {{.StatRb.WinRate}}
历史KD: {{.StatRb.Kd}}
历史游戏时间: {{.StatRb.GameTime}}
全真任务数: {{.StatSb.TotalMission}}
全真胜率: {{.StatSb.WinRate}}
全真KD: {{.StatSb.Kd}}
全真游戏时间: {{.StatSb.GameTime}}

（击杀数/出击数简称为'KA'）
空战街机KA: {{.AviationRateAb.Ka}}
空战历史KA: {{.AviationRateRb.Ka}}
空战全真KA: {{.AviationRateSb.Ka}}

陆战街机KA: {{.GroundRateAb.Ka}}
陆战历史KA: {{.GroundRateRb.Ka}}
陆战全真KA: {{.GroundRateSb.Ka}}

海战街机KA: {{.FleetRateAb.Ka}}
海战历史KA: {{.FleetRateRb.Ka}}
海战全真KA: {{.FleetRateSb.Ka}}

（ThunderSkill效率值需要到其网站更新）
TS街机效率: {{.TsABRate}}%
TS历史效率: {{.TsRBRate}}%
TS全真效率: {{.TsSBRate}}%

数据最后刷新时间: {{.UpdatedAt}}
`

func (u GameUser) ToFriendlyShortString() string {
	return parseTemplate(templateShortStr, u)
}

func (u GameUser) ToFriendlyFullString() string {
	return parseTemplate(templateFullStr, u)
}
