package rank

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constRank"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/base"
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

const (
	MAX_RANK_NUM   = 100
	ROBOT_RANK_NUM = 200
)

type rankAward struct {
	rankType int
	mailName string
}

//排行榜奖励
var award []*rankAward

func (this *RankManager) initRankAwardConf() {
	award = []*rankAward{
		{pb.RANKTYPE_COMBAT, ""},
		{pb.RANKTYPE_LEVEL, ""},
		{pb.RANKTYPE_ARENA, ""},
	}
}

type RankManager struct {
	util.DefaultModule
	managersI.IModule

	updateChan chan RankChange
	loopChan   chan struct{}
}

type RankChange struct {
	RankType int
	Member   interface{}
	Score    int
	IsLogin  bool
	CombatUp bool
}

func NewRankManager(module managersI.IModule) *RankManager {
	return &RankManager{IModule: module}
}

func (this *RankManager) Init() error {
	this.updateChan = make(chan RankChange, 1000)
	go this.updateService()
	//初始假人排行榜
	this.initRobotRank()
	this.initRankAwardConf()
	return nil
}

func (this *RankManager) initRobotRank() {
	logger.Debug("初始化排行榜:pb.RANKTYPE_ARRAY:%v", pb.RANKTYPE_ARRAY)
	//for _, rankType := range pb.RANKTYPE_ARRAY {
	//	if rankType == pb.RANKTYPE_CHAMPIONSHIP {
	//		continue
	//	}
	//	rankKey := rmodel.Rank.GetRankKey(pb.KINGDOM_WEI, rankType)
	//	rankUsers := rmodel.Rank.GetRank(rankKey, 0, 0)
	//	if len(rankUsers) >= 2 {
	//		continue
	//	}
	//	logger.Info("初始化排行版：%v假人数据", rankKey)
	//	startRobotId := RANKROBOT_MAP[rankType]
	//	if startRobotId > 0 {
	//		for i := 0; i < ROBOT_RANK_NUM; i++ {
	//			rmodel.Rank.ZaddRankByType(rankType, pb.KINGDOM_WEI, -(startRobotId + i), gameDb().Robots[startRobotId+i].Sword)
	//		}
	//	}
	//}
}

func (this *RankManager) updateService() {
	for {
		select {
		case msg := <-this.updateChan:
			this.ZAddRank(msg.RankType, base.Conf.ServerId, msg.Member, msg.Score, msg.IsLogin, msg.CombatUp)
		}
	}
}

func (this *RankManager) Append(rankType int, member interface{}, score int, appendNow, isLogin, combatUp bool) {
	if appendNow {
		this.ZAddRank(rankType, base.Conf.ServerId, member, score, isLogin, combatUp)
		return
	}
	select {
	case this.updateChan <- RankChange{RankType: rankType, Member: member, Score: score, IsLogin: isLogin, CombatUp: combatUp}:
	default:
		logger.Warn("rankManager: updatechan is full, please check .")
	}
}

//每日排行 战力 拼接时间戳 用于相同分数玩家 先升到的排到前面
func (this *RankManager) GetUserJointCombat(combat interface{}) float64 {
	beforeSecond := time.Now().Hour()*3600 + time.Now().Minute()*60 + time.Now().Second()
	afterSecond := 86400 - beforeSecond
	if afterSecond%10 == 0 {
		afterSecond += 1
	}

	str := strconv.Itoa(afterSecond)
	for i := len(str); i < 5; i++ {
		str = fmt.Sprintf("%v%v", 0, str)
	}
	strings := fmt.Sprintf("%v.%v", combat, str)
	times, _ := strconv.ParseFloat(strings, 64)
	return times
}

func (this *RankManager) ZAddRank(rankType, serverId int, id interface{}, value int, isLogin, combatUp bool) {

	newScore := value
	if constRank.SORT_TIME_RANK[rankType] {
		timestamps := time.Now().Unix()
		newScore = int(int64(value)*int64(constRank.RANK_SORT_TIME_FIX) + (int64(9999999999) - timestamps))
	}

	if newScore <= 0 {
		return
	}

	rmodel.Rank.ZaddRankByType(rankType, base.Conf.ServerId, id, newScore)

	if isLogin {
		return
	}

	if rankType == pb.RANKTYPE_COMBAT_JEWEL || rankType == pb.RANKTYPE_COMBAT_EQUIP {
		if !combatUp {
			return
		}
	}

	openDay := this.GetSystem().GetServerOpenDaysByServerId(serverId)
	cfg := gamedb.GetDayRankingDayRankingCfg(openDay)
	if cfg == nil {
		//logger.Debug("openDay:%v", openDay)
		return
	}
	if cfg.Type == rankType {
		if this.GetDailyRank().GetAddState() {
			rmodel.Rank.ZAddDailyRankByType(rankType, base.Conf.ServerId, id, this.GetUserJointCombat(value))
		}
		return
	}
	return
}

func (this *RankManager) DelData(rankType int) {
	key := rmodel.Rank.GetRankKey(rankType, base.Conf.ServerId)
	rmodel.Rank.DelData(key)
}

func (this *RankManager) Award() {

	////开服第一天不发排行榜奖励
	//if m.System.GetServerOpenDaysByServerId(base.Conf.ServerId) == 1 {
	//	return
	//}
	//
	//defer func() {
	//	if err := recover(); err != nil {
	//		logger.Error("cronjob rank Award panic: %v, %s", err, debug.Stack())
	//	}
	//}()
	//var activeUserUser map[int]bool
	//go func() {
	//
	//}()

}

//随机时长发邮件发邮件
func (this *RankManager) SendMail(rankType int, activeUser map[int]bool) {

	//awardRankIndex := rankType
	//logger.Info("cron send user rank  reward mail start type:%d", rankType)
	//rankT := gameDb().GetRanks(rankType)
	//if rankT == nil {
	//	return
	//}
	//rankKey := rmodel.Rank.GetRankKey(rankType, base.Conf.ServerId)
	//rankUsers := rmodel.Rank.GetRank(rankKey, 0, rankT.Max-1)
	//
	//combatMap := make(map[int]int)
	//var index = 0
	//var rank = 0
	//for index < len(rankUsers) {
	//	userId := rankUsers[index]
	//	index++
	//	if userId < 0 {
	//		index++
	//		rank++
	//		continue
	//	}
	//	if !activeUser[userId] {
	//		index++
	//		rank++
	//		continue
	//	}
	//	combatMap[userId] = rank + 1
	//	index++
	//	rank++
	//}
	//for id, rank := range combatMap {
	//	awardT := gameDb().GetRankAward(rank, rankType)
	//	if awardT == nil {
	//		logger.Warn("有排名奖励没有配置 type=%v rank=%v", rankType, rank)
	//		continue
	//	}
	//	itemMap := make(gamedb.PropInfos, 0)
	//	for _, award := range awardT.Awards {
	//		info := &gamedb.PropInfo{K: award.ItemId, V: award.Count}
	//		itemMap = append(itemMap, info)
	//	}
	//	rankStr := fmt.Sprintf("%d", rank)
	//	m.Mail.SendSystemMailWithProp(id, pb.MAILTYPE_RANK_MAIL, []string{contentSign, rankStr}, itemMap)
	//}

}

//膜拜奖励
func (this *RankManager) WorshipReward(user *objs.User, op *ophelper.OpBagHelperDefault) error {

	if user.DayStateRecord.RankWorship == 1 {
		return gamedb.ERRAWARDGET
	}

	user.DayStateRecord.RankWorship = 1
	this.GetBag().AddItems(user, gamedb.GetConf().RankWorshipAward, op)

	return nil

}
