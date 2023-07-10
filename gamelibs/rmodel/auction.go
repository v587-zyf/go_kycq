package rmodel

import (
	"cqserver/golibs/logger"
	"fmt"
	"time"
)

const (
	auctionTogether   = "auction_together:%d:%v:%v" //userId:day:Hour
	auctionExpireTime = time.Second * 60 * 60 * 12
)

type AuctionModel struct {
}

var Auction = &AuctionModel{}

func (this *AuctionModel) genAuctionTogether(userId int) string {
	return fmt.Sprintf(auctionTogether, userId, this.getNowDay(), this.getNowHour())
}

func (this *AuctionModel) getNowHour() int {
	return time.Now().Hour()
}

func (this *AuctionModel) getNowDay() int {
	return time.Now().Day()
}

func (this *AuctionModel) AddAuctionTogether(userId, itemId int) {
	key := this.genAuctionTogether(userId)
	redisDb.LPush(key, itemId)
	redisDb.Expire(key, auctionExpireTime)
}

func (this AuctionModel) RangeAuctionTogether(userId int) []int {
	key := this.genAuctionTogether(userId)
	v, err := redisDb.LRange(key).ValuesIntSlice()
	if err != nil {
		logger.Error("RangeAuctionTogether|err:%v", err)
		return []int{}
	}
	return v
}

func (this *AuctionModel) RemAuctionTogether(userId, itemId int) {
	key := this.genAuctionTogether(userId)
	err := redisDb.LRem(key, itemId).Err
	if err != nil {
		logger.Error("RemAuctionTogether|error:%v", err)
	}
}
