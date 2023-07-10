package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/common"
	"strconv"
)

type KyEventHeroPropBase struct {
	*kyEventLog.KyEventPropBase
	HeroIndex int `json:"num1"`
	Combat    int `json:"combat"`
}

func writeUserEvent(user *objs.User, eventName string, eventData interface{}) {

	baseEvent := &kyEventLog.KyEventBase{
		Ouid:       "ouid",
		Timestamp:  int(common.GetNowMillisecond()),
		Did:        "did",
		Event:      eventName,
		Project:    base.Conf.Sdkconfig.KyprojectName,
		Properties: eventData,
	}
	if user != nil {
		baseEvent.Ouid = user.OpenId
		baseEvent.Did = user.DeviceId
	}
	kyEventLog.WriteEventByData(baseEvent)
}

func getEventBaseProp(user *objs.User) *kyEventLog.KyEventPropBase {

	propertiesBase := &kyEventLog.KyEventPropBase{
		Terminal:   user.Origin,
		Plat:       base.Conf.Sdkconfig.KyPlatformId,
		Channel:    strconv.Itoa(user.ChannelId),
		Serverid:   base.Conf.ServerId,
		Source_sid: user.ServerId,
		Openid:     user.OpenId,
		Roleid:     strconv.Itoa(user.Id),
		Rolename:   user.NickName,
		Vip:        user.VipLevel,
		Combat:     user.Combat,
		GuildId:    user.GuildData.NowGuildId,
	}
	if h, ok := user.Heros[constUser.USER_HERO_MAIN_INDEX]; ok {
		propertiesBase.Level = h.ExpLvl
	}
	return propertiesBase
}

func getEventHeroBaseProp(user *objs.User, heroIndex int) *KyEventHeroPropBase {

	propertiesBase := &KyEventHeroPropBase{
		KyEventPropBase: getEventBaseProp(user),
		HeroIndex:       heroIndex,
	}
	if heroIndex == -1 {
		heroIndex = constUser.USER_HERO_MAIN_INDEX
	}
	heroInfo := user.Heros[heroIndex]
	if heroInfo != nil {
		propertiesBase.Combat = heroInfo.Combat
	}
	return propertiesBase
}
