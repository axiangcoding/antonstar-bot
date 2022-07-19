package data

import (
	"axiangcoding/antonstar/api-system/data/schema"
	"context"
)

func SaveSiteNotice(c context.Context, s schema.SiteNotice) (uint, error) {
	err := GetDB().Save(&s).Error
	return s.ID, err
}

func TakeSiteNotice(c context.Context, query schema.SiteNotice) (schema.SiteNotice, error) {
	var findItem schema.SiteNotice
	find := GetDB().Where(&query).Take(&findItem)
	return findItem, find.Error
}

func LastSiteNotice(c context.Context, query schema.SiteNotice) (schema.SiteNotice, error) {
	var findItem schema.SiteNotice
	find := GetDB().Where(&query).Last(&findItem)
	return findItem, find.Error
}
