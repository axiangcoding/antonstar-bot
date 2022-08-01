package service

import (
	"errors"
	"github.com/axiangcoding/ax-web/data"
	"github.com/axiangcoding/ax-web/data/table"
	"gorm.io/gorm"
	"time"
)

func FindGameProfile(nick string) (*table.GameUser, error) {
	db := data.GetDB()
	var find table.GameUser
	if err := db.Where(table.GameUser{Nick: nick}).Take(&find).Error; err != nil {
		return nil, err
	}
	return &find, nil
}

func SaveGameProfile(gameUser table.GameUser) error {
	db := data.GetDB()
	if err := db.Save(&gameUser).Error; err != nil {
		return err
	}
	return nil
}

func UpdateGameProfile(nick string, user table.GameUser) error {
	db := data.GetDB()
	if err := db.Where(table.GameUser{Nick: nick}).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

func CanBeRefresh(nick string) bool {
	profile, err := FindGameProfile(nick)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true
		} else {
			return false
		}
	}
	if time.Now().Sub(profile.UpdatedAt) > time.Hour*24 {
		return true
	}
	return false
}
