package service

import (
	"encoding/json"
	"github.com/axiangcoding/ax-web/data"
	"github.com/axiangcoding/ax-web/data/dal"
	"github.com/axiangcoding/ax-web/data/table"
	"github.com/axiangcoding/ax-web/logging"
	"time"
)

func FindMission(missionId string) (*table.Mission, error) {
	take, err := dal.Q.Mission.Where(dal.Mission.MissionId.Eq(missionId)).Take()
	if err != nil {
		return nil, err
	}
	return take, nil
}

func SubmitMission(missionId string, missionType string) error {
	mission := table.Mission{
		MissionId: missionId,
		Type:      missionType,
		Status:    table.MissionStatusPending,
		Process:   0,
	}
	if err := dal.Q.Mission.Save(&mission); err != nil {
		return err
	}
	return nil
}

func SubmitMissionWithDetail(missionId string, missionType string, detail any) error {
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
	if err := dal.Q.Mission.Save(&mission); err != nil {
		return err
	}
	return nil
}

func RunningMission(missionId string, process float64) error {
	if _, err := dal.Mission.
		Where(dal.Mission.MissionId.Eq(missionId)).
		Updates(table.Mission{
			Status:  table.MissionStatusRunning,
			Process: process}); err != nil {
		return err
	}
	return nil
}

func MustRunningMission(missionId string, process float64) {
	db := data.GetDB()
	if err := db.Where(table.Mission{MissionId: missionId}).
		Updates(table.Mission{Status: table.MissionStatusRunning, Process: process}).Error; err != nil {
		logging.Warn(err)
	}
}

func MustFinishMission(missionId string, status string) {
	db := data.GetDB()
	update := table.Mission{
		Status:       status,
		Process:      100,
		FinishedTime: time.Now(),
	}
	if err := db.Where(table.Mission{MissionId: missionId}).
		Updates(update).Error; err != nil {
		logging.Warn(err)
	}
}

func MustFinishMissionWithResult(missionId string, status string, result any) {
	db := data.GetDB()
	bytes, _ := json.Marshal(result)
	update := table.Mission{
		Status:       status,
		Process:      100,
		Result:       string(bytes),
		FinishedTime: time.Now(),
	}
	if err := db.Where(table.Mission{MissionId: missionId}).
		Updates(update).Error; err != nil {
		logging.Warn(err)
	}
}

func FinishMissionWithResult(missionId string, status string, result any) error {
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
