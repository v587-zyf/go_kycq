package user

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"fmt"
	"time"
)

/**
*  @Description: 获取在线人数
*  @receiver this
*  @return int
**/
func (this *UserManager) GetOnlineTotal() int {
	return len(this.users)
}

func (this *UserManager) GetUserBasicInfo(userId int) *modelGame.UserBasicInfo {
	this.usersMu.Lock()
	defer this.usersMu.Unlock()
	u := this.users[userId]
	resultUserInfo := this.usersBasicInfoMap[userId]
	if u != nil {
		if resultUserInfo == nil {
			resultUserInfo = &modelGame.UserBasicInfo{HeroDisplay: make(map[int]*modelGame.HeroDisplay)}
		}
		resultUserInfo.Id = u.Id
		resultUserInfo.OpenId = u.OpenId
		resultUserInfo.NickName = u.NickName
		resultUserInfo.Vip = u.VipLevel
		resultUserInfo.Combat = u.Combat
		resultUserInfo.Avatar = u.Avatar
		resultUserInfo.ServerId = u.ServerId
		resultUserInfo.LastUpdateTime = u.LastUpdateTime
		resultUserInfo.ChannelId = u.ChannelId
		resultUserInfo.GuildData = u.GuildData
		lv := 0
		if len(u.Heros) > 0 {
			lv = u.Heros[constUser.USER_HERO_MAIN_INDEX].ExpLvl
		}
		resultUserInfo.Level = lv
		for _, v := range u.Heros {
			resultUserInfo.HeroDisplay[v.Index] = &modelGame.HeroDisplay{
				UserId:  u.Id,
				Index:   v.Index,
				Job:     v.Job,
				Sex:     v.Sex,
				Name:    v.Name,
				Combat:  v.Combat,
				Display: v.Display,
				ExpLvl:  v.ExpLvl,
			}
		}
	}
	return resultUserInfo
}

func (this *UserManager) GetAllUsersBasicInfo() map[int]*modelGame.UserBasicInfo {
	defer this.usersMu.Unlock()
	this.usersMu.Lock()
	return this.usersBasicInfoMap
}

func (this *UserManager) GetUser(userId int) *objs.User {
	this.usersMu.Lock()
	user := this.users[userId]
	this.usersMu.Unlock()
	return user
}

func (this *UserManager) GetAllOnlineUserInfo() map[int]*objs.User {
	defer this.usersMu.Unlock()
	this.usersMu.Lock()
	return this.users
}

func (this *UserManager) GetUserByOpenId(openId string, serverId int) *objs.User {
	this.usersMu.Lock()
	if this.usersByOpenId[serverId] == nil {
		this.usersByOpenId[serverId] = make(map[string]*objs.User)
	}
	user := this.usersByOpenId[serverId][openId]
	this.usersMu.Unlock()
	return user
}

func (this *UserManager) getJobCombatType(job int) int {
	if job == pb.JOB_ZHANSHI {
		return pb.RANKTYPE_COMBAT_ZHANSHI
	} else if job == pb.JOB_FASHI {
		return pb.RANKTYPE_COMBAT_FASHI
	} else {
		return pb.RANKTYPE_COMBAT_DAOSHI
	}
}

func (this *UserManager) BuilderBrieUserInfo(userId int) *pb.BriefUserInfo {
	UserInfo := &pb.BriefUserInfo{}
	user := this.GetUser(userId)
	if user != nil {
		UserInfo.Id = int32(user.Id)
		UserInfo.Name = user.NickName
		UserInfo.Lvl = int32(user.Heros[constUser.USER_HERO_MAIN_INDEX].ExpLvl)
		UserInfo.Sex = int32(user.Heros[constUser.USER_HERO_MAIN_INDEX].Sex)
		UserInfo.Vip = int32(user.VipLevel)
		UserInfo.Job = int32(user.Heros[constUser.USER_HERO_MAIN_INDEX].Job)
		UserInfo.Avatar = user.Avatar
		UserInfo.Combat = int64(user.Combat)
		pbDisplay := make(map[int32]*pb.Display)
		for index, data := range user.Heros {
			pbDisplay[int32(index)] = this.GetHeroDisplay(data)
		}
		UserInfo.Display = pbDisplay
		UserInfo.ServerId = int32(user.ServerId)
		UserInfo.MaxLv = int32(this.GetExpPool().GetHeroMaxLv(user))
	} else {
		userBaseInfo := this.GetUserBasicInfo(userId)
		if userBaseInfo != nil {
			UserInfo.Id = int32(userBaseInfo.Id)
			UserInfo.Name = userBaseInfo.NickName
			UserInfo.Avatar = userBaseInfo.Avatar
			UserInfo.Vip = int32(userBaseInfo.Vip)
			if _, ok := userBaseInfo.HeroDisplay[constUser.USER_HERO_MAIN_INDEX]; ok {
				UserInfo.Lvl = int32(userBaseInfo.HeroDisplay[constUser.USER_HERO_MAIN_INDEX].ExpLvl)
				UserInfo.Sex = int32(userBaseInfo.HeroDisplay[constUser.USER_HERO_MAIN_INDEX].Sex)
				UserInfo.Job = int32(userBaseInfo.HeroDisplay[constUser.USER_HERO_MAIN_INDEX].Job)
			}
			UserInfo.Avatar = userBaseInfo.Avatar
			UserInfo.Combat = int64(userBaseInfo.Combat)
			pbDisplay := make(map[int32]*pb.Display)
			maxLv := 0
			for index, data := range userBaseInfo.HeroDisplay {
				if data.ExpLvl > maxLv {
					maxLv = data.ExpLvl
				}
				pbDisplay[int32(index)] = &pb.Display{
					ClothItemId:     int32(data.Display.ClothItemId),
					ClothType:       int32(data.Display.ClothType),
					WeaponItemId:    int32(data.Display.WeaponItemId),
					WeaponType:      int32(data.Display.WeaponType),
					WingId:          int32(data.Display.WingId),
					MagicCircleLvId: int32(data.Display.MagicCircleLvId)}
			}
			UserInfo.Display = pbDisplay
			UserInfo.ServerId = int32(userBaseInfo.ServerId)
			UserInfo.MaxLv = int32(maxLv)
		}
	}

	return UserInfo
}

//玩家不在线就去数据库拉取
func (this *UserManager) BuilderAllUserInfoAndOffline(user *objs.User, rivalUserId int) *pb.BriefUserInfo {
	if rivalUserId < 0 {
		//假人    去取robot表
		return this.GetRobotUserInfo(user, -rivalUserId)
	}
	UserInfo := &pb.BriefUserInfo{}
	rivalUser := this.GetUserManager().BuilderBrieUserInfo(rivalUserId)
	if rivalUser == nil {
		logger.Error("BuilderAllUserInfo nil  rivalUserId:%v", rivalUserId)
		return nil
	}
	UserInfo.Id = rivalUser.Id
	UserInfo.Name = rivalUser.Name
	UserInfo.Avatar = rivalUser.Avatar
	UserInfo.Combat = rivalUser.Combat
	UserInfo.Lvl = rivalUser.MaxLv
	UserInfo.Job = rivalUser.Job
	UserInfo.Sex = rivalUser.Sex
	if rivalUser.MaxLv == 0 {
		logger.Debug("BuilderBrieUserInfo userId:%v rivalUserId:%v", user.Id, rivalUserId)
		UserInfo.Lvl = int32(user.GetMaxHeroLv())
	}

	return UserInfo
}

func (this *UserManager) GetRobotUserInfo(user *objs.User, rivalUserId int) *pb.BriefUserInfo {
	userInfo := &pb.BriefUserInfo{}
	logger.Debug("GetRobotUserInfo  rivalUserId:%v", rivalUserId)
	robotCfg := gamedb.GetRobotRobotCfg(rivalUserId)
	if robotCfg == nil {
		logger.Error("取假人没有取到   rivalUserId:%v", rivalUserId)
		return nil
	}
	userInfo.Id = int32(-rivalUserId)
	//prefix := this.GetSystem().GetPrefix()
	//serverIndex := user.ServerIndex
	//userInfo.Name = fmt.Sprintf("%s%d.%s", prefix, serverIndex, robotCfg.Name)
	userInfo.Name = fmt.Sprintf("%s", robotCfg.Name)
	if len(robotCfg.Icon) > 0 {
		userInfo.Avatar = fmt.Sprintf("%v", robotCfg.Icon[0])
	}
	userInfo.Combat = int64(user.Combat)
	robotCombat := builder.GetCompetitiveRobotCombat(rivalUserId)
	if robotCombat > 0 {
		userInfo.Combat = robotCombat
	}
	userInfo.Lvl = int32(robotCfg.Level)
	userInfo.Job = int32(robotCfg.Job[0])
	userInfo.Sex = int32(robotCfg.Gender[0])
	return userInfo
}

func (this *UserManager) GetAllUserInfoIncludeOfflineUser(userId int) *objs.User {
	if userId < 0 {
		return nil
	}

	userInfo := this.GetUser(userId)
	if userInfo != nil {
		return userInfo
	}

	return this.GetOfflineUserInfo(userId)
}

func (this *UserManager) GetOfflineUserInfo(userId int) *objs.User {

	user := objs.NewUser()
	var err error
	user.User, err = modelGame.GetUserModel().GetByUserId(userId)
	if err != nil {
		logger.Error("GetOfflineUserInfo GetByUserId  userId:%v err:%v", userId, err)
		return nil
	}
	err = this.getHero(user)
	if err == nil {
		return user
	} else {
		logger.Error("GetOfflineUserInfo err:%v  UserId:%v", err, userId)
		return nil
	}

}

/**
 *  @Description: 获取用户在线状态和最后更新时间
 *  @param userId
 *  @return online
 *  @return lastUpdateTime
 */
func (this *UserManager) GetUserOnlineStatus(userId int) (bool, int64) {
	var online bool
	var lastUpdateTime time.Time
	user, ok := this.users[userId]
	if ok {
		online = true
		lastUpdateTime = user.LastUpdateTime
	} else {
		user, ok := this.usersBasicInfoMap[userId]
		if ok {
			lastUpdateTime = user.LastUpdateTime
		}
	}
	return online, lastUpdateTime.Unix()
}

/**
 *  @Description: 获取指定数量用户id
 *  @param getNum	数量
 *  @param hasMap	去除id集合
 *  @return map[int]int
 */
func (this *UserManager) RandGetUserId(getNum int, hasMap map[int]int) map[int]int {
	userIdMap := make(map[int]int)
	num := 0
	if getNum > 0 {
		for userId := range this.usersBasicInfoMap {
			if hasMap != nil {
				if _, ok := hasMap[userId]; ok {
					continue
				}
			}
			userIdMap[userId] = 0
			num++
			if num >= getNum {
				break
			}
		}
	}
	return userIdMap
}
