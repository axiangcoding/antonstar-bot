package display

type QQGroupConfig struct {
	GroupId             int64 `gorm:"uniqueIndex;"`
	BindBiliRoomId      int64
	Banned              bool
	AllowAdminConfig    bool
	Shutdown            bool
	EnableActionQuery   bool
	EnableActionLuck    bool
	EnableActionSetting bool
	EnableCheckBiliRoom bool
	MessageTemplate     int
}

const templateGroupSettingStr = `
{{if .Banned}}==== 本群已被禁用功能 ===={{end}}
QQ群号: {{.GroupId}}
绑定直播间号: https://live.bilibili.com/{{.BindBiliRoomId}}
停用全部功能: {{if .Shutdown}} 是 {{else}} 否 {{end}}
允许管理员配置: {{if .AllowAdminConfig}} 是 {{else}} 否 {{end}}
启用战绩查询功能: {{if .EnableActionQuery}} 是 {{else}} 否 {{end}}
启用气运功能: {{if .EnableActionLuck}} 是 {{else}} 否 {{end}}
启用配置设置功能: {{if .EnableActionSetting}} 是 {{else}} 否 {{end}}
启用直播间检查功能: {{if .EnableCheckBiliRoom}} 是 {{else}} 否 {{end}}
语气类型: {{.MessageTemplate}}
{{if .Banned}}==== 本群已被禁用功能 ===={{end}}
`

func (c QQGroupConfig) ToFriendlyString() string {
	return parseTemplate(templateGroupSettingStr, c)
}
