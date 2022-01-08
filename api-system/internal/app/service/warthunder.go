package service

import (
	"axiangcoding/antonstar/api-system/internal/app/data"
	"axiangcoding/antonstar/api-system/internal/app/data/schema"
	"github.com/gin-gonic/gin"
)

func FindCrawlerData(c *gin.Context, queryId string) (map[string]string, error) {
	cd := schema.CrawlerData{
		QueryID: queryId,
	}
	crawlerData, err := data.FindCrawlerData(c, cd)
	if err != nil {
		return nil, err
	}
	m := map[string]string{}
	for _, datum := range crawlerData {
		m[datum.Source] = datum.Content
	}
	return m, nil
}
