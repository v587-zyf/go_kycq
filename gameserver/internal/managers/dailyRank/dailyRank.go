package dailyRank

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constMail"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"fmt"
	"strconv"
	"time"
)

func NewDailyRankManager(module managersI.IModule) *DailyRankManager {
	return &DailyRankManager{IModule: module}
}

type DailyRankManager struct {
	util.DefaultModule
	managersI.IModule
	AddState bool
}

func (this *DailyRankManager) Init() error {
	this.CheckIsAddCombat()
	return nil
}

func (this *DailyRankManager) LoadRankReq(user *objs.User, ack *pb.DailyRankLoadAck) {
	isShowBefore := this.CheckIsReturnBeforeRank(user.ServerId)
	openDay := this.GetSystem().GetServerOpenDaysByServerId(user.ServerId)
	if isShowBefore {
		openDay -= 1
	}
	logger.Debug("DailyRankManager Load userId:%v isShowBefore:%v  openDay:%v  AddState:%v", user.Id, isShowBefore, openDay, this.AddState)
	cfg := gamedb.GetDayRankingDayRankingCfg(openDay)
	if cfg == nil {
		return
	}
	maxNum := gamedb.GetConf().DayRankingshow
	key := rmodel.Rank.GetDailyRankKey(cfg.Type, base.Conf.ServerId)
	ranks := rmodel.Rank.GetDailyRank(key, maxNum)
	logger.Debug("LoadRankReq  maxNum:%v key:%v  ranks:%v", maxNum, key, ranks)

	haveGetIds, data := this.GetMarkRewardInfos(user, cfg.Type)
	ack.HaveGetIds = haveGetIds
	ack.BuyGiftInfos = data
	ack.Ranks, ack.Self, ack.SelfScore = this.ComputingRankings(ranks, cfg.Type, user.Id)
	ack.Type = int32(cfg.Type)

	return
}

func (this *DailyRankManager) CheckIsReturnBeforeRank(serverId int) bool {
	endTime := gamedb.GetConf().DayRankOpenAndEndTime[1]
	dayRankSendRewardTime := gamedb.GetConf().DayRankSendRewardTime
	if endTime.Hour < 10 {
		if time.Now().Hour() >= dayRankSendRewardTime.Hour && time.Now().Minute() >= dayRankSendRewardTime.Minute {
			if time.Now().Hour() <= endTime.Hour && time.Now().Minute() <= endTime.Minute {
				openDay := this.GetSystem().GetServerOpenDaysByServerId(serverId)
				if openDay > 1 {
					openDay = openDay - 1
				}
				cfg := gamedb.GetDayRankingDayRankingCfg(openDay)
				if cfg == nil {
					return false
				}
				return true
			}
		}
	}
	return false
}

// GetMarkReward
// @Description:获取积分奖励
// @receiver this
func (this *DailyRankManager) GetMarkReward(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.DailyRankGetMarkRewardAck) error {
	if id <= 0 {
		return gamedb.ERRPARAM
	}

	rankCfg, err := this.GetDailyRankType(user.ServerId)
	if err != nil {
		return err
	}

	cfg := gamedb.GetDayRankingMarkDayRankingMarkCfg(id)
	if cfg == nil {
		logger.Error("配置不存在 DayRankingMark id:%v", id)
		return gamedb.ERRPARAM
	}

	if cfg.Type != rankCfg.Type {
		logger.Error("cfg.Type:%v  rankCfg.Type:%v", cfg.Type, rankCfg.Type)
		return gamedb.ERRPARAM
	}
	if user.DailyRankInfo[cfg.Type] == nil {
		user.DailyRankInfo[cfg.Type] = &model.DailyRankInfo{GetDayRewardIds: make(model.IntKv), BuyRewardInfo: make(model.IntKv)}
	}

	if user.DailyRankInfo[cfg.Type].GetDayRewardIds[id] == 1 {
		return gamedb.ERRAWARDGET
	}

	key := rmodel.Rank.GetDailyRankKey(cfg.Type, base.Conf.ServerId)
	score := rmodel.Rank.GetDailyRankSelfRankAndScore(key, user.Id)

	infos := gamedb.GetDailyRankMarkCfg(cfg.Type, int(score))
	canAdd := false
	for _, info := range infos {
		if info.Id == id {
			canAdd = true
			break
		}
	}
	if canAdd {
		this.GetBag().AddItems(user, cfg.Reward, op)
	} else {
		logger.Error("userId:%v  score:%v  id:%v", user.Id, score, id)
		return gamedb.ERRCONDITION
	}
	user.DailyRankInfo[cfg.Type].GetDayRewardIds[id] = 1
	user.Dirty = true
	haveGetIds, _ := this.GetMarkRewardInfos(user, cfg.Type)
	ack.Id = int32(id)
	ack.HaveGetIds = haveGetIds
	_ = this.GetUserManager().SendItemChangeNtf(user, op)
	return nil
}

// 购买每日排行礼包
func (this *DailyRankManager) BuyDailyRankGift(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.DailyRankBuyGiftAck) error {

	if id <= 0 {
		return gamedb.ERRPARAM
	}

	cfg, err := this.GetDailyRankType(user.ServerId)
	if err != nil {
		return err
	}

	giftCfg := gamedb.GetDayRankingGiftDayRankingGiftCfg(id)
	if giftCfg.Type1 != cfg.Type {
		logger.Error("giftCfg.Type:%v cfg.Type:%v", giftCfg.Type1, cfg.Type)
		return gamedb.ERRPARAM
	}

	if giftCfg.Type2 == 2 {
		logger.Error("走支付流程  id:%v", id)
		return gamedb.ERRPARAM
	}

	if user.DailyRankInfo[cfg.Type] == nil {
		user.DailyRankInfo[cfg.Type] = &model.DailyRankInfo{GetDayRewardIds: make(model.IntKv), BuyRewardInfo: make(model.IntKv)}
	}

	if user.DailyRankInfo[cfg.Type].BuyRewardInfo[giftCfg.Id] >= giftCfg.Time {
		return gamedb.ERRPURCHASECAPENOUGH
	}

	//num := int(math.Ceil(float64(giftCfg.Consume) * (float64(giftCfg.Discount) / 100)))
	num := giftCfg.Consume
	ok, _ := this.GetBag().HasEnough(user, pb.ITEMID_INGOT, num)
	if !ok {
		return gamedb.ERRNOTENOUGHGOODS
	}

	err = this.GetBag().Remove(user, op, pb.ITEMID_INGOT, num)
	if err != nil {
		return err
	}

	this.GetBag().AddItems(user, giftCfg.Reward, op)
	user.DailyRankInfo[cfg.Type].BuyRewardInfo[giftCfg.Id] += 1
	user.Dirty = true
	_, data := this.GetMarkRewardInfos(user, cfg.Type)
	ack.Id = int32(id)
	ack.BuyGiftInfos = data
	_ = this.GetUserManager().SendItemChangeNtf(user, op)
	kyEvent.UserGiftBuy(user, pb.MONEYPAYTYPE_DAILYRANKBUYGIFT, cfg.Day, pb.ITEMID_INGOT, num, 0)
	return nil
}

/**
 *  @Description: 校验是否开放，支付金额
 *  @param payNum	支付金额
 *  @return error
 */
func (this *DailyRankManager) PayCheck(user *objs.User, payNum, typeId int) error {

	cfg, err := this.GetDailyRankType(user.ServerId)
	if err != nil {
		return err
	}
	buyCfg := gamedb.GetDailyRankBuyGiftCfg(2, payNum, typeId)
	if buyCfg == nil {
		logger.Error("serverId:%v type:%v payNum:%v", user.ServerId, cfg.Type, payNum)
		return gamedb.ERRPARAM
	}

	if user.DailyRankInfo[cfg.Type].BuyRewardInfo == nil {
		user.DailyRankInfo[cfg.Type].BuyRewardInfo = make(map[int]int)
	}

	if user.DailyRankInfo[cfg.Type].BuyRewardInfo[buyCfg.Id] >= buyCfg.Time {
		return gamedb.ERRBUYTIMESLIMIT
	}

	return nil
}

// 充值回调
func (this *DailyRankManager) PayCallBack(user *objs.User, payNum, typeId int, op *ophelper.OpBagHelperDefault) {

	cfg, err := this.GetDailyRankType(user.ServerId)
	if err != nil {
		return
	}

	buyCfg := gamedb.GetDailyRankBuyGiftCfg(2, payNum, typeId)
	if buyCfg == nil {
		logger.Error("serverId:%v type:%v payNum:%v", user.ServerId, cfg.Type, payNum)
		return
	}

	this.GetBag().AddItems(user, buyCfg.Reward, op)

	if user.DailyRankInfo[cfg.Type].BuyRewardInfo[buyCfg.Id] >= buyCfg.Time {
		logger.Error("购买到达上限 userId:%v  user.DailyRankInfo[cfg.Type:%v].BuyRewardInfo[buyCfg.Id:%v]:%v >= buyCfg.Time:%v", user.Id, cfg.Type, buyCfg.Id, user.DailyRankInfo[cfg.Type].BuyRewardInfo[buyCfg.Id], buyCfg.Time)
		return
	}

	user.DailyRankInfo[cfg.Type].BuyRewardInfo[buyCfg.Id] += 1
	user.Dirty = true
	_, data := this.GetMarkRewardInfos(user, cfg.Type)

	ack := &pb.DailyRankBuyGiftAck{}
	ack.Id = int32(buyCfg.Id)
	ack.BuyGiftInfos = data
	user.Dirty = true
	this.GetUserManager().SendMessage(user, ack, false)
	return
}

func (this *DailyRankManager) GetMarkRewardInfos(user *objs.User, types int) ([]int32, map[int32]int32) {
	if user.DailyRankInfo[types] == nil {
		user.DailyRankInfo[types] = &model.DailyRankInfo{GetDayRewardIds: make(model.IntKv), BuyRewardInfo: make(model.IntKv)}
	}

	haveGetIds := make([]int32, 0)
	for ids := range user.DailyRankInfo[types].GetDayRewardIds {
		haveGetIds = append(haveGetIds, int32(ids))
	}

	data := make(map[int32]int32)
	for k, v := range user.DailyRankInfo[types].BuyRewardInfo {
		data[int32(k)] = int32(v)
	}

	return haveGetIds, data
}

func (this *DailyRankManager) GetDailyRankType(serveId int) (*gamedb.DayRankingDayRankingCfg, error) {

	openDay := this.GetSystem().GetServerOpenDaysByServerId(serveId)
	if this.CheckIsReturnBeforeRank(serveId) {
		openDay -= 1
	}
	rankCfg := gamedb.GetDayRankingDayRankingCfg(openDay)
	if rankCfg == nil {
		logger.Error("GetMarkReward  serverId:%v  openDay:%v", serveId, openDay)
		return nil, gamedb.ERRACTIVITYCLOSE
	}
	return rankCfg, nil
}

func (this *DailyRankManager) SendEndMail() error {
	this.AddState = false
	openDay := this.GetSystem().GetServerOpenDaysByServerId(base.Conf.ServerId)
	logger.Info("发送每日排行奖励  openDay:%v", openDay)
	cfg := gamedb.GetDayRankingDayRankingCfg(openDay)
	if cfg == nil {
		logger.Error("GetDayRankingDayRankingCfg  nil  openDay:%v", openDay)
		return nil
	}

	maxNum := gamedb.GetConf().DayRankingshow
	key := rmodel.Rank.GetDailyRankKey(cfg.Type, base.Conf.ServerId)
	ranks := rmodel.Rank.GetDailyRank(key, maxNum)

	rankInfos, _, _ := this.ComputingRankings(ranks, cfg.Type, 0)
	logger.Info("rankInfos:%v", rankInfos)

	rankEventStr := "["
	for _, rankInfo := range rankInfos {
		cfgs := gamedb.GetDayRankRewardCfgByRank(cfg.Type, int(rankInfo.Rank))
		if cfgs == nil {
			logger.Error("GetDayRankRewardCfgByRank err type:%v rank:%v", cfg.Type, rankInfo.Rank)
			continue
		}
		logger.Debug("每日排行发奖 userId:%v rank:%v openDay:%v type:%v", rankInfo.UserInfo.Id, rankInfo.Rank, openDay, cfg.Type)
		args := []string{strconv.Itoa(int(rankInfo.Rank))}
		this.GetMail().SendSystemMailWithItemInfos(int(rankInfo.UserInfo.Id), constMail.MAILTYPR_DAILY_RANK_REWARD, args, cfgs.Reward)

		userInfo := this.GetUserManager().GetUserBasicInfo(int(rankInfo.UserInfo.Id))
		rankEventStr += fmt.Sprintf(`{"userName":"%s","openId":"%s","roleid":"%d","score":"%d","rank":"%d"},`,
			userInfo.NickName, userInfo.OpenId, userInfo.Id, rankInfo.Score, rankInfo.Rank)
	}
	rankEventStr = rankEventStr[:len(rankEventStr)-1]
	rankEventStr += "]"
	kyEvent.RankingList(cfg.Type, cfg.Name, rankEventStr)

	this.SendPointReward(ranks, cfg.Type)
	return nil
}

func (this *DailyRankManager) CheckIsAddCombat() {

	endTime := gamedb.GetConf().DayRankSendRewardTime
	if time.Now().Hour() >= endTime.Hour && time.Now().Minute() >= endTime.Minute {
		this.AddState = false
	} else {
		this.AddState = true
	}
}

func (this *DailyRankManager) ResAddState() {
	this.AddState = true
}

func (this *DailyRankManager) GetAddState() bool {
	return this.AddState
}

func (this *DailyRankManager) OnlineCheck(user *objs.User) {

	allCfg := gamedb.GetAllDayRankingCfg()
	for _, data := range allCfg {
		if user.DailyRankInfo[data.Type] == nil {
			user.DailyRankInfo[data.Type] = &model.DailyRankInfo{GetDayRewardIds: make(model.IntKv), BuyRewardInfo: make(model.IntKv)}
		}
	}
}

func (this *DailyRankManager) SendPointReward(rankInfos []float64, types int) {

	minScore := gamedb.GetDailyRankMarkMinScore(types)
	logger.Info("未领取积分奖励 发放  types:%v  minScore:%v  rankInfos:%v", types, minScore, rankInfos)
	for i, j := 0, len(rankInfos); i < j; i += 2 {
		rankUserId := int(rankInfos[i])
		rankScore := int(rankInfos[i+1])
		if rankScore < minScore {
			continue
		}
		userInfo := this.GetUserManager().GetUser(rankUserId)
		isOffline := false
		if userInfo == nil {
			isOffline = true
			userInfo = this.GetUserManager().GetOfflineUserInfo(rankUserId)
		}
		if userInfo == nil {
			logger.Error("每日排行积分奖励结算  玩家不存在 rankInfo.UserInfo.Id:%v", rankUserId)
			continue
		}

		datas := gamedb.GetDailyRankMarkCfg(types, rankScore)
		if len(userInfo.DailyRankInfo[types].GetDayRewardIds) == len(datas) {
			logger.Info("userId:%v  score:%v userInfo.DailyRankInfo[types].GetDayRewardIds:%v", rankUserId, rankScore, userInfo.DailyRankInfo[types].GetDayRewardIds)
			continue
		}

		if len(datas) <= 0 {
			continue
		}
		allItems := make(gamedb.ItemInfos, 0)
		for _, data := range datas {
			if userInfo.DailyRankInfo[types].GetDayRewardIds[data.Id] == 0 {
				allItems = append(allItems, data.Reward...)
				userInfo.DailyRankInfo[types].GetDayRewardIds[data.Id] = 1
			}
		}
		if isOffline {
			modelGame.GetUserModel().DbMap().Update(userInfo.User)
		}
		userInfo.Dirty = true
		logger.Debug("每日排行发放未领取积分奖励 userId:%v score:%v userInfo.DailyRankInfo[types].GetDayRewardIds:%v", rankUserId, rankScore, userInfo.DailyRankInfo[types].GetDayRewardIds)
		this.GetMail().SendSystemMailWithItemInfos(rankUserId, constMail.MAILTYPR_DAILY_RANK_MARK_REWARD, []string{}, allItems)
	}
	return
}

func (this *DailyRankManager) ComputingRankings(allUsers []float64, rankType, userId int) ([]*pb.RankInfo, int32, int64) {
	rankInfos := make([]*pb.RankInfo, 0)

	myScore := int64(0)
	ownRank := -1
	logger.Debug("ComputingRankings  allUsers:%v  rankType:%v  userId:%v", allUsers, rankType, userId)
	allCfg := gamedb.GetDayRankRewardCfgByTypes(rankType)
	maxNum := gamedb.GetConf().DayRankingshow
	minRankCfg := gamedb.GetDayRankRewardCfgByTypesAndRank(rankType, maxNum)
	if allCfg == nil || len(allCfg) <= 0 || minRankCfg == nil {
		return rankInfos, -1, myScore
	}

	lastPeopleRank := 0
	lastCfgId := 0
	logger.Debug("  allUsers:%v   len:%v", allUsers, len(allUsers))
	for i, j := 0, len(allUsers); i < j; i += 2 {
		rankUserId := int(allUsers[i])
		rankScore := int(allUsers[i+1])

		if this.GetUserManager().BuilderBrieUserInfo(int(rankUserId)) == nil {
			logger.Debug("GetAllUserInfoIncludeOfflineUser nil userId:%v", rankUserId)
			continue
		}

		if minRankCfg.Least > rankScore {
			continue
		}
		logger.Debug("  userId:%v  score:%v  lastPeopleRank:%v  lastCfgId:%v", rankUserId, rankScore, lastPeopleRank, lastCfgId)
		lastPeopleRank, lastCfgId = gamedb.GetDayRankRewardCfgByType(rankType, rankScore, lastPeopleRank, lastCfgId)
		logger.Debug("  userId:%v  score:%v  lastPeopleRank:%v  lastCfgId:%v", rankUserId, rankScore, lastPeopleRank, lastCfgId)
		if rankUserId == userId {
			myScore = int64(rankScore)
			ownRank = lastPeopleRank
		}
		rankInfos = append(rankInfos, this.GetUserManager().BuildUserRankInfo(int(rankUserId), -1, lastPeopleRank, rankScore))

	}

	if len(rankInfos) > 50 {
		return rankInfos[0:50], int32(ownRank), myScore
	}
	if myScore == 0 {
		key := rmodel.Rank.GetDailyRankKey(rankType, base.Conf.ServerId)
		score := rmodel.Rank.GetDailyRankSelfRankAndScore(key, userId)
		myScore = int64(score)
	}

	return rankInfos, int32(ownRank), myScore

}
