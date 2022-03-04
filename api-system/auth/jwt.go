package auth

import (
	"axiangcoding/antonstar/api-system/data/schema"
	"axiangcoding/antonstar/api-system/logging"
	"axiangcoding/antonstar/api-system/settings"
	"github.com/google/uuid"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret []byte
var expireDuration time.Duration

type UserInfo struct {
	UserID int64  `json:"user_id"`
	Roles  string `json:"roles"`
}

type CustomClaims struct {
	*jwt.StandardClaims
	UserInfo
}

func SetupJwt() {
	jwtSecret = []byte(settings.Config.App.Auth.Secret)
	expireStr := settings.Config.App.Auth.ExpireDuration
	expire, err := time.ParseDuration(expireStr)
	if err != nil {
		logging.Fatal("Config properties: app.token.expire_duration not valid")
	}
	expireDuration = expire
}

// CreateToken generate tokens used for auth
func CreateToken(user schema.User) (string, error) {
	t := jwt.New(jwt.SigningMethodHS256)
	now := time.Now()
	t.Claims = &CustomClaims{
		&jwt.StandardClaims{
			Id:        uuid.NewString(),
			ExpiresAt: now.Add(expireDuration).Unix(),
			IssuedAt:  now.Unix(),
			NotBefore: now.Unix(),
			Issuer:    settings.Config.App.Name,
		},
		UserInfo{
			UserID: user.UserId,
			Roles:  user.Roles,
		},
	}
	return t.SignedString(jwtSecret)
}

// ParseToken parsing token
func ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
	if token != nil {
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// GetUserIdFromToken
// 从Token中获取用户ID
func GetUserIdFromToken(token string) int64 {
	claims, err := ParseToken(token)
	if err != nil {
		return 0
	}
	return claims.UserID
}
