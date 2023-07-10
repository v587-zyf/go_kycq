package rmodel

import (
	"cqserver/golibs/logger"
	"fmt"
	"time"
)

const (
	USER_RANK_KEY       = "user_rank_%d_%d"       //玩家排行榜 (rankType,serverId) hash,{userId,scroe}
	ARENA_RANK_FIGHT    = "arena_rank_fight:%v"   //竞技场排名战斗中
	DAILY_RANK_SAVE     = "daily_rank_save:%v:%v" //openDay:type
	DAILY_USER_RANK_KEY = "daily_user_rank_%d_%d" //每日排行榜玩家排行榜 (rankType,serverId) hash,{userId,scroe} 分数后加时间错
	RANK_REWARD         = "rank_reward_%d"        //排行榜奖励是否领取（rankType）目前只有试炼塔排行奖励如果0点停服，开服补发
)

type RankModel struct {
}

var Rank = &RankModel{}

func GetDailyRankSaveKey(openDay, types int) string {
	return fmt.Sprintf(DAILY_RANK_SAVE, openDay, types)
}

func (this *RankModel) SetDailyRankSave(openDay, types int, data string) {
	key := GetDailyRankSaveKey(openDay, types)
	redisDb.SetWithExpire(key, data, 24*7*time.Hour)
}

func (this *RankModel) GetDailyRankSave(openDay, types int) (string, error) {
	key := GetDailyRankSaveKey(openDay, types)
	return redisDb.Get(key).String()
}

func (this *RankModel) GetRankNum(key string) int {
	return redisDb.ZCard(key)
}

func (this *RankModel) ZAddRank(key, id, value interface{}) {
	redisDb.ZAdd(key, value, id)
}

func (this *RankModel) GetSelfRank(key string, id int) int {
	return redisDb.ZRevrank(key, id)
}

func (this *RankModel) GetSelfScore(key string, id int) (int, error) {
	return redisDb.ZScore(key, id)
}

func (this *RankModel) GetRank(key string, num int) []int {
	ranks, err := redisDb.ZRevrange(key, 0, num).ValuesIntSlice()
	if err != nil {
		logger.Error("获取排行榜数据异常,key:%v,err:%v", key, err)
	}
	return ranks
}

func (this *RankModel) GetDailyRank(key string, num int) []float64 {
	ranks, err := redisDb.ZRevrange(key, 0, num).ValuesFloat64Slice()
	if err != nil {
		logger.Error("获取每日排行榜数据异常,key:%v,err:%v", key, err)
	}
	return ranks
}

func (this *RankModel) GetSelfRankAndScore(key string, id int) (int, int) {
	rank := redisDb.ZRevrank(key, id)
	if rank < 0 {
		return rank, 0
	}
	score, err := redisDb.ZScore(key, id)
	if err != nil {
		fmt.Printf(" GetSelfRankAndScore  err:%v", err)
	}
	return rank, score
}

func (this *RankModel) GetDailyRankSelfRankAndScore(key string, id int) float64 {
	score, err := redisDb.ZScoreByFloat(key, id)
	if err != nil {
		fmt.Printf(" GetSelfRankAndScore  err:%v", err)
	}
	return score
}

/**获取玩家相关排行版，玩家前len/2名和后len/2名*/
func (this *RankModel) GetRankByUserFun(rankKey string, userId int, length int) (int, []int) {
	myNum := this.GetSelfRank(rankKey, userId)
	allRankLen := redisDb.ZCard(rankKey)
	cutLen := int(float32(length) * 0.5)
	preSelNum, afterSelNum := 0, 0
	if myNum >= cutLen && allRankLen-myNum-1 >= cutLen {
		preSelNum = cutLen
		afterSelNum = cutLen
	} else if myNum >= cutLen && allRankLen-myNum-1 < cutLen {

		afterSelNum = allRankLen - myNum - 1
		preSelNum = length - afterSelNum

	} else {
		preSelNum = myNum
		afterSelNum = length - preSelNum
	}
	begining := myNum - preSelNum
	if begining < 0 {
		begining = 0
	}
	rankGuilds, _ := redisDb.ZRevrange(rankKey, begining, myNum+afterSelNum).ValuesIntSlice()
	return begining, rankGuilds
}

/**获取玩家相关排行版，玩家前len名*/
func (this *RankModel) GetRankByUserFun2(rankKey string, userId int, length int) (int, int, []int) {

	myNum := this.GetSelfRank(rankKey, userId)
	begining, ending := 0, 0
	if myNum >= length {
		begining = myNum - length
		ending = myNum
	} else {
		begining = 0
		ending = length
	}
	rankGuilds, _ := redisDb.ZRevrange(rankKey, begining, ending).ValuesIntSlice()
	return begining, ending, rankGuilds
}

/**
 * 写入排行榜数据
 */
func (this *RankModel) ZaddRankByType(rankType, serverId int, id, value interface{}) {
	key := this.GetRankKey(rankType, serverId)
	if key != "" {
		this.ZAddRank(key, id, value)
	}
}

/**
 * 写入每日排行榜数据
 */
func (this *RankModel) ZAddDailyRankByType(rankType, serverId int, id, value interface{}) {
	key := this.GetDailyRankKey(rankType, serverId)
	if key != "" {
		this.ZAddRank(key, id, value)
	}
}

// 获取排名信息 分数段
func (this *RankModel) GetRankBuyTypeAndScoreSection(rankType, serverId int, min, max interface{}) ([]int, error) {
	key := this.GetRankKey(rankType, serverId)
	ranks, err := redisDb.ZRangeByScore(key, min, max).ValuesIntSlice()
	if err != nil {
		logger.Error("GetRankBuyTypeAndScoreSection,key:%v,err:%v", key, err)
	}
	return ranks, err
}

// 获取分数区间
func (this *RankModel) ZrangeBuyScore(rankType, serverId, start, end int) []int {
	rankKey := this.GetRankKey(rankType, serverId)
	if rankKey != "" {
		rankGuilds, _ := redisDb.ZRevrange(rankKey, start, end).ValuesIntSlice()
		return rankGuilds
	}
	return nil
}

func (this *RankModel) GetRankKey(rankType int, serverId int) string {
	return fmt.Sprintf(USER_RANK_KEY, rankType, serverId)
}

func (this *RankModel) GetDailyRankKey(rankType int, serverId int) string {
	return fmt.Sprintf(DAILY_USER_RANK_KEY, rankType, serverId)
}

func (this *RankModel) SetArenaRankFight(rank int, fightUserId int) {

	key := fmt.Sprintf(ARENA_RANK_FIGHT, rank)
	redisDb.SetWithExpire(key, fightUserId, 20*time.Minute)
}

func (this *RankModel) GetArenaRankFight(rank int) (int, error) {

	key := fmt.Sprintf(ARENA_RANK_FIGHT, rank)
	return redisDb.Get(key).Int()
}

func (this *RankModel) DelArenaRankFight(rank int) {

	key := fmt.Sprintf(ARENA_RANK_FIGHT, rank)
	redisDb.Del(key)
}

func (this *RankModel) DelData(key string) {
	redisDb.Del(key)
}

func (this *RankModel) GetRankRewardKey(rankType int) string {
	return fmt.Sprintf(RANK_REWARD, rankType)
}
func (this *RankModel) GetRankReward(rankType int) int {
	key := this.GetRankRewardKey(rankType)
	data, _ := redisDb.Get(key).Int()
	return data
}
func (this *RankModel) SetRankReward(rankType int, value int) {
	key := this.GetRankRewardKey(rankType)
	redisDb.SetWithExpire(key, value, AutoExpireTime)
}
func (this *RankModel) DelRankReward(rankType int) {
	key := this.GetRankRewardKey(rankType)
	redisDb.Del(key)
}
