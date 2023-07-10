package inside

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constMax"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"fmt"
)

func NewInsideManager(module managersI.IModule) *InsideManager {
	return &InsideManager{IModule: module}
}

type InsideManager struct {
	util.DefaultModule
	managersI.IModule
}

const (
	PRIVILEGE_AUTO_UP_NUM = 10
)

/**
 *  @Description: 内功升星
 *  @param user
 *  @param heroIndex
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *InsideManager) InsideUpStar(user *objs.User, heroIndex int, op *ophelper.OpBagHelperDefault) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	inside := hero.Inside

	randMap := make(map[int]int)
	var grade, order int
	for _, v := range pb.INSIDETYPE_ARRAY {
		aId := inside.Acupoint[v]
		insideCfg := gamedb.GetInsideArtInsideArtCfg(aId)
		grade, order = insideCfg.Grade, insideCfg.Order
		maxStar := gamedb.GetInsideMaxStar(grade, order)
		if insideCfg.Star >= maxStar || gamedb.GetInsideArtInsideArtCfg(aId+1) == nil || !this.GetCondition().CheckMulti(user, heroIndex, insideCfg.Condition) {
			continue
		}
		randMap[v] = 0
	}
	if len(randMap) == 0 {
		return gamedb.ERRLVENOUGH
	}
	chance := 100 / len(randMap)
	for pos := range randMap {
		randMap[pos] = chance
	}
	randAcupoint := common.RandWeightByMap(randMap)
	insideId := inside.Acupoint[randAcupoint]
	consumeCfg := gamedb.GetInsideByGradeAndOrder(grade, order, 0)
	if err := this.GetBag().Remove(user, op, consumeCfg.Consume.ItemId, consumeCfg.Consume.Count); err != nil {
		return err
	}

	insideCfg := gamedb.GetInsideArtInsideArtCfg(insideId)
	insideStarWeight := common.RandWeightByMap(gamedb.GetInsideStarWeightMap(grade, order))
	// 判断是否超过当前最大星数
	curGrade, curOrder, curStar := insideCfg.Grade, insideCfg.Order, insideCfg.Star
	oldStar := curStar
	maxStar := gamedb.GetInsideMaxStar(curGrade, curOrder)
	curStar += 1 * insideStarWeight
	if curStar > maxStar {
		curStar = maxStar
	}

	randWeightId := gamedb.GetInsideByGradeAndOrder(curGrade, curOrder, curStar).Id
	inside.Acupoint[randAcupoint] = randWeightId
	kyEvent.InsideStarUp(user, heroIndex, randAcupoint, oldStar, curStar)

	this.GetTask().AddTaskProcess(user, pb.CONDITION_XIU_LIAN_NEI_GONG, 1)
	this.ChangeOperation(user, heroIndex)
	return nil
}

/**
 *  @Description: 内功升重
 *  @param user
 *  @param heroIndex
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *InsideManager) InsideUpOrder(user *objs.User, heroIndex int, op *ophelper.OpBagHelperDefault, ack *pb.InsideUpOrderAck) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	inside := hero.Inside

	insideCfg := gamedb.GetInsideArtInsideArtCfg(inside.Acupoint[pb.INSIDETYPE_ONE])
	if insideCfg.Consume.ItemId != 0 {
		if err := this.GetBag().Remove(user, op, insideCfg.Consume.ItemId, insideCfg.Consume.Count); err != nil {
			return err
		}
	}
	curGrade, curOrder := insideCfg.Grade, insideCfg.Order
	maxStar := gamedb.GetInsideMaxStar(curGrade, curOrder)
	for _, id := range inside.Acupoint {
		if gamedb.GetInsideArtInsideArtCfg(id+1) == nil {
			return gamedb.ERRLVENOUGH
		}
		if gamedb.GetInsideArtInsideArtCfg(id).Star != maxStar {
			return gamedb.ERRPARAM
		}
	}
	for pos := range inside.Acupoint {
		inside.Acupoint[pos]++
	}

	ack.HeroIndex = int32(heroIndex)
	ack.InsideInfo = builder.BuildInsideInfo(hero)

	if insideCfg != nil {
		this.GetTask().AddTaskProcess(user, pb.CONDITION_XIU_LIAN_NEI_GONG_1, -1)
	}
	this.GetCondition().RecordCondition(user, pb.CONDITION_THREE_INSIDE_CHONG, []int{1})
	this.GetTask().SendSpecialCheckTaskInfo(user, pb.CONDITION_UPGRADE_NEI_GONG)
	this.ChangeOperation(user, heroIndex)
	return nil
}

/**
 *  @Description: 内功升阶
 *  @param user
 *  @param heroIndex
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *InsideManager) InsideUpGrade(user *objs.User, heroIndex int, op *ophelper.OpBagHelperDefault, ack *pb.InsideUpGradeAck) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	inside := hero.Inside

	oneInsideId := inside.Acupoint[pb.INSIDETYPE_ONE]
	insideCfg := gamedb.GetInsideArtInsideArtCfg(oneInsideId)
	curGrade, curOrder := insideCfg.Grade, insideCfg.Order
	gradeCfg := gamedb.GetInsideGradeInsideGradeCfg(curGrade)
	if gradeCfg == nil {
		return gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg(fmt.Sprintf(`insideGradeCfg grade:%v`, curGrade))
	}
	if !this.GetCondition().CheckMulti(user, heroIndex, gradeCfg.Condition) {
		return gamedb.ERRCONDITION
	}
	maxStar := gamedb.GetInsideMaxStar(curGrade, curOrder)
	for _, id := range inside.Acupoint {
		if gamedb.GetInsideArtInsideArtCfg(id+1) == nil {
			return gamedb.ERRLVENOUGH
		}
		if gamedb.GetInsideArtInsideArtCfg(id).Star != maxStar {
			return gamedb.ERRPARAM
		}
	}
	if err := this.GetBag().Remove(user, op, gradeCfg.Consume.ItemId, gradeCfg.Consume.Count); err != nil {
		return err
	}
	isUp := common.RandByTenShousand(gradeCfg.Success)
	if isUp {
		for pos := range inside.Acupoint {
			inside.Acupoint[pos]++
		}
	}

	ack.HeroIndex = int32(heroIndex)
	ack.InsideInfo = builder.BuildInsideInfo(hero)
	ack.Res = isUp

	this.GetTask().SendSpecialCheckTaskInfo(user, pb.CONDITION_UPGRADE_NEI_GONG)
	this.ChangeOperation(user, heroIndex)
	return nil
}

/**
 *  @Description: 内功技能升级
 *  @param user
 *  @param heroIndex
 *  @param skillId	技能id
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *InsideManager) InsideSkillUpLv(user *objs.User, heroIndex, skillId int, op *ophelper.OpBagHelperDefault, ack *pb.InsideSkillUpLvAck) error {
	if skillId < 1 {
		return gamedb.ERRPARAM
	}
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	insideSkill := hero.Inside.Skill[skillId]
	if insideSkill == nil {
		hero.Inside.Skill[skillId] = &model.InsideSkill{}
		insideSkill = hero.Inside.Skill[skillId]
	}
	maxLv := gamedb.GetMaxValById(skillId, constMax.MAX_INSIDE_SKILL_LEVEL)
	if insideSkill.Level >= maxLv {
		return gamedb.ERRLVENOUGH
	}

	insideCfg := gamedb.GetInsideSkillBySidAndLv(skillId, insideSkill.Level)
	if len(insideCfg.Condition) > 0 {
		if check := this.GetCondition().CheckMulti(user, heroIndex, insideCfg.Condition); !check {
			return gamedb.ERRCONDITION
		}
	}

	if insideCfg.Type == pb.INSIDESKILLTYPE_ONE {
		remNum := insideCfg.Consume.Count
		hasNum, _ := this.GetBag().GetItemNum(user, insideCfg.Consume.ItemId)
		if hasNum < remNum {
			remNum = hasNum
		}
		if err := this.GetBag().Remove(user, op, insideCfg.Consume.ItemId, remNum); err != nil {
			return gamedb.ERRNOTENOUGHGOODS
		}
		insideSkill.Exp += remNum
		if insideSkill.Exp >= insideCfg.Consume.Count {
			insideSkill.Level++
			insideSkill.Exp = 0
		}
	} else {
		if err := this.GetBag().Remove(user, op, insideCfg.Consume.ItemId, insideCfg.Consume.Count); err != nil {
			return gamedb.ERRNOTENOUGHGOODS
		}
		insideSkill.Level++
	}

	ack.HeroIndex = int32(heroIndex)
	ack.InsideInfo = builder.BuildInsideInfo(hero)
	this.ChangeOperation(user, heroIndex)
	return nil
}

/**
 *  @Description: 内功一键升级
 *  @param user
 *  @param heroIndex
 *  @param op
 *  @return error
 */
func (this *InsideManager) AutoUp(user *objs.User, heroIndex int, op *ophelper.OpBagHelperDefault) error {
	if privilege := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_INSIDE_AUTO_UP); privilege == 0 {
		return gamedb.ERRVIPLVNOTENOUGH
	}
	for i := 0; i < PRIVILEGE_AUTO_UP_NUM; i++ {
		err := this.InsideUpStar(user, heroIndex, op)
		if i == 0 && err != nil {
			return err
		}
		if err != nil {
			break
		}
	}
	return nil
}

func (this *InsideManager) ChangeOperation(user *objs.User, heroIndex int) {
	this.GetUserManager().UpdateCombat(user, -1)
	this.GetCondition().RecordCondition(user, pb.CONDITION_UPGRADE_NEI_GONG, []int{})
	this.GetCondition().RecordCondition(user, pb.CONDITION_THREE_INSIDE_GRADE, []int{})
}
