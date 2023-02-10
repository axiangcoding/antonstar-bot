package service

import (
	"errors"
	"fmt"
	"github.com/axiangcoding/antonstar-bot/data/display"
	"github.com/axiangcoding/antonstar-bot/data/table"
	"github.com/axiangcoding/antonstar-bot/logging"
	"github.com/axiangcoding/antonstar-bot/service/bot"
	"github.com/axiangcoding/antonstar-bot/service/cqhttp"
	"github.com/axiangcoding/antonstar-bot/service/crawler"
	"github.com/google/uuid"
	"github.com/panjf2000/ants/v2"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
	"hash/crc32"
	"strings"
	"time"
)

// DrawNumber 抽一个数字
func DrawNumber(id int64, now time.Time) int32 {
	date := now.Format("2006-01-02")
	sprintf := fmt.Sprintf("%d+%s", id, date)
	hash := crc32.ChecksumIEEE([]byte(sprintf))
	return rand.New(rand.NewSource(uint64(hash))).Int31n(101)
}

func NumberBasedResponse(number int32, template int) string {
	if number == 0 {
		return bot.SelectStaticMessage(template).LuckResp.Is0
	} else if number <= 30 {
		return bot.SelectStaticMessage(template).LuckResp.Between0130
	} else if number <= 50 {
		return bot.SelectStaticMessage(template).LuckResp.Between3050
	} else if number <= 70 {
		return bot.SelectStaticMessage(template).LuckResp.Between5070
	} else if number <= 80 {
		return bot.SelectStaticMessage(template).LuckResp.Between7080
	} else if number <= 95 {
		return bot.SelectStaticMessage(template).LuckResp.Between8095
	} else if number < 100 {
		return bot.SelectStaticMessage(template).LuckResp.Between95100
	} else {
		return bot.SelectStaticMessage(template).LuckResp.Is100
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

// RefreshWTUserInfo 刷新游戏数据
func RefreshWTUserInfo(nickname string, sendForm cqhttp.SendGroupMsgForm) (*string, error) {
	missionId := uuid.NewString()
	form := ScheduleForm{
		SendForm: sendForm,
		Nick:     nickname,
	}
	if err := SubmitMissionWithDetail(missionId, table.MissionTypeUserInfo, form); err != nil {
		return nil, err
	}
	if err := ants.Submit(func() {
		if err := crawler.GetProfileFromWTOfficial(nickname,
			func(status int, user *table.GameUser) {
				switch status {
				case crawler.StatusQueryFailed:
					MustFinishMissionWithResult(missionId, table.MissionStatusFailed, CrawlerResult{
						Found: false,
						Nick:  nickname,
					})
				case crawler.StatusNotFound:
					MustPutRefreshFlag(nickname)
					MustFinishMissionWithResult(missionId, table.MissionStatusSuccess, CrawlerResult{
						Found: false,
						Nick:  nickname,
					})
				case crawler.StatusFound:
					// live, psn等用户的昵称在html中会被cf认为是邮箱而隐藏，这里需要覆盖爬取来的数据
					user.Nick = nickname
					_, err := FindGameProfile(nickname)
					if err != nil {
						if errors.Is(err, gorm.ErrRecordNotFound) {
							MustSaveGameProfile(user)
						} else {
							logging.L().Warn("find game profile failed", logging.Error(err))
						}
					} else {
						MustUpdateGameProfile(nickname, user)
					}

					if err := crawler.GetProfileFromThunderskill(nickname, func(status int, skill *crawler.ThunderSkillResp) {
						skillData := skill.Stats
						data, err := FindGameProfile(nickname)
						data.TsSBRate = skillData.S.Kpd
						data.TsRBRate = skillData.R.Kpd
						data.TsABRate = skillData.A.Kpd
						if err == nil {
							MustUpdateGameProfile(nickname, data)
						}
					}); err != nil {
						logging.L().Warn("failed on update thunder skill profile. ", logging.Error(err))
					}
					MustPutRefreshFlag(nickname)
					MustFinishMissionWithResult(missionId, table.MissionStatusSuccess, CrawlerResult{
						Found: true,
						Nick:  nickname,
						Data:  *user},
					)
				}
			}); err != nil {
			logging.L().Warn("start crawler failed. ", logging.Error(err))
		}
	}); err != nil {
		return nil, err
	}
	return &missionId, nil
}

func DoActionQuery(retMsgForm *cqhttp.SendGroupMsgForm, value string, fullMsg bool) {
	if IsStopGlobalQuery() {
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.StopGlobalQuery
		return
	}

	if value == "我" {
		config := MustFindUserConfig(retMsgForm.UserId)
		if config.BindingGameNick != nil && *config.BindingGameNick != "" {
			value = *config.BindingGameNick
		} else {
			retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.BindingFirst
			return
		}
	}

	if !IsValidNickname(value) {
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.NotValidNickname
		return
	}
	// 检查群查询限制
	if limit, usage, total := CheckGroupTodayQueryLimit(retMsgForm.GroupId); limit {
		retMsgForm.Message = fmt.Sprintf(bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.TodayGroupQueryLimit, usage, total)
		return
	}
	// 检查qq查询限制
	if limit, usage, total := CheckUserTodayQueryLimit(retMsgForm.UserId); limit {
		retMsgForm.Message = fmt.Sprintf(bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.TodayUserQueryLimit, usage, total)
		return
	}
	mId, user, err := QueryWTGamerProfile(value, *retMsgForm)
	if err != nil {
		logging.L().Warn("query WT gamer profile error", logging.Error(err))
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.CanNotRefresh
	}
	if mId != nil {
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.QueryIsRunning
		if err := ants.Submit(func() {
			if err := WaitForCrawlerFinished(*mId, fullMsg); err != nil {
				logging.L().Error("wait for callback error", logging.Error(err))
			}
		}); err != nil {
			logging.L().Error("submit ant job failed", logging.Error(err))
		}
	} else {
		if fullMsg {
			retMsgForm.Message = user.ToFriendlyFullString()
		} else {
			retMsgForm.Message = user.ToFriendlyShortString()
		}
	}
	MustAddUserConfigTodayQueryCount(retMsgForm.UserId, 1)
	MustAddUserConfigTotalQueryCount(retMsgForm.UserId, 1)
	MustAddGroupConfigTodayQueryCount(retMsgForm.GroupId, 1)
	MustAddGroupConfigTotalQueryCount(retMsgForm.GroupId, 1)
}

func DoActionRefresh(retMsgForm *cqhttp.SendGroupMsgForm, value string) {
	if IsStopGlobalQuery() {
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.StopGlobalQuery
		return
	}
	if value == "我" {
		config := MustFindUserConfig(retMsgForm.UserId)
		if config.BindingGameNick != nil && *config.BindingGameNick != "" {
			value = *config.BindingGameNick
		}
	}

	if !IsValidNickname(value) {
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.NotValidNickname
		return
	}
	if !CanBeRefresh(value) {
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.TooShortToRefresh
		return
	}
	// 检查群查询限制
	if limit, usage, total := CheckGroupTodayQueryLimit(retMsgForm.GroupId); limit {
		retMsgForm.Message = fmt.Sprintf(bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.TodayGroupQueryLimit, usage, total)
		return
	}
	// 检查qq查询限制
	if limit, usage, total := CheckUserTodayQueryLimit(retMsgForm.UserId); limit {
		retMsgForm.Message = fmt.Sprintf(bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.TodayUserQueryLimit, usage, total)
		return
	}
	missionId, err := RefreshWTUserInfo(value, *retMsgForm)
	if err != nil {
		logging.L().Warn("refresh WT gamer profile error", logging.Error(err))
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.CanNotRefresh
	}
	retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.QueryIsRunning
	if err := ants.Submit(func() {
		if err := WaitForCrawlerFinished(*missionId, false); err != nil {
			logging.L().Error("wait for callback error", logging.Error(err))
		}
	}); err != nil {
		logging.L().Error("submit ant job failed", logging.Error(err))
	}
	MustAddUserConfigTodayQueryCount(retMsgForm.UserId, 1)
	MustAddUserConfigTotalQueryCount(retMsgForm.UserId, 1)
	MustAddGroupConfigTodayQueryCount(retMsgForm.GroupId, 1)
	MustAddGroupConfigTotalQueryCount(retMsgForm.GroupId, 1)
}

func DoActionDrawCard(retMsgForm *cqhttp.SendGroupMsgForm, value string, id int64) {
	retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.DrawCard
}

func DoActionLuck(retMsgForm *cqhttp.SendGroupMsgForm, value string, id int64) {
	number := DrawNumber(id, time.Now().In(time.FixedZone("CST", 8*3600)))
	retMsgForm.Message = fmt.Sprintf(bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.Luck, number, NumberBasedResponse(number, retMsgForm.MessageTemplate))
}

func DoActionGroupStatus(retMsgForm *cqhttp.SendGroupMsgForm) {
	config := MustFindGroupConfig(retMsgForm.GroupId)
	retMsgForm.Message = config.ToDisplay().ToFriendlyString()
}

func DoActionData(retMsgForm *cqhttp.SendGroupMsgForm, value string) {
	botQueryPrefix := ".cqbot 数据 "
	retMsgForm.MessagePrefix = ""
	opt1 := "导弹数据"
	switch value {
	case opt1:
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.MissileData
	default:
		var lst []string
		lst = append(lst, botQueryPrefix+opt1)
		retMsgForm.Message = fmt.Sprintf(bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.DataOptions, strings.Join(lst, "\n"))
	}
}

func DoActionBinding(retMsgForm *cqhttp.SendGroupMsgForm, value string) {
	profile, err := FindGameProfile(value)
	if err != nil {
		logging.L().Warn("dal failed", logging.Error(err))
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.BindingNickNotExist
		return
	}

	config, err := FindUserConfig(retMsgForm.UserId)
	if err != nil {
		logging.L().Warn("dal failed", logging.Error(err))
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.BindingError
		return
	}
	if config.BindingGameNick != nil && *config.BindingGameNick != "" {
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.BindingExist
		return
	}
	config.BindingGameNick = &profile.Nick
	if err := SaveUserConfig(*config); err != nil {
		logging.L().Warn("dal failed", logging.Error(err))
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.BindingError
		return
	}
	retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.BindingSuccess
}

func DoActionUnbinding(retMsgForm *cqhttp.SendGroupMsgForm) {
	if err := UpdateUserConfigBindingGameNick(retMsgForm.UserId, nil); err != nil {
		logging.L().Warn("dal failed", logging.Error(err))
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.UnbindingError
		return
	}
	retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.UnbindingSuccess
}

func DoActionManager(retMsgForm *cqhttp.SendGroupMsgForm, uc *table.QQUserConfig, value string) {
	// 只有超级管理员可以进行全局设置
	if uc.SuperAdmin == nil || !*uc.SuperAdmin {
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.ConfNotPermit
		return
	}
	botQueryPrefix := ".cqbot 管理 "
	keyCloseResponse := "关闭回复"
	keyOpenResponse := "开启回复"
	keyOpenQuery := "开启查询"
	keyCloseQuery := "关闭查询"
	keySetAdmin := "添加管理员"
	keyUnsetAdmin := "解除管理员"
	switch value {
	case keyOpenResponse:
		MustUpsertGlobalConfig(table.ConfigStopAllResponse, "false")
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.ConfStartGlobalResponse
	case keyCloseResponse:
		MustUpsertGlobalConfig(table.ConfigStopAllResponse, "true")
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.ConfStopGlobalResponse
	case keyOpenQuery:
		MustUpsertGlobalConfig(table.ConfigStopQuery, "false")
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.ConfStartGlobalQuery
	case keyCloseQuery:
		MustUpsertGlobalConfig(table.ConfigStopQuery, "true")
		retMsgForm.Message = bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.ConfStopGlobalQuery
	case keySetAdmin:
	// 	TODO
	case keyUnsetAdmin:
	// 	TODO
	default:
		var lst []string
		lst = append(lst, botQueryPrefix+keyCloseResponse)
		lst = append(lst, botQueryPrefix+keyOpenResponse)
		lst = append(lst, botQueryPrefix+keyOpenQuery)
		lst = append(lst, botQueryPrefix+keyCloseQuery)
		lst = append(lst, botQueryPrefix+keySetAdmin)
		lst = append(lst, botQueryPrefix+keyUnsetAdmin)
		retMsgForm.Message = fmt.Sprintf(bot.SelectStaticMessage(retMsgForm.MessageTemplate).CommonResp.ConfOptions, strings.Join(lst, "\n"))
	}
}

func DoActionGroupManager(retMsgForm *cqhttp.SendGroupMsgForm, uc *table.QQUserConfig, value string) {
	// TODO 群管理
}
