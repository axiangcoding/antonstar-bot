package entity

const (
	SourceGaijin       = "gaijin"
	SourceThunderskill = "thunder_skill"
)

type MQBody struct {
	// 查询的ID
	QueryID string `json:"query_id"`
	// 调用的目标爬虫
	Target []string `json:"target"`
	// 查询玩家的昵称
	Nickname string `json:"nickname"`
}
