package v1

import (
	"axiangcoding/antonstar/api-system/internal/app/service"
	"axiangcoding/antonstar/api-system/pkg/app"
	"axiangcoding/antonstar/api-system/pkg/app/e"
	"github.com/gin-gonic/gin"
)

type UserInfoForm struct {
	// 游戏的昵称
	Nickname string `form:"nickname" binding:"required"`
}

// PostUserInfoQuery
// @Summary  提交一个查询游戏用户的请求
// @Tags     WarThunder
// @Param    form  query     UserInfoForm  true  "param"
// @Success  200   {object}  app.ApiJson   ""
// @Router   /v1/war_thunder/userinfo/query [post]
func PostUserInfoQuery(c *gin.Context) {
	var form UserInfoForm
	err := c.ShouldBindQuery(&form)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}
	info, err := service.RequestUserInfo(c, form.Nickname)
	if err != nil {
		app.BizFailed(c, e.Error, err)
		return
	}
	app.Success(c, info)
}

type UserInfoDetailForm struct {
	QueryID string `json:"query_id" form:"query_id" binding:"required"`
}

// GetUserInfoQuery
// @Summary  获取异步查询结果
// @Tags     WarThunder
// @Param    form  query     UserInfoDetailForm  true  "query userinfo"
// @Success  200   {object}  app.ApiJson         ""
// @Router   /v1/war_thunder/userinfo/query [get]
func GetUserInfoQuery(c *gin.Context) {
	var form UserInfoDetailForm
	err := c.ShouldBindQuery(&form)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}
	data, err := service.FindCrawlerData(c, form.QueryID)
	if err != nil {
		app.BizFailed(c, e.Error, err)
		return
	}
	app.Success(c, data)
}
