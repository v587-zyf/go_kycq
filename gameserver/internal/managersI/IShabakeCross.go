package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
)

type IShaBaKeCrossManager interface {
	LoadCross(user *objs.User, ack *pb.ShaBaKeInfoCrossAck) error

	//判断沙巴克是否开启
	JudgeCrossIsOpen(user *objs.User) error

	//进入跨服沙巴克
	EnterCrossShaBakeFight(user *objs.User) error

	//战斗结束回调
	CrossShaBakeFightEndNtf(msg *pbserver.ShabakeCrossFightEndNtf)

	//开启推送
	CrossShaBakeOpenOrCloseNtf(isOpen bool) error

	SendFirstGuildInfoToCcs(guildId, benFuShaBake int)

	//设置第一门派的信息
	SetFirstGuildInfo(msg *pbserver.CcsToGsBroadShaBakeFirstGuildInfo)

	BuildBackGuildInfo(users []int) []*pbserver.Info
}
