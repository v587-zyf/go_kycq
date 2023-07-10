package rmodel

import (
	"cqserver/golibs/logger"
	"fmt"
	"time"
)

const (
	CROSSSHABAKEFIRSEGUILDINFO = "cross_shabake_first_guild_info:%v"
	SHABAKEFIRSEGUILDINFO      = "shabake_first_guild_info:%v"
	SHABAKEISEND               = "shabake_is_end:%v:%v"       //day:serverId
	CROSSSHABAKEISEND          = "cross_shabake_is_end:%v:%v" //day:serverId
)

type ShabakeModel struct {
}

var Shabake = &ShabakeModel{}

func (this *ShabakeModel) GetCrossShabakeFirstGuildInfoKey() string {
	return fmt.Sprintf(CROSSSHABAKEFIRSEGUILDINFO)
}

func (this *ShabakeModel) GetShabakeFirstGuildInfoKey() string {
	return fmt.Sprintf(SHABAKEFIRSEGUILDINFO)
}

func (this *ShabakeModel) SetCrossShabakeFirstGuildInfo(data string) {
	key := this.GetCrossShabakeFirstGuildInfoKey()
	redisDb.Set(key, data)
}

func (this ShabakeModel) GetCrossShabakeFirstGuildInfo() string {
	key := this.GetCrossShabakeFirstGuildInfoKey()
	v, err := redisDb.Get(key).String()
	if err != nil {
		logger.Error("GetCrossShabakeFirstGuildInfo|err:%v", err)
		return ""
	}
	return v
}

func (this *ShabakeModel) SetShabakeFirstGuildInfo(data string) {
	key := this.GetShabakeFirstGuildInfoKey()
	redisDb.Set(key, data)
}

func (this ShabakeModel) GetShabakeFirstGuildInfo() string {
	key := this.GetShabakeFirstGuildInfoKey()
	v, err := redisDb.Get(key).String()
	if err != nil {
		logger.Error("GetShabakeFirstGuildInfo|err:%v", err)
		return ""
	}
	return v
}

func (this *ShabakeModel) SetShaBakeIsEnd(serverId, state int) {
	redisDb.Set(fmt.Sprintf(SHABAKEISEND, time.Now().Day(), serverId), state)
	redisDb.Expire(SHABAKEISEND, 24*time.Hour)
}

func (this *ShabakeModel) GetShaBakeIsEnd(serverId int) int {
	num, _ := redisDb.Get(fmt.Sprintf(SHABAKEISEND, time.Now().Day(), serverId)).IntDef(0)
	return num
}

func (this *ShabakeModel) SetCrossShaBakeIsEnd(serverId, state int) {
	redisDb.Set(fmt.Sprintf(CROSSSHABAKEISEND, time.Now().Day(), serverId), state)
	redisDb.Expire(CROSSSHABAKEISEND, 24*time.Hour)
}

func (this *ShabakeModel) GetCrossShaBakeIsEnd(serverId int) int {
	num, _ := redisDb.Get(fmt.Sprintf(CROSSSHABAKEISEND, time.Now().Day(), serverId)).IntDef(0)
	return num
}
