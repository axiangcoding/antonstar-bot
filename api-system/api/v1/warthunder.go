package v1

import (
	"axiangcoding/antonstar/api-system/internal/app/entity"
	"axiangcoding/antonstar/api-system/internal/app/service"
	"axiangcoding/antonstar/api-system/pkg/app"
	"axiangcoding/antonstar/api-system/pkg/app/e"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserInfoForm struct {
	Nickname string `form:"nickname" binding:"required"`
}

// UserInfoQuery
// @Summary
// @Tags     WarThunder
// @Param    form  query     UserInfoForm  true  "query userinfo"
// @Success  200   {object}  app.ApiJson         ""
// @Router   /v1/war_thunder/userinfo/query [get]
func UserInfoQuery(c *gin.Context) {
	var form UserInfoForm
	err := c.ShouldBindQuery(&form)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}
	queryID := uuid.NewString()
	body1 := entity.MQBody{
		QueryID:  queryID,
		Source:   entity.SourceGaijin,
		Nickname: form.Nickname,
	}
	body2 := entity.MQBody{
		QueryID:  queryID,
		Source:   entity.SourceThunderskill,
		Nickname: form.Nickname,
	}
	err = service.SendMessage(body1)
	if err != nil {
		app.BizFailed(c, e.Error, err)
		return
	}
	err = service.SendMessage(body2)
	if err != nil {
		app.BizFailed(c, e.Error, err)
		return
	}
	app.Success(c, map[string]string{
		"send":     "success",
		"query_id": queryID},
	)
}

type UserInfoDetailForm struct {
	QueryID string `json:"query_id" form:"query_id" binding:"required"`
}

// UserInfoDetail
// @Summary
// @Tags     WarThunder
// @Param    form  query     UserInfoDetailForm  true  "query userinfo"
// @Success  200   {object}  app.ApiJson   ""
// @Router   /v1/war_thunder/userinfo/detail [get]
func UserInfoDetail(c *gin.Context) {
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
