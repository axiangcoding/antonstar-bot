package app

import (
	"errors"
	"fmt"
	"github.com/axiangcoding/antonstar-bot/internal/entity/e"
	"github.com/axiangcoding/antonstar-bot/pkg/logging"
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

const (
	validFailedErrMsg = "Valid Error: [%s] should be match tag [%s]"
)

func generateErrJson(errs []error) *ErrJson {
	if len(errs) == 0 {
		return nil
	}
	var errMessages []string
	if true {
		for _, err := range errs {
			var validErrors validator.ValidationErrors
			if errors.As(err, &validErrors) {
				for _, err := range err.(validator.ValidationErrors) {
					errMessages = append(errMessages, fmt.Sprintf(validFailedErrMsg, err.Field(), err.Tag()))
				}
			} else {
				errMessages = append(errMessages, err.Error())
			}
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
	logging.L().Warn("Biz failed",
		logging.Any("errCode", errCode),
		logging.Errors("errors", err))
	HttpResponse(c, http.StatusOK, errCode, generateErrJson(err))
	c.Abort()
}

// BadRequest
// bad request response
// 返回错误参数请求
func BadRequest(c *gin.Context, errCode int, err ...error) {
	logging.L().Warn("Bad request",
		logging.Any("errCode", errCode),
		logging.Any("errors", err))
	HttpResponse(c, http.StatusBadRequest, errCode, generateErrJson(err))
	c.Abort()
}

// ServerFailed
// server internal failed response
// 返回服务器内部故障
func ServerFailed(c *gin.Context, errCode int, err ...error) {
	logging.L().Error("Server failed",
		logging.Any("errCode", errCode),
		logging.Any("errors", err))
	HttpResponse(c, http.StatusInternalServerError, errCode, generateErrJson(err))
	c.Abort()
}

// Unauthorized
// authorized failed response
// 返回权限不足
func Unauthorized(c *gin.Context, errCode int, err ...error) {
	logging.L().Warn("Unauthorized",
		logging.Any("errCode", errCode),
		logging.Any("errors", err))
	HttpResponse(c, http.StatusUnauthorized, errCode, generateErrJson(err))
	c.Abort()
}

// Forbidden
// authorized forbidden response
// 返回被禁止访问
func Forbidden(c *gin.Context, errCode int, err ...error) {
	logging.L().Warn("Forbidden",
		logging.Any("errCode", errCode),
		logging.Any("errors", err))
	HttpResponse(c, http.StatusForbidden, errCode, generateErrJson(err))
	c.Abort()
}
