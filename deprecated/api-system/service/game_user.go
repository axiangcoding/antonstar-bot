package service

import (
	"axiangcoding/antonstar/api-system/data"
	"axiangcoding/antonstar/api-system/data/schema"
	"axiangcoding/antonstar/api-system/entity/app"
	"github.com/gin-gonic/gin"
	"time"
)

func GetGameUsers(c *gin.Context, pagination app.Pagination) (map[string]interface{}, error) {
	var users []schema.GameUser
	var total int64
	start, limit := pagination.ToOffsetLimit()
	sortSql := pagination.GetSortSql()
	filterSql := pagination.GetFilterSql()
	err := data.GetDB().Where(filterSql).Order(sortSql).Offset(start).Limit(limit).Find(&users).
		Offset(-1).Limit(-1).Count(&total).Error
	if err != nil {
		return nil, err
	}
	var displays []map[string]interface{}

	for _, item := range users {
		lastUpdate := time.Now()
		if item.UpdatedAt.IsZero() {
			lastUpdate = item.CreatedAt
		}
		displays = append(displays, map[string]interface{}{
			"nick":       item.Nick,
			"clan":       item.Clan,
			"level":      item.Level,
			"title":      item.Title,
			"ts_ab_rate": item.TsABRate,
			"ts_rb_rate": item.TsRBRate,
			"ts_sb_rate": item.TsSBRate,
			"as_ab_rate": item.AsABRate,
			"as_rb_rate": item.AsRBRate,
			"as_sb_rate": item.AsSBRate,
			"banned":     item.Banned,
			"lastUpdate": lastUpdate,
		})
	}
	return map[string]interface{}{
		"total": total,
		"users": displays,
	}, nil
}
