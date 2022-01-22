package data

import (
	"axiangcoding/antonstar/api-system/internal/app/data/schema"
	"context"
)

func QueryShortCrawlerData(c context.Context, crawlerData schema.CrawlerData) ([]schema.ShortCrawlerData, error) {
	var records []schema.ShortCrawlerData
	err := GetDB().Model(&schema.CrawlerData{}).Where(&crawlerData).Order("updated_at desc").Order("created_at desc").Find(&records).Error
	return records, err
}

func QueryCrawlerData(c context.Context, crawlerData schema.CrawlerData) ([]schema.CrawlerData, error) {
	var records []schema.CrawlerData
	err := GetDB().Where(&crawlerData).Find(&records).Error
	return records, err
}

func FindLastCrawlerData(c context.Context, crawlerData schema.CrawlerData) (schema.CrawlerData, error) {
	var record schema.CrawlerData
	err := GetDB().Where(&crawlerData).Last(&record).Error
	return record, err
}

func SaveCrawlerData(c context.Context, crawlerData schema.CrawlerData) (schema.CrawlerData, error) {
	err := GetDB().Save(&crawlerData).Error
	return crawlerData, err
}
