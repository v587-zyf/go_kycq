package challengeCcs

import (
	"cqserver/crosscenterserver/internal/managersI"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/publicCon/constAuction"
	"cqserver/gamelibs/publicCon/constChallenge"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pbserver"
	"fmt"
	"math/rand"
	"runtime/debug"
	"sort"
	"sync"
	"time"
)

type ChallengeCcsManager struct {
	util.DefaultModule
	managersI.IModule
	GroupsPeople        map[int][]*modelCross.Challenge         //key:crossFsId
	GroupsPeopleByRound map[int]map[int][]*modelCross.Challenge //key:crossFsId:round
	GroupsPeopleInfo    map[int]map[int]*modelCross.Challenge
	opChan              chan opMsg
	round               int //当前第几轮
	mu                  sync.RWMutex
}

func NewChallengeCcsManager(m managersI.IModule) *ChallengeCcsManager {
	return &ChallengeCcsManager{
		IModule:             m,
		GroupsPeople:        make(map[int][]*modelCross.Challenge, 0),
		GroupsPeopleByRound: make(map[int]map[int][]*modelCross.Challenge, 0),
		GroupsPeopleInfo:    make(map[int]map[int]*modelCross.Challenge, 0),
		opChan:              make(chan opMsg, 100),
		round:               1,
	}
}

type opMsg struct {
	opType        int
	challengeInfo *modelCross.Challenge
}

type PeopleInfo struct {
	UserId int
	Combat int64
}

type ChallengeInfos []*modelCross.Challenge

func (this ChallengeInfos) Len() int {
	return len(this)
}

func (this ChallengeInfos) Less(i, j int) bool {
	if this[i].Combat > this[j].Combat {
		return true
	}
	return false
}

func (this ChallengeInfos) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *ChallengeCcsManager) Init() error {

	go this.dbDataOp()
	nowTime := time.Now().Unix()
	_ = modelCross.GetChallengeModel().DeleteExpiredItem(nowTime)

	season := this.WeekByDate(time.Now().AddDate(0, 0, -1))
	this.GetGsServers().UpServerListTicker()
	groups := this.GetGsServers().GetAllCrossGroupServerInfo()
	logger.Info("init groups:%v", groups)
	isOpen := this.CheckIsOpen()
	if !isOpen {
		return nil
	}

	for crossFsId := range groups {
		if crossFsId <= 0 {
			continue
		}
		allUsers, err := modelCross.GetChallengeModel().GetAllServerApplyUsersList(crossFsId, season)
		logger.Info("challenge init crossFsId:%v  len:%v", crossFsId, len(allUsers))
		if err != nil {
			logger.Error("challenge init err:%v", err)
		}
		this.buildGroupsPeopleInfo(crossFsId)
		for _, userInfo := range allUsers {
			this.GroupsPeopleInfo[crossFsId][userInfo.UserId] = userInfo
		}

	}

	for crossFsId := range groups {
		maxRound, _ := modelCross.GetChallengeModel().GetChallengeNowRound(season, crossFsId)
		if maxRound > 0 {
			go this.buildInitInfo(crossFsId, maxRound, season)
		}
	}
	return nil
}

//
//  BeginGroup
//  @Description: 报名结束后开始分组
//  @receiver this
//  @param roundIndex  当前到了比赛第几轮
//
func (this *ChallengeCcsManager) BeginGroup(roundIndex int) {
	logger.Info("BeginGroup  roundIndex:%v  now:%v", roundIndex, time.Now())
	if !this.CheckIsOpen() {
		return
	}
	this.round = roundIndex
	this.Reset()
	this.InitGroupInfo(roundIndex + 1)
	return
}

//  InSeasonCompetition
//  @Description: 每一轮结束结算
//  @receiver this
//  @param roundIndex
//
func (this *ChallengeCcsManager) InSeasonCompetition(roundIndex int) {
	logger.Info("InSeasonCompetition  roundIndex:%v  now:%v", roundIndex, time.Now())
	if !this.CheckIsOpen() {
		return
	}
	this.round = roundIndex
	for crossFsId := range this.GroupsPeople {
		go this.buildSeasonCompetition(crossFsId)
	}
	return
}

//  InitGroupInfo
//  @Description: 分组前初始化组员信息 并且检查参赛人是否足够(不够需要插入假人)
//  @receiver this
//  @param round
//
func (this *ChallengeCcsManager) InitGroupInfo(round int) {

	season := this.WeekByDate(time.Now().AddDate(0, 0, -1))
	groups := this.GetGsServers().GetAllCrossGroupServerInfo()
	for crossFsId, serveIds := range groups {
		err := modelCross.GetChallengeModel().UpdateRound(round, crossFsId, season)
		if err != nil {
			logger.Error("GetChallengeModel().UpdateRound(round) err:%v", err)
		}
		if len(serveIds) <= 0 {
			continue
		}
		go this.buildInsertRobotUser(round, crossFsId, season, serveIds)
	}

	return
}

//  RandIndex
//  @Description: 随机给玩家排序
//  @receiver this
//  @param infos
//  @param crossFsId
//  @return []*modelCross.Challenge
//
func (this *ChallengeCcsManager) RandIndex(infos []*modelCross.Challenge, crossFsId int) []*modelCross.Challenge {
	if len(infos) < 2 {
		return infos
	}

	rand.Seed(time.Now().Unix())
	//采用rand.Shuffle&#xff0c;将切片随机化处理后返回
	rand.Shuffle(len(infos), func(i, j int) { infos[i], infos[j] = infos[j], infos[i] })

	specialInfo := this.SpecialUserDispose(infos, crossFsId)
	return specialInfo
}

//  SpecialUserDispose
//  @Description: 特殊玩家处理(A、B、C、D四个服是一个跨服组，活动报名截至时，每个服报名战力最高的那2个人，在预选赛、突围赛分组的时候，不要分在一个组)
//  @receiver this
//  @param infos
//  @param crossFsId
//  @return []*modelCross.Challenge
func (this *ChallengeCcsManager) SpecialUserDispose(infos []*modelCross.Challenge, crossFsId int) []*modelCross.Challenge {
	logger.Info("SpecialUserDispose  len(infos):%v ", len(infos))
	if len(infos) < 2 {
		logger.Debug("SpecialUserDispose len(infos):%v < 2", len(infos))
		return infos
	}

	beforePeoples := make(ChallengeInfos, 0)
	beforePeoples = infos
	sort.Sort(beforePeoples)

	towerUserIdsByGroup := make([]*modelCross.Challenge, 0)
	if len(beforePeoples) > 8 {
		towerUserIdsByGroup = beforePeoples[0:8]
	} else {
		logger.Error("beforePeoples:%v", len(beforePeoples))
		return infos
	}

	season := this.WeekByDate(time.Now().AddDate(0, 0, -1))
	if towerUserIdsByGroup == nil || len(towerUserIdsByGroup) <= 0 {
		logger.Debug("GetServerApplyUsersListByCombat season:%v ", season)
		return infos
	}
	if len(towerUserIdsByGroup) < 8 {
		return infos
	}

	userIndex := make(map[int]bool, 0)
	for _, user := range towerUserIdsByGroup {
		if user == nil || user.IsLose >= 1 {
			continue
		}
		userIndex[user.UserId] = true
	}

	//先del 每个服战力前2的玩家
	logger.Debug(" lenInfos:%v   lenTowerUserIdsByGroup:%v  userIndex:%v", len(infos), len(towerUserIdsByGroup), userIndex)
	markInfos := make([]*modelCross.Challenge, 0)
	for _, info := range infos {
		if userIndex[info.UserId] {
			continue
		}
		markInfos = append(markInfos, info)
	}
	infos = markInfos
	logger.Debug(" lenInfos:%v", len(infos))

	markIndex := 6
	for _, userInfo := range towerUserIdsByGroup {
		if markIndex >= len(infos) {
			break
		}
		markUserInfo := make([]*modelCross.Challenge, 0)
		markUserInfo = append(markUserInfo, infos[:markIndex]...)
		markUserInfo = append(markUserInfo, userInfo)
		markUserInfo = append(markUserInfo, infos[markIndex:]...)
		infos = markUserInfo
		markIndex += 6
	}

	logger.Debug("lenInfos:%v", len(infos))
	return infos

}

//  OneRoundEnd
//  @Description:  每轮结算 and  发送失败奖励
//  @receiver this
//  @param infos
//  @param crossFsId
//  @param roundIndex
//
func (this *ChallengeCcsManager) OneRoundEnd(infos []*modelCross.Challenge, crossFsId, roundIndex int) {
	winUserId := 0
	winServerId := 0
	loseUsers := make(map[int][]int32, 0) //key:serverId
	winUsers := make(map[int][]int32, 0)  //key:serverId
	logger.Info("OneRoundEnd len(infos):%v   roundIndex:%v   crossFsId:%v", len(infos), roundIndex, crossFsId)

	for i, j := 0, len(infos); i < j; i += 2 {
		if this.GroupsPeopleByRound[crossFsId] == nil {
			this.GroupsPeopleByRound[crossFsId] = make(map[int][]*modelCross.Challenge)
		}
		if this.GroupsPeopleByRound[crossFsId][this.round+1] == nil {
			this.GroupsPeopleByRound[crossFsId][this.round+1] = make([]*modelCross.Challenge, 0)
		}

		if i > len(infos)-1 || i+1 > len(infos)-1 {
			logger.Error("分组错误 人数超了 有bug")
			continue
		}
		if infos[i] == nil || infos[i+1] == nil {
			continue
		}
		if infos[i].Combat >= infos[i+1].Combat {
			infos[i+1].IsLose = 1
			infos[i+1].LoseRound = this.round
			infos[i+1].WinUserId = infos[i].UserId
			this.UpDataChallengeInfo(constAuction.OpUpdate, infos[i+1])
			this.setGroupPeopleInfo(crossFsId, infos[i+1].UserId, infos[i+1])
			infos[i].Round = this.round + 1

			this.UpDataChallengeInfo(constAuction.OpUpdate, infos[i])
			this.setGroupPeopleInfo(crossFsId, infos[i].UserId, infos[i])
			this.GroupsPeopleByRound[crossFsId][this.round+1] = append(this.GroupsPeopleByRound[crossFsId][this.round+1], infos[i])
			if loseUsers[infos[i+1].ServerId] == nil {
				loseUsers[infos[i+1].ServerId] = make([]int32, 0)
			}
			loseUsers[infos[i+1].ServerId] = append(loseUsers[infos[i+1].ServerId], int32(infos[i+1].UserId))

			if winUsers[infos[i].ServerId] == nil {
				winUsers[infos[i].ServerId] = make([]int32, 0)
			}
			winUsers[infos[i].ServerId] = append(winUsers[infos[i].ServerId], int32(infos[i].UserId))

			if roundIndex == constChallenge.EIGHTH_ROUND {
				winUserId = infos[i].UserId
				winServerId = infos[i].ServerId
			}

		} else {
			infos[i].IsLose = 1
			infos[i].LoseRound = this.round
			infos[i].WinUserId = infos[i+1].UserId
			this.UpDataChallengeInfo(constAuction.OpUpdate, infos[i])
			this.setGroupPeopleInfo(crossFsId, infos[i].UserId, infos[i])
			infos[i+1].Round = this.round + 1
			this.UpDataChallengeInfo(constAuction.OpUpdate, infos[i+1])
			this.setGroupPeopleInfo(crossFsId, infos[i+1].UserId, infos[i+1])
			this.GroupsPeopleByRound[crossFsId][this.round+1] = append(this.GroupsPeopleByRound[crossFsId][this.round+1], infos[i+1])
			if loseUsers[infos[i].ServerId] == nil {
				loseUsers[infos[i].ServerId] = make([]int32, 0)
			}
			loseUsers[infos[i].ServerId] = append(loseUsers[infos[i].ServerId], int32(infos[i].UserId))

			if winUsers[infos[i+1].ServerId] == nil {
				winUsers[infos[i+1].ServerId] = make([]int32, 0)
			}
			winUsers[infos[i+1].ServerId] = append(winUsers[infos[i+1].ServerId], int32(infos[i+1].UserId))

			if roundIndex == constChallenge.EIGHTH_ROUND {
				winUserId = infos[i+1].UserId
				winServerId = infos[i+1].ServerId
			}
		}
	}

	this.SendToGsLoseReward(loseUsers, winUsers, int32(winUserId), int32(winServerId), crossFsId)
}

//发送每一轮失败者奖励
func (this *ChallengeCcsManager) SendToGsLoseReward(loseInfos, winInfos map[int][]int32, winUserId, winServerId int32, crossFsId int) {

	allWinUserIds := make([]int32, 0)
	for _, data := range winInfos {
		allWinUserIds = append(allWinUserIds, data...)
	}
	serverIds := this.GetGsServers().GetCrossGroupServerInfoByCrossFsId(crossFsId)
	if serverIds != nil {
		for _, serverInfo := range serverIds {
			if serverInfo == nil {
				continue
			}
			winUserId2 := int32(0)
			serverId := serverInfo.ServerId
			if serverId == int(winServerId) {
				winUserId2 = winUserId
			}
			logger.Info("SendToGsLoseReward serverId:%v,  winUserId:%v  winServerId:%v  loseInfos[serverId]:%v   allWinUserIds:%v", serverId, winUserId, winServerId, loseInfos[serverId], allWinUserIds)
			_ = this.GetGsServers().SendMessage(serverId, &pbserver.ChallengeSendLoseRewardNtf{RoundIndex: int32(this.round), LoseUers: loseInfos[serverId], WinUserId: winUserId2, WinUsers: allWinUserIds})
		}
	}

}


func (this *ChallengeCcsManager) ToGsUp(crossFsId int) {

	datas := this.GetGsServers().GetCrossGroupServerInfoByCrossFsId(crossFsId)
	for _, data := range datas {
		_ = this.GetGsServers().SendMessage(data.ServerId, &pbserver.ChallengeAppuserUpToGsNtf{})
	}

}

//提前插入擂台赛假人
func (this *ChallengeCcsManager) BeginInsertRobotUser() {
	cfg := gamedb.GetCrossArenaTimeCrossArenaTimeCfg(1)

	t := int(time.Now().Weekday())
	if t == 0 {
		t = 7
	}

	if t != cfg.SignUpEnd {
		return
	}

	logger.Info("提前插入擂台赛假人")
	this.beforeInsertRobot(constChallenge.Group)
}

func (this *ChallengeCcsManager) setRoundPeoples(infos []*modelCross.Challenge, crossFightId, roundIndex int) {
	season := this.WeekByDate(time.Now().AddDate(0, 0, -1))
	users := make([]int, 0)
	for _, v := range infos {
		users = append(users, v.UserId)
	}
	info := &modelCross.ChallengeData{}
	info.CrossFsId = crossFightId
	info.Season = season
	info.Round = roundIndex
	info.UserIds = users
	info.ExpireTime = time.Now().Unix() + 86400*14
	err := modelCross.GetChallengeDataModel().DbMap().Insert(info)
	if err != nil {
		logger.Error("setRoundPeoples err:%v  crossFightId:%v, roundIndex:%v", err, crossFightId, roundIndex)
	}
}

func (this *ChallengeCcsManager) getGroupPeopleInfo(crossFightId, userId int) *modelCross.Challenge {
	this.mu.RLock()
	defer func() {
		this.mu.RUnlock()
	}()

	return this.GroupsPeopleInfo[crossFightId][userId]
}

func (this *ChallengeCcsManager) setGroupPeopleInfo(crossFightId, userId int, info *modelCross.Challenge) {
	this.mu.RLock()
	defer func() {
		this.mu.RUnlock()
	}()

	this.GroupsPeopleInfo[crossFightId][userId] = info
}

//活动开启判断
func (this *ChallengeCcsManager) CheckIsOpen() bool {
	cfg := gamedb.GetCrossArenaTimeCrossArenaTimeCfg(1)

	t := int(time.Now().Weekday())
	if t == 0 {
		t = 7
	}

	if t == cfg.SignUpEnd {
		if time.Now().Hour() > cfg.SignUpEndTime.Hour {
			return true
		}
		if time.Now().Hour() == cfg.SignUpEndTime.Hour {
			if time.Now().Minute() >= cfg.SignUpEndTime.Minute {
				return true
			}
		}

	}
	return false
}

func (this *ChallengeCcsManager) Reset() {
	this.GroupsPeople = make(map[int][]*modelCross.Challenge, 0)
	this.GroupsPeopleByRound = make(map[int]map[int][]*modelCross.Challenge, 0)
	this.GroupsPeopleInfo = make(map[int]map[int]*modelCross.Challenge, 0)
}

//取活动到了第几赛季
func (this *ChallengeCcsManager) WeekByDate(t time.Time) string {
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

func (this *ChallengeCcsManager) beforeInsertRobot(round int) {

	season := this.WeekByDate(time.Now().AddDate(0, 0, -1))
	groups := this.GetGsServers().GetAllCrossGroupServerInfo()
	logger.Info("擂台赛 提前插入假人 begin season:%v  groups:%v", season, groups)
	for crossFsId, serveIds := range groups {
		if len(serveIds) <= 0 {
			continue
		}
		go this.buildInsertRobotUser(round, crossFsId, season, serveIds)
	}
	logger.Info("提前插入假人end")
	return
}

//插入擂台赛假人
func (this *ChallengeCcsManager) buildInsertRobotUser(round, crossFsId int, season string, serveIds map[int]*modelCross.ServerInfo) {

	robotId := 1
	maxRobotId, err := modelCross.GetChallengeModel().MaxRobotId(round, crossFsId, season)
	if err == nil {
		if maxRobotId < 0 {
			robotId = -maxRobotId + 1
		}
	}

	allUsersInfo := make([]*modelCross.Challenge, 0)
	allServeIds := make([]int, 0)
	for _, serveInfo := range serveIds {
		serveId := serveInfo.ServerId
		allServeIds = append(allServeIds, serveId)
	}
	allUsers, _ := modelCross.GetChallengeModel().GetAllServerApplyUsersList(crossFsId, season)
	if len(allUsers) > 0 {
		allUsersInfo = append(allUsersInfo, allUsers...)
	}
	logger.Debug("crossFsId:%v  crossFsId:%v  len(allUsers):%v  len(allUsersInfo):%v", crossFsId, len(allUsers), len(allUsersInfo))
	for len(allUsersInfo) < 256 {
		rand.Seed(int64(robotId))
		randIndex := rand.Intn(len(allServeIds))
		serveId := allServeIds[randIndex]
		logger.Debug("补充假人 robotId:%v  crossFsId:%v serveId:%v", robotId, crossFsId, serveId)
		//补充假人
		robotCfg := gamedb.GetCrossArenaRobotCrossArenaRobotCfg(robotId)
		if robotCfg == nil {
			logger.Error("GetCrossArenaRobotCrossArenaRobotCfg robotId:%v  nil 假人不足", robotId)
			break
		}
		info := &modelCross.Challenge{UserId: -robotCfg.Id, Season: season, NickName: robotCfg.Name, Avatar: robotCfg.Icon, ServerId: serveId, Combat: int64(robotCfg.Combat), ExpireTime: time.Now().Unix() + 86400*14, Round: round, CrossFsId: crossFsId}
		allUsersInfo = append(allUsersInfo, info)
		this.UpDataChallengeInfo(constAuction.OpInsert, info)
		robotId++
	}

	if round > 0 {
		this.GroupsPeople[crossFsId] = make([]*modelCross.Challenge, 0)
		this.buildGroupsPeopleInfo(crossFsId)
		needDel := len(allUsersInfo) - 256
		for _, userInfo := range allUsersInfo {
			if userInfo.UserId < 0 && needDel > 0 {
				needDel--
				continue
			}
			this.GroupsPeople[crossFsId] = append(this.GroupsPeople[crossFsId], userInfo)
			this.GroupsPeopleInfo[crossFsId][userInfo.UserId] = userInfo
		}
		go this.buildSetRoundPeoples(crossFsId, this.GroupsPeople[crossFsId])
		logger.Debug("InitGroupInfo len:%v", len(this.GroupsPeople[crossFsId]))
	}
	return
}

//赛季比拼
func (this *ChallengeCcsManager) buildSeasonCompetition(crossFsId int) {

	peoples := this.GroupsPeopleByRound[crossFsId][this.round]
	logger.Debug("InSeasonCompetition crossFsId:%v  len:%v  this.round:%v", crossFsId, len(peoples), this.round)
	this.OneRoundEnd(peoples, crossFsId, this.round)
	peoples = this.GroupsPeopleByRound[crossFsId][this.round+1]
	if this.round == constChallenge.SECOND_ROUND {
		//第二轮结算后需要重新分组
		peoples = this.RandIndex(peoples, crossFsId)
		this.buildGroupsPeopleByRoundInit(crossFsId, this.round+1)
		this.GroupsPeopleByRound[crossFsId][this.round+1] = peoples
	}
	this.setRoundPeoples(peoples, crossFsId, this.round+1)
	return
}

//第一次分组  随机排序玩家数据
func (this *ChallengeCcsManager) buildSetRoundPeoples(crossFsId int, peoples []*modelCross.Challenge) {

	logger.Debug(" BeginGroup  crossFsId:%v  len:%v", crossFsId, len(peoples))
	peoples = this.RandIndex(peoples, crossFsId)
	logger.Debug(" BeginGroup RandIndex after  crossFsId:%v  len:%v  this.round:%v", crossFsId, len(peoples), this.round)
	if this.round == constChallenge.Group {
		if this.GroupsPeopleByRound[crossFsId] == nil {
			this.GroupsPeopleByRound[crossFsId] = make(map[int][]*modelCross.Challenge)
		}
		if this.GroupsPeopleByRound[crossFsId][constChallenge.FIRST_ROUND] == nil {
			this.GroupsPeopleByRound[crossFsId][constChallenge.FIRST_ROUND] = make([]*modelCross.Challenge, 0)
		}
		this.GroupsPeopleByRound[crossFsId][constChallenge.FIRST_ROUND] = peoples

	}
	this.setRoundPeoples(peoples, crossFsId, constChallenge.FIRST_ROUND)
	return
}

func (this *ChallengeCcsManager) buildInitInfo(crossFsId, maxRound int, season string) {

	if this.GroupsPeople[crossFsId] == nil {
		this.GroupsPeople[crossFsId] = make([]*modelCross.Challenge, 0)
	}
	for roundIndex := constChallenge.FIRST_ROUND; roundIndex <= maxRound; roundIndex++ {
		data, _ := modelCross.GetChallengeDataModel().GetRoundUserIdsByCrossFsId(crossFsId, roundIndex, season)
		if data == nil || data.UserIds == nil {
			continue
		}
		logger.Info("init  roundIndex:%v  crossFsId:%v  season:%v  lendata.UserIds:%v", roundIndex, crossFsId, season, len(data.UserIds))
		for _, userId := range data.UserIds {
			info := this.getGroupPeopleInfo(crossFsId, userId)
			if info == nil {
				continue
			}

			if roundIndex == constChallenge.FIRST_ROUND {
				this.GroupsPeople[crossFsId] = append(this.GroupsPeople[crossFsId], info)
			}

			this.buildGroupsPeopleByRoundInit(crossFsId, roundIndex)
			this.GroupsPeopleByRound[crossFsId][roundIndex] = append(this.GroupsPeopleByRound[crossFsId][roundIndex], info)
		}
	}

}

func (this *ChallengeCcsManager) buildGroupsPeopleByRoundInit(crossFsId, roundIndex int) {

	if this.GroupsPeopleByRound == nil {
		this.GroupsPeopleByRound = make(map[int]map[int][]*modelCross.Challenge)
	}
	if this.GroupsPeopleByRound[crossFsId] == nil {
		this.GroupsPeopleByRound[crossFsId] = make(map[int][]*modelCross.Challenge)
	}
	if this.GroupsPeopleByRound[crossFsId][roundIndex] == nil {
		this.GroupsPeopleByRound[crossFsId][roundIndex] = make([]*modelCross.Challenge, 0)
	}
	return
}

func (this *ChallengeCcsManager) buildGroupsPeopleInfo(crossFsId int) {

	if this.GroupsPeopleInfo[crossFsId] == nil {
		this.GroupsPeopleInfo[crossFsId] = make(map[int]*modelCross.Challenge, 0)
	}
	return
}

func (this *ChallengeCcsManager) dbDataOp() {
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

func (this *ChallengeCcsManager) UpDataChallengeInfo(state int, item *modelCross.Challenge) {
	// update db
	this.opChan <- opMsg{state, item}
}
