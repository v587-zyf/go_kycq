package talent

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constConstant"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

func NewTalentManager(m managersI.IModule) *TalentManager {
	return &TalentManager{
		IModule: m,
	}
}

type TalentManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *TalentManager) Online(user *objs.User) {
	for _, hero := range user.Heros {
		this.ComputePoints(hero)
	}
	this.TalentGeneral(user)
}

/**
 *  @Description: 天赋升级
 *  @param user
 *  @param heroIndex
 *  @param id		talentWay表id
 *  @param isMax	是否升满
 *  @param ack
 *  @return error
 */
func (this *TalentManager) UpLv(user *objs.User, heroIndex, id int, isMax bool, ack *pb.TalentUpLvAck) error {
	if id < 1 {
		return gamedb.ERRPARAM
	}
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	this.ComputePoints(hero)

	talentWayId := gamedb.GetTalentByJobAndId(hero.Job, id)
	if talentWayId == 0 {
		return gamedb.ERRJOB
	}
	heroTalent := hero.Talent
	heroTalentList := heroTalent.TalentList
	heroTalentInfo, ok := heroTalentList[talentWayId]
	if !ok {
		heroTalentList[talentWayId] = &model.TalentUnit{
			Talents: make(model.IntKv),
		}
		heroTalentInfo = heroTalentList[talentWayId]
	}
	oldLv := heroTalentInfo.Talents[id]
	if gamedb.GetTalentLevelTalentLevelCfg(gamedb.GetRealId(id, oldLv+1)) == nil {
		return gamedb.ERRLVENOUGH
	}
	newLv, forNum, consumeNum := oldLv, oldLv+1, 0
	if isMax {
		forNum = constConstant.COMPUTE_TEN_THOUSAND
	}
	for i := oldLv; i < forNum; i++ {
		levelCfg := gamedb.GetTalentLevelTalentLevelCfg(gamedb.GetRealId(id, i))
		if levelCfg == nil {
			break
		}
		newLvCfg := gamedb.GetTalentLevelTalentLevelCfg(gamedb.GetRealId(id, i+1))
		if newLvCfg == nil {
			break
		}
		breakFlag := false
		for talentWayId, count := range levelCfg.Requirement1 {
			heroTalentInfo, ok := heroTalentList[talentWayId]
			if !ok || heroTalentInfo.UsePoints < count {
				breakFlag = true
				break
			}
		}
		useCount := 0
		for _, talentUnit := range heroTalentList {
			useCount += talentUnit.UsePoints
		}
		if useCount < levelCfg.Requirement2 {
			breakFlag = true
			break
		}
		if consumeNum == 0 {
			if heroTalent.SurplusPoints < levelCfg.Count {
				return gamedb.ERRTALENTPOINTNOTENOUGH
			}
		} else {
			if heroTalent.SurplusPoints <= consumeNum {
				breakFlag = true
				break
			}
		}
		if breakFlag {
			break
		}
		newLv++
		consumeNum += levelCfg.Count
	}
	if newLv == oldLv {
		return gamedb.ERRCONDITION
	}

	heroTalent.SurplusPoints -= consumeNum
	heroTalentInfo.UsePoints += consumeNum
	kyEvent.TianFuUp(user, heroIndex, id, heroTalentInfo.Talents[id], newLv)
	heroTalentInfo.Talents[id] = newLv

	ack.HeroIndex = int32(heroIndex)
	ack.Id = int32(id)
	ack.TalentInfo = builder.BuildTalent(hero)

	this.GetUserManager().UpdateCombat(user, heroIndex)
	user.UpdateFightUserHeroIndexFun(heroIndex)
	this.GetTask().UpdateTaskProcess(user, false, false)
	this.GetCondition().RecordCondition(user, pb.CONDITION_TIAN_FU_ALL_LV, []int{})
	return nil
}

/**
 *  @Description: 天赋重置
 *  @param user
 *  @param heroIndex
 *  @param id	talentWay表id（重置所有传-1）
 *  @param ack
 *  @return error
 */
func (this *TalentManager) Reset(user *objs.User, heroIndex, id int, op *ophelper.OpBagHelperDefault, ack *pb.TalentResetAck) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	this.ComputePoints(hero)

	heroTalent := hero.Talent
	if id == -1 {
		if heroTalent.SurplusPoints == heroTalent.GetPoints {
			return gamedb.ERRTALENTRESET
		}
		consume := gamedb.GetConf().ResetTalentAll
		err := this.GetBag().Remove(user, op, consume.ItemId, consume.Count)
		if err != nil {
			return err
		}
		hero.Talent.TalentList = make(map[int]*model.TalentUnit)
		heroTalent.SurplusPoints = heroTalent.GetPoints
	} else {
		if heroTalent.TalentList[id] == nil || heroTalent.TalentList[id].UsePoints == 0 {
			return gamedb.ERRTALENTRESET
		}
		consume := gamedb.GetConf().ResetTalentPart
		err := this.GetBag().Remove(user, op, consume.ItemId, consume.Count)
		if err != nil {
			return err
		}
		heroTalent.SurplusPoints += hero.Talent.TalentList[id].UsePoints
		hero.Talent.TalentList[id] = &model.TalentUnit{Talents: make(model.IntKv)}
	}

	ack.HeroIndex = int32(heroIndex)
	ack.Id = int32(id)
	ack.TalentInfo = builder.BuildTalent(hero)

	this.GetUserManager().UpdateCombat(user, heroIndex)
	user.UpdateFightUserHeroIndexFun(heroIndex)
	this.GetCondition().RecordCondition(user, pb.CONDITION_TIAN_FU_ALL_LV, []int{})
	return nil
}

/**
 *  @Description: 计算用户获取
 *  @param user
 *  @param heroIndex
 */
func (this *TalentManager) ComputePoints(hero *objs.Hero) {
	heroTalent := hero.Talent
	point := 0
	talentGetCfgs := gamedb.GetTalentGetCfgs()
	for i := 1; i <= hero.ExpLvl; i++ {
		cfg, ok := talentGetCfgs[i]
		if !ok {
			break
		}
		point += cfg.TalentCount
	}
	heroTalent.GetPoints = point
	heroTalent.SurplusPoints = point
	for _, talentUnit := range heroTalent.TalentList {
		heroTalent.SurplusPoints -= talentUnit.UsePoints
	}
}
