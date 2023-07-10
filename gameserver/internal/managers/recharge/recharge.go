package recharge

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constOrder"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"fmt"
	"os"
	"strconv"
	"sync/atomic"
	"time"
)

var orederCounter uint32

type RechargeManager struct {
	util.DefaultModule
	managersI.IModule
}

func NewRechargeManager(module managersI.IModule) *RechargeManager {
	return &RechargeManager{IModule: module}
}

func (this *RechargeManager) Init() error {
	return nil
}

func (this *RechargeManager) Online(user *objs.User) {
	this.DispatchEvent(user.Id, nil, this.updateOrder)
}

func (this *RechargeManager) updateOrder(userId int, user *objs.User, data interface{}) {
	if user == nil {
		return
	}
	orderModel := modelGame.GetOrderModel()
	realOrders, err := orderModel.GetNoRechargeOrder(user.Id)
	if err == nil && len(realOrders) > 0 {
		logger.Info("上线，存在还未同步的订单 userId=%v nickName=%v", user.Id, user.NickName)
		for _, order := range realOrders {
			logger.Info("未同步的订单 userId=%v nickName=%v PayMoney=%v rechargeId=%v", user.Id, user.NickName, order.PayMoney, order.RechargeId)
			order.Synced = constOrder.ORDER_SYNCED_YES
			orderModel.Update(order)
			this.payResultOperation(user.Id, user, order)
		}
	}
	this.RechargeReset(user)
}

func (this *RechargeManager) RechargeReset(user *objs.User) {
	openDay := this.GetSystem().GetServerOpenDaysByServerId(base.Conf.ServerId)
	date := common.GetResetTime(time.Now())
	if openDay%gamedb.GetConf().Recharge == 0 && user.DayStateRecord.RechargeResetTime != date {
		user.Recharge = make(model.IntKv)
		user.DayStateRecord.RechargeResetTime = date
		this.GetUserManager().SendMessage(user, &pb.RechargeResetNtf{Recharge: builder.BuildRecharge(user)}, true)
	}
}

/**
 *  @Description: 生成订单号
 *  @return string
 */
func (this *RechargeManager) makeOrderNo() string {
	formatTime := time.Now().Format("20060102150405")
	pid := os.Getpid()
	orderNum := atomic.AddUint32(&orederCounter, 1)
	orderNo := fmt.Sprintf("%14s%05d%05d", formatTime, pid, orderNum)
	return orderNo
}

// 充值检查
func (this *RechargeManager) checkRechargeNum(id int, payNum int) (int, error) {
	//-1 普通充值检查
	conf := gamedb.GetRechargeRechargeCfg(id)
	if conf != nil && conf.Money == payNum {
		return conf.Money, nil
	}
	return 0, gamedb.ERRRECHARGENUM
}

/**
 *  @Description: 充值支付成功回调
 *  @param user
 *  @param count	支付金额
 *  @param ophelper
 *  @param rechargeId	充值id
 */
func (this *RechargeManager) RechargeResult(user *objs.User, count int, op *ophelper.OpBagHelperDefault, rechargeId int) {
	logger.Info("RechargeResult userId=%v count=%v rechargeId=%d rechargeType:%v", user.Id, count, rechargeId)
	rechargeConf := gamedb.GetRechargeRechargeCfg(rechargeId)
	addIngot := 0
	for _, info := range rechargeConf.Reward1 {
		if info.ItemId == pb.ITEMID_INGOT {
			addIngot += info.Count
		}
	}
	this.GetBag().AddItems(user, rechargeConf.Reward1, op)
	if _, ok := user.Recharge[rechargeId]; !ok {
		this.GetBag().AddItems(user, rechargeConf.Reward2, op)
		user.Recharge[rechargeId] = 0
		for _, info := range rechargeConf.Reward2 {
			if info.ItemId == pb.ITEMID_INGOT {
				addIngot += info.Count
			}
		}
	}
	this.GetUserManager().SendMessage(user, &pb.RechargFulfilledNtf{
		Ingot:        int32(addIngot),
		RechargedAll: int32(user.RechargeAll),
		PayMoney:     int32(count),
		Vip:          int32(user.VipLevel),
		VipExp:       int32(user.VipScore),
		RechargeId:   int32(rechargeId),
	}, true)

	openDay := this.GetSystem().GetServerOpenDaysByServerId(user.ServerId)
	conditionData := this.GetCondition().GetConditionData(user, pb.CONDITION_ALL_RECHARGE_DAY, 0)
	if conditionData/10000 != openDay {
		buyDay := 0
		if conditionData != 0 {
			buyDay, _ = strconv.Atoi(strconv.Itoa(conditionData)[4:])
		}
		this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_RECHARGE_DAY, []int{openDay*10000 + buyDay})
	}
}

////开服检测玩家是否存在未处理的模拟充值订单
//func (this *RechargeManager) CheckSimulationPay(user *objs.User) {
//
//	ophelper := ophelper.NewOpBagHelperDefault(ophelper.OpTypeRecharge)
//	realOrders, err := model.GetOrderModel().GetNoRechargeOrder(user.Id)
//	openday := m.System.GetServerOpenDaysByServerId(user.ServerId)
//	if err == nil && len(realOrders) > 0 {
//		logger.Info("上线，存在还未同步的订单 userId=%v nickName=%v", user.Id, user.NickName)
//		payNum := 0
//		for _, v := range realOrders {
//			logger.Info("未同步的订单 userId=%v nickName=%v PayMoney=%v rechargeId=%v", user.Id, user.NickName, v.PayMoney, v.RechargeId)
//			//payNum += v.PayMoney
//			payNum += v.ShowMoney
//			v.Synced = 1
//			model.GetOrderModel().Update(v)
//			//m.Recharge.RechargeResult(user, v.PayMoney, ophelper, v.RechargeId, v.RechargeType, v.Mi, true)
//			m.Recharge.RechargeResult(user, v.ShowMoney, ophelper, v.RechargeId, v.RechargeType, v.Mi, true)
//		}
//		user.SendMessage(0, &pb.RechargFulfilledNtf{
//			Mi:                    int32(user.Mi),
//			RechargedAll:          int32(user.RechargeAll),
//			AddCount:              int32(payNum),
//			Vip:                   int32(user.VipLevel),
//			VipExp:                int32(user.VipScore),
//			Goods:                 ophelper.ToGoodsChangeMessages(),
//			DayRechargeAwardState: rmodel.User.GetUserDayRechargeAwardState(user.Id, openday),
//			DayMultipleState:      rmodel.User.GetUserDayRechargeMultiple(user.Id, openday)})
//	}
//
//	orders, err := model.GetSimulationModel().GetSimulationOrderByUserId(user.Id)
//	if err != nil {
//		return
//	}
//	if len(orders) < 1 {
//		return
//	}
//
//	for _, v := range orders {
//		//m.Recharge.RechargeResult(user, v.PayMoney, ophelper, v.RechargeId, v.RechargeType, 0, false)
//		m.Recharge.RechargeResult(user, v.ShowMoney, ophelper, v.RechargeId, v.RechargeType, 0, false)
//		user.SendMessage(0, &pb.RechargFulfilledNtf{
//			Mi:           int32(user.Mi),
//			RechargedAll: int32(user.RechargeAll),
//			//AddCount:              int32(v.PayMoney),
//			AddCount:              int32(v.ShowMoney),
//			Vip:                   int32(user.VipLevel),
//			VipExp:                int32(user.VipScore),
//			Goods:                 ophelper.ToGoodsChangeMessages(),
//			DayRechargeAwardState: rmodel.User.GetUserDayRechargeAwardState(user.Id, openday),
//			DayMultipleState:      rmodel.User.GetUserDayRechargeMultiple(user.Id, openday)})
//		model.GetSimulationModel().Upgrade(v.Id)
//		logger.Info("simulation success userId:%d num:%d", user.Id, v.PayMoney)
//	}
//}
