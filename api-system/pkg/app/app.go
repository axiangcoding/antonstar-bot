package app

import (
	"axiangcoding/antonstar/api-system/internal/app/conf"
	"axiangcoding/antonstar/api-system/pkg/app/e"
	"axiangcoding/antonstar/api-system/pkg/logging"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type ApiJson struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ErrJson struct {
	Err []string `json:"err"`
}

func generateErrJson(errs []error) *ErrJson {
	if len(errs) == 0 {
		return nil
	}
	hideDetail := conf.Config.App.Response.HideErrorDetails
	var errMessages []string
	if !hideDetail {
		for _, err := range errs {
			errMessages = append(errMessages, err.Error())
		}
	}
	return &ErrJson{Err: errMessages}
}

func generateBadRequestErrJson(err validator.ValidationErrors) *ErrJson {
	var errMessages []string
	for _, v := range err {
		if field, ok := v.(validator.FieldError); ok {
			logging.Info()
			errMessages = append(errMessages,
				fmt.Sprintf("Field validation for '%s' failed on the '%s' tag", field.Field(), field.Tag()))
		}
	}
	return &ErrJson{Err: errMessages}
}

// HttpResponse
// common response
// 返回通用
func HttpResponse(c *gin.Context, httpCode int, msgCode int, data interface{}) {
	c.JSON(httpCode, ApiJson{
		Code: msgCode,
		Msg:  e.CodeText(msgCode),
		Data: data,
	})
}

// Success
// response a success
// 返回成功
func Success(c *gin.Context, data interface{}) {
	HttpResponse(c, http.StatusOK, e.Success, data)
}

// BizFailed
// business failed response
// 返回业务逻辑失败
func BizFailed(c *gin.Context, errCode int, err ...error) {
	HttpResponse(c, http.StatusOK, errCode, generateErrJson(err))
}

// BadRequest
// bad request response
// 返回错误参数请求
func BadRequest(c *gin.Context, errCode int, err error) {
	logging.Warn(err)
	HttpResponse(c, http.StatusBadRequest, errCode, generateBadRequestErrJson(err.(validator.ValidationErrors)))
	c.Abort()
}

// ServerFailed
// server internal failed response
// 返回服务器内部故障
func ServerFailed(c *gin.Context, errCode int, err ...error) {
	HttpResponse(c, http.StatusInternalServerError, errCode, generateErrJson(err))
}

// Unauthorized
// authorized failed response
// 返回权限不足
func Unauthorized(c *gin.Context, errCode int, err ...error) {
	HttpResponse(c, http.StatusUnauthorized, errCode, generateErrJson(err))
	c.Abort()
}

// Forbidden
// authorized forbidden response
// 返回被禁止访问
func Forbidden(c *gin.Context, errCode int, err ...error) {
	HttpResponse(c, http.StatusForbidden, errCode, generateErrJson(err))
	c.Abort()
}
