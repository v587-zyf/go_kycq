package rmodel

import (
	"fmt"
	"time"
)

const (
	LOTTERY_WIN_USER = "lottery_win_user:%v"
)

type LotteryModel struct {
}

var Lottery = &LotteryModel{}

func (this *LotteryModel) getNowDay() int {
	return time.Now().Day()
}

func (this *LotteryModel) getLotteryWinUserKey() string {
	return fmt.Sprintf(LOTTERY_WIN_USER, this.getNowDay())
}

func (this *LotteryModel) SetLotteryWinUser(userId int) {
	redisDb.SetWithExpire(this.getLotteryWinUserKey(), userId, AutoExpireTime)
}

func (this *LotteryModel) GetLotteryWinUser(day int) int {
	num, _ := redisDb.Get(fmt.Sprintf(LOTTERY_WIN_USER, day)).IntDef(-1)
	return num
}
