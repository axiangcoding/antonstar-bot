package v1

import (
	"github.com/axiangcoding/antonstar-bot/entity/app"
	"github.com/axiangcoding/antonstar-bot/entity/e"
	"github.com/axiangcoding/antonstar-bot/service"
	"github.com/gin-gonic/gin"
)

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
		app.BizFailed(c, e.Error, err)
		return
	}
	displayGameUser := profile.ToDisplayGameUser()
	app.Success(c, displayGameUser)
}
