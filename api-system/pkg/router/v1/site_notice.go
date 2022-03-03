package v1

import (
	"axiangcoding/antonstar/api-system/internal/app/data/schema"
	"axiangcoding/antonstar/api-system/internal/app/service"
	"axiangcoding/antonstar/api-system/pkg/app"
	"axiangcoding/antonstar/api-system/pkg/app/e"
	"axiangcoding/antonstar/api-system/pkg/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetLastSiteNotice
// @Summary  获取最新的一条全站公告消息
// @Tags      Site
// @Success   200   {object}  app.ApiJson  ""
// @Router   /v1/site/notice/last [get]
func GetLastSiteNotice(c *gin.Context) {
	notice, err := service.GetLastSiteNotice(c)
	if err != nil {
		app.BizFailed(c, e.Error, err)
		return
	}
	app.Success(c, notice)
}

type NoticeForm struct {
	Content string `json:"content" binding:"required,max=2000" form:"content"`
	Title   string `json:"title" binding:"omitempty,max=30" form:"title"`
}

// PostSiteNotice
// @Summary   新增一条全站公告消息
// @Tags     Site
// @Param     form  body      NoticeForm   true  "form"
// @Success  200  {object}  app.ApiJson  ""
// @Router    /v1/site/notice/ [post]
// @Security  ApiKeyAuth
func PostSiteNotice(c *gin.Context) {
	var form NoticeForm
	err := c.ShouldBindJSON(&form)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}
	userID := auth.GetUserIdFromToken(c.GetHeader(app.AuthHeader))
	info, err := service.SaveSiteNotice(c, schema.SiteNotice{
		Model:        gorm.Model{},
		Title:        form.Title,
		Content:      form.Content,
		EditorUserId: userID,
	})
	if err != nil {
		app.BizFailed(c, e.Error, err)
		return
	}
	app.Success(c, info)
}
