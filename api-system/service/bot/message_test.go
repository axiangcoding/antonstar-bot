package bot

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var tests = []struct {
	msg    string
	action Action
}{
	{
		msg: "[CQ:at,qq=2362794289] 查询 GodFather_33",
		action: Action{Key: ActionQuery,
			Value: "GodFather_33"},
	},
	{
		msg: "[CQ:at,qq=2362794289]     	查询 	GodFather_33",
		action: Action{Key: ActionQuery,
			Value: "GodFather_33"},
	},
	{
		msg: "[CQ:at,qq=2362794289]查询 	GodFather_33",
		action: Action{Key: ActionQuery,
			Value: "GodFather_33"},
	},
}

func TestParseMessageCommand(t *testing.T) {
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			action := ParseMessageCommand(tt.msg)
			assert.Equal(t, tt.action, *action)
		})
	}
}
