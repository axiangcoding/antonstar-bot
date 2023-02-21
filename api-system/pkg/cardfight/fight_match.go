package cardfight

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

type FightMatch struct {
	A CardCarItem
	B CardCarItem
}

func (m FightMatch) Fight() []string {
	res := make([]string, 0)
	res = append(res, m.A.StartFight())
	res = append(res, m.B.StartFight())
	fst, sec := m.decideOrderBothCar()
	res = append(res, fmt.Sprintf("%s 取得了先手！", fst.displayName()))
	maxStep := 10

	for maxStep > 0 {
		serverDown := CalProbabilities(0.01)
		if serverDown {
			res = append(res, "====服务器已断开连接====")
			break
		}
		res = append(res, fst.TakeStepWithCar(sec)...)
		if dead, s := sec.IsDead(); dead {
			res = append(res, s)
			break
		}
		//res = append(res, sec.ModuleStatus())
		res = append(res, sec.TakeStepWithCar(fst)...)
		if dead, s := fst.IsDead(); dead {
			res = append(res, s)
			break
		}
		//res = append(res, fst.ModuleStatus())
		maxStep--
	}
	if maxStep == 0 {
		res = append(res, "对战结束，未分出胜负")
	}
	return res
}

func (m FightMatch) decideOrderBothCar() (*CardCarItem, *CardCarItem) {
	if m.A.velocity > m.B.velocity {
		return &m.A, &m.B
	} else {
		return &m.B, &m.A
	}
}

func GenerateFightText(steps []string) string {
	res := ""
	for i, step := range steps {
		res += fmt.Sprintf("%d - %s\n", i, step)
	}
	return res
}

func CalProbabilities(percent float64) bool {
	if 0 > percent || percent > 1000 {
		return false
	}
	n, _ := rand.Int(rand.Reader, big.NewInt(1000))
	intn := n.Int64()
	return intn <= int64(math.Floor(percent*1000.0))
}
