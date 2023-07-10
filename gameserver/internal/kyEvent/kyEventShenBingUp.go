package kyEvent

import (
	"cqserver/gameserver/internal/objs"
)

type KyEvent_sheng_bin_up struct {
	*KyEventHeroPropBase
	ShengBinType int `json:"type"`      //神兵类型
	BeforeLv     int `json:"before_lv"` //升级前等级
	Lv           int `json:"lv"`        //升级后等级
}

//神兵升级记录
func ShengBinUp(user *objs.User, heroIndex, shengBinType, beforeLv, afterLv int) {
	data := &KyEvent_sheng_bin_up{
		KyEventHeroPropBase: getEventHeroBaseProp(user, heroIndex),
		ShengBinType:        shengBinType,
		BeforeLv:            beforeLv,
		Lv:                  afterLv,
	}
	writeUserEvent(user, "sheng_bin_up", data)
}
