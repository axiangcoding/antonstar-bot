package service

import (
	"encoding/json"
	"github.com/axiangcoding/antonstar-bot/data/dal"
	"github.com/axiangcoding/antonstar-bot/data/table"
	"github.com/axiangcoding/antonstar-bot/logging"
	"time"
)

func FindMission(missionId string) (*table.Mission, error) {
	take, err := dal.Q.Mission.Where(dal.Mission.MissionId.Eq(missionId)).Take()
	if err != nil {
		return nil, err
	}
	return take, nil
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

func MustFinishMissionWithResult(missionId string, status string, result any) {
	bytes, _ := json.Marshal(result)
	update := table.Mission{
		Status:       status,
		Process:      100,
		Result:       string(bytes),
		FinishedTime: time.Now(),
	}
	if _, err := dal.Q.Mission.
		Where(dal.Mission.MissionId.Eq(missionId)).
		Updates(update); err != nil {
		logging.L().Warn("dal failed", logging.Error(err))
	}
}
