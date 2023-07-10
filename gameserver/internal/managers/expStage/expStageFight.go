package expStage

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/protobuf/pb"
)

/**
 *  @Description: 经验副本进入战斗
 *  @param user
 *  @param stageId
 *  @return error
 */
func (this *ExpStageManager) EnterExpStageFight(user *objs.User, stageId int) error {
	if stageId < 1 {
		return gamedb.ERRPARAM
	}
	if this.GetSurplusNum(user) <= 0 {
		return gamedb.ERRFIGHTNUMNOTENOUGH
	}
	conf := gamedb.GetExpStageByStageId(stageId)
	if conf == nil || conf.Layer-user.ExpStage.Layer > 1 {
		return gamedb.ERRPARAM
	}
	if check := this.GetCondition().CheckMulti(user, -1, conf.Condition); !check {
		return gamedb.ERRCONDITION
	}

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
 *  @Description: 经验副本战斗结束回调
 *  @param user
 *  @param monsterNum	击杀怪物数量
 *  @param getExp		获取经验
 *  @param stageId  	关卡id
 */
func (this *ExpStageManager) ExpStageFightResultNtf(user *objs.User, monsterNum, getExp, stageId int) {
	grade := 0
	expStageCfg := gamedb.GetExpStageByStageId(stageId)
	appraiseIndex := 0
	for i, arr := range expStageCfg.Killandappraise {
		if monsterNum < arr[0] {
			break
		}
		grade = arr[1]
		appraiseIndex = i
	}
	getExp += expStageCfg.Appraiseexp[appraiseIndex]
	if monthCardPrivilege := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_EXPSTAGE); monthCardPrivilege != 0 {
		getExp = common.CalcTenThousand(monthCardPrivilege, getExp)
	}
	userExpStage := user.ExpStage
	userExpStage.ExpStages[stageId] = getExp
	userExpStage.Appraise[stageId] = grade
	if userExpStage.Layer < expStageCfg.Layer {
		userExpStage.Layer = expStageCfg.Layer
	}

	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeExpStage)
	this.GetBag().Add(user, op, pb.ITEMID_EXP, getExp)
	this.GetUserManager().SendItemChangeNtf(user, op)

	ntf := &pb.ExpStageFightResultNtf{
		StageId:    int32(stageId),
		Exp:        int64(getExp),
		MonsterNum: int32(monsterNum),
		Grade:      int32(grade),
		Layer:      int32(userExpStage.Layer),
		IsFree:     this.getKillExpStageNum(user) < 1,
	}
	this.ChangeOperation(user, stageId)
	this.ExpStageDareNumNtf(user)
	this.GetUserManager().SendMessage(user, ntf, true)
}

/**
 *  @Description:经验副本获取经验后减少次数
 *  @param user
 */
func (this *ExpStageManager) ExpStageDareNumNtf(user *objs.User) {
	user.ExpStage.DareNum++
	user.Dirty = true
	//var ItemInfos gamedb.ItemInfos
	//layerCfgMap := gamedb.GetExpStageLayers()
	//layerSlice := make([]int, 0)
	//for _, cfg := range layerCfgMap {
	//	layerSlice = append(layerSlice, cfg.Layer)
	//}
	//sort.Ints(layerSlice)
	//for _, layer := range layerSlice {
	//	if check := this.GetCondition().CheckMulti(user, -1, layerCfgMap[layer].Condition); !check {
	//		break
	//	}
	//	ItemInfos = layerCfgMap[layer].Consume
	//}
	//op := ophelper.NewOpBagHelperDefault(constBag.OpTypeExpStage)
	//this.GetBag().RemoveItemsInfos(user, op, ItemInfos)
	//this.GetUserManager().SendItemChangeNtf(user, op)

	this.GetUserManager().SendMessage(user, &pb.ExpStageDareNumNtf{
		DareNum: int32(user.ExpStage.DareNum),
	}, true)
}

/**
 *  @Description: 经验副本扫荡
 *  @param user
 *  @param stageId	关卡id
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *ExpStageManager) ExpStageSweep(user *objs.User, stageId int, op *ophelper.OpBagHelperDefault, ack *pb.ExpStageSweepAck) error {
	userExpStage := user.ExpStage

	if this.GetSurplusNum(user) <= 0 {
		return gamedb.ERRFIGHTNUMNOTENOUGH
	}
	expStageCfg := gamedb.GetExpStageByStageId(stageId)
	if expStageCfg == nil || expStageCfg.Layer > userExpStage.Layer {
		return gamedb.ERRPARAM
	}
	if _, ok := userExpStage.ExpStages[stageId]; !ok {
		return gamedb.ERRPARAM
	}
	if appraise, ok := userExpStage.Appraise[stageId]; !ok || appraise < APPRAOSE_SS {
		return gamedb.ERRPARAM
	}
	//if err := this.GetBag().RemoveItemsInfos(user, op, expStageCfg.Consume); err != nil {
	//	return gamedb.ERRNOTENOUGHGOODS
	//}
	//怪物掉落
	stageCfg := gamedb.GetStageStageCfg(stageId)
	killMonster := 0
	addExp := 0
	for _, monsterInfo := range stageCfg.Monster_group {
		monsterGroupCfg := gamedb.GetMonstergroupMonstergroupCfg(monsterInfo[0])
		randIndex := 0
		if len(monsterGroupCfg.Monsterid) > 1 {
			weightSlice := make([]int, len(monsterGroupCfg.Monsterid))
			for i, w := range monsterGroupCfg.Monsterid {
				weightSlice[i] = w[2]
			}
			randIndex = common.RandWeightByIntSlice(weightSlice)
		}
		monsterId, monsterNum := monsterGroupCfg.Monsterid[randIndex][0], monsterGroupCfg.Monsterid[randIndex][1]
		killMonster += monsterNum
		monsterCfg := gamedb.GetMonsterMonsterCfg(monsterId)
		dropItems, err := gamedb.GetDropItems(monsterCfg.DropId)
		if err != nil {
			continue
		}
		for _, itemInfo := range dropItems {
			addExp += itemInfo.Count * monsterNum
		}
	}
	appraiseIndex := 0
	for i, arr := range expStageCfg.Killandappraise {
		if killMonster < arr[0] {
			break
		}
		appraiseIndex = i
	}
	//关卡掉落
	addExp += expStageCfg.Appraiseexp[appraiseIndex]
	if monthCardPrivilege := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_EXPSTAGE); monthCardPrivilege != 0 {
		addExp = common.CalcTenThousand(monthCardPrivilege, addExp)
	}
	this.GetBag().Add(user, op, pb.ITEMID_EXP, addExp)

	userExpStage.DareNum++
	userExpStage.ExpStages[stageId] = addExp
	this.ChangeOperation(user, stageId)

	ack.StageId = int32(stageId)
	ack.Exp = int64(addExp)
	ack.MonsterNum = int32(expStageCfg.Killandappraise[len(expStageCfg.Killandappraise)-1][0])
	ack.Grade = int32(expStageCfg.Killandappraise[len(expStageCfg.Killandappraise)-1][1])
	ack.DareNum = int32(userExpStage.DareNum)
	return nil
}

func (this *ExpStageManager) ChangeOperation(user *objs.User, stageId int) {
	user.Dirty = true
	this.GetDailyTask().CompletionOfTask(user, pb.DAILYTASKACTIVITYTYPE_JING_YAN_FU_BEN, 1)
	this.GetCondition().RecordCondition(user, pb.CONDITION_CHALLENGE_JIN_YAN_FU_BEN, []int{1})
	this.GetTask().AddTaskProcess(user, pb.CONDITION_CHALLENGE_JIN_YAN_FU_BEN, -1)

	this.GetWarOrder().WriteWarOrderTask(user, pb.WARORDERCONDITION_COMPLETE_COPY, []int{1, constFight.FIGHT_TYPE_EXPBOSS})
	this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_FINISH_COPY, []int{1})
	this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_KILL_STAGE, []int{stageId, 1})
	this.GetKillMonster().WriteKillMonster(user, stageId)
}
