package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IAchievementManager interface {
	Load(user *objs.User, ack *pb.AchievementLoadAck) error

	//领取成就奖励
	GetAward(user *objs.User, id []int32, ack *pb.AchievementGetAwardAck, op *ophelper.OpBagHelperDefault) error

	//激活成就徽章
	ActiveMedal(user *objs.User, id int, ack *pb.ActiveMedalAck) error

	Online(user *objs.User)

	//成就任务完成 通知
	AddAchievementTaskProcess(user *objs.User, conditionType, num int)

	//通知前端下一成就任务
	UpdateAchievementTaskProcess(user *objs.User, conditionType int)
}
