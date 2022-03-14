package middleware

import (
	auth2 "axiangcoding/antonstar/api-system/auth"
	"axiangcoding/antonstar/api-system/entity/app"
	"axiangcoding/antonstar/api-system/entity/e"
	"axiangcoding/antonstar/api-system/service"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"strconv"
	"strings"
)

// AuthCheck 用户权限校验
func AuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader(app.AuthHeader)
		claims, next := checkToken(c, tokenString)
		if !next {
			return
		}
		next = checkTokenInCache(c, tokenString, claims.UserID)
		if !next {
			return
		}
		next = checkPermission(c, claims.Roles)
		if !next {
			return
		}
		c.Next()
	}
}

type CaptchaForm struct {
	CaptchaId  string `binding:"required" json:"captcha_id" form:"captcha_id"`
	CaptchaVal string `binding:"required" json:"captcha_val" form:"captcha_val"`
}

// CaptchaCheck 验证码校验
func CaptchaCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		var form CaptchaForm
		err := c.ShouldBindQuery(&form)
		if err != nil {
			app.BizFailed(c, e.CaptchaNotValid, err)
			c.Abort()
			return
		}
		if captcha.VerifyString(form.CaptchaId, form.CaptchaVal) {
			c.Next()
		} else {
			app.BizFailed(c, e.CaptchaNotValid)
			c.Abort()
		}
	}
}

// checkToken
func checkToken(c *gin.Context, tokenString string) (*auth2.CustomClaims, bool) {
	if tokenString == "" {
		app.Unauthorized(c, e.TokenNotExist)
		return nil, false
	}
	claims, err := auth2.ParseToken(tokenString)
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				// That's not even a token
				// 这不是个token
				app.Unauthorized(c, e.TokenNotValid)
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// token is either expired or not active yet
				// token要么过期了，要么还没生效
				app.Unauthorized(c, e.TokenExpired)
			} else {
				app.Unauthorized(c, e.TokenNotValid)
			}
		}
		return nil, false
	}
	return claims, true
}

// checkTokenInCache
// check token in cache
// 检查缓存中的token
func checkTokenInCache(c *gin.Context, tokenString string, userID int64) bool {
	// check token in cache
	// 检查缓存中的token
	cacheToken, err := service.GetCachedToken(c, strconv.FormatInt(userID, 10))
	if err != nil {
		app.Unauthorized(c, e.TokenExpired, err)
		return false
	}
	if tokenString == cacheToken {
		service.RefreshTokenTTL(c, strconv.FormatInt(userID, 10))
	} else {
		app.Unauthorized(c, e.TokenExpired)
		return false
	}
	return true
}

// checkPermission
// check user permission to access resources
// 检查用户访问资源的权限
func checkPermission(c *gin.Context, roles string) bool {
	roleItems := strings.Split(roles, ",")
	if len(roleItems) == 0 {
		app.Forbidden(c, e.NoPermission)
		return false
	}
	hasPermission := false
	for _, role := range roleItems {
		allowed, err := auth2.GetEnforcer().Enforce(role, c.Request.URL.Path, c.Request.Method)
		if err != nil {
			app.Forbidden(c, e.NoPermission, err)
			return false
		}
		hasPermission = hasPermission || allowed
		if hasPermission {
			break
		}
	}
	if !hasPermission {
		app.Forbidden(c, e.NoPermission)
		return false
	}
	return true
}
