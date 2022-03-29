package data

import (
	"axiangcoding/antonstar/api-system/data/schema"
	"axiangcoding/antonstar/api-system/logging"
	"context"
	"time"
)

func QueryShortCrawlerData(c context.Context, crawlerData schema.CrawlerData) ([]schema.ShortCrawlerData, error) {
	var records []schema.ShortCrawlerData
	err := GetDB().Model(&schema.CrawlerData{}).
		Where(&crawlerData).Order("updated_at desc").
		Order("created_at desc").Limit(7).Find(&records).Error
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

func CountCrawlerQuery(c context.Context, timestamp time.Time) int64 {
	var count int64
	model := GetDB().Model(&schema.CrawlerData{})
	if !timestamp.IsZero() {
		model.Where("to_days(created_at) = to_days(?)", timestamp)
	}
	err := model.Count(&count).Error
	if err != nil {
		logging.Error(err)
	}
	return count
}
