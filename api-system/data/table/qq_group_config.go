package table

import "gorm.io/gorm"

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
	}
}
