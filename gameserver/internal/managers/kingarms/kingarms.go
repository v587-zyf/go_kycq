package kingarms

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

func NewKingarmsManager(module managersI.IModule) *KingarmsManager {
	return &KingarmsManager{IModule: module}
}

type KingarmsManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *KingarmsManager) KingarmsChange(user *objs.User, heroIndex, pos, bagPos int, op *ophelper.OpBagHelperDefault) (error, *pb.SpecialEquipUnit) {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND, nil
	}
	if !pb.KINGARMSTYPE_MAP[pos] {
		return gamedb.ERRPARAM, nil
	}
	item := this.GetBag().GetItemByPosition(user, bagPos)
	if item == nil || item.ItemId == 0 {
		return gamedb.ERRNOTENOUGHGOODS, nil
	}
	itemId := item.ItemId

	conf := gamedb.GetKingarmsKingarmsCfg(itemId)
	if conf == nil {
		logger.Debug("kingarms error itemId is %v", itemId)
		return gamedb.ERRBAGPOSITION, nil
	}
	check := this.GetCondition().CheckMulti(user, heroIndex, conf.Condition)
	if !check {
		return gamedb.ERRCONDITION, nil
	}
	if conf.Type != pos {
		return gamedb.ERREQUIPTYPE, nil
	}
	equips := hero.Kingarms
	err := this.GetBag().Remove(user, op, itemId, 1)
	if err != nil {
		return err, nil
	}
	if equips[pos] != nil && equips[pos].Id != 0 {
		this.GetBag().AddItem(user, op, equips[pos].Id, 1)
	}
	equips[pos] = &model.SpecialEquipUnit{
		Id: itemId,
	}

	this.GetUserManager().UpdateCombat(user, heroIndex)
	return nil, builder.BuilderSpecialEquipUnit(equips[pos])
}

func (this *KingarmsManager) KingarmsRemove(user *objs.User, heroIndex, pos int, op *ophelper.OpBagHelperDefault) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	equips := hero.Kingarms
	if !pb.KINGARMSTYPE_MAP[pos] {
		return gamedb.ERRPARAM
	}
	if equips[pos].Id != 0 {
		err := this.GetBag().AddItem(user, op, equips[pos].Id, 1)
		if err != nil {
			return err
		}
	}
	equips[pos] = &model.SpecialEquipUnit{}
	this.GetUserManager().UpdateCombat(user, heroIndex)
	return nil
}
