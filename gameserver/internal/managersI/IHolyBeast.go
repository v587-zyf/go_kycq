package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IHolyBeastManager interface {
	Load(user *objs.User, ack *pb.HolyBeastLoadInfoAck)
	//圣兽激活
	Activate(user *objs.User, heroIndex, types int, ack *pb.HolyBeastActivateAck, op *ophelper.OpBagHelperDefault) error

	//圣兽升星
	UpStar(user *objs.User, heroIndex, types int, ack *pb.HolyBeastUpStarAck, op *ophelper.OpBagHelperDefault) error

	//特殊星  特殊属性选择
	ChooseProp(user *objs.User, heroIndex, types, chooseIndex int, ack *pb.HolyBeastChoosePropAck, op *ophelper.OpBagHelperDefault) error

	Rest(user *objs.User, heroIndex, types int, ack *pb.HolyBeastRestAck, op *ophelper.OpBagHelperDefault) error
}
