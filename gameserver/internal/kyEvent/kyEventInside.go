package kyEvent

import "cqserver/gameserver/internal/objs"

type KyEvent_inside_star_up struct {
	*KyEventHeroPropBase
	Acupoint   int `json:"config_id1"` //穴位
	BeforeStar int `json:"before_lv"`  //升级前星数
	Star       int `json:"lv"`         //升级后星数
}

//内功升星
func InsideStarUp(user *objs.User, heroIndex, acupoint, beforeStar, star int) {
	data := &KyEvent_inside_star_up{
		KyEventHeroPropBase: getEventHeroBaseProp(user, heroIndex),
		Acupoint:            acupoint,
		BeforeStar:          star - 1,
		Star:                star,
	}
	writeUserEvent(user, "inside_star_up", data)
}
