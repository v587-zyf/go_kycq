package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gameserver/internal/objs"
	"fmt"
)

type KyEvent_change_hero_combat struct {
	*kyEventLog.KyEventPropBase
	HeroJob     int    `json:"config_id1"`      //武将职业id
	HeroSex     int    `json:"config_id2"`      //武将性别
	HeroIndex   int    `json:"num1"`            //武将索引
	Combat      int    `json:"combat"`          //战力
	Combat_Parm string `json:"log_event_parm1"` //战力组成 如：[2|300,3|400,4|5000]
}

func HeroCombat(user *objs.User, heroIndex int) {
	hero := user.Heros[heroIndex]
	data := &KyEvent_change_hero_combat{
		KyEventPropBase: getEventBaseProp(user),
		HeroJob:         hero.Job,
		HeroSex:         hero.Sex,
		HeroIndex:       heroIndex,
		Combat:          hero.Combat,
	}
	combatStr := "["
	if len(hero.ModuleCombat) > 0 {
		for m, c := range hero.ModuleCombat {
			combatStr += fmt.Sprintf(`%d|%d,`, m, c)
		}
		combatStr = combatStr[:len(combatStr)-1]
	}
	combatStr += "]"
	data.Combat_Parm = combatStr
	writeUserEvent(user, "change_hero_combat", data)
}
