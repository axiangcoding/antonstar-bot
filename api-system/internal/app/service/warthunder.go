package service

import (
	"axiangcoding/antonstar/api-system/internal/app/data"
	"axiangcoding/antonstar/api-system/internal/app/data/schema"
	"axiangcoding/antonstar/api-system/pkg/logging"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func FindCrawlerData(c *gin.Context, queryId string) (map[string]interface{}, error) {
	cd := schema.CrawlerData{
		QueryID: queryId,
	}
	crawlerData, err := data.FindCrawlerData(c, cd)
	if err != nil {
		return nil, err
	}
	m := map[string]interface{}{}
	for _, datum := range crawlerData {
		var itemMap = map[string]interface{}{}
		err := json.Unmarshal([]byte(datum.Content), &itemMap)
		if err != nil {
			logging.Error("Parse crawler data content json error: %s", err)
		}
		m[datum.Source] = itemMap

	}
	return m, nil
}
