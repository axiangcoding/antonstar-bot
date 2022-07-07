package v1

import (
	"axiangcoding/antonstar/api-system/auth"
	"axiangcoding/antonstar/api-system/data/schema"
	"axiangcoding/antonstar/api-system/entity/app"
	"axiangcoding/antonstar/api-system/entity/e"
	"axiangcoding/antonstar/api-system/service"
	"github.com/gin-gonic/gin"
)

type BugReportForm struct {
	// 报告的类型
	Type string `json:"type" binding:"required" form:"type"`
	// 报告的标题
	Title string `json:"title" binding:"required,max=30" form:"title"`
	// 报告的详情
	Content string `json:"content" binding:"omitempty,max=2000" form:"content"`
	// 是否匿名
	Anonymous bool `json:"anonymous" binding:"omitempty" form:"anonymous"`
}

// PostBugReport
// @Summary  提交一条问题反馈
// @Tags     BugReport API
// @Param    form  body      BugReportForm  true  "form"
// @Success  200   {object}  app.ApiJson    ""
// @Router   /v1/bug_report/ [post]
func PostBugReport(c *gin.Context) {
	var form BugReportForm
	err := c.ShouldBindJSON(&form)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}
	userID := auth.GetUserIdFromToken(c.GetHeader(app.AuthHeader))
	if form.Anonymous {
		userID = 0
	}
	err = service.SaveBugReport(c, schema.BugReport{
		Type:    form.Type,
		Title:   form.Title,
		Content: form.Content,
		UserId:  userID,
	})
	if err != nil {
		app.BizFailed(c, e.Error, err)
		return
	}
	app.Success(c, nil)
}
