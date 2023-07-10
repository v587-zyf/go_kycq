package rank

import (
	"cqserver/gamelibs/publicCon/constRank"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/base"
)

func (this *RankManager) GetKey(t int) string {
	return rmodel.Rank.GetRankKey(t, base.Conf.ServerId)
}

func (this *RankManager) GetCount(rankType int) int {
	key := rmodel.Rank.GetRankKey(rankType, base.Conf.ServerId)
	return rmodel.Rank.GetRankNum(key)
}

func (this *RankManager) LoadRank(rankType, count int) []int {
	key := rmodel.Rank.GetRankKey(rankType, base.Conf.ServerId)
	ranks := rmodel.Rank.GetRank(key, count)
	if constRank.SORT_TIME_RANK[rankType] {
		if ranks != nil && len(ranks) > 0 {
			for i := 1; i < len(ranks); i += 2 {
				ranks[i] = ranks[i] / constRank.RANK_SORT_TIME_FIX
			}
		}
	}
	return ranks
}

func (this *RankManager) GetRanking(rankType, id int) int {
	key := rmodel.Rank.GetRankKey(rankType, base.Conf.ServerId)
	return rmodel.Rank.GetSelfRank(key, id)
}

func (this *RankManager) GetRankByScore(rankType int, start int, end int) []int {
	return rmodel.Rank.ZrangeBuyScore(rankType, base.Conf.ServerId, start, end)
}

func (this *RankManager) GetRankScore(rankType, userId int) int {
	key := rmodel.Rank.GetRankKey(rankType, base.Conf.ServerId)
	score, _ := rmodel.Rank.GetSelfScore(key, userId)
	return score
}

func (this *RankManager) GetMyRank(userId int, rankType int) int {
	rankKey := rmodel.Rank.GetRankKey(rankType, base.Conf.ServerId)
	myRank := rmodel.Rank.GetSelfRank(rankKey, userId)
	return myRank
}