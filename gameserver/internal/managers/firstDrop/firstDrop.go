package firstDrop

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"time"
)

type FirstDropManager struct {
	util.DefaultModule
	managersI.IModule
}

func NewFirstDropManager(module managersI.IModule) *FirstDropManager {
	return &FirstDropManager{
		IModule: module,
	}
}

func (this *FirstDropManager) LoadInfo(user *objs.User, types int, ack *pb.FirstDropLoadAck) {
	allCfg := gamedb.GetFirstDropItemInfoByTypes(types)
	if len(allCfg) <= 0 {
		return
	}
	ack.GetDropItemInfo = make(map[int32]int32)
	ack.AllDropItemGetCount = make(map[int32]int32)
	for _, data := range allCfg {
		id := data.Id
		if user.FirstDropItemInfo[id] >= 1 {
			ack.GetDropItemInfo[int32(id)] = int32(user.FirstDropItemGet[id])
		}
	}
	allGetNums := rmodel.FirstDrop.GetFirstDropAllItemGetNum(types)
	for itemId, num := range allGetNums {
		cfg := gamedb.GetFirstDropFirstDropCfg(itemId)
		if cfg == nil {
			continue
		}
		ack.AllDropItemGetCount[int32(itemId)] = int32(num)
	}
	ack.Types = int32(types)
	return
}

//领奖
func (this *FirstDropManager) GetAward(user *objs.User, id int, ack *pb.GetFirstDropAwardAck, op *ophelper.OpBagHelperDefault) error {
	cfg := gamedb.GetFirstDropFirstDropCfg(id)
	if cfg == nil {
		return gamedb.ERRPARAM
	}

	if user.FirstDropItemInfo[id] <= 0 {
		return gamedb.ERRPARAM
	}

	if user.FirstDropItemGet[id] > 0 {
		return gamedb.ERRHAVEGETREWARD
	}

	this.GetBag().AddItem(user, op, cfg.Reward.ItemId, cfg.Reward.Count)
	user.FirstDropItemGet[id] = 1
	rmodel.FirstDrop.SetFirstDropItemGetNum(cfg.Type, id, 1)
	ack.GetDropItemInfo = make(map[int32]int32)
	ack.DropItemGetCount = make(map[int32]int32)
	ack.GetDropItemInfo[int32(id)] = 1
	num := rmodel.FirstDrop.GetFirstDropItemGetNum(cfg.Type, id)
	if num >= cfg.Count && cfg.Count > 0 {
		ntf := &pb.GetAllFirstDropAwardNtf{}
		ntf.DropItemGetCount = make(map[int32]int32)
		ntf.DropItemGetCount[int32(id)] = int32(num)
		this.BroadcastAll(ntf)
	}
	ack.DropItemGetCount[int32(id)] = int32(num)
	ack.Types = int32(cfg.Type)
	user.Dirty = true
	return nil
}

//一键领奖
func (this *FirstDropManager) GetAllAward(user *objs.User, types int, ack *pb.GetAllFirstDropAwardAck, op *ophelper.OpBagHelperDefault) error {
	allCfg := gamedb.GetFirstDropItemInfoByTypes(types)
	if len(allCfg) <= 0 {
		return gamedb.ERRPARAM
	}
	for _, info := range gamedb.GetConf().FirstDropOpenTime {
		if info[0] == allCfg[0].Type {
			if !this.checkActivityIsOpenByTypes(info[2], info[1]) {
				return gamedb.ERRPARAM
			}
		}
	}

	if len(user.FirstDropItemInfo) <= 0 {
		return gamedb.ERRACTIVITYNOTOPEN
	}

	ntf := &pb.GetAllFirstDropAwardNtf{}
	ack.GetDropItemInfo = make(map[int32]int32)
	ack.DropItemGetCount = make(map[int32]int32)
	ntf.DropItemGetCount = make(map[int32]int32)
	for _, data := range allCfg {
		id := data.Id
		if data.Count > 0 {
			num := rmodel.FirstDrop.GetFirstDropItemGetNum(types, id)
			if num >= data.Count {
				continue
			}
		}
		if user.FirstDropItemInfo[data.Id] >= 1 && user.FirstDropItemGet[id] <= 0 {
			this.GetBag().AddItem(user, op, data.Reward.ItemId, data.Reward.Count)
			user.FirstDropItemGet[id] = 1
			rmodel.FirstDrop.SetFirstDropItemGetNum(types, id, 1)
			ack.GetDropItemInfo[int32(id)] = 1
			num := rmodel.FirstDrop.GetFirstDropItemGetNum(types, id)
			ack.DropItemGetCount[int32(id)] = int32(num)
			if num >= data.Count && data.Count > 0 {
				ntf.DropItemGetCount[int32(id)] = int32(num)
			}
		}
	}
	ack.Types = int32(types)
	if len(ntf.DropItemGetCount) > 0 {
		this.BroadcastAll(ntf)
	}
	if len(ack.GetDropItemInfo) > 0 {
		goods := op.ToChangeItems()
		ack.Goods = goods
	}
	user.Dirty = true
	return nil
}

func (this *FirstDropManager) CheckIsFirstDrop(user *objs.User, itemIds map[int]int) {
	ids := make([]int32, 0)
	for itemId := range itemIds {

		dropCfg := gamedb.GetFirstDropItemInfoByItemId(itemId)
		if dropCfg == nil {
			continue
		}
		if user.FirstDropItemInfo[dropCfg.Id] >= 1 {
			continue
		}

		mark := false
		for _, info := range gamedb.GetConf().FirstDropOpenTime {
			if len(info) < 3 {
				logger.Error("game表 FirstDropOpenTime 配置错误")
				break
			}

			if this.checkActivityIsOpenByTypes(info[2], info[1]) {
				mark = true
			}
		}
		if !mark {
			continue
		}
		user.FirstDropItemInfo[dropCfg.Id] = 1
		ids = append(ids, int32(dropCfg.Id))
	}
	user.Dirty = true
	if len(ids) > 0 {
		ntf := &pb.FirstDropRedPointNtf{}
		ntf.Items = ids
		this.GetUserManager().SendMessage(user, ntf, true)
	}
	return
}

func (this *FirstDropManager) checkActivityIsOpenByTypes(continuesTimes, limitDay int) bool {

	openDay := this.GetSystem().GetServerOpenDaysByServerId(base.Conf.ServerId)
	if openDay < limitDay {
		return false
	}

	if continuesTimes < 0 {
		return true
	}
	serverOpenTime := this.GetSystem().GetServerOpenTimeByServerId(base.Conf.ServerId)
	newServerOpenTime := serverOpenTime.AddDate(0, 0, limitDay-1)
	day := continuesTimes / 86400
	lastSecond := continuesTimes % 86400
	hour := lastSecond / 3600
	lastSecond1 := lastSecond % 3600
	min := lastSecond1 / 60
	openTime := []int{day, hour, min}
	lastTimes := common.CalcEndTime(newServerOpenTime, openTime, 5)
	if time.Now().Unix() > lastTimes {
		return false
	}
	return true
}

//一键领奖红包
func (this *FirstDropManager) GetAllRedPacketAward(user *objs.User, infos []int32, op *ophelper.OpBagHelperDefault) error {
	if len(infos) <= 0 {
		return gamedb.ERRPARAM
	}
	allGetValue := 0

	for i, j := 0, len(infos); i < j; i += 2 {

		itemId := int(infos[i])
		itemNum := int(infos[i+1])

		itemConf := gamedb.GetItemBaseCfg(itemId)
		isEnough, _ := this.GetBag().HasEnough(user, itemId, itemNum)
		if !isEnough {
			continue
		}

		if itemConf.Type == pb.ITEMTYPE_RED_PACKET {
			getValue := itemConf.EffectVal * itemNum
			if itemConf.Type == pb.ITEMTYPE_RED_PACKET {
				vipNum := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_RED_PACKET_GET_NUM)

				if user.RedPacketUseNum+getValue > (vipNum+gamedb.GetConf().RedRecoveryCount) * 100{
					logger.Debug("RedPacketUseNum:%v  getValue:%v vipNum:%v  RedRecoveryCount:%v  RedRate:%v", user.RedPacketUseNum, getValue, vipNum, gamedb.GetConf().RedRecoveryCount, gamedb.GetConf().RedRate)
					continue
				}
				err := this.GetBag().Remove(user, op, itemId, itemNum)
				if err != nil {
					continue
				}
				user.RedPacketNum += getValue
				user.RedPacketUseNum += getValue
			}
			user.RechargeAll += getValue
			allGetValue += getValue
			this.GetBag().Add(user, op, pb.ITEMID_INGOT, common.CeilFloat64(float64(getValue)))
			this.GetBag().Add(user, op, pb.ITEMID_VIP_EXP, common.CeilFloat64(float64(getValue)/100*float64(gamedb.GetConf().Vip)))
		}
	}

	ntf := &pb.UserRechargeNumNtf{
		RechargeNum:  int32(user.RechargeAll),
		RedPacketNum: int32(user.RedPacketNum),
	}
	this.GetUserManager().SendMessage(user, ntf, true)
	this.GetRecharge().ContRechargeWrite(user, allGetValue)
	user.Dirty = true
	return nil
}

func (this *FirstDropManager) Reset(user *objs.User, data int) {

	if user.RedPacketUseNumReset != data {
		user.RedPacketUseNum = 0
		user.RedPacketUseNumReset = data
	}
	user.Dirty = true
	return
}
