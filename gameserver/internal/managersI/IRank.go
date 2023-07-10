package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
)

type IRank interface {
	/**
	 *  @Description:  添加到排行版
	 *  @param rankType	排行榜类型
	 *  @param member 成员（一般为userId）
	 *  @param score  成员积分
	 */
	Append(rankType int, member interface{}, score int, appendNow, isLogin, combatUp bool)
	// 获取排行榜key
	GetKey(t int) string
	// 获取榜总数
	GetCount(key int) int
	// 获取榜单
	LoadRank(key, count int) []int
	// 获取榜单排名
	GetRanking(rankType, id int) int
	// 根据索引获取区间（分数由高到低排序）
	GetRankByScore(rankType int, start int, end int) []int
	// 获取积分
	GetRankScore(key, id int) int
	// 删除 redis 中的 del
	DelData(rankType int)
	//膜拜奖励
	WorshipReward(user *objs.User, op *ophelper.OpBagHelperDefault) error
    //用于排行榜分数后 加 时间戳
	GetUserJointCombat(combat interface{}) float64
}
