package expPool

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

func NewExperiencePoolManager(m managersI.IModule) *ExperiencePoolManager {
	expPool := &ExperiencePoolManager{}
	expPool.IModule = m
	return expPool
}

type ExperiencePoolManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *ExperiencePoolManager) Load(user *objs.User, ack *pb.ExpPoolLoadAck) error {

	worldLv, _, _, _ := this.GetExpWorldLvlAndLvLimit(user)
	ack.ExpPool = int64(user.Exp)
	ack.WorlLvl = int32(worldLv)
	return nil
}

//
//  Upgrade
//  @Description: 经验池升级
//  @param index  第几个英雄
//
func (this *ExperiencePoolManager) Upgrade(user *objs.User, op *ophelper.OpBagHelperDefault, ack *pb.ExpPoolUpGradeAck, index, times int) error {
	if times <= 0 {
		return gamedb.ERRPARAM
	}

	heroInfo := user.Heros[index]
	if heroInfo == nil {
		logger.Error("ExperiencePoolManager Upgrade  get heroInfo nil  userId:%v  index:%v", user.Id, index)
		return gamedb.ERRHERONOTFOUND
	}
	beforeLv := heroInfo.ExpLvl

	for i := 1; i <= times; i++ {

		nexUpLv := heroInfo.ExpLvl + 1

		expUpCfg := gamedb.GetExpPoolCfg(index, heroInfo.ExpLvl)
		nextCfg := gamedb.GetExpPoolCfg(index, nexUpLv)
		if expUpCfg == nil || nextCfg == nil {
			logger.Info("ExperiencePoolManager Upgrade have to maxLv userId:%v  expLv:%v  nexUpLv:%v", user.Id, heroInfo.ExpLvl, nexUpLv)
			break
		}

		err := this.UpgradeCheck(user, op, expUpCfg)
		if err != nil {
			break
		}

		heroInfo.ExpLvl = nexUpLv
	}
	kyEvent.ExpLvUp(user, index, beforeLv)
	user.Dirty = true

	this.UpGradeTask(user, index)
	this.GetUserManager().UpdateCombat(user, index)
	ack.HeroIndex = int32(index)
	ack.Lvl = int32(user.Heros[index].ExpLvl)
	kyEvent.UserHeroLvUp(user, index)
	return nil
}

func (this *ExperiencePoolManager) AutoUpExpLv(user *objs.User) {
	index := constUser.USER_HERO_MAIN_INDEX
	if len(user.Heros) > 1 {
		return
	}
	heroInfo := user.Heros[index]
	if heroInfo == nil {
		return
	}

	nowLv := heroInfo.ExpLvl
	nowExp := user.Exp

	nowCfg := gamedb.GetExpPoolCfg(index, nowLv)
	if nowCfg.Condition != nil || len(nowCfg.Condition) > 0 {
		return
	}

	allLvCfg := gamedb.GetExpPoolCfgByHeroIndex(index)
	delExp := 0
	afterLv := nowLv
	for i := nowLv; i < len(allLvCfg); i++ {
		cfg := gamedb.GetExpPoolCfg(index, i)
		if cfg != nil {
			if cfg.Condition != nil || len(cfg.Condition) > 0 {
				break
			}
			if cfg.Item[0].Count > nowExp {
				break
			}
			if cfg.Item[0].Count <= nowExp {
				delExp += cfg.Item[0].Count
				afterLv = i + 1
				nowExp -= cfg.Item[0].Count
			}
		}
	}

	if delExp <= 0 {
		return
	}

	enough, _ := this.GetBag().HasEnough(user, pb.ITEMID_EXP, delExp)
	if !enough {
		return
	}
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeArea)
	this.GetBag().Remove(user, op, pb.ITEMID_EXP, delExp)
	user.Heros[index].ExpLvl = afterLv
	user.Dirty = true
	this.UpGradeTask(user, index)
	kyEvent.UserHeroLvUp(user, index)
	this.GetUserManager().SendItemChangeNtf(user, op)
	ack := &pb.ExpPoolUpGradeAck{}
	ack.HeroIndex = int32(index)
	ack.Lvl = int32(afterLv)
	this.GetUserManager().SendMessage(user, ack, true)
}

func (this *ExperiencePoolManager) UpGradeTask(user *objs.User, index int) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("ExperiencePoolManager UpGradeTask Panic Error. %T", err)
		}
	}()
	countLv := 0
	for _, v := range user.Heros {
		countLv += v.ExpLvl
	}
	this.GetTask().UpdateTaskProcess(user, true, false)
	this.GetRank().Append(pb.RANKTYPE_LEVEL, user.Id, countLv, false, false, false)
	this.GetCondition().RecordCondition(user, pb.CONDITION_ONE_HERO_LV, []int{})
	this.GetCondition().RecordCondition(user, pb.CONDITION_THREE_HERO_LV, []int{})
}

//  UpgradeCheck
//  @Description:升级前置条件检查
func (this *ExperiencePoolManager) UpgradeCheck(user *objs.User, op *ophelper.OpBagHelperDefault, expUpCfg *gamedb.ExpLevelLevelCfg) error {
	//升级的前置条件
	if len(expUpCfg.Condition) > 0 {
		for types, lvl := range expUpCfg.Condition {
			_, state := this.GetCondition().Check(user, -1, types, lvl)
			if !state {
				logger.Info("玩家不满足升级前置条件  userId:%v types:%v lvl:%v", user.Id, types, lvl)
				return gamedb.ERRCANNOTUPGRADE
			}
		}
	}

	for _, info := range expUpCfg.Item {
		enough, _ := this.GetBag().HasEnough(user, info.ItemId, info.Count)
		if !enough {
			return gamedb.ERRNOTENOUGHGOODS
		}
	}

	for _, info := range expUpCfg.Item {
		err := this.GetBag().Remove(user, op, info.ItemId, info.Count)
		if err != nil {
			return err
		}
	}
	return nil
}

//  GetExpWorldLvl
//  @Description: 获取当前开服天数世界等级  and  最大存储上限
//  worldLvl  根据开服天数决定的世界等级
//  types -1 不增不减  1:增加
//  maxLimit.Limit  根据玩家三个角色中最大的一个等级决定的每日经验池存储上线
//  addPercent 增加or减少的 系数
func (this *ExperiencePoolManager) GetExpWorldLvlAndLvLimit(user *objs.User) (int, int, int, float64) {

	openDay := this.GetSystem().GetServerOpenDaysByServerId(user.ServerId)
	worldLvl := gamedb.GetExpWorldLevel(openDay)

	maxLv := 0
	for _, info := range user.Heros {
		if info.ExpLvl >= maxLv {
			maxLv = info.ExpLvl
		}
	}

	maxLimit := gamedb.GetExpPoolExpPoolCfg(maxLv)
	if maxLimit == nil {
		maxLimit = gamedb.GetExpPoolExpPoolCfg(gamedb.GetExpPoolLimitMaxLen())
	}

	types := -1
	if maxLv >= worldLvl {
		types = 1
	}

	if maxLv < worldLvl {
		types = 2
	}

	differLv := worldLvl - maxLv
	if differLv < 0 {
		differLv = -differLv
	}
	logger.Debug("GetExpWorldLvlAndLvLimit userId:%v   worldLv:%v  types:%v  differLv:%v", user.Id, worldLvl, types, differLv)
	addPercent := gamedb.GetExpPoolWorldLvBuff(types, differLv)

	return worldLvl, maxLimit.Limit, types, addPercent
}

func (this *ExperiencePoolManager) GetHeroMaxLv(user *objs.User) int {
	//获取武将信息
	if user.Heros == nil || len(user.Heros) == 0 {
		heros, err := modelGame.GetHeroModel().GetHerosByUserId(user.Id)
		if err != nil {
			logger.Error("GetHerosByUserId  err:%v", err)
			return 0
		}
		for _, v := range heros {
			user.Heros[v.Index] = objs.NewHero(v)
		}
	}

	lv := 0
	for _, v := range user.Heros {
		if v.ExpLvl > lv {
			lv = v.ExpLvl
		}
	}
	return lv
}

/**
 *  ExpPilUse
 *  @Description: 使用经验丹
 *  @param user
 *  @param itemId
 *  @param itemNum
 *  @param op
 *  @return error
**/
func (this *ExperiencePoolManager) ExpPillUse(user *objs.User, itemId int, itemNum int, op *ophelper.OpBagHelperDefault) error {
	var err error
	expPillCfg := gamedb.GetExpPillExpPillCfg(itemId)
	if expPillCfg == nil {
		err = gamedb.ERRPARAM
	}
	addExp := 0
	expPillAddExp := gamedb.GetItemBaseCfg(itemId).EffectVal
	if len(user.Heros) == 1 {
		flag := true
		// 校验条件
		toLv := 0
		conditionSlice := make([]int, 1)
		for _, t := range expPillCfg.Condition {
			if t[0] == pb.CONDITION_LEVEL_BETWEEN {
				conditionSlice, flag = this.GetCondition().CheckBySlice(user, -1, t)
				if toLv < t[2] {
					toLv = t[2]
				}
			} else {
				_, flag = this.GetCondition().CheckBySlice(user, -1, t)
			}
			if !flag {
				err = gamedb.ERRCONDITION
				break
			}
		}
		if err != nil {
			return err
		}
		if toLv == 0 {
			return gamedb.ERRPARAM
		}
		// 计算所需经验
		startLv := conditionSlice[0]
		needExp := 0
		for i := startLv; i <= toLv+1; i++ {
			levelCfg := gamedb.GetExpLevelLevelCfg(i)
			for _, info := range levelCfg.Item {
				if info.ItemId == pb.ITEMID_EXP {
					needExp += info.Count
				}
			}
		}
		needExp -= user.Exp
		// 转换数量
		needNum := common.CeilFloat64(float64(needExp)/float64(expPillAddExp))
		if needNum <= itemNum {
			err = this.GetBag().Remove(user, op, itemId, needNum)
			addExp += expPillAddExp * needNum
		} else {
			err = this.GetBag().Remove(user, op, itemId, itemNum)
			addExp += expPillAddExp * itemNum
		}
	} else {
		err = this.GetBag().Remove(user, op, itemId, itemNum)
		addExp += expPillAddExp * itemNum
	}

	if err != nil {
	    return err
	}
	this.GetBag().Add(user, op, pb.ITEMID_EXP, addExp)

	return err
}
