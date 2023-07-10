package arena

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func (this *ArenaManager) AddArenaRank(user *objs.User) error {
	err := this.checkArena(user, false)
	if err != nil {
		return err
	}
	userArenaRanking := this.GetRank().GetRanking(ArenaRankType, user.Id)
	if userArenaRanking < 0 {
		count := this.GetRank().GetCount(pb.RANKTYPE_ARENA)
		this.GetRank().Append(ArenaRankType, user.Id, this.rankScore(count), false, false, false)
	}
	return nil
}

func (this *ArenaManager) GetPbThree() []*pb.RankInfo {
	pbArr := make([]*pb.RankInfo, 0)
	// 获取前三名
	ranks := this.GetRank().GetRankByScore(ArenaRankType, 0, 2)

	for i, l := 0, len(ranks); i < l; i += 2 {
		rankUserId := ranks[i]
		rankScore := ranks[i+1]
		pbArr = append(pbArr, this.GetUserManager().BuildUserRankInfo(rankUserId, -1, (i/2)+1, rankScore))
	}
	return pbArr
}
