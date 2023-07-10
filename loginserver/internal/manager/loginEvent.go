package manager

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/golibs/common"
	"cqserver/loginserver/conf"
)

func writeUserEvent(openId string, did string, channel string, eventName string) {

	baseEvent := &kyEventLog.KyEventBase{
		Ouid:      openId,
		Timestamp: int(common.GetNowMillisecond()),
		Did:       did,
		Event:     eventName,
		Project:   conf.Conf.Sdkconfig.KyprojectName,
		Properties: &kyEventLog.KyEventPropBase{
			Plat:       conf.Conf.Sdkconfig.KyPlatformId,
			Channel:    channel,
			Serverid:   0,
			Source_sid: 0,
			Openid:     openId,
		},
	}
	kyEventLog.WriteEventByData(baseEvent)
}
