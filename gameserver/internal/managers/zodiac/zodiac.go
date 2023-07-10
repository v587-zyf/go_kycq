package zodiac

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

func NewZodiacManager(module managersI.IModule) *ZodiacManager {
	return &ZodiacManager{IModule: module}
}

type ZodiacManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *ZodiacManager) ZodiacChange(user *objs.User, heroIndex, pos, bagPos int, op *ophelper.OpBagHelperDefault) (error, *pb.SpecialEquipUnit) {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND, nil
	}
	if !pb.ZODIACTYPE_MAP[pos] {
		return gamedb.ERRPARAM, nil
	}
	item := this.GetBag().GetItemByPosition(user, bagPos)
	if item == nil || item.ItemId == 0 {
		return gamedb.ERRNOTENOUGHGOODS, nil
	}
	itemId := item.ItemId

	conf := gamedb.GetZodiacEquipZodiacEquipCfg(itemId)
	check := this.GetCondition().CheckMulti(user, heroIndex, conf.Condition)
	if !check {
		return gamedb.ERRCONDITION, nil
	}
	if conf.Type != pos {
		return gamedb.ERREQUIPTYPE, nil
	}
	equips := hero.Zodiacs
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
	user.UpdateFightUserHeroIndexFun(heroIndex)
	return nil, builder.BuilderSpecialEquipUnit(equips[pos])
}

func (this *ZodiacManager) ZodiacRemove(user *objs.User, heroIndex, pos int, op *ophelper.OpBagHelperDefault) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}
	equips := hero.Zodiacs
	if !pb.ZODIACTYPE_MAP[pos] {
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
	user.UpdateFightUserHeroIndexFun(heroIndex)
	return nil
}
