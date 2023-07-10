package equip

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
)

/**
 *  @Description: 装备洗练
 *  @param user
 *  @param heroIndex
 *  @param pos			装备位置
 *  @param propIndex	属性下标（初始化传-1）
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *EquipManager) Clear(user *objs.User, heroIndex, pos, propIndex int, op *ophelper.OpBagHelperDefault, ack *pb.ClearAck) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}

	if propIndex != -1 && len(hero.EquipClear[pos]) <= 0 {
		return gamedb.ERRPARAM
	}
	if propIndex == -1 && len(hero.EquipClear[pos]) > 0 {
		return gamedb.ERRPARAM
	}

	minGrade := gamedb.GetConf().EquipClearGradeMin
	equip := hero.Equips[pos]
	if equip.ItemId == 0 {
		return gamedb.ERRNOTWEAREQUIP
	}
	equipCfg := gamedb.GetEquipEquipCfg(equip.ItemId)
	if equipCfg.Class < minGrade {
		return gamedb.ERREQUIPLV
	}
	consumeCfg := gamedb.GetConf().EquipClearConsume
	if err := this.GetBag().RemoveItemsInfos(user, op, consumeCfg); err != nil {
		return err
	}

	conf := gamedb.GetWashByPosAndGrade(EQUIP_POS_TO_TYPE[pos], equipCfg.Class)
	if conf == nil {
		logger.Error("Clear GetWashCfg err;pos:%v grade:%v", pos, equipCfg.Class)
		return gamedb.ERRPARAM
	}
	randNum := 1
	hasPropMap := make(map[int]int)
	if propIndex == -1 {
		randNum = gamedb.GetConf().EquipClearPropMax
	} else {
		for _, clearUnit := range hero.EquipClear[pos] {
			if clearUnit.PropId == 0 {
				continue
			}
			hasPropMap[clearUnit.PropId] = clearUnit.Value
		}
	}
	oldPropMap := make(map[int]int)
	for _, unit := range hero.EquipClear[pos] {
		oldPropMap[unit.PropId] = unit.Value
	}
	propSlice := this.randClearPropSlice(hero, conf.Washrand_group, hasPropMap, propIndex, pos, equipCfg.Class, randNum)
	if propIndex == -1 {
		hero.EquipClear[pos] = propSlice
	} else {
		hero.EquipClear[pos][propIndex] = propSlice[0]
	}
	propMap := make(map[int]int)
	for _, unit := range hero.EquipClear[pos] {
		propMap[unit.PropId] = unit.Value
	}
	kyEvent.Clear(user, heroIndex, pos, oldPropMap, propMap)

	ack.HeroIndex = int32(heroIndex)
	ack.Pos = int32(pos)
	ack.Goods = op.ToChangeItems()
	ack.EquipClear = &pb.EquipClearArr{EquipClearInfo: builder.BuildEquipClearUnit(hero.EquipClear[pos])}

	this.GetUserManager().UpdateCombat(user, heroIndex)
	return nil
}

func (this *EquipManager) randClearPropSlice(hero *objs.Hero, washRandGroup gamedb.IntSlice, hasPropMap map[int]int, propIndex, pos, grade, randNum int) []*model.EquipClearUnit {
	propSlice := make([]*model.EquipClearUnit, 0)
	// 颜色随机
	randQualityMap := make(map[int]int)
	for _, v := range washRandGroup {
		randCfg := gamedb.GetWashrandRandCfg(v)
		randQualityMap[v] = randCfg.Weight
	}
	for i := 0; i < randNum; i++ {
		randId := common.RandWeightByMap(randQualityMap)
		randPropCfg := gamedb.GetWashrandRandCfg(randId)
		// 属性随机
		weightSlice := make([]int, len(randPropCfg.Attribute))
		for kk, vv := range randPropCfg.Attribute {
			if _, ok := hasPropMap[vv[0]]; ok {
				if propIndex == -1 {
					continue
				} else {
					heroClear := hero.EquipClear[pos][propIndex]
					if vv[0] != heroClear.PropId || (vv[0] == heroClear.PropId && heroClear.Value >= vv[3]) {
						continue
					}
				}
			}
			weightSlice[kk] = vv[1]
		}
		randPropIndex := common.RandWeightByIntSlice(weightSlice)
		pid := randPropCfg.Attribute[randPropIndex][0]
		val := common.RandNum(randPropCfg.Attribute[randPropIndex][2], randPropCfg.Attribute[randPropIndex][3])
		hasPropMap[pid] = val
		propSlice = append(propSlice, &model.EquipClearUnit{
			Grade:  grade,
			Color:  randPropCfg.Quality,
			PropId: pid,
			Value:  val,
		})
	}
	return propSlice
}
