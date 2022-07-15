package service

import (
	"encoding/json"
	"errors"
	"github.com/axiangcoding/ax-web/data"
	"github.com/axiangcoding/ax-web/data/display"
	"github.com/axiangcoding/ax-web/data/table"
	"github.com/axiangcoding/ax-web/mq"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func QueryWTGamerProfile(nickname string) (bool, *display.GameUser, error) {
	db := data.GetDB()
	find := table.GameUser{}
	err := db.Take(&find, table.GameUser{Nick: nickname}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := RefreshWTGamerProfile(nickname); err != nil {
			return false, nil, err
		}
		return true, nil, nil
	} else {
		user := find.ToDisplayGameUser()
		return false, &user, nil
	}

}

func RefreshWTGamerProfile(nickname string) error {
	db := data.GetDB()
	missionId := uuid.NewString()
	cb := mq.CrawBody{
		MissionId: missionId,
		Nickname:  nickname,
	}
	detailJson, err := json.Marshal(cb)
	if err != nil {
		return err
	}

	mission := table.Mission{
		MissionId: missionId,
		Type:      table.MissionTypeWTProfile,
		Status:    table.MissionStatusPending,
		Process:   0,
		Detail:    string(detailJson),
	}
	if err := SendMessage(cb); err != nil {
		return err
	}
	if err := db.Save(&mission).Error; err != nil {
		return err
	}
	return nil
}
