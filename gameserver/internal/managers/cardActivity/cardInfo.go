package cardActivity

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constAuction"
	"cqserver/gamelibs/publicCon/constCard"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"fmt"
	"math/rand"
	"runtime/debug"
	"sync"
	"time"
)

func NewCardActivityManager(m managersI.IModule) *CardActivity {
	return &CardActivity{
		IModule:        m,
		opChan:         make(chan opMsg, constAuction.DbChanSize),
		myCardInfo:     make(map[int][]*pb.CardInfoUnit),
		serverCardInfo: make([]*pb.CardInfoUnit, 0),
	}
}

type CardActivity struct {
	util.DefaultModule
	managersI.IModule
	opChan         chan opMsg
	myCardInfo     map[int][]*pb.CardInfoUnit
	serverCardInfo []*pb.CardInfoUnit
	mu             sync.RWMutex
}

type opMsg struct {
	opType   int
	cardItem *modelGame.Card
}

func (this *CardActivity) Init() error {
	go this.dbDataOp()
	this.initCardInfo()
	return nil
}

func (this *CardActivity) Load(user *objs.User, ack *pb.CardActivityInfosAck) error {

	season, _, _ := this.getDrawSeason(user)
	logger.Debug("CardActivity season:%v", season)
	if season == -1 {
		return nil
	}
	user.CardInfo.Season = season
	ack.Integral = int32(user.CardInfo.Integral)
	ack.TotalDrawCardTimes = int32(user.CardInfo.DrawTimes)
	ack.HaveGetIndex, ack.MyDrawInfo, ack.ServerDrawInfo = this.getMyCardInfoAndSysInfo(user)
	ack.NowSeason = int32(season)
	ack.MergeMark = int32(user.CardInfo.MergeMark)
	user.Dirty = true
	return nil
}

func (this *CardActivity) Draw(user *objs.User, times int, ack *pb.CardActivityApplyGetAck, op *ophelper.OpBagHelperDefault) error {
	if times != constCard.OneTime {
		if times != constCard.TenTime {
			return gamedb.ERRPARAM
		}
	}

	if user.CardInfo.DrawTimes >= gamedb.GetConf().DrawMax {
		return gamedb.ERRCARDERR
	}

	var err error
	if times == constCard.OneTime {
		//单抽
		err = this.drawRemove(user, times, op)
	} else if times == constCard.TenTime {
		//十连抽
		if user.CardInfo.DrawTimes+times > gamedb.GetConf().DrawMax {
			times = gamedb.GetConf().DrawMax - user.CardInfo.DrawTimes
		}
		err = this.drawRemove(user, times, op)
	} else {
		err = gamedb.ERRPARAM
	}
	if err != nil {
		return err
	}
	drawTime := int(time.Now().Unix())
	for i := constCard.OneTime; i <= times; i++ {

		rand.Seed(time.Now().UnixNano() + int64(i))

		cfg, err := this.randLowOrHighPool(user)
		if err != nil {
			return err
		}
		logger.Debug("神机宝库 抽奖 time:%v", i)
		itemId, count := common.RandWeightBySlice2(cfg.Probability)
		if itemId <= 0 || count <= 0 {
			logger.Error("抽卡随机道具错误 RandWeightBySlice2 Probability:%v  itemId:%v  count:%v", cfg.Probability, itemId, count)
			continue
		}
		ack.Cards = append(ack.Cards, int32(itemId), int32(count))
		items := gamedb.ItemInfos{}
		items = append(items, &gamedb.ItemInfo{ItemId: itemId, Count: count})
		this.GetBag().AddItems(user, items, op)
		this.addItemIntoCard(user, itemId, count, times, drawTime)
		this.GetAnnouncement().SendSystemChat(user, pb.SCROLINGTYPE_CHOUKA, itemId, -1)
	}
	user.CardInfo.DrawTimes += times
	user.CardInfo.Integral += gamedb.GetConf().DrawMark * times
	user.Dirty = true
	ack.Type = int32(times)
	ack.CardTime = int32(user.CardInfo.DrawTimes)
	_, ack.MyDrawInfo, ack.ServerDrawInfo = this.getMyCardInfoAndSysInfo(user)

	//if times > 1 {
	//	allGoods := make([]*pb.ItemUnit, 0)
	//	allGoodsMark := make(map[int32]int64)
	//	allItems := op.ToChangeItems()
	//	for _, data := range allItems.Items {
	//		allGoodsMark[data.ItemId] += data.Count
	//	}
	//	for itemId, count := range allGoodsMark {
	//		allGoods = append(allGoods, &pb.ItemUnit{ItemId: itemId, Count: count})
	//	}
	//	data := op.ToChangeItems()
	//	data.Items = allGoods
	//	ack.Goods = data
	//} else {
	//	ack.Goods = op.ToChangeItems()
	//}
	ack.Goods = op.ToChangeItems()
	ack.Integral = int32(user.CardInfo.Integral)
	_ = this.GetUserManager().SendItemChangeNtf(user, op)
	this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_CARD_ACTIVITY, []int{times})
	return nil
}

//获取积分奖励
func (this *CardActivity) GetReward(user *objs.User, id, times int, ack *pb.GetIntegralAwardAck, op *ophelper.OpBagHelperDefault) error {

	if id <= 0 {
		return gamedb.ERRPARAM
	}

	cfg := gamedb.GetDrawShopDrawShopCfgByShop(user.CardInfo.Season, id)
	if cfg == nil {
		return gamedb.ERRPARAM
	}

	if user.CardInfo.Integral < cfg.Num*times {
		return gamedb.ERRCARDERR1
	}
	user.CardInfo.Integral -= cfg.Num * times
	this.GetBag().AddItem(user, op, cfg.Reward.ItemId, cfg.Reward.Count*times)
	//user.CardInfo.GetAwardIds = append(user.CardInfo.GetAwardIds, id)
	user.Dirty = true
	ack.HaveGetIndex, _, _ = this.getMyCardInfoAndSysInfo(user)
	ack.Integral = int32(user.CardInfo.Integral)
	ack.Goods = op.ToChangeItems()
	return nil
}

func (this *CardActivity) drawRemove(user *objs.User, times int, op *ophelper.OpBagHelperDefault) error {
	var err error
	if times < constCard.TenTime {
		//单抽
		itemInfo := gamedb.GetConf().Draw[0]
		if has, _ := this.GetBag().HasEnough(user, itemInfo.ItemId, itemInfo.Count*times); !has {
			itemInfo = gamedb.GetConf().Draw[2]
			if has, _ = this.GetBag().HasEnough(user, itemInfo.ItemId, itemInfo.Count*times); !has {
				return gamedb.ERRNOTENOUGHGOODS
			}
			err = this.GetBag().Remove(user, op, itemInfo.ItemId, itemInfo.Count*times)
			return err
		}
		err = this.GetBag().Remove(user, op, itemInfo.ItemId, itemInfo.Count*times)
	} else {
		itemInfo := gamedb.GetConf().Draw[1]
		if has, _ := this.GetBag().HasEnough(user, itemInfo.ItemId, itemInfo.Count); !has {
			itemInfo = gamedb.GetConf().Draw[3]
			if has, _ = this.GetBag().HasEnough(user, itemInfo.ItemId, itemInfo.Count); !has {
				return gamedb.ERRNOTENOUGHGOODS
			}
			err = this.GetBag().Remove(user, op, itemInfo.ItemId, itemInfo.Count)
			return err
		}
		err = this.GetBag().Remove(user, op, itemInfo.ItemId, itemInfo.Count)
	}

	return err
}

func (this *CardActivity) getDrawSeason(user *objs.User) (int, int, int) {
	serverId := base.Conf.ServerId
	if user != nil {
		serverId = user.ServerId
	}

	openDay := this.GetSystem().GetMergerServerOpenDaysByServerId(serverId)
	if user != nil && user.CardInfo.MergeMark == 1 {
		openDay = this.GetSystem().GetServerOpenDaysByServerIdByExcursionTime(serverId, 0)
	}
	_, season, restOpenDay := gamedb.GetDrawNowSeason(openDay)
	return season, openDay, restOpenDay
}

func (this *CardActivity) randLowOrHighPool(user *objs.User) (*gamedb.DrawDrawCfg, error) {

	cfg1 := gamedb.GetDrawCfgBySeasonAndType(user.CardInfo.Season, constCard.LowPool)
	cfg2 := gamedb.GetDrawCfgBySeasonAndType(user.CardInfo.Season, constCard.HighPool)
	if cfg1 == nil || cfg2 == nil {
		return nil, gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg("赛季" + fmt.Sprintf("%v", user.CardInfo.Season))
	}
	randNum := rand.Intn(cfg1.Weight+cfg2.Weight) + 1
	logger.Debug("randLowOrHighPool before  userId:%v  randNum:%v  user.CardInfo.AddWeigh:%v", user.Id, randNum, user.CardInfo.AddWeight)
	if randNum >= 0 && randNum <= cfg2.Weight+user.CardInfo.AddWeight {
		user.CardInfo.AddWeight = 0
		return cfg2, nil
	}
	user.CardInfo.AddWeight += cfg1.Weight1
	logger.Debug("randLowOrHighPool after   userId:%v user.CardInfo.AddWeigh:%v", user.Id, user.CardInfo.AddWeight)
	return cfg1, nil
}

func (this *CardActivity) buildInfo(user *objs.User) ([]int32, []*pb.CardInfoUnit, []*pb.CardInfoUnit) {
	HaveGetIndex := make([]int32, 0)
	MyDrawInfo := make([]*pb.CardInfoUnit, 0)
	ServerDrawInfo := make([]*pb.CardInfoUnit, 0)
	for _, id := range user.CardInfo.GetAwardIds {
		HaveGetIndex = append(HaveGetIndex, int32(id))
	}

	data, err := modelGame.GetCardModel().GetMyDrawCardInfos(user.CardInfo.Season, user.Id, 50)
	if err != nil {
		logger.Error("GetMyDrawCardInfos err:%v   user.CardInfo.Season:%v, user.Id:%v", err, user.CardInfo.Season, user.Id)
		return HaveGetIndex, MyDrawInfo, ServerDrawInfo
	}
	for _, info := range data {
		MyDrawInfo = append(MyDrawInfo, &pb.CardInfoUnit{ItemId: int32(info.ItemId), Time: int32(info.DrawTime), UserName: info.NickName, Type: int32(info.DrawType), Count: int32(info.Count)})
	}
	cfg := gamedb.GetConf().DrawRecord
	data, err = modelGame.GetCardModel().GetAllDrawCardInfos(user.CardInfo.Season, 20, cfg[0], cfg[1])
	if err != nil {
		logger.Error("GetAllDrawCardInfos err:%v   user.CardInfo.Season:%v, user.Id:%v", err, user.CardInfo.Season, user.Id)
		return HaveGetIndex, MyDrawInfo, ServerDrawInfo
	}
	for _, info := range data {
		ServerDrawInfo = append(ServerDrawInfo, &pb.CardInfoUnit{ItemId: int32(info.ItemId), Time: int32(info.DrawTime), UserName: info.NickName, Type: int32(info.DrawType), Count: int32(info.Count)})
	}
	return HaveGetIndex, MyDrawInfo, ServerDrawInfo
}

func (this *CardActivity) Rest(user *objs.User, isEnter bool) {
	data := common.GetResetTime(time.Now())
	logger.Debug("CardActivity userId:%v  data:%v   user.CardInfo.DayResDay:%v", user.Id, data, user.CardInfo.DayResDay)
	if data != user.CardInfo.DayResDay {
		user.CardInfo.DrawTimes = 0
		user.CardInfo.DayResDay = data
	}

	season, _, _ := this.getDrawSeason(user)
	if season != user.CardInfo.Season {
		//user.CardInfo.Integral = 0
		user.CardInfo.AddWeight = 0
		user.CardInfo.Season = season
		if user.Id != base.Conf.ServerId {
			user.CardInfo.MergeMark = 2
			season1, _, _ := this.getDrawSeason(user)
			user.CardInfo.Season = season1
		}
		this.resetInit()
	}
	user.Dirty = true
	if season == -1 && !isEnter {
		this.BroadcastAll(&pb.CardCloseNtf{IsClose: true})
		return
	}
	if !isEnter {
		ack := &pb.CardActivityInfosAck{}
		err := this.Load(user, ack)
		if err != nil {
			return
		}
		_ = this.GetUserManager().SendMessage(user, ack, true)
	}
}

func (this *CardActivity) resetInit() {
	defer this.mu.Unlock()
	this.mu.Lock()
	this.myCardInfo = make(map[int][]*pb.CardInfoUnit)
	this.serverCardInfo = make([]*pb.CardInfoUnit, 0)
}

func (this *CardActivity) addItemIntoCard(user *objs.User, itemId, count, drawType, drawTime int) {
	cfg := gamedb.GetItemBaseCfg(itemId)
	if cfg == nil {
		logger.Error("addItemIntoCard itemId:%v Get err", itemId)
		return
	}
	this.addToCardInfo(user, drawType, itemId, count)
	item := &modelGame.Card{UserId: user.Id, NickName: user.NickName, DrawTime: drawTime, ItemId: itemId, Count: count, DrawType: drawType, Season: user.CardInfo.Season, ExpireTime: time.Now().Unix() + 86400*7, ItemQuality: cfg.Quality}
	// update db
	this.opChan <- opMsg{constAuction.OpInsert, item}
}

func (this *CardActivity) dbDataOp() {
	logger.Debug("CardActivity dbDataOp run")

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("card dbDataOp panic: %v, time: %v, stack tace: %v\n", err, time.Now(), string(debug.Stack()))
		}
	}()

	for {
		select {
		case msg := <-this.opChan:
			switch msg.opType {
			case constAuction.OpInsert:
				logger.Debug("card insert: %+v", *msg.cardItem)
				err := modelGame.GetCardModel().DbMap().Insert(msg.cardItem)
				if err != nil {
					logger.Error("insert card data: %v err: %v", *msg.cardItem, err)
				}
			case constAuction.OpUpdate:
				logger.Debug("card update: %+v", *msg.cardItem)
				_, err := modelGame.GetCardModel().DbMap().Update(msg.cardItem)
				if err != nil {
					logger.Error("update card data: %v err: %v", *msg.cardItem, err)
				}
			default:
				logger.Error("Unknown db operation type: %v", msg.opType)
			}
		}
	}
}

/**添加到系统消息*/
func (this *CardActivity) addToCardInfo(user *objs.User, types, itemId, count int) {
	defer this.mu.Unlock()
	this.mu.Lock()
	baseCfg := gamedb.GetItemBaseCfg(itemId)
	if baseCfg == nil {
		return
	}

	cfg := gamedb.GetConf().DrawRecord

	now := int(time.Now().Unix())
	cardAtInfoUnit := &pb.CardInfoUnit{
		ItemId:   int32(itemId),
		Count:    int32(count),
		Time:     int32(now),
		UserName: user.NickName,
		Type:     int32(types),
	}
	if this.myCardInfo[user.Id] == nil {
		this.myCardInfo[user.Id] = make([]*pb.CardInfoUnit, 0)
	}

	this.myCardInfo[user.Id] = append(this.myCardInfo[user.Id], cardAtInfoUnit)
	if len(this.myCardInfo[user.Id]) > 20 {
		size := len(this.myCardInfo[user.Id])
		delNum := size - 20
		this.myCardInfo[user.Id] = this.myCardInfo[user.Id][delNum:size]
	}

	if baseCfg.Quality >= cfg[0] && baseCfg.Quality <= cfg[1] {
		this.serverCardInfo = append(this.serverCardInfo, cardAtInfoUnit)
		if len(this.serverCardInfo) > 20 {
			size := len(this.serverCardInfo)
			delNum := size - 20
			this.serverCardInfo = this.serverCardInfo[delNum:size]
		}
	}
	return
}

func (this *CardActivity) getMyCardInfoAndSysInfo(user *objs.User) (haveGetIndex []int32, myCardInfo, serverInfo []*pb.CardInfoUnit) {
	defer this.mu.RUnlock()
	this.mu.RLock()
	HaveGetIndex := make([]int32, 0)
	for _, id := range user.CardInfo.GetAwardIds {
		HaveGetIndex = append(HaveGetIndex, int32(id))
	}
	if this.myCardInfo[user.Id] == nil {
		this.myCardInfo[user.Id] = make([]*pb.CardInfoUnit, 0)
	}

	return HaveGetIndex, this.myCardInfo[user.Id], this.serverCardInfo

}

func (this *CardActivity) initCardInfo() {
	season, _, _ := this.getDrawSeason(nil)

	allBaseUserInfo := this.GetUserManager().GetAllUsersBasicInfo()
	for userId := range allBaseUserInfo {
		if this.myCardInfo[userId] == nil {
			this.myCardInfo[userId] = make([]*pb.CardInfoUnit, 0)
		}
		data, err := modelGame.GetCardModel().GetMyDrawCardInfos(season, userId, 20)
		if err == nil {
			for _, info := range data {
				this.myCardInfo[userId] = append(this.myCardInfo[userId], &pb.CardInfoUnit{ItemId: int32(info.ItemId), Time: int32(info.DrawTime), UserName: info.NickName, Type: int32(info.DrawType), Count: int32(info.Count)})
			}
		}
	}

	cfg := gamedb.GetConf().DrawRecord
	data, err := modelGame.GetCardModel().GetAllDrawCardInfos(season, 20, cfg[0], cfg[1])
	if err == nil {
		for _, info := range data {
			this.serverCardInfo = append(this.serverCardInfo, &pb.CardInfoUnit{ItemId: int32(info.ItemId), Time: int32(info.DrawTime), UserName: info.NickName, Type: int32(info.DrawType), Count: int32(info.Count)})
		}
	}

}
