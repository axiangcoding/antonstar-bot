package service

import (
	"axiangcoding/antonstar/api-system/internal/app/data"
	"axiangcoding/antonstar/api-system/internal/app/data/schema"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetLastSiteNotice(c *gin.Context) (map[string]interface{}, error) {
	notice, err := data.LastSiteNotice(c, schema.SiteNotice{})
	return map[string]interface{}{
		"create_at":      notice.CreatedAt,
		"title":          notice.Title,
		"editor_user_id": strconv.FormatInt(notice.EditorUserId, 10),
		"content":        notice.Content,
	}, err
}

func SaveSiteNotice(c *gin.Context, notice schema.SiteNotice) (map[string]interface{}, error) {
	id, err := data.SaveSiteNotice(c, notice)
	return map[string]interface{}{
		"id": id,
	}, err
}
