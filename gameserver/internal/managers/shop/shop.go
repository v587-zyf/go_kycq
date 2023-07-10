package shop

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constShop"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"time"
)

type Shop struct {
	util.DefaultModule
	managersI.IModule
}

func NewShop(module managersI.IModule) *Shop {
	p := &Shop{IModule: module}
	return p
}

func (this *Shop) OnLine(user *objs.User) {
	date := common.GetResetTime(time.Now())
	this.ResetShop(user, date, false)
}

func (this *Shop) ResetShop(user *objs.User, date int, reset bool) {
	userShop := user.Shops
	if userShop.ShopItem == nil {
		for _, v := range pb.SHOPTYPE_ARRAY {
			userShop.ShopItem[v] = make(model.IntKv)
		}
	}
	resetWeek := common.GetYearWeek(time.Now())
	if userShop.ResetTime != resetWeek {
		userShop.ResetTime = resetWeek
		userShop.ShopItem[pb.SHOPTYPE_WEEK_LIMIT] = make(model.IntKv)
		if reset {
			this.GetUserManager().SendMessage(user, &pb.ShopWeekResetNtf{ShopInfo: builder.BuildShopByType(userShop.ShopItem[pb.SHOPTYPE_WEEK_LIMIT])}, true)
		}
	}
	
	//配置表中已经没有的商品 从玩家数据中删除
	for _, v := range user.Shops.ShopItem {
		for id, _ := range v {
			if gamedb.GetShopTypeCfg(id) == nil {
				delete(v, id)
			}
		}
	}
}

/**
 *  @Description: 商城购买
 *  @param user
 *  @param id		购买的配置id
 *  @param buyNum	购买数量
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *Shop) Buy(user *objs.User, id, buyNum int, op *ophelper.OpBagHelperDefault, ack *pb.ShopBuyAck) error {
	if buyNum < 1 {
		return gamedb.ERRPARAM
	}
	if err := this.GetCondition().CheckFunctionOpen(user, pb.FUNCTIONID_SHOP_OPEN); err != nil {
		return err
	}

	cfg := gamedb.GetShopTypeCfg(id)
	if cfg == nil {
		return gamedb.ERRPARAM
	}

	userShop, ok := user.Shops.ShopItem[cfg.Type]
	if !ok {
		user.Shops.ShopItem[cfg.Type] = make(model.IntKv)
		userShop = user.Shops.ShopItem[cfg.Type]
	}
	if userShop[id]+buyNum > cfg.Time && cfg.Time != 0 {
		return gamedb.ERRBUYUPPERLIMIT
	}
	money := 0
	discount := 1.0
	if cfg.BuyType != constShop.FREE_BUY_TYPE {
		buyNum := float64(cfg.Price.Count * buyNum)
		if cfg.Discount != 0 {
			discount = float64(cfg.Discount) / 10
		}
		if cfg.Type == pb.SHOPTYPE_WEEK_LIMIT {
			if vipDiscount := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_SHOP_DISCOUNT); vipDiscount != 0 {
				//discount *= float64(vipDiscount) / 10
				discount = float64(vipDiscount) / 10
			}
		}
		money = common.CeilFloat64(buyNum * discount)
		if err := this.GetBag().Remove(user, op, cfg.Price.ItemId, money); err != nil {
			return err
		}
	}
	if err := this.GetBag().Add(user, op, cfg.Item.ItemId, cfg.Item.Count*buyNum); err != nil {
		return err
	}
	userShop[id] += buyNum
	user.Dirty = true

	kyEvent.ShopBuy(user, cfg.Item.ItemId, buyNum, cfg.Type, cfg.BuyType, discount, map[int]int{cfg.Price.ItemId: money})
	ack.Id = int32(id)
	ack.BuyNum = int32(userShop[id])
	ack.Goods = op.ToChangeItems()
	this.GetWarOrder().WriteWarOrderTask(user, pb.WARORDERCONDITION_BUY_SHOP_GOODS, []int{buyNum, cfg.Type})
	return nil
}
