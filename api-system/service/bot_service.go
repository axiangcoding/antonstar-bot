package service

import (
	"errors"
	"fmt"
	"github.com/axiangcoding/ax-web/data/display"
	"github.com/axiangcoding/ax-web/data/table"
	"github.com/axiangcoding/ax-web/service/cqhttp"
	"github.com/axiangcoding/ax-web/service/crawler"
	"github.com/google/uuid"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
	"hash/crc32"
	"time"
)

// DrawNumber 抽一个数字
func DrawNumber(id int64, now time.Time) int32 {
	date := now.Format("2006-01-02")
	sprintf := fmt.Sprintf("%d+%s", id, date)
	hash := crc32.ChecksumIEEE([]byte(sprintf))
	return rand.New(rand.NewSource(uint64(hash))).Int31n(101)
}

func NumberBasedResponse(number int32) string {
	if number == 0 {
		return "你是个诡计多端的0"
	} else if number <= 30 {
		return "就这？"
	} else if number <= 60 {
		return "至少不算丢人"
	} else if number <= 80 {
		return "有没有考虑过转发抽奖"
	} else if number < 100 {
		return "分我点运气呗"
	} else {
		return "你是一个个个100"
	}
}

// QueryWTGamerProfile 查询系统中已有的玩家的游戏资料。如果资料不存在，则调用爬虫爬取
func QueryWTGamerProfile(nickname string, sendForm cqhttp.SendGroupMsgForm) ([]string, *display.GameUser, error) {
	find, err := FindGameProfile(nickname)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		if missionIds, err := RefreshWTGamerProfile(nickname, sendForm); err != nil {
			return nil, nil, err
		} else {
			return missionIds, nil, nil
		}
	} else {
		user := find.ToDisplayGameUser()
		return nil, &user, nil
	}
}

// RefreshWTGamerProfile 请求爬虫，爬取玩家的游戏资料
func RefreshWTGamerProfile(nickname string, sendForm cqhttp.SendGroupMsgForm) ([]string, error) {
	defaultProject := "crawler"

	missionId := uuid.NewString()
	form := ScheduleForm{
		SendForm:  sendForm,
		Project:   defaultProject,
		Spider:    crawler.SourceGaijin,
		MissionId: missionId,
		Nick:      nickname,
	}
	if err := SubmitMission(missionId, table.MissionTypeWTProfile, form); err != nil {
		return nil, err
	}
	if err := RequestCrawlerSpider(form); err != nil {
		return nil, err
	}
	if err := RunningMission(missionId, 10); err != nil {
		return nil, err
	}
	// 执行查询时记录时间，不允许短期内重复刷新
	PutRefreshFlag(nickname)

	missionId2 := uuid.NewString()
	form2 := ScheduleForm{
		SendForm:  sendForm,
		Project:   defaultProject,
		Spider:    crawler.SourceThunderSkill,
		MissionId: missionId2,
		Nick:      nickname,
	}
	if err := SubmitMission(missionId2, table.MissionTypeWTProfile, form2); err != nil {
		return nil, err
	}
	if err := RequestCrawlerSpider(form2); err != nil {
		return nil, err
	}
	if err := RunningMission(missionId2, 10); err != nil {
		return nil, err
	}
	return []string{missionId, missionId2}, nil
}
