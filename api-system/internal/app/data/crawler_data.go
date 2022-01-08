package data

import (
	"axiangcoding/antonstar/api-system/internal/app/data/schema"
	"context"
)

func FindCrawlerData(ctx context.Context, crawlerData schema.CrawlerData) ([]schema.CrawlerData, error) {
	var records []schema.CrawlerData
	err := GetDB().Where(&crawlerData).Find(&records).Error
	return records, err
}
