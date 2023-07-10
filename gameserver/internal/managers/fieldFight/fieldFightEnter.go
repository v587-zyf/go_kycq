package fieldFight

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
	"encoding/json"
	"strconv"
)

func (this *FieldManager) EnterFieldFight(user *objs.User, challengeUid, isBeatBack int) error {

	logger.Debug("EnterFieldFight  userId:%v name:%v  challengeUid:%v, isBeatBack:%v", user.Id, user.NickName, challengeUid, isBeatBack)
	if isBeatBack < 0 || isBeatBack > 1 {
		return gamedb.ERRPARAM
	}

	err := this.GetCondition().CheckFunctionOpen(user, pb.FUNCTIONID_FIELD_FIGHT)
	if err != nil {
		return err
	}

	openDay := this.GetOpenDayByReduceTime()

	if isBeatBack == 1 {
		beatBackUsers, _ := rmodel.FieldFight.GetFieldFightDefeatOwnerUsers(user.Id, openDay)
		if beatBackUsers == nil {
			logger.Error("反击失败 user.Id:%v, openDay:%v  challengeUid:%v", user.Id, openDay, challengeUid)
			return gamedb.ERRPARAM
		}

		isCanBeatBack := false
		for backUserId := range beatBackUsers {
			if backUserId == challengeUid {
				isCanBeatBack = true
			}
		}
		if !isCanBeatBack {
			logger.Error("反击失败 user.Id:%v, openDay:%v  beatBackUsers:%v  challengeUid:%v", user.Id, openDay, beatBackUsers, challengeUid)
			return gamedb.ERRPARAM
		}

	}

	fieldFightTimesCfg := gamedb.GetConf().FieldFightMaxNum
	haveChallengeTimes := user.FieldFight.HaveChallengeTimes
	lastChallengeTimes := this.GetChallengeTimes(fieldFightTimesCfg[0]+this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_FIELDFIGHT_FREENUM), haveChallengeTimes)

	if lastChallengeTimes <= 0 {
		return gamedb.ERRNOTENOUGHTIMES
	}

	if !this.GetFight().CheckInFightBefore(user, constFight.FIGHT_TYPE_FIELD_STAGE) {
		logger.Error("玩家已经在战斗中了")
		return gamedb.ERRUSERINFIGHT
	}

	//创建战斗
	fightId, err := this.GetFight().CreateFight(constFight.FIGHT_TYPE_FIELD_STAGE, []byte(strconv.Itoa(isBeatBack)))
	if err != nil {
		logger.Error("创建战斗失败:%v  fightId:%v", err, fightId)
		return err
	}

	//对手进入战斗
	err = this.GetFight().EnterFightByFightIdForUserRobot(challengeUid, fightId, constFight.FIGHT_TYPE_FIELD_STAGE, constFight.FIGHT_TEAM_ZERO)
	if err != nil {
		logger.Error("创建战斗失败:%v  fightId:%v", err, fightId)
		return err
	}
	//玩家进入战斗
	err = this.GetFight().EnterFightByFightId(user, constFight.FIGHT_TYPE_FIELD_STAGE, fightId)
	if err != nil {
		logger.Error("玩家进入战斗:%v  fightId:%v", err, fightId)
		return err
	}

	if challengeUid > 0 {
		chanllengeUserInfo := this.GetUserManager().GetUserBasicInfo(challengeUid)
		if chanllengeUserInfo != nil {
			kyEvent.FieldFightStart(user, chanllengeUserInfo.OpenId, chanllengeUserInfo.Id, chanllengeUserInfo.NickName, chanllengeUserInfo.Combat)
		}
	} else {
		kyEvent.FieldFightStart(user, "", -challengeUid, "", 0)
	}
	logger.Debug("EnterFieldFight")
	this.GetWarOrder().WriteWarOrderTask(user, pb.WARORDERCONDITION_FIELDFIGHT_NUM, []int{1})
	logger.Debug("EnterFieldFight")
	return nil
}

func (this *FieldManager) FieldFightFightEndResult(user *objs.User, result bool, challengeUid, isBeatBack int) error {
	logger.Debug("野战战斗 回调 userId:%v  name:%v  result:%v  challengeUid:%v  isBeatBack:%v", user.Id, user.NickName, result, challengeUid, isBeatBack)
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeFieldFightChallengeWinReward)
	fieldFightTimesCfg, _, openDay, err := this.CheckGameCfg()
	if err != nil {
		return err
	}

	// 推送消息
	arenaFightNtf := &pb.FieldFightNtf{}

	if result {
		arenaFightNtf.Result = 1

		challengeInfo := this.GetUserManager().GetUserBasicInfo(challengeUid)
		//赢了扣挑战次数  isBeatBack 1是反击 0不是反击
		if isBeatBack == 0 {
			//不是反击挑战扣次数
			user.FieldFight.HaveChallengeTimes += 1
			user.Dirty = true
			if challengeInfo != nil {
				logger.Info("FieldManager EnterFieldFight userId:%v  击败 challengeUid:%v ", user.Id, challengeUid)
				rmodel.FieldFight.SetFieldFightDefeatOwnerUsers(challengeUid, openDay, user.Id, user.NickName)
			}
		}
		difficultState := -1
		allRivalUserInfo := make([]*pb.FieldFightRivalUserInfo, 0)
		data, err := rmodel.FieldFight.GetFieldFightSaveBeforeRefRivals(user.Id)
		if err == nil {
			err = json.Unmarshal([]byte(data), &allRivalUserInfo)
			if err == nil {
				for _, info := range allRivalUserInfo {
					if int(info.RivalUserId) == challengeUid {
						difficultState = int(info.RivalDifficult)
					}
				}
			}
		}
		if difficultState < 0 {
			logger.Info("取劲敌玩家 数据err:%v allRivalUserInfo:%v ", err, allRivalUserInfo)
			difficultState = 3
		}

		rewardCfg := gamedb.GetFieldFightFieldBaseCfg(difficultState)
		if rewardCfg != nil {
			this.GetBag().AddItems(user, rewardCfg.Reward, op)
		}

		//反击奖励
		if isBeatBack == 1 {
			backFight := gamedb.GetConf().FieldFightBack
			if backFight != nil {

				this.GetBag().AddItems(user, backFight, op)
			}
			//del 曾经挑战失败的玩家
			rmodel.FieldFight.DelFieldFightDefeatOwnerUsers(user.Id, challengeUid, openDay)
		}
		//挑战完一个难度的玩家后 补足一个相应难度的玩家
		if isBeatBack == 0 {
			this.ComplementPeopleByDifficult(user, openDay, difficultState, 1, challengeUid, allRivalUserInfo)
		}
	}

	//今天剩余可挑战次数
	haveChallengeTimes := user.FieldFight.HaveChallengeTimes
	lastChallengeTimes := this.GetChallengeTimes(fieldFightTimesCfg[0]+this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_FIELDFIGHT_FREENUM), haveChallengeTimes)
	arenaFightNtf.RemainChallengeTimes = int32(lastChallengeTimes)
	//推送道具变化
	this.GetUserManager().SendItemChangeNtf(user, op)

	allRivalUserInfo := make([]*pb.FieldFightRivalUserInfo, 0)
	haveAppearUser := make(map[int]bool)
	robotId := 1
	data, err := rmodel.FieldFight.GetFieldFightSaveBeforeRefRivals(user.Id)
	if err == nil {
		err = json.Unmarshal([]byte(data), &allRivalUserInfo)
		if err == nil {
			for _, info := range allRivalUserInfo {
				if info.RivalUserId == 0 {
					logger.Error("野战 劲敌玩家 id 错误 玩家id:%v  劲敌id:%v", user.Id, info.RivalUserId)
					continue
				}
				_, arenaFightNtf.ListInfo, _ = this.GetListInfo(user, arenaFightNtf.ListInfo, haveAppearUser, robotId, int(info.RivalUserId), int(info.RivalDifficult))
			}
		}
	}
	if challengeUid > 0 {
		chanllengeUserInfo := this.GetUserManager().GetUserBasicInfo(challengeUid)
		if chanllengeUserInfo != nil {
			kyEvent.FieldFightEnd(user, int(arenaFightNtf.Result), chanllengeUserInfo.OpenId, chanllengeUserInfo.Id, chanllengeUserInfo.NickName, chanllengeUserInfo.Combat)
		}
	} else {
		kyEvent.FieldFightEnd(user, int(arenaFightNtf.Result), "", -challengeUid, "", 0)
	}

	arenaFightNtf.Goods = op.ToChangeItems()
	BeatBackUserInfo := this.GetRivalUserInfos(user.Id, openDay)
	arenaFightNtf.BeatBackOwnUserInfo = BeatBackUserInfo
	//每日任务 完成一次通知
	this.GetDailyTask().CompletionOfTask(user, pb.DAILYTASKACTIVITYTYPE_YE_ZHAN, 1)
	logger.Debug("arenaFightNtf:%v", arenaFightNtf)
	err = this.GetUserManager().SendMessage(user, arenaFightNtf, true)
	logger.Debug("fieldFight result:%v user.GateSessionId:%v   err:%v", result, user.GateSessionId, err)
	if result {
		challengeUser := this.GetUserManager().GetUser(challengeUid)
		if challengeUser != nil {
			ntf := &pb.BeatBackInfoNtf{}
			ntf.BeatBackOwnUserInfo = this.GetRivalUserInfos(challengeUid, openDay)
			this.GetUserManager().SendMessage(challengeUser, ntf, true)
		}
	}
	return nil
}
