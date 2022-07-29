package bot

import (
	"regexp"
	"strings"
)

var (
	MessageGetPrimaryInfoPattern = regexp.MustCompile(`^.*\[.*\](.*)$`)
)

var (
	ActionUnknown  = "unknown"
	ActionQuery    = "query"
	ActionRefresh  = "refresh"
	ActionReport   = "report"
	ActionDrawCard = "drawCard"
)

type Action struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

func ParseMessageCommand(msg string) *Action {
	sub := MessageGetPrimaryInfoPattern.FindStringSubmatch(msg)
	if len(sub) <= 1 {
		return nil
	}

	split := strings.Fields(strings.TrimSpace(sub[1]))
	if len(split) == 0 {
		return nil
	}
	var key string

	switch split[0] {
	case "查询":
		key = ActionQuery
		break
	case "刷新":
		key = ActionRefresh
		break
	case "举办":
		key = ActionReport
		break
	case "举报":
		key = ActionReport
		break
	case "抽卡":
		key = ActionDrawCard
		break
	default:
		key = ActionUnknown
	}
	if len(split) == 1 {
		return &Action{Key: key}
	}
	return &Action{
		Key:   key,
		Value: split[1],
	}
}
