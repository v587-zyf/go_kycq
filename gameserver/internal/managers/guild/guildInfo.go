package guild

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"time"
)

func (this *GuildManager) Init() error {

	this.CreateRobotGuild()
	allGuildInfos, _ := modelGame.GetGuildModel().GetAllGuildInfos()
	if allGuildInfos != nil {
		for _, info := range allGuildInfos {
			if info.IsDelete > 0 {
				continue
			}
			this.guildInfo[info.GuildId] = info
		}
	}
	return nil
}

func (this *GuildManager) BuildGuildInfo(user *objs.User, guildInfo *modelGame.Guild) *pb.GuildInfo {
	onLineUserCount := 0
	guildMemberInfo := make([]*pb.GuildMenberInfo, 0)
	for i, j := 0, len(guildInfo.Positions); i < j; i += 2 {
		userId := guildInfo.Positions[i]
		position := guildInfo.Positions[i+1]
		memberInfo := &pb.GuildMenberInfo{UserId: int32(userId), Position: int32(position)}
		if userId < 0 {
			cfg := gamedb.GetGuildRobotGuildRobotCfg(-userId)
			if cfg != nil {
				memberInfo.Lv = int32(cfg.Level)
				memberInfo.NickName = cfg.Name
			}
		} else {
			memberUser := this.GetUserManager().GetUser(userId)
			if memberUser != nil {
				onLineUserCount += 1
				memberInfo = this.GetMemberInfo(int32(userId), int32(position), int32(0), int32(memberUser.GuildData.GuildCapital), int32(memberUser.GuildData.ContributionValue),
					int32(this.GetExpPool().GetHeroMaxLv(memberUser)), int32(memberUser.Combat), memberUser.NickName, memberUser.Avatar, int32(memberUser.Heros[constUser.USER_HERO_MAIN_INDEX].Job), int32(memberUser.Heros[constUser.USER_HERO_MAIN_INDEX].Sex))
			} else {
				memberUser1 := this.GetUserManager().GetUserBasicInfo(userId)
				if memberUser1 != nil {
					memberInfo = this.GetMemberInfo(int32(userId), int32(position), int32(memberUser1.LastUpdateTime.Unix()), int32(memberUser1.GuildData.GuildCapital), int32(memberUser1.GuildData.ContributionValue),
						int32(memberUser1.Level), int32(memberUser1.Combat), memberUser1.NickName, memberUser1.Avatar, int32(memberUser1.HeroDisplay[constUser.USER_HERO_MAIN_INDEX].Job), int32(memberUser1.HeroDisplay[constUser.USER_HERO_MAIN_INDEX].Sex))
				}

			}
		}

		guildMemberInfo = append(guildMemberInfo, memberInfo)
	}

	guildLv, _ := gamedb.GetGuildMemberLimit(guildInfo.GuildContributionValue)

	_, positionCount, _ := this.GetGuildPositionInfo(user.GuildData.NowGuildId)
	return &pb.GuildInfo{GuildId: int32(guildInfo.GuildId), GuildName: guildInfo.GuildName, GuildLv: int32(guildLv), JoinCd: int32(user.GuildData.JoinCD), GuildMenberInfo: guildMemberInfo, Notice: guildInfo.Notice, IsAutoAgree: int32(guildInfo.AutoAgree), PositionCount: positionCount, OnlineUser: int32(onLineUserCount), GuildContributionValue: int32(guildInfo.GuildContributionValue), Combat: int64(guildInfo.JoinCombat), ServerId: int32(guildInfo.ServerId)}
}

func (this *GuildManager) GetMemberInfo(userId, pos, offUnx, capital, ContributionValue, lv, combat int32, name, avatar string, job, sex int32) *pb.GuildMenberInfo {

	return &pb.GuildMenberInfo{
		UserId:            userId,
		Position:          pos,
		OfflineTime:       offUnx,
		GuildCapital:      capital,
		GuildContribution: ContributionValue,
		NickName:          name,
		Avatar:            avatar,
		Lv:                lv,
		Combat:            combat,
		Job:               job,
		Sex:               sex,
	}
}

//更行玩家身上数据
func (this *GuildManager) UpdateUserGuildDataInfo(user *objs.User, guildId, position int, isSetJoinCd, isResetContribute, isCreateGuild, isOfflineUser bool) {
	if user == nil {
		logger.Error("UpdateUserGuildDataInfo user == nil")
		return
	}
	if user.GuildData.NowGuildId > 0 {
		user.GuildData.BeforeGuildId = user.GuildData.NowGuildId
	}
	user.GuildData.NowGuildId = guildId
	user.GuildData.Position = position
	if isCreateGuild {
		user.GuildData.MyCreateId = guildId
	}
	if isResetContribute {
		user.GuildData.ContributionValue = 0
	}
	if isSetJoinCd {
		user.GuildData.JoinCD = int(time.Now().Unix()) + gamedb.GetConf().ChangeGuildInterval
	}
	user.Dirty = true
	if isOfflineUser {
		_, err := modelGame.GetUserModel().DbMap().Update(user.User)
		if err != nil {
			logger.Error("存储离线玩家数据失败  userId:%v err:%v", user.Id, err)
		}
	} else {
		_ = this.GetUserManager().Save(user, true)
	}
}

func (this *GuildManager) GetGuildPositionInfo(guildId int) (userPosition, positionCount map[int32]int32, err error) {
	userPosition = make(map[int32]int32)
	positionCount = make(map[int32]int32)
	guildInfo := this.GetGuild().GetGuildInfo(guildId)
	if guildInfo == nil {
		return userPosition, positionCount, gamedb.ERRNOGUILD
	}

	for i, j := 0, len(guildInfo.Positions); i < j; i += 2 {
		userId := guildInfo.Positions[i]
		position := guildInfo.Positions[i+1]
		userPosition[int32(userId)] = int32(position)
		positionCount[int32(position)] += 1
	}
	return userPosition, positionCount, nil
}

func (this *GuildManager) GetApplyUserLists(guildId int) []*pb.BriefUserInfo {
	infos := make([]*pb.BriefUserInfo, 0)
	allUser, err := rmodel.Guild.GetGuildApplyUserIds(guildId)
	if err != nil {
		return infos
	}
	for userId := range allUser {
		info := this.GetUserManager().BuilderBrieUserInfo(userId)
		if info == nil {
			logger.Error("GetAllUserInfoIncludeOfflineUser nil userId:%v", userId)
			continue
		}
		infos = append(infos, &pb.BriefUserInfo{
			Id: int32(info.Id), Name: info.Name, Lvl: int32(info.MaxLv), Combat: int64(info.Combat),
		})
	}
	return infos
}

//
//  CheckChairmanIsLoggedIn
//  @Description: //处理弹劾掌门
//  @param userId 掌门id
//  @param guildId 所在门派id
//
func (this *GuildManager) CheckChairmanIsLoggedIn(userId, guildId int) {
	checkTime := time.Now().Unix()
	logger.Info("CheckChairmanIsLoggedIn 有人弹劾会长 处理弹劾协程  userId:%v, guildId:%v  checkTime:%v", userId, guildId, checkTime)
	ticker1 := time.NewTicker(time.Second * 30)

	for {
		select {

		case <-ticker1.C:
			logger.Debug("掌门:%v 上线检查  处理让位 checkTime:%v  nowTime:%v", userId, checkTime, time.Now().Unix())
			if int(time.Now().Unix()-checkTime) >= gamedb.GetConf().ImpeachSuccess*60 {
				logger.Info("掌门:%v  被弹劾半小时后还没上线  处理让位", userId)
				this.ChairmanIdAbdicate(guildId)
				_ = rmodel.Guild.DelGuildImpeachPresident(userId)
				return
			}
			userInfo := this.GetUserManager().GetUser(userId)
			if userInfo != nil {
				_ = rmodel.Guild.DelGuildImpeachPresident(userId)
				logger.Info("掌门上线:%v 弹劾失败", userId)
				return
			}
		}
	}
}

//处理掌门让位
func (this *GuildManager) ChairmanIdAbdicate(guildId int) {
	logger.Info("弹劾成功 处理掌门让位  guildId:%v", guildId)
	guildInfo := this.GetGuild().GetGuildInfo(guildId)
	if guildInfo == nil {
		logger.Error("GetGuildInfo nil  guildId:%v  ", guildId)
		return
	}

	userPosition := make(map[int]int)
	positionCount := make(map[int]int)

	for i, j := 0, len(guildInfo.Positions); i < j; i += 2 {
		userId := guildInfo.Positions[i]
		position := guildInfo.Positions[i+1]
		userPosition[userId] = position
		positionCount[position] += 1
	}
	onUserCombat := make(map[int]int)

	for userId := range userPosition {
		userInfo := this.GetUserManager().GetUser(userId)
		if userInfo != nil {
			onUserCombat[userId] = userInfo.Combat
		} else {
			userInfo1 := this.GetUserManager().GetUserBasicInfo(userId)
			if userInfo1 != nil && int(time.Now().Unix())-int(userInfo1.LastUpdateTime.Unix()) < 72*3600 {
				onUserCombat[userId] = userInfo1.Combat
			}
		}
	}
	logger.Info("userPosition:%v positionCount:%v ", userPosition, positionCount)
	logger.Info("onUserCombat:%v  ", onUserCombat)
	maxUserId := this.GetMaxCombatUserByPosition(guildId, onUserCombat)
	logger.Info("maxUserId:%v", maxUserId)
	if maxUserId > 0 {
		this.TopOffPosition(maxUserId, guildInfo)
		return
	}
	return
}

//顶掉会长的位置
func (this *GuildManager) TopOffPosition(maxUserId int, guildInfo *modelGame.Guild) {
	logger.Info("TopOffPosition maxUserId :%v, guildInfo:%v", maxUserId, guildInfo)
	positionSlice := make([]int, 0)

	for i, j := 0, len(guildInfo.Positions); i < j; i += 2 {
		flag := false
		userId := guildInfo.Positions[i]
		pos := guildInfo.Positions[i+1]
		logger.Debug("TopOffPosition userId:%v pos:%v", userId, pos)
		if userId == guildInfo.ChairmanId {
			pos = pb.GUILDPOSITION_CHENGYUAN
			flag = true
		}
		if userId == maxUserId {
			pos = pb.GUILDPOSITION_HUIZHANG
			flag = true
		}
		if flag {
			info := this.GetUserManager().GetUser(userId)
			if info == nil {
				info = this.GetUserManager().GetOfflineUserInfo(userId)
				info.GuildData.Position = pos
				_, err := modelGame.GetUserModel().DbMap().Update(info.User)
				if err != nil {
					logger.Error("会长被弹劾  顶掉会长 err:%v的位置err:%v userId:%v", guildInfo.ChairmanId, err, maxUserId)
					return
				}
			} else {
				info.GuildData.Position = pos
				info.Dirty = true
			}
		}

		positionSlice = append(positionSlice, userId, pos)
	}
	logger.Info("positionSlice:%v maxUserId:%v", positionSlice, maxUserId)
	infos := make([]int32, 0)
	infos = append(infos, int32(maxUserId), pb.GUILDPOSITION_HUIZHANG, int32(guildInfo.ChairmanId), pb.GUILDPOSITION_CHENGYUAN)
	guildInfo.Positions = positionSlice
	guildInfo.ChairmanId = maxUserId
	this.SetGuildInfo(guildInfo)
	this.BroadcastGuildInfoNtf(guildInfo.GuildId, pb.GUILDBROCASTTYPE_TAN_HE_HUI_ZHANG, infos, nil)
}

func (this *GuildManager) SetGuildInfo(guildInfo *modelGame.Guild) {
	this.Lock()
	defer this.Unlock()
	this.guildInfo[guildInfo.GuildId] = guildInfo
	err := modelGame.GetGuildModel().Update(guildInfo)
	if err != nil {
		logger.Error("GetGuildModel().Update err:%v  guildId:%v", err, guildInfo.GuildId)
	}
}

func (this *GuildManager) DelGuildInfo(guildInfo *modelGame.Guild) {
	this.Lock()
	defer this.Unlock()
	if _, ok := this.guildInfo[guildInfo.GuildId]; ok {
		delete(this.guildInfo, guildInfo.GuildId)
	}
}

func (this *GuildManager) GetGuildInfo(guildId int) *modelGame.Guild {
	this.RLock()
	defer this.RUnlock()
	info := this.guildInfo[guildId]
	if info == nil {
		logger.Debug("GetGuildInfo  guildId:%v  nil", guildId)
		return nil
	}
	return info
}

func (this *GuildManager) GetAllGuildInfos() map[int]*modelGame.Guild {
	this.RLock()
	defer this.RUnlock()
	data := make(map[int]*modelGame.Guild)
	for k, v := range this.guildInfo {
		data[k] = v
	}
	return data
}

func (this *GuildManager) ResetGuildBonfireDonateInfo() {
	this.Lock()
	defer this.Unlock()
	for _, info := range this.guildInfo {
		info.DonateUsers = make(model.IntSlice, 0)
		info.DonateTimes = make(model.IntKv, 0)
	}
}

func (this *GuildManager) GetGuildName(userId int) string {
	guildName := ""
	if userId <= 0 {
		return guildName
	}
	dataUserInfo := this.GetUserManager().GetUserBasicInfo(userId)
	if dataUserInfo != nil {
		guildInfo := this.GetGuild().GetGuildInfo(dataUserInfo.GuildData.NowGuildId)
		logger.Debug("NowGuildId:%v", dataUserInfo.GuildData.NowGuildId)
		if guildInfo != nil {
			guildName = guildInfo.GuildName
		}
	}
	return guildName

}

func (this *GuildManager) BroadcastChatToGuildUsers(user *objs.User, protoMsg *pb.ChatMessageNtf) error {
	if user.GuildData.NowGuildId <= 0 {
		logger.Error("玩家没有门派")
		return gamedb.ERRHAVENOGUILD
	}
	guildInfo := this.GetGuildInfo(user.GuildData.NowGuildId)
	if guildInfo == nil {
		logger.Error("门派不存在 guildId:%v", user.GuildData.NowGuildId)
		return gamedb.ERRNOGUILD
	}

	for i, j := 0, len(guildInfo.Positions); i < j; i += 2 {
		userId := guildInfo.Positions[i]
		protoMsg.ToId = int32(userId)
		guildUser := this.GetUserManager().GetUser(userId)
		if guildUser != nil {
			this.GetUserManager().SendMessage(guildUser, protoMsg, true)
		}
	}
	return nil
}

func (this *GuildManager) GetGuildHuiAndFuHuiUserIds(guildId int) []int {

	data := make([]int, 0)
	err, memberInfos, _, _ := this.GetGuild().GetGuildMemberInfo(guildId)
	if err != nil {
		logger.Error("GetGuildMemberInfo err:%v  guildId:%v", err, guildId)
		return data
	}
	hui := memberInfos[pb.GUILDPOSITION_HUIZHANG]
	fuHui := memberInfos[pb.GUILDPOSITION_FUHUIZHANG]

	for _, userId := range hui {
		data = append(data, userId, pb.GUILDPOSITION_HUIZHANG)
	}
	for _, userId := range fuHui {
		data = append(data, userId, pb.GUILDPOSITION_FUHUIZHANG)
	}
	return data
}

func (this *GuildManager) BroadcastGuildInfoNtf(guildId, types int, infos []int32, memberInfo []*pb.GuildMenberInfo) {
	if guildId <= 0 {
		logger.Error("玩家没有门派")
		return
	}
	guildInfo := this.GetGuildInfo(guildId)
	if guildInfo == nil {
		logger.Error("门派不存在 guildId:%v", guildId)
		return
	}
	logger.Info("guildId:%v  guildInfo.Positions:%v", guildId, guildInfo.Positions)
	_, positionCount, _ := this.GetGuildPositionInfo(guildId)
	for i, j := 0, len(guildInfo.Positions); i < j; i += 2 {
		userId := guildInfo.Positions[i]
		guildUser := this.GetUserManager().GetUser(userId)
		if guildUser != nil {
			this.GetUserManager().SendMessage(guildUser, &pb.BroadcastGuildChangeNtf{Types: int32(types), UserInfos: infos, PositionCount: positionCount, GuildMenberInfo: memberInfo}, true)
		}
	}
	return
}

//取最高职位的最高战力的玩家
func (this *GuildManager) GetMaxCombatUserByPosition(guildId int, onUserCombat map[int]int) int {

	err, guildMemberInfo, _, _ := this.GetGuild().GetGuildMemberInfo(guildId)
	logger.Debug("guildMemberInfo :%v, err:%v", guildMemberInfo, err)

	maxCombat := 0
	maxUserId := 0
	for i := pb.GUILDPOSITION_FUHUIZHANG; i <= pb.GUILDPOSITION_CHENGYUAN; i++ {
		data := guildMemberInfo[i]
		if data == nil {
			continue
		}
		for _, userId := range data {
			combat := onUserCombat[userId]
			if combat > maxCombat {
				maxCombat = combat
				maxUserId = userId
			}
		}
		break
	}
	return maxUserId
}
