package base

import (
	"cqserver/golibs/common"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"time"
)

type IFightDamageRank interface {
	FightDamageRankSetDamage(actor Actor, value int64)
	FightDamageRankGetRank(isForce bool) []int
	FightDamageRankGetRankInfos(actor Actor) nw.ProtoMessage
	FightDamageRankSetRankInterval(interval int)
	FightDamageRankReset()
}

//
//  @Description: 战斗伤害排行
// 获取排行版时进行排序，可根据参数强制排序，默认间隔1000毫秒
//
type FightDamageRank struct {
	users            map[int]Actor
	damageRank       *common.SortedMap
	lastSortTime     time.Time
	sortRankInterval int64 //排名间隔时间（毫秒）
}

func NewFightDamageRank() *FightDamageRank {
	fightDamageRank := &FightDamageRank{
		users: make(map[int]Actor),
		damageRank: &common.SortedMap{
			M: make(map[int]int64),
		},
		sortRankInterval: 1000,
	}
	return fightDamageRank
}

func (fd *FightDamageRank) FightDamageRankReset() {
	fd.users = make(map[int]Actor)
	fd.damageRank = &common.SortedMap{
		M: make(map[int]int64),
	}
}

func (fd *FightDamageRank) FightDamageRankSetRankInterval(interval int) {
	fd.sortRankInterval = int64(interval)
}

func (fd *FightDamageRank) FightDamageRankSetDamage(actor Actor, value int64) {
	user := fd.users[actor.GetUserId()]
	//玩家第一造成伤害或者玩家退出游戏 重新进入游戏
	if user == nil || user.GetObjId() != actor.GetObjId() {
		fd.users[actor.GetUserId()] = actor
	}
	fd.damageRank.M[actor.GetUserId()] += value
}

func (fd *FightDamageRank) FightDamageRankGetRank(isForce bool) []int {
	if isForce || fd.lastSortTime.IsZero() || (time.Now().Sub(fd.lastSortTime).Milliseconds()) > fd.sortRankInterval {

		return fd.damageRank.Sort()
	}
	return fd.damageRank.GetNowRank()
}

func (fd *FightDamageRank) FightDamageRankGetRankInfos(actor Actor) *pb.FightHurtRankAck {

	rank := fd.FightDamageRankGetRank(false)
	ack := &pb.FightHurtRankAck{
		Ranks: make([]*pb.FightRankUnit, len(rank)),
	}
	for k, v := range rank {
		ack.Ranks[k] = &pb.FightRankUnit{
			Rank:   int32(k + 1),
			Name:   fd.users[v].NickName(),
			Score:  fd.damageRank.M[v],
			UserId: int32(fd.users[v].GetUserId()),
		}
		if actor != nil && fd.users[v].GetUserId() == actor.GetUserId() {
			ack.MyRank = ack.Ranks[k]
		}
	}
	return ack
}
