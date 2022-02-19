package service

import (
	"axiangcoding/antonstar/api-system/internal/app/data"
	"axiangcoding/antonstar/api-system/internal/app/data/schema"
	"axiangcoding/antonstar/api-system/internal/app/entity"
	"axiangcoding/antonstar/api-system/pkg/auth"
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

func UserRegister(ctx *gin.Context, ur entity.UserRegister) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(ur.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	var user = schema.User{
		UserName:    ur.UserName,
		Email:       sql.NullString{String: ur.Email, Valid: ur.Email != ""},
		Phone:       sql.NullString{String: ur.Phone, Valid: ur.Phone != ""},
		Password:    string(hashedPassword),
		InvitedCode: ur.InvitedCode,
		Roles:       schema.UserRoleOrdinary,
		Status:      schema.UserStatusNoVerify,
	}
	user.GenerateId()
	user.NickName = sql.NullString{String: "用户" + strconv.FormatInt(user.UserId, 10)}
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
	return result, err
}

func FindValueExist(c *gin.Context, key string, val string) bool {
	user := schema.User{}
	switch key {
	case "username":
		user.UserName = val
		break
	case "email":
		user.Email = sql.NullString{String: val, Valid: val != ""}
	}
	findUser := data.FindUser(c, user)
	return findUser
}
