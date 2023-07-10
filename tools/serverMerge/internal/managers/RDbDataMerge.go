package managers

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gamelibs/publicCon/constMail"
	"cqserver/golibs/logger"
	"cqserver/golibs/redisdb"
	"cqserver/golibs/util"
	"cqserver/tools/serverMerge/internal/rmodel"
	"fmt"
	"strconv"
	"strings"
	"time"
)

//玩家redis数据合并
type RdbDataMerge struct {
	util.DefaultModule
	serverInfos map[int]*modelCross.ServerInfo
}

func NewRdbDataMerge() *RdbDataMerge {
	return &RdbDataMerge{
		serverInfos: make(map[int]*modelCross.ServerInfo),
	}
}

func (this *RdbDataMerge) Init() error {
	return nil
}

//开始合并数据
func (this *RdbDataMerge) Merge(mergeServerInfos []modelCross.ServerInfo) bool {

	//清空redis的数据
	logger.Info("置空新redis数据库中数据")
	rmodel.RedisDb().FlushDb()
	//开始redis数据合并
	for _, v := range mergeServerInfos {
		logger.Info("服务器[%v][%v]redis数据合并开始", v.ServerId, v.RedisAddr)
		redisAddr := strings.Split(v.RedisAddr, ":")
		if len(redisAddr) < 3 {
			logger.Error("redis 数据合并，服务器:%v redis配置错误，redis:%v", v.Id, v.RedisAddr)
			return false
		}
		db := 0
		if len(redisAddr) == 4 {
			db, _ = strconv.Atoi(redisAddr[3])
		}
		redisDb := &redisdb.RedisDb{}
		redisDb.Init("tcp", redisAddr[0]+":"+redisAddr[1], redisAddr[2], db)

		//日常任务
		if ok := this.dailyTaskMerge(v.ServerId, redisDb); !ok {
			return false
		}

		//竞技场赛季奖励结算
		this.competitiveMergeSendMail(v.ServerId, redisDb)
		//首领击杀
		this.killBoss(redisDb)
		//首暴
		this.firstDrop(redisDb)

		logger.Info("服务器[%v][%v]redis数据合并结束", v.ServerId, v.RedisAddr)

	}
	logger.Info("redis数据合并结束")
	return true
}

//日常任务数据合并
func (this *RdbDataMerge) dailyTaskMerge(serverId int, redisCs *redisdb.RedisDb) bool {

	logger.Info("模块日常任务数据合并开始")
	date := time.Now().Add(-time.Hour * 5).Format("060102")
	keys, err := redisCs.Keys(fmt.Sprintf(rmodel.GuildDailyCondition_, date))
	if err != nil {
		logger.Error("模块日常任务数据合并获取keys错误:%v", err)
		return false
	}

	for _, key := range keys {
		taskProcessMap, err := redisCs.HgetallIntMap(key)
		if err != nil {
			logger.Error("模块日常任务数据合并获取key:%v错误:%v", key, err)
			return false
		}
		//写入新redis
		if len(taskProcessMap) > 0 {
			_, err := rmodel.RedisDb().HmsetIntMap(key, taskProcessMap)
			if err != nil {
				logger.Error("模块日常任务数据合并设置key:%v错误:%v", key, err)
				return false
			}
		}
	}
	logger.Info("模块日常任务数据合并结束")
	return true
}

//竞技场积分榜  发送奖励
func (this *RdbDataMerge) competitiveMergeSendMail(serverId int, redisCs *redisdb.RedisDb) bool {

	season, _, openDay := m.BaseFunctionMerge.getCurrentSeason(serverId, false)
	key := fmt.Sprintf(rmodel.CompetitiveSeasonRankInfo, season)
	key1 := fmt.Sprintf(rmodel.CompetitiveSeasonSendRewardMark, season)
	ranks, err := redisCs.ZRevrange(key, 0, 1000).ValuesIntSlice()
	haveSendSeason, _ := redisCs.Get(key1).IntDef(0)
	logger.Info("竞技场 发送赛季奖励  haveSendSeason:%v  serverId:%v  season:%v  key:%v  ranks:%v", haveSendSeason, serverId, season, key, ranks)
	if haveSendSeason == season {
		return false
	}

	if err != nil {
		logger.Error("GetSeasonRankInfos获取排行榜数据异常,key:%v,err:%v", key, err)
		return false
	}

	if len(ranks) <= 0 {
		return false
	}
	rank := 0
	for i, j := 0, len(ranks); i < j; i += 2 {
		rank++
		userId := ranks[i]
		reward, _ := gamedb.GetCompetitveSeasonEndReward(rank)
		//rewardInfo := gamedb.ItemInfos{}
		//rewardInfo = reward
		logger.Info("competitiveMergeSendMail userId:%v  rank:%v reward:%v", userId, rank, reward)
		if len(reward) > 0 {
			rewards := make(map[int]int)
			userInfo := m.DbDataMerge.GetUserInfoByUserId(userId)
			if userInfo == nil {
				continue
			}
			dailyReward := m.BaseFunctionMerge.GetDailyReward(userInfo, redisCs, season, openDay)
			//monthCardPrivilege := m.BaseFunctionMerge.GetMonthCardPrivilege(userInfo, pb.VIPPRIVILEGE_COMPETITVE_DAILY_REWARD)
			//if monthCardPrivilege > 0 {
			//	for _, v := range rewardInfo {
			//		count := common.CalcTenThousand(monthCardPrivilege, v.Count)
			//		v.Count = count
			//	}
			//}
			for _, v := range reward {
				rewards[v.ItemId] += v.Count
			}
			m.BaseFunctionMerge.SendMail(userId, constMail.COMPETITVE_RANK_REWARD, rewards, nil)
			if len(dailyReward) > 0 {
				m.BaseFunctionMerge.SendMail(userId, constMail.MAILTYPE_COMPETITIVE_DAILY_REWARD, dailyReward, nil)
			}
		}
	}
	return true
}

//首领击杀
func (this *RdbDataMerge) killBoss(redisDb *redisdb.RedisDb) bool {
	logger.Info("首领击杀合并redis开始")
	keys, err := redisDb.Keys(fmt.Sprintf(rmodel.Boss_Kill_, constFight.FIGHT_TYPE_PERSON_BOSS))
	if err != nil {
		logger.Error("首领击杀合并获取keys err:%v", err)
		return false
	}
	for _, key := range keys {
		dataMap, err := redisDb.HgetallIntMap(key)
		if err != nil {
			logger.Error("首领击杀合并获取数据错误key:%v err:%v", key, err)
			return false
		}
		if len(dataMap) > 0 {
			_, err := rmodel.RedisDb().HmsetIntMap(key, dataMap)
			if err != nil {
				logger.Error("首领redis合并设置key:%v err:%v", key, err)
				return false
			}
		}
	}
	logger.Info("首领击杀合并redis结束")
	return true
}

//首领击杀
func (this *RdbDataMerge) firstDrop(redisDb *redisdb.RedisDb) bool {
	logger.Info("首爆合并redis开始")
	keys, err := redisDb.Keys(rmodel.First_drop_item_all_get_num)
	if err != nil {
		logger.Error("首爆合并获取keys err:%v", err)
		return false
	}
	for _, key := range keys {
		dataMap, err := redisDb.HgetallIntMap(key)
		if err != nil {
			logger.Error("首爆合并获取数据错误key:%v err:%v", key, err)
			return false
		}
		if len(dataMap) > 0 {
			_, err := rmodel.RedisDb().HmsetIntMap(key, dataMap)
			if err != nil {
				logger.Error("首爆redis合并设置key:%v err:%v", key, err)
				return false
			}
		}
	}
	logger.Info("首爆合并redis结束")
	return true
}
