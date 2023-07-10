package kyEvent

import (
	"cqserver/gamelibs/eventLog/kyEventLog"
	"cqserver/gameserver/internal/objs"
)

//拍卖行购买
type KyEvent_auction struct {
	*kyEventLog.KyEventPropBase
	SellPeople int `json:"info"`
	BuyPeople  int `json:"buy_people"`
	ItemType   int `json:"main_type"`
	ItemId     int `json:"item_id"`
	ItemCount  int `json:"item_num"`
	Price      int `json:"num1"`
}

//拍卖行上架
type KyEvent_auction_up struct {
	AuctionType int `json:"auction_type"` //1:个人上架 2:系统上架
	UpUser      int `json:"up_user"`      //上架人的userId, -1:世界首领掉落  -2:沙巴克掉落
	UpPrice     int `json:"num1"`         //上架价格
	ItemId      int `json:"item_id"`
	Count       int `json:"item_num"`
}

//拍卖行购买
func Auction(user *objs.User, sellPeople, buyPeople, itemType, itemId, itemCount, price int) {
	data := &KyEvent_auction{
		KyEventPropBase: getEventBaseProp(user),
		SellPeople:      sellPeople,
		BuyPeople:       buyPeople,
		ItemType:        itemType,
		ItemId:          itemId,
		ItemCount:       itemCount,
		Price:           price,
	}
	writeUserEvent(user, "auction", data)
}

//拍卖行上架
func AuctionUp(upUser, upPrice, itemId, itemCount, auctionType int) {
	data := &KyEvent_auction_up{
		UpUser:      upUser,
		UpPrice:     upPrice,
		ItemId:      itemId,
		Count:       itemCount,
		AuctionType: auctionType,
	}
	writeUserEvent(nil, "auction_up", data)
}
