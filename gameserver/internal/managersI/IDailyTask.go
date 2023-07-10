package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IDailyTaskManager interface {
	DailyTaskLoad(user *objs.User, ack *pb.DailyTaskLoadAck) error

	//购买每天任务的 挑战次数
	BuyChallengeTime(user *objs.User, activityId int, ack *pb.BuyChallengeTimeAck, op *ophelper.OpBagHelperDefault) error

	//每日任务 领取周或者日 阶段奖励
	GetReward(user *objs.User, id, types int, ack *pb.GetAwardAck, op *ophelper.OpBagHelperDefault) error

	//资源找回
	GetResourcesBackReward(user *objs.User, useIngot, activityId, backTimes int, ack *pb.ResourcesBackGetRewardAck, op *ophelper.OpBagHelperDefault) error

	//任务完成通知
	CompletionOfTask(user *objs.User, types, times int)

	InfoInit(user *objs.User)

	Reset(user *objs.User, date int, isReset bool)

	OnLine(user *objs.User)

	SendReward()

	GetDiffDaysBySecond(t1, t2 int64) int

	OfflineSaveDailyTaskInfo(user *objs.User)

	//一键资源找回
	GetResourcesBackAllReward(user *objs.User, ack *pb.ResourcesBackGetAllRewardAck, op *ophelper.OpBagHelperDefault) error
}
