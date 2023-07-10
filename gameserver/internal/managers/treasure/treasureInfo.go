package treasure

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constAuction"
	"cqserver/gamelibs/publicCon/constMail"
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
	"strconv"
	"time"
)

type Treasure struct {
	util.DefaultModule
	managersI.IModule
	opChan chan opMsg
}

func NewTreasure(module managersI.IModule) *Treasure {
	p := &Treasure{IModule: module, opChan: make(chan opMsg, constAuction.DbChanSize)}
	return p
}

type opMsg struct {
	opType   int
	cardItem *modelGame.Treasure
}

func (this *Treasure) Init() error {
	go this.dbDataOp()
	return nil
}

func (this *Treasure) GetNowSeason(serverId int) int {
	openDay := this.GetTreasureOpenDay(serverId)
	if gamedb.GetTreasureNowSeason(openDay) == nil {
		return -1
	}
	return gamedb.GetTreasureNowSeason(openDay).Type1
}

func (this *Treasure) Load(user *objs.User, ack *pb.TreasureInfosAck) error {
	user.TreasureInfo.Season = this.GetNowSeason(user.ServerId)
	if user.TreasureInfo.Season <= 0 {
		return nil
	}
	haveBuyTimes, allChooseItems, haveGetItems, allGetRound := this.buildInfo(user)
	MyDrawInfo, ServerDrawInfo := this.buildModelInfo(user)
	ack.ChoosItems = allChooseItems
	ack.HaveGetItems = haveGetItems
	ack.HaveBuyTimes = haveBuyTimes
	ack.HaveGetRoundId = allGetRound
	ack.TreasureTimes = int32(user.TreasureInfo.AllUseTimes)
	ack.MyTreasureInfo = MyDrawInfo
	ack.ServerTreasureInfo = ServerDrawInfo
	ack.Season = int32(user.TreasureInfo.Season)
	ack.PopUpState = int32(user.TreasureInfo.PopUpState)
	return nil
}

//
//  SetPopUp
//  @Description: 设置玩家上线弹框提示
//  @receiver this
//  @param state 状态 1， 0
//
func (this *Treasure) SetPopUp(user *objs.User, state int, ack *pb.SetTreasurePopUpStateAck) {

	openDay := this.GetTreasureOpenDay(user.ServerId)
	user.TreasureInfo.PopUpState = state
	user.TreasureInfo.PopUpResOpenDay = openDay
	user.Dirty = true
	ack.State = int32(state)
	return
}

//花费元宝购买寻龙令
func (this *Treasure) BuyTreasureItem(user *objs.User, ack *pb.BuyTreasureItemAck, op *ophelper.OpBagHelperDefault) error {

	cfg := gamedb.GetConf().XunlongBuy[0]
	cfg1 := gamedb.GetConf().XunlongConsume
	if user.TreasureInfo.BuyTimes == nil {
		user.TreasureInfo.BuyTimes = make(map[int]int)
	}
	if user.TreasureInfo.BuyTimes[pb.ITEMID_INGOT] >= cfg[3] {
		return gamedb.ERRTREASURE
	}

	if has, _ := this.GetBag().HasEnough(user, pb.ITEMID_INGOT, cfg[2]); !has {
		return gamedb.ERRNOTENOUGHGOODS
	}
	err := this.GetBag().Remove(user, op, pb.ITEMID_INGOT, cfg[2])
	if err != nil {
		return err
	}
	//op1 := ophelper.NewOpBagHelperDefault(constBag.OpTreasureBuyXunLongLin)
	err = this.GetBag().AddItem(user, op, cfg1[0], cfg[1])
	if err != nil {
		return err
	}
	user.TreasureInfo.BuyTimes[pb.ITEMID_INGOT] += 1

	buyTimes, _, _, _ := this.buildInfo(user)
	ack.HaveBuyTimes = buyTimes
	ack.Goods = op.ToChangeItems()
	return nil
}

//  ChooseTreasureItem
//  @Description: 转盘奖励选择
//  @receiver this
//  @param types 挡位
//  @param itemId
//  @param ack
//  @return error
//
func (this *Treasure) ChooseTreasureItem(user *objs.User, types int, indexReq []int32, isReplace, replaceIndex int, ack *pb.ChooseTreasureAwardAck) error {
	if user.TreasureInfo.Season <= 0 {
		return gamedb.ERRACTIVITYCLOSE
	}

	if len(indexReq) < 0 {
		return gamedb.ERRPARAM
	}

	if user.TreasureInfo.AllUseTimes >= pb.TREASURETYPE_ONE_ROUND_TIMES {
		return gamedb.ERRPARAM

	}

	cfg := gamedb.GetTreasureCfgBySeasonAndType(user.TreasureInfo.Season, types)
	if cfg == nil {
		return gamedb.ERRPARAM
	}
	if user.TreasureInfo.ChooseItems == nil {
		user.TreasureInfo.ChooseItems = make(map[int]model.IntSlice)
	}
	if user.TreasureInfo.HaveRandomItems == nil {
		user.TreasureInfo.HaveRandomItems = make(map[int]model.IntSlice)
	}
	if user.TreasureInfo.ChooseItems[types] == nil {
		user.TreasureInfo.ChooseItems[types] = make([]int, 0)
	}

	if len(indexReq) > 3 {
		return gamedb.ERRTREASURE1
	}

	if user.TreasureInfo.AllUseTimes != 0 && user.TreasureInfo.AllUseTimes%pb.TREASURETYPE_ONE_ROUND_TIMES != 0 {
		return gamedb.ERRTREASURE2
	}
	itemIds := make([]int32, 0)
	for _, index := range indexReq {
		mark := false
		markItemId := -1
		for index1, data := range cfg.Reward {
			if index1 == int(index) {
				mark = true
				markItemId = data[0]
				itemIds = append(itemIds, int32(markItemId))
			}
		}
		if !mark {
			logger.Error("ChooseTreasureItem 参数错误  season:%v types:%v,  index:%v itemId:%v", user.TreasureInfo.Season, types, indexReq, markItemId)
			return gamedb.ERRPARAM
		}
	}
	user.TreasureInfo.ChooseItems[types] = make([]int, 0)
	for _, id := range indexReq {
		user.TreasureInfo.ChooseItems[types] = append(user.TreasureInfo.ChooseItems[types], int(id))
	}
	ack.ItemId = itemIds
	ack.Index = indexReq
	ack.Type = int32(types)
	_, chooseInfo, haveGetItems, _ := this.buildInfo(user)
	ack.ChoosItems = chooseInfo
	ack.HaveGetItems = haveGetItems
	ack.IsReplace = int32(isReplace)
	ack.ReplaceIndex = int32(replaceIndex)
	user.Dirty = true
	return nil
}

func (this *Treasure) GetTreasureIntegralAward(user *objs.User, index int, ack *pb.GetTreasureIntegralAwardAck, op *ophelper.OpBagHelperDefault) error {
	if index <= 0 {
		return gamedb.ERRPARAM
	}
	cfg := gamedb.GetXunlongRoundsXunlongRoundsCfg(index)
	if cfg == nil {
		return gamedb.ERRPARAM
	}
	for _, v := range user.TreasureInfo.AllGetRound {
		if v == index {
			return gamedb.ERRTREASURE6
		}
	}

	if user.TreasureInfo.Season != cfg.Type1 || user.TreasureInfo.AllUseTimes < cfg.Rounds {
		return gamedb.ERRPARAM
	}

	for _, data := range cfg.Reward {
		err := this.GetBag().AddItem(user, op, data.ItemId, data.Count)
		if err != nil {
			logger.Error("GetTreasureIntegralAward AddItem  userId:%v err:%v", user.Id, err)
			return err
		}
	}

	user.TreasureInfo.AllGetRound = append(user.TreasureInfo.AllGetRound, index)
	ack.TreasureTimes = int32(user.TreasureInfo.AllUseTimes)
	_, _, _, allGetRound := this.buildInfo(user)
	ack.HaveGetIndex = allGetRound
	return nil

}

//开始抽奖
func (this *Treasure) ApplyGet(user *objs.User, ack *pb.TreasureApplyGetAck, op *ophelper.OpBagHelperDefault) error {
	if user.TreasureInfo.Season <= 0 {
		return gamedb.ERRACTIVITYCLOSE
	}
	//if user.TreasureInfo.AllUseTimes > 0 {
	//	if user.TreasureInfo.AllUseTimes%pb.TREASURETYPE_ONE_ROUND_TIMES == 0 {
	//		return gamedb.ERRTREASURE4
	//	}
	//}

	allChooseItems := 0
	for _, data := range user.TreasureInfo.ChooseItems {
		if data != nil {
			allChooseItems += len(data)
		}
	}
	if allChooseItems == 0 {
		return gamedb.ERRTREASURE4
	}

	for _, data := range user.TreasureInfo.HaveRandomItems {
		if data != nil {
			allChooseItems += len(data)
		}
	}
	if allChooseItems < pb.TREASURETYPE_ONE_ROUND_TIMES {
		return gamedb.ERRTREASURE3
	}

	if user.TreasureInfo.AllUseTimes > 0 {
		times := (user.TreasureInfo.AllUseTimes + 1) % pb.TREASURETYPE_ONE_ROUND_TIMES
		if times == 0 {
			times = pb.TREASURETYPE_ONE_ROUND_TIMES
		}

		logger.Debug("user.TreasureInfo.AllUseTimes:%v  times:%v", user.TreasureInfo.AllUseTimes, times)
		cfg := gamedb.GetXunlongPrXunlongPrCfg(times)
		if cfg == nil {
			logger.Error("配置错误 GetXunLongPrXunLongPrCfg times:%v", times)
			return gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg("item" + strconv.Itoa(times))
		}

		if has, _ := this.GetBag().HasEnough(user, cfg.Consume.ItemId, cfg.Consume.Count); !has {
			logger.Error("寻龙令不足 cfg.Id:%v  cfg.Consume.ItemId:%v, cfg.Consume.Count:%v", cfg.Time, cfg.Consume.ItemId, cfg.Consume.Count)
			return gamedb.ERRNOTENOUGHGOODS
		}
		err := this.GetBag().Remove(user, op, cfg.Consume.ItemId, cfg.Consume.Count)
		if err != nil {
			return err
		}
	}

	allData, err := this.buildRandomItemInfos(user)
	if err != nil {
		return err
	}
	rand.Seed(time.Now().UnixNano())
	itemId, itemCount, itemIndex, types := common.RandWeightBySlice3(allData)
	if types <= 0 {
		return gamedb.ERRPARAM
	}
	this.addApplyItem(user, types, itemId, itemCount, itemIndex, op)
	//-----
	logger.Debug("itemId:%v", itemId)
	user.TreasureInfo.AllUseTimes += 1
	ack.TreasureTimes = int32(user.TreasureInfo.AllUseTimes)
	ack.Items = append(ack.Items, int32(itemIndex))
	ack.Goods = op.ToChangeItems()
	_, ack.ChoosItems, ack.HaveGetItems, _ = this.buildInfo(user)
	ack.MyTreasureInfo, ack.ServerTreasureInfo = this.buildModelInfo(user)
	ack.RandomType = int32(types)
	//if user.TreasureInfo.AllUseTimes%pb.TREASURETYPE_ONE_ROUND_TIMES == 0 {
	//	//下一轮重置
	//	user.TreasureInfo.ChooseItems = user.TreasureInfo.HaveRandomItems
	//	user.TreasureInfo.HaveRandomItems = make(map[int]model.IntSlice)
	//}
	user.Dirty = true
	_ = this.GetUserManager().SendItemChangeNtf(user, op)
	this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_TREASURE, []int{1})
	return nil
}

func (this *Treasure) addApplyItem(user *objs.User, types, itemId, itemCount, itemIndex int, op *ophelper.OpBagHelperDefault) int {
	itemInfos := make(gamedb.ItemInfos, 0)
	itemInfos = append(itemInfos, &gamedb.ItemInfo{ItemId: itemId, Count: itemCount})
	this.GetBag().AddItems(user, itemInfos, op)
	this.addItemIntoTreasure(user, itemId, itemCount, int(time.Now().Unix()))

	if user.TreasureInfo.HaveRandomItems == nil {
		user.TreasureInfo.HaveRandomItems = make(map[int]model.IntSlice, 0)
	}
	if user.TreasureInfo.HaveRandomItems[types] == nil {
		user.TreasureInfo.HaveRandomItems[types] = make([]int, 0)
	}

	haveRandomItemsIndex := make([]int, 0)
	for _, id := range user.TreasureInfo.HaveRandomItems[types] {
		haveRandomItemsIndex = append(haveRandomItemsIndex, id)
	}
	haveRandomItemsIndex = append(haveRandomItemsIndex, itemIndex)
	user.TreasureInfo.HaveRandomItems[types] = haveRandomItemsIndex
	slices := make([]int, 0)
	for _, v := range user.TreasureInfo.ChooseItems[types] {
		if v != itemIndex {
			slices = append(slices, v)
		}
	}
	user.TreasureInfo.ChooseItems[types] = slices
	return itemIndex
}

//数据组装
func (this *Treasure) buildInfo(user *objs.User) (map[int32]int32, map[int32]*pb.ChooseInfo, map[int32]*pb.ChooseInfo, []int32) {
	haveBuyTimes := make(map[int32]int32)
	allChooseItems := make(map[int32]*pb.ChooseInfo)
	haveGetItems := make(map[int32]*pb.ChooseInfo)
	allGetRound := make([]int32, 0)
	for k, v := range user.TreasureInfo.BuyTimes {
		haveBuyTimes[int32(k)] = int32(v)
	}
	for k, v := range user.TreasureInfo.ChooseItems {
		if allChooseItems[int32(k)] == nil {
			allChooseItems[int32(k)] = &pb.ChooseInfo{}
			allChooseItems[int32(k)].Items = make([]int32, 0)
		}
		for _, v1 := range v {
			allChooseItems[int32(k)].Items = append(allChooseItems[int32(k)].Items, int32(v1))
		}
	}
	for k, v := range user.TreasureInfo.HaveRandomItems {
		if haveGetItems[int32(k)] == nil {
			haveGetItems[int32(k)] = &pb.ChooseInfo{}
			haveGetItems[int32(k)].Items = make([]int32, 0)
		}
		for _, v1 := range v {
			haveGetItems[int32(k)].Items = append(haveGetItems[int32(k)].Items, int32(v1))
		}
	}
	for _, v := range user.TreasureInfo.AllGetRound {
		allGetRound = append(allGetRound, int32(v))
	}

	return haveBuyTimes, allChooseItems, haveGetItems, allGetRound
}

//重置
func (this *Treasure) Reset(user *objs.User, isEnter bool) {
	openDay := this.GetTreasureOpenDay(user.ServerId)
	if openDay != user.TreasureInfo.PopUpResOpenDay {
		user.TreasureInfo.PopUpState = 0
	}
	season := this.GetNowSeason(user.ServerId)
	if user.TreasureInfo.Season != season {
		user.TreasureInfo = &model.TreasureInfo{}
		user.TreasureInfo.Season = season
	}
	user.Dirty = true
	if season == -1 && !isEnter {
		this.BroadcastAll(&pb.TreasureCloseNtf{IsClose: true})
		return
	}
	if !isEnter {
		ack := &pb.TreasureInfosAck{}
		err := this.Load(user, ack)
		if err != nil {
			return
		}
		_ = this.GetUserManager().SendMessage(user, ack, true)
	}
}

/**
 *  @Description: 寻龙探宝校验是否开放，支付金额
 *  @param payNum	支付金额
 *  @return error
 */
func (this *Treasure) PayCheck(user *objs.User, payNum int) error {
	canBuyTimes := 0
	cfg := gamedb.GetConf().XunlongBuy
	for _, data := range cfg {
		if data[0] == 2 {
			if data[2] == payNum {
				canBuyTimes = data[3]
			}
		}
	}
	if canBuyTimes <= 0 {
		logger.Error("寻龙探宝 购买寻龙令 充值回调 err userId:%v canBuyTimes:%v  payNum:%v", user.Id, canBuyTimes, payNum)
		return gamedb.ERRBUYNUM
	}

	if user.TreasureInfo.BuyTimes == nil {
		user.TreasureInfo.BuyTimes = make(map[int]int)
	}
	if user.TreasureInfo.BuyTimes[payNum] >= canBuyTimes {
		return gamedb.ERRPURCHASECAPENOUGH
	}
	return nil
}

func (this *Treasure) PayCallBack(user *objs.User, payNum int, op *ophelper.OpBagHelperDefault) {

	addNum := 0
	cfg := gamedb.GetConf().XunlongBuy
	for _, data := range cfg {
		if data[0] == 2 {
			if data[2] == payNum {
				addNum = data[1]
			}
		}
	}
	if addNum <= 0 {
		logger.Error("寻龙探宝 购买寻龙令 充值回调 err userId:%v addNum:%v  payNum:%v", user.Id, addNum, payNum)
		return
	}
	cfg1 := gamedb.GetConf().XunlongConsume
	logger.Info("充值购买寻龙令回调 userId:%v payNum:%v ItemId:%v addNum:%v", user.Id, payNum, cfg1[0], addNum)
	itemInfos := make(gamedb.ItemInfos, 0)
	itemInfos = append(itemInfos, &gamedb.ItemInfo{ItemId: cfg1[0], Count: addNum})
	this.GetBag().AddItems(user, itemInfos, op)
	if user.TreasureInfo.BuyTimes == nil {
		user.TreasureInfo.BuyTimes = make(map[int]int)
	}
	user.TreasureInfo.BuyTimes[payNum] += 1
	user.Dirty = true
	ack := &pb.BuyTreasureItemAck{}
	buyTimes, _, _, _ := this.buildInfo(user)
	ack.HaveBuyTimes = buyTimes
	this.GetUserManager().SendMessage(user, ack, true)
	return
}

func (this *Treasure) DrawLoad(user *objs.User, ack *pb.TreasureDrawInfoAck) {
	user.TreasureInfo.Season = this.GetNowSeason(user.ServerId)
	MyDrawInfo, ServerDrawInfo := this.buildModelInfo(user)
	ack.MyTreasureInfo = MyDrawInfo
	ack.ServerTreasureInfo = ServerDrawInfo
}

//活动结束,玩家未领取的轮次奖励,邮件发送
func (this *Treasure) SendTreasureMail() {
	openDay := this.GetTreasureOpenDay(base.Conf.ServerId)
	cfg := gamedb.GetXunLongBeforeDaySeason(openDay - 1)
	logger.Info("寻龙探宝 阶段奖励发放 serverId:%v  openDay:%v", base.Conf.ServerId, openDay-1)
	if cfg == nil {
		logger.Info("活动还未结束")
		return
	}
	offlineTimes := gamedb.GetXunLongTypeTime(cfg.Type1)

	allUserInfos := this.GetUserManager().GetAllUsersBasicInfo()
	for _, data := range allUserInfos {
		if int(time.Now().Unix()-data.LastUpdateTime.Unix()) > offlineTimes {
			continue
		}
		userInfo := this.GetUserManager().GetUser(data.Id)
		if userInfo == nil {
			userInfo = this.GetUserManager().GetOfflineUserInfo(data.Id)
		}
		if userInfo == nil {
			continue
		}

		if userInfo.TreasureInfo.AllUseTimes <= 0 {
			continue
		}
		round := userInfo.TreasureInfo.AllUseTimes
		if round <= 0 {
			continue
		}

		itemInfos := gamedb.GetXunLongRoundCfgByType(cfg.Type1, round, userInfo.TreasureInfo.AllGetRound)
		if len(itemInfos) > 0 {
			_ = this.GetMail().SendSystemMailWithItemInfos(userInfo.Id, constMail.MAILTYPE_TREASURE_STAGE_REWARD, nil, itemInfos)
		}
	}
}

//寻龙探宝 获取开服天数
func (this *Treasure) GetTreasureOpenDay(serverId int) int {

	return this.GetSystem().GetServerOpenDaysByServerIdByExcursionTime(serverId, 0)
}

func (this *Treasure) buildRandomItemInfos(user *objs.User) ([][]int, error) {

	applyTimes := user.TreasureInfo.AllUseTimes + 1
	if user.TreasureInfo.AllUseTimes >= 12 {
		applyTimes = (user.TreasureInfo.AllUseTimes + 1) % 12
		if applyTimes == 0 {
			applyTimes = 12
		}
	}
	allData := make([][]int, 0)
	cfg := gamedb.GetXunlongPrXunlongPrCfg(applyTimes)
	if cfg == nil {
		return nil, gamedb.ERRSETTINGNOTFOUND
	}
	for _, data := range cfg.Probability {
		chooseData := user.TreasureInfo.ChooseItems[data[0]]
		logger.Debug("userId:%v  data[0]:%v  chooseData:%v", user.Id, data[0], chooseData)
		for _, index := range chooseData {
			randomInfo := make([]int, 0)
			cfg := gamedb.GetTreasureCfgBySeasonAndType(user.TreasureInfo.Season, data[0])
			if cfg == nil {
				return nil, gamedb.ERRSETTINGNOTFOUND
			}
			//道具Id,道具数量，道具权重,道具索引(对应配置表),奖池类型
			randomInfo = append(randomInfo, cfg.Reward[index][0], cfg.Reward[index][1], data[1], index, data[0])
			allData = append(allData, randomInfo)
		}
	}
	logger.Debug("寻龙权重  第%v次抽奖 allData:%v", applyTimes, allData)
	return allData, nil
}

func (this *Treasure) buildModelInfo(user *objs.User) ([]*pb.TreasureInfoUnit, []*pb.TreasureInfoUnit) {
	MyDrawInfo := make([]*pb.TreasureInfoUnit, 0)
	ServerDrawInfo := make([]*pb.TreasureInfoUnit, 0)

	data, err := modelGame.GetTreasureModel().GetMyTreasureInfos(user.TreasureInfo.Season, user.Id, 50)
	if err != nil {
		logger.Error("GetMyDrawCardInfos err:%v   user.CardInfo.Season:%v, user.Id:%v", err, user.CardInfo.Season, user.Id)
		return MyDrawInfo, ServerDrawInfo
	}
	for _, info := range data {
		MyDrawInfo = append(MyDrawInfo, &pb.TreasureInfoUnit{ItemId: int32(info.ItemId), Count: int32(info.Count), Time: int32(info.DrawTime), UserName: info.NickName})
	}
	cfg := gamedb.GetConf().XunlongRecord
	data, err = modelGame.GetTreasureModel().GetAllTreasureInfos(user.TreasureInfo.Season, 20, cfg[0], cfg[1])
	if err != nil {
		logger.Error("GetAllDrawCardInfos err:%v   user.CardInfo.Season:%v, user.Id:%v", err, user.CardInfo.Season, user.Id)
		return MyDrawInfo, ServerDrawInfo
	}
	for _, info := range data {
		ServerDrawInfo = append(ServerDrawInfo, &pb.TreasureInfoUnit{ItemId: int32(info.ItemId), Count: int32(info.Count), Time: int32(info.DrawTime), UserName: info.NickName})
	}
	return MyDrawInfo, ServerDrawInfo
}

func (this *Treasure) addItemIntoTreasure(user *objs.User, itemId, count, drawTime int) {
	cfg := gamedb.GetItemBaseCfg(itemId)
	if cfg == nil {
		logger.Error("addItemIntoTreasure itemId:%v Get err", itemId)
		return
	}
	item := &modelGame.Treasure{UserId: user.Id, NickName: user.NickName, DrawTime: drawTime, ItemId: itemId, Count: count, Season: user.TreasureInfo.Season, ExpireTime: time.Now().Unix() + 86400*7, ItemQuality: cfg.Quality}
	// update db
	this.opChan <- opMsg{constAuction.OpInsert, item}
}

func (this *Treasure) dbDataOp() {
	logger.Debug("Treasure dbDataOp run")

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
					logger.Error("insert Treasure data: %v err: %v", *msg.cardItem, err)
				}
			case constAuction.OpUpdate:
				logger.Debug("card update: %+v", *msg.cardItem)
				_, err := modelGame.GetCardModel().DbMap().Update(msg.cardItem)
				if err != nil {
					logger.Error("update Treasure data: %v err: %v", *msg.cardItem, err)
				}
			default:
				logger.Error("Unknown db operation type: %v", msg.opType)
			}
		}
	}
}
