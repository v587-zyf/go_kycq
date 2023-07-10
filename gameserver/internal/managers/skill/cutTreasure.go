package skill

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
	"time"
)

/**
 *  @Description: 切割技能
 *  @param user
 *  @param op
 *  @return error
 */
func (this *SkillManager) CutTreasureUpLv(user *objs.User, op *ophelper.OpBagHelperDefault) error {
	if gamedb.GetCutTreasureByLv(user.CutTreasure+1) == nil {
		return gamedb.ERRLVENOUGH
	}
	cutTreasureCfg := gamedb.GetCutTreasureByLv(user.CutTreasure)
	if !this.GetCondition().CheckMulti(user, constUser.USER_HERO_MAIN_INDEX, cutTreasureCfg.Condition) {
		return gamedb.ERRCONDITION
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, cutTreasureCfg.Item); err != nil {
		return err
	}
	user.CutTreasure++
	user.Dirty = true
	kyEvent.CutTreasureLvUp(user, user.CutTreasure-1, user.CutTreasure)
	this.GetTask().AddTaskProcess(user, pb.CONDITION_UPGRADE_SHENG_DAO_SKILL, -1)
	this.GetCondition().RecordCondition(user, pb.CONDITION_UPGRADE_SHENG_DAO_SKILL, []int{})
	return nil
}

/**
 *  @Description: 返回切割技能(skillLevel表id)
 *  @param user
 *  @return int
 */
func (this *SkillManager) GetCutTreasureSkill(user *objs.User) int {
	hero := user.Heros[constUser.USER_HERO_MAIN_INDEX]
	cutTreasureCfg := gamedb.GetCutTreasureByLv(user.CutTreasure)
	skillId := 0
	switch hero.Job {
	case pb.JOB_ZHANSHI:
		skillId = cutTreasureCfg.SkillZhan
	case pb.JOB_FASHI:
		skillId = cutTreasureCfg.SkillFa
	case pb.JOB_DAOSHI:
		skillId = cutTreasureCfg.SkillDao
	}
	return skillId
}

func (this *SkillManager) CutTreasureSkillUse(user *objs.User) (int32, error) {
	now := time.Now()
	if user.CutTreasureUseEndCd > now.Unix() {
		return 0, gamedb.ERRSKILLCASTBYCD
	}

	stageConf := gamedb.GetStageStageCfg(user.FightStageId)
	if stageConf == nil {
		return 0, gamedb.ERRUNKNOW
	}
	if len(gamedb.GetMaptypeGameCfg(stageConf.Type).CanUseDDQG) > 0 && gamedb.GetMaptypeGameCfg(stageConf.Type).CanUseDDQG[0] != 1 {
		return 0, gamedb.ERRSKILLCANNOTUSE
	}

	cutTreasureConf := gamedb.GetCutTreasureByLv(user.CutTreasure)
	if cutTreasureConf == nil {
		return 0, gamedb.ERRPARAM
	}

	err := this.GetFight().UseCutTreasure(user)
	if err != nil {
		return 0, err
	}
	user.CutTreasureUseEndCd = now.Unix() + int64(cutTreasureConf.CooldownTime)
	user.Dirty = true
	return int32(int(now.Unix()) + cutTreasureConf.CooldownTime), nil
}

func (this *SkillManager) ClearTreasureCd(user *objs.User, stageId int) {

	for _, v := range gamedb.GetConf().ResetCutCD {
		if v == stageId {
			user.CutTreasureUseEndCd = 0
			this.GetUserManager().SendMessage(user, &pb.CutTreasureUseAck{UseTime: int32(0), CdEndTime: int32(0)}, true)
			return
		}
	}

	if this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_CUTTREASURE_AUTO) != 0 {
		stageConf := gamedb.GetStageStageCfg(stageId)
		if stageConf == nil {
			return
		}
		mapTypeConf := gamedb.GetMaptypeGameCfg(stageConf.Type)
		if mapTypeConf == nil {
			return
		}
		if mapTypeConf.ResetCutCD == 1 {
			user.CutTreasureUseEndCd = 0
			this.GetUserManager().SendMessage(user, &pb.CutTreasureUseAck{UseTime: int32(0), CdEndTime: int32(0)}, true)
		}
	}
}
