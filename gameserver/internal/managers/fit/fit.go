package fit

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"fmt"
	"time"
)

type FitManager struct {
	util.DefaultModule
	managersI.IModule
}

func NewFitManager(module managersI.IModule) *FitManager {
	f := &FitManager{IModule: module}
	return f
}

/**
 *  @Description: 合体升级
 *  @param user
 *  @param fitId	合体id(目前只有1)
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *FitManager) UpLv(user *objs.User, fitId int, op *ophelper.OpBagHelperDefault, ack *pb.FitUpLvAck) error {
	if fitId < 0 {
		return gamedb.ERRPARAM
	}
	userFitLv := user.Fit.Lv
	lv, ok := userFitLv[fitId]
	if !ok {
		lv = 0
	}
	if gamedb.GetFitLevelFitLevelCfg(gamedb.GetRealId(fitId, lv+1)) == nil {
		return gamedb.ERRLVENOUGH
	}
	lvCfg := gamedb.GetFitLevelFitLevelCfg(gamedb.GetRealId(fitId, lv))
	if !this.GetCondition().CheckMulti(user, -1, lvCfg.Condition) {
		return gamedb.ERRCONDITION
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, lvCfg.Item); err != nil {
		return err
	}
	if lv == 0 {
		user.Fit.Skills[constFight.FIT_SKILL_ZHUDONG_ID] = &model.FitSkill{
			Lv:   1,
			Star: 1,
		}
	}
	userFitLv[fitId]++
	kyEvent.FitLvUp(user, userFitLv[fitId])
	ack.FitId = int32(fitId)
	ack.FitLvId = int32(userFitLv[fitId])

	this.GetUserManager().UpdateCombat(user, -1)
	this.GetTask().UpdateTaskProcess(user, true, false)
	this.GetCondition().RecordCondition(user, pb.CONDITION_FIT_LV, []int{})
	return nil
}

/**
 *  @Description: 合体技能激活
 *  @param user
 *  @param fitSkillId
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *FitManager) SkillActive(user *objs.User, fitSkillId int, op *ophelper.OpBagHelperDefault, ack *pb.FitSkillActiveAck) error {
	if fitSkillId < 1 {
		return gamedb.ERRPARAM
	}
	userFit := user.Fit
	if _, ok := userFit.Skills[fitSkillId]; ok {
		return gamedb.ERRREPEATACTIVE
	}
	skillCfg := gamedb.GetFitSkillFitSkillCfg(fitSkillId)
	if skillCfg == nil {
		return gamedb.ERRSKILLNOTFOUND
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, skillCfg.Item); err != nil {
		return err
	}

	userFit.Skills[fitSkillId] = &model.FitSkill{Star: 1, Lv: 1}
	user.Dirty = true

	ack.FitSkillId = int32(fitSkillId)
	this.GetUserManager().UpdateCombat(user, -1)
	return nil
}

/**
 *  @Description: 合体技能升级
 *  @param user
 *  @param fitSkillId
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *FitManager) SkillUpLv(user *objs.User, fitSkillId int, op *ophelper.OpBagHelperDefault, ack *pb.FitSkillUpLvAck) error {
	if fitSkillId < 1 {
		return gamedb.ERRPARAM
	}
	userFitSkill := user.Fit.Skills
	skill, ok := userFitSkill[fitSkillId]
	if !ok {
		return gamedb.ERRSKILLNOTSTUDY
	}
	if gamedb.GetFitSkillLevelFitSkillLevelCfg(gamedb.GetRealId(fitSkillId, skill.Lv+1)) == nil {
		return gamedb.ERRLVENOUGH
	}
	lvCfg := gamedb.GetFitSkillLevelFitSkillLevelCfg(gamedb.GetRealId(fitSkillId, skill.Lv))
	if lvCfg == nil {
		return gamedb.ERRPARAM
	}
	if check := this.GetCondition().CheckMulti(user, -1, lvCfg.Condition); !check {
		return gamedb.ERRCONDITION
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, lvCfg.Item); err != nil {
		return err
	}
	skill.Lv++
	user.Dirty = true

	ack.FitSkillId = int32(fitSkillId)
	ack.FitSkillLv = int32(skill.Lv)
	kyEvent.FitSkillLvUp(user, fitSkillId, skill.Lv)
	this.IsUpDateCombat(user, fitSkillId)
	return nil
}

/**
 *  @Description: 技能升星
 *  @param user
 *  @param fitSkillId
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *FitManager) SkillUpStar(user *objs.User, fitSkillId int, op *ophelper.OpBagHelperDefault, ack *pb.FitSkillUpStarAck) error {
	if fitSkillId < 1 {
		return gamedb.ERRPARAM
	}
	userFitSkill := user.Fit.Skills
	skill, ok := userFitSkill[fitSkillId]
	if !ok {
		return gamedb.ERRSKILLNOTSTUDY
	}
	if gamedb.GetFitSkillStarFitSkillStarCfg(gamedb.GetRealId(fitSkillId, skill.Star+1)) == nil {
		return gamedb.ERRLVENOUGH
	}
	starCfg := gamedb.GetFitSkillStarFitSkillStarCfg(gamedb.GetRealId(fitSkillId, skill.Star))
	if check := this.GetCondition().CheckMulti(user, -1, starCfg.Condition); !check {
		return gamedb.ERRCONDITION
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, starCfg.Item); err != nil {
		return err
	}
	skill.Star++
	user.Dirty = true

	ack.FitSkillId = int32(fitSkillId)
	ack.FitSkillStar = int32(skill.Star)
	kyEvent.FitSkillStarUp(user, fitSkillId, skill.Star)
	this.IsUpDateCombat(user, fitSkillId)
	return nil
}

/**
 *  @Description: 合体技能替换
 *  @param user
 *  @param skillId
 *  @param slot
 *  @param ack
 *  @return error
 */
func (this *FitManager) SkillChange(user *objs.User, skillId, slot int, ack *pb.FitSkillChangeAck) error {
	if skillId < 1 || slot < 1 {
		return gamedb.ERRPARAM
	}
	userFit := user.Fit
	skills := userFit.Skills
	skillBag := userFit.SkillBag
	if _, ok := skills[skillId]; !ok {
		return gamedb.ERRSKILLNOTSTUDY
	}

	skillCfg := gamedb.GetFitSkillFitSkillCfg(skillId)
	if skillCfg == nil {
		return gamedb.ERRPARAM
	}
	if skillCfg.Type == pb.FITSKILLTYPE_ZHUDONG {
		return gamedb.ERRSKILLNOTWEAR
	}

	slotCfg := gamedb.GetFitSkillSlotFitSkillSlotCfg(slot)
	if slotCfg == nil {
		return gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg(fmt.Sprintf(`fitSkillSlot slot:%v`, slot))
	}
	if check := this.GetCondition().CheckMulti(user, -1, slotCfg.Condition); !check {
		return gamedb.ERRCONDITION
	}
	oldSkill := skillBag[slot]
	skillBag[slot] = skillId
	kyEvent.FitSkillChange(user, oldSkill, skillId)

	ack.FitSkillId = int32(skillId)
	ack.FitSkillSlot = int32(slot)
	this.GetUserManager().UpdateCombat(user, -1)
	return nil
}

/**
 *  @Description: 合体技能重置
 *  @param user
 *  @param skillId
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *FitManager) SkillReset(user *objs.User, skillId int, op *ophelper.OpBagHelperDefault, ack *pb.FitSkillResetAck) error {
	if skillId < 1 {
		return gamedb.ERRPARAM
	}
	userFit := user.Fit
	skill, ok := userFit.Skills[skillId]
	if !ok {
		return gamedb.ERRSKILLNOTSTUDY
	}
	if skill.Star <= 1 && skill.Lv <= 1 {
		return gamedb.ERRPARAM
	}

	resetConsume := gamedb.GetConf().ResetFitSkill
	if err := this.GetBag().Remove(user, op, resetConsume.ItemId, resetConsume.Count); err != nil {
		return err
	}
	addMap := make(gamedb.ItemInfos, 0)
	for i := skill.Star; i > 1; i-- {
		starCfg := gamedb.GetFitSkillStarFitSkillStarCfg(gamedb.GetRealId(skillId, i-1))
		for _, info := range starCfg.Item {
			addMap = append(addMap, &gamedb.ItemInfo{
				ItemId: info.ItemId,
				Count:  info.Count,
			})
		}
	}
	for i := skill.Lv; i > 1; i-- {
		starCfg := gamedb.GetFitSkillLevelFitSkillLevelCfg(gamedb.GetRealId(skillId, i-1))
		for _, info := range starCfg.Item {
			addMap = append(addMap, &gamedb.ItemInfo{
				ItemId: info.ItemId,
				Count:  info.Count,
			})
		}
	}
	if len(addMap) <= 0 {
		return gamedb.ERRRESETSKILL
	}
	this.GetBag().AddItems(user, addMap, op)
	userFit.Skills[skillId] = &model.FitSkill{Star: 1, Lv: 1}

	ack.FitSkillId = int32(skillId)
	ack.FitSkillLv = 1
	ack.FitSkillStar = 1
	this.GetUserManager().UpdateCombat(user, -1)
	return nil
}

/**
 *  @Description: 合体时装升级
 *  @param user
 *  @param fashionId
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *FitManager) FashionUpLv(user *objs.User, fashionId int, op *ophelper.OpBagHelperDefault, ack *pb.FitFashionUpLvAck) error {
	if fashionId < 1 {
		return gamedb.ERRPARAM
	}
	userFit := user.Fit
	lv, ok := userFit.Fashion[fashionId]
	if !ok {
		lv = 0
	}
	if gamedb.GetFitFashionLevelFitFashionLevelCfg(gamedb.GetRealId(fashionId, lv+1)) == nil {
		return gamedb.ERRLVENOUGH
	}
	lvCfg := gamedb.GetFitFashionLevelFitFashionLevelCfg(gamedb.GetRealId(fashionId, lv))
	if err := this.GetBag().RemoveItemsInfos(user, op, lvCfg.Item); err != nil {
		return err
	}
	userFit.Fashion[fashionId]++
	kyEvent.FitFashionLvUp(user, fashionId, userFit.Fashion[fashionId])
	ack.FitFashionId = int32(fashionId)
	ack.FitFashionLv = int32(userFit.Fashion[fashionId])
	this.GetUserManager().UpdateCombat(user, -1)
	return nil
}

/**
 *  @Description: 合体时装穿戴
 *  @param user
 *  @param fashionId
 *  @param ack
 *  @return error
 */
func (this *FitManager) FashionChange(user *objs.User, fashionId int, ack *pb.FitFashionChangeAck) error {
	if fashionId < 1 {
		return gamedb.ERRPARAM
	}
	if _, ok := user.Fit.Fashion[fashionId]; !ok {
		return gamedb.ERRFASHIONNOACTIVE
	}

	userWear := user.Wear
	if userWear.FitFashionId == fashionId {
		userWear.FitFashionId = 0
	} else {
		userWear.FitFashionId = fashionId
	}
	user.Dirty = true

	ack.FitFashionId = int32(userWear.FitFashionId)
	return nil
}

/**
 *  @Description: 进入合体
 *  @param user
 *  @param fitId	合体id(目前只有1)
 *  @return error
 */
func (this *FitManager) EnterFit(user *objs.User, fitId int, ack *pb.FitEnterAck) error {
	if fitId < 1 {
		return gamedb.ERRPARAM
	}
	userFit := user.Fit
	lv, ok := userFit.Lv[fitId]
	if !ok {
		return gamedb.ERRPARAM
	}
	timeNow := int(time.Now().Unix())
	lvCfg := gamedb.GetFitLevelFitLevelCfg(gamedb.GetRealId(fitId, lv))
	if timeNow < userFit.CdEnd {
		return gamedb.ERRFITCD
	}

	err := this.GetFight().UserFitReq(user)
	if err != nil {
		return err
	}
	userFit.CdStart = timeNow
	//userFit.CdEnd = timeNow
	userFit.CdEnd = timeNow + lvCfg.CooldownTime
	user.Dirty = true

	ack.CdStartTime = int32(userFit.CdStart)
	ack.CdEndTime = int32(userFit.CdEnd)
	return nil
}

func (this *FitManager) FitCancel(user *objs.User) error {
	err := this.GetFight().UserFitCacelReq(user)
	if err != nil {
		return err
	}
	return nil
}

func (this *FitManager) IsUpDateCombat(user *objs.User, fitSkillId int) {
	flag := false
	for _, skillId := range user.Fit.SkillBag {
		if fitSkillId == skillId {
			flag = true
			break
		}
	}
	if flag || fitSkillId == constFight.FIT_SKILL_ZHUDONG_ID {
		this.GetUserManager().UpdateCombat(user, -1)
	}
}
