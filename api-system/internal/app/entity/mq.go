package entity

const (
	SourceGaijin       = "gaijin"
	SourceThunderskill = "thunderskill"
)

type MQBody struct {
	QueryID  string `json:"query_id"`
	Source   string `json:"source"`
	Nickname string `json:"nickname"`
}
