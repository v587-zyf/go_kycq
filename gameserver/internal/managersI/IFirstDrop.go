package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IFirstDropManager interface {
	CheckIsFirstDrop(user *objs.User, itemIds map[int]int)

	LoadInfo(user *objs.User, types int, ack *pb.FirstDropLoadAck)

	GetAward(user *objs.User, id int, ack *pb.GetFirstDropAwardAck, op *ophelper.OpBagHelperDefault) error

	GetAllAward(user *objs.User, types int, ack *pb.GetAllFirstDropAwardAck, op *ophelper.OpBagHelperDefault) error

	GetAllRedPacketAward(user *objs.User, infos []int32, op *ophelper.OpBagHelperDefault) error

	Reset(user *objs.User, data int)
}
