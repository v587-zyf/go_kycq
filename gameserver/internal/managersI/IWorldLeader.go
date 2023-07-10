package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
)

type IWorldLeaderManager interface {
	LoadWorldLeader(user *objs.User, ack *pb.LoadWorldLeaderAck) error

	//挑战首领
	WorldLeaderEnter(user *objs.User, stageId int, ack *pb.WorldLeaderEnterAck) error

	//伤害排名信息
	WorldLeaderRankInfo(user *objs.User, stageId int, ack *pb.GetWorldLeaderRankInfoAck) error

	//推送挑战开始
	SendClientWorldLeaderStart(stageId int)

	//推送世界首领排行
	WorldLeaderFightRankNtf(msg *pbserver.WorldLeaderFightRankNtf)

	//推送世界首领挑战结果
	EndWorldLeaderBossNtf(msg *pbserver.WorldLeaderFightEndNtf)

	Reset()
}
