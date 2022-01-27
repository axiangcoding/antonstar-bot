package v1

import (
	"axiangcoding/antonstar/api-system/internal/app/service"
	"axiangcoding/antonstar/api-system/pkg/app"
	"axiangcoding/antonstar/api-system/pkg/app/e"
	"github.com/gin-gonic/gin"
	"time"
)

type UserInfoForm struct {
	// 游戏的昵称
	Nickname string `form:"nickname" binding:"required,max=20"`
}

// GetUserInfoQueries
// @Summary  查询游戏昵称的所有query_id
// @Tags     WarThunder
// @Param    form  query     UserInfoForm  true  "param"
// @Success  200   {object}  app.ApiJson        ""
// @Router   /v1/war_thunder/userinfo/queries [get]
func GetUserInfoQueries(c *gin.Context) {
	var form UserInfoForm
	err := c.ShouldBindQuery(&form)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}
	info, err := service.GetAllUserInfo(c, form.Nickname)
	if err != nil {
		app.BizFailed(c, e.Error, err)
		return
	}
	app.Success(c, info)
}

// PostUserInfoRefresh
// @Summary  刷新一个游戏用户数据的最新数据
// @Tags     WarThunder
// @Param    form  query     UserInfoForm  true  "param"
// @Success  200   {object}  app.ApiJson   ""
// @Router   /v1/war_thunder/userinfo/refresh [post]
func PostUserInfoRefresh(c *gin.Context) {
	var form UserInfoForm
	err := c.ShouldBindQuery(&form)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}
	err = service.CheckReachRefreshLimit(c)
	if err != nil {
		app.BizFailed(c, e.ReachRefreshLimit)
		return
	}
	info, err := service.RefreshUserInfo(c, form.Nickname)
	if err != nil {
		app.BizFailed(c, e.Error, err)
		return
	}
	app.Success(c, info)
}

type UserInfoDetailForm struct {
	QueryID string `json:"query_id" form:"query_id" binding:"required"`
}

// GetUserInfo
// @Summary  获取异步查询结果
// @Tags     WarThunder
// @Param    form  query     UserInfoDetailForm  true  "query userinfo"
// @Success  200   {object}  app.ApiJson         ""
// @Router   /v1/war_thunder/userinfo [get]
func GetUserInfo(c *gin.Context) {
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

type GetQueryCountForm struct {
	Timestamp time.Time `form:"timestamp" json:"timestamp"`
}

// GetQueryCount
// @Summary  查询query的数量
// @Tags     WarThunder
// @Param    form  query     GetQueryCountForm  true  "param"
// @Success  200   {object}  app.ApiJson   ""
// @Router   /v1/war_thunder/userinfo/query/count [get]
func GetQueryCount(c *gin.Context) {
	var form GetQueryCountForm
	err := c.ShouldBindQuery(&form)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}
	info := service.CountCrawlerQuery(c, form.Timestamp)
	app.Success(c, info)
}
