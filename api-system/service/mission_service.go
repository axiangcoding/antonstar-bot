package service

import (
	"encoding/json"
	"github.com/axiangcoding/ax-web/data"
	"github.com/axiangcoding/ax-web/data/table"
	"time"
)

func FindMission(missionId string) (*table.Mission, error) {
	db := data.GetDB()
	var find table.Mission
	if err := db.Where(table.Mission{MissionId: missionId}).Take(&find).Error; err != nil {
		return nil, err
	}
	return &find, nil
}

func SubmitMission(missionId string, missionType string, detail any) error {
	db := data.GetDB()
	bytes, err := json.Marshal(detail)
	if err != nil {
		return err
	}
	mission := table.Mission{
		MissionId: missionId,
		Type:      missionType,
		Status:    table.MissionStatusPending,
		Process:   0,
		Detail:    string(bytes),
	}
	if err := db.Save(&mission).Error; err != nil {
		return err
	}
	return nil
}

func RunningMission(missionId string, process float64) error {
	db := data.GetDB()
	if err := db.Where(table.Mission{MissionId: missionId}).
		Updates(table.Mission{Status: table.MissionStatusRunning, Process: process}).Error; err != nil {
		return err
	}
	return nil
}

func FinishMission(missionId string, status string, result any) error {
	db := data.GetDB()

	bytes, err := json.Marshal(result)
	if err != nil {
		return err
	}
	update := table.Mission{
		Status:       status,
		Process:      100,
		Result:       string(bytes),
		FinishedTime: time.Now(),
	}
	if err := db.Where(table.Mission{MissionId: missionId}).
		Updates(update).Error; err != nil {
		return err
	}
	return nil
}
