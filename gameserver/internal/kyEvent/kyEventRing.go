package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gameserver/internal/objs"
)

type KyEvent_ring_phantom_star_up struct {
	*KyEventHeroPropBase
	RingPos    int `json:"config_id1"` //特戒部位
	BeforeStar int `json:"before_lv"`  //升级前星数
	Star       int `json:"lv"`         //升级后星数
}

type KyEvent_ring_strengthen struct {
	*KyEventHeroPropBase
	RingPos  int `json:"config_id1"` //特戒部位
	BeforeLv int `json:"before_lv"`  //升级前等级
	Lv       int `json:"lv"`         //升级后等级
}

type KyEvent_ring_fuse struct {
	*kyEventLog.KyEventPropBase
	Id      int         `json:"id"`      //合成后特戒id
	Consume map[int]int `json:"consume"` //消耗材料
}

//特戒幻灵升星
func RingPhantomStarUp(user *objs.User, heroIndex, ringPos, star int) {
	data := &KyEvent_ring_phantom_star_up{
		KyEventHeroPropBase: getEventHeroBaseProp(user, heroIndex),
		RingPos:             ringPos,
		BeforeStar:          star - 1,
		Star:                star,
	}
	writeUserEvent(user, "ring_phantom_star_up", data)
}

//特戒强化
func RingStrengthen(user *objs.User, heroIndex, ringPos, lv int) {
	data := &KyEvent_ring_strengthen{
		KyEventHeroPropBase: getEventHeroBaseProp(user, heroIndex),
		RingPos:             ringPos,
		BeforeLv:            lv - 1,
		Lv:                  lv,
	}
	writeUserEvent(user, "ring_strengthen", data)
}

//特戒融合
func RingFuse(user *objs.User, id int, consume map[int]int) {
	data := &KyEvent_ring_fuse{
		KyEventPropBase: getEventBaseProp(user),
		Id:              id,
		Consume:         consume,
	}
	writeUserEvent(user, "ring_fuse", data)
}
