package v1

import (
	"axiangcoding/antonstar/api-system/auth"
	"axiangcoding/antonstar/api-system/entity"
	"axiangcoding/antonstar/api-system/entity/app"
	"axiangcoding/antonstar/api-system/entity/e"
	"axiangcoding/antonstar/api-system/service"
	"github.com/gin-gonic/gin"
)

// LoginForm 目前支持用户名登录
type LoginForm struct {
	// 用户名
	UserName string `binding:"username,required" json:"username" form:"username"`
	// 密码
	Password string `binding:"password,required" json:"password" form:"password"`
}

// UserLogin
// @Summary  用户登录
// @Tags      User API
// @Param    form  body      LoginForm               true  "register form"
// @Param    form  query     middleware.CaptchaForm  true  "captcha"
// @Success  200   {object}  app.ApiJson             ""
// @Failure  400   {object}  app.ErrJson             ""
// @Router   /v1/user/login [post]
func UserLogin(c *gin.Context) {
	form := LoginForm{}
	err := c.ShouldBindJSON(&form)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}

	login := entity.UserLogin{
		UserName: form.UserName,
		Password: form.Password,
	}

	token, err := service.UserLogin(c, login)
	if err != nil {
		app.BizFailed(c, e.LoginFailed, err)
		return
	}
	app.Success(c, map[string]string{app.AuthHeader: token})
}

type RegisterForm struct {
	// 用户名
	UserName string `binding:"username,required" json:"username" form:"username"`
	// 密码
	Password string `binding:"password,required" json:"password" form:"password"`
	// 邮箱
	Email string `binding:"omitempty,email" json:"email" form:"email"`
	// 电话
	Phone string `binding:"omitempty,e164" json:"phone"`
	// 邀请码
	InvitedCode string `binding:"omitempty" json:"invited_code" form:"invited_code"`
	// 头像url地址
	AvatarUrl string `binding:"omitempty,url" json:"avatar_url"`
}

// UserRegister
// @Summary  用户注册
// @Tags      User API
// @Param    form  body      RegisterForm            true  "form"
// @Param    form  query     middleware.CaptchaForm  true  "captcha"
// @Success  200   {object}  app.ApiJson             ""
// @Failure  400   {object}  app.ErrJson             ""
// @Router   /v1/user/register [post]
func UserRegister(c *gin.Context) {
	regForm := RegisterForm{}
	err := c.ShouldBindJSON(&regForm)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}
	register := entity.UserRegister{
		UserName:    regForm.UserName,
		Email:       regForm.Email,
		Phone:       regForm.Phone,
		Password:    regForm.Password,
		InvitedCode: regForm.InvitedCode,
	}
	id, err := service.UserRegister(c, register)
	if err != nil {
		app.BizFailed(c, e.RegisterFailed, err)
		return
	}
	app.Success(c, map[string]string{"id": id})
}

// UserLogout
// @Summary   用户注销
// @Tags     User API
// @Success   200   {object}  app.ApiJson  ""
// @Failure   400   {object}  app.ErrJson  ""
// @Router    /v1/user/logout [post]
// @Security  ApiKeyAuth
func UserLogout(c *gin.Context) {
	result, err := service.UserLogout(c, c.GetHeader(app.AuthHeader))
	if err != nil {
		app.BizFailed(c, e.Error, err)
		return
	}
	if result == 0 {
		app.BizFailed(c, e.TokenNotValid)
		return
	}
	app.Success(c, nil)
}

type KeyFieldExistForm struct {
	// 字段名称，比如username和email
	Key string `json:"key" form:"key" binding:"required,oneof=username email"`
	// 字段值
	Value string `json:"value" form:"value" binding:"required"`
}

// IsKeyFieldValueExist
// @Summary  判断主要的用户信息的值是否存在
// @Tags      User API
// @Param    form  body      KeyFieldExistForm  true  "form"
// @Success   200  {object}  app.ApiJson  ""
// @Failure   400  {object}  app.ErrJson  ""
// @Router   /v1/user/value/exist [post]
func IsKeyFieldValueExist(c *gin.Context) {
	form := KeyFieldExistForm{}
	err := c.ShouldBindJSON(&form)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}
	existed := service.FindValueExist(c, form.Key, form.Value)
	app.Success(c, map[string]interface{}{
		"exists": existed,
	})
}

type IdForm struct {
	// 如果不传入user_id参数，则查token所代表的个人信息，如果传入，则查其他用户的
	UserId int64 `json:"user_id" form:"user_id" binding:"omitempty"`
}

// UserInfo
// @Summary   获取用户信息
// @Tags     User API
// @Param     form  query     IdForm       true  "form"
// @Success   200  {object}  app.ApiJson  ""
// @Failure   400  {object}  app.ErrJson  ""
// @Router    /v1/user/info [post]
// @Security  ApiKeyAuth
func UserInfo(c *gin.Context) {
	form := IdForm{}
	err := c.ShouldBindQuery(&form)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}
	info, err := service.UserInfo(c, c.GetHeader(app.AuthHeader), form.UserId)
	if err != nil {
		app.BizFailed(c, e.Error, err)
		return
	}
	app.Success(c, info)
}

// GetUserWTQueryHistory
// @Summary   获取用户的查询历史记录
// @Tags     User API
// @Success  200   {object}  app.ApiJson        ""
// @Failure  400   {object}  app.ErrJson        ""
// @Router    /v1/user/wt_query/history [get]
// @Security  ApiKeyAuth
func GetUserWTQueryHistory(c *gin.Context) {
	userID := auth.GetUserIdFromToken(c.GetHeader(app.AuthHeader))
	history, err := service.GetWTQueryHistory(c, userID)
	if err != nil {
		app.BizFailed(c, e.Error, err)
		return
	}
	app.Success(c, history)
}
