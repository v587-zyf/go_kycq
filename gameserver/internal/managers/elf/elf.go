package elf

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
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

const (
	EQUIP_GRADE_ONE = 1
	EFL_LV_DEF      = 1
)

var ELF_SKILLBAG_DEF = map[int]int{
	1: 110000,
	2: 110001,
	3: 110002,
}

type Elf struct {
	util.DefaultModule
	managersI.IModule
}

func NewElf(module managersI.IModule) *Elf {
	return &Elf{IModule: module}
}

func (this *Elf) Online(user *objs.User) {
	curLv, curExp := user.Elf.Lv, user.Elf.Exp
	nextCfg := gamedb.GetElfGrowElfGrowCfg(curLv + 1)
	if nextCfg == nil {
		return
	}
	lvCfg := gamedb.GetElfGrowElfGrowCfg(curLv)
	if lvCfg == nil {
		return
	}
	for nextCfg != nil && curExp >= nextCfg.Experience {
		if !this.GetCondition().CheckMulti(user, -1, lvCfg.Condition) {
			break
		}
		curLv++
		curExp -= nextCfg.Experience
		lvCfg = gamedb.GetElfGrowElfGrowCfg(curLv)
		nextCfg = gamedb.GetElfGrowElfGrowCfg(curLv + 1)
	}
	user.Elf.Lv = curLv
	user.Elf.Exp = curExp
}

/**
 *  @Description: 精灵喂养
 *  @param user
 *  @param op
 *  @param positions	背包中道具位置
 *  @return error
 */
func (this *Elf) Feed(user *objs.User, op *ophelper.OpBagHelperDefault, positions []int32, ack *pb.ElfFeedAck) error {
	curLv, curExp := user.Elf.Lv, user.Elf.Exp
	beforeLv := curLv
	isUpdateElf := false
	if curLv == 0 {
		//初始化
		isUpdateElf = true
		curLv = EFL_LV_DEF
		for pos, skillId := range ELF_SKILLBAG_DEF {
			user.Elf.SkillBag[pos] = skillId
		}
	}
	nextLvCfg := gamedb.GetElfGrowElfGrowCfg(curLv + 1)
	//if nextLvCfg == nil {
	//	return gamedb.ERRLVENOUGH
	//}
	lvCfg := gamedb.GetElfGrowElfGrowCfg(curLv)
	if lvCfg == nil {
		return gamedb.ERRPARAM
	}
	if !this.GetCondition().CheckMulti(user, -1, lvCfg.Condition) {
		return gamedb.ERRCONDITION
	}

	for _, v := range positions {
		itemId := this.GetBag().GetItemByPosition(user, int(v)).ItemId
		itemBaseCfg := gamedb.GetItemBaseCfg(itemId)
		if itemBaseCfg == nil || (itemBaseCfg.ElfExperience == 0 && len(itemBaseCfg.ElfRecover) < 1) {
			logger.Error("itemConf error itemId:%v bagPos:%v", itemId, v)
			return gamedb.ERRPARAM
		}
	}

	addMap := make(map[int]int)
	removeMap := make(map[int]int)
	for _, v := range positions {
		bagItem := this.GetBag().GetItemByPosition(user, int(v))
		itemId, itemCount := bagItem.ItemId, bagItem.Count
		itemBaseCfg := gamedb.GetItemBaseCfg(itemId)
		if nextLvCfg == nil {
			curExp += itemCount * itemBaseCfg.ElfExperience
			removeMap[itemId] += itemCount
			for _, info := range itemBaseCfg.ElfRecover {
				addMap[info.ItemId] += info.Count * itemCount
			}
		} else {
			for this.GetCondition().CheckMulti(user, -1, lvCfg.Condition) {
				needExp := lvCfg.Experience - curExp
				needCount := common.CeilFloat64(float64(needExp) / float64(itemBaseCfg.ElfExperience))
				if itemCount >= needCount {
					isUpdateElf = true
					curExp = needCount*itemBaseCfg.ElfExperience - needExp
					removeMap[itemId] += needCount
					for _, info := range itemBaseCfg.ElfRecover {
						addMap[info.ItemId] += info.Count * needCount
					}
					itemCount -= needCount
					if gamedb.GetElfGrowElfGrowCfg(curLv+1) == nil {
						break
					}
					curLv += 1
					lvCfg = gamedb.GetElfGrowElfGrowCfg(curLv)
				} else {
					curExp += itemCount * itemBaseCfg.ElfExperience
					removeMap[itemId] += itemCount
					for _, info := range itemBaseCfg.ElfRecover {
						addMap[info.ItemId] += info.Count * itemCount
					}
					break
				}
			}
		}
	}
	if len(removeMap) > 0 {
		removeItems := make(gamedb.ItemInfos, len(removeMap))
		removeIndex := 0
		for itemId, count := range removeMap {
			removeItems[removeIndex] = &gamedb.ItemInfo{
				ItemId: itemId,
				Count:  count,
			}
			removeIndex++
		}
		if err := this.GetBag().RemoveItemsInfos(user, op, removeItems); err != nil {
			return err
		}
	}
	if len(addMap) > 0 {
		recoverLimitCfgMap := make(map[int]int)
		vipPrivilege := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_ELF_RECOVER_LIMIT_ADD)
		for _, itemInfo := range gamedb.GetConf().ElfRecoverLimit {
			itemCount := itemInfo.Count
			if vipPrivilege > 0 {
				itemCount += vipPrivilege
			}
			recoverLimitCfgMap[itemInfo.ItemId] += itemCount
		}

		if user.Elf.RecoverLimit == nil {
			user.Elf.RecoverLimit = make(model.IntKv)
		}
		userRecoverLimit := user.Elf.RecoverLimit
		addItems := make(gamedb.ItemInfos, len(addMap))
		addIndex := 0
		for itemId, count := range addMap {
			addNum := count
			if n, ok := recoverLimitCfgMap[itemId]; ok {
				canAddNum := int(float64(n) - float64(userRecoverLimit[itemId]))
				if canAddNum < 0 {
					addNum = 0
				} else if canAddNum < addNum {
					addNum = canAddNum
				}
				userRecoverLimit[itemId] += addNum
			}
			addItems[addIndex] = &gamedb.ItemInfo{
				ItemId: itemId,
				Count:  addNum,
			}
			addIndex++
		}
		this.GetBag().AddItems(user, addItems, op)
		logger.Debug("精灵回收今日货币:%v", user.Elf.RecoverLimit)
	}

	user.Elf.Lv = curLv
	user.Elf.Exp = curExp

	ack.Lv = int32(user.Elf.Lv)
	ack.Exp = int32(user.Elf.Exp)
	ack.Goods = op.ToChangeItems()
	ack.ReceiveLimit = builder.BuildElfReceiveLimit(user.Elf.RecoverLimit)

	if isUpdateElf {
		this.GetFight().UpdateUserElf(user)
	}
	this.GetUserManager().UpdateCombat(user, -1)
	this.GetTask().AddTaskProcess(user, pb.CONDITION_UPGRADE_JIN_LIN_LV, -1)
	this.GetCondition().RecordCondition(user, pb.CONDITION_UPGRADE_JIN_LIN_LV, []int{})
	kyEvent.JinLinUp(user, beforeLv, user.Elf.Lv)
	return nil
}

/**
 *  @Description: 精灵学习、升级技能
 *  @param user
 *  @param skillId
 *  @return error
 */
func (this *Elf) SkillUpLv(user *objs.User, skillId int, ack *pb.ElfSkillUpLvAck) error {
	cfg := gamedb.GetSkillSkillCfg(skillId)
	if cfg == nil {
		return gamedb.ERRSKILLNOTFOUND
	}

	userElf := user.Elf
	lv, ok := userElf.Skills[skillId]
	maxLv := gamedb.GetMaxValById(skillId, constMax.MAX_ELF_SKILL_LEVEL)
	if ok && lv >= maxLv {
		return gamedb.ERRLVENOUGH
	}

	var skillCfg *gamedb.ElfSkillElfGrowCfg
	if !ok {
		skillCfg = gamedb.GetElfSkillBySkillIdAndLv(skillId, 0)
	} else {
		skillCfg = gamedb.GetElfSkillBySkillIdAndLv(skillId, lv)
	}
	if check := this.GetCondition().CheckMulti(user, -1, skillCfg.Condition); !check {
		return gamedb.ERRCONDITION
	}
	userElf.Skills[skillId]++
	user.Dirty = true

	this.GetFight().UpdateUserElf(user)

	ack.SkillId = int32(skillId)
	ack.SkillLv = int32(userElf.Skills[skillId])
	ack.SkillBag = builder.BuildElfSkillBag(userElf.SkillBag)
	return nil
}

/**
 *  @Description: 精灵技能换位置
 *  @param user
 *  @param skillId
 *  @param pos
 *  @param ack
 *  @return error
 */
func (this *Elf) SkillChangePos(user *objs.User, skillId, pos int, ack *pb.ElfSkillChangePosAck) error {
	//userElf := user.Elf
	//
	//if _, ok := userElf.Skills[skillId]; !ok {
	//	return gamedb.ERRSKILLNOTSTUDY
	//}
	//oldSkillId, ok := userElf.SkillBag[pos]
	//if _, ok := userElf.Skills[oldSkillId]; !ok {
	//	return gamedb.ERRSKILLNOTSTUDY
	//}
	//if !ok {
	//	return gamedb.ERRSKILLPOS
	//}
	//for p, sid := range userElf.SkillBag {
	//	if skillId == sid {
	//		userElf.SkillBag[p] = oldSkillId
	//		break
	//	}
	//}
	//userElf.SkillBag[pos] = skillId
	//user.Dirty = true
	//
	//ack.SkillBag = builder.BuildElfSkillBag(userElf.SkillBag)
	return nil
}
