package spendRebates

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constMail"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

var noOpTArr = map[int]int{
	constBag.OpTypeAuctionPutAwayItem: 0,
	constBag.OpTypeAuctionGotItem:     0,
	constBag.OpTypeAuctionbidding:     0,
}

func NewSpendRebatesManager(m managersI.IModule) *SpendRebatesManager {
	return &SpendRebatesManager{IModule: m}
}

type SpendRebatesManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *SpendRebatesManager) Online(user *objs.User) {
	this.ResetSpendRebate(user, false)
}

func (this *SpendRebatesManager) ResetSpendRebate(user *objs.User, reset bool) {
	userSpendRebates := user.SpendRebates
	openDay := this.GetSystem().GetServerOpenDaysByServerId(user.ServerId)
	spendRebateCfgs := gamedb.GetSpendRebateCfgs()
	if userSpendRebates.Cycle < 1 {
		for _, cfg := range spendRebateCfgs {
			if openDay <= cfg.Time[1] && openDay >= cfg.Time[0] {
				userSpendRebates.Cycle = cfg.Type1
				break
			}
		}
	} else {
		//是否进入下个周期
		spendRebateByType := gamedb.GetSpendRebateByType(userSpendRebates.Cycle)
		if spendRebateByType != nil && openDay > spendRebateByType.Time[1] {
			addMap := make(map[int]int)
			for _, cfg := range spendRebateCfgs {
				err := this.checkReceiveCondition(user, cfg, userSpendRebates)
				if err != nil {
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
				err := this.GetMail().SendSystemMail(user.Id, constMail.MAILTYPE_SPENDREBATES_REWARD, []string{}, bags, 0)
				if err != nil {
					logger.Error("spendRebates sendMail err:%v", err)
				}
			}

			spendRebateByType = gamedb.GetSpendRebateByType(userSpendRebates.Cycle + 1)
			if spendRebateByType != nil {
				userSpendRebates.Cycle = spendRebateByType.Type1
				userSpendRebates.Ingot = 0
			} else {
				userSpendRebates.Cycle = -1
			}
			if reset {
				this.spendRebatesNtf(user, userSpendRebates)
			}
		}
	}
}

/**
 *  @Description: 消费返利领取奖励
 *  @param user
 *  @param id	奖励id
 *  @param op
 *  @return error
 */
func (this *SpendRebatesManager) Reward(user *objs.User, id int, op *ophelper.OpBagHelperDefault) error {
	userSpendRebates := user.SpendRebates

	cfg := gamedb.GetSpendrebatesSpendrebatesCfg(id)
	err := this.checkReceiveCondition(user, cfg, userSpendRebates)
	if err != nil {
		return err
	}

	this.GetBag().AddItems(user, cfg.Reward, op)
	userSpendRebates.Reward[id] = 0
	user.Dirty = true
	return nil
}

func (this *SpendRebatesManager) checkReceiveCondition(user *objs.User, cfg *gamedb.SpendrebatesSpendrebatesCfg, userSpendRebates *model.SpendRebates) error {
	if cfg == nil || cfg.Type1 != userSpendRebates.Cycle {
		return gamedb.ERRPARAM
	}
	if _, ok := userSpendRebates.Reward[cfg.Id]; ok {
		return gamedb.ERRREPEATRECEIVE
	}
	if check := this.GetCondition().CheckMulti(user, -1, cfg.Condition); !check {
		return gamedb.ERRCONDITION
	}
	return nil
}

/**
 *  @Description: 消费返利更新数据
 *  @param user
 *  @param op
 *  @param num
 */
func (this *SpendRebatesManager) UpdateSpendRebates(user *objs.User, op *ophelper.OpBagHelperDefault, num int) {
	userSpendRebates := user.SpendRebates
	opT := op.GetOpType()
	if _, ok := noOpTArr[opT]; !ok {
		userSpendRebates.Ingot += num
		userSpendRebates.CountIngot += num
		this.spendRebatesNtf(user, userSpendRebates)
	}
	user.Dirty = true
}

func (this *SpendRebatesManager) spendRebatesNtf(user *objs.User, userSpendRebates *model.SpendRebates) {
	this.GetUserManager().SendMessage(user, &pb.SpendRebatesNtf{
		CountIngot: int32(userSpendRebates.CountIngot),
		Ingot:      int32(userSpendRebates.Ingot),
		Cycle:      int32(userSpendRebates.Cycle),
	}, true)
}
