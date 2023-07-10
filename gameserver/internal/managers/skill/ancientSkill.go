package skill

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constConstant"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

/**
 *  @Description: 远古神技激活
 *  @param user
 *  @param skillId	技能id
 *  @param op
 *  @return error
 */
func (this *SkillManager) AncientSkillActive(user *objs.User, heroIndex int, op *ophelper.OpBagHelperDefault) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}

	levelCfg := gamedb.GetAncientSkillLevelAncientSkillLevelCfg(constConstant.COMPUTE_TEN_THOUSAND)
	if levelCfg == nil {
		return gamedb.ERRPARAM
	}
	if check := this.GetCondition().CheckMulti(user, heroIndex, levelCfg.Condition); !check {
		return gamedb.ERRCONDITION
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, levelCfg.Cost); err != nil {
		return err
	}

	if hero.AncientSkill == nil {
		hero.AncientSkill = &model.AncientSkill{}
	}
	if hero.AncientSkill.Level > 0 || hero.AncientSkill.Grade > 0 {
		return gamedb.ERRREPEATACTIVE
	}
	hero.AncientSkill = &model.AncientSkill{
		Level: 1,
		Grade: 1,
	}
	user.Dirty = true

	this.GetUserManager().UpdateCombat(user, -1)
	this.GetTask().UpdateTaskProcess(user, false, false)
	this.GetCondition().RecordCondition(user, pb.CONDITION_ACTIVE_YUAN_GU_SKILL_NUM, []int{})
	return nil
}

/**
 *  @Description: 远古神技升级
 *  @param user
 *  @param skillId	技能id
 *  @param op
 *  @return error
 */
func (this *SkillManager) AncientSkillUpLv(user *objs.User, heroIndex int, op *ophelper.OpBagHelperDefault) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	skillLv := hero.AncientSkill.Level
	if gamedb.GetAncientSkillLevelAncientSkillLevelCfg(constConstant.COMPUTE_TEN_THOUSAND+skillLv+1) == nil {
		return gamedb.ERRLVENOUGH
	}
	skillInfoCfg := gamedb.GetAncientSkillLevelAncientSkillLevelCfg(constConstant.COMPUTE_TEN_THOUSAND + skillLv)
	if check := this.GetCondition().CheckMulti(user, heroIndex, skillInfoCfg.Condition); !check {
		return gamedb.ERRCONDITION
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, skillInfoCfg.Cost); err != nil {
		return err
	}

	hero.AncientSkill.Level++
	user.Dirty = true

	this.GetUserManager().UpdateCombat(user, -1)
	return nil
}

/**
 *  @Description: 远古神技升阶
 *  @param user
 *  @param skillId	技能id
 *  @param op
 *  @return error
 */
func (this *SkillManager) AncientSkillUpGrade(user *objs.User, heroIndex int, op *ophelper.OpBagHelperDefault) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	skillGrade := hero.AncientSkill.Grade
	if gamedb.GetAncientSkillGradeAncientSkillGradeCfg(constConstant.COMPUTE_TEN_THOUSAND+skillGrade+1) == nil {
		return gamedb.ERRLVENOUGH
	}
	skillInfoCfg := gamedb.GetAncientSkillGradeAncientSkillGradeCfg(constConstant.COMPUTE_TEN_THOUSAND + skillGrade)
	if check := this.GetCondition().CheckMulti(user, heroIndex, skillInfoCfg.Condition); !check {
		return gamedb.ERRCONDITION
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, skillInfoCfg.Cost); err != nil {
		return err
	}

	hero.AncientSkill.Grade++
	user.Dirty = true

	this.GetUserManager().UpdateCombat(user, -1)
	return nil
}
