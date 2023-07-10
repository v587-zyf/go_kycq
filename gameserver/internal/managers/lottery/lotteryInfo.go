package lottery

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelGame"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"math"
	"math/rand"
	"time"
)

func (this *LotteryManager) getAllUserHaveBuyNum() map[int]int {
	defer this.locker.Unlock()
	this.locker.Lock()
	return this.userHaveBuyNum
}

func (this *LotteryManager) getUserHaveBuyNumByUserId(userId int) int {
	defer this.locker.Unlock()
	this.locker.Lock()
	return this.userHaveBuyNum[userId]
}

func (this *LotteryManager) addHaveBuyNum(userId, num int) {
	defer this.locker.Unlock()
	this.locker.Lock()
	this.haveBuyNums += num
	this.userHaveBuyNum[userId] += num
}

func (this *LotteryManager) getHaveBuyNum() int {
	defer this.locker.Unlock()
	this.locker.Lock()
	return this.haveBuyNums
}

func (this *LotteryManager) setLotteryInfo(info *modelGame.Lottery) {
	defer this.locker.Unlock()
	this.locker.Lock()
	this.lotteryInfo = append(this.lotteryInfo, info)
}

func (this *LotteryManager) getAllLotteryInfo() []*modelGame.Lottery {
	defer this.locker.Unlock()
	this.locker.Lock()
	return this.lotteryInfo
}

func (this *LotteryManager) setTakeNumber(num int) {
	defer this.locker.Unlock()
	this.locker.Lock()
	this.takeNumber[num] = true
}

func (this *LotteryManager) getTakeNumber(num int) bool {
	defer this.locker.Unlock()
	this.locker.Lock()
	return this.takeNumber[num]
}

func (this *LotteryManager) setGetAwardState(userId int, state bool) {
	defer this.locker.Unlock()
	this.locker.Lock()
	this.getAwardState[userId] = state
}

func (this *LotteryManager) getGetAwardState(userId int) bool {
	defer this.locker.Unlock()
	this.locker.Lock()
	return this.getAwardState[userId]
}

func (this *LotteryManager) getAllGetAwardState() map[int]bool {
	defer this.locker.Unlock()
	this.locker.Lock()
	return this.getAwardState
}

//随机获得抽奖号
func (this *LotteryManager) randomGetTakeNumber() int {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(10000) + 1
	if !this.getTakeNumber(randNum) {
		this.setTakeNumber(randNum)
		return randNum
	}
	randNum1 := 0
	for i := 1; i <= gamedb.GetConf().GrowCount; i++ {
		randNum1 = rand.Intn(10000) + 1
		if !this.getTakeNumber(randNum1) {
			break
		}
	}
	this.setTakeNumber(randNum1)
	return randNum1
}

func (this *LotteryManager) checkIsOpenByOpenDay() bool {

	limitOpenDay := gamedb.GetConf().YaoCaiTime
	continuousOpenDay := gamedb.GetConf().YaoCaiTime1
	openDay := this.GetSystem().GetServerOpenDaysByServerId(base.Conf.ServerId)
	if openDay < limitOpenDay {
		return false
	}
	if openDay > continuousOpenDay[0]+limitOpenDay {
		return false
	}

	return true
}

func (this *LotteryManager) checkIsOpen(user *objs.User) bool {

	limitOpenDay := gamedb.GetConf().YaoCaiTime
	continuousOpenDay := gamedb.GetConf().YaoCaiTime1
	openDay := this.GetSystem().GetServerOpenDaysByServerId(base.Conf.ServerId)
	if openDay < limitOpenDay {
		logger.Debug("LotteryManager checkIsOpen  openDay:%v limitOpenDay:%v", openDay, limitOpenDay)
		return false
	}
	if openDay > continuousOpenDay[0]+limitOpenDay {
		logger.Debug("LotteryManager checkIsOpen  openDay:%v limitOpenDay:%v  continuousOpenDay[0]:%v", openDay, limitOpenDay, continuousOpenDay[0])
		return false
	}

	buyTime := gamedb.GetConf().BuyTime
	openTime := buyTime[0].GetSecondsFromZero()
	closeTime := buyTime[1].GetSecondsFromZero()

	now := common.GetTimeSeconds(time.Now())

	if openTime > now {
		logger.Debug("LotteryManager buyTime:%v  openTime:%v  now:%v", buyTime, openTime, now)
		return false
	}

	if now > closeTime {
		logger.Debug("LotteryManager buyTime:%v  closeTime:%v  now:%v", buyTime, closeTime, now)
		return false
	}
	return true
}

func (this *LotteryManager) buildMyLotteryInfos(user *objs.User) ([]*pb.LotteryInfo, []*pb.LotteryInfo, []*pb.LotteryInfo) {

	myLotteryInfo := make([]*pb.LotteryInfo, 0)
	allLotteryInfo := make([]*pb.LotteryInfo, 0)
	winLotteryInfo := make([]*pb.LotteryInfo, 0)
	lotteryInfos := this.getAllLotteryInfo()
	winUserId := this.getLotteryWinUser()

	for _, v := range lotteryInfos {
		if v.UserId == user.Id {
			myLotteryInfo = append(myLotteryInfo, &pb.LotteryInfo{UserId: int32(user.Id), UserName: user.NickName, AwardNumber: int32(v.AwardNumber), ShareNum: int32(v.Share)})
		}
		userInfo := this.GetUserManager().GetUserBasicInfo(v.UserId)
		if userInfo != nil {
			if v.UserId == winUserId {
				winLotteryInfo = append(winLotteryInfo, &pb.LotteryInfo{UserId: int32(userInfo.Id), UserName: userInfo.NickName, AwardNumber: int32(v.AwardNumber), ShareNum: int32(v.Share), Combat: int64(userInfo.Combat)})
			}
			allLotteryInfo = append(allLotteryInfo, &pb.LotteryInfo{UserId: int32(v.UserId), UserName: userInfo.NickName, AwardNumber: int32(v.AwardNumber), ShareNum: int32(v.Share)})
		}
	}
	return myLotteryInfo, allLotteryInfo, winLotteryInfo
}

func (this *LotteryManager) getWinAwardsInfo() []*pb.ItemUnit {
	luckyReward := gamedb.GetConf().LuckyReward
	items := make([]*pb.ItemUnit, 0)
	if len(luckyReward) > 0 {
		for _, v := range luckyReward {
			pbItem := &pb.ItemUnit{ItemId: int32(v.ItemId), Count: int64(v.Count)}
			items = append(items, pbItem)
		}
	}
	return items
}

func (this *LotteryManager) getLoseAwardsInfo(userId int) []*pb.ItemUnit {
	onePrice := gamedb.GetConf().OnePrice         //一份额价格
	notLuckyCoin := gamedb.GetConf().NotLuckyGold //未中奖投入1元宝的返还比例
	items := make([]*pb.ItemUnit, 0)
	num := this.getUserHaveBuyNumByUserId(userId)
	if num > 0 {
		loseReturnNun := int64(math.Floor(float64(num*onePrice[0].Count) * notLuckyCoin))
		items = append(items, &pb.ItemUnit{ItemId: int32(onePrice[0].ItemId), Count: loseReturnNun})
	}
	loseReward := gamedb.GetConf().NotLuckReward //未奖奖励
	if len(loseReward) > 0 {
		for _, v := range loseReward {
			pbItem := &pb.ItemUnit{ItemId: int32(v.ItemId), Count: int64(v.Count)}
			items = append(items, pbItem)
		}
	}
	return items
}
