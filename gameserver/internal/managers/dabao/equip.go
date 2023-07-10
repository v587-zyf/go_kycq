package dabao

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

/**
 *  @Description: 打宝神器升级
 *  @param user
 *  @param equipT
 *  @param op
 *  @return error
 */
func (this *DaBao) UpEquip(user *objs.User, equipT int, op *ophelper.OpBagHelperDefault) error {
	if user.DaBaoEquip == nil {
		user.DaBaoEquip = make(model.IntKv)
	}
	userEquips := user.DaBaoEquip
	if gamedb.GetDaBaoEquipByTypeAndLv(equipT, userEquips[equipT]+1) == nil {
		return gamedb.ERRLVENOUGH
	}
	equipCfg := gamedb.GetDaBaoEquipByTypeAndLv(equipT, userEquips[equipT])
	if equipCfg == nil {
		return gamedb.ERRPARAM
	}
	flag := true
	for _, condition := range equipCfg.Condition {
		if _, check := this.GetCondition().CheckBySlice(user, -1, condition); !check {
			flag = false
			break
		}
	}
	if !flag {
		return gamedb.ERRCONDITION
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, equipCfg.Cost); err != nil {
		return err
	}

	userEquips[equipT]++
	user.Dirty = true

	this.GetUserManager().UpdateCombat(user, -1)
	this.GetTask().UpdateTaskProcess(user, false, false)
	this.GetCondition().RecordCondition(user, pb.CONDITION_DAO_BAO_SHEN_QI_ALL_LV, []int{})
	return nil
}
