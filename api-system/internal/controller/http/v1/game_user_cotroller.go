package v1

import (
	"errors"
	"github.com/axiangcoding/antonstar-bot/internal/data/display"
	"github.com/axiangcoding/antonstar-bot/internal/entity/app"
	"github.com/axiangcoding/antonstar-bot/internal/entity/e"
	"github.com/axiangcoding/antonstar-bot/internal/service"
	"github.com/gin-gonic/gin"
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
