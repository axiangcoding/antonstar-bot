package v1

import (
	"github.com/axiangcoding/ax-web/entity/app"
	"github.com/axiangcoding/ax-web/entity/e"
	"github.com/gin-gonic/gin"
)

type CommonParam struct {
	// param1, min 10 words and max 255 words
	Param1 string `json:"param1" form:"param1" binding:"min=10,max=255"`
	// param2, required
	Param2 string `json:"param2" form:"param2" binding:"required"`
	// param3, if it's null, validate nothing. if it's not null, must match email regex
	Param3 string `json:"param3" form:"param3" binding:"omitempty,email"`
}

// DemoGet
// @Summary  Demo for Get
// @Tags     Demo API
// @Param    param  query     CommonParam  true  "getParam"
// @Success  200    {object}  app.ApiJson  ""
// @Router   /v1/demo/get [get]
func DemoGet(c *gin.Context) {
	var param CommonParam
	err := c.ShouldBindQuery(&param)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}
	app.Success(c, param)
}

// DemoPost
// @Summary  Demo for Post
// @Tags     Demo API
// @Param    param  body      CommonParam  true  "getParam"
// @Success  200    {object}  app.ApiJson  ""
// @Router   /v1/demo/post [post]
func DemoPost(c *gin.Context) {
	var param CommonParam
	err := c.ShouldBindJSON(&param)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}
	app.Success(c, param)
}
