package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IArenaManager interface {
	Online(user *objs.User)
	// 每日重置
	ResetArena(user *objs.User)
	// 打开页面
	ArenaOpen(user *objs.User, ack *pb.ArenaOpenAck) error
	// 进入战斗
	EnterArenaFight(user *objs.User, challengeUid, challengeRanking int) error
	// 战斗回调
	ArenaFightResult(user *objs.User, result bool, challengeUid int)
	// 购买次数
	BuyArenaFightNum(user *objs.User, op *ophelper.OpBagHelperDefault, ack *pb.BuyArenaFightNumAck) error
	// 刷新对手
	RefArenaRank(user *objs.User, ack *pb.RefArenaRankAck) error
	// 加入榜单
	AddArenaRank(user *objs.User) error

	LeaveFight(challengeRanking int)
}
