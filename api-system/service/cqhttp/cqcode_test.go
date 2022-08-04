package cqhttp

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var tests1 = []string{
	"[CQ:at,qq=2362794289] 你在说些什么我怎么不知道\\r\\n啊啊啊啊，焯",
	"你在说些什么我怎么不知道 [CQ:at,qq=2362794289] \\r\\n啊啊啊啊，焯",
	"你在说些什么我怎么不知道 [CQ:at,qq=2362794289] [CQ:at,qq=2362794289] [CQ:at,qq=2362794289] [CQ:at,qq=2362794289]",
}

var tests2 = []string{
	"你在说些什么我怎么不知道\\r\\n啊啊啊啊，焯",
	"你在说些什么我怎么不知道]",
	"[猪鼻",
}

func TestContainsCqCode(t *testing.T) {
	for i, tt := range tests1 {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if !MustContainsCqCode(tt) {
				t.Fail()
			}
		})
	}

	for i, tt := range tests2 {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if MustContainsCqCode(tt) {
				t.Fail()
			}
		})
	}
}

var tests3 = []struct {
	msg     string
	extract string
}{
	{msg: "[CQ:at,qq=2362794289] 你在说些什么我怎么不知道\\r\\n啊啊啊啊，焯", extract: "CQ:at,qq=2362794289"},
	{msg: "你在说些什么我怎么不知道\\r\\n啊啊啊啊，焯", extract: ""},
}

func TestMustGetCqCode(t *testing.T) {
	for i, tt := range tests3 {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			qq := MustGetCqCode(tt.msg)
			assert.Equal(t, qq, tt.extract)
		})
	}
}

var tests4 = []struct {
	msg     string
	extract int64
}{
	{msg: "[CQ:at,qq=2362794289] 你在说些什么我怎么不知道\\r\\n啊啊啊啊，焯", extract: int64(2362794289)},
	{msg: "你在说些什么我怎么不知道\\r\\n啊啊啊啊，焯", extract: int64(0)},
	{msg: "！[CQ:at,qq=all] 试试", extract: int64(0)},
}

func TestGetAtQQ(t *testing.T) {
	for i, tt := range tests4 {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			qq := MustGetCqCodeAtQQ(tt.msg)
			assert.Equal(t, qq, tt.extract)
		})
	}
}

var tests5 = []string{
	".cqbot 查询",
	".cqbot 帮助",
	".cqbot",
	" .cqbot 帮助",
}

func TestMustContainsTrigger(t *testing.T) {
	for i, tt := range tests5 {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			bContains := MustContainsTrigger(tt)
			if !bContains {
				t.Fail()
			}
		})
	}
}
