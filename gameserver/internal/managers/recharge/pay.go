package recharge

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/ptsdk"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constOrder"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gamelibs/rmodelCross"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"fmt"
	"sort"
	"time"
)

/**
 *  @Description: GM命令支付
 *  @param user
 *  @param rechargeId	支付id
 *  @return error
 */
func (this *RechargeManager) TestPay(user *objs.User, rechargeId int) error {
	rechageConf := gamedb.GetRechargeRechargeCfg(rechargeId)
	if rechageConf == nil {
		return gamedb.ERRPARAM
	}
	_, err, _ := this.ApplyPay(user, int32(rechargeId), rechageConf.Money)
	if err != nil {
		logger.Error("模拟充值异常%v", err)
		return err
	}
	//orderNo, err := this.ApplyPay(user, int32(rechargeId), rechageConf.Money)
	//if err != nil {
	//	logger.Error("模拟充值异常%v", err)
	//	return err
	//}
	//order, err := modelGame.GetOrderModel().GetOrderByOrderNo(orderNo)
	//this.DispatchEvent(order.UserId, order, this.payResultOperation)
	return nil
}

/**
 *  @Description: 小伙伴hgame充值相关-发起订单
 *  @param user
 *  @param rechargeId	充值id
 *  @param payNum		金额
 *  @return string		订单号
 *  @return error
 */
func (this *RechargeManager) ApplyPay(user *objs.User, rechargeId int32, payNum int) (string, error, bool) {

	logger.Info("RechargeManager.ApplyPay 玩家：%v, payNum=%v rechargeId=%v ", user.IdName(), payNum, rechargeId)

	paydata, err, isPayToken, _ := this.Pay(user, payNum, pb.MONEYPAYTYPE_RECHARGE, int(rechargeId), false)
	if err != nil {
		return "", err, isPayToken
	}
	return paydata, nil, isPayToken
}

/**
 *  @Description: 统一下单接口
 *  @param user
 *  @param payNum	支付金额
 *  @param payType	支付模块(枚举中的MoneyPayType)
 *  @param typeId	支付类型对应id(如每日礼包,dailyPack表id)
 *  @return string	订单号
 *  @return error
 */
func (this *RechargeManager) Pay(user *objs.User, payNum, payType, typeId int, fromBg bool) (string, error, bool, *modelGame.OrderDb) {
	timeNow := time.Now()
	if timeNow.Unix()-user.RechargeTime.Unix() <= 3 {
		return "", gamedb.ERRPAYTOOFAST, false, nil
	}
	user.RechargeTime = timeNow

	logger.Info("RechargeManager.MoneyPay 玩家：%v, payNum=%v type=%v typeIid=%v,来自平台：%v", user.IdName(), payNum, payType, typeId, fromBg)
	// 生成订单编号
	orderNo := this.makeOrderNo()

	// 生成订单
	var order = &modelGame.OrderDb{
		OpenId:      user.OpenId,
		UserId:      user.Id,
		ServerIndex: user.ServerIndex,
		ServerId:    user.ServerId,
		OrderNo:     orderNo,
		PayMoney:    payNum, // 花费金额(实际支付金额)
		CreateTime:  timeNow,
		PayModule:   payType,
		PayModuleId: typeId,
		OrderDis:    fmt.Sprintf("充值%d元", payNum),
	}

	//特殊订单处理并校验充值金额
	var err error
	switch payType {
	case pb.MONEYPAYTYPE_VIP:
		//VIP 目前没有直购物品
	case pb.MONEYPAYTYPE_LIMITED_GIFT:
		//限时礼包
		order.OrderDis, err = this.GetGift().LimitedGiftCheckBuy(user, typeId, payNum)
	case pb.MONEYPAYTYPE_WARORDER_LUXURY:
		//豪华战令
		err = this.GetWarOrder().WarOrderCheckBuyLuxury(user, payNum)
		order.OrderDis = gamedb.GetConf().WarOrderLuxuryDisplay
	case pb.MONEYPAYTYPE_RECHARGE:
		//充值
		var money int
		money, err = this.checkRechargeNum(int(typeId), payNum)
		if err == nil {
			conf := gamedb.GetRechargeRechargeCfg(typeId)
			order.OrderDis = conf.Display
		}
		order.RechargeId = typeId
		order.Ingot = money * 100
	case pb.MONEYPAYTYPE_WARORDER_EXP:
		//战令经验
		err = this.GetWarOrder().WarOrderCheckBuyExp(user, payNum)
		order.OrderDis = gamedb.GetConf().WarOrderExpBuyDisplay
	case pb.MONEYPAYTYPE_DAILY_PACK:
		//每日礼包
		err = this.GetDailyPack().DailyPackCheckBuy(user, typeId, payNum)
		if err == nil {
			dailyPackCfg := gamedb.GetDailypackDailypackCfg(typeId)
			order.OrderDis = dailyPackCfg.Display
		}
	case pb.MONEYPAYTYPE_GROW_FUND:
		//成长基金
		err = this.GetGrowFund().GrowFundCheckBuy(user, payNum)
		order.OrderDis = gamedb.GetConf().OpenGiftDisplay
	case pb.MONEYPAYTYPE_MONTH_CARD:
		//月卡
		err = this.GetMonthCard().MonthCardCheckBuy(user, typeId, payNum)
		if err == nil {
			monthCardCfg := gamedb.GetMonthCardMonthCardCfg(typeId)
			order.OrderDis = monthCardCfg.Display
		}
	case pb.MONEYPAYTYPE_TREASURE:
		//寻龙探宝  购买寻龙令
		err = this.GetTreasure().PayCheck(user, payNum)
		order.OrderDis = gamedb.GetConf().XunlongDisplay
	case pb.MONEYPAYTYPE_DAILYRANKBUYGIFT:
		//每日排行礼包购买
		err = this.GetDailyRank().PayCheck(user, payNum, typeId)
		if err == nil {
			cfg := gamedb.GetDayRankingGiftDayRankingGiftCfg(typeId)
			order.OrderDis = cfg.Display
		}
	case pb.MONEYPAYTYPE_SEVENINVESTMENT:
		//七日投资购买
		err = this.GetSevenInvestment().SevenPayCheck(user, payNum)
		order.OrderDis = gamedb.GetConf().InvestCostDisplay
	case pb.MONEYPAYTYPE_OPEN_GIFT:
		//开服礼包
		err = this.GetGift().OpenGiftBuyCheck(user, typeId, payNum)
		if err == nil {
			openGiftCfg := gamedb.GetOpenGiftOpenGiftCfg(typeId)
			order.OrderDis = openGiftCfg.Display
		}
	case pb.MONEYPAYTYPE_FIRST_RECHARGE:
		//首充
		err = this.GetFirstRecharge().FirstRechargePayCheckPay(user, typeId, payNum)
		if err == nil {
			firstRechargeCfg := gamedb.GetFirstRechargTypeFirstRechargTypeCfg(typeId)
			order.OrderDis = firstRechargeCfg.Display
		}
	}
	if err != nil {
		return "", err, false, nil
	}

	//支付代币
	isPayToken, isPayTokenFlag := constOrder.PAY_TOKEN_NO, false
	hasPayTokenMap := make(map[int]int)
	hasPayTokenSlice := make([]int, 0)
	hasPayTokenNumMap := make(map[int]int)
	firstChargeDiscountMap := make(map[int]int)

	for _, item := range user.Bag {
		itemId := item.ItemId
		if itemId < 1 {
			continue
		}
		itemBaseCfg := gamedb.GetItemBaseCfg(itemId)
		if itemBaseCfg == nil {
			continue
		}
		if itemBaseCfg.Type == pb.ITEMTYPE_FIRST_RECHARGE_DISCOUNT {
			firstChargeDiscountMap[itemBaseCfg.EffectVal] = item.ItemId
			continue
		}
		if itemBaseCfg.Type != pb.ITEMTYPE_PAY_TOKEN {
			continue
		}
		effectValue := itemBaseCfg.EffectVal
		hasPayTokenMap[effectValue] = itemId
		hasPayTokenNumMap[itemId] = item.Count
		hasPayTokenSlice = append(hasPayTokenSlice, effectValue)
	}

	//首充优惠
	switch payType {
	case pb.MONEYPAYTYPE_FIRST_RECHARGE:
		userFirstRecharge := user.FirstRecharge
		if !userFirstRecharge.IsRecharge && len(firstChargeDiscountMap) > 0 && userFirstRecharge.Discount == 0 {
			firstChargeSlice := make([]int, 0)
			for discount := range firstChargeDiscountMap {
				firstChargeSlice = append(firstChargeSlice, discount)
			}
			sort.Ints(firstChargeSlice)
			if enough, _ := this.GetBag().HasEnough(user, firstChargeDiscountMap[firstChargeSlice[0]], 1); enough {
				payNum = common.CeilFloat64(float64(payNum) * (float64(firstChargeSlice[0]) / 100))
				order.PayMoney = payNum
				userFirstRecharge.Discount = firstChargeDiscountMap[firstChargeSlice[0]]
			}
		}
	}

	if !fromBg && len(hasPayTokenSlice) > 0 {
		sort.Ints(hasPayTokenSlice)
		removeItem := make(gamedb.ItemInfos, 0)
		needPayToken := payNum
		for _, money := range hasPayTokenSlice {
			if needPayToken-money < 0 {
				continue
			}
			itemId := hasPayTokenMap[money]
			needNum := common.FloorFloat64(float64(needPayToken) / float64(money))
			consumeNum := needNum
			if hasPayTokenNumMap[itemId] < needNum {
				consumeNum = hasPayTokenNumMap[itemId]
			}
			removeItem = append(removeItem, &gamedb.ItemInfo{
				ItemId: itemId,
				Count:  consumeNum,
			})
			needPayToken -= money * consumeNum
			if needPayToken == 0 {
				break
			}
			if itemId, ok := hasPayTokenMap[needPayToken]; ok {
				removeItem = append(removeItem, &gamedb.ItemInfo{
					ItemId: itemId,
					Count:  1,
				})
				needPayToken = 0
				break
			}
		}
		if needPayToken == 0 {
			op := ophelper.NewOpBagHelperDefault(constBag.OpTypePayToken)
			if err := this.GetBag().RemoveItemsInfos(user, op, removeItem); err == nil {
				isPayToken = constOrder.PAY_TOKEN_YES
				isPayTokenFlag = true
				order.FinishTime = timeNow
				this.GetUserManager().SendItemChangeNtf(user, op)
			}
		}
	}
	order.IsPayToken = isPayToken

	//保存订单
	err = modelGame.GetOrderModel().Create(order)
	if err != nil {
		return "", err, isPayTokenFlag, order
	}
	kyEvent.UserStartOrder(user, order)
	//使用支付代币，直接执行后续操作
	payData := ""
	if !fromBg {
		if isPayTokenFlag {
			this.DispatchEvent(order.UserId, order, this.payResultOperation)
		} else {
			trailServer := rmodelCross.GetSystemSeting().GetSystemSettingConverInt(rmodelCross.SYSTEM_SETTING_TRIAL_SERVER)
			payData = ptsdk.GetSdk().GetRechargeData(base.Conf.ServerId, user.NickName, user.Heros[constUser.USER_HERO_MAIN_INDEX].ExpLvl, order, trailServer == 1)
			if base.Conf.Sandbox {
				this.DispatchEvent(order.UserId, order, this.payResultOperation)
			}
		}
	}

	return payData, nil, isPayTokenFlag, order
}
