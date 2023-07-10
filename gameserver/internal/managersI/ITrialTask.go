package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type ITrialTaskManager interface {
	TrialTaskLoad(user *objs.User, ack *pb.TrialTaskInfoAck)

	//某个任务完成奖励
	GetTrialTaskAward(user *objs.User, id int, ack *pb.TrialTaskGetAwardAck, op *ophelper.OpBagHelperDefault) error

	//获取阶段奖励
	GetStageAward(user *objs.User, id int, ack *pb.TrialTaskGetStageAwardAck, op *ophelper.OpBagHelperDefault) error

	OfflineSaveTrialTaskInfo(user *objs.User)

	SendTrialTaskInfoNtf(user *objs.User, conditionType int)

	SendReward()

	OnlineCheck(user *objs.User)
}
