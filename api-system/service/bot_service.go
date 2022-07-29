package service

import (
	"errors"
	"fmt"
	"github.com/axiangcoding/ax-web/data"
	"github.com/axiangcoding/ax-web/data/display"
	"github.com/axiangcoding/ax-web/data/table"
	"github.com/google/uuid"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
	"hash/crc32"
	"time"
)

// DrawNumber 抽一个数字
func DrawNumber(id int, now time.Time) int32 {
	date := now.Format("2006-01-02")
	sprintf := fmt.Sprintf("%d+%s", id, date)
	hash := crc32.ChecksumIEEE([]byte(sprintf))
	return rand.New(rand.NewSource(uint64(hash))).Int31n(101)
}

// QueryWTGamerProfile 查询系统中已有的玩家的游戏资料。如果资料不存在，则调用爬虫爬取
// FIXME: 如果达到限制次数，应该直接告知上层应用
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

// RefreshWTGamerProfile 请求爬虫，爬取玩家的游戏资料
// FIXME: 限制同一天内对同个nickname的请求次数；保存数据到game_user表中
func RefreshWTGamerProfile(nickname string) error {
	missionId := uuid.NewString()

	defaultProject := "crawler"

	form := ScheduleForm{
		Project:   defaultProject,
		Spider:    "gaijin",
		MissionId: missionId,
		Nick:      nickname,
	}
	if err := SubmitMission(missionId, table.MissionTypeWTProfile, form); err != nil {
		return err
	}
	if err := RequestCrawlerSpider(form); err != nil {
		return err
	}
	if err := RunningMission(missionId, 10); err != nil {
		return err
	}
	return nil
}
