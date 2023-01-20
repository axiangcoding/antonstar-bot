package cqhttp

import (
	"github.com/axiangcoding/ax-web/logging"
	"regexp"
	"strconv"
)

var (
	CqCodePattern     = regexp.MustCompile(`^.*\[(.*)\].*$`)
	CqCodeAtQQPattern = regexp.MustCompile(`^CQ:at,qq=(\d+)$`)
	CqTriggerPattern  = regexp.MustCompile(`^\s*\.?cqbot.*$`)
)

func MustContainsTrigger(message string) bool {
	return CqTriggerPattern.MatchString(message)
}

func MustContainsCqCode(message string) bool {
	return CqCodePattern.MatchString(message)
}

func MustGetCqCode(message string) string {
	sub := CqCodePattern.FindStringSubmatch(message)
	if len(sub) <= 1 {
		return ""
	}
	return sub[1]
}

func MustGetCqCodeAtQQ(message string) int64 {
	cqCode := MustGetCqCode(message)
	sub := CqCodeAtQQPattern.FindStringSubmatch(cqCode)
	if len(sub) <= 1 {
		return 0
	}
	parseInt, err := strconv.ParseInt(sub[1], 10, 64)
	if err != nil {
		logging.Errorf("parse qq string: %s error. %s", sub[1], err)
		return 0
	}
	return parseInt
}
