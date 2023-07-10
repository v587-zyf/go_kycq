package lottery

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constMail"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"math"
	"math/rand"
	"time"
)

func NewLotteryManager(module managersI.IModule) *LotteryManager {
	return &LotteryManager{
		IModule:        module,
		locker:         util.NewLocker(0),
		lotteryInfo:    make([]*modelGame.Lottery, 0),
		takeNumber:     make(map[int]bool),
		userHaveBuyNum: make(map[int]int),
		getAwardState:  make(map[int]bool),
	}
}

type LotteryManager struct {
	util.DefaultModule
	managersI.IModule
	locker         *util.Locker
	lotteryInfo    []*modelGame.Lottery
	takeNumber     map[int]bool //已使用的奖号
	haveBuyNums    int          //所有玩家已购买的份数
	userHaveBuyNum map[int]int  //玩家已经购买的份数
	getAwardState  map[int]bool //领奖状态
}

func (this *LotteryManager) Init() error {

	datas, err := modelGame.GetLotteryModel().GetAllLotteryInfos()
	if err != nil {
		logger.Error("GetAllLotteryInfos err:%v", err)
		return err
	}

	for _, data := range datas {
		this.lotteryInfo = append(this.lotteryInfo, data)
		this.userHaveBuyNum[data.UserId] += data.Share
		this.haveBuyNums += data.Share
		this.takeNumber[data.AwardNumber] = true
	}

	for userId := range this.userHaveBuyNum {
		userInfo := this.GetUserManager().GetOfflineUserInfo(userId)
		if userInfo == nil {
			continue
		}
		this.getAwardState[userId] = userInfo.LotteryInfo.IsGetAward == 1

	}
	logger.Info("userHaveBuyNum:%v", this.userHaveBuyNum)
	return nil
}

func (this *LotteryManager) Load(user *objs.User, ack *pb.LotteryInfoAck) {

	ack.PopUpState = int32(user.LotteryInfo.PopUpState)
	ack.GoodLuckState = int32(user.LotteryInfo.GoodLuckState)
	myLotteryInfo, allLotteryInfo, winLotteryInfo := this.buildMyLotteryInfos(user)
	ack.MyLotteryInfos = myLotteryInfo
	ack.AllLotteryInfos = allLotteryInfo
	ack.WinLotteryInfos = winLotteryInfo
	winUserId := this.getLotteryWinUser()
	if winUserId > 0 {
		ack.WinUserInfo = this.GetUserManager().BuilderBrieUserInfo(winUserId)
	}
	return
}

func (this *LotteryManager) getLotteryWinUser() int {
	winUserId := 0
	openAwardTime := gamedb.GetConf().RewardTime
	if common.GetTimeSeconds(time.Now()) > openAwardTime[0].GetSecondsFromZero() {
		winUserId = rmodel.Lottery.GetLotteryWinUser(time.Now().Day())
	} else {
		winUserId = rmodel.Lottery.GetLotteryWinUser(time.Now().Day() - 1)
	}
	return winUserId
}

//接好运
func (this *LotteryManager) GetGoodLucky(user *objs.User, ack *pb.GetGoodLuckAck, op *ophelper.OpBagHelperDefault) error {
	if !this.checkIsOpenByOpenDay() {
		return gamedb.ERRCHALLENGE2
	}

	if user.LotteryInfo.GoodLuckState >= 1 {
		return gamedb.ERRHAVEGETREWARD
	}

	luckyItem := gamedb.GetConf().LuckyItem
	this.GetBag().AddItems(user, luckyItem, op)
	user.LotteryInfo.GoodLuckState = 1
	ack.State = 1
	user.Dirty = true
	return nil
}

//购买份额
func (this *LotteryManager) LotteryBuyNums(user *objs.User, nums int, ack *pb.LotteryBuyNumsAck, op *ophelper.OpBagHelperDefault) error {

	if !this.checkIsOpen(user) {
		return gamedb.ERRCHALLENGE2
	}

	if nums <= 0 {
		return gamedb.ERRPARAM
	}

	allHaveBuyNum := this.getHaveBuyNum()
	allCanBuyNun := gamedb.GetConf().AllBuy
	if nums > allCanBuyNun-allHaveBuyNum {
		logger.Error("allHaveBuyNum:%v  allCanBuyNun:%v", allHaveBuyNum, allCanBuyNun)
		return gamedb.ERRBUYUPPERLIMIT
	}

	ownHaveBuyNum := this.getUserHaveBuyNumByUserId(user.Id)
	if nums > gamedb.GetConf().BuyNum-ownHaveBuyNum {
		return gamedb.ERRBUYUPPERLIMIT
	}

	items := gamedb.GetConf().OnePrice
	for _, info := range items {
		if has, _ := this.GetBag().HasEnough(user, info.ItemId, info.Count*nums); !has {
			return gamedb.ERRNOTENOUGHGOODS
		}
	}

	for _, info := range items {
		_ = this.GetBag().Remove(user, op, info.ItemId, info.Count*nums)
	}

	data := make([]*pb.LotteryInfo, 0)
	for i := 1; i <= nums; i++ {
		//随机生成奖号
		awardNum := this.randomGetTakeNumber()
		msg := &modelGame.Lottery{
			UserId:      user.Id,
			AwardNumber: awardNum,
			Share:       nums,
			Day:         time.Now().Day(),
		}
		this.setLotteryInfo(msg)
		err := modelGame.GetLotteryModel().DbMap().Insert(msg)
		if err != nil {
			logger.Error("GetLotteryModel().DbMap().Insert  msg:%v  err:%v", msg, err)
			return err
		}
		data = append(data, &pb.LotteryInfo{UserId: int32(user.Id), UserName: user.NickName, AwardNumber: int32(awardNum), ShareNum: 1})
	}
	this.addHaveBuyNum(user.Id, nums)
	this.setGetAwardState(user.Id, user.LotteryInfo.IsGetAward == 1)
	ack.LotteryInfos = data
	ntf := &pb.BrocastBuyNumsNtf{}
	ntf.LotteryInfos = data
	this.BroadcastAll(ntf)
	user.Dirty = true
	return nil
}

//获奖信息
func (this *LotteryManager) GetAwardInfo(user *objs.User, ack *pb.LotteryInfo1Ack) {
	userId := this.getLotteryWinUser()
	if userId == user.Id {
		ack.IsWin = 1
		ack.Items = this.getWinAwardsInfo()
		winLotteryInfo := make([]*pb.LotteryInfo, 0)
		lotteryInfos := this.getAllLotteryInfo()
		for _, v := range lotteryInfos {
			if v.UserId == userId {
				userInfo := this.GetUserManager().GetUserBasicInfo(userId)
				if userInfo != nil {
					winLotteryInfo = append(winLotteryInfo, &pb.LotteryInfo{UserId: int32(v.UserId), UserName: userInfo.NickName, AwardNumber: int32(v.AwardNumber), ShareNum: int32(v.Share)})
				}
			}
		}
		ack.WinLotteryInfos = winLotteryInfo
		return
	}
	ack.Items = this.getLoseAwardsInfo(user.Id)
	return
}

//领取投注奖励
func (this *LotteryManager) GetLotteryAward(user *objs.User, ack *pb.LotteryGetEndAwardAck, op *ophelper.OpBagHelperDefault) error {
	if user.LotteryInfo.IsGetAward >= 1 {
		return gamedb.ERRAWARDGET
	}

	ownHaveBuyNum := this.getUserHaveBuyNumByUserId(user.Id)
	if ownHaveBuyNum <= 0 {
		return gamedb.ERRLOTTERY
	}
	userId := this.getLotteryWinUser()
	awards := make([]*pb.ItemUnit, 0)
	if userId == user.Id {
		awards = this.getWinAwardsInfo()
	} else {
		awards = this.getLoseAwardsInfo(user.Id)
	}

	if len(awards) <= 0 {
		return gamedb.ERRGETCONDITIONERR2
	}
	items := make(gamedb.ItemInfos, 0)
	for _, v := range awards {
		items = append(items, &gamedb.ItemInfo{ItemId: int(v.ItemId), Count: int(v.Count)})
	}
	this.GetBag().AddItems(user, items, op)
	user.LotteryInfo.IsGetAward = 1
	ack.GetState = 1
	this.setGetAwardState(user.Id, true)
	ack.Items = awards
	return nil
}

func (this *LotteryManager) OpenAward() {
	//每天限购份额	100
	//保底修正参数	0.3
	//成长修正参数	0.01
	//概率修正公式	单个玩家实际中奖公式=当天单个玩家购买份额/{当天总玩家购买份额/【MIN(保底修正参数+当天总玩家购买份额/每天限购份额/0.05*成长修正参数+当天总玩家购买份额/每天限购份额,1)*每天限购份额】*100}

	infos := make(map[int]int, 0)
	allCanBuyNun := float64(gamedb.GetConf().AllBuy)        //每天限购份额
	param := float64(gamedb.GetConf().LimitCount) / 10000.0 //保底修正参数（万分比）
	param1 := float64(gamedb.GetConf().GrowCount) / 10000.0 //成长修正参数（万分比）
	allBuyNum := float64(this.getHaveBuyNum())
	allUserHaveBuyNum := this.getAllUserHaveBuyNum()

	if len(allUserHaveBuyNum) <= 0 {
		this.BroadcastAll(&pb.LotteryEnd{})
		return
	}
	for userId, num := range allUserHaveBuyNum {
		afterNum := float64(num) / (allBuyNum / (math.Min(param+allBuyNum/allCanBuyNun/0.05*param1+allBuyNum/allCanBuyNun, 1) * allCanBuyNun) * 100)
		infos[userId] = int(afterNum * 100)
	}

	rand.Seed(time.Now().Unix())
	winUserId := common.RandWeightByMap(infos)
	logger.Info("摇彩  所有人购买的总份数:%v  中奖玩家:%v  allUserHaveBuyNum:%v   玩家权重:%v", allBuyNum, winUserId, allUserHaveBuyNum, infos)
	rmodel.Lottery.SetLotteryWinUser(winUserId)
	lotteryInfos := this.getAllLotteryInfo()
	winLotteryInfo := make([]*pb.LotteryInfo, 0)
	for _, v := range lotteryInfos {
		if v.UserId == winUserId {
			userInfo := this.GetUserManager().GetUserBasicInfo(winUserId)
			if userInfo != nil {
				winLotteryInfo = append(winLotteryInfo, &pb.LotteryInfo{UserId: int32(v.UserId), UserName: userInfo.NickName, AwardNumber: int32(v.AwardNumber), ShareNum: int32(v.Share)})
			}
		}
	}

	this.BroadcastAll(&pb.LotteryEnd{
		WinLotteryInfos: winLotteryInfo,
		WinUserInfo:     this.GetUserManager().BuilderBrieUserInfo(winUserId),
	})
	winReward := gamedb.GetConf().LuckyReward
	if len(winReward) > 0 {
		winUserInfo := this.GetUserManager().GetAllUserInfoIncludeOfflineUser(winUserId)
		this.GetAnnouncement().SendSystemChat(winUserInfo, pb.SCROLINGTYPE_LOTTERY, winReward[0].ItemId, -1)
	}

	return
}

func (this *LotteryManager) SendReward() {
	winUserId := this.getLotteryWinUser()
	infos := this.getAllGetAwardState()
	logger.Debug("infos:%v  winUserId:%v", infos, winUserId)
	for userId, state := range infos {
		if state {
			continue
		}
		if userId == winUserId {
			luckyReward := gamedb.GetConf().LuckyReward //中奖奖励
			luckyReward = append(luckyReward, gamedb.GetConf().LuckyGold...)
			this.GetMail().SendSystemMailWithItemInfos(winUserId, constMail.LOTTERY_WIN_AWARD, []string{}, luckyReward)
			continue
		}

		awards := this.getLoseAwardsInfo(userId)
		if len(awards) <= 0 {
			continue
		}
		items := make(gamedb.ItemInfos, 0)
		for _, v := range awards {
			items = append(items, &gamedb.ItemInfo{ItemId: int(v.ItemId), Count: int(v.Count)})
		}
		this.GetMail().SendSystemMailWithItemInfos(userId, constMail.LOTTERY_LOSE_AWARD, []string{}, items)
	}
	//清空内存
	this.Reset()
}

func (this *LotteryManager) SetPopState(user *objs.User, ack *pb.SetLotteryPopUpStateAck) {

	user.LotteryInfo.PopUpState = 1
	ack.State = 1
}

func (this *LotteryManager) OnlineCheck(user *objs.User) {

	if user.LotteryInfo.ResetDay != time.Now().Day() {
		user.LotteryInfo.ResetDay = time.Now().Day()
		user.LotteryInfo.PopUpState = 0
		user.LotteryInfo.GoodLuckState = 0
		user.LotteryInfo.IsGetAward = 0
	}
}

func (this *LotteryManager) Reset() {
	this.takeNumber = make(map[int]bool, 0)
	this.lotteryInfo = make([]*modelGame.Lottery, 0)
	this.haveBuyNums = 0
	this.userHaveBuyNum = make(map[int]int, 0)
	_ = modelGame.GetLotteryModel().DeleteAllItem()
	return
}

func (this *LotteryManager) UserReset(user *objs.User) {
	user.LotteryInfo.IsGetAward = 0
	user.LotteryInfo.PopUpState = 0
	user.LotteryInfo.GoodLuckState = 0
	user.LotteryInfo.ResetDay = time.Now().Day()
	user.Dirty = true
	ack := &pb.LotteryInfoAck{}
	this.Load(user, ack)
	this.GetUserManager().SendMessage(user, ack, true)
	return
}
