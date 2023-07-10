package firstRecharge

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"time"
)

type FirstRechargeManager struct {
	util.DefaultModule
	managersI.IModule
}

func NewFirstRechargeManager(module managersI.IModule) *FirstRechargeManager {
	return &FirstRechargeManager{IModule: module}
}

/**
 *  @Description: 首充支付检查
 *  @param user
 *  @param typeId	首充id
 *  @param payNum	充值金额
 *  @return error
 */
func (this *FirstRechargeManager) FirstRechargePayCheckPay(user *objs.User, typeId, payNum int) error {
	firstRechargeCfg := gamedb.GetFirstRechargTypeFirstRechargTypeCfg(typeId)
	if firstRechargeCfg == nil {
		return gamedb.ERRPARAM
	}
	if payNum != firstRechargeCfg.Money {
		return gamedb.ERRBUYNUM
	}
	userFirstRecharge := user.FirstRecharge
	if userFirstRecharge.IsRecharge && userFirstRecharge.Discount != 0 {
		return gamedb.ERRREPEATBUY
	}
	return nil
}

/**
 *  @Description: 首充支付成功
 *  @param user
 *  @param typeId	首充id
 *  @param op
 */
func (this *FirstRechargeManager) FirstRechargePayOperation(user *objs.User, typeId int, op *ophelper.OpBagHelperDefault) {
	userFirstRecharge := user.FirstRecharge
	if userFirstRecharge.IsRecharge && userFirstRecharge.Discount != 0 {
		return
	}
	rechargeTypeCfg := gamedb.GetFirstRechargTypeFirstRechargTypeCfg(typeId)
	if rechargeTypeCfg == nil {
		logger.Warn("首充配置未找到 userId:%v id:%v", user.Id, typeId)
		return
	}
	this.GetBag().AddItems(user, rechargeTypeCfg.Reward, op)
	this.UpdateFirstRechargeStatus(user)
	if userFirstRecharge.Discount == 0 {
		userFirstRecharge.Discount = 100
	} else {
		this.GetBag().Remove(user, op, userFirstRecharge.Discount, 1)
		this.GetUserManager().SendItemChangeNtf(user, op)
	}
}

/**
 *  @Description: 领取礼包
 *  @param user
 *  @param day	天数
 *  @param op
 *  @return error
 */
func (this *FirstRechargeManager) Reward(user *objs.User, day int, op *ophelper.OpBagHelperDefault) error {
	if day < 1 {
		return gamedb.ERRPARAM
	}
	userFirstRecharge := user.FirstRecharge
	if !userFirstRecharge.IsRecharge {
		return gamedb.ERRRECEIVEAFTERRECHARGING
	}
	if _, ok := userFirstRecharge.Days[day]; ok {
		return gamedb.ERRREPEATRECEIVE
	}

	openTime := time.Unix(int64(userFirstRecharge.OpenDay), 0)
	openDay := common.GetTheDays(openTime)
	if openDay == 1 {
		if openTime.Hour() < common.DAILY_RESET_HOUR_NEW {
			if time.Now().Hour() > 12 {
				openDay = 1
			} else {
				openDay = 2
			}
		}
	}
	if day > openDay {
		return gamedb.ERRREWARDTIME
	}
	cfg := gamedb.GetFirstRechargeFirstRechargCfg(day)
	if cfg == nil {
		return gamedb.ERRPARAM
	}
	this.GetBag().AddItems(user, cfg.Reward, op)

	userFirstRecharge.Days[day] = 0
	user.Dirty = true
	kyEvent.UserFirstRechargeReward(user, cfg.Day, cfg.Day)
	return nil
}

/**
 *  支付调用一次
 *  @Description: 更新首充状态
 *  @param user
 */
func (this *FirstRechargeManager) UpdateFirstRechargeStatus(user *objs.User) {
	if user.FirstRecharge.IsRecharge {
		return
	}
	userFirstRecharge := user.FirstRecharge
	userFirstRecharge.IsRecharge = true
	userFirstRecharge.OpenDay = int(time.Now().Unix())
	user.Dirty = true
	this.GetUserManager().SendMessage(user, &pb.FirstRechargeNtf{IsRecharge: true, OpenDay: int64(userFirstRecharge.OpenDay)}, true)
}
