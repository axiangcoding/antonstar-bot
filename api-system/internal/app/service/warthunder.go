package service

import (
	"axiangcoding/antonstar/api-system/internal/app/data"
	"axiangcoding/antonstar/api-system/internal/app/data/schema"
	"axiangcoding/antonstar/api-system/internal/app/entity"
	"axiangcoding/antonstar/api-system/pkg/logging"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// RequestUserInfo
// 1. 查看该用户的请求是否在1天以内
// 2. 如果在一天以内，不做请求，返回最近的一个queryId
// 3. 如果在一天以外，做新的请求，返回新的queryId
func RequestUserInfo(c *gin.Context, nick string) (map[string]interface{}, error) {
	find, err := data.FindLastCrawlerData(c, schema.CrawlerData{Nick: nick})
	if err != nil {
		// 如果未找到记录，做一次查询
		if errors.Is(err, gorm.ErrRecordNotFound) {
			crawler, err := sendNewRequestToCrawler(c, nick)
			return map[string]interface{}{
				"send":       true,
				"query_id":   crawler.QueryID,
				"created_at": crawler.CreatedAt,
			}, err
		} else {
			return nil, err
		}
	} else {
		now := time.Now()
		createdAt := find.CreatedAt
		// 如果两个针对一个nick的请求小于24小时，则不重新发送请求
		if now.Sub(createdAt).Hours() <= 24 {
			return map[string]interface{}{
				"send":       false,
				"query_id":   find.QueryID,
				"created_at": createdAt,
			}, nil
		} else {
			crawler, err := sendNewRequestToCrawler(c, nick)
			return map[string]interface{}{
				"send":       true,
				"query_id":   crawler.QueryID,
				"created_at": crawler.CreatedAt,
			}, err
		}
	}
}

func sendNewRequestToCrawler(c *gin.Context, nick string) (*schema.CrawlerData, error) {
	queryID := uuid.NewString()
	// 先保存请求信息
	saved, err := data.SaveCrawlerData(c, schema.CrawlerData{
		Nick:    nick,
		QueryID: queryID,
		Source:  entity.SourceGaijin,
	})
	saved, err = data.SaveCrawlerData(c, schema.CrawlerData{
		Nick:    nick,
		QueryID: queryID,
		Source:  entity.SourceThunderskill,
	})
	if err != nil {
		return nil, err
	}
	body1 := entity.MQBody{
		QueryID:  queryID,
		Source:   entity.SourceGaijin,
		Nickname: nick,
	}
	body2 := entity.MQBody{
		QueryID:  queryID,
		Source:   entity.SourceThunderskill,
		Nickname: nick,
	}
	err = SendMessages(body1, body2)
	if err != nil {
		return nil, err
	}
	return &saved, nil
}

func FindCrawlerData(c *gin.Context, queryId string) (map[string]interface{}, error) {
	cd := schema.CrawlerData{
		QueryID: queryId,
	}
	crawlerData, err := data.QueryCrawlerData(c, cd)
	if err != nil {
		return nil, err
	}
	m := map[string]interface{}{}
	for _, datum := range crawlerData {
		var itemMap = map[string]interface{}{}
		itemMap["found"] = datum.Found
		itemMap["updated_at"] = datum.UpdatedAt
		if datum.Found {
			err := json.Unmarshal([]byte(datum.Content), &itemMap)
			if err != nil {
				logging.Errorf("Parse crawler data content json error: %s", err)
			}
		}
		m[datum.Source] = itemMap

	}
	return m, nil
}
