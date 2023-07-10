package equip

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

var (
	EQUIP_POS_TO_TYPE = map[int]int{
		pb.EQUIPPOS_ONE:   pb.EQUIPTYPE_WEAPON_R,
		pb.EQUIPPOS_TWO:   pb.EQUIPTYPE_WEAPON_L,
		pb.EQUIPPOS_THREE: pb.EQUIPTYPE_HELMET,
		pb.EQUIPPOS_FOUR:  pb.EQUIPTYPE_CLOTHES,
		pb.EQUIPPOS_FIVE:  pb.EQUIPTYPE_BELT,
		pb.EQUIPPOS_SIX:   pb.EQUIPTYPE_SHOES,
		pb.EQUIPPOS_SEVEN: pb.EQUIPTYPE_RING,
		pb.EQUIPPOS_EIGHT: pb.EQUIPTYPE_BRACELET,
		pb.EQUIPPOS_NINE:  pb.EQUIPTYPE_RING,
		pb.EQUIPPOS_TEN:   pb.EQUIPTYPE_BRACELET,
	}
)

func NewEquipManager(module managersI.IModule) *EquipManager {
	return &EquipManager{IModule: module}
}

type EquipManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *EquipManager) Online(user *objs.User) {}

/**
 *  @Description: 装备更换
 *  @param user
 *  @param heroIndex
 *  @param op
 *  @param equipPos		装备位置
 *  @param EquipBagPos	背包位置
 *  @return error
 */
func (this *EquipManager) EquipChange(user *objs.User, heroIndex int, op *ophelper.OpBagHelperDefault, equipPos int, EquipBagPos int) error {
	heros := user.Heros[heroIndex]
	if heros == nil {
		return gamedb.ERRHERONOTFOUND
	}
	equips := heros.Equips
	if !pb.EQUIPPOS_MAP[equipPos] {
		return gamedb.ERRPARAM
	}
	item := this.GetBag().GetItemByPosition(user, EquipBagPos)
	if item == nil || item.ItemId == 0 {
		return gamedb.ERRNOTENOUGHGOODS
	}
	equipType := EQUIP_POS_TO_TYPE[equipPos]
	err := this.checkWearCondition(user, heroIndex, equipType, item.ItemId)
	if err != nil {
		return err
	}
	newEquip, err := this.GetBag().EquipChange(user, op, EquipBagPos, equipType, equips[equipPos])
	if err != nil {
		return err
	}
	if newEquip == nil {
		return gamedb.ERRUNKNOW
	}
	equips[equipPos] = newEquip
	this.GetUserManager().SendDisplay(user)
	this.GetTask().AddTaskProcess(user, pb.CONDITION_WEAR_EQUIP, -1)
	this.GetUserManager().UpdateCombat(user, heroIndex)
	this.GetTask().SendSpecialCheckTaskInfo(user, pb.CONDITION_ALL_HEROS_WEAR_ASSIGN_EQUIP)
	this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_HEROS_WEAR_ASSIGN_EQUIP, []int{})
	return nil
}

/**
 *  @Description: 装备一键更换
 *  @param user
 *  @param heroIndex
 *  @param op
 *  @return []int	更换了哪些位置
 *  @return error
 */
func (this *EquipManager) EquipChangeOneKey(user *objs.User, heroIndex int, op *ophelper.OpBagHelperDefault) ([]int, error) {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return nil, gamedb.ERRHERONOTFOUND
	}
	heroEquips, heroJob := hero.Equips, hero.Job
	userBag, userEquipBag := user.Bag, user.EquipBag
	newEquips := make(map[int]*model.Item)
	changePos := make([]int, 0)
	for _, v := range userBag {
		if v.EquipIndex <= 0 {
			continue
		}
		equipCfg := gamedb.GetEquipEquipCfg(v.ItemId)
		equipPos, equipItemId, equipIndex := equipCfg.Type, v.ItemId, v.EquipIndex
		err := this.checkWearCondition(user, heroIndex, equipPos, equipItemId)
		if err != nil {
			continue
		}
		equipPos, tmpEquip := this.compareEquip(heroJob, userEquipBag, newEquips, equipPos, heroEquips)
		isGood := this.EquipCompareCombat(heroJob, userEquipBag[equipIndex], tmpEquip)
		if isGood {
			newEquip, _ := this.GetBag().EquipChange(user, op, v.Position, EQUIP_POS_TO_TYPE[equipPos], heroEquips[equipPos])
			heroEquips[equipPos] = newEquip
			changePos = append(changePos, equipPos)
		}
	}

	this.GetUserManager().UpdateCombat(user, heroIndex)
	this.GetUserManager().SendDisplay(user)

	this.GetTask().AddTaskProcess(user, pb.CONDITION_WEAR_EQUIP, -1)
	this.GetTask().SendSpecialCheckTaskInfo(user, pb.CONDITION_ALL_HEROS_WEAR_ASSIGN_EQUIP)
	this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_HEROS_WEAR_ASSIGN_EQUIP, []int{})
	logger.Debug("一键换转，更换装备位置,玩家：%v,位置%v", user.NickName, changePos)
	return changePos, nil
}

/**
 *  @Description: 装备卸下
 *  @param user
 *  @param heroIndex
 *  @param pos	装备位置
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *EquipManager) EquipRemove(user *objs.User, heroIndex, pos int, op *ophelper.OpBagHelperDefault, ack *pb.EquipRemoveAck) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	if !pb.EQUIPPOS_MAP[pos] {
		return gamedb.ERRNOTEQUIP
	}
	_, err := this.GetBag().EquipChange(user, op, -1, EQUIP_POS_TO_TYPE[pos], hero.Equips[pos])
	if err != nil {
		return err
	}
	hero.Equips[pos] = &model.Equip{}

	ack.Pos = int32(pos)
	ack.HeroIndex = int32(heroIndex)
	this.GetUserManager().SendDisplay(user)
	this.GetUserManager().UpdateCombat(user, heroIndex)
	return nil
}

func (this *EquipManager) compareEquip(heroJob int, userEquipBag model.EquipBag, newEquips map[int]*model.Item, equipPos int, heroEquips model.Equips) (int, *model.Equip) {
	var tmpEquip *model.Equip
	newEquip, has := newEquips[equipPos]
	if has {
		tmpEquip = userEquipBag[newEquip.EquipIndex]
	} else {
		switch equipPos {
		case pb.EQUIPTYPE_RING:
			tmpEquip = heroEquips[pb.EQUIPPOS_SEVEN]
		case pb.EQUIPTYPE_BRACELET:
			tmpEquip = heroEquips[pb.EQUIPPOS_EIGHT]
		}
	}
	switch {
	case equipPos == pb.EQUIPTYPE_RING && this.EquipCompareCombat(heroJob, tmpEquip, heroEquips[pb.EQUIPPOS_NINE]):
		equipPos = pb.EQUIPPOS_NINE
	case equipPos == pb.EQUIPTYPE_BRACELET && this.EquipCompareCombat(heroJob, tmpEquip, heroEquips[pb.EQUIPPOS_TEN]):
		equipPos = pb.EQUIPPOS_TEN
	}
	tmpEquip = heroEquips[equipPos]
	return equipPos, tmpEquip
}

/**
 *  @Description: 比较两个装备战斗力
 *  @param job
 *  @param equip1
 *  @param equip2
 *  @return bool	1好返回true,2好返回false
 */
func (this *EquipManager) EquipCompareCombat(job int, equip1, equip2 *model.Equip) bool {
	if equip1 == nil {
		return false
	}
	if equip2 == nil {
		return true
	}
	equip1ItemId, equip2ItemId := equip1.ItemId, equip2.ItemId
	equip1Conf := gamedb.GetEquipEquipCfg(equip1ItemId)
	if equip1Conf == nil {
		return false
	}
	equip2Conf := gamedb.GetEquipEquipCfg(equip2ItemId)
	if equip2Conf == nil {
		return true
	}
	equip1PropMap := make(map[int]int)
	for pid, pVal := range equip1Conf.Properties {
		equip1PropMap[pid] += pVal
	}
	for pid, pVal := range equip1Conf.PropertiesStar {
		equip1PropMap[pid] += pVal
	}
	equip1Combat := builder.CalcCombat(job, equip1ItemId, -1, equip1PropMap, equip1.RandProps)
	equip2PropMap := make(map[int]int)
	for pid, pVal := range equip2Conf.Properties {
		equip2PropMap[pid] += pVal
	}
	for pid, pVal := range equip2Conf.PropertiesStar {
		equip2PropMap[pid] += pVal
	}
	equip2Combat := builder.CalcCombat(job, equip2ItemId, -1, equip2PropMap, equip2.RandProps)
	if equip1Combat > equip2Combat {
		return true
	}
	return false
}

func (this *EquipManager) checkWearCondition(user *objs.User, heroIndex, equipType, itemId int) error {
	equipT := gamedb.GetEquipEquipCfg(itemId)
	if equipT == nil {
		return gamedb.ERRSETTINGNOTFOUND
	}
	if equipT.Type != equipType {
		return gamedb.ERREQUIPTYPE
	}
	if isOk := this.GetCondition().CheckMulti(user, heroIndex, equipT.Condition); !isOk {
		return gamedb.ERRCONDITION
	}
	return nil
}
