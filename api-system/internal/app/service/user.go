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
	findUser, err := data.FindUser(ctx, user)
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
	result, err := DeleteCachedToken(c, strconv.FormatInt(claims.UserID, 10))
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
	findUser := data.ExistUser(c, user)
	return findUser
}

func UserInfo(c *gin.Context, token string, id int64) (map[string]interface{}, error) {
	isPrivate := false
	if id == 0 {
		id = auth.GetUserIdFromToken(token)
		isPrivate = true
	}
	user, err := data.FindUser(c, schema.User{UserId: id})

	m := map[string]interface{}{
		"user_id":    strconv.FormatInt(user.UserId, 10),
		"username":   user.UserName,
		"nickname":   user.NickName.String,
		"avatar_url": user.AvatarUrl,
		"roles":      user.Roles,
		"status":     user.Status,
	}
	// 如果是查自己的数据，可以适当添加一些记录
	if isPrivate {
		m["create_at"] = user.CreatedAt
	}
	return m, err
}
