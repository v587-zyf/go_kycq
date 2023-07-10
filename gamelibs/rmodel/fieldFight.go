package rmodel

import (
	"fmt"
	"time"
)

const (
	FieldFightChangeRivalCd         = "field_fight_change_rival_cd:%v:%v"        // openDay:userId  野战刷新对手cd
	FieldFightDefeatOwnerUsers      = "field_fight_defeat_owner_users:%v:%v"     // openDay:userId  挑战失败的玩家id
	FieldFightMatchUserFiveAmCombat = "field_fight_match_user_five_am_combat:%v" //openDay //每日玩家5点的战力 hmset
	FieldFightSaveBeforeRefRivals   = "field_fight_save_before_ref_rivals"       //openDay:userId //之前刷新的挑战玩家

)

type FieldFightModel struct {
}

var FieldFight = &FieldFightModel{}

func (this *FieldFightModel) GetDay() int {
	return time.Now().Day()
}

func (this *FieldFightModel) GetFieldFightChangeRivalCdKey(userId, openDay int) string {
	return fmt.Sprintf(FieldFightChangeRivalCd, openDay, userId)
}

func (this *FieldFightModel) GetFieldFightDefeatOwnerUsersKey(userId, openDay int) string {
	return fmt.Sprintf(FieldFightDefeatOwnerUsers, openDay, userId)
}

func (this *FieldFightModel) GetFieldFightMatchUserFiveAmCombatKey(openDay int) string {
	return fmt.Sprintf(FieldFightMatchUserFiveAmCombat, openDay)
}

func (this *FieldFightModel) GetFieldFightSaveBeforeRefRivalsKey() string {
	return fmt.Sprintf(FieldFightSaveBeforeRefRivals)
}

//野战刷新对手cd
func (this *FieldFightModel) SetFieldFightChangeRivalCd(userId, openDay, cd int) {
	key := this.GetFieldFightChangeRivalCdKey(userId, openDay)
	redisDb.Set(key, cd)
	redisDb.Expire(key, 2*24*time.Hour)
}

func (this *FieldFightModel) GetFieldFightChangeRivalCd(userId, openDay int) int {
	key := this.GetFieldFightChangeRivalCdKey(userId, openDay)
	num, _ := redisDb.Get(key).IntDef(0)
	return num
}

//自己挑战失败的玩家信息记录
func (this *FieldFightModel) SetFieldFightDefeatOwnerUsers(userId, openDay, rivalUserId int, nickName string) {
	key := this.GetFieldFightDefeatOwnerUsersKey(userId, openDay)
	redisDb.Hmset(key, rivalUserId, nickName)
	redisDb.Expire(key, 2*24*time.Hour)
}

func (this *FieldFightModel) GetFieldFightDefeatOwnerUsers(userId, openDay int) (map[int]string, error) {
	key := this.GetFieldFightDefeatOwnerUsersKey(userId, openDay)
	data, err := redisDb.HGetAllIntAndStringMap(key)
	return data, err
}

func (this *FieldFightModel) DelFieldFightDefeatOwnerUsers(userId, challengeId, openDay int) {
	key := this.GetFieldFightDefeatOwnerUsersKey(userId, openDay)
	state, _ := redisDb.Hexists(key, challengeId)
	if state == 1 {
		redisDb.HDel(key, challengeId)
	}
}

func (this *FieldFightModel) SetFieldFightMatchUserFiveAmCombat(openDay int, userInfo map[int]int) {
	key := this.GetFieldFightMatchUserFiveAmCombatKey(openDay)
	redisDb.HmsetIntMap(key, userInfo)
	redisDb.Expire(key, 2*24*time.Hour)
}

func (this *FieldFightModel) SetFieldFightMatchUserFiveAmCombat1(openDay int, userId, combat int) {
	key := this.GetFieldFightMatchUserFiveAmCombatKey(openDay)
	redisDb.Hmset(key, userId, combat)
	redisDb.Expire(key, 2*24*time.Hour)
}

func (this *FieldFightModel) GetFieldFightMatchUserFiveAmCombat(openDay, userId int) int {
	key := this.GetFieldFightMatchUserFiveAmCombatKey(openDay)
	num, _ := redisDb.HgetIntDef(key, userId, 0)
	return num
}

func (this *FieldFightModel) SetFieldFightSaveBeforeRefRivals(userId int, data string) {
	key := this.GetFieldFightSaveBeforeRefRivalsKey()
	redisDb.Hmset(key, userId, data)
	redisDb.Expire(key, 2*24*time.Hour)
}

func (this *FieldFightModel) GetFieldFightSaveBeforeRefRivals(userId int) (data string, err error) {
	key := this.GetFieldFightSaveBeforeRefRivalsKey()
	return redisDb.HgetStr(key, userId)
}

func (this *FieldFightModel) GetIsHaveFieldFightSaveBeforeRefRivals(userId int) (state int, err error) {
	key := this.GetFieldFightSaveBeforeRefRivalsKey()
	return redisDb.Hexists(key, userId)
}
