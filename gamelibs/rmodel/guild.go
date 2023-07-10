package rmodel

import (
	"fmt"
	"time"
)

const (
	GUILDAPLLUSERIDSLIST  = "guild_apply_user_ids_list:%v" //guildId
	GUILDIMPEACHPRESIDENT = "guild_impeach_president"      //userId
)

type GuildModel struct {
}

var Guild = &GuildModel{}

func (this *GuildModel) GetDay() int {
	return time.Now().Day()
}

func (this *GuildModel) GetSetApplyUserIdsKey(guildId int) string {
	return fmt.Sprintf(GUILDAPLLUSERIDSLIST, guildId)
}

func (this *GuildModel) GetGuildImpeachPresidentKey() string {
	return fmt.Sprintf(GUILDIMPEACHPRESIDENT)
}

func (this *GuildModel) SetGuildApplyUserId(guildId, applyUserId int) {
	key := this.GetSetApplyUserIdsKey(guildId)
	redisDb.Hmset(key, applyUserId, 1)
	redisDb.Expire(key, 24*time.Hour)
}

func (this *GuildModel) GetGuildApplyUserIds(guildId int) (map[int]int, error) {
	key := this.GetSetApplyUserIdsKey(guildId)
	return redisDb.HgetallIntMap(key)
}

func (this *GuildModel) IsExistUserId(guildId, applyUserId int) bool {
	key := this.GetSetApplyUserIdsKey(guildId)
	state, err := redisDb.Hexists(key, applyUserId)
	if err != nil {
		return false
	}
	return state == 1
}

func (this *GuildModel) DelApplyUserId(guildId, applyUserId int) error {
	key := this.GetSetApplyUserIdsKey(guildId)
	_, err := redisDb.HDel(key, applyUserId)
	if err != nil {
		return err
	}
	return nil
}

//弹劾会长
func (this *GuildModel) SetGuildImpeachPresident(userId, times int) {
	key := this.GetGuildImpeachPresidentKey()
	redisDb.Hmset(key, userId, times)
	redisDb.Expire(key, time.Second*60*30)
}

func (this *GuildModel) IsExistGuildImpeachPresident(userId int) bool {
	key := this.GetGuildImpeachPresidentKey()
	state, err := redisDb.Hexists(key, userId)
	if err != nil {
		return false
	}
	return state == 1
}

func (this *GuildModel) DelGuildImpeachPresident(userId int) error {
	key := this.GetGuildImpeachPresidentKey()
	_, err := redisDb.HDel(key, userId)
	if err != nil {
		return err
	}
	return nil
}
