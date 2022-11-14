package bot

import "regexp"

var (
	// MessageGetAtPrimaryMsgPattern at时获取主要消息
	MessageGetAtPrimaryMsgPattern  = regexp.MustCompile(`^.*\[.*\](.*)$`)
	MessageGetCmdPrimaryMsgPattern = regexp.MustCompile(`^\s*\.cqbot\s*(.*)$`)
)

var (
	ActionUnknown     = "unknown"
	ActionQuery       = "query"
	ActionFullQuery   = "fullQuery"
	ActionRefresh     = "refresh"
	ActionReport      = "report"
	ActionDrawCard    = "drawCard"
	ActionLuck        = "luck"
	ActionVersion     = "version"
	ActionGetHelp     = "getHelp"
	ActionGroupStatus = "groupStatus"
)

type Action struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type StaticMessage struct {
	Id         int    `json:"id"`
	Mode       string `json:"mode"`
	CommonResp struct {
		Common               string `json:"common"`
		Report               string `json:"report"`
		CanNotRefresh        string `json:"can_not_refresh"`
		TooShortToRefresh    string `json:"too_short_to_refresh"`
		QueryIsRunning       string `json:"query_is_running"`
		NotValidNickname     string `json:"not_valid_nickname"`
		GetHelp              string `json:"get_help"`
		DrawCard             string `json:"draw_card"`
		Luck                 string `json:"luck"`
		GroupGetBanned       string `json:"group_get_banned"`
		UserGetBanned        string `json:"user_get_banned"`
		TodayUserQueryLimit  string `json:"today_user_query_limit"`
		TodayGroupQueryLimit string `json:"today_group_query_limit"`
		Version              string `json:"version"`
		LiveBroadcast        string `json:"live_broadcast"`
	} `json:"common_resp"`
	LuckResp struct {
		Is0          string `json:"is_0"`
		Between0130  string `json:"between_1_30"`
		Between3050  string `json:"between_30_50"`
		Between5070  string `json:"between_50_70"`
		Between7080  string `json:"between_70_80"`
		Between8095  string `json:"between_80_95"`
		Between95100 string `json:"between_95_100"`
		Is100        string `json:"is_100"`
	} `json:"luck_resp"`
}
