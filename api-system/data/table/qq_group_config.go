package table

import (
	"github.com/axiangcoding/antonstar-bot/data/display"
	"gorm.io/gorm"
)

type QQGroupConfig struct {
	gorm.Model
	GroupId             int64 `gorm:"uniqueIndex;"`
	BindBiliRoomId      int64
	Banned              *bool
	AllowAdminConfig    *bool
	Shutdown            *bool
	EnableActionQuery   *bool
	EnableActionLuck    *bool
	EnableActionSetting *bool
	EnableCheckBiliRoom *bool
	MessageTemplate     int
	TodayQueryCount     int
	OneDayQueryLimit    int
	TotalQueryCount     int
	TodayUsageCount     int
	OneDayUsageLimit    int
	TotalUsageCount     int
}

func DefaultGroupConfig(groupId int64) QQGroupConfig {
	trueVal := true
	falseVal := false
	return QQGroupConfig{
		GroupId:             groupId,
		Banned:              &falseVal,
		AllowAdminConfig:    &trueVal,
		Shutdown:            &falseVal,
		EnableActionQuery:   &trueVal,
		EnableActionLuck:    &trueVal,
		EnableActionSetting: &falseVal,
		EnableCheckBiliRoom: &falseVal,
		TodayQueryCount:     0,
		OneDayQueryLimit:    50,
		TotalQueryCount:     0,
		TodayUsageCount:     0,
		OneDayUsageLimit:    200,
		TotalUsageCount:     0,
	}
}

func (c QQGroupConfig) ToDisplay() display.QQGroupConfig {
	return display.QQGroupConfig{
		GroupId:             c.GroupId,
		BindBiliRoomId:      c.BindBiliRoomId,
		Banned:              *c.Banned,
		AllowAdminConfig:    *c.AllowAdminConfig,
		Shutdown:            *c.Shutdown,
		EnableActionQuery:   *c.EnableActionQuery,
		EnableActionLuck:    *c.EnableActionLuck,
		EnableActionSetting: *c.EnableActionSetting,
		EnableCheckBiliRoom: *c.EnableCheckBiliRoom,
		MessageTemplate:     c.MessageTemplate,
	}
}
