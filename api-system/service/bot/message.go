package bot

import (
	"regexp"
	"strings"
)

const (
	RespDontKnowAction    = "我不道你在说什么，请按照指令提问，注意不要缺少空格"
	RespHelp              = "我不道你想干啥，输入”帮助“查看可用的命令"
	RespReport            = "举报已被记录，该举报仅代表玩家意见，仅作为参考，不是官方实锤"
	RespCanNotRefresh     = "目前无法查询，请稍后重试"
	RespTooShortToRefresh = "对不起，距上次刷新间隔太短，不允许刷新"
	RespRunningQuery      = "正在发起查询，请耐心等待..."
	RespNotAValidNickname = "我说你这id不对吧，别逗我玩"
	RespGetHelp           = "我啷个晓得怎么帮你，找33去"
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
	ActionGetHelp  = "getHelp"
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
	case "帮助":
		key = ActionGetHelp
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
