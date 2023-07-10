package rmodel

import (
	"cqserver/gamelibs/beans"
	"cqserver/golibs/redisdb"
)

var (
	redisNewDb = &redisdb.RedisDb{}
)

const (
	GuildDailyCondition                = "guild_dai_condi:%s:%d"                     //每日政务的完成情况 hash,{conditionID,num},(date,userId)
	GuildDailyCondition_               = "guild_dai_condi:%s:*"                      //每日政务的完成情况 hash,{conditionID,num},(date,userId)
	CompetitiveSeasonRankInfo          = "competitve_season_rank_info:%v"            //season 竞技场赛季的排名 过期时间14天
	CompetitiveSeasonSendRewardMark    = "competitve_season_send_reward_mark:%v"     //serverid
	CompetitiveSeasonRankInfoByOpenDay = "competitve_season_rank_info_by_openDay:%v" //season 竞技场赛季的排名 过期时间14天

	Boss_Kill_ = "boss_kill_num_%d_*" //boss击杀 hash {fightType} {userId,fightNum}
	First_drop_item_all_get_num = "first_drop_item_all_get_num:*"
)

func Init(rc *beans.RedisConfig, sId int) error {
	return redisNewDb.Init(rc.Network, rc.Address, rc.Password, rc.DB)
}

func RedisDb() *redisdb.RedisDb {
	return redisNewDb
}
