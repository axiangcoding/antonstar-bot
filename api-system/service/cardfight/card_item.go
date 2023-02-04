package cardfight

import (
	"fmt"
	"unicode/utf8"
)

const (
	CardItemCar = iota
	CardItemJet
)

type CardItem struct {
	name string
	// 用户名称
	user string
	// 卡牌类型
	cardType int
	// 成员熟练度
	memberProficiency float64
}

func (i *CardItem) displayName() string {
	var shortName string
	maxLen := 4
	if utf8.RuneCountInString(i.user) > maxLen {
		runes := []rune(i.user)
		shortName = string(runes[:maxLen]) + "..."
	} else {
		shortName = i.user
	}
	return fmt.Sprintf("%s(%s)", i.name, shortName)
}

func (i *CardItem) TakeStepWithCar(item *CardCarItem) string {
	return ""
}

func (i *CardItem) IsDead() (bool, string) {
	return true, ""
}

func (i *CardItem) StartFight() string {
	msgTemplate := "%s 进入了战斗"
	return fmt.Sprintf(msgTemplate, i.displayName())
}

func (i *CardItem) ModuleStatus() string {
	msg := "%s 当前模块状态"
	return fmt.Sprintf(msg, i.displayName())
}
