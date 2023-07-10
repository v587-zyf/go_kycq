package arena

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constMail"
	"cqserver/gamelibs/publicCon/constMax"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"strconv"
	"time"
)

const (
	ArenaRankType  = pb.RANKTYPE_ARENA
	RANK_SCORE_FIX = 10000000
	ArenaRankMax   = 998
)

// 排行榜单规则：分数越高 排名越高
func NewArenaManager(m managersI.IModule) *ArenaManager {
	arena := &ArenaManager{}
	arena.IModule = m
	return arena
}

type ArenaManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *ArenaManager) Online(user *objs.User) {
	this.ResetArena(user)
}

func (this *ArenaManager) ResetArena(user *objs.User) {
	userArena := user.Arena
	timeNowFive := common.GetDateIntByOffset5(time.Now())
	if userArena.DareDate != timeNowFive {
		userArena.DareDate = timeNowFive
		// 战斗次数
		userArena.DareNum = gamedb.GetConf().ArenaChallengeNum
		// 购买次数
		userArena.BuyDareNum = 0
		// 发奖励
		userId := user.Id
		userArenaRanking := this.GetRank().GetRanking(pb.RANKTYPE_ARENA, userId)
		// 没上榜不发
		if userArenaRanking < 0 {
			return
		}

		userArenaRanking += 1
		arenaRankConfs := gamedb.GetArenaRank()
		var arenaRewardConf *gamedb.ArenaRankArenaRankCfg
		for _, arenaRankCfg := range arenaRankConfs {
			min, max := arenaRankCfg.RankMin, arenaRankCfg.RankMax
			for ; min <= max; min++ {
				if userArenaRanking == min {
					arenaRewardConf = arenaRankCfg
				}
			}
		}
		if arenaRewardConf == nil {
			logger.Debug("not in arenaRank conf")
			return
		}
		// 整理发放奖励
		rewardItems := make([]*model.Item, 0)
		for _, itemInfo := range arenaRewardConf.Reward {
			rewardItems = append(rewardItems, &model.Item{
				ItemId: itemInfo.ItemId,
				Count:  itemInfo.Count,
			})
		}
		// 发放这个档次奖励
		this.GetMail().SendSystemMail(userId, constMail.ARENA_RANK_ID, []string{strconv.Itoa(userArenaRanking)}, rewardItems,0)
	}
}

func (this *ArenaManager) ArenaOpen(user *objs.User, ack *pb.ArenaOpenAck) error {
	// 校验开启条件
	err := this.checkArena(user, false)
	if err != nil {
		return err
	}

	userArena := user.Arena

	randArenaRival, err := this.RandArenaRival(user)
	if err != nil {
		return err
	}
	ack.ArenaRank = randArenaRival
	ack.Three = this.GetPbThree()
	ack.DareNum = int32(userArena.DareNum)
	ack.BuyDareNum = int32(userArena.BuyDareNum)
	ack.BuyDareNums = int32(userArena.BuyDareNums)

	userRanking := this.GetRank().GetRanking(ArenaRankType, user.Id)
	if userRanking >= 0 {
		userRanking += 1
	}
	ack.Ranking = int32(userRanking)
	return nil
}

func (this *ArenaManager) BuyArenaFightNum(user *objs.User, op *ophelper.OpBagHelperDefault, ack *pb.BuyArenaFightNumAck) error {
	err := this.checkArena(user, false)
	if err != nil {
		return err
	}

	userArena := user.Arena

	arenaBuyMax := gamedb.GetMaxValById(0, constMax.MAX_ARENA_BUY_NUM)
	buyNum := userArena.BuyDareNum + 1
	// 校验今日购买次数
	if buyNum > arenaBuyMax {
		return gamedb.ERRENOUGHTIMES
	}

	buyConf := gamedb.GetArenaBuyArenaBuyCfg(buyNum)
	err = this.GetBag().Remove(user, op, buyConf.Cost.ItemId, buyConf.Cost.Count)
	if err != nil {
		return err
	}

	userArena.BuyDareNum = buyNum
	userArena.BuyDareNums += 1
	user.Dirty = true

	ack.Goods = op.ToChangeItems()
	ack.DareNum = int32(userArena.DareNum)
	ack.BuyDareNum = int32(userArena.BuyDareNum)
	ack.BuyDareNums = int32(userArena.BuyDareNums)
	return nil
}

func (this *ArenaManager) RefArenaRank(user *objs.User, ack *pb.RefArenaRankAck) error {
	randArenaRival, err := this.RandArenaRival(user)
	if err != nil {
		return err
	}
	ack.ArenaRank = randArenaRival
	ack.Three = this.GetPbThree()
	return nil
}

// 随机四个挑战|碾压对手
func (this *ArenaManager) RandArenaRival(user *objs.User) ([]*pb.ArenaRank, error) {
	userId := user.Id
	pbArenaRank := make([]*pb.ArenaRank, 0)
	userArenaRanking := this.GetRank().GetRanking(ArenaRankType, userId)
	count := this.GetRank().GetCount(ArenaRankType)
	if count == 0 {
		return pbArenaRank, nil
	}
	// 防止榜单人数不足指定值
	if userArenaRanking > ArenaRankMax || userArenaRanking < 0 {
		if count > ArenaRankMax {
			userArenaRanking = ArenaRankMax
		} else {
			userArenaRanking = count
		}
	}
	// 配置
	arenaRankConf := gamedb.GetArenaMatch()
	var userArenaRankingCfg *gamedb.ArenaMatchArenaMatchCfg
	for _, cfg := range arenaRankConf {
		if userArenaRanking+1 >= cfg.RankMin && userArenaRanking+1 <= cfg.RankMax {
			userArenaRankingCfg = cfg
			break
		}
	}
	if userArenaRankingCfg == nil {
		logger.Error("ArenaMatch not found ranking is %v", userArenaRanking)
		return nil, gamedb.ERRPARAM
	}
	// 排名比指定值高  取一个碾压
	if userArenaRanking < ArenaRankMax {
		start := userArenaRanking
		end := start + userArenaRankingCfg.RangeLow
		if end > ArenaRankMax {
			end = ArenaRankMax
		}
		pbArenaRank = this.randRankUsers(end, start, 1, userArenaRanking, pbArenaRank)
	}
	// 3|4 挑战者
	end := userArenaRanking
	start := end - userArenaRankingCfg.RangeHigh
	if userArenaRanking < 5 {
		start = 0
		end = 5
	} else {
		if start < 0 {
			start = 0
		}
	}
	pbArenaRank = this.randRankUsers(end, start, 4, userArenaRanking, pbArenaRank)

	return pbArenaRank, nil
}

func (this *ArenaManager) randRankUsers(end, start, needNum int, userArenaRanking int, pbArenaRank []*pb.ArenaRank) []*pb.ArenaRank {
	hasMap := make(map[int]int)
	if len(pbArenaRank) != 0 {
		for _, arenaRank := range pbArenaRank {
			hasMap[int(arenaRank.Ranking-1)] = int(arenaRank.Ranking - 1)
		}
	}
	for i := 0; i < 10000; i++ {
		ranking := common.RandNum(start, end)
		if ranking == userArenaRanking {
			continue
		}
		if _, have := hasMap[ranking]; have {
			continue
		}
		redisRank := this.GetRank().GetRankByScore(ArenaRankType, ranking, ranking)
		if len(redisRank) != 2 {
			continue
		}
		hasMap[ranking] = ranking
		rankUid := redisRank[0]

		pbArenaRank = append(pbArenaRank, &pb.ArenaRank{
			Ranking:  int32(ranking + 1),
			Userinfo:this.GetUserManager().BuilderBrieUserInfo(rankUid),
		})
		if len(pbArenaRank) >= needNum {
			break
		}
	}
	return pbArenaRank
}

func (this *ArenaManager) checkArena(user *objs.User, checkRoll bool) error {
	openCfg := gamedb.GetFunctionFunctionCfg(pb.FUNCTIONTYPE_ARENA)
	isOpen := this.GetCondition().CheckMulti(user, -1, openCfg.Condition)
	if !isOpen {
		return gamedb.ERRNOTACTIVE
	}
	// 校验碾压
	if checkRoll {
		rollCfg := gamedb.GetFunctionFunctionCfg(pb.FUNCTIONTYPE_ARENAROLL)
		isRollOpen := this.GetCondition().CheckMulti(user, -1, rollCfg.Condition)
		if !isRollOpen {
			return gamedb.ERRNOTACTIVE
		}
	}
	return nil
}

func (this *ArenaManager) rankScore(rank int) int {
	return 1 * RANK_SCORE_FIX / (rank + 1)
}

//func AddRedis() {
//	allUsers, _ := modelGame.GetUserModel().LoadAllUsers()
//	for i := 0; i < len(allUsers)-1; i++ {
//		if allUsers[i].Id == 161000 {
//			this.GetRank().Append(ArenaRankType, allUsers[i].Id, 206)
//			continue
//		}
//		logger.Info("-----------------添加玩家：%v竞技场排行版", allUsers[i].Id)
//		this.GetRank().Append(ArenaRankType, allUsers[i].Id, i+1)
//	}
//}
