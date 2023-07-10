package growFund

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

type GrowFundManager struct {
	util.DefaultModule
	managersI.IModule
}

func NewGrowFundManager(module managersI.IModule) *GrowFundManager {
	return &GrowFundManager{IModule: module}
}

/**
 *  @Description: 成长基金，校验支付金额
 *  @param payNum	支付金额
 *  @return error
 */
func (this *GrowFundManager) GrowFundCheckBuy(user *objs.User, payNum int) error {
	if user.GrowFund.IsBuy {
		return gamedb.ERRREPEATBUY
	}
	payMoney := gamedb.GetConf().GrowFund
	if payMoney != payNum {
		return gamedb.ERRBUYNUM
	}
	return nil
}

func (this *GrowFundManager) GrowFundBuyOperation(user *objs.User, op *ophelper.OpBagHelperDefault) {
	if user.GrowFund.IsBuy {
		return
	}
	growFundGet := gamedb.GetConf().GrowFundGet
	this.GetBag().Add(user, op, growFundGet.ItemId, growFundGet.Count)
	user.GrowFund.IsBuy = true
	user.Dirty = true
	this.GetUserManager().SendMessage(user, &pb.GrowFundBuyAck{IsBuy: true}, true)
}

/**
 *  @Description: 成长基金领取奖励
 *  @param user
 *  @param id	奖励id
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *GrowFundManager) Reward(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.GrowFundRewardAck) error {
	userGrowFund := user.GrowFund
	if !userGrowFund.IsBuy {
		return gamedb.ERRPARAM
	}
	if _, ok := userGrowFund.Ids[id]; ok {
		return gamedb.ERRREPEATRECEIVE
	}

	growFundCfg := gamedb.GetGrowFundGrowFundCfg(id)
	if growFundCfg == nil {
		return gamedb.ERRPARAM
	}
	if check := this.GetCondition().CheckMulti(user, -1, growFundCfg.Condition); !check {
		return gamedb.ERRCONDITION
	}
	this.GetBag().AddItems(user, growFundCfg.Reward, op)
	userGrowFund.Ids[id] = 0
	user.Dirty = true

	ack.Id = int32(id)
	ack.Goods = op.ToChangeItems()
	return nil
}