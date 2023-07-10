package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gameserver/internal/objs"
)

type KyEvent_atlas_star_up struct {
	*kyEventLog.KyEventPropBase
	AtlasType int `json:"type"`       //图鉴类型
	AtlasId   int `json:"config_id1"` //图鉴id
	BeforeLv  int `json:"before_lv"`  //升级前等级
	Lv        int `json:"lv"`         //升级后等级
}

type KyEvent_atlas_change struct {
	*KyEventHeroPropBase
	AtlasId int `json:"config_id1"` //图鉴id
}

//图鉴升级记录
func AtlasStarUp(user *objs.User, atlasT, atlasId, lv int) {
	data := &KyEvent_atlas_star_up{
		KyEventPropBase: getEventBaseProp(user),
		AtlasType:       atlasT,
		AtlasId:         atlasId,
		BeforeLv:        lv - 1,
		Lv:              lv,
	}
	writeUserEvent(user, "atlas_star_up", data)
}

//图鉴镶嵌记录
func AtlasChange(user *objs.User, heroIndex, atlasId int) {
	data := &KyEvent_atlas_change{
		KyEventHeroPropBase: getEventHeroBaseProp(user, heroIndex),
		AtlasId:             atlasId,
	}
	writeUserEvent(user, "atlas_change", data)
}
