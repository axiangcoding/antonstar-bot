package service

import (
	"axiangcoding/antonstar/api-system/internal/app/data"
	"axiangcoding/antonstar/api-system/internal/app/data/schema"
	"axiangcoding/antonstar/api-system/internal/app/entity"
	"axiangcoding/antonstar/api-system/pkg/auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

func UserRegister(ctx *gin.Context, ur entity.UserRegister) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ur.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user := schema.User{
		UserName: ur.UserName,
		Email:    ur.Email,
		Phone:    ur.Phone,
		Password: string(hashedPassword),
		Roles:    schema.UserRoleOrdinary,
	}
	user.GenerateId()
	return data.UserRegister(ctx, user)
}

func UserLogin(ctx *gin.Context, login entity.UserLogin) (string, error) {
	user := schema.User{
		UserName: login.UserName,
	}
	findUser, err := data.UserLogin(ctx, user)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(findUser.Password), []byte(login.Password))
	if err != nil {
		return "", err
	}
	token, err := auth.CreateToken(findUser)
	if err != nil {
		return "", err
	}
	err = CacheToken(ctx, strconv.FormatInt(findUser.UserId, 10), token)
	if err != nil {
		return "", err
	}
	return token, nil
}

func UserLogout(c *gin.Context, token string) (int64, error) {

	claims, _ := auth.ParseToken(token)
	result, err := DeleteCachedToken(c, claims.Id)
	if err != nil {
		return 0, err
	}
	return result, nil
}
