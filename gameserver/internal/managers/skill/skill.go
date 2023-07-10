package skill

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gamelibs/publicCon/constMax"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

type SkillManager struct {
	util.DefaultModule
	managersI.IModule
}

func NewSkillManager(module managersI.IModule) *SkillManager {
	return &SkillManager{IModule: module}
}

/**
 *  @Description: 技能升级
 *  @param user
 *  @param op
 *  @param req
 *  @param ack
 *  @return error
 */
func (this *SkillManager) UpLv(user *objs.User, op *ophelper.OpBagHelperDefault, req *pb.SkillUpLvReq, ack *pb.SkillUpLvAck) error {
	skillId := int(req.SkillId)
	userSkill, _, hero, skillCfg, err := this.GetSkill(user, int(req.HeroIndex), skillId, true)
	if err != nil {
		return err
	}
	if skillCfg == nil {
		return gamedb.ERRSKILLNOTFOUND
	}
	if skillCfg.Job != hero.Job {
		return gamedb.ERRJOB
	}

	skill, ok := userSkill[skillId]
	maxLv := gamedb.GetMaxValById(skillId, constMax.MAX_SKILL_LEVEL)
	if ok && skill.Lv >= maxLv {
		return gamedb.ERRLVENOUGH
	}

	var skillLvCfg *gamedb.SkillLevelSkillCfg
	if !ok {
		skillLvCfg = gamedb.GetSkillLvConf(skillId, 0)
	} else {
		skillLvCfg = gamedb.GetSkillLvConf(skillId, skill.Lv)
	}
	if skillLvCfg.Level_open > hero.ExpLvl {
		return gamedb.ERRPLAYERLVNOTENOUGH
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, skillLvCfg.Consume); err != nil {
		return err
	}

	lv := 1
	if ok {
		lv = userSkill[skillId].Lv + 1
	}
	userSkill[skillId] = &model.SkillUnit{
		Id: skillId,
		Lv: lv,
	}

	skillT := skillCfg.Type
	if skillT == pb.SKILLTYPE_UNIQUE {
		hero.UniqueSkills = userSkill
	} else {
		hero.Skills = userSkill
	}
	user.Dirty = true

	ack.HeroIndex = req.HeroIndex
	ack.SkillType = int32(skillT)
	ack.Skill = builder.BuildSkillUnit(userSkill[skillId])
	newSkillLvConf := gamedb.GetSkillLvConf(userSkill[skillId].Id, userSkill[skillId].Lv)
	if newSkillLvConf != nil && len(newSkillLvConf.Attribute) > 0 {
		this.GetUserManager().UpdateCombat(user, hero.Index)
	}
	kyEvent.SkillLvUp(user, int(req.HeroIndex), skillId, userSkill[skillId].Lv-1, userSkill[skillId].Lv)
	this.GetTask().SendSpecialCheckTaskInfo(user, pb.CONDITION_UPGRADE_SKILL)
	user.UpdateFightUserHeroIndexFun(hero.Index)
	logger.Debug("userId:%v  user.MainLineTask.TaskId:%v  process:%v", user.Id, user.MainLineTask.TaskId, user.MainLineTask.Process)

	return nil
}

/**
 *  @Description: 技能变换位置
 *  @param user
 *  @param req
 *  @param ack
 *  @return error
 */
func (this *SkillManager) ChangePos(user *objs.User, req *pb.SkillChangePosReq, ack *pb.SkillChangePosAck) error {
	skillId, newPos := int(req.SkillId), int(req.Pos)
	if skillId < 1 {
		return gamedb.ERRPARAM
	}
	userSkill, userSkillBag, hero, skillCfg, err := this.GetSkill(user, int(req.HeroIndex), skillId, true)
	if err != nil {
		return err
	}
	skillT := skillCfg.Type
	if skillT == pb.SKILLTYPE_PASSIVE || skillT == pb.SKILLTYPE_PASSIVE2 {
		return gamedb.ERRSKILLNOTWEAR
	}
	// 查看技能是否学习
	_, ok := userSkill[skillId]
	if !ok {
		return gamedb.ERRSKILLNOTSTUDY
	}
	// 查看要设置的技能是否已在槽位
	for pos, sid := range userSkillBag {
		if skillId == sid {
			userSkillBag[pos] = userSkillBag[newPos]
			break
		}
	}
	userSkillBag[newPos] = skillId
	if skillT == pb.SKILLTYPE_UNIQUE {
		hero.UniqueSkillBag = userSkillBag
	} else {
		hero.SkillBag = userSkillBag
	}
	user.Dirty = true

	ack.HeroIndex = req.HeroIndex
	ack.SkillType = int32(skillT)
	ack.SkillBags = builder.BuildSkillBag(hero, skillT)
	return nil
}

/**
 *  @Description: 技能穿戴丶卸下
 *  @param user
 *  @param req
 *  @param ack
 *  @return error
 */
func (this *SkillManager) ChangeWear(user *objs.User, req *pb.SkillChangeWearReq, ack *pb.SkillChangeWearAck) error {
	skillId := int(req.SkillId)
	userSkill, userSkillBag, hero, skillCfg, err := this.GetSkill(user, int(req.HeroIndex), skillId, true)
	if err != nil {
		return err
	}
	skillT := skillCfg.Type
	if skillT == pb.SKILLTYPE_PASSIVE || skillT == pb.SKILLTYPE_PASSIVE2 {
		return gamedb.ERRSKILLNOTWEAR
	}
	// 查看技能是否学习
	_, ok := userSkill[skillId]
	if !ok {
		return gamedb.ERRSKILLNOTSTUDY
	}
	defSkillId := constFight.JOB_SKILL_MAP[hero.Job]
	if skillId == defSkillId {
		return gamedb.ERRCHANGEDEFSKILL
	}
	// 已装备，卸下
	skillPosArr := pb.SKILLPOS_ARRAY
	if skillT == pb.SKILLTYPE_UNIQUE {
		skillPosArr = pb.UNIQUESKILLPOS_ARRAY
	}
	wearPos, newPos := 0, 0
	for _, pos := range skillPosArr {
		if userSkillBag[pos] == 0 && newPos == 0 {
			newPos = pos
		}
		if userSkillBag[pos] == skillId {
			wearPos = pos
			break
		}
	}
	if wearPos != 0 {
		userSkillBag[wearPos] = 0
	} else {
		// 未装备，检查位置是否为空
		if newPos == 0 {
			return gamedb.ERRWEARENOUGH
		}
		userSkillBag[newPos] = skillId
	}

	if skillT == pb.SKILLTYPE_UNIQUE {
		hero.UniqueSkillBag = userSkillBag
	} else {
		hero.SkillBag = userSkillBag
	}
	user.Dirty = true

	ack.HeroIndex = req.HeroIndex
	ack.SkillType = int32(skillT)
	ack.SkillBags = builder.BuildSkillBag(hero, skillT)
	this.GetFight().UpdateUserInfoToFight(user, map[int]bool{hero.Index: true}, true)
	this.GetTask().SendSpecialCheckTaskInfo(user, pb.CONDITION_LEARN_TO_WEAR_SKILL)
	logger.Debug("userId:%v  user.MainLineTask.TaskId:%v  process:%v", user.Id, user.MainLineTask.TaskId, user.MainLineTask.Process)
	//user.UpdateFightUserHeroIndexFun(hero.Index)
	return nil
}

/**
 *  @Description: 技能重置
 *  @param user
 *  @param op
 *  @param req
 *  @param ack
 *  @return error
 */
func (this *SkillManager) Reset(user *objs.User, op *ophelper.OpBagHelperDefault, req *pb.SkillResetReq, ack *pb.SkillResetAck) error {
	skillT := int(req.SkillType)
	userSkills, userSkillBag, hero, _, err := this.GetSkill(user, int(req.HeroIndex), skillT, true)
	if err != nil {
		return err
	}
	defSkillId := constFight.JOB_SKILL_MAP[hero.Job]
	addMap := make(gamedb.ItemInfos, 0)
	for _, skillUnit := range userSkills {
		if skillUnit.Id == defSkillId || skillUnit.Lv == 0 {
			continue
		}
		for i := 1; i <= skillUnit.Lv; i++ {
			skillLvCfg := gamedb.GetSkillLvConf(skillUnit.Id, i-1)
			for _, itemInfo := range skillLvCfg.Consume {
				addMap = append(addMap, &gamedb.ItemInfo{
					ItemId: itemInfo.ItemId,
					Count:  itemInfo.Count,
				})
			}
		}
	}
	if len(addMap) <= 0 {
		return gamedb.ERRRESETSKILL
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, gamedb.GetConf().SkillResetConsume); err != nil {
		return err
	}

	this.GetBag().AddItems(user, addMap, op)
	if skillT != pb.SKILLTYPE_UNIQUE {
		for _, v := range pb.SKILLPOS_ARRAY {
			userSkillBag[v] = 0
			if v == pb.SKILLPOS_SIX {
				userSkillBag[v] = defSkillId
			}
		}
	} else {
		for _, v := range pb.UNIQUESKILLPOS_ARRAY {
			userSkillBag[v] = 0
		}
	}
	userSkills = make(model.Skills)
	if skillT == pb.SKILLTYPE_UNIQUE {
		hero.UniqueSkills = userSkills
		hero.UniqueSkillBag = userSkillBag
	} else {
		userSkills[defSkillId] = &model.SkillUnit{Id: defSkillId, Lv: 1}
		hero.Skills = userSkills
		hero.SkillBag = userSkillBag
	}
	user.Dirty = true

	ack.HeroIndex = req.HeroIndex
	ack.SkillType = req.SkillType
	ack.SkillBags = builder.BuildSkillBag(hero, skillT)
	ack.Skills = builder.BuildSkills(hero, skillT)
	user.UpdateFightUserHeroIndexFun(hero.Index)
	this.GetUserManager().UpdateCombat(user, hero.Index)
	//this.GetFight().UpdateUserInfoToFight(user, map[int]bool{hero.Index:true})
	return nil
}

/**
 *  @Description: 使用技能
 *  @param user
 *  @param req
 *  @return error
 */
func (this *SkillManager) UseSkill(user *objs.User, req *pb.SkillUseReq) error {
	skillId := int(req.SkillId)
	userSkills, _, hero, _, err := this.GetSkill(user, int(req.HeroIndex), skillId, false)
	if err != nil {
		return err
	}
	skill, ok := userSkills[skillId]
	if !ok {
		return gamedb.ERRSKILLNOTSTUDY
	}

	skillLvConf := gamedb.GetSkillLvConf(skillId, skill.Lv)
	startTime := common.GetNowMillisecond()
	endTime := startTime + int64(skillLvConf.CD)
	skill.StartTime = startTime
	skill.EndTime = endTime
	hero.Skills[skillId] = skill
	user.Dirty = true

	this.GetUserManager().SendMessage(user, &pb.SkillUseNtf{
		HeroIndex: req.HeroIndex,
		SkillId:   int32(skillId),
		StartTime: startTime,
		EndTime:   endTime,
	}, true)
	return nil
}

func (this *SkillManager) GetSkill(user *objs.User, heroIndex, skillId int, getBag bool) (model.Skills, model.IntKv, *objs.Hero, *gamedb.SkillSkillCfg, error) {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return nil, nil, nil, nil, gamedb.ERRHERONOTFOUND
	}
	skillT := skillId
	skillCfg := gamedb.GetSkillSkillCfg(skillId)
	if skillCfg != nil {
		skillT = skillCfg.Type
	}
	var skill model.Skills
	var skillBag model.IntKv
	if skillT == pb.SKILLTYPE_UNIQUE {
		skill = hero.UniqueSkills
		if getBag {
			skillBag = hero.UniqueSkillBag
		}
	} else {
		skill = hero.Skills
		if getBag {
			skillBag = hero.SkillBag
		}
	}
	return skill, skillBag, hero, skillCfg, nil
}
