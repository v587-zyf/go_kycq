package user

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"github.com/go-sql-driver/mysql"
	"math/rand"
	"time"
)

// 记录本服玩家
func (this *UserManager) AddLocalServerInfo(user *objs.User) {
	if user.NickName == "" {
		return
	}
	userBasicInfo := this.usersBasicInfoMap[user.Id]
	if userBasicInfo == nil {
		this.usersBasicInfoMap[user.Id] = &modelGame.UserBasicInfo{}
		userBasicInfo = this.usersBasicInfoMap[user.Id]
	}
	userBasicInfo.Id = user.Id
	userBasicInfo.OpenId = user.OpenId
	userBasicInfo.NickName = user.NickName
	userBasicInfo.Vip = user.VipLevel
	userBasicInfo.Combat = user.Combat
	userBasicInfo.Avatar = user.Avatar
	userBasicInfo.ServerId = user.ServerId
	userBasicInfo.LastUpdateTime = user.LastUpdateTime
	userBasicInfo.ChannelId = user.ChannelId
	userBasicInfo.GuildData = user.GuildData
	userBasicInfo.HeroDisplay = make(map[int]*modelGame.HeroDisplay)
	for _, v := range user.Heros {
		userBasicInfo.HeroDisplay[v.Index] = &modelGame.HeroDisplay{
			UserId:  user.Id,
			Index:   v.Index,
			Job:     v.Job,
			Sex:     v.Sex,
			Name:    v.Name,
			Combat:  v.Combat,
			Display: v.Display,
			ExpLvl:  v.ExpLvl,
		}
	}
	if userBasicInfo.HeroDisplay[constUser.USER_HERO_MAIN_INDEX] != nil {
		userBasicInfo.Level = userBasicInfo.HeroDisplay[constUser.USER_HERO_MAIN_INDEX].ExpLvl
	}
	this.hasRoleName[user.NickName] = user.NickName
}

func (this *UserManager) CreateRole(user *objs.User, nickName string, avatar string, sex, job int) error {

	nickName = common.StringValid(nickName)
	if _, ok := this.hasRoleName[nickName]; ok {
		return gamedb.ERROCCUPIEDNICKNAME
	}
	err := base.CheckName(nickName)
	if err != nil {
		return err
	}
	user.NickName = nickName
	if avatar != "" {
		user.Avatar = avatar
	} else {
		if sex == 0 {
			user.Avatar = "head.d/head18.jpg"
		} else {
			user.Avatar = "head.d/head03.jpg"
		}
	}
	this.Save(user, true)

	heroIndex, err1 := this.CreateHero(user, sex, job)

	if err1 != nil {
		return err1
	}

	this.UpdateUserCombatForIsLogin(user, heroIndex, true, false)
	this.GetFieldFight().OnLine(user)

	//记录玩家名字
	this.usersMu.Lock()
	this.AddLocalServerInfo(user)
	this.usersMu.Unlock()
	this.GetTlog().PlayerRegister(user)
	kyEvent.UserCreate(user)
	//记录登录日志
	kyEvent.UserLogin(user)
	logger.Info("CreateRole OK. userId=%v nickName=%v", user.Id, user.NickName)
	return nil
}

func (this *UserManager) CreateHero(user *objs.User, sex, job int) (int, error) {

	for _, v := range user.Heros {
		if v.Job == job {
			return 0, gamedb.ERRJOBHAS
		}
	}

	hero := modelGame.NewHero(user.Id, sex, job)
	hero.Index = len(user.Heros) + 1
	user.Heros[hero.Index] = objs.NewHero(hero)
	err := modelGame.GetHeroModel().Create(hero)
	if err == nil && hero.Index != constUser.USER_HERO_MAIN_INDEX {
		this.UpdateUserCombatForIsLogin(user, hero.Index, false, false)
		kyEvent.UserHeroCreate(user, hero.Index)
	}
	if len(user.Heros) >= 2 {
		this.GetTask().AddTaskProcess(user, pb.CONDITION_UNLOCK_SECOND_HERO, -1)
	}
	if len(user.Heros) >= 3 {
		this.GetTask().AddTaskProcess(user, pb.CONDITION_UNLOCK_THREE_HERO, -1)
	}

	this.GetCondition().RecordCondition(user, pb.CONDITION_UNLOCK_SECOND_HERO, []int{0})
	if err != nil {
		logger.Error("创建武将数据异常,玩家：%v-%v,err:%v", user.Id, user.NickName, err)
	}

	return hero.Index, err
}

func (this *UserManager) AddUser(user *objs.User) {
	this.usersMu.Lock()
	this.users[user.Id] = user
	if this.usersByOpenId[user.ServerId] == nil {
		this.usersByOpenId[user.ServerId] = make(map[string]*objs.User)
	}
	this.usersByOpenId[user.ServerId][user.OpenId] = user
	logger.Info("usermanager user add,userId:%v,openId:%v,serverid:%v", user.Id, user.OpenId, user.ServerId)
	this.AddLocalServerInfo(user)
	this.usersMu.Unlock()
}

func (this *UserManager) saveInDB(user *objs.User) {
	this.GetWarOrder().WriteWarOrderTask(user, pb.WARORDERCONDITION_ONLINE, []int{int(time.Now().Unix()) - int(user.LastUpdateTime.Unix())})
	user.LastUpdateTime = time.Now()
	this.saveModuleUserInfo(user)
	_, err := modelGame.GetUserModel().DbMap().Update(user.User)
	if err != nil {
		//检查mysql中是否有重复字段。
		me, _ := err.(*mysql.MySQLError)
		logger.Info("saveuserinfo UserManager:Save:userId:%v NickName:%v Avatar:%v err:%v mysqlErr:%v", user.Id, user.NickName, user.Avatar, err, me)
		return
	}
	//保存武将信息
	for _, v := range user.Heros {
		_, err := modelGame.GetHeroModel().DbMap().Update(v.Hero)
		if err != nil {
			logger.Error("保存武将信息异常：%v", err)
		}
	}

	logger.Info("UserManager Save :id:%d,nickName:%s,vip:%v,mi:%v", user.Id, user.NickName, user.VipLevel, user.Ingot)
	user.Dirty = false
}

func (this *UserManager) Save(user *objs.User, force bool) error {
	if force || (user.Dirty && time.Since(user.LastUpdateTime) > time.Second*5) {
		this.saveInDB(user)
		this.SyncUser(user, 0)
	}
	return nil
}

func (this *UserManager) RandName(user *objs.User, sex int, ack *pb.RandNameAck) error {
	userInfo := this.GetUserBasicInfo(user.Id)
	if userInfo != nil && len(user.NickName) > 0 {
		return gamedb.ERRPARAM
	}
	if len(user.Heros) > 0 {
		return gamedb.ERRROLEEXISTS
	}
	nameSlice := gamedb.GetRoleName(sex)
	randNum := 500
	hasName := 0
	for i := 1; i <= randNum; i++ {
		name := nameSlice[rand.Intn(len(nameSlice))]
		_, has := this.hasRoleName[name]
		if !has {
			ack.Names = append(ack.Names, name)
			hasName++
		}
		if hasName >= 5 {
			break
		}
	}
	return nil
}

//存储玩家模块数据
func (this *UserManager) saveModuleUserInfo(user *objs.User) {
	//每日任务 资源找回处理
	this.GetDailyTask().OfflineSaveDailyTaskInfo(user)
	this.GetTrialTask().OfflineSaveTrialTaskInfo(user)
	return
}
