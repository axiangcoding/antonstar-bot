package service

import (
	"errors"
	"fmt"
	"github.com/axiangcoding/ax-web/data/display"
	"github.com/axiangcoding/ax-web/data/table"
	"github.com/axiangcoding/ax-web/logging"
	"github.com/axiangcoding/ax-web/service/cqhttp"
	"github.com/axiangcoding/ax-web/tool"
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
func QueryWTGamerProfile(nickname string, sendForm cqhttp.SendGroupMsgForm) (*string, *display.GameUser, error) {
	find, err := FindGameProfile(nickname)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil, err
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		if missionId, err := RefreshWTUserInfo(nickname, sendForm); err != nil {
			return nil, nil, err
		} else {
			return missionId, nil, nil
		}
	} else {
		user := find.ToDisplayGameUser()
		return nil, &user, nil
	}
}

func RefreshWTUserInfo(nickname string, sendForm cqhttp.SendGroupMsgForm) (*string, error) {
	missionId := uuid.NewString()
	form := ScheduleForm{
		SendForm: sendForm,
		Nick:     nickname,
	}
	if err := SubmitMissionWithDetail(missionId, table.MissionTypeUserInfo, form); err != nil {
		return nil, err
	}
	tool.GoWithRecover(func() {
		if err := GetUserInfoFromWarThunder(missionId, nickname); err != nil {
			logging.Warn("start crawler failed. ", err)
		}
	})
	return &missionId, nil
}
