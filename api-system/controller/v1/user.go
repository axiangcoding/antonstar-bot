package v1

import (
	"errors"
	"fmt"
	"github.com/axiangcoding/ax-web/entity/app"
	"github.com/axiangcoding/ax-web/entity/e"
	"github.com/axiangcoding/ax-web/service"
	"github.com/axiangcoding/axth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LoginParam struct {
	LoginName string `json:"loginName" form:"loginName" binding:"required"`
	Password  string `json:"password" form:"password" binding:"required"`
}

// UserLogin
// @Summary
// @Tags     User API
// @Param    param  body      LoginParam   true  "login param"
// @Success  200  {object}  app.ApiJson  ""
// @Router   /v1/user/login [post]
func UserLogin(c *gin.Context) {
	var param LoginParam
	err := c.ShouldBindJSON(&param)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}
	user, err := service.UserLogin(c, param.LoginName, param.Password)
	if err != nil {
		if errors.Is(err, axth.ErrUserPasswordNotMatched) {
			app.BizFailed(c, e.UserPasswordNotMatched)
		} else {
			app.BizFailed(c, e.UserLoginFailed, err)
		}
		return
	}
	session := sessions.Default(c)
	session.Set("userId", user.UserID)
	err = session.Save()
	if err != nil {
		app.BizFailed(c, e.UserLoginFailed, err)
		return
	}

	app.Success(c, map[string]interface{}{
		"user": user,
	})
}

type RegisterParam struct {
	DisplayName string `json:"displayName" form:"displayName" binding:"required"`
	Email       string `json:"email" form:"email" binding:"required,email"`
	Phone       string `json:"phone" form:"phone" binding:"omitempty"`
	Password    string `json:"password" form:"password" binding:"required"`
}

// UserRegister
// @Summary
// @Tags     User API
// @Param    param  body      RegisterParam  true  "register param"
// @Success  200    {object}  app.ApiJson    ""
// @Router   /v1/user/register [post]
func UserRegister(c *gin.Context) {
	var param RegisterParam
	err := c.ShouldBindJSON(&param)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}
	ru := axth.RegisterUser{
		UserID:      uuid.NewString(),
		DisplayName: param.DisplayName,
		Email:       param.Email,
		Phone:       param.Phone,
		Password:    param.Password,
	}
	exist, err := service.CheckUserExist(c, ru)
	if err != nil {
		app.BizFailed(c, e.Error, err)
		return
	}
	if exist {
		app.BizFailed(c, e.UserExist)
		return
	}
	reg, err := service.UserRegister(c, ru)
	if err != nil {
		app.BizFailed(c, e.UserRegisterFailed, err)
		return
	}
	app.Success(c, reg)
}

// UserMe
// @Summary
// @Tags     User API
// @Success  200    {object}  app.ApiJson  ""
// @Router   /v1/user/me [post]
func UserMe(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("userId")
	user, err := service.FindUser(c, fmt.Sprintf("%v", userId))
	if err != nil {
		app.BizFailed(c, e.Error, err)
		return
	}
	app.Success(c, map[string]interface{}{
		"user": user,
	})
}
