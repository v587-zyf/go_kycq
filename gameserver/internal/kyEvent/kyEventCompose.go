package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/objs"
)

type KyEvent_Compose struct {
	*kyEventLog.KyEventPropBase
	Item_id      int         `json:"item_id"`    //合成后道具
	Item_count   int         `json:"item_num"` //合成数量
	Consume_item map[int]int `json:"info"`    //合成消耗道具
}

func Compose(user *objs.User, itemId, composeNum int, consume gamedb.ItemInfos) {
	consumeMap := make(map[int]int)
	for _, info := range consume {
		consumeMap[info.ItemId] += info.Count
	}
	data := &KyEvent_Compose{
		KyEventPropBase: getEventBaseProp(user),
		Item_id:         itemId,
		Item_count:      composeNum,
		Consume_item:    consumeMap,
	}
	writeUserEvent(user, "compose", data)
}
