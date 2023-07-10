package managers

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constAuction"
	"cqserver/gamelibs/publicCon/constMail"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/redisdb"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"cqserver/tools/serverMerge/internal/model"
	"cqserver/tools/serverMerge/internal/rmodel"
	"fmt"
	"math"
	"time"
)

//玩家redis数据合并
type BaseFunctionMerge struct {
	util.DefaultModule
}

func NewBaseFunctionMerge() *BaseFunctionMerge {
	return &BaseFunctionMerge{}
}

func (this *BaseFunctionMerge) SendMail(userId, mailId int, items map[int]int, items1 map[int]int) {

	infos := gamedb.GetModelBagInfo(items, items1)

	mailConf := gamedb.GetMailMailCfg(mailId)
	mail := &modelGame.Mail{
		UserId:     userId,
		MailID:     mailId,
		Sender:     mailConf.FromName,
		Title:      mailConf.Title,
		Content:    mailConf.Content,
		Status:     pb.MAILSTATUS_UNREAD,
		ExpireAt:   time.Now().AddDate(0, 0, mailConf.ExpireDays),
		CreatedAt:  time.Now(),
		Args:       []string{},
		ItemSource: 0,
		Items:      infos,
	}

	if infos != nil && len(infos) > 0 {
		dataModel := model.GetMailModel()
		err := dataModel.InsertNewData(mail)
		if err != nil {
			logger.Error("sendSystemMail to userId:%v  mailId:%v send mail err:", userId, mailId, err)
			return
		}
	}
	return
}

func (this *BaseFunctionMerge) getOpenDayByServerId(serverId int) int {

	serverInfo, err := modelCross.GetServerInfoModel().GetServerInfoByServerId(serverId)
	if err != nil {
		logger.Error("GetServerOpenDaysByServerId DB Error: %v serverId:%v", err, serverId)
		return -1
	}

	return common.GetTheDays(serverInfo.OpenTime)
}

//获取当前赛季 and 赛季第几天
func (this *BaseFunctionMerge) getCurrentSeason(serverId int, needBeforeDay bool) (season, day, openDay int) {
	openDay = this.getOpenDayByServerId(serverId)
	period := gamedb.GetConf().CompetitveSeason
	logger.Info("GetCurrentSeason  openDay:%v  period:%v  needBeforeDay:%v", openDay, period, needBeforeDay)
	if period <= 0 {
		logger.Error("GetLenCompetitiveCompetitiveCfg 配置表错误")
		period = 7
	}

	if needBeforeDay {
		//取前一天的开服天数
		openDay -= 1
	}

	before := openDay / period
	after := openDay % period

	if before <= 0 {
		return 1, after, openDay
	}

	if after == 0 {
		return before, period, openDay
	}

	return before + 1, after, openDay
}

func (this *BaseFunctionMerge) GetMonthCardPrivilege(user *BaseUserInfo, privilege int) int {
	p := 0
	if cfg := gamedb.GetVipLvlCfg(user.VipLv); cfg != nil {
		if num, ok := cfg.Privilege[privilege]; ok {
			p = num
		}
	}
	p += this.GetPrivilege(user, privilege)
	return p
}

func (this *BaseFunctionMerge) GetPrivilege(user *BaseUserInfo, privilege int) int {
	num := 0
	timeNow := int(time.Now().Unix())
	for t, unit := range user.MonthCard.MonthCards {
		if unit.EndTime >= timeNow {
			cfg := gamedb.GetMonthCardByType(t)
			if cfg != nil {
				if _, ok := cfg.Privilege[privilege]; ok {
					num += cfg.Privilege[privilege]
				}
			}
		}
	}
	return num
}

//获取竞技场每日奖励
func (this *BaseFunctionMerge) GetDailyReward(userInfo *BaseUserInfo, redisCs *redisdb.RedisDb, currentSeason, openDay int) map[int]int {
	data := make(map[int]int)
	if userInfo.IsHaveGetDailyReward >= 1 {
		logger.Debug("userId:%v  BeforeDayRewardGetState:%v", userInfo.UserId, userInfo.IsHaveGetDailyReward)
		return data
	}
	key := fmt.Sprintf(rmodel.CompetitiveSeasonRankInfoByOpenDay, openDay-1)
	nowUserScore, _ := redisCs.ZScore(key, userInfo.UserId)
	state, comCfg, _ := gamedb.GetCompetitveCfgByScore(nowUserScore)
	logger.Debug("GetCompetitiveDailyReward user.Id:%v, currentSeason:%v, openDay:%v  nowUserScore:%v  state:%v, comCfg:%v", userInfo.UserId, currentSeason, openDay, nowUserScore, state, comCfg)
	if state {
		privilege := m.BaseFunctionMerge.GetMonthCardPrivilege(userInfo, pb.VIPPRIVILEGE_COMPETITVE_DAILY_REWARD)
		for _, info := range comCfg {
			count := info.Count
			if privilege != 0 {
				count = common.CalcTenThousand(privilege, count)
			}
			data[info.ItemId] = info.Count
		}
	}
	return data
}

//金锭换算成元宝
func (this *BaseFunctionMerge) buildGoldIngotCalc(jinDingNum int) int {

	ingotRate := gamedb.GetConf().YuanBaoRate
	ingotNum := int(math.Ceil(float64(jinDingNum) / float64(ingotRate[1]) * float64(ingotRate[0])))
	logger.Info("金锭换算成元宝 jinDingNum:%v  ingotNum:%v  ingotRate:%v", jinDingNum, ingotNum, ingotRate)
	return ingotNum
}

//发送世界拍卖行 道具
func (this *BaseFunctionMerge) SendAuctionItems(data []modelGame.AuctionItem) {
	for _, v := range data {
		if v.Status > constAuction.OnAuction {
			continue
		}

		items := make(map[int]int)
		if v.NowBidPlayerId <= 0 {
			aInfo := m.DbDataMerge.GetUserInfoByUserId(v.AuctionUserId)
			if aInfo == nil {
				continue
			}
			//没人拍  道具返回给上架玩家
			items[v.ItemId] = v.ItemCount
			logger.Info("auctionId:%v  NowBidPlayerId:%v  items:%v", v.Id, v.AuctionUserId, items)
			this.SendMail(v.AuctionUserId, constMail.MAILTYPE_HEFU_AUCTION_DEAL, items, nil)
			continue
		}

		aInfo := m.DbDataMerge.GetUserInfoByUserId(v.NowBidPlayerId)
		if aInfo != nil {
			//发送给最后竞拍玩家道具
			items = make(map[int]int)
			items[v.ItemId] = v.ItemCount
			logger.Info("auctionId:%v  NowBidPlayerId:%v  items:%v", v.Id, v.NowBidPlayerId, items)
			this.SendMail(v.NowBidPlayerId, constMail.MAILTYPE_HEFU_AUCTION_DEAL4, items, nil)
		}

		//发送给上架玩家元宝
		aInfo = m.DbDataMerge.GetUserInfoByUserId(v.AuctionUserId)
		if aInfo != nil {
			money := v.NowBidPrice - int(math.Ceil(float64(v.NowBidPrice*(gamedb.GetConf().AuctionWorldTax))/10000.0))
			logger.Info("auctionId:%v  AuctionUserId:%v  money:%v", v.Id, v.AuctionUserId, money)
			monthCardPrivilege := m.BaseFunctionMerge.GetMonthCardPrivilege(aInfo, pb.VIPPRIVILEGE_AUCTION_SERVICE_CHARGE)
			if monthCardPrivilege > 0 {
				count := int(math.Ceil(float64(v.NowBidPrice*(gamedb.GetConf().AuctionWorldTax))/10000.0) * (1 - (float64(monthCardPrivilege) / float64(10000))))
				money = v.NowBidPrice - count
			}
			items = make(map[int]int)
			items[pb.ITEMID_INGOT] = money
			this.SendMail(v.AuctionUserId, constMail.MAILTYPE_HEFU_AUCTION_DEAL1, items, nil)
		}
	}
	return
}

//发送公会拍卖行 道具
func (this *BaseFunctionMerge) SendGuildAuctionItems(data []modelGame.GuildAuctionItem) {
	for _, v := range data {
		if v.Status > constAuction.OnAuction {
			continue
		}

		items := make(map[int]int)
		if v.NowBidPlayerId <= 0 {
			//没人拍  流拍
			continue
		}

		aInfo := m.DbDataMerge.GetUserInfoByUserId(v.NowBidPlayerId)
		if aInfo != nil {
			//发送给最后竞拍玩家道具
			items = make(map[int]int)
			items[v.ItemId] = v.ItemCount
			this.SendMail(v.NowBidPlayerId, constMail.MAILTYPE_HEFU_AUCTION_DEAL2, items, nil)
		}

		//分红处理
		money := this.buildGoldIngotCalc(v.NowBidPrice)
		logger.Info("money:%v", money)
		for _, userId := range v.CanGetRedAward {
			aInfo = m.DbDataMerge.GetUserInfoByUserId(userId)
			if aInfo != nil {
				items = make(map[int]int)
				items[pb.ITEMID_INGOT] = money
				this.SendMail(userId, constMail.MAILTYPE_HEFU_AUCTION_DEAL3, items, nil)
			}
		}
	}
	return
}
