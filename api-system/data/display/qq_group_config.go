package display

import (
	"bytes"
	"fmt"
	"github.com/axiangcoding/ax-web/logging"
	"text/template"
)

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
停用全部功能: {{.Shutdown}}
允许管理员配置: {{.AllowAdminConfig}}
启用战绩查询功能: {{.EnableActionQuery}}
启用气运功能: {{.EnableActionLuck}}
启用配置设置功能: {{.EnableActionSetting}}
启用直播间检查功能: {{.EnableCheckBiliRoom}}
语气类型: {{.MessageTemplate}}
{{if .Banned}}==== 本群已被禁用功能 ===={{end}}
`

func (c QQGroupConfig) ToFriendlyString() string {
	fmt.Println(c.Banned)
	var buf bytes.Buffer

	t, err := template.New("display").Parse(templateGroupSettingStr)
	if err != nil {
		logging.Warn(err)
		return "Error"
	}
	if err := t.Execute(&buf, c); err != nil {
		logging.Error(err)
		return "Error"
	}
	return buf.String()
}
