package service

import (
	"axiangcoding/antonstar/api-system/data"
	"axiangcoding/antonstar/api-system/data/schema"
	"github.com/gin-gonic/gin"
)

func SaveBugReport(c *gin.Context, item schema.BugReport) error {
	err := data.GetDB().Save(&item).Error
	if err != nil {
		return err
	}
	return nil
}
