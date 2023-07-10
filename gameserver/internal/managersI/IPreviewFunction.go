package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IPreviewFunctionManager interface {

	Load(user *objs.User, ack *pb.PreviewFunctionLoadAck)


	GetReward(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.PreviewFunctionGetAck) error

	//功能预览 设置点击状态
	SetPointId(user *objs.User, pointId int, ack *pb.PreviewFunctionPointAck) error
}
