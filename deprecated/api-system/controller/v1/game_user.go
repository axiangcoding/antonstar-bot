package v1

import (
	"axiangcoding/antonstar/api-system/entity/app"
	"axiangcoding/antonstar/api-system/entity/e"
	"axiangcoding/antonstar/api-system/service"
	"github.com/gin-gonic/gin"
)

// GetGameUsers
// @Summary  查询游戏昵称的所有query_id
// @Tags     GameUser API
// @Param    form  query     app.Pagination  true  "param"
// @Success  200   {object}  app.ApiJson     ""
// @Router   /v1/game_users [get]
func GetGameUsers(c *gin.Context) {
	var pagination app.Pagination
	err := c.ShouldBindQuery(&pagination)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}
	users, err := service.GetGameUsers(c, pagination)
	if err != nil {
		app.BizFailed(c, e.Error, err)
		return
	}
	app.Success(c, users)
}
