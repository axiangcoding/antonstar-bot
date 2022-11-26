package bot

import (
	"encoding/json"
	"github.com/axiangcoding/ax-web/logging"
	"github.com/axiangcoding/ax-web/static"
	"strings"
)

func SelectStaticMessage(id int) StaticMessage {
	var filename string
	if id == 0 {
		filename = "default.json"
	} else if id == 1 {
		filename = "two_dim.json"
	} else {
		filename = "default.json"
	}
	bytes := static.MustReadMessageFileAsBytes(filename)
	msg := StaticMessage{}
	err := json.Unmarshal(bytes, &msg)
	if err != nil {
		logging.Warn(err)
	}
	return msg
}

func ParseMessageCommand(msg string) *Action {
	sub := MessageGetCmdPrimaryMsgPattern.FindStringSubmatch(msg)
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
	case "完整查询":
		key = ActionFullQuery
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
	case "气运":
		key = ActionLuck
		break
	case "运气":
		key = ActionLuck
		break
	case "版本":
		key = ActionVersion
		break
	case "群状态":
		key = ActionGroupStatus
		break
	case "数据":
		key = ActionData
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
