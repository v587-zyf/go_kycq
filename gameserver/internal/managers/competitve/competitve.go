package competitve

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"sync"
	"time"
)

func NewCompetitveManager(m managersI.IModule) *CompetitveManager {
	competive := &CompetitveManager{
		WinState: make(map[int]int),
	}
	competive.IModule = m
	return competive
}

type CompetitveManager struct {
	util.DefaultModule
	managersI.IModule
	WinState map[int]int
	Mu       sync.RWMutex
}

func (this *CompetitveManager) Init() error {
	//清除redis 中老数据
	this.CheckUser()
	return nil
}

func (this *CompetitveManager) setWinState(userId, state int) {
	defer this.Mu.Unlock()
	this.Mu.Lock()
	this.WinState[userId] = state
}

func (this *CompetitveManager) getWinState(userId int) int {
	defer this.Mu.RUnlock()
	this.Mu.RLock()
	return this.WinState[userId]
}

//
//  LoadInfo
//  @Description:
//  @receiver this
//
func (this *CompetitveManager) LoadInfo(user *objs.User, ack *pb.CompetitveLoadAck) error {

	competitiveTimesCfg, _, openDay, err := this.CheckGameCfg(user)
	if err != nil {
		return err
	}
	//今天剩余可挑战次数
	haveChallengeTimes := user.CompetitiveInfo.HaveChallengeTimes
	vipAddLv := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_COMPETITVE_FREENUM)
	lastChallengeTimes := this.GetChallengeTimes(competitiveTimesCfg[0]+vipAddLv, haveChallengeTimes)
	//今天已购买次数
	haveBuyTimes := user.CompetitiveInfo.BuyTimes
	vipBuyAddLv := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_COMPETITVE_BUYNUM)
	lastCanBuyTimes := this.GetChallengeTimes(competitiveTimesCfg[1]+vipBuyAddLv, haveBuyTimes)

	currentSeason, day, _ := this.GetCurrentSeason(user.ServerId, false)        //当前赛季
	seasonRankInfos := rmodel.Competitve.GetSeasonRankInfos(currentSeason, 100) //赛季排名
	_, nowUserScore := rmodel.Competitve.GetSeasonSelfRankAndScore(user.Id, currentSeason)
	ack.UserScore = int32(nowUserScore)
	ack.SeasonTimes = int32(user.SeasonTimes)            //赛季场次
	ack.SessionWinTimes = int32(user.SeasonWinTimes)     //赛季胜场
	ack.RemainChallengeTimes = int32(lastChallengeTimes) // 剩余可挑战次数
	ack.TodayCanBuyTimes = int32(lastCanBuyTimes)        //剩余可购买次数
	if openDay > 1 {
		lastDayUserScore := rmodel.Competitve.GetCompetitiveScoreByOpenDay(user.Id, openDay-1)
		state, _, cfg := gamedb.GetCompetitveCfgByScore(lastDayUserScore)
		if state {
			ack.YestardayReward = int32(cfg.Id) //取出来是配置表对应id 昨天的奖励
		}
	}
	ack.SeasonRank = this.GetSeasonRankInfo(seasonRankInfos) //当前赛季排名
	if currentSeason > 1 {
		lastSeasonRankInfos := rmodel.Competitve.GetSeasonRankInfos(currentSeason-1, 100) //上一赛季
		rank, score := rmodel.Competitve.GetSeasonSelfRankAndScore(user.Id, currentSeason-1)
		ack.LastSeasonRank = this.GetSeasonRankInfo(lastSeasonRankInfos) //上一赛季排名
		ack.LastSeasonUserRank = int32(rank)
		ack.LastSeasonUserRankScore = int32(score)
	}
	if day > 1 {
		ack.BeginTimes = int32(common.GetZeroTimeUnix(-(day - 1))) //赛季开始时间戳
	} else {
		ack.BeginTimes = int32(common.GetZeroTimeUnix(0))
	}
	logger.Debug("openDay:%v   user.CompetitiveInfo.NowSeason:%v currentSeason:%v", openDay, user.CompetitiveInfo.NowSeason, currentSeason)
	//user.CompetitiveInfo.NowSeason = currentSeason
	return nil
}

//购买挑战次数
func (this *CompetitveManager) BuyCompetitiveChallengeNum(user *objs.User, op *ophelper.OpBagHelperDefault, ack *pb.BuyCompetitveChallengeTimesAck) error {

	competitiveTimesCfg, CompetitiveCostCfg, _, err := this.CheckGameCfg(user)
	if err != nil {
		return err
	}
	//今天已购买次数
	if user.CompetitiveInfo.BuyTimes >= competitiveTimesCfg[1]+this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_COMPETITVE_BUYNUM) {
		logger.Error("购买次数到达上线")
		return gamedb.ERRBUYUPPERLIMIT
	}

	//今天已挑战次数
	if user.CompetitiveInfo.HaveChallengeTimes <= 0 {
		logger.Error("BuyCompetitiveChallengeNum userId:%v  挑战次数已满", user.Id)
		return gamedb.ERRENOUGHTIMES
	}

	itemId := 0
	count := 0
	for _, v := range CompetitiveCostCfg {
		ok, _ := this.GetBag().HasEnough(user, v.ItemId, v.Count)
		if ok {
			itemId = v.ItemId
			count = v.Count
			break
		}
	}

	if itemId == 0 || count == 0 {
		return gamedb.ERRNOTENOUGHGOODS
	}

	err = this.GetBag().Remove(user, op, itemId, count)
	if err != nil {
		return err
	}

	user.CompetitiveInfo.HaveChallengeTimes -= 1
	user.CompetitiveInfo.BuyTimes += 1
	vipAdd := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_COMPETITVE_FREENUM)
	ack.ResidueTimes = int32(this.GetChallengeTimes(competitiveTimesCfg[0]+vipAdd, user.CompetitiveInfo.HaveChallengeTimes)) //剩余挑战次数
	vipBuyAddLv := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_COMPETITVE_BUYNUM)
	ack.TodayCanBuyTimes = int32(this.GetChallengeTimes(competitiveTimesCfg[1]+vipBuyAddLv, user.CompetitiveInfo.BuyTimes)) //今日可购买次数
	ack.Goods = op.ToChangeItems()
	return nil
}

//领取每日奖励
func (this *CompetitveManager) GetCompetitiveDailyReward(user *objs.User, op *ophelper.OpBagHelperDefault, ack *pb.GetCompetitveDailyRewardAck) error {

	if user.CompetitiveInfo.BeforeDayRewardGetState >= 1 {
		return gamedb.ERRAWARDGET

	}
	openDay := this.GetCompetitiveOpenDay(user.ServerId)
	if openDay == 1 {
		user.CompetitiveInfo.BeforeDayRewardGetState = 1
		ack.HaveGetRewardState = 1
		return nil
	}
	currentSeason, _, _ := this.GetCurrentSeason(user.ServerId, true)

	nowUserScore := rmodel.Competitve.GetCompetitiveScoreByOpenDay(user.Id, openDay-1)

	state, comCfg, _ := gamedb.GetCompetitveCfgByScore(nowUserScore)
	logger.Debug("GetCompetitiveDailyReward user.Id:%v, currentSeason:%v, openDay:%v  nowUserScore:%v  state:%v, comCfg:%v", user.Id, currentSeason, openDay, nowUserScore, state, comCfg)
	if state {
		privilege := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_COMPETITVE_DAILY_REWARD)
		for _, info := range comCfg {
			count := info.Count
			if privilege != 0 {
				count = common.CalcTenThousand(privilege, count)
			}
			_ = this.GetBag().Add(user, op, info.ItemId, count)
		}
	}
	user.CompetitiveInfo.BeforeDayRewardGetState = 1
	user.Dirty = true
	ack.HaveGetRewardState = int32(user.CompetitiveInfo.BeforeDayRewardGetState)
	ack.Goods = op.ToChangeItems()
	return nil
}

//  RefCompetitiveRival
//  @Description:
//  @section:当前分数段区间  match: 当前段位配置表id//
func (this *CompetitveManager) RefCompetitiveRival(user *objs.User, ack *pb.RefCompetitveRankAck) error {

	competitiveTimesCfg := gamedb.GetConf().CompetitveTimes
	//今天已挑战次数
	haveChallengeTimes := user.CompetitiveInfo.HaveChallengeTimes
	if haveChallengeTimes >= competitiveTimesCfg[0]+this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_COMPETITVE_FREENUM) {
		logger.Error("RefCompetitiveRival userId:%v  挑战次数已满   haveChallengeTimes:%v  vipAdd:%v", user.Id, haveChallengeTimes, this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_COMPETITVE_FREENUM))
		return gamedb.ERRENOUGHTIMES
	}
	curSeason, _, _ := this.GetCurrentSeason(user.ServerId, false)
	_, beforeNowScore := rmodel.Competitve.GetSeasonSelfRankAndScore(user.Id, curSeason)
	nowScore := int(beforeNowScore)
	section, matchId := gamedb.GetCompetitveMatchRuleSectionCfg(nowScore)
	logger.Debug("竞技场匹配对手 userId:%v 匹配分数区间section:%v  配置表Id_matchId:%v  nowScore:%v  curSeason:%v", user.Id, section, matchId, nowScore, curSeason)
	competitiveCfgLen, _, _, err := this.CheckCfg()
	if err != nil {
		return err
	}
	if len(section) >= 2 {
		userIds, userIdInfos, _ := this.GetCompetitiveNowSeasonScoreSection(curSeason, section[0], section[1], user.Id)
		logger.Debug("curSeason:%v, section[0]:%v, section[1]:%v  userIds:%v  userIdInfos:%v", curSeason, section[0], section[1], userIds, userIdInfos)
		if len(userIds) > 0 {
			return this.buildCompetitiveInfo(user.Id, userIds, userIdInfos, ack)
		} else {
			for i := matchId; i > 0; i-- {
				//降低段位寻找
				if i <= 1 {
					logger.Info("没法向下匹配 只能向上匹配")
					continue
				} else {
					beforeCfg := gamedb.GetCompetitveCompetitveCfg(i - 1)
					afterCfg := gamedb.GetCompetitveCompetitveCfg(i)
					userIds, userIdInfos, _ := this.GetCompetitiveNowSeasonScoreSection(curSeason, beforeCfg.Mark, afterCfg.Mark, user.Id)
					logger.Debug("竞技场没有匹配到当前自己分数段的人  userId:%v  真实段位:%v  reduceMatch:%v  match:%v 去寻找  curSeason:%v, beforeCfg.Mark:%v, afterCfg.Mark:%v  userIds:%v", user.Id, matchId, i-1, i, curSeason, beforeCfg.Mark, afterCfg.Mark, userIds)
					if len(userIds) > 0 {
						return this.buildCompetitiveInfo(user.Id, userIds, userIdInfos, ack)
					}
				}
			}
			//降低段位没找到  那么升高段位去找
			for i := matchId; i < competitiveCfgLen; i++ {
				//降低段位寻找
				beforeCfg := gamedb.GetCompetitveCompetitveCfg(i)
				afterCfg := gamedb.GetCompetitveCompetitveCfg(i + 1)
				userIds, userIdInfos, _ := this.GetCompetitiveNowSeasonScoreSection(curSeason, beforeCfg.Mark, afterCfg.Mark, user.Id)
				logger.Debug("竞技场没有匹配到当前自己分数段的人降低段位也没找到那么升高段位去找  userId:%v  真实段位:%v  match:%v  addMatch:%v 去寻找 competitiveCfgsLen:%v   curSeason:%v, beforeCfg.Mark:%v, afterCfg.Mark:%v  userIds:%v", user.Id, matchId, i, i+1, competitiveCfgLen, curSeason, beforeCfg.Mark, afterCfg.Mark, userIds)
				if len(userIds) > 0 {
					return this.buildCompetitiveInfo(user.Id, userIds, userIdInfos, ack)
				}
			}
			logger.Debug("竞技场向上向下都没找到人userId:%v   匹配假人", user.Id)
			//向上向下都没匹配到  去匹配假人
			robotInfo := gamedb.GetCompetitiveRobotUserByScore(nowScore)
			if robotInfo != nil {
				logger.Info("竞技场 添加假人成功  score:%v  robotId:%v", nowScore, robotInfo.Id)
				ack.UserInfo, ack.Score, _ = builder.BuildRobotUserInfo(robotInfo.Id, base.Conf.ServerId)
				return nil
			}
		}
	}
	return nil
}

//  RefCompetitiveRival
//  @Description:
//  @section:竞技场新的匹配规则
func (this *CompetitveManager) RefCompetitiveRivalNew(user *objs.User, ack *pb.RefCompetitveRankAck) error {
	curSeason, _, _ := this.GetCurrentSeason(user.ServerId, false)
	_, beforeNowScore := rmodel.Competitve.GetSeasonSelfRankAndScore(user.Id, curSeason)
	nowScore := int(beforeNowScore)
	section, matchId := gamedb.GetCompetitveMatchRuleSectionCfg(nowScore)
	//1. 玩家首次匹配竞技场，必定匹配到机器人
	lastMarkUserId := rmodel.Competitve.GetLastMarkUserId(user.Id)
	if user.SeasonTimes <= 0 || user.SeasonLoseContinueTimes >= gamedb.GetConf().Elo {
		robotInfo := gamedb.GetCompetitiveRobotUserByScoreSlice(section[0], section[1])
		if robotInfo != nil {
			logger.Info("竞技场 添加假人成功  score:%v  robotId:%v", nowScore, robotInfo.Id)
			rmodel.Competitve.SetLastMarkUserId(user.Id, -robotInfo.Id)
			ack.UserInfo, ack.Score, _ = builder.BuildRobotUserInfo(robotInfo.Id, base.Conf.ServerId)
			if user.SeasonLoseContinueTimes >= gamedb.GetConf().Elo {
				user.SeasonLoseContinueTimes = 0
			}
			return nil
		}
	}

	//2.玩家段位 <= xx 时 一定匹配到机器人，反之优先匹配真人

	if nowScore <= gamedb.GetConf().CompetitveLevel[0] {
		//匹配假人
		robotInfo := gamedb.GetCompetitiveRobotUserByScoreSlice(section[0], section[1])
		if robotInfo != nil {
			logger.Info("竞技场 添加假人成功  score:%v  robotId:%v", nowScore, robotInfo.Id)
			rmodel.Competitve.SetLastMarkUserId(user.Id, -robotInfo.Id)
			ack.UserInfo, ack.Score, _ = builder.BuildRobotUserInfo(robotInfo.Id, base.Conf.ServerId)
			return nil
		}
	}

	//3.匹配真人

	competitiveCfgLen, _, _, err := this.CheckCfg()
	if err != nil {
		return err
	}
	if len(section) >= 2 {
		afterUsers := make([]int, 0)
		userIds, userIdInfos, _ := this.GetCompetitiveNowSeasonScoreSection(curSeason, section[0], section[1], user.Id)
		logger.Debug("curSeason:%v, section[0]:%v, section[1]:%v  userIds:%v  userIdInfos:%v  lastMarkUserId:%v", curSeason, section[0], section[1], userIds, userIdInfos, lastMarkUserId)
		for _, userId := range userIds {
			if userId != lastMarkUserId {
				afterUsers = append(afterUsers, userId)
			}
		}
		logger.Debug("afterUsers:%v", afterUsers)
		userIds = afterUsers
		if len(userIds) > 0 {
			return this.buildCompetitiveInfo(user.Id, userIds, userIdInfos, ack)
		} else {
			for i := matchId; i > 0; i-- {
				//降低段位寻找
				if i <= 1 {
					logger.Info("没法向下匹配 只能向上匹配")
					continue
				} else {
					beforeCfg := gamedb.GetCompetitveCompetitveCfg(i - 1)
					afterCfg := gamedb.GetCompetitveCompetitveCfg(i)
					userIds, userIdInfos, _ := this.GetCompetitiveNowSeasonScoreSection(curSeason, beforeCfg.Mark, afterCfg.Mark, user.Id)
					logger.Debug("竞技场没有匹配到当前自己分数段的人  userId:%v  真实段位:%v  reduceMatch:%v  match:%v 去寻找  curSeason:%v, beforeCfg.Mark:%v, afterCfg.Mark:%v  userIds:%v", user.Id, matchId, i-1, i, curSeason, beforeCfg.Mark, afterCfg.Mark, userIds)
					if len(userIds) > 0 {
						return this.buildCompetitiveInfo(user.Id, userIds, userIdInfos, ack)
					}
				}
			}
			//降低段位没找到  那么升高段位去找
			for i := matchId; i < competitiveCfgLen; i++ {
				//降低段位寻找
				beforeCfg := gamedb.GetCompetitveCompetitveCfg(i)
				afterCfg := gamedb.GetCompetitveCompetitveCfg(i + 1)
				userIds, userIdInfos, _ := this.GetCompetitiveNowSeasonScoreSection(curSeason, beforeCfg.Mark, afterCfg.Mark, user.Id)
				logger.Debug("竞技场没有匹配到当前自己分数段的人降低段位也没找到那么升高段位去找  userId:%v  真实段位:%v  match:%v  addMatch:%v 去寻找 competitiveCfgsLen:%v   curSeason:%v, beforeCfg.Mark:%v, afterCfg.Mark:%v  userIds:%v", user.Id, matchId, i, i+1, competitiveCfgLen, curSeason, beforeCfg.Mark, afterCfg.Mark, userIds)
				if len(userIds) > 0 {
					return this.buildCompetitiveInfo(user.Id, userIds, userIdInfos, ack)
				}
			}
			logger.Debug("竞技场向上向下都没找到人userId:%v   匹配假人", user.Id)
			//向上向下都没匹配到  去匹配假人
			robotInfo := gamedb.GetCompetitiveRobotUserByScoreSlice(section[0], section[1])
			if robotInfo != nil {
				logger.Info("竞技场 添加假人成功  score:%v  robotId:%v", nowScore, robotInfo.Id)
				rmodel.Competitve.SetLastMarkUserId(user.Id, -robotInfo.Id)
				ack.UserInfo, ack.Score, _ = builder.BuildRobotUserInfo(robotInfo.Id, base.Conf.ServerId)
				return nil
			}
		}
	}

	return nil
}

func (this *CompetitveManager) GetCompetitiveNowSeasonScoreSection(curSeason, min, max, myUserId int) ([]int, map[int]int, error) {
	rankInfos, err := rmodel.Competitve.GetSeasonRankInfosBuyScoreSection(curSeason, min, max)
	userIds := make([]int, 0)
	userIdInfos := make(map[int]int)
	for i, j := 0, len(rankInfos); i < j; i += 2 {
		if rankInfos[i] > 0 {
			userId := int(rankInfos[i])
			if userId == myUserId {
				//过滤掉自己
				continue
			}
			userIds = append(userIds, userId)
			userIdInfos[userId] = int(rankInfos[i+1])
		}
	}

	return userIds, userIdInfos, err
}

//获取当前赛季 and 赛季第几天
func (this *CompetitveManager) GetCurrentSeason(serverId int, needBeforeDay bool) (season, day, openDay int) {
	openDay = this.GetCompetitiveOpenDay(serverId)
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

func (this *CompetitveManager) GetChallengeTimes(cfgTimes, times int) int {

	lastTimes := cfgTimes - times
	if lastTimes <= 0 {
		lastTimes = 0
	}
	return lastTimes
}

func (this *CompetitveManager) CheckGameCfg(user *objs.User) (competitveTimesCfg gamedb.IntSlice, CompetitiveCostCfg gamedb.ItemInfos, openDay int, err error) {

	openDay = this.GetCompetitiveOpenDay(user.ServerId)
	competitveTimesCfg = gamedb.GetConf().CompetitveTimes
	if len(competitveTimesCfg) < 2 {
		return nil, nil, openDay, gamedb.ERRGAMECFGERR
	}
	CompetitiveCostCfg = gamedb.GetConf().CompetitveCost
	if len(CompetitiveCostCfg) < 2 {
		return nil, nil, openDay, gamedb.ERRGAMECFGERR
	}
	return competitveTimesCfg, CompetitiveCostCfg, openDay, nil
}

//判断赛季是否结束
func (this *CompetitveManager) JudgeSeasonIsOver() bool {
	season, day, _ := this.GetCurrentSeason(base.Conf.ServerId, false)
	if season <= 0 || day != 1 {
		// ！= 1 因为是第8天的5点执行
		logger.Debug(" SendSeasonEndReward  赛季未结束 season:%v, day:%v", season, day)
		return false
	}
	return true
}

func (this *CompetitveManager) CheckCfg() (int, gamedb.IntSlice, gamedb.IntSlice, error) {

	competitiveCfgsLen := gamedb.GetLenCompetitveCompetitveCfg()
	levelCfg := gamedb.GetConf().CompetitveLevel
	rankCfg := gamedb.GetConf().CompetitveRank
	if levelCfg == nil || rankCfg == nil || len(levelCfg) < 1 || len(rankCfg) < 2 {
		logger.Error("RefCompetitiveRival  竞技场匹配对手 配置err levelCfg:%v  rankCfg:%v", levelCfg, rankCfg)
		return 0, nil, nil, gamedb.ERRENOUGHTIMES
	}
	return competitiveCfgsLen, levelCfg, rankCfg, nil
}

//背包使用道具购买挑战次数
func (this *CompetitveManager) BagUserItemAddChallengeNum(user *objs.User) error {

	competitiveTimesCfg, _, _, err := this.CheckGameCfg(user)
	if err != nil {
		return err
	}
	ack := &pb.BuyCompetitveChallengeTimesAck{}
	user.CompetitiveInfo.BuyTimes += 1
	user.CompetitiveInfo.HaveChallengeTimes -= 1
	vipAdd := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_COMPETITVE_FREENUM)
	ack.ResidueTimes = int32(this.GetChallengeTimes(competitiveTimesCfg[0]+vipAdd, user.CompetitiveInfo.HaveChallengeTimes)) //剩余挑战次数
	ack.TodayCanBuyTimes = int32(this.GetChallengeTimes(competitiveTimesCfg[1], user.CompetitiveInfo.BuyTimes))              //今日可购买次数
	return nil
}

func (this *CompetitveManager) OnLine(user *objs.User) {
	this.DayReset(user, false)
}

//每日重置玩家数据
func (this *CompetitveManager) DayReset(user *objs.User, isReset bool) {
	data := common.GetResetTime(time.Now())
	currentSeason, _, _ := this.GetCurrentSeason(user.ServerId, false)
	if data == user.CompetitiveInfo.DayResDay && currentSeason == user.CompetitiveInfo.NowSeason {
		return
	}
	logger.Debug("DayReset userId:%v  data:%v  currentSeason:%v  user.CompetitiveInfo.DayResDay:%v   user.CompetitiveInfo.NowSeason:%v", user.Id, data, currentSeason, user.CompetitiveInfo.DayResDay, user.CompetitiveInfo.NowSeason)
	if data != user.CompetitiveInfo.DayResDay {
		user.CompetitiveInfo.BeforeDayRewardGetState = 0
		user.CompetitiveInfo.HaveChallengeTimes = 0
		user.CompetitiveInfo.BuyTimes = 0
		user.CompetitiveInfo.DayResDay = data
	}
	if currentSeason != user.CompetitiveInfo.NowSeason {
		user.SeasonTimes = 0
		user.SeasonWinTimes = 0
		user.CompetitiveInfo.NowSeason = currentSeason
	}
	user.Dirty = true
	if isReset {
		ack := &pb.CompetitveLoadAck{}
		err := this.LoadInfo(user, ack)
		if err == nil {
			_ = this.GetUserManager().SendMessage(user, ack, true)
		}
	}
	return
}

func (this *CompetitveManager) CheckUser() {
	currentSeason, _, _ := this.GetCurrentSeason(base.Conf.ServerId, false)    //当前赛季
	seasonRankInfos := rmodel.Competitve.GetSeasonRankInfos(currentSeason, -1) //赛季排名

	allData := make([]int, 0)
	oldData := make([]int, 0)
	if len(seasonRankInfos) <= 0 {
		return
	}
	logger.Debug("checker seasonRankInfos:%v", seasonRankInfos)
	for i, j := 0, len(seasonRankInfos); i < j; i += 2 {
		if seasonRankInfos[i] > 0 {
			userId := int(seasonRankInfos[i])
			score := int(seasonRankInfos[i+1])
			userInfo := this.GetUserManager().GetUserBasicInfo(userId)
			if userInfo != nil {
				allData = append(allData, userId, score)
			} else {
				oldData = append(oldData, userId, score)
			}
		}
	}
	logger.Debug("checker allData:%v  oldData:%v", allData, oldData)
	if len(oldData) <= 0 {
		return
	}

	rmodel.Competitve.DelSeasonRank(currentSeason)
	for i, j := 0, len(allData); i < j; i += 2 {
		userId := allData[i]
		score := allData[i+1]
		rmodel.Competitve.ZAddSeasonRank(currentSeason, userId, score)
	}
	return
}

//
//  CompetitiveMultipleClaim
//  @Description:竞技场多倍领取
//
func (this *CompetitveManager) CompetitiveMultipleClaim(user *objs.User, op *ophelper.OpBagHelperDefault, ack *pb.CompetitveMultipleClaimAck) error {
	num := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_COMPETITVE_MULTIPLE_GET)
	if num <= 0 {
		return gamedb.ERRGETCONDITIONERR2
	}
	competitiveTimesCfg, _, openDay, err := this.CheckGameCfg(user)
	if err != nil {
		return err
	}
	//今天剩余可挑战次数
	haveChallengeTimes := user.CompetitiveInfo.HaveChallengeTimes
	if haveChallengeTimes <= 0 {
		return gamedb.ERRNOTENOUGHTIMES
	}
	vipAddLv := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_COMPETITVE_FREENUM)
	lastChallengeTimes := this.GetChallengeTimes(competitiveTimesCfg[0]+vipAddLv, haveChallengeTimes)
	if lastChallengeTimes <= 0 {
		return gamedb.ERRCOMPETITVEMULTIPLEERR
	}

	addTimes := 0
	otherTimes := gamedb.GetConf().CompetitveVip
	addTimes = otherTimes
	if lastChallengeTimes < otherTimes {
		addTimes = otherTimes - lastChallengeTimes
	}
	if addTimes <= 0 {
		return gamedb.ERRCOMPETITVEMULTIPLEERR
	}
	season, _, openDay := this.GetCurrentSeason(user.ServerId, false)
	_, beforeSeasonUserScore := rmodel.Competitve.GetSeasonSelfRankAndScore(user.Id, season)
	seasonUserScore := int(beforeSeasonUserScore)
	_, _, cfg := gamedb.GetCompetitveCfgByScore(seasonUserScore)
	items := make(gamedb.ItemInfos, 0)
	if this.getWinState(user.Id) == 1 {
		afterScore := this.GetRank().GetUserJointCombat(seasonUserScore + cfg.MarkWin*addTimes)
		rmodel.Competitve.ZAddSeasonRankInfo(season, openDay, user.Id, afterScore)
		user.SeasonWinTimes += addTimes
		this.GetCondition().RecordCondition(user, pb.CONDITION_COMPETIVE_ALL_WIN, []int{})
		this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_COMPETIVE_WIN, []int{addTimes})
		for _, info := range cfg.RewardWin {
			items = append(items, &gamedb.ItemInfo{ItemId: info.ItemId, Count: info.Count * addTimes})
		}
		this.GetBag().AddItems(user, items, op)
	} else {
		afterScore := this.GetRank().GetUserJointCombat(seasonUserScore + cfg.MarkLoss*addTimes)
		rmodel.Competitve.ZAddSeasonRankInfo(season, openDay, user.Id, afterScore)
		for _, info := range cfg.RewardLoss {
			items = append(items, &gamedb.ItemInfo{ItemId: info.ItemId, Count: info.Count * addTimes})
		}
		this.GetBag().AddItems(user, cfg.RewardLoss, op)
	}
	user.CompetitiveInfo.HaveChallengeTimes += addTimes
	user.SeasonTimes += addTimes
	_, beforeSeasonUserScore = rmodel.Competitve.GetSeasonSelfRankAndScore(user.Id, season)
	seasonUserScore = int(beforeSeasonUserScore)
	ack.SeasonScore = int32(seasonUserScore)
	//每日任务 完成一次通知
	this.GetDailyTask().CompletionOfTask(user, pb.DAILYTASKACTIVITYTYPE_JING_JI_CHANG, addTimes)
	//通知任务系统 挑战一次竞技场
	this.GetCondition().RecordCondition(user, pb.CONDITION_CHALLENGE_JIN_JI_CHANG, []int{addTimes})
	this.GetTask().AddTaskProcess(user, pb.CONDITION_CHALLENGE_JIN_JI_CHANG, -1)
	user.Dirty = true
	return nil
}

//背包使用道具购买挑战次数 检查
func (this *CompetitveManager) BagUserItemAddChallengeNumCheck(user *objs.User) error {
	competitiveTimesCfg, _, _, err := this.CheckGameCfg(user)
	if err != nil {
		return err
	}
	//今天已购买次数
	haveBuyTimes := user.CompetitiveInfo.HaveChallengeTimes
	if haveBuyTimes >= competitiveTimesCfg[1]+this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_COMPETITVE_BUYNUM) {
		logger.Error("购买次数到达上线")
		return gamedb.ERRBUYUPPERLIMIT
	}

	//今天已挑战次数
	haveChallengeTimes := user.CompetitiveInfo.HaveChallengeTimes
	if haveChallengeTimes <= 0 {
		logger.Error("BuyCompetitiveChallengeNum userId:%v  挑战次数已满", user.Id)
		return gamedb.ERRENOUGHTIMES
	}
	return nil
}

//竞技场 获取开服天数
func (this *CompetitveManager) GetCompetitiveOpenDay(serverId int) int {

	mergerId, mergerTime := this.GetSystem().GetServerMergerIdAndMergerTime(serverId)
	if mergerId > 0 {
		lenServers, _ := modelCross.GetServerInfoModel().GetAllMergerServerIds(mergerId)
		if lenServers != nil && len(lenServers) > 1 {
			return common.GetNumberOfDaysDifference(mergerTime, time.Now())
		}
	}
	return this.GetSystem().GetServerOpenDaysByServerIdByExcursionTime(serverId, 0)
}
