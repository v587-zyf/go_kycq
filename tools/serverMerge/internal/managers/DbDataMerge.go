package managers

import (
	"cqserver/gamelibs/modelGame"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/tools/serverMerge/internal/base"
	"cqserver/tools/serverMerge/internal/model"
	"database/sql"
	"errors"
	"fmt"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

type MonthCardUnit struct {
	StartTime int
	EndTime   int
}
type MonthCard struct {
	ResetTime  int
	MonthCards map[int]*MonthCardUnit //品质，信息
}

type BaseUserInfo struct {
	UserId               int
	VipLv                int
	IsHaveGetDailyReward int
	MonthCard            *MonthCard
}

//玩家数据合并
type DbDataMerge struct {
	util.DefaultModule
	mergeServerUsers map[int]map[int]int //map[合并的服务器]map[玩家id]玩家原始服务器id
	mergerUserVipLv  map[int]*BaseUserInfo
}

func NewDbDataMerge() *DbDataMerge {
	return &DbDataMerge{}
}

func (this *DbDataMerge) Init() error {
	return nil
}

//开始合并数据
func (this *DbDataMerge) Merge() bool {

	defer func() {
		if r := recover(); r != nil {
			stackBytes := debug.Stack()
			logger.Error("panic DbDataMerge:%v,%s", r, stackBytes)
		}
	}()

	this.mergeServerUsers = make(map[int]map[int]int)
	this.mergerUserVipLv = make(map[int]*BaseUserInfo)
	//获取玩家数据
	for k, _ := range base.Conf.DbConfigs {
		if k == model.NEW_SERVER {
			continue
		}
		k1 := strings.Split(k, "_")
		if len(k1) == 2 && k1[0] == model.DB_SERVER {
			serverId, _ := strconv.Atoi(k1[1])
			logger.Info("合并老服：%v数据开始", k1[1])
			//合并玩家数据
			userIds, err := this.userMerge(k, serverId)
			if err != nil {
				return false
			}
			if len(userIds) == 0 {
				logger.Warn("合并老服：%v数据,没有符合条件的玩家数据", k)
				continue
			}
			this.mergeServerUsers[serverId] = make(map[int]int, 0)
			uidSlice := make([]int, 0)
			for uid, sid := range userIds {
				this.mergeServerUsers[serverId][uid] = sid
				uidSlice = append(uidSlice, uid)
			}
			//武将表
			if ok := this.heroMerge(k, uidSlice); !ok {
				return false
			}

			//挖矿表
			if ok := this.miningMerge(k, uidSlice); !ok {
				return false
			}

			//邮件表
			if ok := this.mailMerge(k, uidSlice); !ok {
				return false
			}

			//门派
			if ok := this.guildMerge(k); !ok {
				return false
			}

			//世界拍卖行
			if ok := this.worldAuctionMerge(k); !ok {
				return false
			}

			//门派拍卖行
			if ok := this.guildAuctionMerge(k); !ok {
				return false
			}

			//抽卡信息
			if ok := this.cardMerge(k); !ok {
				return false
			}

			//寻龙探宝信息
			if ok := this.treasureMerge(k); !ok {
				return false
			}

			//订单信息
			if ok := this.orderMerge(k, uidSlice); !ok {
				return false
			}

			logger.Info("合并老服：%v数据结束", k1[1])
		}
	}
	return true
}

/**合并玩家数据*/
func (this *DbDataMerge) userMerge(dbKey string, serverId int) (map[int]int, error) {
	logger.Info("用户数据表合并开始")

	dataModel := model.GetUserModel()
	var activeTime int64 = 99999999999999
	if base.Conf.ActiveDay > 0 {
		activeTime = time.Now().Unix() - int64(base.Conf.ActiveDay*86400)
	}

	users, err := dataModel.LoadAllUsers(dbKey, base.Conf.ActiveDay, base.Conf.RechargeMin, base.Conf.LevelMin, base.Conf.CombatMin)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("load user err,db:%v,err:%", dbKey, err)
		return nil, errors.New(fmt.Sprintf("合并老服：%v数据拉取玩家数据错误！", dbKey))
	}

	activeUser, rechargeUser, levelUser, combatUser := 0, 0, 0, 0
	userIds := make(map[int]int, 0)
	for _, u := range users {
		if base.Conf.ActiveDay > 0 && u.OfflineTime.Unix() > activeTime {
			activeUser += 1
		}

		if base.Conf.CombatMin > 0 && u.Combat > base.Conf.CombatMin {
			combatUser += 1
		}
		//记录玩家serverId
		if u.ServerId == 0 {
			u.ServerId = serverId
		}
		u = this.competitiveInit(u)
		this.setMergeUserInfo(u)
		u.CompetitiveInfo.BeforeDayRewardGetState = 1
		u = this.limitedDataChange(u)
		u.Recharge = make(map[int]int) //充值双倍奖励
		//武林至尊数据修正数据修正
		userIds[u.Id] = u.ServerId
		err := dataModel.InsertNewData(&u)
		if err != nil {
			logger.Error("合并老服：%v，插入：%v,出现错误err:%v", dbKey, u, err)
			return nil, errors.New(fmt.Sprintf("合并老服：%v，出现错误err:%v", dbKey, err))
		}
		logger.Info("有效玩家详细数据：%v", u)
	}
	logger.Info("合并老服：%v玩家数据，有效玩家数据%v,其中充值大于[%v]的玩家[%v],最近[%v]天登录玩家[%v],等级大于[%v]的玩家[%v],战力大于[%v]的玩家[%v]", dbKey, len(users), base.Conf.RechargeMin, rechargeUser, base.Conf.ActiveDay, activeUser, base.Conf.LevelMin, levelUser, base.Conf.CombatMin, combatUser)
	return userIds, nil
}

/*武将表数据**/
func (this *DbDataMerge) heroMerge(dbKey string, userIds []int) bool {
	logger.Info("武将数据表合并开始")

	dataModel := model.GetHeroModel()
	datas, err := dataModel.GetDatas(dbKey, userIds)
	if err == sql.ErrNoRows {
		logger.Warn("合并老服：%v武将表数数据,没有拉取到武将数数据", dbKey)
	}
	if err != nil && err != sql.ErrNoRows {
		logger.Error("合并老服：%v武将表数据错误：%v", dbKey, err)
		return false
	}
	for _, v := range datas {
		err := dataModel.InsertNewData(&v)
		if err != nil {
			logger.Error("合并老服：%v，插入武将表数据err:%v", dbKey, err)
			return false
		}
		logger.Info("武将数据表详细数据：%v", v)
	}
	logger.Info("武将数据表合并结束,有效数据:%v条", len(datas))
	return true
}

/*挖矿表数据**/
func (this *DbDataMerge) miningMerge(dbKey string, userIds []int) bool {
	logger.Info("挖矿数据表合并开始")

	dataModel := model.GetMiningModel()
	datas, err := dataModel.GetDatas(dbKey, userIds)
	if err == sql.ErrNoRows {
		logger.Warn("合并老服：%v挖矿表数数据,没有拉取到挖矿数数据", dbKey)
	}
	if err != nil && err != sql.ErrNoRows {
		logger.Error("合并老服：%v挖矿表数据错误：%v", dbKey, err)
		return false
	}
	for _, v := range datas {
		err := dataModel.InsertNewData(&v)
		if err != nil {
			logger.Error("合并老服：%v，插入挖矿表数据err:%v", dbKey, err)
			return false
		}
		logger.Info("挖矿数据表详细数据：%v", v)
	}
	logger.Info("挖矿数据表合并结束,有效数据:%v条", len(datas))
	return true
}

/*邮件表数据**/
func (this *DbDataMerge) mailMerge(dbKey string, userIds []int) bool {
	logger.Info("邮件数据表合并开始")

	dataModel := model.GetMailModel()
	datas, err := dataModel.GetDatas(dbKey, userIds)
	if err == sql.ErrNoRows {
		logger.Warn("合并老服：%v邮件表数数据,没有拉取到邮件数数据", dbKey)
	}
	if err != nil && err != sql.ErrNoRows {
		logger.Error("合并老服：%v邮件表数据错误：%v", dbKey, err)
		return false
	}
	for _, v := range datas {
		err := dataModel.InsertNewData(&v)
		if err != nil {
			logger.Error("合并老服：%v，插入邮件表数据err:%v", dbKey, err)
			return false
		}
		logger.Info("邮件数据表详细数据：%v", v)
	}
	logger.Info("邮件数据表合并结束,有效数据:%v条", len(datas))
	return true
}

/*门派表数据**/
func (this *DbDataMerge) guildMerge(dbKey string) bool {
	logger.Info("门派数据表合并开始")

	dataModel := model.GetGuildModel()
	datas, err := dataModel.GetDatas(dbKey)
	if err == sql.ErrNoRows {
		logger.Warn("合并老服：%v门派表数数据,没有拉取到门派数数据", dbKey)
	}
	if err != nil && err != sql.ErrNoRows {
		logger.Error("合并老服：%v门派表数据错误：%v", dbKey, err)
		return false
	}
	for _, v := range datas {
		err := dataModel.InsertNewData(&v)
		if err != nil {
			logger.Error("合并老服：%v，插入门派表数据err:%v", dbKey, err)
			return false
		}
		logger.Info("门派数据表详细数据：%v", v)
	}
	logger.Info("门派数据表合并结束,有效数据:%v条", len(datas))
	return true
}

/*世界拍卖行表数据**/
func (this *DbDataMerge) worldAuctionMerge(dbKey string) bool {
	logger.Info("世界拍卖行数据表合并开始")

	dataModel := model.GetWorldAuctionModel()
	datas, err := dataModel.GetDatas(dbKey)
	if err == sql.ErrNoRows {
		logger.Warn("合并老服：%v世界拍卖行数数据,没有拉取到门派数数据", dbKey)
	}
	if err != nil && err != sql.ErrNoRows {
		logger.Error("合并老服：%v世界拍卖行数据错误：%v", dbKey, err)
		return false
	}
	m.BaseFunctionMerge.SendAuctionItems(datas)
	logger.Info("世界拍卖行数据表合并结束,有效数据:%v条", len(datas))
	return true
}

/*门派拍卖行表数据**/
func (this *DbDataMerge) guildAuctionMerge(dbKey string) bool {
	logger.Info("门派拍卖行数据表合并开始")

	dataModel := model.GetGuildAuctionModel()
	datas, err := dataModel.GetDatas(dbKey)
	if err == sql.ErrNoRows {
		logger.Warn("合并老服：%v门派拍卖行数数据,没有拉取到门派数数据", dbKey)
	}
	if err != nil && err != sql.ErrNoRows {
		logger.Error("合并老服：%v门派拍卖行数据错误：%v", dbKey, err)
		return false
	}
	m.BaseFunctionMerge.SendGuildAuctionItems(datas)
	logger.Info("门派拍卖行数据表合并结束,有效数据:%v条", len(datas))
	return true
}

/*拍卖行竞拍信息表数据**/
func (this *DbDataMerge) auctionBidInfoMerge(dbKey string) bool {
	logger.Info("拍卖行竞拍信息数据表合并开始")

	dataModel := model.GetAuctionBidModel()
	datas, err := dataModel.GetDatas(dbKey)
	if err == sql.ErrNoRows {
		logger.Warn("合并老服：%v拍卖行竞拍信息数据,没有拉取到门派数数据", dbKey)
	}
	if err != nil && err != sql.ErrNoRows {
		logger.Error("合并老服：%v拍卖行竞拍信息数据错误：%v", dbKey, err)
		return false
	}
	for _, v := range datas {
		err := dataModel.InsertNewData(&v)
		if err != nil {
			logger.Error("合并老服：%v，插入拍卖行竞拍信息表数据err:%v", dbKey, err)
			return false
		}
		logger.Info("拍卖行竞拍信息数据表详细数据：%v", v)
	}
	logger.Info("拍卖行竞拍信息数据表合并结束,有效数据:%v条", len(datas))
	return true
}

/*抽卡表数据**/
func (this *DbDataMerge) cardMerge(dbKey string) bool {
	logger.Info("抽卡数据表合并开始")
	dataModel := model.GetCardModel()
	datas, err := dataModel.GetDatas(dbKey)
	if err == sql.ErrNoRows {
		logger.Warn("合并老服：%v抽卡数据,没有拉取到数据", dbKey)
	}
	if err != nil && err != sql.ErrNoRows {
		logger.Error("合并老服：%v抽卡数据错误：%v", dbKey, err)
		return false
	}
	for _, v := range datas {
		err := dataModel.InsertNewData(&v)
		if err != nil {
			logger.Error("合并老服：%v，插入抽卡表数据err:%v", dbKey, err)
			return false
		}
		logger.Info("抽卡数据表详细数据：%v", v)
	}
	logger.Info("抽卡数据表合并结束,有效数据:%v条", len(datas))
	return true
}

/*寻龙探宝表数据**/
func (this *DbDataMerge) treasureMerge(dbKey string) bool {
	logger.Info("寻龙探宝数据表合并开始")
	dataModel := model.GetTreasureModel()
	datas, err := dataModel.GetDatas(dbKey)
	if err == sql.ErrNoRows {
		logger.Warn("合并老服：%v寻龙探宝数据,没有拉取到数据", dbKey)
	}
	if err != nil && err != sql.ErrNoRows {
		logger.Error("合并老服：%v寻龙探宝数据错误：%v", dbKey, err)
		return false
	}
	for _, v := range datas {
		err := dataModel.InsertNewData(&v)
		if err != nil {
			logger.Error("合并老服：%v，插入寻龙探宝表数据err:%v", dbKey, err)
			return false
		}
		logger.Info("寻龙探宝数据表详细数据：%v", v)
	}
	logger.Info("寻龙探宝数据表合并结束,有效数据:%v条", len(datas))
	return true
}

/**合并充值表数据**/
func (this *DbDataMerge) orderMerge(dbKey string, userIds []int) bool {
	logger.Info("充值表数据数据合并开始")

	dataModel := model.GetOrderModel()
	datas, err := dataModel.GetOrderData(dbKey, userIds)
	if err == sql.ErrNoRows {
		logger.Warn("合并老服：%v充值表数数据,没有拉取到充值表数数据", dbKey)
	}
	if err != nil && err != sql.ErrNoRows {
		logger.Error("合并老服：%v充值表数据错误：%v", dbKey, err)
		return false
	}
	for _, v := range datas {
		err := dataModel.InsertNewData(&v)
		if err != nil {
			logger.Error("合并老服：%v，插入充值表数据err:%v", dbKey, err)
			return false
		}
		logger.Info("充值表详细数据：%v", v)
	}
	logger.Info("充值表数据数据合并结束,有效数据:%v条", len(datas))
	return true
}

//竞技场合服数据处理
func (this *DbDataMerge) competitiveInit(u modelGame.User) modelGame.User {

	//u.CompetitiveInfo.BeforeDayRewardGetState = 0
	u.CompetitiveInfo.HaveChallengeTimes = 0
	u.CompetitiveInfo.BuyTimes = 0
	u.CompetitiveInfo.DayResDay = common.GetResetTime(time.Now())
	u.SeasonTimes = 0
	u.SeasonWinTimes = 0
	u.CompetitiveInfo.NowSeason = 1
	return u
}

func (this *DbDataMerge) setMergeUserInfo(u modelGame.User) {

	monthCard := &MonthCard{}
	monthCard.MonthCards = make(map[int]*MonthCardUnit)
	monthCard.ResetTime = u.MonthCard.ResetTime
	for k, data := range u.MonthCard.MonthCards {
		monthCard.MonthCards[k] = &MonthCardUnit{StartTime: data.StartTime, EndTime: data.EndTime}
	}
	u1 := &BaseUserInfo{
		UserId:               u.Id,
		VipLv:                u.VipLevel,
		IsHaveGetDailyReward: u.CompetitiveInfo.BeforeDayRewardGetState,
		MonthCard:            monthCard,
	}

	this.mergerUserVipLv[u.Id] = u1
}

//限时礼包合服数据处理
func (this *DbDataMerge) limitedDataChange(u modelGame.User) modelGame.User {
	userLimited := u.LimitedGift
	nowTime := int(time.Now().Unix())
	mergeData := make(map[int]int)
	for t, lvInfo := range userLimited.List {
		if lv, ok := userLimited.TLv[t]; ok && lvInfo[lv].EndTime >= nowTime && !lvInfo[lv].IsBuy {
			mergeData[t] = lvInfo[lv].Grade
		}
	}
	userLimited.MergeData = mergeData
	return u
}

func (this *DbDataMerge) GetUserInfoByUserId(userId int) *BaseUserInfo {
	return this.mergerUserVipLv[userId]
}
