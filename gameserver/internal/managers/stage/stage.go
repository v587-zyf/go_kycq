package stage

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"time"
)

func NewStageManager(module managersI.IModule) *StageManager {
	return &StageManager{IModule: module}
}

type StageManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *StageManager) Online(user *objs.User) {

	if user.StageId == 0 {
		user.StageId = gamedb.GetConf().StageInitId
		user.StageWave = 0
	}
	user.StageCumulationTime = 0
	user.StageExpCumulationTime = 0
	if !user.OfflineTime.IsZero() {

		user.StageNormalStartTime = user.OfflineTime
		user.StageExpNormalStartTime = user.OfflineTime
	} else {
		user.StageNormalStartTime = time.Now()
		user.StageExpNormalStartTime = time.Now()
	}

	if user.HookMapBag == nil {
		user.HookMapBag = make(model.Bag, 0)
	}

	this.calcHookMapReawrd(user)
}

func (this *StageManager) LeaveStage(user *objs.User) {

	if user.FightId > 0 {
		stageConf := gamedb.GetStageStageCfg(user.FightStageId)
		if stageConf != nil && stageConf.Type == constFight.FIGHT_TYPE_STAGE {
			user.StageCumulationTime = int(time.Now().Unix() - user.StageNormalStartTime.Unix())
			user.StageNormalStartTime = time.Now()
			user.StageExpCumulationTime = int(time.Now().Unix() - user.StageExpNormalStartTime.Unix())
			user.StageExpNormalStartTime = time.Now()
		}
	}
}

func (this *StageManager) StageFightStartReq(user *objs.User) error {

	rewards := this.calcHookMapReawrd(user)
	if rewards != nil && len(rewards) > 0 {
		bagChangeNtf := &pb.StageBagChangeNtf{
			HookupTime: int32(user.HookMapTime),
			IsOnline:   false,
		}
		for _, v := range user.HookMapBag {
			bagChangeNtf.Items = append(bagChangeNtf.Items, &pb.ItemUnit{
				ItemId: int32(v.ItemId), Count: int64(v.Count),
			})
		}

		this.GetUserManager().SendMessage(user, bagChangeNtf, true)
	}

	user.StageNormalStartTime = time.Now()
	user.StageExpNormalStartTime = time.Now()
	//创建战斗
	fightId, err := this.GetFight().CreateFight(user.StageId, nil)
	if err != nil {
		return err
	}

	//玩家进入战斗
	err = this.GetFight().EnterFightByFightId(user, user.StageId, fightId)
	if err != nil {
		return err
	}
	return nil
	//this.GetFight().EnterStage(user)
}

func (this *StageManager) StageFightEndReq(user *objs.User, op *ophelper.OpBagHelperDefault) error {

	stageConf := gamedb.GetStageStageCfg(user.StageId)
	if stageConf == nil {
		return gamedb.ERRSETTINGNOTFOUND
	}

	newWave := user.StageWave + 1

	hookMapConf := gamedb.GetHookMapHookMapCfg(user.StageId)
	if hookMapConf == nil {
		logger.Error("挂机关卡配置错误,关卡：%v", user.StageId)
		return nil
	}
	if newWave > hookMapConf.Num {
		newWave = hookMapConf.Num
	}
	user.StageWave = newWave

	rewards := this.calcHookMapReawrd(user)
	if rewards != nil && len(rewards) > 0 {
		for _, v := range rewards {

			op.BuildItemGetDisplay(v.ItemId, v.Count)
		}

		bagChangeNtf := &pb.StageBagChangeNtf{
			HookupTime: int32(user.HookMapTime),
			IsOnline:   false,
		}
		for _, v := range user.HookMapBag {
			bagChangeNtf.Items = append(bagChangeNtf.Items, &pb.ItemUnit{
				ItemId: int32(v.ItemId), Count: int64(v.Count),
			})
		}

		this.GetUserManager().SendMessage(user, bagChangeNtf, true)
	}

	ntf := &pb.StageFightEndNtf{StageId: int32(user.StageId), Wave: int32(user.StageWave), Goods: op.ToChangeItems(), OnlyUpdate: true, Result: pb.RESULTFLAG_SUCCESS}
	this.GetUserManager().SendMessage(user, ntf, true)

	return nil
}

func (this *StageManager) calcHookMapReawrd(user *objs.User) []*gamedb.ItemInfo {

	hookMapConf := gamedb.GetHookMapHookMapCfg(user.StageId)
	if hookMapConf == nil {
		logger.Error("挂机关卡配置错误,关卡：%v", user.StageId)
		return nil
	}

	//最大挂机时间
	maxMinu := gamedb.GetConf().HookMapMax / 60
	if user.HookMapTime >= maxMinu {
		return nil
	}

	ExpHookTime := user.StageExpCumulationTime + int(time.Now().Unix()-user.StageExpNormalStartTime.Unix())
	ExpHookMin := ExpHookTime / 60

	hookTime := user.StageCumulationTime + int(time.Now().Unix()-user.StageNormalStartTime.Unix())
	hookMin := hookTime / gamedb.GetConf().HookMapDrop

	if ExpHookMin+user.HookMapTime > maxMinu {
		ExpHookMin = maxMinu - user.HookMapTime
		hookMin = ExpHookMin * 60 / gamedb.GetConf().HookMapDrop
	}

	logger.Debug("计算挂机道具奖励，玩家：%v,时间：%v,奖励次数：%v", user.IdName(), hookTime, hookMin)
	itemRewards := make([]*gamedb.ItemInfo, 0)
	if hookMin >= 1 {
		for i := 0; i <= hookMin; i++ {
			reward, err := gamedb.GetDropItems(hookMapConf.Drop)
			if err != nil {
				logger.Error("挂机关卡配置获取错误,关卡：%v，dropId:%v", user.StageId, hookMapConf.Drop)
				return nil
			}
			//item := make(map[int]int, 0)
			//for _, v := range reward {
			//	item[v.ItemId] = v.Count
			//}
			//logger.Info("-------------次数：%v-------------奖励：%v", i, item)
			itemRewards = append(itemRewards, reward...)
		}

		user.StageNormalStartTime = time.Now()
		user.StageCumulationTime = hookTime % 60
		for _, v := range itemRewards {

			itemConf := gamedb.GetItemBaseCfg(v.ItemId)
			hasAdd := false
			if itemConf.CountLimit <= 0 || itemConf.CountLimit > 1 {
				for _, hasItem := range user.HookMapBag {
					if v.ItemId == hasItem.ItemId {
						hasItem.Count += v.Count
						hasAdd = true
						break
					}
				}
			}
			if !hasAdd {
				user.HookMapBag = append(user.HookMapBag, &model.Item{ItemId: v.ItemId, Count: v.Count})
			}
		}
	}

	if ExpHookMin >= 1 {
		expReward := 0
		user.StageExpNormalStartTime = time.Now()
		user.StageExpCumulationTime = ExpHookTime % 60

		expReward = hookMapConf.Name * ExpHookMin
		if privilege := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_HANGUP_EXP); privilege != 0 {
			expReward = common.CalcTenThousand(privilege, expReward)
		}
		user.HookMapTime += ExpHookMin
		itemRewards = append(itemRewards, &gamedb.ItemInfo{ItemId: pb.ITEMID_EXP, Count: expReward})
		hasAdd := false
		for _, hasItem := range user.HookMapBag {
			if hasItem.ItemId == pb.ITEMID_EXP {
				hasItem.Count += expReward
				hasAdd = true
				break
			}
		}
		if !hasAdd {
			user.HookMapBag = append(user.HookMapBag, &model.Item{ItemId: pb.ITEMID_EXP, Count: expReward})
		}
		logger.Debug("计算挂机经验奖励，玩家：%v,时间：%v,获得经验", user.IdName(), ExpHookMin, expReward)
	}

	return itemRewards
}

func (this *StageManager) StartStageBossFight(user *objs.User) error {
	hookMapConf := gamedb.GetHookMapHookMapCfg(user.StageId)
	stageConf := gamedb.GetStageStageCfg(user.StageId)
	if stageConf == nil || hookMapConf == nil {
		logger.Error("玩家申请挂机boss战，配置异常，stage：%v", user.StageId)
		return gamedb.ERRSETTINGNOTFOUND
	}

	if hookMapConf.StageId2 <= 0 {
		logger.Error("玩家申请挂机boss战，已经是最后已关了stage：%v", user.StageId)
		return gamedb.ERRMAXSTAGE
	}

	if user.StageWave < hookMapConf.Num {
		logger.Error("玩家申请挂机boss战，当前关卡未通关")
		return gamedb.ERRSTAGENOTPASS
	}

	if !this.GetFight().CheckInFightBefore(user, hookMapConf.StageId2) {
		return gamedb.ERRUSERINFIGHT
	}

	//创建战斗
	fightId, err := this.GetFight().CreateFight(hookMapConf.StageId2, nil)
	if err != nil {
		return err
	}

	//玩家进入战斗
	err = this.GetFight().EnterFightByFightId(user, hookMapConf.StageId2, fightId)
	if err != nil {
		return err
	}

	return nil

}

func (this *StageManager) StageBossKillResult(user *objs.User, stageId, result int, items map[int]int) {

	goods := ophelper.CreateGoodsChangeNtf(items)
	ntf := &pb.StageFightEndNtf{StageId: int32(user.StageId), Wave: int32(user.StageWave), Goods: goods, OnlyUpdate: false, Result: int32(result)}
	this.GetUserManager().SendMessage(user, ntf, true)
}

func (this *StageManager) StageBossKill(user *objs.User) {
	soloBossConf := gamedb.GetHookMapHookMapCfg(user.StageId)
	logger.Debug("StageBossKillResult userId:%v  user.StageId:%v user.StageId2:%v soloBossConf.StageId2:%v", user.Id, user.StageId, user.StageId2, soloBossConf.StageId2)
	if soloBossConf.StageId2 != user.FightStageId {
		logger.Error("战斗发来玩家挂机boss战斗结束，stage异常,玩家：%v，挂机：%v,挂机boss:%v，战斗服发来：%v", user.IdName(), user.StageId, soloBossConf.Stage_id, soloBossConf.StageId2, user.FightStageId)
		ntf := &pb.StageFightEndNtf{StageId: int32(user.StageId), Wave: int32(user.StageWave), OnlyUpdate: true, Result: pb.RESULTFLAG_SUCCESS}
		this.GetUserManager().SendMessage(user, ntf, true)
		return
	}
	user.StageWave = 0
	user.StageId2 = soloBossConf.StageId2
	if soloBossConf.StageId3 > 0 {
		user.StageId = soloBossConf.StageId3
	}
	//每日任务 完成一次通知
	this.GetDailyTask().CompletionOfTask(user, pb.DAILYTASKACTIVITYTYPE_GUA_JI_BOSS, 1)
	this.GetCondition().RecordCondition(user, pb.CONDITION_KILL_BOSS_NUM, []int{1})
	logger.Debug("user.StageId2 :%v", user.StageId2)
	this.GetTask().UpdateTaskProcess(user, true, false)
	this.GetCondition().RecordCondition(user, pb.CONDITION_KILL_SHOU_LIN, []int{})
	ntf := &pb.StageFightEndNtf{StageId: int32(user.StageId), Wave: int32(user.StageWave), OnlyUpdate: true, Result: pb.RESULTFLAG_SUCCESS}
	this.GetUserManager().SendMessage(user, ntf, true)
}

/**
 *  @Description: 获取boss掉落
 *  @param stageId
 *  @return gamedb.ItemInfos
 *  @return error
 */
func (this *StageManager) GetBossDropItem(user *objs.User, stageId int) (gamedb.ItemInfos, map[int]int, error) {
	stageCfg := gamedb.GetStageStageCfg(stageId)
	if len(stageCfg.Monster_group) > 2 {
		return nil, nil, gamedb.ERRPARAM
	}
	monsterConf := gamedb.GetMonsterMonsterCfg(stageCfg.Monster_group[0][0])
	pickMax := gamedb.GetRedPacketDropMax(this.GetSystem().GetServerOpenDaysByServerId(user.ServerId))
	isFirst := this.GetCondition().GetConditionData(user, pb.CONDITION_ALL_KILL_STAGE, stageId) == 0
	dropItems, _, err := gamedb.GetMonsterDrop(monsterConf.Monsterid, user.RedPacketItem.PickNum, pickMax, user.RedPacketItem.PickInfo, isFirst, stageCfg.Type)
	if err != nil {
		return nil, nil, err
	}

	if len(dropItems) <= 0 {
		logger.Error("申请怪物掉落，随机物品为空,:%+v", monsterConf.Monsterid)
		return nil, nil, gamedb.ERRUNKNOW
	}

	hasRedPacketChange := false
	for _, v := range dropItems {
		itemConf := gamedb.GetItemBaseCfg(int(v.ItemId))
		if itemConf.Type == pb.ITEMTYPE_RED_PACKET {
			hasRedPacketChange = true
			user.RedPacketItem.PickNum += itemConf.EffectVal
			if monsterConf != nil && len(monsterConf.DropSpecial) > 0 {
				user.RedPacketItem.PickInfo[monsterConf.DropSpecial[0]] += 1
			}
		}
	}
	if hasRedPacketChange {
		this.FSSendMessage(user.FightId, user.FightStageId, &pbserver.GsToFsPickRedPacketInfo{
			UserId: int32(user.Id),
			RedPacket: &pbserver.ActorRedPacket{
				PickNum:   int32(user.RedPacketItem.PickNum),
				PickMax:   int32(gamedb.GetRedPacketDropMax(this.GetSystem().GetServerOpenDaysByServerId(user.ServerId))),
				PickInfos: common.ConvertMapIntToInt32(user.RedPacketItem.PickInfo),
			},
		})
		this.GetUserManager().SendMessage(user, &pb.UserRedPacketGetNumNtf{int32(user.RedPacketItem.PickNum)}, true)
	}

	return dropItems, nil, nil
}

func (this *StageManager) GetStageHookMapReward(user *objs.User, op *ophelper.OpBagHelperDefault) error {

	if len(user.HookMapBag) <= 0 {
		return gamedb.ERRPARAM
	}
	items := make([]*gamedb.ItemInfo, len(user.HookMapBag))
	for k, v := range user.HookMapBag {
		items[k] = &gamedb.ItemInfo{ItemId: v.ItemId, Count: v.Count}
	}
	user.HookMapBag = make(model.Bag, 0)
	user.HookMapTime = 0
	this.GetBag().AddItems(user, items, op)
	return nil
}
