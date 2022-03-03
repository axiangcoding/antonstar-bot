package data

import (
	"axiangcoding/antonstar/api-system/internal/app/data/schema"
	"context"
	"errors"
	"gorm.io/gorm"
	"strconv"
)

func UserRegister(ctx context.Context, user schema.User) (string, error) {
	err := GetDB().Save(&user).Error
	if err != nil {
		return "", err
	}
	id := strconv.FormatInt(user.UserId, 10)
	return id, err
}

func FindUser(ctx context.Context, user schema.User) (schema.User, error) {
	var found schema.User
	take := GetDB().Where(user).Take(&found)
	if errors.Is(take.Error, gorm.ErrRecordNotFound) {
		return found, take.Error
	}
	return found, nil
}

// ExistUser
// 检查用户 是否存在
func ExistUser(c context.Context, user schema.User) bool {
	take := GetDB().Where(user).Take(&schema.User{})
	return take.RowsAffected == 1
}
