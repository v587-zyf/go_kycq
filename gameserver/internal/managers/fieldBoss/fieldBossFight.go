package fieldBoss

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"time"
)

/**
 *  @Description: 野外boss进入战斗
 *  @param user
 *  @param stageId
 *  @return error
 */
func (this *FieldBoss) EnterFieldBossFight(user *objs.User, stageId int) error {
	if stageId < 1 {
		return gamedb.ERRPARAM
	}
	if this.GetSurplusNum(user) <= 0 {
		return gamedb.ERRNOTENOUGHTIMES
	}

	if !user.FieldBoss.FirstReceive {
		logger.Error("玩家未过引导:%v", user.Id)
		return gamedb.ERRCONDITION
	}

	fieldBossCfg := gamedb.GetFieldBossFieldBossCfg(stageId)
	if fieldBossCfg == nil {
		return gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg("fieldBoss stageId:%v", stageId)
	}
	if check := this.GetCondition().CheckMulti(user, -1, fieldBossCfg.Condition); !check {
		return gamedb.ERRCONDITION
	}

	err := this.GetFight().EnterResidentFightByStageId(user, stageId, 0)
	if err != nil {
		return err
	}
	user.FieldBoss.DareNum++
	user.Dirty = true

	this.GetCondition().RecordCondition(user, pb.CONDITION_CHALLENGE_FIELD_LEADER, []int{1})
	this.GetTask().AddTaskProcess(user, pb.CONDITION_CHALLENGE_FIELD_LEADER, -1)
	return nil
}

/**
 *  @Description: 野外boss战斗回调
 *  @param user
 *  @param winUserId	归属奖励用户id
 *  @param stageId
 *  @param op
 *  @return error
 */
func (this *FieldBoss) FieldBossFightEndAck(user *objs.User, winUserId, stageId int, items map[int]int) error {
	ntf := &pb.FieldBossFightResultNtf{
		StageId: int32(stageId),
		Result:  pb.RESULTFLAG_FAIL,
	}
	if winUserId != user.Id {
		op := ophelper.NewOpBagHelperDefault(constBag.OpTypeFieldBossFight)
		stageCfg := gamedb.GetFieldBossFieldBossCfg(stageId)
		this.GetBag().AddItems(user, stageCfg.JoinDrop, op)
		ntf.Goods = op.ToChangeItems()
		this.GetUserManager().SendItemChangeNtf(user, op)
	} else {
		ntf.Result = pb.RESULTFLAG_SUCCESS
		ntf.Goods = ophelper.CreateGoodsChangeNtf(items)
		this.GetCondition().RecordCondition(user, pb.CONDITION_KILL_BOSS_NUM, []int{1})
		this.GetWarOrder().WriteWarOrderTask(user, pb.WARORDERCONDITION_KILL_MONSTER, []int{1, constFight.FIGHT_TYPE_FIELDBOSS})
		this.GetTask().AddTaskProcess(user, pb.CONDITION_KILL_YE_WAI_BOSS, -1)
		//this.GetAnnouncement().FightSendSystemChat(user, items, stageId, pb.SCROLINGTYPE_KILL_GET)
		this.GetFirstDrop().CheckIsFirstDrop(user, items)

	}
	this.GetDailyTask().CompletionOfTask(user, pb.DAILYTASKACTIVITYTYPE_DUO_REN_YE_WAI_BOSS, 1)
	ntf.Winner = this.GetUserManager().BuilderBrieUserInfo(winUserId)
	ntf.DareNum = int32(user.FieldBoss.DareNum)
	this.GetUserManager().SendMessage(user, ntf, true)
	return nil
}

/**
 *  @Description: 用户中途离开
 *  @param user
 *  @param stageId
 */
func (this *FieldBoss) UserLeave(user *objs.User, stageId int) {
	userFieldBoss := user.FieldBoss
	userFieldBoss.CD[stageId] = int(time.Now().Unix())
	user.Dirty = true
}
