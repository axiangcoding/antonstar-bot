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
	case "完整查询":
		key = ActionFullQuery
	case "刷新":
		key = ActionRefresh
	case "举办":
		key = ActionReport
	case "举报":
		key = ActionReport
	case "抽卡":
		key = ActionDrawCard
	case "帮助":
		key = ActionGetHelp
	case "气运":
		key = ActionLuck
	case "运气":
		key = ActionLuck
	case "版本":
		key = ActionVersion
	case "群状态":
		key = ActionGroupStatus
	case "数据":
		key = ActionData
	case "管理":
		key = ActionManager
	case "绑定":
		key = ActionBinding
	case "解绑":
		key = ActionUnbinding
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
