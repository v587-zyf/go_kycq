package fieldFight

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constConstant"
	"cqserver/gamelibs/publicCon/constField"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"database/sql"
	"encoding/json"
	"time"
)

func NewFieldFightManager(m managersI.IModule) *FieldManager {
	field := &FieldManager{}
	field.IModule = m
	return field
}

type FieldManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *FieldManager) Init() error {

	this.CheckUser()
	return nil

}

func (this *FieldManager) LoadInfo(user *objs.User, ack *pb.FieldFightLoadAck) error {

	fieldFightTimesCfg, _, openDay, err := this.CheckGameCfg()
	if err != nil {
		return err
	}

	this.GetFieldFight().JudgeIsSetUserFiveAmCombat(user)
	haveChallengeTimes := user.FieldFight.HaveChallengeTimes
	lastChallengeTimes := this.GetChallengeTimes(fieldFightTimesCfg[0]+this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_FIELDFIGHT_FREENUM), haveChallengeTimes)
	haveBuyTimes := user.FieldFight.HaveBuyTimes
	lastCanBuyTimes := this.GetChallengeTimes(fieldFightTimesCfg[1]+this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_FIELDFIGHT_BUYNUM), haveBuyTimes)
	canRefreshTime := rmodel.FieldFight.GetFieldFightChangeRivalCd(user.Id, openDay)
	BeatBackUserInfo := this.GetRivalUserInfos(user.Id, openDay)

	ack.ListInfo = this.GetRefRivalsUsersListInfo(user, openDay)
	ack.BeatBackOwnUserInfo = BeatBackUserInfo
	ack.ChangeRivalCd = int32(canRefreshTime)            //可以刷新对手的时间戳
	ack.RemainChallengeTimes = int32(lastChallengeTimes) //今天剩余可挑战次数
	ack.TodayCanBuyTimes = int32(lastCanBuyTimes)        //今天剩余可购买次数
	ack.MyCombat = int32(user.Combat)

	return nil
}

//购买挑战次数
func (this *FieldManager) BuyFieldFightChallengeNum(user *objs.User, op *ophelper.OpBagHelperDefault, ack *pb.BuyFieldFightChallengeTimesAck) error {

	fieldFightTimesCfg, fieldFightCostCfg, _, err := this.CheckGameCfg()
	if err != nil {
		return err
	}

	haveBuyTimes := user.FieldFight.HaveBuyTimes
	lastCanBuyTimes := this.GetChallengeTimes(fieldFightTimesCfg[1]+this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_FIELDFIGHT_BUYNUM), haveBuyTimes)
	if lastCanBuyTimes <= 0 {
		return gamedb.ERRPURCHASECAPENOUGH
	}
	itemId := 0
	count := 0
	for _, v := range fieldFightCostCfg {
		if ok, _ := this.GetBag().HasEnough(user, v.ItemId, v.Count); !ok {
			continue
		}
		itemId = v.ItemId
		count = v.Count
		break
	}

	if itemId == 0 || count == 0 {
		return gamedb.ERRNOTENOUGHGOODS
	}

	err = this.GetBag().Remove(user, op, itemId, count)
	if err != nil {
		return err
	}

	user.FieldFight.HaveBuyTimes += 1
	user.FieldFight.HaveChallengeTimes -= 1
	user.Dirty = true
	ack.ResidueTimes = int32(this.GetChallengeTimes(fieldFightTimesCfg[0]+this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_FIELDFIGHT_FREENUM), user.FieldFight.HaveChallengeTimes)) //剩余挑战次数
	ack.TodayCanBuyTimes = int32(this.GetChallengeTimes(fieldFightTimesCfg[1]+this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_FIELDFIGHT_BUYNUM), user.FieldFight.HaveBuyTimes))    //今日可购买次数
	return nil
}

//刷新劲敌
func (this *FieldManager) RivalUser(user *objs.User, ack *pb.RefFieldFightRivalUserAck, isLoginCheck bool) error {

	openDay := this.GetOpenDayByReduceTime()
	cd := rmodel.FieldFight.GetFieldFightChangeRivalCd(user.Id, openDay)
	if int(time.Now().Unix()) < cd {
		return gamedb.ERRREFCD
	}

	baseInfo, err := this.GetRivalUserNumAndCombat(openDay, user.Id)
	if err != nil {
		return err
	}

	//刷新劲敌
	allNightMareUserIds, allDifficultyUserIds, allSimpleUserIds := this.GetAllDifficultUsers(user, baseInfo)
	allUserIds := make([]*pb.FieldFightRivalUserInfo, 0)
	allMatchUserIds := make([]int, 0)
	allMatchUserIds = append(allMatchUserIds, allNightMareUserIds...)
	allMatchUserIds = append(allMatchUserIds, allDifficultyUserIds...)
	allMatchUserIds = append(allMatchUserIds, allSimpleUserIds...)
	if !isLoginCheck {
		rmodel.FieldFight.SetFieldFightChangeRivalCd(user.Id, openDay, int(time.Now().Unix())+gamedb.GetConf().FieldFightCd)
	}
	ack.ChangeRivalCd = int32(rmodel.FieldFight.GetFieldFightChangeRivalCd(user.Id, openDay))
	allUserIds, ack.ListInfo = this.SoreUserInfoByCombat(user, allMatchUserIds)

	bytes, _ := json.Marshal(allUserIds)
	rmodel.FieldFight.SetFieldFightSaveBeforeRefRivals(user.Id, string(bytes))
	return nil

}

//每日重置玩家数据
func (this *FieldManager) DayReset(user *objs.User, isReset bool) {
	data := common.GetResetTime(time.Now())
	logger.Debug("DayReset userId:%v  data:%v   user.FieldFight.DayResDay:%v", user.Id, data, user.FieldFight.DayResDay)
	if data != user.FieldFight.DayResDay {
		user.FieldFight.HaveChallengeTimes = 0
		user.FieldFight.HaveBuyTimes = 0
		user.FieldFight.DayResDay = data
	}
	user.Dirty = true
	if isReset {
		ack := &pb.FieldFightLoadAck{}
		err := this.LoadInfo(user, ack)
		if err == nil {
			_ = this.GetUserManager().SendMessage(user, ack, true)
		}
	}
	return
}

func (this *FieldManager) GetAllDifficultUsers(user *objs.User, baseInfo map[int][]int) ([]int, []int, []int) {

	allNightMareUserIds := make([]int, 0)
	allDifficultyUserIds := make([]int, 0)
	allSimpleUserIds := make([]int, 0)
	matchUserIds := make([]int, 0)

	haveAppearUser, haveAppearRobot := this.buildFieldHaveAppearUserInfo(user.Id)

	for difficult, data := range baseInfo {
		index := difficult - 1
		matchUserIds, haveAppearUser, haveAppearRobot = this.MatchRivalPeople(user, difficult, data[0]/2, data[1], haveAppearUser, haveAppearRobot)
		if index == constField.NightMare {
			allNightMareUserIds = matchUserIds
		} else if index == constField.Difficulty {
			allDifficultyUserIds = matchUserIds
		} else if index == constField.Simple {
			allSimpleUserIds = matchUserIds
		}
	}

	for _, v := range allNightMareUserIds {
		logger.Debug(" 野战 噩梦玩家列表 allNightMareUserIds: userId:%v", v)
	}
	for _, v := range allDifficultyUserIds {
		logger.Debug(" 野战 困难玩家列表 allDifficultyUserIds: userId:%v", v)
	}
	for _, v := range allSimpleUserIds {
		logger.Debug(" 野战 简单玩家列表 allSimpleUserIds: userId:%v", v)
	}
	return allNightMareUserIds, allDifficultyUserIds, allSimpleUserIds
}

//匹配劲敌玩家
func (this *FieldManager) MatchRivalPeople(user *objs.User, difficulty, needNum, needCombat int, haveAppearUsers []int, haveAppearRobot map[int]bool) ([]int, []int, map[int]bool) {
	logger.Debug("匹配劲敌玩家 userId:%v difficulty:%v, needNum:%v, needCombat:%v ", user.Id, difficulty, needNum, needCombat)
	logger.Debug("haveAppearUsers:%v  haveAppearRobot:%v", haveAppearUsers, haveAppearRobot)
	matchUserInfos := make([]int, 0)
	rivalUserInfos := make([]int, 0)
	isHaveUp := false
	//先向下取
	rivalUserInfosBefore, err := rmodel.Rank.GetRankBuyTypeAndScoreSection(pb.RANKTYPE_COMBAT, user.ServerId, -1, int(needCombat))
	//过滤上次列表劲敌
	rivalUserInfos, haveAppearRobot = this.ScreenOutRivalUserInfo1(user, rivalUserInfosBefore, haveAppearUsers, haveAppearRobot)
	if err != nil || rivalUserInfos == nil || len(rivalUserInfos) <= 0 {
		//向上取
		logger.Debug("野战 刷新劲敌 找不到比自己战力低的 那去找比自己战力高的 err:%v rivalUserInfos:%v  rivalUserInfosBefore:%v", err, rivalUserInfos, rivalUserInfosBefore)
		rivalUserInfosBefore, err = rmodel.Rank.GetRankBuyTypeAndScoreSection(pb.RANKTYPE_COMBAT, user.ServerId, int(needCombat), "+inf")
		rivalUserInfos, haveAppearRobot = this.ScreenOutRivalUserInfo1(user, rivalUserInfosBefore, haveAppearUsers, haveAppearRobot)
		isHaveUp = true
		logger.Debug("after rivalUserInfos:%v  haveAppearRobot:%v", rivalUserInfos, haveAppearRobot)
		if err != nil || rivalUserInfos == nil || len(rivalUserInfos) <= 0 {
			//匹配假人
			matchUserInfos, haveAppearRobot = this.matchRobotUser(difficulty, needNum, needCombat, matchUserInfos, haveAppearRobot)
			return matchUserInfos, haveAppearUsers, haveAppearRobot
		}
	}

	lenRivalUserInfos := len(rivalUserInfos)
	realNeedNumber := lenRivalUserInfos - needNum
	logger.Debug("len(rivalUserInfos):%v  rivalUserInfos:%v    realNeedNumber:%v  isHaveUp:%v", lenRivalUserInfos, rivalUserInfos, realNeedNumber, isHaveUp)
	if realNeedNumber < 0 {
		matchUserInfos = this.buildMatchUserInfo(difficulty, matchUserInfos, rivalUserInfos)
		haveAppearUsers = append(haveAppearUsers, matchUserInfos...)
		//剩余需要匹配的人数
		lastNeedMatchNum := needNum - lenRivalUserInfos
		if isHaveUp {
			//匹配假人
			matchUserInfos, haveAppearRobot = this.matchRobotUser(difficulty, lastNeedMatchNum, needCombat, matchUserInfos, haveAppearRobot)
			return matchUserInfos, haveAppearUsers, haveAppearRobot
		}

		//向上取
		rivalUserInfosBefore, err = rmodel.Rank.GetRankBuyTypeAndScoreSection(pb.RANKTYPE_COMBAT, user.ServerId, needCombat, "+inf")
		rivalUserInfos, haveAppearRobot = this.ScreenOutRivalUserInfo1(user, rivalUserInfosBefore, haveAppearUsers, haveAppearRobot)

		logger.Debug("向上取出 %v 人  rivalUserInfos:%v   lastNeedMatchNum:%v", len(rivalUserInfos), rivalUserInfos, lastNeedMatchNum)
		realNeedNumber = len(rivalUserInfos) - lastNeedMatchNum

		if realNeedNumber < 0 {

			matchUserInfos = this.buildMatchUserInfo(difficulty, matchUserInfos, rivalUserInfos)
			haveAppearUsers = append(haveAppearUsers, matchUserInfos...)

			lastNeedRobotNum := lastNeedMatchNum - len(rivalUserInfos)
			//补充假人
			logger.Debug("向上取出来的人不够  需要补充 %v 个假人", lastNeedRobotNum)
			matchUserInfos, haveAppearRobot = this.matchRobotUser(difficulty, lastNeedRobotNum, needCombat, matchUserInfos, haveAppearRobot)
			return matchUserInfos, haveAppearUsers, haveAppearRobot

		} else {
			rivalUserInfos = rivalUserInfos[realNeedNumber:]
			matchUserInfos = this.buildMatchUserInfo(difficulty, matchUserInfos, rivalUserInfos)
			haveAppearUsers = append(haveAppearUsers, matchUserInfos...)
			return matchUserInfos, haveAppearUsers, haveAppearRobot
		}

	} else {
		rivalUserInfos = rivalUserInfos[lenRivalUserInfos-needNum:]
		matchUserInfos = this.buildMatchUserInfo(difficulty, matchUserInfos, rivalUserInfos)
		haveAppearUsers = append(haveAppearUsers, matchUserInfos...)
		return matchUserInfos, haveAppearUsers, haveAppearRobot
	}
}

//补足相应困难度玩家
func (this *FieldManager) ComplementPeopleByDifficult(user *objs.User, openDay, difficult, needNum, loseUserId int, allRivalUserInfo []*pb.FieldFightRivalUserInfo) {
	logger.Debug("补足相应困难度玩家 userId:%v 困难度:%v needNum:%v, loseUserId:%v", user.Id, difficult, needNum, loseUserId)
	haveAppearRobot := make(map[int]bool)
	haveAppearUsers := make([]int, 0)
	matchUserIds := make([]int, 0)
	for _, data := range allRivalUserInfo {
		if data.RivalUserId < 0 {
			haveAppearRobot[int(-data.RivalUserId)] = true
		}
		haveAppearUsers = append(haveAppearUsers, int(data.RivalUserId))
		if int(data.RivalUserId) == loseUserId {
			continue
		}
		matchUserIds = append(matchUserIds, int(data.RivalUserId))
	}
	fiveCombat := rmodel.FieldFight.GetFieldFightMatchUserFiveAmCombat(openDay, user.Id)
	fightUserCfg := gamedb.GetConf().FieldFightLevel
	rate := float32(1)
	for index, data := range fightUserCfg {
		if index == difficult-1 {
			rate = float32(data[0]) / constConstant.WAN_FEN_BI
		}
	}
	needCombat := int(float32(fiveCombat) * rate)

	addMatchUserId, _, _ := this.MatchRivalPeople(user, difficult, needNum, needCombat, haveAppearUsers, haveAppearRobot)
	logger.Debug("补玩家 addMatchUserId:%v", addMatchUserId)
	matchUserIds = append(matchUserIds, addMatchUserId...)
	logger.Debug("after 补玩家 matchUserIds:%v", matchUserIds)

	allUserIds := make([]*pb.FieldFightRivalUserInfo, 0)
	allUserIds, _ = this.SoreUserInfoByCombat(user, matchUserIds)

	bytes, _ := json.Marshal(allUserIds)
	rmodel.FieldFight.SetFieldFightSaveBeforeRefRivals(user.Id, string(bytes))
	return

}

//
//  GetOpenDayByReduceTime
//  @Description: 获取偏移时间 获得的开服天数
//
func (this *FieldManager) GetOpenDayByReduceTime() int {

	return this.GetSystem().GetServerOpenDaysByServerIdByExcursionTime(base.Conf.ServerId, 0)

}

//购买挑战次数
func (this *FieldManager) BuyFieldFightChallengeNumByBag(user *objs.User) error {

	fieldFightTimesCfg, _, _, err := this.CheckGameCfg()
	if err != nil {
		return err
	}

	ack := pb.BuyFieldFightChallengeTimesAck{}

	user.FieldFight.HaveBuyTimes += 1
	user.FieldFight.HaveChallengeTimes -= 1

	ack.ResidueTimes = int32(this.GetChallengeTimes(fieldFightTimesCfg[0]+this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_FIELDFIGHT_FREENUM), user.FieldFight.HaveChallengeTimes)) //剩余挑战次数
	ack.TodayCanBuyTimes = int32(this.GetChallengeTimes(fieldFightTimesCfg[1]+this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_FIELDFIGHT_BUYNUM), user.FieldFight.HaveBuyTimes))    //今日可购买次数
	return nil
}

//定时5点存储活跃玩家战力 用于匹配对手
func (this *FieldManager) FiveAmRecordUserCombat() {
	//7天活跃的玩家
	allUsers, err := modelGame.GetUserModel().LoadAllUsersCombat(604800)
	if err != nil && err != sql.ErrNoRows {
		logger.Error("FiveAmRecordUserCombat err:%v", err)
		return
	}
	userCombatInfo := make(map[int]int)
	for _, userInfo := range allUsers {
		userCombatInfo[userInfo.Id] = userInfo.Combat
	}
	openDay := this.GetFieldFight().GetOpenDayByReduceTime()
	logger.Debug("FiveAmRecordUserCombat openDay:%v  userCombatInfo:%v", openDay, userCombatInfo)
	rmodel.FieldFight.SetFieldFightMatchUserFiveAmCombat(openDay, userCombatInfo)
}

//长时间不活跃玩家 上线时存储他的当前战力 用于匹配野战玩家 and 判断是否需要刷新对手
func (this *FieldManager) JudgeIsSetUserFiveAmCombat(user *objs.User) {
	if !this.checkIconIsOpen(user) {
		return
	}
	isRival := this.setUserMatchCombat(user)
	//判断是否需要刷新对手
	state, _ := rmodel.FieldFight.GetIsHaveFieldFightSaveBeforeRefRivals(user.Id)
	if state == 0 || isRival {
		ack := &pb.RefFieldFightRivalUserAck{}
		err := this.RivalUser(user, ack, true)
		if err != nil {
			logger.Error("玩家上线刷新劲敌玩家失败  err:%v", err)
		}
	}
	return
}

func (this *FieldManager) setUserMatchCombat(user *objs.User) bool {

	isRival := false
	openDay := this.GetFieldFight().GetOpenDayByReduceTime()
	combat := rmodel.FieldFight.GetFieldFightMatchUserFiveAmCombat(openDay, user.Id)
	if combat == 0 {
		isRival = true
		//长时间不活跃玩家 上线时存储他的当前战力 用于匹配野战玩家
		rmodel.FieldFight.SetFieldFightMatchUserFiveAmCombat1(openDay, user.Id, user.Combat)
	}
	return isRival
}

func (this *FieldManager) checkIconIsOpen(user *objs.User) bool {
	funCfg := gamedb.GetFunctionFunctionCfg(115)
	if funCfg != nil {
		state := this.GetCondition().CheckMulti(user, -1, funCfg.Condition)
		if !state {
			return false
		}
	}
	return true
}

func (this *FieldManager) OnLine(user *objs.User) {
	//是否需要初始化 玩家5点存的战力 和 劲敌列表 判断
	this.GetFieldFight().JudgeIsSetUserFiveAmCombat(user)
	this.DayReset(user, false)
}
