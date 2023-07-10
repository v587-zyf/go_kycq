package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gameserver/internal/objs"
)

type KyEvent_shop_buy struct {
	*kyEventLog.KyEventPropBase
	ItemId   int         `json:"item_id"`   //商品id
	ShopType int         `json:"main_type"` //商城类型
	BuyType  int         `json:"sub_type"`  //购买类型
	BuyNum   int         `json:"item_num"`   //购买数量
	Discount float64     `json:"num1"`  //折扣
	Consume  map[int]int `json:"info"`   //购买消耗
}

//商场购买
func ShopBuy(user *objs.User, itemId, buyNum, shopT, buyT int, discount float64, consume map[int]int) {
	data := &KyEvent_shop_buy{
		KyEventPropBase: getEventBaseProp(user),
		ItemId:          itemId,
		ShopType:        shopT,
		BuyType:         buyT,
		BuyNum:          buyNum,
		Discount:        discount,
		Consume:         consume,
	}
	writeUserEvent(user, "shop_buy", data)
}
