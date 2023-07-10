package recharge

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constOrder"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"runtime/debug"
	"time"
)

/**
 *  @Description: 购买成功，hgame回调我们
 *  @param w
 *  @param r
 */
func (this *RechargeManager) NotifyBuy(req *pbserver.RechageCcsToGsReq) error {
	defer func() {
		if r := recover(); r != nil {
			stackBytes := debug.Stack()
			logger.Error("panic NotifyBuy:%v,%s", r, stackBytes)
		}
	}()

	gameOrderNo := req.GameOrder
	order, err := modelGame.GetOrderModel().GetOrderByOrderNo(gameOrderNo)
	if err != nil {
		logger.Error("[NOTIFY_BUY_ERROR] 订单异常，err:%v,  gameOrderNo:%s", gameOrderNo, err)
		return gamedb.ERRRECHARGEORDER
	}

	// 订单是否已处理过
	if order.FinishTime.Year() > 2000 {
		logger.Error("[NOTIFY_BUY_ERROR] 订单已处理, userId:%d, gameOrderNo:%s", order.UserId, gameOrderNo)
		return gamedb.ERRRECHARGEHASFINISH
	}

	money := int(req.Money)
	// 普通充值处理
	if order.PayMoney != int(req.Money) {
		logger.Error("[NOTIFY_BUY_ERROR] 订单金额不符: %d,%d, userId:%d, gameOrderNo:%s", order.PayMoney, money, order.UserId, gameOrderNo)
		return gamedb.ERRRECHARGEMONEY
	}

	order.PlatformOrderNo = req.Oid
	order.FinishTime = time.Now()

	//记录跨服组充值额度
	this.DispatchEvent(order.UserId, order, this.payResultOperation)
	logger.Info("[NOTIFY_BUY_SUCCESS] 购买元宝成功, userId:%d, gameOrderNo:%s payMoney:%v", order.UserId, gameOrderNo, money)
	return nil
}

/**
 *  @Description: 支付完成后续操作
 *  @param userId
 *  @param user
 *  @param data
 */
func (this *RechargeManager) payResultOperation(userId int, user *objs.User, data interface{}) {

	order := data.(*modelGame.OrderDb)
	// 更新数据库
	if user != nil {
		order.Synced = constOrder.ORDER_SYNCED_YES
		user.LastRechargeTime = order.FinishTime.Unix()
	} else {
		order.Synced = constOrder.ORDER_SYNCED_NO
	}

	err := modelGame.GetOrderModel().Update(order)
	if err != nil {
		logger.Error("[NOTIFY_BUY_ERROR] 更新订单DB处理结果失败: %s, userId:%d, gameOrderNo:%s", err.Error(), order.UserId, order.OrderNo)
		return
	}

	// 更新玩家数据
	if user != nil {
		logger.Info("Recharge OK.user is online userId=%v %v %v %v %v %v", user.Id, user.NickName, order.PayMoney, order.OrderNo, order.PlatformOrderNo)
		this.SendMsgToCCS(0, &pbserver.SetDayRechargeNumNtf{ServerId: int32(order.ServerId), RechargeNum: int32(order.PayMoney)})
		ophelper := ophelper.NewOpBagHelperDefault(constBag.OpTypeRecharge)
		//m.Recharge.RechargeResult(user, order.PayMoney, ophelper, order.RechargeId, int(order.RechargeType), order.Mi, true)
		//根据订单类型处理
		switch order.PayModule {
		case pb.MONEYPAYTYPE_VIP:
			//暂无vip
		case pb.MONEYPAYTYPE_LIMITED_GIFT:
			//限时礼包
			this.GetGift().LimitedGiftBuyOperation(user, order.PayModuleId, ophelper)
		case pb.MONEYPAYTYPE_WARORDER_LUXURY:
			//豪华战令
			this.GetWarOrder().WarOrderBuyLuxuryOperation(user)
		case pb.MONEYPAYTYPE_RECHARGE:
			//充值
			this.RechargeResult(user, order.PayMoney, ophelper, order.RechargeId)
		case pb.MONEYPAYTYPE_WARORDER_EXP:
			//战令经验
			this.GetWarOrder().WarOrderBuyExpOperation(user)
		case pb.MONEYPAYTYPE_DAILY_PACK:
			//每日、每周礼包
			this.GetDailyPack().DailyPackBuyOperation(user, order.PayModuleId, ophelper)
			kyEvent.UserGiftBuy(user, pb.MONEYPAYTYPE_DAILY_PACK, order.PayModuleId, 0, order.PayMoney, 0)
		case pb.MONEYPAYTYPE_GROW_FUND:
			//成长基金
			this.GetGrowFund().GrowFundBuyOperation(user, ophelper)
		case pb.MONEYPAYTYPE_MONTH_CARD:
			//月卡
			this.GetMonthCard().MonthCardBuyOperation(user, order.PayModuleId, ophelper)
		case pb.MONEYPAYTYPE_TREASURE:
			//寻龙探宝
			this.GetTreasure().PayCallBack(user, order.PayMoney, ophelper)
		case pb.MONEYPAYTYPE_DAILYRANKBUYGIFT:
			//每日排行充值
			this.GetDailyRank().PayCallBack(user, order.PayMoney, order.PayModuleId, ophelper)
			kyEvent.UserGiftBuy(user, pb.MONEYPAYTYPE_DAILYRANKBUYGIFT, order.PayModuleId, 0, order.PayMoney, 0)
		case pb.MONEYPAYTYPE_SEVENINVESTMENT:
			//七日奖励
			this.GetSevenInvestment().SevenPayCallBack(user)
		case pb.MONEYPAYTYPE_OPEN_GIFT:
			//开服礼包
			this.GetGift().OpenGiftBuyOperation(user, order.PayModuleId, ophelper)
			kyEvent.UserGiftBuy(user, pb.MONEYPAYTYPE_OPEN_GIFT, order.PayModuleId, 0, order.PayMoney, 0)
		case pb.MONEYPAYTYPE_FIRST_RECHARGE:
			//首充
			this.GetFirstRecharge().FirstRechargePayOperation(user, order.PayModuleId, ophelper)
		}
		user.RechargeAll += order.PayMoney * 100                                                  //累充
		user.DayStateRecord.DailyRecharge += order.PayMoney                                       //每日充值
		this.GetBag().Add(user, ophelper, pb.ITEMID_VIP_EXP, order.PayMoney*gamedb.GetConf().Vip) //vip
		this.ContRechargeWrite(user, order.PayMoney*100)                                          //连续充值
		this.GetUserManager().SendMessage(user, &pb.UserRechargeNumNtf{
			RechargeNum:   int32(user.RechargeAll),
			RedPacketNum:  int32(user.RedPacketNum),
			DailyRecharge: int32(user.DayStateRecord.DailyRecharge),
		}, true)
		this.GetUserManager().SendItemChangeNtf(user, ophelper)
		user.Dirty = true
		//恺英日志
		kyEvent.UserOrder(user, order)
	} else {
		logger.Info("Recharge OK but user is offline userId=%v %v", order.UserId, order.PayMoney, order.OrderNo, order.PlatformOrderNo)
	}
}
