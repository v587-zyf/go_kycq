package rmodel

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/golibs/logger"
	"fmt"
	"time"
)

const (
	CompetitiveSeasonRankInfo          = "competitve_season_rank_info:%v"            //season 竞技场赛季的排名 过期时间14天
	CompetitiveSeasonRankInfoByOpenDay = "competitve_season_rank_info_by_openDay:%v" //season 竞技场赛季的排名 过期时间14天
	CompetitiveSeasonSendRewardMark    = "competitve_season_send_reward_mark:%v"     //serverid
	CompetitiveSeasonMarkUserIds       = "competitve_season_mark_user_ids"
)

type CompetitveModel struct {
}

var Competitve = &CompetitveModel{}

//赛季排名 key
func (this *CompetitveModel) GetZAddSeasonRankKey(season int) string {
	return fmt.Sprintf(CompetitiveSeasonRankInfo, season)
}

func (this *CompetitveModel) GetZAddSeasonRankKeyByOpenDay(openDay int) string {
	return fmt.Sprintf(CompetitiveSeasonRankInfoByOpenDay, openDay)
}

func (this *CompetitveModel) CompetitiveSeasonSendRewardMark(serverId int) string {
	return fmt.Sprintf(CompetitiveSeasonSendRewardMark, serverId)
}

func (this *CompetitveModel) CompetitiveSeasonMarkUserIds() string {
	return fmt.Sprintf(CompetitiveSeasonMarkUserIds)
}

//------- 赛季排名处理

func (this *CompetitveModel) ZAddSeasonRankInfo(season, openDay int, id, value interface{}) {
	key := this.GetZAddSeasonRankKey(season)
	key1 := this.GetZAddSeasonRankKeyByOpenDay(openDay)
	redisDb.ZAdd(key, value, id)
	redisDb.Expire(key, 14*24*time.Hour)
	redisDb.ZAdd(key1, value, id)
	redisDb.Expire(key1, 14*24*time.Hour)
}

//获取赛季排名信息
func (this *CompetitveModel) GetSeasonRankInfos(season, num int) []float64 {
	key := this.GetZAddSeasonRankKey(season)
	ranks, err := redisDb.ZRevrange(key, 0, num).ValuesFloat64Slice()
	if err != nil {
		logger.Error("GetSeasonRankInfos获取排行榜数据异常,key:%v,err:%v", key, err)
	}
	return ranks
}

//获取赛季排名信息
func (this *CompetitveModel) GetSeasonRankInfosBuyScoreSection(season, min, max int) ([]float64, error) {
	key := this.GetZAddSeasonRankKey(season)
	ranks, err := redisDb.ZRangeByScore(key, min, max).ValuesFloat64Slice()
	if err != nil {
		logger.Error("GetSeasonRankInfos获取排行榜数据异常,key:%v,err:%v", key, err)
	}
	return ranks, err
}

//获取赛季玩家的排名和积分
func (this *CompetitveModel) GetSeasonSelfRankAndScore(userId, season int) (int, float64) {
	key := this.GetZAddSeasonRankKey(season)
	rank := redisDb.ZRevrank(key, userId)
	if rank < 0 {
		return rank, 0
	}
	score, err := redisDb.ZScoreByFloat(key, userId)
	if err != nil {
		fmt.Printf(" GetSeasonSelfRankAndScore  err:%v", err)
	}
	return rank, score
}

//根据开服天数获取竞技场分数
func (this *CompetitveModel) GetCompetitiveScoreByOpenDay(userId, openDay int) int {

	var err error
	seasonDay := gamedb.GetConf().CompetitveSeason
	score := float64(0)
	num := 0

	before := openDay / seasonDay
	after := openDay % seasonDay

	seasonNowDay := 0
	seasonNowDay = after
	if before <= 0 {
		seasonNowDay = after
	}

	if after == 0 {
		seasonNowDay = seasonDay
	}

	if seasonNowDay != seasonDay {
		seasonDay = seasonNowDay
	}
	for i := openDay; i > 0; i-- {
		num++
		key := this.GetZAddSeasonRankKeyByOpenDay(i)
		score, err = redisDb.ZScoreByFloat(key, userId)
		if err != nil {
			fmt.Printf(" GetSeasonSelfRankAndScore  err:%v", err)
		}
		if score > 0 || num >= seasonDay {
			break
		}
	}

	return int(score)
}

//-----服务器启动 删除user表不存在的玩家 redis中的数据
func (this *CompetitveModel) ZAddSeasonRank(season int, id, value interface{}) {
	key := this.GetZAddSeasonRankKey(season)
	redisDb.ZAdd(key, value, id)
	redisDb.Expire(key, 14*24*time.Hour)
}

func (this *CompetitveModel) DelSeasonRank(season int) {
	key := this.GetZAddSeasonRankKey(season)
	redisDb.Del(key)
}

//竞技场赛季发奖记录
func (this *CompetitveModel) SetCompetitiveSeasonSendRewardMark(serverId, season int) {
	key := this.CompetitiveSeasonSendRewardMark(serverId)
	redisDb.Set(key, season)
}

func (this *CompetitveModel) GetCompetitiveSeasonSendRewardMark(serverId int) int {
	key := this.CompetitiveSeasonSendRewardMark(serverId)
	num, _ := redisDb.Get(key).IntDef(0)
	return num
}

func (this *CompetitveModel) SetLastMarkUserId(uerId, markUserId int) {
	key := this.CompetitiveSeasonMarkUserIds()
	redisDb.Hmset(key, uerId, markUserId)
}

func (this *CompetitveModel) GetLastMarkUserId(uerId int) int {
	key := this.CompetitiveSeasonMarkUserIds()
	num, _ := redisDb.HgetIntDef(key, uerId, 0)
	return num
}
