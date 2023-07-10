package sevenInvestment

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

func NewSevenInvestmentManager(module managersI.IModule) *SevenInvestmentManager {
	return &SevenInvestmentManager{IModule: module}
}

type SevenInvestmentManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *SevenInvestmentManager) Load(user *objs.User, ack *pb.SevenInvestmentLoadAck) {

	ack.ActivateDay = int32(user.SevenInvestment.BuyOpenDay)
	ack.HaveGetIds = this.buildHaveGetIds(user)
	return
}

func (this *SevenInvestmentManager) GetAward(user *objs.User, op *ophelper.OpBagHelperDefault, id int, ack *pb.GetSevenInvestmentAwardAck) error {
	if id <= 0 {
		return gamedb.ERRPARAM
	}
	if user.SevenInvestment.BuyOpenDay <= 0 {
		return gamedb.ERRRECEIVEAFTERRECHARGING
	}

	openDay := this.GetSystem().GetServerOpenDaysByServerId(user.ServerId)
	canGetId := openDay - user.SevenInvestment.BuyOpenDay + 1
	logger.Debug("获取七日奖励  userId:%v  id:%v  openDay:%v  canGetId:%v", user.Id, id, openDay, canGetId)
	if id > canGetId {
		return gamedb.ERRPARAM
	}
	for _, ids := range user.SevenInvestment.GetAwardIds {
		if ids == id {
			return gamedb.ERRGETCONDITIONERR1
		}
	}

	cfg := gamedb.GetSevenDayInvestSevenDayInvestCfg(id)
	if cfg == nil {
		logger.Error("GetSevenDayInvestSevenDayInvestCfg  nil id:%v", id)
		return gamedb.ERRPARAM
	}

	this.GetBag().AddItems(user, cfg.Rewards, op)

	user.SevenInvestment.GetAwardIds = append(user.SevenInvestment.GetAwardIds, id)
	user.Dirty = true
	ack.HaveGetIds = this.buildHaveGetIds(user)
	return nil
}

func (this *SevenInvestmentManager) SevenPayCheck(user *objs.User, payNum int) error {

	if user.SevenInvestment.BuyOpenDay > 0 {
		return gamedb.ERRSEVENBUYERR
	}

	if gamedb.GetConf().InvestCost != payNum {
		return gamedb.ERRPARAM
	}

	return nil
}

func (this *SevenInvestmentManager) SevenPayCallBack(user *objs.User) {
	if user.SevenInvestment.BuyOpenDay > 0 {
		return
	}
	user.SevenInvestment.BuyOpenDay = this.GetSystem().GetServerOpenDaysByServerId(user.ServerId)
	user.Dirty = true
	ack := &pb.SevenInvestmentLoadAck{}
	this.Load(user, ack)
	this.GetUserManager().SendMessage(user, ack, true)
	this.GetAnnouncement().SendSystemChat(user, pb.SCROLINGTYPE_GOU_MAI_QI_RI_TOU_ZHI, -1, -1)
}

func (this *SevenInvestmentManager) buildHaveGetIds(user *objs.User) []int32 {
	ids := make([]int32, 0)
	for _, data := range user.SevenInvestment.GetAwardIds {
		ids = append(ids, int32(data))
	}
	return ids
}
