package rmodel

import (
	"cqserver/golibs/logger"
	"fmt"
	"time"
)

const (
	WORLDLEADERRANKINFO     = "world_leader_rank_info:%v:%v"
	WORLDLEADERRHP          = "world_leader_hp:%v:%v"
	WORLDLEADERRGUILDNUMBER = "world_leader_guild_number:%v:%v:%v"
	WORLDLEADERRGUILD       = "world_leader_guild:%v:%v"
)

type WorldLeaderModel struct {
}

var WorldLeader = &WorldLeaderModel{}

func (this *WorldLeaderModel) GetWorldLeaderRankInfoKey(openDay, stageId int) string {
	return fmt.Sprintf(WORLDLEADERRANKINFO, openDay, stageId)
}

func (this *WorldLeaderModel) GetWorldLeaderHpKey(openDay, stageId int) string {
	return fmt.Sprintf(WORLDLEADERRHP, openDay, stageId)
}

func (this *WorldLeaderModel) GetWorldLeaderGuildNumberKey(openDay, stageId, guildId int) string {
	return fmt.Sprintf(WORLDLEADERRGUILDNUMBER, openDay, stageId, guildId)
}

func (this *WorldLeaderModel) GetWorldLeaderGuildKey(openDay, stageId int) string {
	return fmt.Sprintf(WORLDLEADERRGUILD, openDay, stageId)
}

func (this *WorldLeaderModel) SetWorldLeaderRankInfo(stageId int, data string) {
	key := this.GetWorldLeaderRankInfoKey(time.Now().Day(), stageId)
	redisDb.Set(key, data)
	redisDb.Expire(key, 24*time.Hour)
}

func (this *WorldLeaderModel) GetWorldLeaderRankInfo(stageId int) string {
	key := this.GetWorldLeaderRankInfoKey(time.Now().Day(), stageId)
	data, _ := redisDb.Get(key).String()
	return data
}

func (this *WorldLeaderModel) SetWorldLeaderHpInfo(stageId, hp int) {
	key := this.GetWorldLeaderHpKey(time.Now().Day(), stageId)
	redisDb.Set(key, hp)
	redisDb.Expire(key, 24*time.Hour)
}

func (this *WorldLeaderModel) GetWorldLeaderHpInfo(stageId int) int {
	key := this.GetWorldLeaderHpKey(time.Now().Day(), stageId)
	data, _ := redisDb.Get(key).IntDef(0)
	return data
}

func (this *WorldLeaderModel) SetWorldLeaderEnterGuild(stageId, guildId int) {
	key := this.GetWorldLeaderGuildKey(time.Now().Day(), stageId)
	redisDb.Hmset(key, guildId, 1)
	redisDb.Expire(key, 24*time.Hour)
}

func (this *WorldLeaderModel) GetWorldLeaderEnterGuilds(stageId int) map[int]int {
	key := this.GetWorldLeaderGuildKey(time.Now().Day(), stageId)
	data, _ := redisDb.HgetallIntMap(key)
	return data
}

func (this *WorldLeaderModel) SetWorldLeaderEnterGuildNumber(stageId, guildId, userId int) {
	key := this.GetWorldLeaderGuildNumberKey(time.Now().Day(), stageId, guildId)
	redisDb.LPush(key, userId)
	redisDb.Expire(key, 24*time.Hour)
}

func (this *WorldLeaderModel) GetWorldLeaderEnterGuildNumber(stageId, guildId int) []int {
	key := this.GetWorldLeaderGuildNumberKey(time.Now().Day(), stageId, guildId)

	v, err := redisDb.LRange(key).ValuesIntSlice()
	if err != nil {
		logger.Error("RangeAuctionTogether|err:%v", err)
		return []int{}
	}

	return v
}



