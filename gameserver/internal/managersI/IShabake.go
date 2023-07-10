package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

type IShabakeManager interface {
	Load(user *objs.User, ack *pb.ShaBaKeInfoAck) error

	//判断沙巴克是否开启
	JudgeIsOpen(user *objs.User) error

	//进入沙巴克
	EnterShabakeFight(user *objs.User) error

	//战斗结束回调
	ShabakeFightEndAck(rank, guildRank, otherUser []int) error

	//开启推送
	ShabakeOpenOrCloseNtf(isOpen bool) error
}
