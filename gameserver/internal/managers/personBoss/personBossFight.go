package personBoss

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

/**
 *  @Description: 个人boss进入战斗
 *  @param user
 *  @param stageId
 *  @return error
 */
func (this *PersonBoss) EnterPersonBossFightReq(user *objs.User, stageId int, ack *pb.EnterPersonBossFightAck) error {
	_, hasFightNum, err := this.CheckFight(user, stageId)
	if err != nil {
		return err
	}
	fightId, err := this.GetFight().CreateFight(stageId, nil)
	if err != nil {
		return err
	}
	err = this.GetFight().EnterFightByFightId(user, stageId, fightId)
	if err != nil {
		return err
	}

	ack.StageId = int32(stageId)
	ack.DareNum = int32(hasFightNum)
	return nil
}

/**
 *  @Description: 个人boss战斗回调
 *  @param user
 *  @param stageId
 *  @param isWin	输赢
 *  @param op
 */
func (this *PersonBoss) PersonBossFightResultNtf(user *objs.User, stageId int, isWin bool, items map[int]int) {
	ntf := &pb.PersonBossFightResultNtf{
		StageId: int32(stageId),
		Result:  pb.RESULTFLAG_FAIL,
	}

	if isWin {
		if n, ok := user.PersonBosses.DareNum[stageId]; !ok || n == 0 {
		//	op := ophelper.NewOpBagHelperDefault(constBag.OpTypePersonBossDare)
		//	addItems := make(gamedb.ItemInfos, 0)
		//	for itemId, count := range items {
		//		addItems = append(addItems, &gamedb.ItemInfo{
		//			ItemId: itemId,
		//			Count:  count,
		//		})
		//		items[itemId] += count
		//	}
		//	this.GetBag().AddItems(user, addItems, op)
		//	this.GetUserManager().SendItemChangeNtf(user, op)
			user.PersonBosses.DareNum[stageId] = 1
		}
		ntf.Result = pb.RESULTFLAG_SUCCESS
		ntf.Goods = ophelper.CreateGoodsChangeNtf(items)
		this.GetFirstDrop().CheckIsFirstDrop(user, items)
	}
	ntf.DareNum = int32(this.GetBossKillNum(user.Id, stageId))
	this.GetUserManager().SendMessage(user, ntf, true)
}

func (this *PersonBoss) KillMonsterChangeDareNum(user *objs.User, stageId, hasFightNum int) {
	if _, ok := user.PersonBosses.DareNum[stageId]; !ok {
		user.PersonBosses.DareNum[stageId] = 0
	}
	this.ChangeOperation(user, true, stageId, hasFightNum)
	this.GetUserManager().SendMessage(user, &pb.PersonBossDareNumNtf{
		StageId: int32(stageId),
		DareNum: int32(hasFightNum + 1),
	}, true)
}

/**
 *  @Description: 个人boss扫荡
 *  @param user
 *  @param stageId
 *  @param op
 *  @return error
 */
func (this *PersonBoss) PersonBossSweep(user *objs.User, stageId int, op *ophelper.OpBagHelperDefault, ack *pb.PersonBossSweepAck) error {
	if _, ok := user.PersonBosses.DareNum[stageId]; !ok {
		return gamedb.ERRPARAM
	}
	_, hasFightNum, err := this.CheckFight(user, stageId)
	if err != nil {
		return err
	}
	dropItems, _, err := this.GetStageManager().GetBossDropItem(user, stageId)
	if err != nil {
		return err
	}

	this.GetBag().AddItems(user, dropItems, op)
	isBagFull := !this.GetBag().CheckHasEnoughPos(user, dropItems)
	items := make(map[int]int)
	for _, dropItem := range dropItems {
		items[dropItem.ItemId] += dropItem.Count
	}
	if isBagFull {
		ack.Goods = ophelper.CreateGoodsChangeNtf(items)
		ack.IsBagFull = pb.ISBAGFULLSTATE_FULL
	} else {
		ack.Goods = op.ToChangeItems()
		ack.IsBagFull = pb.ISBAGFULLSTATE_NOT_FULL
	}
	this.GetFirstDrop().CheckIsFirstDrop(user, items)
	//user.RedPacketItem.PickInfo = pickInfo
	this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_KILL_STAGE, []int{stageId, 1})
	this.ChangeOperation(user, true, stageId, hasFightNum)

	ack.StageId = int32(stageId)
	ack.DareNum = int32(hasFightNum + 1)
	return nil
}

func (this *PersonBoss) CheckFight(user *objs.User, stageId int) (*gamedb.PersonalBossPersonalBossCfg, int, error) {
	personBossCfg := gamedb.GetSingleBossByStage(stageId)
	if check := this.GetCondition().CheckMulti(user, -1, personBossCfg.Condition); !check {
		return personBossCfg, 0, gamedb.ERRCONDITION
	}

	if !this.GetFight().CheckInFightBefore(user, stageId) {
		return personBossCfg, 0, gamedb.ERRUSERINFIGHT
	}

	hasFightNum := this.GetBossKillNum(user.Id, stageId)
	if hasFightNum >= personBossCfg.MaxNum {
		return personBossCfg, 0, gamedb.ERRFIGHTNUMNOTENOUGH
	}
	return personBossCfg, hasFightNum, nil
}

func (this *PersonBoss) ChangeOperation(user *objs.User, isWin bool, stageId, hasFightNum int) {
	user.Dirty = true
	if isWin {
		this.writeBossKillNum(user.Id, stageId, hasFightNum+1)
		this.GetWarOrder().WriteWarOrderTask(user, pb.WARORDERCONDITION_KILL_MONSTER, []int{1, constFight.FIGHT_TYPE_PERSON_BOSS})
		this.GetCondition().RecordCondition(user, pb.CONDITION_KILL_BOSS_NUM, []int{1})
	}

	this.GetTask().AddTaskProcess(user, pb.CONDITION_CHALLENGE_ONE_TIME_GE_REN_BOSS, -1)
	this.GetDailyTask().CompletionOfTask(user, pb.DAILYTASKACTIVITYTYPE_GE_REN_BOSS, 1)
}
