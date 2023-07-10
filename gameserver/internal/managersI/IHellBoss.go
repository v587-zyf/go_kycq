package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IHellBoss interface {
	Online(user *objs.User)
	ResetHellBoss(user *objs.User, date int)
	Load(user *objs.User, floor int) []*pb.HellBossNtf
	BuyNum(user *objs.User, use bool, buyNum int, op *ophelper.OpBagHelperDefault) error

	EnterHellBossFight(user *objs.User, stageId, helpUserId int) error
	HellBossFightResult(user *objs.User, stageId, winUserId int, items map[int]int, toHelpUserId int)

	GetSurplusNum(user *objs.User) int
}
