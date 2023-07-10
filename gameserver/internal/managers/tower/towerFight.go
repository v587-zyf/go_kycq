package tower

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constMax"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"strconv"
)

// 挑战boss
func (this *Tower) EnterTowerFight(user *objs.User) error {
	towerLv := user.Tower.TowerLv
	maxLv := gamedb.GetMaxValById(0, constMax.MAX_TOWER_LEVEL)
	if towerLv > maxLv {
		return gamedb.ERRMAXSTAGE
	}
	towerConf := gamedb.GetTowerTowerCfg(towerLv)
	if towerConf == nil {
		return gamedb.ERRUNKNOW
	}
	//if check := this.GetCondition().CheckMulti(user, -1, towerConf.Condition); !check {
	//	return gamedb.ERRCONDITION
	//}
	if !this.GetFight().CheckInFightBefore(user, towerConf.Stage) {
		return gamedb.ERRUSERINFIGHT
	}

	fightId, err := this.GetFight().CreateFight(towerConf.Stage, nil)
	if err != nil {
		return err
	}
	err = this.GetFight().EnterFightByFightId(user, towerConf.Stage, fightId)
	if err != nil {
		return err
	}
	return nil
}

// 接收客户端 继续战斗
func (this *Tower) TowerFightContinue(user *objs.User) error {
	if user.FightId <= 0 {
		return gamedb.ERRFIGHTID
	}
	towerConf := gamedb.GetTowerTowerCfg(user.Tower.TowerLv)
	if towerConf == nil {
		return gamedb.ERRMAXLV
	}
	return this.EnterTowerFight(user)
	////推送爬塔继续
	//this.FSSendMessage(user.FightId, user.FightStageId, &pbserver.FSContinueFightReq{
	//	StageId: int32(towerConf.Stage),
	//})
}

// 挑战boss后，战斗服务响应结果
func (this *Tower) TowerFightEndAck(user *objs.User, isWin bool, items map[int]int) error {
	if !isWin {
		this.GetUserManager().SendMessage(user, &pb.TowerFightResultNtf{
			Result: 0,
		}, true)
		return nil
	}

	logger.Debug("userId:%v  user.Tower.TowerLv:%v  taskId:%v", user.Id, user.Tower.TowerLv, user.MainLineTask.TaskId)
	towerLv := user.Tower.TowerLv
	if towerLv != 1 {
		towerLv -= 1
	}
	towerConf := gamedb.GetTowerTowerCfg(towerLv)
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeTowerFight)

	this.GetBag().AddItems(user, towerConf.Reward, op)
	this.GetUserManager().SendItemChangeNtf(user, op)
	//if towerConf.RewardLotteryTimes > 0 {
	//	user.Tower.LotteryNum += towerConf.RewardLotteryTimes
	//}
	opGoods := op.ToChangeItems()
	if len(items) > 0 {
		opMsg := ophelper.CreateGoodsChangeNtf(items)
		for _, v := range opMsg.Items {
			opGoods.Items = append(opGoods.Items, v)
		}
		this.GetFirstDrop().CheckIsFirstDrop(user, items)
	}
	this.GetUserManager().SendMessage(user, &pb.TowerFightResultNtf{
		Result:  pb.RESULTFLAG_SUCCESS,
		Goods:   opGoods,
		TowerLv: int32(user.Tower.TowerLv),
	}, true)
	return nil
}

func (this *Tower) KillMonsterChangeDareNum(user *objs.User) {
	this.ChangeOperation(user, 1)
	this.GetUserManager().SendMessage(user, &pb.TowerLvNtf{
		TowerLv: int32(user.Tower.TowerLv),
	}, true)
}

/**
 *  @Description: 通天塔碾压
 *  @param user
 *  @return error
 */
func (this *Tower) TowerSweep(user *objs.User, op *ophelper.OpBagHelperDefault, ack *pb.TowerSweepAck) error {
	towerLv := user.Tower.TowerLv
	maxLv := gamedb.GetMaxValById(0, constMax.MAX_TOWER_LEVEL)
	if towerLv > maxLv {
		return gamedb.ERRMAXSTAGE
	}

	sweepMaxLv := towerLv + gamedb.GetConf().TowSweepMax
	if sweepMaxLv > maxLv {
		sweepMaxLv = maxLv
	}
	okNum, addMap := 0, make(map[int]int)
	for i := towerLv; i <= sweepMaxLv; i++ {
		lv := i
		if lv != 1 {
			lv -= 1
		}
		towerConf := gamedb.GetTowerTowerCfg(i)
		cfgCombat, _ := strconv.Atoi(towerConf.Crush)
		if user.Combat < cfgCombat {
			break
		}
		upperConf := gamedb.GetTowerTowerCfg(lv)
		dropItems, _, err := this.GetStageManager().GetBossDropItem(user, upperConf.Stage)
		if err != nil {
			break
		}
		for _, itemInfo := range dropItems {
			addMap[itemInfo.ItemId] += itemInfo.Count
		}
		for _, itemInfo := range upperConf.Reward {
			addMap[itemInfo.ItemId] += itemInfo.Count
		}
		okNum++
		//user.RedPacketItem.PickInfo = pickInfo
	}
	if okNum == 0 {
		return gamedb.ERRCOMBAT
	}
	addItems := make(gamedb.ItemInfos, 0)
	for itemId, count := range addMap {
		addItems = append(addItems, &gamedb.ItemInfo{
			ItemId: itemId,
			Count:  count,
		})
	}
	this.GetFirstDrop().CheckIsFirstDrop(user, addMap)
	this.GetBag().AddItems(user, addItems, op)
	this.ChangeOperation(user, okNum)
	//this.GetUserManager().SendItemChangeNtf(user, op)

	ack.TowerLv = int32(user.Tower.TowerLv)
	ack.Goods = op.ToChangeItems()
	return nil
}

func (this *Tower) ChangeOperation(user *objs.User, changeNum int) {
	user.Dirty = true
	//主线任务
	user.Tower.TowerLv += changeNum

	//user.Tower.LotteryNum = user.Tower.TowerLv - 1/gamedb.GetConf().LotteryChance
	this.GetTask().AddTaskProcess(user, pb.CONDITION_TOWER, -1)
	//每日任务 完成一次通知
	this.GetDailyTask().CompletionOfTask(user, pb.DAILYTASKACTIVITYTYPE_TONG_TIAN_TA, changeNum)
	//成就
	this.GetCondition().RecordCondition(user, pb.CONDITION_KILL_BOSS_NUM, []int{changeNum})
	//记录到排行榜
	this.GetRank().Append(pb.RANKTYPE_TOWER, user.Id, user.Tower.TowerLv-1, false, false, false)

	this.GetCondition().RecordCondition(user, pb.CONDITION_TOWER, []int{})
}
