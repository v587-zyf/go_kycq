package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IAppletsManager interface {
	EnterAppletsReq(user *objs.User, appletsType int, ntf *pb.AppletsEnergyNtf) error

	//领取魔法射击  杀怪奖励
	AppletsReceiveReq(user *objs.User, receiveId int, ack *pb.AppletsReceiveAck, op *ophelper.OpBagHelperDefault) error

	//魔法射击定时奖励获取
	CronGetAwardReq(user *objs.User, id, index int, ack *pb.CronGetAwardAck, op *ophelper.OpBagHelperDefault) error

	//通关结算奖励
	EndResultReq(user *objs.User, appletsType, id int, ack *pb.EndResultAck, op *ophelper.OpBagHelperDefault) error

	CronAddPhysicalPower(user *objs.User)

	OnlineAddPhysicalPower(user *objs.User)
}
