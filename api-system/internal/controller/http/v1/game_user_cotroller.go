package v1

import (
	"errors"
	"github.com/axiangcoding/antonstar-bot/internal/data/display"
	"github.com/axiangcoding/antonstar-bot/internal/data/table"
	"github.com/axiangcoding/antonstar-bot/internal/entity/app"
	"github.com/axiangcoding/antonstar-bot/internal/entity/e"
	"github.com/axiangcoding/antonstar-bot/internal/service"
	"github.com/axiangcoding/antonstar-bot/pkg/crawler"
	"github.com/axiangcoding/antonstar-bot/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/panjf2000/ants/v2"
	"gorm.io/gorm"
)

type ProfileResp struct {
	Found   bool              `json:"found"`
	Profile *display.GameUser `json:"profile,omitempty"`
}

// GameUserProfile
// @Summary  获取游戏内玩家数据
// @Tags     GameUser API
// @Param    nick  query     string       true  "user nickname"
// @Success  200   {object}  app.ApiJson  ""
// @Router   /v1/wt/profile [get]
func GameUserProfile(c *gin.Context) {
	nick := c.Query("nick")
	profile, err := service.FindGameProfile(nick)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			app.Success(c, ProfileResp{
				Found: false,
			})
			return
		}
		app.BizFailed(c, e.Error, err)
		return
	}
	displayGameUser := profile.ToDisplayGameUser()
	app.Success(c, ProfileResp{
		Found:   true,
		Profile: &displayGameUser,
	})
}

// UpdateGameUserProfile
// @Summary  更新游戏内玩家数据
// @Tags     GameUser API
// @Param    nick  query     string       true  "user nickname"
// @Success  200   {object}  app.ApiJson  ""
// @Router   /v1/wt/profile/update [post]
func UpdateGameUserProfile(c *gin.Context) {
	nickname := c.Query("nick")

	if !service.IsValidNickname(nickname) {
		app.Success(c, map[string]any{
			"refresh": false,
			"reason":  "not a valid nickname",
		})
		return
	}
	if !service.CanBeRefresh(nickname) {
		app.Success(c, map[string]any{
			"refresh": false,
			"reason":  "refresh too often",
		})
		return
	}

	missionId := uuid.NewString()
	form := service.ScheduleForm{
		Nick: nickname,
	}
	if err := service.SubmitMissionWithDetail(missionId, table.MissionTypeUserInfo, form); err != nil {
		app.BizFailed(c, e.Error, err)
		return
	}
	if err := ants.Submit(func() {
		if err := crawler.GetProfileFromWTOfficial(nickname,
			func(status int, user *table.GameUser) {
				switch status {
				case crawler.StatusQueryFailed:
					service.MustFinishMissionWithResult(missionId, table.MissionStatusFailed, service.CrawlerResult{
						Found: false,
						Nick:  nickname,
					})
				case crawler.StatusNotFound:
					service.MustPutRefreshFlag(nickname)
					service.MustFinishMissionWithResult(missionId, table.MissionStatusSuccess, service.CrawlerResult{
						Found: false,
						Nick:  nickname,
					})
				case crawler.StatusFound:
					// live, psn等用户的昵称在html中会被cf认为是邮箱而隐藏，这里需要覆盖爬取来的数据
					user.Nick = nickname
					_, err := service.FindGameProfile(nickname)
					if err != nil {
						if errors.Is(err, gorm.ErrRecordNotFound) {
							service.MustSaveGameProfile(user)
						} else {
							logging.L().Warn("find game profile failed", logging.Error(err))
						}
					} else {
						service.MustUpdateGameProfile(nickname, user)
					}

					if err := crawler.GetProfileFromThunderskill(nickname, func(status int, skill *crawler.ThunderSkillResp) {
						skillData := skill.Stats
						data, err := service.FindGameProfile(nickname)
						data.TsSBRate = skillData.S.Kpd
						data.TsRBRate = skillData.R.Kpd
						data.TsABRate = skillData.A.Kpd
						if err == nil {
							service.MustUpdateGameProfile(nickname, data)
						}
					}); err != nil {
						logging.L().Warn("failed on update thunder skill profile. ", logging.Error(err))
					}
					service.MustPutRefreshFlag(nickname)
					service.MustFinishMissionWithResult(missionId, table.MissionStatusSuccess, service.CrawlerResult{
						Found: true,
						Nick:  nickname,
						Data:  *user},
					)
				}
			}); err != nil {
			logging.L().Warn("start crawler failed. ", logging.Error(err))
		}
	}); err != nil {
		app.BizFailed(c, e.Error, err)
		return
	}
	app.Success(c, map[string]any{
		"refresh":   true,
		"missionId": missionId,
	})
}
