package vipBoss

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

/**
 *  @Description: vipBoss进入战斗
 *  @param user
 *  @param stageId
 *  @return error
 */
func (this *VipBossManager) EnterVipBossFight(user *objs.User, stageId int) error {
	_, _, err := this.CheckFight(user, stageId)
	if err != nil {
		return err
	}

	//userVipBoss.DareNum[stageId]++
	//user.Dirty = true

	if !this.GetFight().CheckInFightBefore(user, stageId) {
		return gamedb.ERRUSERINFIGHT
	}

	fightId, err := this.GetFight().CreateFight(stageId, nil)
	if err != nil {
		return err
	}
	err = this.GetFight().EnterFightByFightId(user, stageId, fightId)
	if err != nil {
		return err
	}
	return nil
}

/**
 *  @Description: vipBoss战斗回调
 *  @param user
 *  @param isWin	输赢
 *  @param stageId
 *  @param op
 */
func (this *VipBossManager) VipBossFightResultNtf(user *objs.User, isWin bool, stageId int, items map[int]int) {
	ntf := &pb.VipBossFightResultNtf{
		StageId: int32(stageId),
		Result:  pb.RESULTFLAG_FAIL,
	}

	if isWin {
		ntf.Result = pb.RESULTFLAG_SUCCESS
		ntf.Goods = ophelper.CreateGoodsChangeNtf(items)
		this.GetFirstDrop().CheckIsFirstDrop(user, items)
	}
	ntf.DareNum = int32(user.VipBosses.DareNum[stageId])
	this.GetUserManager().SendMessage(user, ntf, true)
}

func (this *VipBossManager) KillMonsterChangeDareNum(user *objs.User, stageId int) {
	this.ChangeOperation(user, true, stageId)
	this.GetUserManager().SendMessage(user, &pb.VipBossDareNumNtf{
		StageId: int32(stageId),
		DareNum: int32(user.VipBosses.DareNum[stageId]),
	}, true)
}

/**
 *  @Description: vipBoss扫荡
 *  @param user
 *  @param stageId
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *VipBossManager) VipBossSweep(user *objs.User, stageId int, op *ophelper.OpBagHelperDefault, ack *pb.VipBossSweepAck) error {
	//if privilege := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_COPY_SWEEP); privilege == 0 {
	//	return gamedb.ERRVIPLVNOTENOUGH
	//}
	userVipBoss, vipBossCfg, err := this.CheckFight(user, stageId)
	if err != nil {
		return err
	}
	conVal, ok := vipBossCfg.Condition[pb.CONDITION_USER_VIP_LV]
	if !ok {
		return gamedb.ERRPARAM
	}
	if conVal >= user.VipLevel {
		return gamedb.ERRPARAM
	}

	dropItems, _, err := this.GetStageManager().GetBossDropItem(user, stageId)
	if err != nil {
		return err
	}

	//user.RedPacketItem.PickInfo = pickInfo

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
	this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_KILL_STAGE, []int{stageId, 1})
	this.ChangeOperation(user, true, stageId)

	ack.StageId = int32(stageId)
	ack.DareNum = int32(userVipBoss.DareNum[stageId])
	return nil
}

func (this *VipBossManager) CheckFight(user *objs.User, stageId int) (*model.VipBosses, *gamedb.VipBossVipBossCfg, error) {
	userVipBoss := user.VipBosses
	vipBossCfg := gamedb.GetVipBossByStageId(stageId)
	if vipBossCfg == nil {
		return nil, nil, gamedb.ERRPARAM
	}
	if check := this.GetCondition().CheckMulti(user, -1, vipBossCfg.Condition); !check {
		return userVipBoss, vipBossCfg, gamedb.ERRVIPLVNOTENOUGH
	}
	if user.VipBosses.DareNum[stageId] >= vipBossCfg.MaxNum {
		return userVipBoss, vipBossCfg, gamedb.ERRFIGHTNUMNOTENOUGH
	}
	return userVipBoss, vipBossCfg, nil
}

func (this *VipBossManager) ChangeOperation(user *objs.User, isWin bool, stageId int) {
	user.Dirty = true
	if isWin {
		user.VipBosses.DareNum[stageId]++
		this.GetWarOrder().WriteWarOrderTask(user, pb.WARORDERCONDITION_KILL_MONSTER, []int{1, constFight.FIGHT_TYPE_VIPBOSS})
		this.GetCondition().RecordCondition(user, pb.CONDITION_KILL_BOSS_NUM, []int{1})
	}
	//通知任务系统，挑战一次vipBoss
	this.GetTask().AddTaskProcess(user, pb.CONDITION_CHALLENGE_VIP_BOSS, -1)
	this.GetDailyTask().CompletionOfTask(user, pb.DAILYTASKACTIVITYTYPE_VIP_GE_REN_BOSS, 1)
}
