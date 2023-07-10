package recharge

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constMail"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"time"
)

func (this *RechargeManager) ContRechargeReset(user *objs.User, reset bool) {
	userContRecharge := user.ContRecharge
	if userContRecharge == nil {
		user.ContRecharge = &model.ContRecharge{Receive: make(model.IntKv), Day: make(model.IntKv)}
		userContRecharge = user.ContRecharge
	}
	contRechargeCfgs := gamedb.GetContRechargeTypes()
	openDay := this.GetSystem().GetServerOpenDaysByServerId(user.ServerId)
	flag := false
	for t, cfg := range contRechargeCfgs {
		if openDay <= cfg.Time[1] && openDay >= cfg.Time[0] {
			flag = true
			if userContRecharge.Cycle != t {
				addMap := make(map[int]int)
				for _, cfg := range gamedb.GetContRechargeByType(t) {
					if err := this.checkReceiveCondition(cfg, userContRecharge); err != nil {
						continue
					}
					for _, info := range cfg.Reward {
						addMap[info.ItemId] += info.Count
					}
				}
				if len(addMap) > 0 {
					bags := make([]*model.Item, 0)
					for itemId, count := range addMap {
						bags = append(bags, &model.Item{
							ItemId: itemId,
							Count:  count,
						})
					}
					err := this.GetMail().SendSystemMail(user.Id, constMail.MAILTYPE_CONTRECHARGE_REWARD, []string{}, bags, 0)
					if err != nil {
						logger.Error("contRecharge sendMail err:%v", err)
					}
				}

				userContRecharge.Cycle = t
				userContRecharge.Day = make(model.IntKv)
				userContRecharge.Receive = make(model.IntKv)
				if reset {
					this.GetUserManager().SendMessage(user, &pb.ContRechargeCycleNtf{Cycle: int32(userContRecharge.Cycle)}, true)
				}
				break
			}
		}
	}
	if !flag {
		userContRecharge.Cycle = -1
		if reset {
			this.GetUserManager().SendMessage(user, &pb.ContRechargeCycleNtf{Cycle: int32(userContRecharge.Cycle)}, true)
		}
	}
}

/**
 *  @Description: 连续充值记录
 *  @param user
 *  @param buyNum	充值金额
 */
func (this *RechargeManager) ContRechargeWrite(user *objs.User, buyNum int) {
	userContRecharge := user.ContRecharge
	day := common.GetResetTime(time.Now())
	userContRecharge.Day[day] += buyNum
	this.GetUserManager().SendMessage(user, &pb.ContRechargeNtf{Recharge: builder.BuildContRechargeRecharge(userContRecharge.Day)}, true)
	user.Dirty = true
}

/**
 *  @Description: 连续充值领取奖励
 *  @param user
 *  @param contRechargeId
 *  @param op
 *  @return error
 */
func (this *RechargeManager) ContRechargeReceive(user *objs.User, contRechargeId int, op *ophelper.OpBagHelperDefault) error {
	userContRecharge := user.ContRecharge

	contRechargeCfg := gamedb.GetContRechargeContRechargeCfg(contRechargeId)
	err := this.checkReceiveCondition(contRechargeCfg, userContRecharge)
	if err != nil {
		return err
	}

	this.GetBag().AddItems(user, contRechargeCfg.Reward, op)
	userContRecharge.Receive[contRechargeId] = 0
	this.GetAnnouncement().SendSystemChat(user, pb.SCROLINGTYPE_LIAN_CHONG_HAO_LI, -1, -1)
	user.Dirty = true
	return nil
}

func (this *RechargeManager) checkReceiveCondition(contRechargeCfg *gamedb.ContRechargeContRechargeCfg, userContRecharge *model.ContRecharge) error {
	if contRechargeCfg == nil || contRechargeCfg.Type != userContRecharge.Cycle {
		return gamedb.ERRPARAM
	}
	okDay := 0
	for _, payNum := range userContRecharge.Day {
		if payNum/100 >= gamedb.GetConf().ContRecharge {
			okDay++
		}
	}
	if okDay < contRechargeCfg.Day {
		return gamedb.ERRNOTENOUGHTIMES
	}
	if _, ok := userContRecharge.Receive[contRechargeCfg.Id]; ok {
		return gamedb.ERRREPEATRECEIVE
	}
	return nil
}
