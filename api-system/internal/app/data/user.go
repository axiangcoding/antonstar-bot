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

func UserLogin(ctx context.Context, user schema.User) (schema.User, error) {
	var findUser schema.User
	find := GetDB().Where(user).Take(&findUser)
	if errors.Is(find.Error, gorm.ErrRecordNotFound) {
		return findUser, find.Error
	}
	return findUser, nil
}
