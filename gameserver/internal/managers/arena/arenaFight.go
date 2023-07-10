package arena

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
	"strconv"
)

func (this *ArenaManager) EnterArenaFight(user *objs.User, challengeUid, challengeRanking int) error {
	// 校验开启条件
	err := this.checkArena(user, false)
	if err != nil {
		return err
	}

	userArena := user.Arena

	fightUserId, _ := rmodel.Rank.GetArenaRankFight(challengeRanking)
	if fightUserId > 0 {
		return gamedb.ERRRANKFIGHT
	}

	// 校验战斗次数
	if userArena.DareNum <= 0 && userArena.BuyDareNums <= 0 {
		return gamedb.ERRNOTENOUGHTIMES
	}

	// 校验挑战者排名
	userRanking := this.GetRank().GetRanking(ArenaRankType, challengeUid)
	if userRanking+1 != challengeRanking {
		return gamedb.ERRRANKINGACHANGE
	}

	//创建战斗
	fightId, err := this.GetFight().CreateFight(constFight.FIGHT_TYPE_ARENA_STAGE, []byte(strconv.Itoa(challengeRanking)))
	if err != nil {
		return err
	}
	//对手进入战斗
	err = this.GetFight().EnterFightByFightIdForUserRobot(challengeUid, fightId, constFight.FIGHT_TYPE_ARENA_STAGE, constFight.FIGHT_TEAM_ZERO)
	if err != nil {
		return err
	}
	//玩家进入战斗
	err = this.GetFight().EnterFightByFightId(user, constFight.FIGHT_TYPE_ARENA_STAGE, fightId)
	if err != nil {
		return err
	}

	//设置排名挑战中
	rmodel.Rank.SetArenaRankFight(challengeRanking, user.Id)

	//// 进入战斗
	//this.ArenaFightResult(user, true, challengeUid)

	return nil
}

func (this *ArenaManager) ArenaFightResult(user *objs.User, result bool, challengeRank int) {
	userArena := user.Arena
	arenaFightNtf := &pb.ArenaFightNtf{}
	// 减少用户战斗次数
	if userArena.DareNum <= 0 {
		userArena.BuyDareNums -= 1
	} else {
		userArena.DareNum -= 1
	}
	user.Dirty = true

	//设置排名挑战中
	rmodel.Rank.DelArenaRankFight(challengeRank)
	// 返回结果
	if !result {
		arenaFightNtf.Result = pb.RESULTFLAG_FAIL
	} else {
		arenaFightNtf.Result = pb.RESULTFLAG_SUCCESS
	}
	// 不管输赢 发放奖励
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeArena)
	arenaReward := gamedb.GetConf().ArenaReward
	this.GetBag().AddItems(user, arenaReward, op)

	arenaFightNtf.Goods = op.ToChangeItems()
	//推送道具变化
	this.GetUserManager().SendItemChangeNtf(user, op)
	// 更新排名
	if result {
		this.UpArenaRank(user, challengeRank)
	}
	// 推送消息
	this.GetUserManager().SendMessage(user, arenaFightNtf, true)
}

func (this *ArenaManager) LeaveFight(challengeRanking int) {
	rmodel.Rank.DelArenaRankFight(challengeRanking)
}

// 更新榜单
func (this *ArenaManager) UpArenaRank(user *objs.User, challengeRank int) {
	challengeRank -= 1
	userId := user.Id
	userRanking := this.GetRank().GetRanking(ArenaRankType, userId)
	challengeRankInfo := this.GetRank().GetRankByScore(ArenaRankType, challengeRank, challengeRank)
	if len(challengeRankInfo) == 2 && challengeRank < userRanking {
		// 交换名次
		this.GetRank().Append(ArenaRankType, userId, this.rankScore(challengeRank), false, false, false)
		this.GetRank().Append(ArenaRankType, challengeRankInfo[0], this.rankScore(userRanking), false, false, false)
	}
}
