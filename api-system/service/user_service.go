package service

import (
	"github.com/axiangcoding/ax-web/auth"
	"github.com/axiangcoding/axth"
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context, loginName string, password string) (*axth.DisplayUser, error) {
	enforcer := auth.GetAxthEnforcer()
	user, err := enforcer.LoginWithEmail(loginName, password)
	return user, err
}

func CheckUserExist(c *gin.Context, user axth.RegisterUser) (bool, error) {
	enforcer := auth.GetAxthEnforcer()
	exist1, err := enforcer.CheckEmailExist(user.Email)
	exist2, err := enforcer.CheckUserIdExist(user.UserID)
	exist3, err := enforcer.CheckPhoneExist(user.Phone)
	if err != nil {
		return false, err
	}
	return exist1 || exist2 || exist3, nil
}

func UserRegister(c *gin.Context, user axth.RegisterUser) (bool, error) {
	enforcer := auth.GetAxthEnforcer()
	register, err := enforcer.Register(user)
	if err != nil {
		return false, err
	}
	return register, nil
}

func FindUser(c *gin.Context, userId string) (*axth.DisplayUser, error) {
	enforcer := auth.GetAxthEnforcer()
	user, err := enforcer.FindUser(userId)
	return user, err
}
