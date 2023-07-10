package rmodel

import (
	"fmt"
	"time"
)

const (
	challengexiazhu = "challenge_xia_zhu:%v:%v" //userId:day:Hour
	expireTime      = time.Hour * 24 * 3
)

type ChallengeModel struct {
}

var Challenge = &ChallengeModel{}

func (this *ChallengeModel) genChallengeXiaZhu(roundIndex int, season string) string {
	return fmt.Sprintf(challengexiazhu, roundIndex, season)
}

func (this *ChallengeModel) HSetBottomUser(roundIndex, userId, bottomUser int, season string) {
	key := this.genChallengeXiaZhu(roundIndex, season)
	redisDb.Hmset(key, userId, bottomUser)
	redisDb.Expire(key, expireTime)
}

func (this ChallengeModel) CheckBottomUserIsExist(roundIndex, userId int, season string) int {
	key := this.genChallengeXiaZhu(roundIndex, season)
	state, _ := redisDb.Hexists(key, userId)
	return state
}

func (this ChallengeModel) HGetAllBottomUsers(roundIndex int, season string) map[int]int {
	key := this.genChallengeXiaZhu(roundIndex, season)
	state, _ := redisDb.HgetallIntMap(key)
	return state
}
