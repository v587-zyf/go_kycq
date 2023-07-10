package fieldFight

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constConstant"
	"cqserver/gamelibs/publicCon/constField"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"encoding/json"
	"sort"
)

//获取每个困难度 玩家的人数 和 匹配战力
func (this *FieldManager) GetRivalUserNumAndCombat(openDay, userId int) (map[int][]int, error) {

	baseInfo := make(map[int][]int) //key:困难度 [需要的人数,匹配战力]
	fiveCombat := rmodel.FieldFight.GetFieldFightMatchUserFiveAmCombat(openDay, userId)
	fightUserCfg := gamedb.GetConf().FieldFightLevel
	for index, info := range fightUserCfg {
		if len(info) < 2 {
			logger.Error("game表 FieldFightLevel 配置错误  长度小于二  info[0]:概率 info[1]:数量")
			return nil, gamedb.ERRGAMECFGERR
		}
		rate := float32(info[0]) / constConstant.WAN_FEN_BI
		needCombat := float32(fiveCombat) * rate

		//这边  人数 * 2 是因为取出来的玩家数组 是userId,combat,userId1,combat,1
		needNumber := info[1] * 2
		if baseInfo[index+1] == nil {
			baseInfo[index+1] = make([]int, 0)
		}
		baseInfo[index+1] = append(baseInfo[index+1], needNumber, int(needCombat))
		logger.Debug("野战 匹配玩家列表 困难程度:%v 战力系数:%v 每天5点的战力:%v  需要匹配的战力:%v  需要匹配的人数:%v", index, rate, fiveCombat, needCombat, needNumber)
	}
	logger.Debug("baseInfo:%v", baseInfo)
	return baseInfo, nil
}

//获取野战挑战玩家列表
func (this *FieldManager) GetRefRivalsUsersListInfo(user *objs.User, openDay int) []*pb.FieldFightListInfo {

	allMatchUserIds := make([]int, 0)
	allRivalUserInfo := make([]*pb.FieldFightRivalUserInfo, 0)
	data, err := rmodel.FieldFight.GetFieldFightSaveBeforeRefRivals(user.Id)
	if err == nil && len(data) > 0 {
		err = json.Unmarshal([]byte(data), &allRivalUserInfo)
		if err == nil {
			for _, info := range allRivalUserInfo {
				if info.RivalUserId == 0 {
					logger.Error("野战 劲敌玩家 id 错误 玩家id:%v  劲敌id:%v", user.Id, info.RivalUserId)
					continue
				}
				allMatchUserIds = append(allMatchUserIds, int(info.RivalUserId))
			}
		}
	}

	//刷新劲敌
	allUserIds, list := this.SoreUserInfoByCombat(user, allMatchUserIds)
	if len(allUserIds) > 0 {
		bytes, _ := json.Marshal(allUserIds)
		rmodel.FieldFight.SetFieldFightSaveBeforeRefRivals(user.Id, string(bytes))
		return list
	} else {
		ack := &pb.RefFieldFightRivalUserAck{}
		err = this.RivalUser(user, ack, true)
		list = ack.ListInfo
		return list
	}
}

//老数据 开服自动删除  （策划懒-- 不肯手动清库 -^-）
func (this *FieldManager) CheckUser() {

	key := rmodel.Rank.GetRankKey(pb.RANKTYPE_COMBAT, base.Conf.ServerId)
	allData := rmodel.Rank.GetRank(key, -1)

	newData := make([]int, 0)
	if len(allData) <= 0 {
		return
	}
	changeData := make([]int, 0)
	for i, j := 0, len(allData); i < j; i += 2 {
		if allData[i] > 0 {
			userId := allData[i]
			score := allData[i+1]
			userInfo := this.GetUserManager().GetAllUserInfoIncludeOfflineUser(userId)
			if userInfo != nil {
				newData = append(newData, userId, score)
			} else {
				changeData = append(changeData, userId, score)
			}
		}
	}
	if len(changeData) <= 0 {
		return
	}

	this.GetRank().DelData(pb.RANKTYPE_COMBAT)
	for i, j := 0, len(newData); i < j; i += 2 {
		userId := newData[i]
		score := newData[i+1]
		rmodel.Rank.ZaddRankByType(pb.RANKTYPE_COMBAT, base.Conf.ServerId, userId, score)
	}
	return
}

func (this *FieldManager) SoreUserInfoByCombat(user *objs.User, allUserIds []int) ([]*pb.FieldFightRivalUserInfo, []*pb.FieldFightListInfo) {
	newAllUserIds := make([]*pb.FieldFightRivalUserInfo, 0)
	data := make([]*pb.FieldFightListInfo, 0)
	fightUserCfg := gamedb.GetConf().FieldFightLevel
	NightMare := 1  //野战 噩梦级别的玩家
	Difficulty := 2 //困难级别
	//Simple := 2     //简单
	for index, info := range fightUserCfg {
		if index == constField.NightMare {
			NightMare = int(info[1])
		}
		if index == constField.Difficulty {
			Difficulty = int(info[1])
		}
		//if index == constField.Simple {
		//	Simple = int(info[1])
		//}

	}
	field := make(model.FieldUserCombatSlice, 0)
	for _, userId := range allUserIds {
		userInfo := this.GetUserManager().BuilderAllUserInfoAndOffline(user, userId)
		if userInfo != nil {
			field = append(field, &model.FieldUserCombat{UserId: userId, Combat: int(userInfo.Combat), Avatar: userInfo.Avatar, NickName: userInfo.Name, UserLv: userInfo.Lvl, Job: userInfo.Job, Sex: userInfo.Sex})
		}
	}
	sort.Sort(field)

	num := 0
	for _, info := range field {
		num++
		if num <= NightMare {
			newAllUserIds = append(newAllUserIds, &pb.FieldFightRivalUserInfo{RivalUserId: int32(info.UserId), RivalDifficult: constField.NightMare + 1})
			data = append(data, this.buildRivalUserInfo1(info, 1))
		}
		if num > NightMare && num <= NightMare+Difficulty {
			newAllUserIds = append(newAllUserIds, &pb.FieldFightRivalUserInfo{RivalUserId: int32(info.UserId), RivalDifficult: constField.Difficulty + 1})
			data = append(data, this.buildRivalUserInfo1(info, 2))
		}
		if num > NightMare+Difficulty {
			newAllUserIds = append(newAllUserIds, &pb.FieldFightRivalUserInfo{RivalUserId: int32(info.UserId), RivalDifficult: constField.Simple + 1})
			data = append(data, this.buildRivalUserInfo1(info, 3))
		}
	}

	return newAllUserIds, data
}

func (this *FieldManager) buildRivalUserInfo1(info *model.FieldUserCombat, difficult int) *pb.FieldFightListInfo {
	return &pb.FieldFightListInfo{
		DifficultyLevel: int32(difficult),
		Avatar:          info.Avatar,
		NickName:        info.NickName,
		Combat:          int32(info.Combat),
		UserLv:          info.UserLv,
		UserId:          int32(info.UserId),
		Job:             info.Job,
		Sex:             info.Sex,
	}
}

func (this *FieldManager) buildFieldHaveAppearUserInfo(userId int) ([]int, map[int]bool) {

	haveAppearUser := make([]int, 0)
	haveAppearRobot := make(map[int]bool)

	//openDay := this.GetFieldFight().GetOpenDayByReduceTime()
	data, err := rmodel.FieldFight.GetFieldFightSaveBeforeRefRivals(userId)
	if err == nil {
		allRivalUserInfo := make([]*pb.FieldFightRivalUserInfo, 0)
		err = json.Unmarshal([]byte(data), &allRivalUserInfo)
		if err == nil {
			for _, info := range allRivalUserInfo {
				haveAppearUser = append(haveAppearUser, int(info.RivalUserId))
				if info.RivalUserId < 0 {
					haveAppearRobot[-int(info.RivalUserId)] = true
				}
			}
		}
	}
	return haveAppearUser, haveAppearRobot
}

//击败自己的玩家 信息组装
func (this *FieldManager) GetRivalUserInfos(userId, openDay int) []*pb.FieldFightBeatBackUserInfo {
	data := make([]*pb.FieldFightBeatBackUserInfo, 0)
	rivalUsers, err := rmodel.FieldFight.GetFieldFightDefeatOwnerUsers(userId, openDay)
	if err == nil {
		for userId, nickName := range rivalUsers {
			data = append(data, &pb.FieldFightBeatBackUserInfo{
				UserId:   int32(userId),
				NickName: nickName,
			})
		}
	} else {
		logger.Error("GetFieldFightDefeatOwnerUsers err:%v", err)
	}
	return data
}

func (this *FieldManager) CheckGameCfg() (gamedb.IntSlice, gamedb.ItemInfos, int, error) {

	openDay := this.GetOpenDayByReduceTime()
	fieldFightTimesCfg := gamedb.GetConf().FieldFightMaxNum
	if len(fieldFightTimesCfg) < 2 {
		logger.Error("len(FieldFightMaxNum):%v < 2 ", len(gamedb.GetConf().CompetitveTimes))
		return nil, nil, openDay, gamedb.ERRGAMECFGERR
	}
	fieldFightCostCfg := gamedb.GetConf().FieldFightBuyMaxNum
	if len(fieldFightCostCfg) < 2 {
		logger.Error("len(FieldFightBuyMaxNum):%v < 2 ", len(gamedb.GetConf().CompetitveTimes))
		return nil, nil, openDay, gamedb.ERRGAMECFGERR
	}
	return fieldFightTimesCfg, fieldFightCostCfg, openDay, nil
}

func (this *FieldManager) GetChallengeTimes(cfgTimes, times int) int {

	lastTimes := cfgTimes - times
	if lastTimes <= 0 {
		lastTimes = 0
	}
	return lastTimes
}

//@Description:  过滤掉上次玩家的劲敌列表
func (this *FieldManager) ScreenOutRivalUserInfo1(user *objs.User, rivalUserInfos []int, haveAppear []int, haveAppearRobot map[int]bool) ([]int, map[int]bool) {

	rivalUserInfos1 := make([]int, 0)
	haveAppearState := make(map[int]bool)

	for _, userId := range haveAppear {
		haveAppearState[userId] = true
	}

	if rivalUserInfos == nil {
		return rivalUserInfos1, haveAppearRobot
	}

	for i, j := 0, len(rivalUserInfos); i < j; i += 2 {
		userId := rivalUserInfos[i]
		if userId < 0 {
			haveAppearRobot[-userId] = true
		}
		if userId == user.Id {
			continue
		}
		if haveAppearState[userId] {
			continue
		}
		rivalUserInfos1 = append(rivalUserInfos1, rivalUserInfos[i])
	}

	return rivalUserInfos1, haveAppearRobot
}

func (this *FieldManager) GetListInfo(user *objs.User, ack []*pb.FieldFightListInfo, haveAppearUser map[int]bool, robotId, rivalUserId, rivalDifficult int) (*pb.BriefUserInfo, []*pb.FieldFightListInfo, int) {

	userInfo := this.GetUserManager().BuilderAllUserInfoAndOffline(user, int(rivalUserId))
	if userInfo != nil {
		nickName := userInfo.Name
		if haveAppearUser[rivalUserId] {
			//这边做假假人处理   (只有两个玩家的时候显示列表没有那么多玩家 全都是添加的这一个玩家id  但是玩家的基本信息换成假人的)
			cfg := gamedb.GetRobotRobotCfg(robotId)
			if cfg != nil {
				nickName = cfg.Name
			}
			robotId++
		}
		haveAppearUser[rivalUserId] = true
		ack = append(ack, &pb.FieldFightListInfo{DifficultyLevel: int32(rivalDifficult), Avatar: userInfo.Avatar, NickName: nickName, Combat: int32(userInfo.Combat), UserLv: int32(userInfo.Lvl), UserId: int32(rivalUserId), Job: int32(userInfo.Job), Sex: int32(userInfo.Sex)})
	} else {
		logger.Error("RefFieldFightRivalUser getUserInfo nil  userId:%v", rivalUserId)
		return nil, ack, robotId
	}
	return userInfo, ack, robotId
}

func (this *FieldManager) buildMatchUserInfo(difficulty int, matchUserInfos, rivalUserInfos []int) []int {

	for _, userId := range rivalUserInfos {
		logger.Debug("收纳的玩家 困难度:%v userId:%v", difficulty, userId)
		matchUserInfos = append(matchUserInfos, userId)
	}
	return matchUserInfos
}

func (this *FieldManager) matchRobotUser(difficult, needNum, combat int, matchUserInfos []int, haveAppearRobot map[int]bool) ([]int, map[int]bool) {
	if haveAppearRobot == nil {
		haveAppearRobot = make(map[int]bool)
	}

	for i := 1; i <= needNum; i++ {
		robotId := gamedb.GetRandomRobotIdCfg(combat, haveAppearRobot)
		if robotId == 0 {
			logger.Debug("野战真人不足  补充 困难度:%v 假人。。。但是没有匹配到合适假人 needNum:%v  Combat:%v", difficult, needNum, combat)
			continue
		}
		logger.Debug("野战 补充假人 困难度:%v  需要%v个假人  匹配战力:%v  matchUserInfos:%v  haveAppearRobot:%v", difficult, needNum, combat, matchUserInfos, haveAppearRobot)
		matchUserInfos = append(matchUserInfos, -robotId)
		haveAppearRobot[robotId] = true
	}
	return matchUserInfos, haveAppearRobot
}
