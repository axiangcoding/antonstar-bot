package service

import (
	"axiangcoding/antonstar/api-system/data"
	"axiangcoding/antonstar/api-system/data/schema"
	"axiangcoding/antonstar/api-system/logging"
	"axiangcoding/antonstar/api-system/mq"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func GetAllUserInfo(c *gin.Context, nick string, userID int64) (map[string][]interface{}, error) {
	lst := make(map[string][]interface{})
	crawlerData, err := data.QueryShortCrawlerData(c, schema.CrawlerData{Nick: nick, Source: schema.SourceGaijin})
	if err != nil {
		return lst, err
	}
	for _, datum := range crawlerData {
		source := datum.Source
		lst[source] = append(lst[source], map[string]interface{}{
			"query_id":   datum.QueryID,
			"updated_at": datum.UpdatedAt,
			"created_at": datum.CreatedAt,
			"found":      datum.Found,
			"status":     datum.Status,
			"source":     source,
		})
	}
	// 如果是登录用户搜索，则记录搜索记录
	if userID != 0 {
		data.GetDB().Save(&schema.SearchHistory{
			Type:    schema.SearchHistoryTypeCrawlerQuery,
			UserId:  userID,
			Context: nick,
			Found:   len(lst) != 0,
		})
	}
	return lst, nil
}

func CheckReachRefreshLimit(c *gin.Context) error {
	limit := GetRefreshLimit(c)
	size := data.CountCrawlerQuery(c, time.Now())
	if (size / 2) >= int64(limit) {
		return errors.New("reach daily refresh limit")
	}
	return nil
}

// RefreshUserInfo
// 1. 查看该用户的请求是否在1天以内
// 2. 如果在一天以内，不做请求，返回最近的一个queryId
// 3. 如果在一天以外，做新的请求，返回新的queryId
func RefreshUserInfo(c *gin.Context, nick string, triggerUser int64) (map[string]interface{}, error) {
	find, err := data.FindLastCrawlerData(c, schema.CrawlerData{Nick: nick})
	if err != nil {
		// 如果未找到记录，做一次查询
		if errors.Is(err, gorm.ErrRecordNotFound) {
			crawler, err := sendNewRequestToCrawler(c, nick, triggerUser)
			return map[string]interface{}{
				"refresh":    true,
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
				"refresh":    false,
				"query_id":   find.QueryID,
				"created_at": createdAt,
			}, nil
		} else {
			crawler, err := sendNewRequestToCrawler(c, nick, triggerUser)
			return map[string]interface{}{
				"refresh":    true,
				"query_id":   crawler.QueryID,
				"created_at": crawler.CreatedAt,
			}, err
		}
	}
}

func sendNewRequestToCrawler(c *gin.Context, nick string, triggerUser int64) (*schema.CrawlerData, error) {
	queryID := uuid.NewString()
	// 先保存请求信息
	saved, err := data.SaveCrawlerData(c, schema.CrawlerData{
		Nick:          nick,
		QueryID:       queryID,
		Source:        schema.SourceGaijin,
		Status:        schema.CrawlerStatusRunning,
		TriggerUserId: triggerUser,
	})
	saved, err = data.SaveCrawlerData(c, schema.CrawlerData{
		Nick:          nick,
		QueryID:       queryID,
		Source:        schema.SourceThunderskill,
		Status:        schema.CrawlerStatusRunning,
		TriggerUserId: triggerUser,
	})
	if err != nil {
		return nil, err
	}
	body := mq.CrawBody{
		QueryID:  queryID,
		Target:   []string{schema.SourceGaijin, schema.SourceThunderskill},
		Nickname: nick,
	}
	err = SendMessage(body)
	if err != nil {
		return nil, err
	}
	return &saved, nil
}

func FindCrawlerData(c *gin.Context, queryId string) (map[string]interface{}, error) {
	crawlerData, err := data.QueryCrawlerData(c, schema.CrawlerData{QueryID: queryId})
	if err != nil {
		return nil, err
	}
	m := map[string]interface{}{}
	for _, datum := range crawlerData {
		var itemMap = map[string]interface{}{}
		itemMap["found"] = datum.Found
		itemMap["created_at"] = datum.CreatedAt
		itemMap["updated_at"] = datum.UpdatedAt
		itemMap["status"] = datum.Status
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

func CountCrawlerQuery(c *gin.Context, timestamp time.Time) map[string]interface{} {
	visit := data.CountCrawlerQuery(c, timestamp)
	return map[string]interface{}{
		"count": visit,
	}
}
