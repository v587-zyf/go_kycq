package auction

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constAuction"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"database/sql"
	"math"
	"sync"
	"time"
)

type AuctionManager struct {
	util.DefaultModule
	managersI.IModule

	sync.RWMutex // 拍卖专用锁
	//世界拍卖行MAXid
	worldMaxId int
	//门派拍卖行MAXid
	guildMaxId int
	// 随机数种子
	randSeed int64
	//世界拍卖行
	worldAuctionData map[int][]*modelGame.AuctionItem // key:itemId

	// 门派拍卖
	guildAuctionInfos map[int]map[int]*modelGame.GuildAuctionItem // 门派拍卖信息，main map: key->guildId, sub map: key->auctionId

	//insert world data
	opChan      chan opMsg
	guildOpChan chan guildOpMsg

	AuctionMoney int

	ReturnMoney int
}

func NewAuctionManager(module managersI.IModule) *AuctionManager {
	auctionManager := &AuctionManager{
		worldAuctionData: make(map[int][]*modelGame.AuctionItem),
		// 初始化门派拍卖数据
		guildAuctionInfos: make(map[int]map[int]*modelGame.GuildAuctionItem),

		opChan:      make(chan opMsg, constAuction.DbChanSize),
		guildOpChan: make(chan guildOpMsg, constAuction.DbChanSize),
	}
	auctionManager.IModule = module
	return auctionManager
}

func (this *AuctionManager) ProcessAuctionInfoReq(user *objs.User, auctionType int, ack *pb.AuctionInfoNtf) error {
	if auctionType == constAuction.WorldAuction || auctionType == constAuction.GuildAuction {
		ack = this.ProcessAuctionInfoNtf(auctionType, user.GuildData.NowGuildId, ack)
	} else {
		return gamedb.ERRPARAM
	}
	return nil
}

//兑换金锭  (拍卖行法定交易货币)
func (this *AuctionManager) ConversionGoldIngot(user *objs.User, num int, op *ophelper.OpBagHelperDefault, ack *pb.ConversionGoldIngotAck) error {
	if num <= 0 {
		return gamedb.ERRPARAM
	}
	rechargeRate := gamedb.GetConf().JinDingRate
	ingotRate := gamedb.GetConf().YuanBaoRate
	logger.Debug("rechargeRate:%v  ingotRate:%v", rechargeRate, ingotRate)
	canConversionNum := int(math.Ceil(float64(user.RechargeAll-user.RedPacketNum-user.HaveUseRecharge)/float64(100)/float64(rechargeRate[0])) * float64(rechargeRate[1]))
	logger.Info("兑换金锭  userId:%v  RechargeAll:%v  RedPacketNum:%v  HaveUseRecharge:%v  canConversionNum:%v  num:%v", user.Id, user.RechargeAll, user.RedPacketNum, user.HaveUseRecharge, canConversionNum, num)
	if num > canConversionNum {
		return gamedb.ERRDUIHUANUP
	}

	needRemove := int(math.Ceil(float64(num) / float64(ingotRate[1]) * float64(ingotRate[0])))
	ok, _ := this.GetBag().HasEnough(user, pb.ITEMID_INGOT, needRemove)
	if !ok {
		return gamedb.ERRNOTENOUGHGOODS
	}
	this.GetBag().Remove(user, op, pb.ITEMID_INGOT, needRemove)

	this.GetBag().AddItem(user, op, pb.ITEMID_GOLD_INGOT, num)

	haveUseRecharge := int(math.Ceil(float64(num)/float64(rechargeRate[1]))) * 100

	user.HaveUseRecharge += haveUseRecharge
	ack.HaveUseRecharge = int32(user.HaveUseRecharge)
	return nil
}

//金锭换算成元宝
func (this *AuctionManager) buildGoldIngotCalc(jinDingNum int) int {

	ingotRate := gamedb.GetConf().YuanBaoRate
	ingotNum := int(math.Ceil(float64(jinDingNum) / float64(ingotRate[1]) * float64(ingotRate[0])))
	logger.Info("金锭换算成元宝 jinDingNum:%v  ingotNum:%v  ingotRate:%v", jinDingNum, ingotNum, ingotRate)
	return ingotNum
}

//
//  UpItemToAuction
//  @Description:上架拍卖行
//
func (this *AuctionManager) UpItemToAuction(user *objs.User, req *pb.AuctionPutawayItemReq, ack *pb.AuctionPutawayItemNtf, op *ophelper.OpBagHelperDefault) error {
	logger.Debug("UpItemToAuction run, userId: %v, artifactUuid: %v, itemId: %v  pos:%v  count:%v", user.Id, req.ItemId, req.ItemId, req.Position, req.Count)
	ack.Code = constAuction.AuctionItemOk
	if req.Count <= 0 || req.ItemId <= 0 || req.Price <= 0 {
		return gamedb.ERRPARAM
	}

	itemId := int(req.ItemId)
	itemCfg := gamedb.GetAuctionAuctioinCfg(itemId)
	itemBaseCfg := gamedb.GetItemBaseCfg(itemId)
	if itemCfg == nil || itemBaseCfg == nil {
		ack.Code = constAuction.AuctionItemErrConf
		return nil
	}

	if int(req.Count) > itemCfg.MaxNum {
		ack.Code = constAuction.AuctionOverCount
		return nil
	}

	nowPrice := int(req.Price)
	if nowPrice < itemCfg.LowPrice*int(req.Count) || nowPrice > itemCfg.HighPrice*int(req.Count) {
		ack.Code = constAuction.AuctionSetPriceErr
		return nil
	}
	var err error
	//每小时上架物品限制
	limitTimes := gamedb.GetConf().AuctionNum
	alreadyAuction := rmodel.Auction.RangeAuctionTogether(user.Id)
	logger.Info("numberLimit:%d  together:%d  count:%v", limitTimes, len(alreadyAuction), req.Count)
	if len(alreadyAuction)+1 > limitTimes {
		ack.Code = constAuction.AuctionTogetherLimit
		return nil
	}

	enough, _ := this.GetBag().HasEnough(user, itemId, int(req.Count))
	if !enough {
		ack.Code = constAuction.AuctionItemErrNoItem
	}

	if itemBaseCfg.Type == pb.ITEMTYPE_EQUIP {

		if req.Count == 1 {
			err = this.GetBag().RemoveByPosition(user, op, itemId, int(req.Count), int(req.Position))
			if err != nil {
				return err
			}
		}
		if req.Count >= 2 {
			for count := 0; count < int(req.Count); count++ {
				pos, _, _ := this.GetBag().GetEquipItemInfos(user, itemId)
				err = this.GetBag().RemoveByPosition(user, op, itemId, 1, pos)
				if err != nil {
					return err
				}
			}
		}
	} else {
		err = this.GetBag().Remove(user, op, itemId, int(req.Count))
		if err != nil {
			return err
		}
	}
	rmodel.Auction.AddAuctionTogether(user.Id, itemId)

	auctionItemInfo := &modelGame.AuctionItem{}
	auctionItemInfo.Id = this.GetIdGenerator().GetNextWorldId()
	auctionItemInfo.PutAwayPrice = nowPrice
	auctionItemInfo.AuctionUserId = user.Id
	auctionItemInfo.AuctionSrc = constAuction.AuctionSrcPlayer
	auctionItemInfo.ItemId = itemId
	auctionItemInfo.ItemCount = int(req.Count)
	auctionItemInfo.AuctionTime = int(time.Now().Unix())
	auctionItemInfo.AuctionDuration = itemCfg.Time * 60
	auctionItemInfo.Status = constAuction.OnAuction
	auctionItemInfo.AuctionType = constAuction.WorldAuction

	ack.AuctionInfos = this.buildAuctionInfo(auctionItemInfo)
	//存入数据库
	this.addItemIntoAuction(auctionItemInfo)

	kyEvent.AuctionUp(user.Id, nowPrice, itemId, int(req.Count), constAuction.WorldAuction)

	return nil
}

//
//  ProcessBidInfo
//  @Description:指定物品竞拍信息
//
func (this *AuctionManager) ProcessBidInfo(user *objs.User, auctionId, auctionType int, ack *pb.BidInfoNtf) error {
	if gamedb.GetItemBaseCfg(auctionId) == nil {
		return gamedb.ERRPARAM
	}
	if auctionType == constAuction.WorldAuction || auctionType == constAuction.GuildAuction {
		ack.AuctionInfo, _, _ = this.GetAuctionInfoByAuctionId(auctionType, auctionId)
		ack.AuctionType = int32(auctionType)
	} else {
		return gamedb.ERRPARAM
	}
	return nil
}

// 处理玩家竞价
func (this *AuctionManager) ProcessBid(user *objs.User, req *pb.BidReq, ack *pb.BidNtf, op *ophelper.OpBagHelperDefault) error {
	if req.AuctionType < 0 || req.AuctionId < 0 || req.IsBuyNow < 0 {
		return gamedb.ERRPARAM
	}
	logger.Debug("ProcessBid  AuctionType:%v", req.AuctionType)
	if req.AuctionType == constAuction.WorldAuction {
		this.processWorldGsBidNtf(user, req, ack, op)
	} else if req.AuctionType == constAuction.GuildAuction {
		this.processGuidAuctionGsBidNtf(user, req, ack, op)
	} else {
		return gamedb.ERRPARAM
	}
	ack.AuctionId = int32(req.AuctionId)
	ack.AuctionType = req.AuctionType
	return nil
}

//  MyBidNtf
//  @Description: 玩家竞拍物品信息
func (this *AuctionManager) MyBidNtf(user *objs.User, ack *pb.MyBidNtf) {
	this.RLock()
	defer this.RUnlock()
	for _, data := range this.worldAuctionData {
		for _, info := range data {
			if info.NowBidPlayerId == user.Id {
				ack.MyBidInfos = append(ack.MyBidInfos, this.buildAuctionInfo(info))
			}
		}
	}

	for _, data := range this.guildAuctionInfos[user.GuildData.NowGuildId] {
		if data.NowBidPlayerId == user.Id {
			ack.MyBidInfos = append(ack.MyBidInfos, this.buildGuildAuctionInfo(data))
		}
	}
	return
}

//  MyBidNtf
//  @Description: 玩家上架的物品信息
func (this *AuctionManager) MyPutAwayItemInfo(user *objs.User, ack *pb.MyPutAwayItemInfoAck) {

	infos, err := modelGame.GetAuctionItemModel().GetAuctionItemByAuctionUser(user.Id, 30)
	if err != nil {
		logger.Error("GetAuctionItemByAuctionUser err:%v  userId:%v", err, user.Id)
		return
	}
	for _, data := range infos {
		ack.MyBidInfosNotInBid = append(ack.MyBidInfosNotInBid, this.buildAuctionInfo(data))
	}

	infos, err = modelGame.GetAuctionItemModel().GetAuctionItemByAuctionUserInBid(user.Id)
	if err != nil {
		logger.Error("GetAuctionItemByAuctionUser err:%v  userId:%v", err, user.Id)
		return
	}
	for _, data := range infos {
		ack.MyBidInfosInBid = append(ack.MyBidInfosInBid, this.buildAuctionInfo(data))
	}

	return
}

func (this *AuctionManager) GetUserBidInfos(user *objs.User, ack *pb.MyBidInfoItemAck) {
	bidInfos, err := modelGame.GetAuctionBidModel().GetAllUserBidItemsByUserId(user.Id)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("GetAllUserBidItemsByUserId err: %v, userId: %d", err, user.Id)
		return
	}
	for _, info := range bidInfos {
		ack.MyBidInfos = append(ack.MyBidInfos, this.buildBidInfo(info))
	}
	return
}

//拍卖行红点处理
func (this *AuctionManager) BroadcastAuctionRedPointNtf(guildId, types, isBright int) {
	if guildId <= 0 {
		logger.Error("玩家没有门派")
		return
	}
	guildInfo := this.GetGuild().GetGuildInfo(guildId)
	if guildInfo == nil {
		logger.Error("门派不存在 guildId:%v", guildId)
		return
	}
	msg := &pb.RedPointStateNtf{Type: int32(types), IsBright: int32(isBright)}
	logger.Info("guildId:%v  guildInfo.Positions:%v", guildId, guildInfo.Positions)
	for i, j := 0, len(guildInfo.Positions); i < j; i += 2 {
		userId := guildInfo.Positions[i]
		guildUser := this.GetUserManager().GetUser(userId)
		if guildUser != nil {
			this.GetUserManager().SendMessage(guildUser, msg, true)
		}
	}
	return
}

//拍卖行红点处理
func (this *AuctionManager) SendAuctionRedPointNtf(userId, types, isBright int) {
	userInfo := this.GetUserManager().GetUser(userId)
	if userInfo != nil {
		msg := &pb.RedPointStateNtf{Type: int32(types), IsBright: int32(isBright)}
		this.GetUserManager().SendMessage(userInfo, msg, true)
	}
	return
}
