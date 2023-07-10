package competitve

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"strconv"
)

func (this *CompetitveManager) EnterCompetitveFight(user *objs.User, challengeUid, challengeRanking int) error {
	if challengeUid == 0 {
		return gamedb.ERRPARAM
	}

	competitiveTimesCfg := gamedb.GetConf().CompetitveTimes
	_, _, openDay := this.GetCurrentSeason(user.ServerId, false)

	//今天已挑战次数
	haveChallengeTimes := user.CompetitiveInfo.HaveChallengeTimes
	logger.Debug("EnterCompetitiveFight user.Id:%v, openDay:%v  haveChallengeTimes:%v competitiveTimesCfg[0]:%v this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_COMPETITVE_FREENUM):%v", user.Id, openDay, haveChallengeTimes, competitiveTimesCfg[0], this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_COMPETITVE_FREENUM))
	if haveChallengeTimes >= competitiveTimesCfg[0]+this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_COMPETITVE_FREENUM) {
		logger.Error("BuyCompetitiveChallengeNum userId:%v  挑战次数已满", user.Id)
		return gamedb.ERRENOUGHTIMES
	}

	if !this.GetFight().CheckInFightBefore(user, constFight.FIGHT_TYPE_ARENA_STAGE) {
		return gamedb.ERRUSERINFIGHT
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
	user.CompetitiveInfo.HaveChallengeTimes += 1
	user.SeasonTimes += 1

	if challengeUid > 0 {
		chanllengeUserInfo := this.GetUserManager().GetUserBasicInfo(challengeUid)
		if chanllengeUserInfo != nil {
			kyEvent.ArenaFightStart(user, chanllengeUserInfo.OpenId, chanllengeUserInfo.Id, chanllengeUserInfo.NickName, chanllengeUserInfo.Combat)
		}
	} else {
		kyEvent.ArenaFightStart(user, "", -challengeUid, "", 0)
	}
	this.GetWarOrder().WriteWarOrderTask(user, pb.WARORDERCONDITION_COMPETIVE_NUM, []int{1})
	this.GetCondition().RecordCondition(user, pb.CONDITION_CHALLENGE_JIN_JI_CHANG, []int{1})
	this.GetTask().AddTaskProcess(user, pb.CONDITION_CHALLENGE_JIN_JI_CHANG, -1)
	return nil
}

func (this *CompetitveManager) CompetitveFightEndResult(user *objs.User, result bool) {
	logger.Info("竞技场结算 userId:%v  result:%v", user.Id, result)
	season, _, openDay := this.GetCurrentSeason(user.ServerId, false)
	arenaFightNtf := &pb.ArenaFightNtf{}
	if !result {
		arenaFightNtf.Result = pb.RESULTFLAG_FAIL
		user.SeasonLoseContinueTimes += 1
		user.CompetitiveInfo.ContinuityWin = 0
	} else {
		arenaFightNtf.Result = pb.RESULTFLAG_SUCCESS
		user.SeasonLoseContinueTimes = 0
		user.CompetitiveInfo.ContinuityWin += 1
		this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_COMPETIVE_WIN, []int{1})
	}
	this.setWinState(user.Id, int(arenaFightNtf.Result))
	_, beforeSeasonUserScore := rmodel.Competitve.GetSeasonSelfRankAndScore(user.Id, season)
	seasonUserScore := int(beforeSeasonUserScore)
	_, _, cfg := gamedb.GetCompetitveCfgByScore(seasonUserScore)
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeCompetitveEndReward)
	if result {
		afterScore := this.GetRank().GetUserJointCombat(seasonUserScore + cfg.MarkWin)
		rmodel.Competitve.ZAddSeasonRankInfo(season, openDay, user.Id, afterScore)
		user.SeasonWinTimes += 1
		this.GetCondition().RecordCondition(user, pb.CONDITION_COMPETIVE_ALL_WIN, []int{})
		this.GetBag().AddItems(user, cfg.RewardWin, op)
	} else {
		afterScore := this.GetRank().GetUserJointCombat(seasonUserScore + cfg.MarkLoss)
		rmodel.Competitve.ZAddSeasonRankInfo(season, openDay, user.Id, afterScore)
		this.GetBag().AddItems(user, cfg.RewardLoss, op)
	}
	rank, _ := rmodel.Competitve.GetSeasonSelfRankAndScore(user.Id, season)
	if rank >= 0 {
		rank += 1
	}
	arenaFightNtf.MyRank = int32(rank)
	arenaFightNtf.Goods = op.ToChangeItems()
	_, beforeSeasonUserScore = rmodel.Competitve.GetSeasonSelfRankAndScore(user.Id, season)
	seasonUserScore = int(beforeSeasonUserScore)
	arenaFightNtf.SeasonScore = int32(seasonUserScore)

	kyEvent.ArenaFightEnd(user, int(arenaFightNtf.Result), rank, seasonUserScore)
	err := this.GetUserManager().SendMessage(user, arenaFightNtf, true)
	if err != nil {
		logger.Error("competitive send end ntf err:%v", err)
	}
	_ = this.GetUserManager().SendItemChangeNtf(user, op)
	//每日任务 完成一次通知
	this.GetDailyTask().CompletionOfTask(user, pb.DAILYTASKACTIVITYTYPE_JING_JI_CHANG, 1)
	return
}
