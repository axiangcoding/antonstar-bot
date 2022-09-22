package service

import (
	"github.com/axiangcoding/ax-web/data"
	"github.com/axiangcoding/ax-web/data/table"
)

func FindGroupConfig(groupId int64) (*table.QQGroupConfig, error) {
	db := data.GetDB()
	var find table.QQGroupConfig
	if err := db.Where(table.QQGroupConfig{GroupId: groupId}).Take(&find).Error; err != nil {
		return nil, err
	}
	return &find, nil
}

func SaveGroupConfig(gc table.QQGroupConfig) error {
	db := data.GetDB()
	if err := db.Save(&gc).Error; err != nil {
		return err
	}
	return nil
}
