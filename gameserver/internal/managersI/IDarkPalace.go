package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IDarkPalaceManager interface {
	Online(user *objs.User)
	ResetDarkPalace(user *objs.User, date int)
	Load(floor int, ack *pb.DarkPalaceLoadAck)
	BuyNum(user *objs.User, use bool, buyNum int, op *ophelper.OpBagHelperDefault, ack *pb.DarkPalaceBuyNumAck) error
	SendDarkPalaceBossNtf(darkPalaceBossNtf *pb.DarkPalaceBossNtf)
	GetSurplusNum(user *objs.User) int

	EnterDarkPalaceFight(user *objs.User, stageId int, helpUserId int) error
	DarkPalaceFightResultNtf(user *objs.User, stageId, winUserId int, items map[int]int, toHelpUserId int)
}
