package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pbserver"
)

type IDaBaoManager interface {
	Online(user *objs.User)

	UpEquip(user *objs.User, equipT int, op *ophelper.OpBagHelperDefault) error

	EnterMystery(user *objs.User, stageId int) error
	//MysteryResult(user *objs.User, stageId int, isWin bool, items map[int]int)
	ResumeEnergy(user *objs.User)
	SyncEnergy(user *objs.User, changeNum int)
	EnergyItemUse(user *objs.User, itemId int, op *ophelper.OpBagHelperDefault) error

	//打宝秘境广播
	SendSystemDropItem(user *objs.User, stageId int, replyItems map[int32]*pbserver.ItemUnitForPickUp)
}
