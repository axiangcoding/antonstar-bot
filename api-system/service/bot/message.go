package bot

import (
	"regexp"
	"strings"
)

var (
	MessageGetPrimaryInfoPattern = regexp.MustCompile(`^.*\[.*\](.*)$`)
)

var (
	ActionUnknown = "unknown"
	ActionQuery   = "query"
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
	if len(split) < 2 {
		return nil
	}
	var key string

	switch split[0] {
	case "查询":
		key = ActionQuery
		break
	default:
		key = ActionUnknown
	}

	return &Action{
		Key:   key,
		Value: split[1],
	}
}
