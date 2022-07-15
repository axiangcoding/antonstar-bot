package table

import (
	"gorm.io/gorm"
	"time"
)

const (
	MissionTypeWTProfile = "wt_profile"
)

const (
	MissionStatusUnknown = "unknown"
	MissionStatusPending = "pending"
	MissionStatusRunning = "running"
	MissionStatusSuccess = "success"
	MissionStatusFailed  = "failed"
)

type Mission struct {
	gorm.Model
	MissionId    string `gorm:"uniqueIndex;size:255"`
	Type         string `gorm:"index;size:255"`
	Status       string `gorm:"index;size:255"`
	FinishedTime time.Time
	Process      float64
	Detail       string
	Result       string
}
