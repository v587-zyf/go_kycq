package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
)

type KyEvent_item_coin struct {
	*kyEventLog.KyEventPropBase
	ChangeType      int    `json:"type"`            //	变动类型(1:增加 2:减少)	数值	是
	Coin_id         int    `json:"coin_id"`         //货币id	数值	是
	Coin_num        int    `json:"coin_num"`        //变动数量	数值	是
	Coin_value      int    `json:"coin_value"`      //当前数量	数值	是
	Log_event_parm1 string `json:"log_event_parm1"` //货币流动一级原因	字符	是
	Log_event_parm2 string `json:"log_event_parm2"` //道具流动二级原因	字符	是
}

type KyEvent_item struct {
	*kyEventLog.KyEventPropBase
	ChangeType      int    `json:"type"`            //	变动类型(1:增加 2:减少)	数值	是
	Main_type       int    `json:"main_type"`       //道具大类	数值	是
	Sub_type        int    `json:"sub_type"`        //道具小类	数值	是
	Item_id         int    `json:"item_id"`         //道具id	数值	是
	Item_num        int    `json:"item_num"`        //数量	数值	是
	Item_value      int    `json:"item_value"`      //动作后的物品存量	数值	是
	Log_event_parm1 string `json:"log_event_parm1"` //道具流动一级原因	字符	是
	Log_event_parm2 string `json:"log_event_parm2"` //道具流动二级原因	字符	是
}

func ItemChange(user *objs.User, itemId, count, afterCount, Reason int, reason2 int, addOrReduce bool) {
	if itemId == pb.ITEMID_GOLD || itemId == pb.ITEMID_INGOT {
		ItemCoinChange(user, itemId, count, afterCount, Reason, reason2, addOrReduce)
	} else {
		itemNormalChange(user, itemId, count, afterCount, Reason, reason2, addOrReduce)
	}
}

func ItemCoinChange(user *objs.User, itemId, count, afterCount, reason int, reason2 int, addOrReduce bool) {
	data := &KyEvent_item_coin{
		KyEventPropBase: getEventBaseProp(user),
		ChangeType:      1,
		Coin_id:         itemId,
		Coin_num:        count,
		Coin_value:      afterCount,
	}
	if !addOrReduce {
		data.ChangeType = 2
	}
	reasonLv1, reasonLv2 := ophelper.GetResaon(reason, reason2)
	data.Log_event_parm1 = reasonLv1
	data.Log_event_parm2 = reasonLv2
	writeUserEvent(user, "change_coin", data)
}

func itemNormalChange(user *objs.User, itemId, count, afterCount, reason, reason2 int, addOrReduce bool) {
	itemConf := gamedb.GetItemBaseCfg(itemId)
	if itemConf == nil {
		logger.Error("道具配置错误：%v", itemId)
		return
	}
	data := &KyEvent_item{
		KyEventPropBase: getEventBaseProp(user),
		ChangeType:      1,
		Main_type:       itemConf.Type,
		Sub_type:        0,
		Item_id:         itemId,
		Item_num:        count,
		Item_value:      afterCount,
	}
	if !addOrReduce {
		data.ChangeType = 2
	}
	reasonLv1, reasonLv2 := ophelper.GetResaon(reason, reason2)
	data.Log_event_parm1 = reasonLv1
	data.Log_event_parm2 = reasonLv2
	writeUserEvent(user, "change_item", data)
}
