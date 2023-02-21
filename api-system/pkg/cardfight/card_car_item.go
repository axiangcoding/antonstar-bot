package cardfight

import (
	"fmt"
	"math"
	"math/rand"

	"strings"
)

type CardCarItem struct {
	CardItem
	// 机动
	velocity float64
	// 火力
	firepower float64
	// 防护
	protection float64
	// 车辆模块
	module CardCarModule
}

type CardCarModule struct {
	// 成员
	Crew float64
	// 履带
	Track float64
	// 垂稳
	GunSteady float64
	// 炮管
	Barrel float64
	// 引擎
	Motor float64
	// 弹药架
	AmmunitionRacks float64
}

func InitCarItem(name string,
	user string, memberProficiency float64,
	velocity float64, firepower float64, protection float64) *CardCarItem {
	return &CardCarItem{
		CardItem: CardItem{
			name:              name,
			user:              user,
			cardType:          CardItemCar,
			memberProficiency: memberProficiency,
		},
		velocity:   velocity,
		firepower:  firepower,
		protection: protection,
		module: CardCarModule{
			Crew:            100,
			Track:           100,
			GunSteady:       100,
			Barrel:          100,
			Motor:           100,
			AmmunitionRacks: 100,
		},
	}
}

func (i *CardCarItem) TakeStepWithCar(enemy *CardCarItem) []string {
	res := make([]string, 0)
	if i.module.Barrel <= 33 {
		res = append(res, "由于炮管故障无法攻击，"+i.repair())
		res = append(res, i.ModuleStatus())
	} else if i.module.GunSteady <= 33 {
		res = append(res, "由于垂稳故障无法攻击，"+i.repair())
		res = append(res, i.ModuleStatus())
	} else {
		res = append(res, i.attack(enemy))
		res = append(res, enemy.ModuleStatus())
	}
	return res
}

func (i *CardCarItem) IsDead() (bool, string) {
	if i.module.Crew <= 33 {
		msg := "%s 成员不足，退出战斗"
		return true, fmt.Sprintf(msg, i.displayName())
	}
	if i.module.AmmunitionRacks <= 0 {
		msg := "%s 弹药殉爆！"
		return true, fmt.Sprintf(msg, i.displayName())
	}
	return false, ""
}

func (i *CardCarItem) ModuleStatus() string {
	res := make([]string, 0)
	module := i.module
	if module.AmmunitionRacks < 100 {
		res = append(res, fmt.Sprintf("弹药架（%.2f%%）", math.Max(module.AmmunitionRacks, 0)))
	}
	if module.Crew < 100 {
		res = append(res, fmt.Sprintf("成员（%.2f%%）", math.Max(module.Crew, 0)))
	}
	if module.Barrel < 100 {
		res = append(res, fmt.Sprintf("炮管（%.2f%%）", math.Max(module.Barrel, 0)))
	}
	if module.GunSteady < 100 {
		res = append(res, fmt.Sprintf("垂稳（%.2f%%）", math.Max(module.GunSteady, 0)))
	}
	if module.Motor < 100 {
		res = append(res, fmt.Sprintf("引擎（%.2f%%）", math.Max(module.Motor, 0)))
	}
	if module.Track < 100 {
		res = append(res, fmt.Sprintf("履带（%.2f%%）", math.Max(module.Track, 0)))
	}
	if len(res) > 0 {
		return fmt.Sprintf("%s 状态：%s", i.displayName(), strings.Join(res, "，"))
	} else {
		return fmt.Sprintf("%s 完好无损", i.displayName())
	}
}

func (i *CardCarItem) attack(enemy *CardCarItem) string {
	text := i.attackModules(40, i.firepower, enemy.protection, enemy.velocity, &enemy.module)
	msgTemplate := "%s 攻击 %s，%s"
	return fmt.Sprintf(msgTemplate, i.displayName(), enemy.displayName(), text)
}

func (i *CardCarItem) repair() string {
	text := i.repairModules(30, i.memberProficiency, &i.module)
	msgTemplate := "%s 进行了维修，修复了 %s"
	return fmt.Sprintf(msgTemplate, i.displayName(), text)
}

func (i *CardCarItem) attackModules(baseDamage float64, attack float64, defend float64, dodge float64, module *CardCarModule) string {
	res := make([]string, 0)
	baseProb := 0.4
	prob := baseProb * (2 - dodge*0.1)
	if CalProbabilities(prob * 0.2) {
		var damage float64
		var txt string
		if CalProbabilities(0.5) {
			damage = 100
			txt = "弹药架（-100% 暴击！）"
		} else {
			damage = calDamage(baseDamage, attack, defend)
			txt = fmt.Sprintf("弹药架（-%.2f%%）", damage)
		}
		module.AmmunitionRacks -= damage

		if damage == 100 {
			return txt
		}
		res = append(res, txt)
	}

	if CalProbabilities(prob) {
		damage := calDamage(baseDamage, attack, defend)
		module.Crew -= damage
		res = append(res, fmt.Sprintf("成员（-%.2f%%）", damage))
	}
	if CalProbabilities(prob * 1.1) {
		damage := calDamage(baseDamage, attack, defend)
		module.Track -= damage
		res = append(res, fmt.Sprintf("履带（-%.2f%%）", damage))
	}
	if CalProbabilities(prob) {
		damage := calDamage(baseDamage, attack, defend)
		module.GunSteady -= damage
		res = append(res, fmt.Sprintf("垂稳（-%.2f%%）", damage))

	}
	if CalProbabilities(prob) {
		damage := calDamage(baseDamage, attack, defend)
		module.Barrel -= damage
		res = append(res, fmt.Sprintf("炮管（-%.2f%%）", damage))
	}
	if CalProbabilities(prob) {
		damage := calDamage(baseDamage, attack, defend)
		module.Motor -= damage
		res = append(res, fmt.Sprintf("引擎（-%.2f%%）", damage))
	}

	if len(res) > 0 {
		return "对其造成了 " + strings.Join(res, "，")
	} else {
		return "未能击穿他们的装甲"
	}
}

func (i *CardCarItem) repairModules(baseRepair float64, Proficiency float64, module *CardCarModule) string {
	res := make([]string, 0)
	repair := baseRepair * (0.1 * Proficiency)

	if repair > 0 {
		f := math.Min(100-module.AmmunitionRacks, repair)
		module.AmmunitionRacks += f
		repair -= f
		res = append(res, "弹药架")
	}
	if repair > 0 {
		f := math.Min(100-module.Crew, repair)
		module.Crew += f
		repair -= f
		res = append(res, "成员")
	}
	if repair > 0 {
		f := math.Min(100-module.Barrel, repair)
		module.Barrel += f
		repair -= f
		res = append(res, "炮管")
	}
	if repair > 0 {
		f := math.Min(100-module.GunSteady, repair)
		module.GunSteady += f
		repair -= f
		res = append(res, "垂稳")
	}
	if repair > 0 {
		f := math.Min(100-module.Motor, repair)
		module.Motor += f
		repair -= f
		res = append(res, "引擎")
	}
	if repair > 0 {
		f := math.Min(100-module.Track, repair)
		module.Track += f
		repair -= f
		res = append(res, "履带")
	}
	return strings.Join(res, "，")
}

func calDamage(baseDamage float64, attack float64, defend float64) float64 {
	damage := 40.0
	if baseDamage != 0.0 {
		damage = baseDamage
	}

	if defend > attack {
		damage *= 0.5
	} else {
		damage *= 1 + (attack-defend)*0.1
	}
	floatMin := 0.7
	floatMax := 1.1
	finalDamage := (floatMin + rand.Float64()*(floatMax-floatMin)) * damage
	return finalDamage
}
