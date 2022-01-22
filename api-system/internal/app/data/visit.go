package data

import (
	"axiangcoding/antonstar/api-system/internal/app/data/schema"
	"context"
)

func SaveVisit(ctx context.Context, visit schema.Visit) error {
	err := GetDB().Save(&visit).Error
	return err
}
