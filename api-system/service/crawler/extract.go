package crawler

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/axiangcoding/ax-web/data/table"
	"github.com/gocolly/colly/v2"
	"strconv"
	"strings"
	"time"
)

func ExtractGaijinData(e *colly.HTMLElement) table.GameUser {
	var data table.GameUser
	dom := e.DOM
	data.Nick = strings.TrimSpace(dom.Find("li[class=user-profile__data-nick]").Text())
	data.Clan = dom.Find("a[class=user-profile__data-link]").Text()
	var clanUrl string
	clanUrlSuffix := dom.Find("a[class=user-profile__data-link]").AttrOr("href", "")
	if clanUrlSuffix == "" {
		clanUrl = ""
	} else {
		clanUrl = "https://warthunder.com" + clanUrlSuffix
	}
	data.ClanUrl = clanUrl
	length := dom.Find("div[class=user-profile__data-nick--banned]").Length()
	banned := length == 1
	data.Banned = &banned
	data.RegisterDate = extractRegisterDate(dom.Find("li[class=user-profile__data-regdate]").Text())
	data.Title = extractTitle(dom.Find("li[class=user-profile__data-item]").Eq(0).Text())
	data.Level = extractLevel(dom.Find("li[class=user-profile__data-item]").Eq(1).Text())

	userStatKeys := extractTableKeys(
		dom.Find("div[class='user-stat__list-row user-stat__list-row--with-head']>" +
			"ul[class='user-stat__list user-stat__list--titles']>li"))
	data.StatAb = extractUserStat(userStatKeys,
		dom.Find("div[class='user-stat__list-row user-stat__list-row--with-head']>"+
			"ul[class='user-stat__list arcadeFightTab']>li"))

	data.StatRb = extractUserStat(userStatKeys,
		dom.Find("div[class='user-stat__list-row user-stat__list-row--with-head']>"+
			"ul[class='user-stat__list historyFightTab']>li"))

	data.StatSb = extractUserStat(userStatKeys,
		dom.Find("div[class='user-stat__list-row user-stat__list-row--with-head']>"+
			"ul[class='user-stat__list simulationFightTab']>li"))

	aviationDom := dom.Find("div[class='user-profile__stat user-stat user-stat--tabs']>div[class='user-stat__list-row']").Eq(0)
	userRateAviationKeys := extractTableKeys(
		aviationDom.Find("ul[class='user-stat__list user-stat__list--titles']>li"))
	data.AviationRateAb = extractAviationUserRate(userRateAviationKeys, aviationDom.Find("ul[class='user-stat__list arcadeFightTab']>li"))
	data.AviationRateRb = extractAviationUserRate(userRateAviationKeys, aviationDom.Find("ul[class='user-stat__list historyFightTab']>li"))
	data.AviationRateSb = extractAviationUserRate(userRateAviationKeys, aviationDom.Find("ul[class='user-stat__list simulationFightTab']>li"))

	groundDom := dom.Find("div[class='user-profile__stat user-stat user-stat--tabs']>div[class='user-stat__list-row']").Eq(1)
	userRateGroundKeys := extractTableKeys(
		groundDom.Find("ul[class='user-stat__list user-stat__list--titles']>li"))
	data.GroundRateAb = extractGroundUserRate(userRateGroundKeys, groundDom.Find("ul[class='user-stat__list arcadeFightTab']>li"))
	data.GroundRateRb = extractGroundUserRate(userRateGroundKeys, groundDom.Find("ul[class='user-stat__list historyFightTab']>li"))
	data.GroundRateSb = extractGroundUserRate(userRateGroundKeys, groundDom.Find("ul[class='user-stat__list simulationFightTab']>li"))

	fleetDom := dom.Find("div[class='user-profile__stat user-stat user-stat--tabs']>div[class='user-stat__list-row']").Eq(2)
	userRateFleetKeys := extractTableKeys(
		fleetDom.Find("ul[class='user-stat__list user-stat__list--titles']>li"))
	data.FleetRateAb = extractFleetUserRate(userRateFleetKeys, fleetDom.Find("ul[class='user-stat__list arcadeFightTab']>li"))
	data.FleetRateRb = extractFleetUserRate(userRateFleetKeys, fleetDom.Find("ul[class='user-stat__list historyFightTab']>li"))
	data.FleetRateSb = extractFleetUserRate(userRateFleetKeys, fleetDom.Find("ul[class='user-stat__list simulationFightTab']>li"))
	return data
}

func extractRegisterDate(str string) time.Time {
	str = strings.TrimSpace(str)
	str = strings.Split(str, " ")[1]
	parse, _ := time.Parse("02.01.2006", str)
	return parse
}

func extractTitle(str string) string {
	trim := strings.TrimSpace(str)
	return strings.ReplaceAll(trim, "\\t", "")
}

func extractLevel(str string) int {
	str = strings.TrimSpace(str)
	str = strings.Split(str, " ")[1]
	i, _ := strconv.Atoi(str)
	return i
}

func extractTableKeys(s *goquery.Selection) []string {
	var keys []string
	s.Each(func(i int, sub *goquery.Selection) {
		keys = append(keys, strings.TrimSpace(sub.Text()))
	})
	return keys
}

func extractUserStat(keys []string, s *goquery.Selection) table.UserStat {
	mp := make(map[string]string)
	s.Each(func(i int, sub *goquery.Selection) {
		mp[keys[i]] = strings.TrimSpace(sub.Text())
	})
	return table.UserStat{
		TotalMission:         parseCommonNumber(mp["任务总数"]),
		WinRate:              parseWinRate(mp["作战胜率"]),
		GroundDestroyCount:   parseCommonNumber(mp["地面单位摧毁数"]),
		FleetDestroyCount:    parseCommonNumber(mp["水面单位摧毁数"]),
		GameTime:             mp["游戏时间"],
		AviationDestroyCount: parseCommonNumber(mp["空中单位摧毁数"]),
		WinCount:             parseCommonNumber(mp["胜利场次"]),
		SliverEagleEarned:    parseSENumber(mp["银狮获得数"]),
		DeadCount:            parseCommonNumber(mp["阵亡数"]),
	}
}

func extractGroundUserRate(keys []string, s *goquery.Selection) table.GroundRate {
	mp := make(map[string]string)
	s.Each(func(i int, sub *goquery.Selection) {
		mp[keys[i]] = strings.TrimSpace(sub.Text())
	})
	return table.GroundRate{
		GameCount:              parseCommonNumber(mp["参战次数(陆战)"]),
		GroundVehicleGameCount: parseCommonNumber(mp["参战次数(地面载具)"]),
		TDGameCount:            parseCommonNumber(mp["参战次数(坦克歼击车)"]),
		HTGameCount:            parseCommonNumber(mp["参战次数(重型坦克)"]),
		SPAAGameCount:          parseCommonNumber(mp["参战次数(自行防空炮)"]),
		GameTime:               mp["游戏时长(陆战)"],
		GroundVehicleGameTime:  mp["游戏时长(地面单位)"],
		TDGameTime:             mp["游戏时长(坦克歼击车)"],
		HTGameTime:             mp["游戏时长(重型坦克)"],
		SPAAGameTime:           mp["游戏时长(自行防空炮)"],
		TotalDestroyCount:      parseCommonNumber(mp["击毁目标总计"]),
		AviationDestroyCount:   parseCommonNumber(mp["空中单位摧毁数"]),
		GroundDestroyCount:     parseCommonNumber(mp["地面单位摧毁数"]),
		FleetDestroyCount:      parseCommonNumber(mp["水面单位摧毁数"]),
	}
}

func extractAviationUserRate(keys []string, s *goquery.Selection) table.AviationRate {
	mp := make(map[string]string)
	s.Each(func(i int, sub *goquery.Selection) {
		mp[keys[i]] = strings.TrimSpace(sub.Text())
	})
	return table.AviationRate{
		GameCount:            parseCommonNumber(mp["参战次数(空战)"]),
		FighterGameCount:     parseCommonNumber(mp["参战次数(战斗机)"]),
		BomberGameCount:      parseCommonNumber(mp["参战次数(轰炸机)"]),
		AttackerGameCount:    parseCommonNumber(mp["参战次数(攻击机)"]),
		GameTime:             mp["游戏时长(空战)"],
		FighterGameTime:      mp["游戏时长(战斗机)"],
		BomberGameTime:       mp["游戏时长(轰炸机)"],
		AttackerGameTime:     mp["游戏时长(攻击机)"],
		TotalDestroyCount:    parseCommonNumber(mp["击毁目标总计"]),
		AviationDestroyCount: parseCommonNumber(mp["空中单位摧毁数"]),
		GroundDestroyCount:   parseCommonNumber(mp["地面单位摧毁数"]),
		FleetDestroyCount:    parseCommonNumber(mp["水面单位摧毁数"]),
	}
}

func extractFleetUserRate(keys []string, s *goquery.Selection) table.FleetRate {
	mp := make(map[string]string)
	s.Each(func(i int, sub *goquery.Selection) {
		mp[keys[i]] = strings.TrimSpace(sub.Text())
	})
	return table.FleetRate{
		GameCount:               parseCommonNumber(mp["参战次数(海军)"]),
		FleetGameCount:          parseCommonNumber(mp["参战次数(舰船)"]),
		TorpedoBoatGameCount:    parseCommonNumber(mp["参战次数(鱼雷艇)"]),
		GunboatGameCount:        parseCommonNumber(mp["参战次数(炮艇)"]),
		TorpedoGunboatGameCount: parseCommonNumber(mp["参战次数(鱼雷炮艇)"]),
		SubmarineHuntGameCount:  parseCommonNumber(mp["参战次数(猎潜艇))"]),
		DestroyerGameCount:      parseCommonNumber(mp["参战次数(驱逐舰)"]),
		NavyBargeGameCount:      parseCommonNumber(mp["参战次数(海军驳渡船)"]),
		GameTime:                mp["游戏时长(海战)"],
		FleetGameTime:           mp["游戏时长(船舰)"],
		TorpedoBoatGameTime:     mp["游戏时长(鱼雷艇)"],
		GunboatGameTime:         mp["游戏时长(炮艇)"],
		TorpedoGunboatGameTime:  mp["游戏时长(鱼雷炮艇)"],
		SubmarineHuntGameTime:   mp["游戏时长(猎潜艇)"],
		DestroyerGameTime:       mp["游戏时长(驱逐舰)"],
		NavyBargeGameTime:       mp["游戏时长(海军驳渡船)"],
		TotalDestroyCount:       parseCommonNumber(mp["击毁目标总计"]),
		AviationDestroyCount:    parseCommonNumber(mp["空中单位摧毁数"]),
		GroundDestroyCount:      parseCommonNumber(mp["地面单位摧毁数"]),
		FleetDestroyCount:       parseCommonNumber(mp["水面单位摧毁数"]),
	}
}

func parseWinRate(str string) float64 {
	str = strings.ReplaceAll(str, "%", "")
	num, _ := strconv.Atoi(str)
	return float64(num) / 100
}

func parseCommonNumber(str string) int {
	str = strings.ReplaceAll(str, ",", "")
	atoi, _ := strconv.Atoi(str)
	return atoi
}
func parseSENumber(str string) int64 {
	str = strings.ReplaceAll(str, ",", "")
	atoi, _ := strconv.ParseInt(str, 10, 64)
	return atoi
}
