package v1

import (
	"axiangcoding/antonstar/api-system/internal/app/entity"
	"axiangcoding/antonstar/api-system/internal/app/service"
	"axiangcoding/antonstar/api-system/pkg/app"
	"axiangcoding/antonstar/api-system/pkg/app/e"
	"axiangcoding/antonstar/api-system/pkg/logging"
	"github.com/gin-gonic/gin"
	"time"
)

type PostVisitForm struct {
	// 客户端生成id
	ClientId string `form:"client_id" json:"client_id"`
	// 用户id
	UserId int64 `form:"user_id" json:"user_id"`
	// 访问页面
	Page string `form:"page" json:"page"`
}

// PostVisit
// @Summary  登记访问信息
// @Tags      Visit
// @Param    form  body      PostVisitForm  true  "query userinfo"
// @Success  200   {object}  app.ApiJson    ""
// @Router   /v1/visits/visit [post]
func PostVisit(c *gin.Context) {
	var form PostVisitForm
	err := c.ShouldBindJSON(&form)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}
	visit := entity.AddVisit{
		ClientId:  form.ClientId,
		UserId:    form.UserId,
		Page:      form.Page,
		VisitTime: time.Now(),
	}
	// 异步登记访问信息
	ccp := c.Copy()
	go func() {
		err = service.AddVisit(ccp, visit)
		if err != nil {
			logging.Errorf("save visit error: %s", err)
		}
	}()
	app.Success(c, nil)
}

// GetVisits
// @Summary
// @Tags     Visit
// @Param     form  query     string       true  "query userinfo"
// @Success   200   {object}  app.ApiJson  ""
// @Router    /v1/visits [get]
// @Security  ApiKeyAuth
func GetVisits(c *gin.Context) {

}

type GetVisitsCountForm struct {
	Timestamp time.Time `form:"timestamp" json:"timestamp"`
}

// GetVisitCount
// @Summary
// @Tags     Visit
// @Param    form  query     GetVisitsCountForm  true  "query userinfo"
// @Success  200   {object}  app.ApiJson         ""
// @Router   /v1/visits/count [get]
func GetVisitCount(c *gin.Context) {
	var form GetVisitsCountForm
	err := c.ShouldBindQuery(&form)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}
	info := service.CountVisit(c, form.Timestamp)
	app.Success(c, info)
}
