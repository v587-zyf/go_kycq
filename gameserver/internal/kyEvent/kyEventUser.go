package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/objs"
)

type KyEvent_roleCreate struct {
	*kyEventLog.KyEventPropBase
	HeroJob   int `json:"config_id1"`
	HeroSex   int `json:"config_id2"`
	HeroIndex int `json:"num1"`
}

type KyEvent_hero_create struct {
	*kyEventLog.KyEventPropBase
	HeroJob   int `json:"config_id1"`
	HeroSex   int `json:"config_id2"`
	HeroIndex int `json:"num1"`
}

type KyEvent_heroLvUp struct {
	*KyEventHeroPropBase
	Level     int `json:"num2"` //当前等级	数值	是
}

type KyEvent_online_count struct {
	Plat       int `json:"plat"`       //平台	数值	是
	Channel    int `json:"channel"`    //渠道	数值	是
	Serverid   int `json:"_serverid"`  //	区服	数值	是
	Online_num int `json:"online_num"` //在线人数	数值	是
}

func UserCreate(user *objs.User) {
	data := &KyEvent_roleCreate{
		KyEventPropBase: getEventBaseProp(user),
		HeroJob:         user.Heros[constUser.USER_HERO_MAIN_INDEX].Job,
		HeroSex:         user.Heros[constUser.USER_HERO_MAIN_INDEX].Sex,
		HeroIndex:       constUser.USER_HERO_MAIN_INDEX,
	}
	writeUserEvent(user, "create", data)
}

func UserHeroCreate(user *objs.User, heroIndex int) {
	data := &KyEvent_roleCreate{
		KyEventPropBase: getEventBaseProp(user),
		HeroJob:         user.Heros[heroIndex].Job,
		HeroSex:         user.Heros[heroIndex].Sex,
		HeroIndex:       heroIndex,
	}
	writeUserEvent(user, "hero_create", data)
}

func UserHeroLvUp(user *objs.User, heroIndex int) {

	data := &KyEvent_heroLvUp{
		KyEventHeroPropBase: getEventHeroBaseProp(user,heroIndex),
		Level:           user.Heros[heroIndex].ExpLvl,
	}
	writeUserEvent(user, "hero_level_up", data)

}

func OnlineTotal(total int) {
	data := &KyEvent_online_count{
		Plat:       base.Conf.Sdkconfig.KyPlatformId,
		Channel:    0,
		Serverid:   base.Conf.ServerId,
		Online_num: total,
	}
	writeUserEvent(nil, "online_count", data)
}
