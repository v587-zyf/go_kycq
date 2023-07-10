package auction

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constAuction"
	"cqserver/gamelibs/publicCon/constMail"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"database/sql"
	"fmt"
	"math"
	"runtime/debug"
	"strconv"
	"time"
)

type opMsg struct {
	opType      int
	auctionItem *modelGame.AuctionItem
}

type guildOpMsg struct {
	opType      int
	auctionItem *modelGame.GuildAuctionItem
}

func (this *AuctionManager) Init() error {
	// 初始化MaxId
	maxId, err := modelGame.GetAuctionItemModel().GetMaxId()
	if err != nil {
		logger.Info("GetAuctionItemModel().GetMaxId() err:%v", err)
		maxId = 0
	}
	guildMaxId, err := modelGame.GetGuildAuctionItemModel().GetMaxId()
	if err != nil {
		logger.Info("GetGuildAuctionItemModel().GetMaxId() err:%v", err)
		guildMaxId = 0
	}
	this.AuctionMoney = gamedb.GetConf().AuctionMoney
	this.ReturnMoney = pb.ITEMID_INGOT
	this.initAuctionData()
	this.GetIdGenerator().InitWorldNowId(maxId, guildMaxId)
	go this.dbDataOp()
	go this.guildDbDataOp()

	return nil
}

//
//  initWorldAuctionData
//  @Description:初始化 世界and门派拍卖行数据
//
func (this *AuctionManager) initAuctionData() {
	allItems, err := modelGame.GetAuctionItemModel().GetAuctionItem()
	if err != nil {
		logger.Error("initWorldAuctionData|GetAuctionItem error:%v", err)
		return
	}
	for _, item := range allItems {
		this.AddWorldAuction(item)
	}
	allGuildAuctionInfos, err := modelGame.GetGuildAuctionItemModel().GetAllAuctionItems()
	if err != nil {
		fmt.Printf("guild auction GetAllAuctionItems error: %v", err)
		return
	} else {
		for _, info := range allGuildAuctionInfos {
			this.AddGuildAuction(info)
		}
	}

}

//上架世界拍卖行
func (this *AuctionManager) AddWorldAuction(item *modelGame.AuctionItem) {
	this.Lock()
	defer this.Unlock()
	if this.worldAuctionData[item.ItemId] == nil {
		this.worldAuctionData[item.ItemId] = make([]*modelGame.AuctionItem, 0)
	}
	this.worldAuctionData[item.ItemId] = append(this.worldAuctionData[item.ItemId], item)
}

//上架门派拍卖行
func (this *AuctionManager) AddGuildAuction(item *modelGame.GuildAuctionItem) {
	this.Lock()
	defer this.Unlock()
	if this.guildAuctionInfos[item.AuctionGuild] == nil {
		this.guildAuctionInfos[item.AuctionGuild] = make(map[int]*modelGame.GuildAuctionItem)
	}
	this.guildAuctionInfos[item.AuctionGuild][item.Id] = item

}

func (this *AuctionManager) AlterMyBidInfo(userId, auctionType int) {

}

func (this *AuctionManager) addItemIntoAuction(item *modelGame.AuctionItem) {
	this.AddWorldAuction(item)
	// update db
	this.opChan <- opMsg{constAuction.OpInsert, item}
}

func (this *AuctionManager) dbDataOp() {
	logger.Debug("dbDataOp run")

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("world auction dbDataOp panic: %v, time: %v, stack tace: %v\n", err, time.Now(), string(debug.Stack()))
		}
	}()

	for {
		select {
		case msg := <-this.opChan:
			switch msg.opType {
			case constAuction.OpInsert:
				logger.Debug("world auction insert: %+v", *msg.auctionItem)
				err := modelGame.GetAuctionItemModel().DbMap().Insert(msg.auctionItem)
				if err != nil {
					logger.Error("insert auction data: %v err: %v", *msg.auctionItem, err)
				}
			case constAuction.OpUpdate:
				logger.Debug("world auction update: %+v", *msg.auctionItem)
				_, err := modelGame.GetAuctionItemModel().DbMap().Update(msg.auctionItem)
				if err != nil {
					logger.Error("update auction data: %v err: %v", *msg.auctionItem, err)
				}
			default:
				logger.Error("Unknown db operation type: %v", msg.opType)
			}
		}
	}
}

func (this *AuctionManager) guildDbDataOp() {
	logger.Info("guildDbDataOp run")

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("guild auction guildDbDataOp panic: %v, time: %v, stack trace: %v\n", err, time.Now(), string(debug.Stack()))
		}
	}()

	for {
		select {
		case msg := <-this.guildOpChan:
			switch msg.opType {
			case constAuction.OpInsert:
				logger.Debug("guild auction insert: %+v", *msg.auctionItem)
				err := modelGame.GetGuildAuctionItemModel().DbMap().Insert(msg.auctionItem)
				if err != nil {
					logger.Error("insert guild auction data: %v err: %v", *msg.auctionItem, err)
				}
			case constAuction.OpUpdate:
				logger.Debug("guild auction update: %+v", *msg.auctionItem)
				_, err := modelGame.GetGuildAuctionItemModel().DbMap().Update(msg.auctionItem)
				if err != nil {
					logger.Error("update guild auction data: %v err: %v", *msg.auctionItem, err)
				}
			default:
				logger.Error("Unknown guild db operation type: %v", msg.opType)
			}
		}
	}
}

func (this *AuctionManager) ProcessAuctionInfoNtf(auctionType, guildId int, ack *pb.AuctionInfoNtf) *pb.AuctionInfoNtf {
	// constAuction.WorldAuction
	logger.Debug("  ProcessAuctionInfoNtf auctionType:%v, guildId:%v", auctionType, guildId)
	ack.AuctionType = int32(auctionType)
	this.RLock()
	defer this.RUnlock()
	if auctionType == constAuction.WorldAuction {
		if len(this.worldAuctionData) == 0 {
			logger.Debug("拍卖行没有东西")
			return ack
		}
		for _, data := range this.worldAuctionData {
			for _, v := range data {
				ack.AuctionInfos = append(ack.AuctionInfos, this.buildAuctionInfo(v))
			}
		}
	} else {
		if len(this.guildAuctionInfos[guildId]) == 0 {
			logger.Debug("门派拍卖行没有东西  guildId:%v", guildId)
			return ack
		}
		for _, data := range this.guildAuctionInfos[guildId] {
			ack.AuctionInfos = append(ack.AuctionInfos, this.buildGuildAuctionInfo(data))
		}
	}

	return ack
}

func (this *AuctionManager) buildAuctionInfo(info *modelGame.AuctionItem) *pb.AuctionItemInfo {
	temp := &pb.AuctionItemInfo{}
	temp.AuctionId = int64(info.Id)
	temp.ItemId = int32(info.ItemId)
	temp.NowBidUserId = int64(info.NowBidPlayerId)
	temp.NowBidPrice = int32(info.NowBidPrice)
	temp.AuctionTime = int64(info.AuctionTime)
	temp.AuctionDuration = int32(info.AuctionDuration)
	temp.AuctionType = int32(info.AuctionType)
	temp.AuctionSrc = int32(info.AuctionSrc)
	temp.DropState = int32(info.Status)
	temp.ItemCount = int32(info.ItemCount)
	temp.PutAwayPrice = int32(info.PutAwayPrice)

	if temp.NowBidUserId > 0 {
		bidderInfo := this.GetUserManager().GetUserBasicInfo(int(temp.NowBidUserId))
		temp.NowBidderAvatar = bidderInfo.Avatar
		temp.NowBidderNickname = bidderInfo.NickName
	}

	for _, userId := range info.AllBidUsers {
		temp.HaveBidUsers = append(temp.HaveBidUsers, int32(userId))
	}

	if info.ExpireTime > 0 {
		temp.FinBidTimes = int64(info.ExpireTime) - int64(constAuction.AuctionDataSaveDuration)
	}

	return temp
}

func (this *AuctionManager) buildGuildAuctionInfo(info *modelGame.GuildAuctionItem) *pb.AuctionItemInfo {
	temp := &pb.AuctionItemInfo{}
	temp.AuctionId = int64(info.Id)
	temp.ItemId = int32(info.ItemId)
	temp.NowBidUserId = int64(info.NowBidPlayerId)
	temp.NowBidPrice = int32(info.NowBidPrice)
	temp.AuctionTime = int64(info.AuctionTime)
	temp.AuctionDuration = int32(info.AuctionDuration)
	temp.AuctionType = int32(info.AuctionType)
	temp.ItemCount = int32(info.ItemCount)

	if temp.NowBidUserId > 0 {
		bidderInfo := this.GetUserManager().GetUserBasicInfo(int(temp.NowBidUserId))
		temp.NowBidderAvatar = bidderInfo.Avatar
		temp.NowBidderNickname = bidderInfo.NickName
	}
	for _, userId := range info.AllBidUsers {
		temp.HaveBidUsers = append(temp.HaveBidUsers, int32(userId))
	}
	return temp
}

func (this *AuctionManager) buildBidInfo(info *modelGame.AuctionBid) *pb.AuctionBidInfo {
	temp := &pb.AuctionBidInfo{}
	temp.Id = int64(info.Id)
	temp.UserId = int32(info.UserId)
	temp.AuctionId = int32(info.AuctionId)
	temp.AuctionType = int32(info.AuctionType)
	temp.ItemId = int32(info.ItemId)
	temp.FirstBidTime = int64(info.FirstBidTime)
	temp.FinallyBidTime = int64(info.FinallyBidTime)
	//if info.FinalBidUserId == info.UserId && info.Status == constAuction.ItemSold {
	//}
	temp.ExpireTime = int64(info.ExpireTime)
	_, _, state := this.GetAuctionInfoByAuctionId(info.AuctionType, info.AuctionId)
	temp.State = int32(state)

	return temp
}

func (this *AuctionManager) GetAuctionInfoByAuctionId(auctionType, auctionId int) (*pb.AuctionItemInfo, *modelGame.AuctionItem, int) {
	this.Lock()
	defer this.Unlock()
	state := constAuction.ItemSold
	itemInfo := &pb.AuctionItemInfo{}
	if auctionType == constAuction.WorldAuction {
		if len(this.worldAuctionData) == 0 {
			logger.Debug("拍卖行没有东西")
			data, err := modelGame.GetAuctionItemModel().GetAuctionItemByAuctionId(auctionId)
			if err == nil {
				state = data.Status
			}
			return nil, nil, state
		}
		for _, data := range this.worldAuctionData {
			for _, v := range data {
				if auctionId == v.Id {
					itemInfo = this.buildAuctionInfo(v)
					state = v.Status
					return itemInfo, v, state
				}
			}
		}
	}

	if auctionType == constAuction.GuildAuction {
		if len(this.guildAuctionInfos) == 0 {
			logger.Debug("公会拍卖行没有东西")
			data, err := modelGame.GetGuildAuctionItemModel().GetAuctionItemByAuctionId(auctionId)
			if err == nil {
				state = data.Status
			}
			return nil, nil, state
		}
		for _, data := range this.guildAuctionInfos {
			for _, v := range data {
				if auctionId == int(v.Id) {
					itemInfo = this.buildGuildAuctionInfo(v)
					state = v.Status
					return itemInfo, nil, state
				}
			}
		}
	}
	return nil, nil, state
}

func (this *AuctionManager) delFromMemory(auctionType, itemId, auctionId int) {
	this.Lock()
	defer this.Unlock()

	if len(this.worldAuctionData) == 0 || this.worldAuctionData[itemId] == nil {
		logger.Debug("拍卖行没有东西")
		return
	}
	indexMark := -1
	for index, data := range this.worldAuctionData[itemId] {
		if auctionId == int(data.Id) {
			indexMark = index
		}
	}
	if indexMark == -1 {
		logger.Error("delFromMemory 拍卖行达到一口价 del 内存数据 失败  auctionType:%v, itemId:%v, auctionId:%v", auctionType, itemId, auctionId)
		return
	}
	this.worldAuctionData[itemId] = append(this.worldAuctionData[itemId][0:indexMark],
		this.worldAuctionData[itemId][indexMark+1:]...)
	return
}

//
//  processWorldGsBidNtf
//  @Description:处理世界拍卖行玩家竞价
//
func (this *AuctionManager) processWorldGsBidNtf(user *objs.User, req *pb.BidReq, ack *pb.BidNtf, op *ophelper.OpBagHelperDefault) {
	logger.Debug("processWorldGsBidNtf AuctionType:%v  AuctionId:%v", int(req.AuctionType), int(req.AuctionId))
	_, auctionInfo, _ := this.GetAuctionInfoByAuctionId(int(req.AuctionType), int(req.AuctionId))
	if auctionInfo == nil {
		ack.Code = int32(constAuction.CodeItemSold)
		return
	}

	if auctionInfo.AuctionUserId == user.Id {
		ack.Code = int32(constAuction.CodeCanNotBidMyselfItem)
		return
	}

	conf := gamedb.GetAuctionAuctioinCfg(auctionInfo.ItemId)
	if conf == nil {
		logger.Error("cannot find artifact conf, itemId: %v", auctionInfo.ItemId)
		ack.Code = constAuction.CodeUnknown
		return
	}

	newPrice := auctionInfo.NowBidPrice + (conf.Price3 * auctionInfo.ItemCount)
	if auctionInfo.NowBidPrice == 0 {
		//newPrice = auctionInfo.PutAwayPrice + (conf.Price3 * auctionInfo.ItemCount)
		newPrice = auctionInfo.PutAwayPrice
	}

	//超过一口价
	if newPrice > conf.Price1*auctionInfo.ItemCount {
		newPrice = conf.Price1 * auctionInfo.ItemCount
	}
	if user.GoldIngot < newPrice {
		ack.Code = constAuction.CodeMiNoEnough
		return
	}

	isFirst := false
	if auctionInfo.NowBidPlayerId <= 0 {
		isFirst = true
	}

	// 检查是否在正常竞标时间内
	nowTs := time.Now().Unix()
	if nowTs > int64(auctionInfo.AuctionTime)+int64(auctionInfo.AuctionDuration) {
		logger.Warn("world auction not in auction time, in preparation time, userId: %v, auctionId: %v", user.Id, req.AuctionId)
		ack.Code = constAuction.CodeNotInAuctionTime
		return
	}

	// 检查是否超过起拍价
	if newPrice < auctionInfo.PutAwayPrice {
		logger.Warn("bid price if less than start price, start price: %v, bid price: %v", auctionInfo.PutAwayPrice, newPrice)
		ack.Code = constAuction.CodeBidPriceLowerThanStartPrice
		return
	}

	//超过一口价 或 一口价购买
	if newPrice >= conf.Price1*auctionInfo.ItemCount || req.IsBuyNow == 1 {
		newPrice = conf.Price1 * auctionInfo.ItemCount
	}

	err := this.GetBag().Remove(user, op, this.AuctionMoney, newPrice)
	if err != nil {
		ack.Code = constAuction.CodeMiNoEnough
		return
	}

	this.SendAuctionRedPointNtf(auctionInfo.AuctionUserId, pb.REDPOINTTYPE_UP_ITEM_BE_BID, pb.REDPOINTSTATE_BRIGHT)

	if conf.Price1 > 0 { // 如果拍卖物品有一口价则判断是否达到一口价
		if newPrice >= conf.Price1*auctionInfo.ItemCount || req.IsBuyNow == 1 { // 如果达到一口价或者一口价
			ack.Code = constAuction.CodeReachBuyNowPrice
			ack.IsBuyNow = 1
			// 更新内存和数据库
			if !isFirst {
				auctionInfo.LastBidPlayerId = auctionInfo.NowBidPlayerId
				auctionInfo.LastBidPrice = auctionInfo.NowBidPrice
			}
			auctionInfo.NowBidPrice = conf.Price1 * auctionInfo.ItemCount
			auctionInfo.NowBidPlayerId = user.Id
			auctionInfo.Status = constAuction.ItemSold
			auctionInfo.ExpireTime = int(time.Now().Unix() + constAuction.AuctionDataSaveDuration)
			if !this.checkUserExists(auctionInfo.AllBidUsers, user.Id) {
				auctionInfo.AllBidUsers = append(auctionInfo.AllBidUsers, user.Id)
			}

			this.opChan <- opMsg{constAuction.OpUpdate, auctionInfo}

			ack.AuctionInfo = this.buildAuctionInfo(auctionInfo)

			this.NotifyAllBidUsersUpdateBidInfo(int(req.AuctionType), int(req.AuctionId), int(auctionInfo.ItemId), int(auctionInfo.LastBidPlayerId), int(auctionInfo.LastBidPrice), constAuction.CodeItemSold, auctionInfo.AllBidUsers, false)

			//给上架玩家元宝
			this.AuctionItemEndNtf(constAuction.AuctionResultSuccess, 0, 0, 0, auctionInfo)
			//给最后竞拍者道具
			this.SendUserItem(auctionInfo.AuctionUserId, auctionInfo.NowBidPlayerId, auctionInfo.ItemId, auctionInfo.ItemCount, auctionInfo.NowBidPrice, false, false)
			this.UpdateUserBidInfo(user.Id, auctionInfo.Status, auctionInfo.Id, auctionInfo.AuctionType, auctionInfo.ItemId, auctionInfo.ItemCount)
			//del 内存数据
			this.delFromMemory(int(req.AuctionType), int(auctionInfo.ItemId), int(req.AuctionId))
			this.SendAuctionRedPointNtf(auctionInfo.AuctionUserId, pb.REDPOINTTYPE_UP_ITEM_SOLD, pb.REDPOINTSTATE_BRIGHT)
			return
		}
	}

	if auctionInfo.NowBidPrice >= newPrice { // 如果竞价被超越
		ack.Code = constAuction.CodePriceBeyond
		ack.AuctionInfo = this.buildAuctionInfo(auctionInfo)
	} else { // 竞价成功，更新拍卖物品信息
		ack.Code = constAuction.CodeSuccess
		if !isFirst {
			auctionInfo.LastBidPlayerId = auctionInfo.NowBidPlayerId
			auctionInfo.LastBidPrice = auctionInfo.NowBidPrice
		}
		auctionInfo.NowBidPrice = newPrice
		auctionInfo.NowBidPlayerId = user.Id

		if !this.checkUserExists(auctionInfo.AllBidUsers, int(user.Id)) {
			auctionInfo.AllBidUsers = append(auctionInfo.AllBidUsers, int(user.Id))
		}

		this.opChan <- opMsg{constAuction.OpUpdate, auctionInfo}

	}
	this.UpdateUserBidInfo(user.Id, auctionInfo.Status, auctionInfo.Id, auctionInfo.AuctionType, auctionInfo.ItemId, auctionInfo.ItemCount)
	this.NotifyAllBidUsersUpdateBidInfo(int(req.AuctionType), int(req.AuctionId), int(auctionInfo.ItemId), int(auctionInfo.LastBidPlayerId), int(auctionInfo.LastBidPrice), constAuction.CodePriceBeyond, auctionInfo.AllBidUsers, false)
	ack.AuctionInfo = this.buildAuctionInfo(auctionInfo)
	return
}

func (this *AuctionManager) getGuildAuctionInfos(guildId, auctionId int) *modelGame.GuildAuctionItem {
	this.Lock()
	defer this.Unlock()
	if this.guildAuctionInfos[guildId] == nil {
		return nil
	}
	if _, ok := this.guildAuctionInfos[guildId][auctionId]; !ok {
		return nil
	}
	return this.guildAuctionInfos[guildId][auctionId]
}

//
//  processGuidAuctionGsBidNtf
//  @Description:处理门派拍卖行玩家竞价
//
func (this *AuctionManager) processGuidAuctionGsBidNtf(user *objs.User, req *pb.BidReq, ack *pb.BidNtf, op *ophelper.OpBagHelperDefault) {
	ack.AuctionType = req.AuctionType
	ack.IsBuyNow = req.IsBuyNow
	guildId := user.GuildData.NowGuildId
	auctionInfo := this.getGuildAuctionInfos(guildId, int(req.AuctionId))

	if auctionInfo == nil {
		ack.Code = constAuction.CodeItemSold
		return
	}
	isFirst := false
	if auctionInfo.NowBidPlayerId <= 0 {
		isFirst = true
	}

	conf := gamedb.GetGuildAuctionGuildAuctionCfg(int(auctionInfo.ItemId))
	if conf == nil {
		logger.Error("cannot find artifact conf, itemId: %v", auctionInfo.ItemId)
		ack.Code = constAuction.CodeUnknown
		return
	}

	newPrice := auctionInfo.NowBidPrice + conf.Price3*auctionInfo.ItemCount
	if auctionInfo.NowBidPrice == 0 {
		//newPrice = (conf.Price2 + conf.Price3) * auctionInfo.ItemCount
		newPrice = conf.Price2 * auctionInfo.ItemCount
	}

	if newPrice > conf.Price1*auctionInfo.ItemCount {
		newPrice = conf.Price1 * auctionInfo.ItemCount
	}

	if user.GoldIngot < newPrice {
		ack.Code = constAuction.CodeMiNoEnough
		return
	}

	// 检查是否在正常竞标时间内
	nowTs := time.Now().Unix()
	if nowTs > auctionInfo.AuctionTime+int64(auctionInfo.AuctionDuration) {
		logger.Debug("guild auction not in auction time, in preparation time, userId: %v, auctionId: %v", user.Id, req.AuctionId)
		ack.Code = constAuction.CodeNotInAuctionTime
		return
	}

	// 检查是否超过起拍价
	if newPrice < conf.Price2*auctionInfo.ItemCount {
		logger.Debug("bid price if less than start price, start price: %v, bid price: %v", conf.Price2*auctionInfo.ItemCount, newPrice)
		ack.Code = constAuction.CodeBidPriceLowerThanStartPrice
		return
	}

	if newPrice >= conf.Price1*auctionInfo.ItemCount || req.IsBuyNow == 1 {
		newPrice = conf.Price1 * auctionInfo.ItemCount
	}

	err := this.GetBag().Remove(user, op, this.AuctionMoney, newPrice)
	if err != nil {
		ack.Code = constAuction.CodeMiNoEnough
		return
	}

	if conf.Price1 > 0 { // 如果有一口价则判断是否达到一口价
		if newPrice >= conf.Price1*auctionInfo.ItemCount || req.IsBuyNow == 1 {
			ack.Code = constAuction.CodeReachBuyNowPrice
			ack.IsBuyNow = 1

			// update memory and db
			if !isFirst {
				auctionInfo.LastBidPlayerId = auctionInfo.NowBidPlayerId
				auctionInfo.LastBidPrice = auctionInfo.NowBidPrice
			}
			auctionInfo.NowBidPrice = conf.Price1 * auctionInfo.ItemCount
			auctionInfo.NowBidPlayerId = user.Id
			auctionInfo.Status = constAuction.ItemSold
			auctionInfo.ExpireTime = int(time.Now().Unix() + constAuction.AuctionDataSaveDuration)
			if !this.checkUserExists(auctionInfo.AllBidUsers, int(user.Id)) {
				auctionInfo.AllBidUsers = append(auctionInfo.AllBidUsers, int(user.Id))
			}

			// 一口价玩家获得道具ntf

			this.guildOpChan <- guildOpMsg{constAuction.OpUpdate, auctionInfo}

			ack.AuctionInfo = this.buildGuildAuctionInfo(auctionInfo)

			this.NotifyAllBidUsersUpdateBidInfo(int(req.AuctionType), int(req.AuctionId), auctionInfo.ItemId, auctionInfo.LastBidPlayerId, auctionInfo.LastBidPrice, constAuction.CodeItemSold, auctionInfo.AllBidUsers, true)

			guildUsers := this.GetCanGetRedRewardUses(guildId, auctionInfo)
			// 门派分红
			returnNum := this.CalculateFenHonNum(guildId, auctionInfo.NowBidPrice, len(guildUsers))
			logger.Info("CalculateFenHonNum  returnNum:%v, guildUsers:%v  auctionInfo.NowBidPrice:%v", returnNum, guildUsers, auctionInfo.NowBidPrice)
			if guildUsers == nil {
				logger.Error("门派分红 getGuildMember nil  id:%v guildId:%v", auctionInfo.Id, auctionInfo.AuctionGuild)
			}

			if guildUsers != nil && returnNum > 0 {
				for _, userId := range guildUsers {
					this.AuctionItemEndNtf(constAuction.AuctionResultSuccess, userId, returnNum, auctionInfo.ItemId, nil)
				}
			}
			//给最后竞拍者道具
			this.SendUserItem(auctionInfo.DropState, int(auctionInfo.NowBidPlayerId), auctionInfo.ItemId, auctionInfo.ItemCount, auctionInfo.NowBidPrice, false, true)
			this.UpdateUserBidInfo(user.Id, auctionInfo.Status, auctionInfo.Id, auctionInfo.AuctionType, auctionInfo.ItemId, auctionInfo.ItemCount)
			delete(this.guildAuctionInfos[guildId], int(req.AuctionId))
			if len(this.guildAuctionInfos[guildId]) <= 0 {
				this.BroadcastAuctionRedPointNtf(guildId, pb.REDPOINTTYPE_OWN_GUILD_AUCTION_NO_ITEMS, pb.REDPOINTSTATE_NO_BRIGHT)
			}
			return
		}
	}

	if auctionInfo.NowBidPrice >= newPrice { // 竞价已被超越
		ack.Code = constAuction.CodePriceBeyond
		ack.AuctionInfo = this.buildGuildAuctionInfo(auctionInfo)
	} else { // 竞价成功，更新竞价物品的信息
		ack.Code = constAuction.CodeSuccess
		if !isFirst {
			auctionInfo.LastBidPlayerId = auctionInfo.NowBidPlayerId
			auctionInfo.LastBidPrice = auctionInfo.NowBidPrice
		}
		auctionInfo.NowBidPrice = newPrice
		auctionInfo.NowBidPlayerId = user.Id
		if ok, delta := this.calRemainSecsByNow(auctionInfo.AuctionTime, auctionInfo.AuctionDuration); ok {
			auctionInfo.AuctionDuration = delta
		}

		if !this.checkUserExists(auctionInfo.AllBidUsers, user.Id) {
			auctionInfo.AllBidUsers = append(auctionInfo.AllBidUsers, user.Id)
		}

		this.guildOpChan <- guildOpMsg{constAuction.OpUpdate, auctionInfo}

	}
	this.UpdateUserBidInfo(user.Id, auctionInfo.Status, auctionInfo.Id, auctionInfo.AuctionType, auctionInfo.ItemId, auctionInfo.ItemCount)
	this.NotifyAllBidUsersUpdateBidInfo(int(req.AuctionType), int(req.AuctionId), int(auctionInfo.ItemId), int(auctionInfo.LastBidPlayerId), int(auctionInfo.LastBidPrice), constAuction.CodePriceBeyond, auctionInfo.AllBidUsers, true)
	ack.AuctionInfo = this.buildGuildAuctionInfo(auctionInfo)
	return
}

func (this *AuctionManager) checkUserExists(userIds []int, userId int) bool {
	for _, v := range userIds {
		if v == userId {
			return true
		}
	}

	return false
}

// 通知已竞拍玩家新的拍卖信息
func (this *AuctionManager) NotifyAllBidUsersUpdateBidInfo(auctionType, auctionId, itemId, lastBidPlayerId, lastBidPrice, code int, allBidUsers []int, isGuildAuction bool) {
	ntf := &pb.BidItemUpdateNtf{}
	ntf.AuctionId = int32(auctionId)
	ntf.ItemStatus = int32(code)
	ntf.NewInfo, _, _ = this.GetAuctionInfoByAuctionId(auctionType, auctionId)
	ntf.LastBidUserId = int32(lastBidPlayerId)
	ntf.AuctionType = int32(auctionType)
	logger.Debug("NotifyAllBidUsersUpdateBidInfo lastBidPlayerId:%v  lastBidPrice:%v  auctionType:%v, auctionId:%v", lastBidPlayerId, lastBidPrice,
		auctionType, auctionId)
	if lastBidPlayerId > 0 && lastBidPrice > 0 {

		logger.Info("return back userId: %v mi num: %v", lastBidPlayerId, lastBidPrice)
		if isGuildAuction {
			conf := gamedb.GetGuildAuctionGuildAuctionCfg(itemId)
			if conf == nil {
				logger.Error("cannot get artifact quality conf, itemId: %v, cannot return back user mi", itemId)
			} else {
				returnNum := lastBidPrice
				returnItem := gamedb.ItemInfos{&gamedb.ItemInfo{ItemId: this.AuctionMoney, Count: returnNum}}
				this.GetMail().SendSystemMailWithItemInfos(lastBidPlayerId, constMail.MAILTYPE_AUCTION_BID_RETURN_MI, []string{strconv.Itoa(returnNum)}, returnItem)
			}
		} else {
			conf := gamedb.GetAuctionAuctioinCfg(itemId)
			if conf == nil {
				logger.Error("cannot get artifact quality conf, itemId: %v, cannot return back user mi", itemId)
			} else {
				returnNum := lastBidPrice
				returnItem := gamedb.ItemInfos{&gamedb.ItemInfo{ItemId: this.AuctionMoney, Count: returnNum}}
				this.GetMail().SendSystemMailWithItemInfos(lastBidPlayerId, constMail.MAILTYPE_AUCTION_BID_RETURN_MI, []string{strconv.Itoa(returnNum)}, returnItem)
			}
		}
	}

	for _, userId := range allBidUsers {
		this.GetUserManager().SendMessageByUserId(userId, ntf)
	}
}

//玩家上架物品被拍掉,给他钱
func (this *AuctionManager) AuctionItemEndNtf(auctionResult, userId, returnNum, auctionItemId int, auctionInfo *modelGame.AuctionItem) {
	logger.Debug("HandleCrossAuctionItemEndNtf run")

	switch auctionResult {
	case constAuction.AuctionResultSuccess:

		//门派分红
		if auctionInfo == nil {
			conf := gamedb.GetGuildAuctionGuildAuctionCfg(auctionItemId)
			if conf == nil {
				logger.Error("cannot get artifact quality conf, itemId: %v, cannot send user mi", auctionItemId)
			} else {
				getNum := this.buildGoldIngotCalc(returnNum)
				returnItem := gamedb.ItemInfos{&gamedb.ItemInfo{ItemId: this.ReturnMoney, Count: getNum}}
				this.GetMail().SendSystemMailWithItemInfos(userId, constMail.MAILTYPE_AUCTION_FENHONG, []string{strconv.Itoa(returnNum)}, returnItem)
			}
			return
		}
		logger.Info("userId: %v, auction auctionId: %v sales success, return mi", auctionInfo.AuctionUserId, auctionInfo.Id)
		// 拍卖成功，发放拍卖玩家元宝
		if auctionInfo.AuctionType == constAuction.WorldAuction && auctionInfo.AuctionSrc == constAuction.AuctionSrcPlayer { // 玩家拍卖的物品
			money := auctionInfo.NowBidPrice - int(math.Ceil(float64(auctionInfo.NowBidPrice*(gamedb.GetConf().AuctionWorldTax))/10000.0))
			logger.Info("money:%v", money)
			aUser := this.GetUserManager().GetAllUserInfoIncludeOfflineUser(auctionInfo.AuctionUserId)
			if aUser != nil {
				if monthCardPrivilege := this.GetVipManager().GetPrivilege(aUser, pb.VIPPRIVILEGE_AUCTION_SERVICE_CHARGE); monthCardPrivilege != 0 {
					count := int(math.Ceil(float64(auctionInfo.NowBidPrice*(gamedb.GetConf().AuctionWorldTax))/10000.0) * (1 - (float64(monthCardPrivilege) / float64(10000))))
					money = auctionInfo.NowBidPrice - count
				}
				this.GetWarOrder().WriteWarOrderTask(aUser, pb.WARORDERCONDITION_AUCTION_SELL, []int{1})
			} else {
				n := rmodel.WarOrder.GetTask(pb.WARORDERCONDITION_AUCTION_SELL, userId)
				if n != 0 {
					n += 1
				} else {
					n = 1
				}
				rmodel.WarOrder.SetTask(pb.WARORDERCONDITION_AUCTION_SELL, userId, n)
			}
			returnNum := money
			logger.Info("send userId: %v mi num: %v  auctionInfo.NowBidPrice:%v  money:%v", auctionInfo.AuctionUserId, returnNum, auctionInfo.NowBidPrice, money)
			itemId := auctionInfo.ItemId
			userId := auctionInfo.AuctionUserId
			conf := gamedb.GetAuctionAuctioinCfg(itemId)
			if conf == nil {
				logger.Error("cannot get artifact quality conf, itemId: %v, cannot send user mi", itemId)
			} else {
				getNum := this.buildGoldIngotCalc(returnNum)
				returnItem := gamedb.ItemInfos{&gamedb.ItemInfo{ItemId: this.ReturnMoney, Count: getNum}}
				this.GetMail().SendSystemMailWithItemInfos(int(auctionInfo.AuctionUserId), constMail.MAILTYPE_AUCTION_SUCCESS, []string{strconv.Itoa(returnNum)}, returnItem)
			}
			rmodel.Auction.RemAuctionTogether(userId, itemId)
			return
		}
	}
	return
}

// 计算拍卖剩余时间是否小于配置的时间
// return: true, 重置剩余配置的时间段所需的auctionDuration 或者 false, -1
func (this *AuctionManager) calRemainSecsByNow(auctionTs int64, auctionDuration int) (bool, int) {
	nowTs := time.Now().Unix()
	limitTs := auctionTs + int64(auctionDuration)
	deltaT := int(limitTs - nowTs)
	if deltaT < 30 {
		return true, int(nowTs + 30 - auctionTs)
	}

	return false, -1
}

// 定时检查在拍卖的物品是否拍卖时间结束
func (this *AuctionManager) CheckAuctionItemTask() {
	//logger.Debug("CheckAuctionItemTask")
	nowTs := time.Now().Unix()
	this.checkAllAuctionItem(nowTs)
	this.checkGuildAuctionItem(nowTs)
}

func (this *AuctionManager) checkAllAuctionItem(limitTs int64) {
	this.Lock()
	defer this.Unlock()
	if len(this.worldAuctionData) == 0 {
		return
	}

	for itemId, data := range this.worldAuctionData {
		temp := make([]*modelGame.AuctionItem, 0)
		for _, info := range data {
			if int64(info.AuctionTime)+int64(info.AuctionDuration) <= limitTs { // 拍卖时间结束
				if info.NowBidPlayerId > 0 { // 有人竞拍
					if info.AuctionSrc == constAuction.AuctionSrcPlayer { // 玩家拍卖
						// 发放拍卖玩家元宝
						this.AuctionItemEndNtf(constAuction.AuctionResultSuccess, 0, 0, 0, info)
					}
					// 发放竞拍玩家道具
					this.SendUserItem(info.AuctionUserId, info.NowBidPlayerId, info.ItemId, info.ItemCount, info.NowBidPrice, false, false)
					info.Status = constAuction.ItemSold
					info.ExpireTime = int(limitTs + constAuction.AuctionDataSaveDuration)
					this.UpdateUserBidInfo(info.NowBidPlayerId, info.Status, info.Id, info.AuctionType, info.ItemId, info.ItemCount)

				} else { // 流拍
					if info.AuctionSrc == constAuction.AuctionSrcPlayer && info.AuctionType == constAuction.WorldAuction {
						this.SendUserItem(info.AuctionUserId, info.AuctionUserId, info.ItemId, info.ItemCount, info.NowBidPrice, true, false)
						info.Status = constAuction.PassIn
						info.ExpireTime = int(limitTs + constAuction.AuctionDataSaveDuration)
					} else { // 流拍
						info.Status = constAuction.PassIn
						info.ExpireTime = int(limitTs + constAuction.AuctionDataSaveDuration)
					}
					//this.UpdateUserBidInfo(0, info.Status, info.Id, info.AuctionType, info.ItemId)
				}
				// update db
				logger.Debug("============>delete item: %+v", info)
				this.opChan <- opMsg{constAuction.OpUpdate, info}
				this.SendAuctionRedPointNtf(info.AuctionUserId, pb.REDPOINTTYPE_UP_ITEM_SOLD, pb.REDPOINTSTATE_BRIGHT)
			} else { // 保留没过期的数据
				temp = append(temp, info)
			}
		}
		this.worldAuctionData[itemId] = temp
	}
}

// 定时检查门派拍卖物品拍卖时间
func (this *AuctionManager) checkGuildAuctionItem(nowTs int64) {
	this.Lock()
	defer this.Unlock()
	if len(this.guildAuctionInfos) == 0 {
		return
	}

	for gd, info := range this.guildAuctionInfos {
		for aId, auctionInfo := range info {

			if auctionInfo.AuctionTime+int64(auctionInfo.AuctionDuration) <= nowTs { // 拍卖时间结束
				if auctionInfo.NowBidPlayerId > 0 { // 有人竞拍
					// 门派分红
					nowUserInfo := this.GetUserManager().GetAllUserInfoIncludeOfflineUser(auctionInfo.NowBidPlayerId)
					if nowUserInfo == nil {
						logger.Error("GetAllUserInfoIncludeOfflineUser nil  userId:%v", auctionInfo.NowBidPlayerId)
						continue
					}
					guildUsers := this.GetCanGetRedRewardUses(nowUserInfo.GuildData.NowGuildId, auctionInfo)
					returnNum := this.CalculateFenHonNum(nowUserInfo.GuildData.NowGuildId, auctionInfo.NowBidPrice, len(guildUsers))
					logger.Info("CalculateFenHonNum  returnNum:%v, guildUsers:%v  auctionInfo.NowBidPrice:%v", returnNum, guildUsers, auctionInfo.NowBidPrice)
					if guildUsers == nil {
						logger.Error("门派分红 getGuildMember nil  id:%v guildId:%v", auctionInfo.Id, nowUserInfo.GuildData.NowGuildId)
						continue
					}
					if returnNum > 0 {
						for _, userId := range guildUsers {
							this.AuctionItemEndNtf(constAuction.AuctionResultSuccess, userId, returnNum, auctionInfo.ItemId, nil)
						}
					}
					//给最后竞拍者道具
					this.SendUserItem(auctionInfo.DropState, auctionInfo.NowBidPlayerId, auctionInfo.ItemId, auctionInfo.ItemCount, auctionInfo.NowBidPrice, false, true)

					auctionInfo.Status = constAuction.ItemSold
					auctionInfo.ExpireTime = int(nowTs) + constAuction.AuctionDataSaveDuration
					this.UpdateUserBidInfo(int(auctionInfo.NowBidPlayerId), auctionInfo.Status, auctionInfo.Id, auctionInfo.AuctionType, auctionInfo.ItemId, auctionInfo.ItemCount)
				} else { // 流拍
					auctionInfo.Status = constAuction.PassIn
					auctionInfo.ExpireTime = int(nowTs) + constAuction.AuctionDataSaveDuration
					//this.UpdateUserBidInfo(0, auctionInfo.Status, auctionInfo.Id, auctionInfo.AuctionType, auctionInfo.ItemId)
				}
				// delete item
				this.guildOpChan <- guildOpMsg{constAuction.OpUpdate, auctionInfo}
				delete(this.guildAuctionInfos[gd], aId)
				if len(this.guildAuctionInfos[gd]) <= 0 {
					this.BroadcastAuctionRedPointNtf(gd, pb.REDPOINTTYPE_OWN_GUILD_AUCTION_NO_ITEMS, pb.REDPOINTSTATE_NO_BRIGHT)
				}
			}
		}
	}
}

//给玩家物品
func (this *AuctionManager) SendUserItem(auctionUserId, getUserId, itemId, itemCount, nowBidPrice int, isBeat, isGuildAuction bool) {
	logger.Debug("SendUserItem run  getUserId:%v, itemId:%v isBeat:%v  itemCount:%v", getUserId, itemId, isBeat, itemCount)
	returnItem := gamedb.ItemInfos{&gamedb.ItemInfo{ItemId: itemId, Count: itemCount}}

	mailType := constMail.MAILTYPR_BID_SUCCESS
	if isBeat {
		mailType = constMail.MAILTYPE_AUCTION_LIUPAI
	}
	if isGuildAuction {
		conf := gamedb.GetGuildAuctionGuildAuctionCfg(itemId)
		if conf != nil {
			this.GetMail().SendSystemMailWithItemInfos(getUserId, mailType, []string{conf.Name}, returnItem)
		}
	} else {
		conf := gamedb.GetAuctionAuctioinCfg(itemId)
		if conf != nil {
			this.GetMail().SendSystemMailWithItemInfos(getUserId, mailType, []string{conf.Name}, returnItem)
		}
	}

	getUserInfo := this.GetUserManager().GetAllUserInfoIncludeOfflineUser(getUserId)
	auctionType := constAuction.WorldAuction
	if isGuildAuction {
		auctionType = constAuction.GuildAuction
	}

	kyEvent.Auction(getUserInfo, auctionUserId, getUserId, auctionType, itemId, itemCount, nowBidPrice)

	if !isBeat {
		user := this.GetUserManager().GetUser(getUserId)
		if user != nil {
			this.GetWarOrder().WriteWarOrderTask(user, pb.WARORDERCONDITION_AUCTION_BUY, []int{1})
		} else {
			n := rmodel.WarOrder.GetTask(pb.WARORDERCONDITION_AUCTION_BUY, getUserId)
			if n != 0 {
				n += 1
			} else {
				n = 1
			}
			rmodel.WarOrder.SetTask(pb.WARORDERCONDITION_AUCTION_BUY, getUserId, n)
		}
	}
}

//  DropItemToGuildAuction
//  @Description:  处理门派神兵掉落
//  @param guildId 门派id
//  @param itemId 物品id
//  @param count 数量
//  @param canGetRedReward   可以获得分红的人
//
func (this *AuctionManager) DropItemToGuildAuction(guildId, itemId, count, dropState int, canGetRedReward model.IntSlice) {
	logger.Debug("HandleGsGuildAuctionDropNtf run")

	if this.GetGuild().GetGuildInfo(guildId) == nil {
		logger.Error("添加物品进公会拍卖行 公会:%v不存在", guildId)
		return
	}

	nowTs := time.Now().Unix()
	conf := gamedb.GetGuildAuctionGuildAuctionCfg(itemId)
	itemCfg := gamedb.GetItemBaseCfg(itemId)
	if conf != nil && itemCfg != nil {
		info := &modelGame.GuildAuctionItem{}
		info.Id = this.GetIdGenerator().GetNextGuildId()
		info.ItemId = itemId
		info.ItemCount = count
		info.AuctionTime = nowTs
		info.AuctionDuration = conf.Time * 60
		info.Status = constAuction.OnAuction
		info.AuctionGuild = guildId
		info.AuctionType = constAuction.GuildAuction
		info.CanGetRedAward = canGetRedReward
		info.DropState = dropState
		// 更新数据库
		this.guildOpChan <- guildOpMsg{constAuction.OpInsert, info}
		this.AddGuildAuction(info)
		kyEvent.AuctionUp(dropState, conf.Price2*count, itemId, count, constAuction.GuildAuction)
	} else {
		logger.Error("GetAuctionConfById error, conf is nil, itemId: %v", itemId)
	}
}

//计算门派拍卖行分红
func (this *AuctionManager) CalculateFenHonNum(guildId, nowBidPrice, canGetRedNum int) int {
	// 门派分红
	if canGetRedNum <= 0 {
		logger.Info("CalculateFenHonNum  canGetRedNum:%v", canGetRedNum)
		return 0
	}
	returnNum := int(math.Ceil(float64(nowBidPrice) * float64(1-float64(gamedb.GetConf().AuctionUnionTax)/10000.0) * float64(float64(gamedb.GetConf().AuctionShare)/10000.0) / float64(canGetRedNum)))
	logger.Info("guidId:%v  nowBidPrice:%v  金锭:returnNum:%v", guildId, nowBidPrice, returnNum)
	return returnNum
}

func (this *AuctionManager) UpdateUserBidInfo(bidUserId, state, auctionId, auctionType, itemId, count int) {

	bidId, err := modelGame.GetAuctionBidModel().GetUserBidIdByUserIdAndAuctionId(bidUserId, auctionId, auctionType)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("GetUserBidIdByUserIdAndAuctionId err: %v, auctionId: %d, auctionType: %v", err, auctionId, auctionType)
		return
	}
	if err == sql.ErrNoRows || bidId == nil {
		bidInfo := &modelGame.AuctionBid{}
		bidInfo.AuctionId = auctionId
		bidInfo.ItemId = itemId
		bidInfo.ItemCount = count
		bidInfo.UserId = bidUserId
		bidInfo.FirstBidTime = time.Now().Unix()
		bidInfo.AuctionType = auctionType
		bidInfo.Status = state
		bidInfo.FinalBidUserId = bidUserId
		bidInfo.ExpireTime = int(time.Now().Unix() + constAuction.AuctionDataSaveDuration)
		bidInfo.FinallyBidTime = time.Now().Unix()
		err := modelGame.GetAuctionBidModel().DbMap().Insert(bidInfo)
		if err != nil {
			logger.Error("insert err: %v, insert info: %v", *bidInfo)
		}
		return
	}

	bidId.Status = state
	bidId.FinalBidUserId = bidUserId
	bidId.FinallyBidTime = time.Now().Unix()

	_, err = modelGame.GetAuctionBidModel().DbMap().Update(bidId)
	if err != nil {
		logger.Error("UpdateStatus err:%v auctionInfo.Id:%v, auctionInfo.AuctionType:%v  state:%v", err, auctionId, auctionType, state)
	}

	allInfo, err := modelGame.GetAuctionBidModel().GetUserBidIdByUserIdsAndAuctionId(auctionId, auctionType)
	if err != nil {
		logger.Error("GetUserBidIdByUserIdsAndAuctionId err:%v", err)
		return
	}
	for _, v := range allInfo {
		v.FinallyBidTime = time.Now().Unix()
		v.Status = state
		_, err = modelGame.GetAuctionBidModel().DbMap().Update(v)
		if err != nil {
			logger.Error("UpdateStatus err:%v auctionInfo.Id:%v, auctionInfo.AuctionType:%v  state:%v", err, auctionId, auctionType, state)
		}
	}

}

//获取能够获得分红的玩家
func (this *AuctionManager) GetCanGetRedRewardUses(guildId int, auctionInfo *modelGame.GuildAuctionItem) []int {

	inGuild := make(map[int]bool)
	_, guildUsers, _, _ := this.GetGuild().GetGuildMemberInfo(guildId)
	for _, userIds := range guildUsers {
		for _, userId := range userIds {
			inGuild[userId] = true
		}
	}
	logger.Info("GetCanGetRedRewardUses guildMember:%v   auctionInfo.CanGetRedAward:%v", inGuild, auctionInfo.CanGetRedAward)
	canGetRedUsers := make([]int, 0)
	for _, userId := range auctionInfo.CanGetRedAward {
		if inGuild[userId] {
			canGetRedUsers = append(canGetRedUsers, userId)
		}
	}
	return canGetRedUsers

}
