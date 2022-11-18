package table

import "gorm.io/gorm"

const (
	ConfigStopQuery = "stop_query"
)

type GlobalConfig struct {
	gorm.Model
	Key   string `gorm:"uniqueIndex;"`
	Value string
}
