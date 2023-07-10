package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
	"time"
)

type IChallengeManager interface {
	//跨服擂台赛报名
	SetApplyUserInfo(user *objs.User, ack *pb.ApplyChallengeAck) error

	//下注    bottomUser:投注人 userId
	BottomPour(user *objs.User, bottomUser int, op *ophelper.OpBagHelperDefault, ack *pb.BottomPourAck) error

	//当前轮参加的玩家
	EachRoundPeople(user *objs.User, ack *pb.ChallengeEachRoundPeopleAck) error

	//每轮结束 给失败者 发奖
	SendLoseReward(loseUserIds, winUserIds []int32, roundIndex, winUserId int)

	LoadInfo(serverId int, ack *pb.ChallengeInfoAck) error

	//报名红点判断
	IsApplyChallenge(user *objs.User) int32

	WeekByDate(t time.Time) string

	//擂台赛 开始,结束,报名结束 广播
	Broadcast()
}
