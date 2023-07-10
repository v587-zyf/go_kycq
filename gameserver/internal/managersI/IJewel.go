package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IJewelManager interface {
	// 镶嵌宝石
	JewelMake(user *objs.User, heroIndex, equipPos, jewelPos, itemId int, op *ophelper.OpBagHelperDefault, ack *pb.JewelMakeAck) error
	// 宝石升级
	JewelUpLv(user *objs.User, heroIndex, equipPos, jewelPos int, op *ophelper.OpBagHelperDefault, ack *pb.JewelUpLvAck) error
	// 宝石替换
	JewelChange(user *objs.User, heroIndex, equipPos, jewelPos, itemId int, op *ophelper.OpBagHelperDefault, ack *pb.JewelChangeAck) error
	// 宝石卸下
	JewelRemove(user *objs.User, heroIndex, equipPos, jewelPos int, op *ophelper.OpBagHelperDefault, ack *pb.JewelRemoveAck) error
	// 宝石一键镶嵌
	JewelMakeAll(user *objs.User, heroIndex int, op *ophelper.OpBagHelperDefault, ack *pb.JewelMakeAllAck) error
}
