package kyEvent

import (
	"cqserver/gameserver/internal/objs"
	"time"
)

type KyEvent_exp_upgrade struct {
	*KyEventHeroPropBase
	BeforeLv int `json:"before_lv"`
	Lv       int `json:"lv"`
	UpTime   int `json:"up_time"`
}

func ExpLvUp(user *objs.User, heroIndex int, beforeLv int) {
	data := &KyEvent_exp_upgrade{
		KyEventHeroPropBase: getEventHeroBaseProp(user, heroIndex),
		BeforeLv:            beforeLv,
		Lv:                  user.Heros[heroIndex].ExpLvl,
		UpTime:              int(time.Now().Unix()),
	}
	writeUserEvent(user, "exp_up", data)
}
