package godEquip

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

type GodEquipManager struct {
	util.DefaultModule
	managersI.IModule
}

func NewGodEquipManager(module managersI.IModule) *GodEquipManager {
	return &GodEquipManager{IModule: module}
}

/**
 *  @Description: 神兵激活
 *  @param user
 *  @param heroIndex
 *  @param id		神兵id
 *  @param op
 *  @return error
 */
func (this *GodEquipManager) GodEquipActive(user *objs.User, heroIndex, id int, op *ophelper.OpBagHelperDefault) error {
	//hero := user.Heros[heroIndex]
	//if hero == nil {
	//	return gamedb.ERRHERONOTFOUND
	//}
	//
	//if _, ok := hero.GodEquips[id]; ok {
	//	return gamedb.ERRREPEATACTIVE
	//}
	//
	//godEquipConf := gamedb.GetGodEquipConfCfg(id)
	//if godEquipConf == nil {
	//	return gamedb.ERRSETTINGNOTFOUND
	//}
	////扣材料
	//err := this.GetBag().Remove(user, op, godEquipConf.Consume.ItemId, godEquipConf.Consume.Count)
	//if err != nil {
	//	return err
	//}
	////激活
	//hero.GodEquips[id] = &model.GodEquip{
	//	Id: id,
	//	Lv: 1,
	//}
	//user.Dirty = true
	//this.GetUserManager().UpdateCombat(user, heroIndex)
	return nil
}

/**
 *  @Description: 神兵激活丶升级
 *  @param user
 *  @param heroIndex
 *  @param id	神兵id
 *  @param op
 *  @return error
 */
func (this *GodEquipManager) GodEquipUpLevel(user *objs.User, heroIndex, id int, op *ophelper.OpBagHelperDefault) error {
	if id < 1 {
		return gamedb.ERRPARAM
	}
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	godEquip, ok := hero.GodEquips[id]
	if !ok {
		hero.GodEquips[id] = &model.GodEquip{Id: id}
		godEquip = hero.GodEquips[id]
	}
	if gamedb.GetGodEquipLevelConfCfg(gamedb.GetRealId(godEquip.Id, godEquip.Lv+1)) == nil {
		return gamedb.ERRLVENOUGH
	}
	godEquipConf := gamedb.GetGodEquipLevelConfCfg(gamedb.GetRealId(godEquip.Id, godEquip.Lv))
	if check := this.GetCondition().CheckMulti(user, heroIndex, godEquipConf.Condition); !check {
		return gamedb.ERRCONDITION
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, godEquipConf.Consume); err != nil {
		return err
	}
	godEquip.Lv += 1

	kyEvent.ShengBinUp(user, heroIndex, id, godEquip.Lv-1, godEquip.Lv)

	this.GetUserManager().UpdateCombat(user, heroIndex)
	this.GetTask().AddTaskProcess(user, pb.CONDITION_UPGRADE_SHEN_BIN, 1)
	this.GetTask().AddTaskProcess(user, pb.CONDITION_UPGRADE_SHEN_BIN_1, -1)
	this.GetCondition().RecordCondition(user, pb.CONDITION_UPGRADE_SHEN_BIN_1, []int{})
	return nil
}
