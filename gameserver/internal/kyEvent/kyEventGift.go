package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gameserver/internal/objs"
)

type KyEvent_get_gift struct {
	*kyEventLog.KyEventPropBase
	Module   int `json:"type"`       //礼包类型	数值	是	跟充值的type一样的
	GiftId   int `json:"config_id1"` //礼包id	数值	是
	GiftLv   int `json:"num1"`       //礼包等级	数值	是	限时礼包用
	Coin_id  int `json:"coin_id"`    //货币id	数值	是	同货币变化中的货币id，消耗的货币id
	Coin_num int `json:"coin_num"`   //货币数量	数值	是
}

type KyEvent_get_first_pay struct {
	*kyEventLog.KyEventPropBase
	GiftId   int `json:"config_id1"` //礼包id	数值	是
	GetTimes int `json:"num1"`       //第几次领取
}

func UserGiftBuy(user *objs.User, module int, confId int, cost int, costNum int, giftLv int) {
	data := &KyEvent_get_gift{
		KyEventPropBase: getEventBaseProp(user),
		Module:          module,
		GiftId:          confId,
		GiftLv:          giftLv,
		Coin_id:         cost,
		Coin_num:        costNum,
	}
	writeUserEvent(user, "get_gift", data)
}

func UserFirstRechargeReward(user *objs.User, confId int, times int) {
	data := &KyEvent_get_first_pay{
		KyEventPropBase: getEventBaseProp(user),
		GiftId:          confId,
		GetTimes:        times,
	}
	writeUserEvent(user, "get_first_pay", data)
}
