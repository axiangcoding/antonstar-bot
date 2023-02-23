package v1

import (
	"github.com/axiangcoding/antonstar-bot/internal/entity/app"
	"github.com/axiangcoding/antonstar-bot/internal/entity/e"
	"github.com/axiangcoding/antonstar-bot/internal/service"
	"github.com/gin-gonic/gin"
)

// GetMission
// @Summary  获取执行任务状态
// @Tags     Mission API
// @Param    id   query     string       true  "mission id"
// @Success  200  {object}  app.ApiJson  ""
// @Router   /v1/mission [get]
func GetMission(c *gin.Context) {
	id := c.Query("id")
	mission, err := service.FindMission(id)
	if err != nil {
		app.BizFailed(c, e.Error, err)
		return
	}
	app.Success(c, map[string]any{
		"mission_id":    mission.MissionId,
		"created_at":    mission.CreatedAt,
		"updated_at":    mission.UpdatedAt,
		"status":        mission.Status,
		"process":       mission.Process,
		"finished_time": mission.FinishedTime,
	})

}
