package rein

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"time"
)

func NewReinManager(module managersI.IModule) *ReinManager {
	return &ReinManager{IModule: module}
}

type ReinManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *ReinManager) OnLine(user *objs.User) {
	// 重置购买次数
	this.ResetReinCostBuyNum(user)
}

func (this *ReinManager) ResetReinCostBuyNum(user *objs.User) {
	reinCosts := user.ReinCosts
	pbReinCost := make([]*pb.ReinCost, 0)
	fiveTime := common.GetDateIntByOffset5(time.Now())

	isUpdate := false
	for _, reinCost := range reinCosts {
		if fiveTime != reinCost.Date {
			isUpdate = true
			reinCost.Num = 0
			reinCost.Date = fiveTime
		}
		pbReinCost = append(pbReinCost, &pb.ReinCost{
			Id:  int32(reinCost.Id),
			Num: int32(reinCost.Num),
		})
	}
	if isUpdate {
		this.GetUserManager().SendMessage(user, &pb.ReinCostBuyNumRefNtf{ReinCost: pbReinCost}, true)
	}
}

func (this *ReinManager) ReinActive(user *objs.User, ack *pb.ReinActiveAck) error {
	if this.GetCondition().CheckMulti(user, -1, gamedb.GetFunctionFunctionCfg(pb.FUNCTIONTYPE_REIN).Condition) {
		user.Rein = &model.Rein{
			Id: 1,
		}
	} else {
		return gamedb.ERRPLAYERLVNOTENOUGH
	}
	ack.Rein = builder.BuilderRein(user.Rein)
	this.GetUserManager().UpdateCombat(user, -1)
	return nil
}

func (this *ReinManager) Reincarnation(user *objs.User, ack *pb.ReincarnationAck) error {
	rein := user.Rein
	id := rein.Id
	if id == 0 {
		return gamedb.ERRREINNOTACTIVE
	}
	if gamedb.GetReinCfg(id+1) == nil {
		return gamedb.ERRLVENOUGH
	}
	conf := gamedb.GetReinCfg(id)
	if rein.Exp < conf.Exp {
		return gamedb.ERRCULTIVATIONNOTENOUGH
	}
	rein.Id += 1
	rein.Exp -= conf.Exp
	ack.Rein = builder.BuilderRein(rein)
	this.GetUserManager().UpdateCombat(user, -1)
	return nil
}

func (this *ReinManager) ReinCostBuy(user *objs.User, id, num int, use bool, op *ophelper.OpBagHelperDefault, ack *pb.ReinCostBuyAck) error {
	rein := user.Rein
	if rein.Id == 0 {
		return gamedb.ERRREINNOTACTIVE
	}
	if user.ReinCosts[id] == nil {
		user.ReinCosts[id] = &model.ReinCost{Id: id}
	}
	reinCost := user.ReinCosts[id]

	conf := gamedb.GetReinCostCfg(id)
	if conf == nil {
		return gamedb.ERRSETTINGNOTFOUND
	}
	if reinCost.Num+num > conf.Number {
		return gamedb.ERRPURCHASECAPENOUGH
	}
	costItemId, costCount := conf.Cost.ItemId, conf.Cost.Count*num
	rewardItemId, rewardCount := conf.Reward.ItemId, conf.Reward.Count*num
	err := this.GetBag().Remove(user, op, costItemId, costCount)
	if err != nil {
		return gamedb.ERRNOTENOUGHGOODS
	}
	reinCost.Num += num
	reinCost.Date = common.GetDateIntByOffset5(time.Now())
	if use == true {
		rein.Exp += gamedb.GetItemBaseCfg(rewardItemId).EffectVal * rewardCount
	} else {
		err = this.GetBag().Add(user, op, rewardItemId, rewardCount)
		if err != nil {
			return err
		}
	}
	ack.Rein = builder.BuilderRein(rein)
	ack.ReinCost = builder.BuilderReinCost(reinCost)
	return nil
}

func (this *ReinManager) ReinCostUse(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.ReinCostUseAck) error {
	rein := user.Rein
	if rein.Id == 0 {
		return gamedb.ERRREINNOTACTIVE
	}
	reinCostConf := gamedb.GetReinCostCfg(id)
	if reinCostConf == nil {
		return gamedb.ERRPARAM
	}
	itemId, itemCount := reinCostConf.Reward.ItemId, 1
	err := this.GetBag().Remove(user, op, itemId, itemCount)
	if err != nil {
		return err
	}
	rein.Exp += gamedb.GetItemBaseCfg(itemId).EffectVal * itemCount

	ack.Rein = builder.BuilderRein(rein)
	if user.ReinCosts[id] != nil {
		ack.ReinCost = builder.BuilderReinCost(user.ReinCosts[id])
	}
	return nil
}
