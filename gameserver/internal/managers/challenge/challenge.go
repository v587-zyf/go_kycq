package challenge

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/publicCon/constAuction"
	"cqserver/gamelibs/publicCon/constChallenge"
	"cqserver/gamelibs/publicCon/constMail"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"encoding/json"
	"fmt"
	"runtime/debug"
	"strconv"
	"time"
)

func NewChallengeManager(module managersI.IModule) *ChallengeManager {
	return &ChallengeManager{IModule: module, opChan: make(chan opMsg, 50)}
}

type ChallengeManager struct {
	util.DefaultModule
	managersI.IModule
	opChan chan opMsg
}

type opMsg struct {
	opType        int
	challengeInfo *modelCross.Challenge
}

func (this *ChallengeManager) Init() error {
	go this.dbDataOp()
	return nil
}

func (this *ChallengeManager) LoadInfo(serverId int, ack *pb.ChallengeInfoAck) error {
	season := this.WeekByDate(time.Now().AddDate(0, 0, -1))
	if err := this.JudgeIsInApplyTime(serverId); err != nil {
		lastSeason := this.GetBeforeSeason(season)
		season = lastSeason
		logger.Error("展示上一期活动  season:%v  lastSeason:%v", season, lastSeason)
	}

	nowRoundId := this.GetChallengeNowRound()
	logger.Info("season:%v  nowRoundId:%v", season, nowRoundId)
	crossFsId := this.GetSystem().GetServerIndexCrossFsId(serverId)
	if nowRoundId < 0 {
		return nil
	}

	logger.Debug("challenge LoadInfo nowRoundId:%v   season:%v", nowRoundId, season)
	for round := constChallenge.FIRST_ROUND; round <= nowRoundId; round++ {
		nowRoundUserIds, err := modelCross.GetChallengeDataModel().GetRoundUserIdsByCrossFsId(crossFsId, round, season)
		if err != nil {
			logger.Error("GetAllServerApplyUsersListByRound  user.ServerId:%v, round:%v err:%v", serverId, round, err)
			continue
		}
		peoples := &pb.PeopleInfos{}
		logger.Debug("len(nowRoundUserIds.UserIds):%v  round:%v", len(nowRoundUserIds.UserIds), round)
		for _, userId := range nowRoundUserIds.UserIds {
			peoplesInfo := this.GetChallengePeopleInfo1(userId, crossFsId, season)
			if peoplesInfo == nil {
				logger.Error("GetChallengePeopleInfo userId:%v nil", userId)
				continue
			}
			peoples.PeopleInfo = append(peoples.PeopleInfo, peoplesInfo)
		}
		if ack.ChallengePeopleInfo == nil {
			ack.ChallengePeopleInfo = make(map[int32]*pb.PeopleInfos, 0)
		}
		ack.ChallengePeopleInfo[int32(round)] = peoples
	}
	if err := this.JudgeIsInApplyTime(serverId); err == nil {
		allUserIds := rmodel.Challenge.HGetAllBottomUsers(nowRoundId, season)
		for u, u1 := range allUserIds {
			peoplesInfo := this.GetChallengePeopleInfo2(u, crossFsId, season)
			peoplesInfo1 := this.GetChallengePeopleInfo1(u1, crossFsId, season)
			if peoplesInfo == nil || peoplesInfo1 == nil {
				logger.Error("GetChallengePeopleInfo userId:%v nil", u)
				continue
			}
			ack.BottomUserInfo = append(ack.BottomUserInfo, peoplesInfo, peoplesInfo1)
		}
	}

	appUserInfos := make([]*modelCross.Challenge, 0)
	var err error
	if nowRoundId <= 0 {
		appUserInfos, err = modelCross.GetChallengeModel().GetAllServerApplyUsersListByRound1(season, crossFsId)
	} else {
		appUserInfos, err = modelCross.GetChallengeModel().GetAllServerApplyUsersListByRound(season, crossFsId)
	}
	if err != nil {
		logger.Error("GetAllServerApplyUsersListByRound serverId:%v err:%v", serverId, err)
	}
	if appUserInfos != nil && nowRoundId < 9 {
		for _, data := range appUserInfos {
			ack.ApplyUserInfo = append(ack.ApplyUserInfo, &pb.PeopleInfo{UserId: int32(data.UserId), Name: data.NickName, Avatar: data.Avatar, ServerId: int32(data.ServerId), Combat: data.Combat, GuildName: data.GuildName})
		}
	}

	if crossFsId > 0 {
		infos, _ := modelCross.GetServerInfoModel().GetAllServerIdsByCrossFsIds(crossFsId)
		if infos != nil {
			for _, data := range infos {
				if data.IsClose > 0 {
					continue
				}
				ack.JoinServer = append(ack.JoinServer, int32(data.ServerId))
			}
		}
	}
	return nil
}

func (this *ChallengeManager) SetApplyUserInfo(user *objs.User, ack *pb.ApplyChallengeAck) error {

	cfg := gamedb.GetCrossArenaTimeCrossArenaTimeCfg(1)
	t := int(time.Now().Weekday())
	if t == 0 {
		t = 7
	}

	if t < cfg.SignUpBegin || t > cfg.SignUpEnd {
		return gamedb.ERRCHALLENGE6
	}

	if t == cfg.SignUpEnd && common.GetTimeSeconds(time.Now()) > cfg.SignUpEndTime.GetSecondsFromZero() {
		return gamedb.ERRCHALLENGE6
	}

	err1 := this.JudgeIsInApplyTime(user.ServerId)
	if err1 != nil {
		logger.Error("活动未开启")
		return err1
	}

	ok := this.GetCondition().CheckMulti(user, -1, cfg.Condition)
	if !ok {
		return gamedb.ERRTASKISNOTOVER
	}

	if time.Now().Unix()-user.ChallengeApplyTime < 60 {
		return gamedb.ERRCHALLENGE5
	}

	serverInfo := this.GetSystem().GetServerInfoByServerId(user.ServerId)
	guildName := this.GetGuild().GetGuildName(user.Id)
	info := builder.BuildApplyChallengeUserInfoAck(user, serverInfo.CrossFsId, serverInfo.Name, guildName)
	fightUser := this.GetFight().GetFightUserInfo(user, 0, 0, true, false)
	bytes, err := json.Marshal(fightUser)
	if err != nil {
		return gamedb.ERRCHALLENGE1
	}
	ntf := &pb.ChallengeApplyUserInfoNtf{}
	info.FightUserInfo = string(bytes)
	season := this.WeekByDate(time.Now().AddDate(0, 0, -1))
	data, err := modelCross.GetChallengeModel().CheckUserIsInRound(user.ServerId, constChallenge.Group, user.Id, serverInfo.CrossFsId, season)
	if data == nil || err != nil {
		info.ExpireTime = time.Now().Unix() + 86400*14
		info.Season = this.WeekByDate(time.Now().AddDate(0, 0, -1))
		info.GuildName = guildName
		modelCross.GetChallengeModel().DbMap().Insert(info)
	} else {
		data.NickName = user.NickName
		data.Avatar = user.Avatar
		data.Combat = int64(user.Combat)
		data.FightUserInfo = info.FightUserInfo
		data.GuildName = guildName
		data.ExpireTime = time.Now().Unix() + 86400*14
		modelCross.GetChallengeModel().DbMap().Update(data)
	}
	user.ChallengeApplyTime = time.Now().Unix()
	isFirst := false
	nowRoundId := this.GetChallengeNowRound()
	if nowRoundId > 0 {
		appUserInfos, _ := modelCross.GetChallengeModel().GetAllServerApplyUsersListByRound(season, serverInfo.CrossFsId)
		if appUserInfos != nil {
			for _, data := range appUserInfos {
				guildName := this.GetGuild().GetGuildName(data.UserId)
				info := &pb.PeopleInfo{}
				if user.Id == data.UserId {
					isFirst = true
					info = &pb.PeopleInfo{UserId: int32(data.UserId), Name: user.NickName, Avatar: user.Avatar, ServerId: int32(user.ServerId), Combat: int64(user.Combat), GuildName: guildName}
					ack.ApplyUserInfo = append(ack.ApplyUserInfo, info)
				} else {
					info = &pb.PeopleInfo{UserId: int32(data.UserId), Name: data.NickName, Avatar: data.Avatar, ServerId: int32(data.ServerId), Combat: int64(data.Combat), GuildName: guildName}
					ack.ApplyUserInfo = append(ack.ApplyUserInfo, info)
				}
			}
		}
	}
	if !isFirst {
		info := &pb.PeopleInfo{UserId: int32(user.Id), Name: user.NickName, Avatar: user.Avatar, ServerId: int32(user.ServerId), Combat: int64(user.Combat), GuildName: guildName}
		ack.ApplyUserInfo = append(ack.ApplyUserInfo, info)
	}
	ntf.ApplyUserInfo = append(ntf.ApplyUserInfo, &pb.PeopleInfo{UserId: int32(user.Id), Name: user.NickName, Avatar: user.Avatar, ServerId: int32(user.ServerId), Combat: int64(user.Combat), GuildName: guildName})
	this.BroadcastAll(ntf)
	this.SendMsgToCCS(0, &pbserver.ChallengeAppuserUpNtf{CrossFsId: int32(serverInfo.CrossFsId)})
	return nil
}

//当前轮参赛的玩家信息
func (this *ChallengeManager) EachRoundPeople(user *objs.User, ack *pb.ChallengeEachRoundPeopleAck) error {

	if err := this.JudgeIsInApplyTime(user.ServerId); err != nil {
		logger.Error("活动未开启")
		return nil
	}

	season := this.WeekByDate(time.Now().AddDate(0, 0, -1))
	crossFsId := this.GetSystem().GetServerIndexCrossFsId(user.ServerId)
	nowRoundId := this.GetChallengeNowRound()
	if nowRoundId <= 0 {
		return nil
	}

	nowRoundUserIds, err := modelCross.GetChallengeDataModel().GetRoundUserIdsByCrossFsId(crossFsId, nowRoundId, season)
	if err != nil {
		logger.Error("GetAllServerApplyUsersListByRound  user.ServerId:%v, round:%v", user.ServerId, nowRoundId)
		return nil
	}
	peoples := &pb.PeopleInfos{}
	for _, userId := range nowRoundUserIds.UserIds {
		peoplesInfo := this.GetChallengePeopleInfo1(userId, crossFsId, season)
		if peoplesInfo == nil {
			logger.Error("GetChallengePeopleInfo userId:%v nil", userId)
			continue
		}
		peoples.PeopleInfo = append(peoples.PeopleInfo, peoplesInfo)
	}
	allUserIds := rmodel.Challenge.HGetAllBottomUsers(nowRoundId, season)
	for u, u1 := range allUserIds {
		peoplesInfo := this.GetChallengePeopleInfo2(u, crossFsId, season)
		peoplesInfo1 := this.GetChallengePeopleInfo1(u1, crossFsId, season)
		if peoplesInfo == nil || peoplesInfo1 == nil {
			logger.Error("GetChallengePeopleInfo userId:%v nil", u)
			continue
		}
		ack.BottomUserInfo = append(ack.BottomUserInfo, peoplesInfo, peoplesInfo1)
	}
	ack.ChallengePeopleInfo = peoples.PeopleInfo
	ack.NowRound = int32(nowRoundId)

	return nil
}

//下注
func (this *ChallengeManager) BottomPour(user *objs.User, bottomUser int, op *ophelper.OpBagHelperDefault, ack *pb.BottomPourAck) error {

	if bottomUser <= 0 {
		robotCfg := gamedb.GetCrossArenaRobotCrossArenaRobotCfg(-bottomUser)
		if robotCfg == nil {
			return gamedb.ERRPARAM
		}
	}

	if err := this.JudgeIsInApplyTime(user.ServerId); err != nil {
		logger.Error("活动未开启")
		return nil
	}

	season := this.WeekByDate(time.Now().AddDate(0, 0, -1))
	crossFsId := this.GetSystem().GetServerIndexCrossFsId(user.ServerId)

	nowRoundIndex := this.GetChallengeNowRound()
	if nowRoundIndex <= 0 {
		return gamedb.ERRCHALLENGE2
	}
	state, err := modelCross.GetChallengeModel().CheckUserIsInRound1(nowRoundIndex, bottomUser, crossFsId, season)
	if state == nil || err != nil {
		logger.Error("CheckUserIsInRound  crossFsId:%v   user.ServerId:%v, nowRoundIndex:%v, bottomUser:%v state:%v  season:%v  err:%v", crossFsId, user.ServerId, nowRoundIndex, bottomUser, state, season, err)
		return gamedb.ERRPARAM
	}
	if rmodel.Challenge.CheckBottomUserIsExist(nowRoundIndex, user.Id, season) == 1 {
		return gamedb.ERRCHALLENGE3
	}
	cfg := gamedb.GetConf().CrossArenaGamble
	if cfg == nil {
		logger.Error("GetCrossArenaRewardCrossArenaRewardCfg nil state:%v", nowRoundIndex)
		return gamedb.ERRSETTINGNOTFOUND
	}
	ok, _ := this.GetBag().HasEnough(user, cfg[0].ItemId, cfg[0].Count)
	if !ok {
		return gamedb.ERRNOTENOUGHGOODS
	}

	err = this.GetBag().Remove(user, op, cfg[0].ItemId, cfg[0].Count)
	if err != nil {
		return err
	}

	logger.Debug("nowRoundIndex:%v, user.Id:%v, bottomUser:%v, season:%v", nowRoundIndex, user.Id, bottomUser, season)
	rmodel.Challenge.HSetBottomUser(nowRoundIndex, user.Id, bottomUser, season)

	allUserIds := rmodel.Challenge.HGetAllBottomUsers(nowRoundIndex, season)
	for u, u1 := range allUserIds {
		peoplesInfo := this.GetChallengePeopleInfo2(u, crossFsId, season)
		peoplesInfo1 := this.GetChallengePeopleInfo1(u1, crossFsId, season)
		if peoplesInfo == nil || peoplesInfo1 == nil {
			logger.Error("GetChallengePeopleInfo userId:%v nil  u1:%v", u, u1)
			continue
		}
		ack.BottomUserInfo = append(ack.BottomUserInfo, peoplesInfo, peoplesInfo1)
	}
	ack.State = true
	return nil
}

func (this *ChallengeManager) JudgeIsInApplyTime(serverId int) error {
	cfg := gamedb.GetCrossArenaTimeCrossArenaTimeCfg(1)
	if this.GetSystem().GetServerOpenDaysByServerId(serverId) < cfg.OpenDayMin {
		return gamedb.ERRCHALLENGE2
	}

	t := int(time.Now().Weekday())
	if t == 0 {
		t = 7
	}
	//周几限制
	if t < cfg.SignUpBegin || t > cfg.SignUpEnd {
		return gamedb.ERRCHALLENGE2
	}

	if t == cfg.SignUpBegin && common.GetTimeSeconds(time.Now()) < cfg.SignUpBeginTime.GetSecondsFromZero() {

		return gamedb.ERRCHALLENGE2
	}

	if t == cfg.SignUpEnd && common.GetTimeSeconds(time.Now()) > cfg.CloseTime.GetSecondsFromZero() {
		return gamedb.ERRCHALLENGE2
	}

	crossFsId := this.GetSystem().GetServerIndexCrossFsId(serverId)
	if crossFsId <= 0 {
		return gamedb.ERRCHALLENGE2
	}

	return nil
}

//
//  SendLoseReward
//  @Description: 给每一轮失败者 发送失败奖励
//  @receiver this
//  @param userIds
//  @param roundIndex
//
func (this *ChallengeManager) SendLoseReward(loseUserIds, winUsers []int32, roundIndex int, winUserId int) {
	logger.Info(" SendLoseReward  loseUserIds:%v, winUsers:%v  roundIndex:%v, winUserId:%v", loseUserIds, winUsers, roundIndex, winUserId)
	cfg := gamedb.GetCrossArenaRewardCrossArenaRewardCfg(roundIndex)
	if cfg == nil {
		logger.Error("GetCrossArenaRewardCrossArenaRewardCfg nil state:%v", roundIndex)
		return
	}

	for _, userId := range loseUserIds {
		if userId > 0 {
			userInfo := this.GetUserManager().GetAllUserInfoIncludeOfflineUser(int(userId))
			if userInfo != nil {
				kyEvent.CrossChallenge(userInfo, roundIndex, false)
			}
			this.GetMail().SendSystemMailWithItemInfos(int(userId), constMail.MAILTYPE_CHALLENGE, []string{strconv.Itoa(roundIndex)}, cfg.Reward)
		}
	}

	if winUserId > 0 {
		cfg = gamedb.GetCrossArenaRewardCrossArenaRewardCfg(constChallenge.EIGHTH_ROUND + 1)
		if cfg == nil {
			logger.Error("GetCrossArenaRewardCrossArenaRewardCfg nil state:%v", roundIndex)
			return
		}
		this.GetMail().SendSystemMailWithItemInfos(int(winUserId), constMail.MAILTYPE_CHALLENGE_WIN, []string{strconv.Itoa(roundIndex)}, cfg.Reward)
	}
	for _, winId := range winUsers {
		userInfo := this.GetUserManager().GetAllUserInfoIncludeOfflineUser(int(winId))
		if userInfo != nil {
			kyEvent.CrossChallenge(userInfo, roundIndex, true)
		}
	}

	this.SendBottomReward(winUsers)
	this.BroadcastAll(&pb.ChallengeRoundEndNtf{})
	return
}

//发送下注奖励
func (this *ChallengeManager) SendBottomReward(userIds []int32) {
	cfg := gamedb.GetConf().CrossArenaGamble
	if cfg == nil {
		logger.Error("GetCrossArenaRewardCrossArenaRewardCfg nil ")
		return
	}
	crossFsId := this.GetSystem().GetServerIndexCrossFsId(base.Conf.ServerId)
	season := this.WeekByDate(time.Now().AddDate(0, 0, -1))
	nowRoundIndex := this.GetChallengeNowRound()

	winUserState := make(map[int]bool)
	for _, winUserId := range userIds {

		winUserState[int(winUserId)] = true
	}
	bottomUsers := rmodel.Challenge.HGetAllBottomUsers(nowRoundIndex-1, season)
	logger.Debug("跨服擂台赛 发送下注奖励 crossFsId:%v,   nowRoundIndex:%v nowRoundIndex-1:%v  loseUserState:%v  bottomUsers:%v  season:%v", crossFsId, nowRoundIndex, nowRoundIndex-1, winUserState, bottomUsers, season)
	for userId, bUser := range bottomUsers {
		if winUserState[bUser] {
			returnItem := gamedb.ItemInfos{&gamedb.ItemInfo{ItemId: cfg[1].ItemId, Count: cfg[1].Count}}
			this.GetMail().SendSystemMailWithItemInfos(int(userId), constMail.MAILTYPE_BOTTLE_WIN, []string{strconv.Itoa(1)}, returnItem)
		}
	}
}

func (this *ChallengeManager) GetChallengePeopleInfo(userInfo *modelCross.Challenge) *pb.PeopleInfo {
	return &pb.PeopleInfo{Name: userInfo.NickName, Avatar: userInfo.Avatar, ServerId: int32(userInfo.ServerId)}
}

func (this *ChallengeManager) GetChallengePeopleInfo1(userId, crossFsId int, season string) *pb.PeopleInfo {
	userInfo, err := modelCross.GetChallengeModel().GetAllServerApplyUserInfo(crossFsId, userId, season)
	if err != nil {
		logger.Error("GetAllServerApplyUserInfo crossFsId:%v, userId:%v err:%v", crossFsId, userId, err)
		return nil
	}
	return &pb.PeopleInfo{UserId: int32(userInfo.UserId), Name: userInfo.NickName, Avatar: userInfo.Avatar, ServerId: int32(userInfo.ServerId), Combat: userInfo.Combat, GuildName: userInfo.GuildName}
}

func (this *ChallengeManager) GetChallengePeopleInfo2(userId, crossFsId int, season string) *pb.PeopleInfo {
	userInfo := this.GetUserManager().GetUserBasicInfo(userId)
	if userInfo == nil {
		logger.Error("GetUserBasicInfo nil userId:%v", crossFsId, userId)
		return nil
	}
	return &pb.PeopleInfo{UserId: int32(userInfo.Id), Name: userInfo.NickName, Avatar: userInfo.Avatar, ServerId: int32(userInfo.ServerId), Combat: int64(userInfo.Combat)}
}

func (this *ChallengeManager) IsApplyChallenge(user *objs.User) int32 {
	err := this.JudgeIsInApplyTime(user.ServerId)
	if err != nil {
		logger.Error("IsApplyChallenge err:%v", err)
		return 0
	}
	crossFsId := this.GetSystem().GetServerIndexCrossFsId(user.ServerId)
	season := this.WeekByDate(time.Now().AddDate(0, 0, -1))
	userInfo, err := modelCross.GetChallengeModel().GetAllServerApplyUserInfo(crossFsId, user.Id, season)
	if err != nil {
		return 0
	}
	if userInfo == nil {
		return 0
	}
	return 1
}

func (this *ChallengeManager) UpDataChallengeInfo(state int, item *modelCross.Challenge) {
	// update db
	logger.Debug("UpDataChallengeInfo state:%v item:%v", state, item)
	this.opChan <- opMsg{state, item}
}

func (this *ChallengeManager) dbDataOp() {
	logger.Debug("challenge dbDataOp run")

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("ChallengeCcsManager dbDataOp panic: %v, time: %v, stack tace: %v\n", err, time.Now(), string(debug.Stack()))
		}
	}()

	for {
		select {
		case msg := <-this.opChan:
			switch msg.opType {
			case constAuction.OpInsert:
				//logger.Debug("challenge insert: %+v", *msg.challengeInfo)
				err := modelCross.GetChallengeModel().DbMap().Insert(msg.challengeInfo)
				if err != nil {
					logger.Error("insert challenge data: %v err: %v", *msg.challengeInfo, err)
				}
			case constAuction.OpUpdate:
				//logger.Debug("challenge update: %+v", *msg.challengeInfo)
				_, err := modelCross.GetChallengeModel().DbMap().Update(msg.challengeInfo)
				if err != nil {
					logger.Error("update challenge data: %v err: %v", *msg.challengeInfo, err)
				}
			default:
				logger.Error("Unknown db operation type: %v", msg.opType)
			}
		}
	}
}

//当前是一年的第几周
func (this *ChallengeManager) WeekByDate(t time.Time) string {
	yearDay := t.YearDay()
	yearFirstDay := t.AddDate(0, 0, -yearDay+1)
	firstDayInWeek := int(yearFirstDay.Weekday())

	//今年第一周有几天
	firstWeekDays := 1
	if firstDayInWeek != 0 {
		firstWeekDays = 7 - firstDayInWeek + 1
	}
	var week int
	if yearDay <= firstWeekDays {
		week = 1
	} else {
		week = (yearDay-firstWeekDays)/7 + 2
	}
	return fmt.Sprintf("%d", week)
}

func (this *ChallengeManager) GetBeforeSeason(lastSeason string) string {

	t := int(time.Now().Weekday())
	if t == 0 {
		t = 7
	}
	if t == 7 {
		return lastSeason
	}

	season, _ := strconv.Atoi(lastSeason)
	season -= 1
	return fmt.Sprintf("%v", season)
}

func (this *ChallengeManager) GetChallengeNowRound() int {

	round := -1

	if err := this.JudgeIsInApplyTime(base.Conf.ServerId); err != nil {
		return constChallenge.EIGHTH_ROUND + 1
	}

	openCfg := gamedb.GetCrossArenaTimeCrossArenaTimeCfg(1)
	if openCfg == nil {
		logger.Error("擂台赛 配置错误")
		return round
	}

	round = 0
	for num := constChallenge.Group; num <= constChallenge.EIGHTH_ROUND; num++ {
		if num == constChallenge.Group {
			t := int(time.Now().Weekday())
			if t == 0 {
				t = 7
			}
			if t == openCfg.SignUpEnd {
				if this.GetDaySecond(openCfg.SignUpEndTime.Hour, openCfg.SignUpEndTime.Minute, openCfg.SignUpEndTime.Second) {
					round = num + 1
				}
			}
		} else {
			cfg1 := gamedb.GetCrossArenaCrossArenaCfg(num)
			if cfg1 != nil {
				if this.GetDaySecond(cfg1.IntervalTime.Hour, cfg1.IntervalTime.Minute, cfg1.IntervalTime.Second) {
					round = num + 1
				}
			}
		}

	}
	return round
}

func (this *ChallengeManager) GetDaySecond(hour, min, sec int) bool {
	second1 := int(time.Now().Unix()) - common.GetZeroTimeUnix(0)
	second2 := hour*60*60 + min*60 + sec
	return second1 >= second2
}

func (this *ChallengeManager) Broadcast() {

	err := this.JudgeIsInApplyTime(base.Conf.ServerId)
	if err != nil {
		return
	}
	ack := &pb.ChallengeInfoAck{}
	_ = this.LoadInfo(base.Conf.ServerId, ack)
	this.BroadcastAll(ack)
}
