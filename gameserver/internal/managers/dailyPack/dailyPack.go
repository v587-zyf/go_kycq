package dailyPack

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"time"
)

func NewDailyPackManager(module managersI.IModule) *DailyPackManager {
	return &DailyPackManager{IModule: module}
}

type DailyPackManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *DailyPackManager) Online(user *objs.User) {
	this.ResetDailyPack(user, false)
}

func (this *DailyPackManager) ResetDailyPack(user *objs.User, reset bool) {
	ntfMap := make(map[int32]int32)
	ntfMap[pb.RESETTYPE_DAILY_PACK_DAY] = pb.RESETTYPE_DAILY_PACK_DAY
	timeNow := time.Now()
	for _, t := range pb.DAILYPACKTYPE_ARRAY {
		dailyPack := user.DailyPack[t]
		if dailyPack == nil {
			user.DailyPack[t] = &model.DailyPackUnit{BuyIds: make(model.IntKv)}
			dailyPack = user.DailyPack[t]
		}
		switch t {
		case pb.DAILYPACKTYPE_DAY:
			date := common.GetResetTime(timeNow)
			if dailyPack.ResetTime != date {
				dailyPack.ResetTime = date
				dailyPack.BuyIds = make(model.IntKv)
			}
		case pb.DAILYPACKTYPE_WEEK:
			resetWeek := common.GetYearWeek(timeNow)
			if dailyPack.ResetWeek != resetWeek {
				dailyPack.ResetWeek = resetWeek
				dailyPack.BuyIds = make(model.IntKv)
				ntfMap[pb.RESETTYPE_DAILY_PACK_WEEK] = pb.RESETTYPE_DAILY_PACK_WEEK
			}
		}
	}
	if reset {
		this.GetUserManager().SendMessage(user, &pb.ResetNtf{Type: ntfMap, NewDayTime: int32(time.Now().Unix())}, true)
	}
}

/**
 *  @Description: 每日礼包购买
 *  @param user
 *  @param id	礼包id
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *DailyPackManager) Buy(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.DailyPackBuyAck) error {
	if id <= 0 {
		return gamedb.ERRPARAM
	}
	dailyPackCfg := gamedb.GetDailypackDailypackCfg(id)
	if dailyPackCfg == nil || dailyPackCfg.Type2 != pb.DAILYPACKBUYTYPE_INGOT {
		return gamedb.ERRPARAM
	}
	userDailyPack := user.DailyPack[dailyPackCfg.Type]
	buyNum, ok := userDailyPack.BuyIds[id]
	if !ok {
		buyNum = 0
	}
	if buyNum >= dailyPackCfg.Condition {
		return gamedb.ERRBUYUPPERLIMIT
	}
	if err := this.GetBag().Remove(user, op, pb.ITEMID_INGOT, dailyPackCfg.Price1); err != nil {
		return err
	}
	this.GetBag().AddItems(user, dailyPackCfg.Reward, op)
	userDailyPack.BuyIds[id]++
	user.Dirty = true

	ack.Id = int32(id)
	ack.BuyNum = int32(userDailyPack.BuyIds[id])
	ack.Goods = op.ToChangeItems()
	kyEvent.UserGiftBuy(user, pb.MONEYPAYTYPE_DAILY_PACK, dailyPackCfg.Id, pb.ITEMID_INGOT, dailyPackCfg.Price1, 0)
	return nil
}

/**
 *  @Description: 每日礼包，校验购买次数，支付金额
 *  @param user
 *  @param typeId	dailyPack表id
 *  @param payNum	支付金额
 *  @return error
 */
func (this *DailyPackManager) DailyPackCheckBuy(user *objs.User, typeId, payNum int) error {
	if typeId <= 0 || payNum <= 0 {
		return gamedb.ERRPARAM
	}
	dailyPackCfg := gamedb.GetDailypackDailypackCfg(typeId)
	if dailyPackCfg == nil || dailyPackCfg.Type2 != pb.DAILYPACKBUYTYPE_MONEY {
		return gamedb.ERRPARAM
	}
	userDailyPack := user.DailyPack[dailyPackCfg.Type]
	buyNum, ok := userDailyPack.BuyIds[typeId]
	if !ok {
		buyNum = 0
	}
	if buyNum >= dailyPackCfg.Condition {
		return gamedb.ERRBUYUPPERLIMIT
	}
	payMoney := dailyPackCfg.Price1
	if payMoney != payNum {
		return gamedb.ERRBUYNUM
	}
	return nil
}

/**
 *  @Description: 每日、每周礼包购买后续操作
 *  @param user
 *  @param payModuleId	dailyPack表id
 *  @param op
 */
func (this *DailyPackManager) DailyPackBuyOperation(user *objs.User, payModuleId int, op *ophelper.OpBagHelperDefault) {
	dailyPackCfg := gamedb.GetDailypackDailypackCfg(payModuleId)
	userDailyPack := user.DailyPack[dailyPackCfg.Type]
	buyNum, ok := userDailyPack.BuyIds[payModuleId]
	if !ok {
		buyNum = 0
	}
	if buyNum >= dailyPackCfg.Condition {
		return
	}
	this.GetBag().AddItems(user, dailyPackCfg.Reward, op)
	userDailyPack.BuyIds[payModuleId]++
	this.GetUserManager().SendMessage(user, &pb.DailyPackBuyAck{Id: int32(payModuleId), BuyNum: int32(buyNum + 1), Goods: op.ToChangeItems()}, true)
}
