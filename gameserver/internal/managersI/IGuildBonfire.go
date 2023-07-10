package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IGuildBonfireManager interface {
	LoadInfo(user *objs.User, ack *pb.GuildBonfireLoadAck) error

	GuildAddExpPercent(user *objs.User, op *ophelper.OpBagHelperDefault, consumptionType int, ack *pb.GuildBonfireAddExpAck) error

	//玩家再篝火圈内 add exp
	AddUserExp(userIds []int)

	//玩家离开篝火圈 停止add exp
	StopAddUserExp(userId int)

	//活动结束时 调用
	StopAllUserAdd()

	//判断门派篝火活动是否开启
	JudgeGuildBonfireIsOpen(user *objs.User) bool

	//进入战斗
	EnterGuildBonfireFight(user *objs.User) error

	//战斗回调
	GuildBonfireFightResult(userIds []int)

	//活动开启or关闭推送
	GuildBonfireIsOpenNtf(isOpen bool)
}
