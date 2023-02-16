package main

import (
	"fmt"
	"github.com/axiangcoding/antonstar-bot/service/cardfight"
)

// WIP 卡牌对战程序
func main() {
	card99A := cardfight.InitCarItem("ZTZ99A", "用户一号", 10, 9, 7, 6)
	cardL2A6 := cardfight.InitCarItem("豹2A6", "用户二号", 10, 7, 8, 6)

	match := cardfight.FightMatch{
		A: *card99A,
		B: *cardL2A6,
	}
	result := match.Fight()

	txt := cardfight.GenerateFightText(result)
	fmt.Println(txt)
}
