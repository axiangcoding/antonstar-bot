package table

import "gorm.io/gorm"

const (
	ConfigStopQuery       = "stop_query"
	ConfigStopAllResponse = "stop_all_response"
)

type GlobalConfig struct {
	gorm.Model
	Key   string `gorm:"uniqueIndex;"`
	Value string
}
