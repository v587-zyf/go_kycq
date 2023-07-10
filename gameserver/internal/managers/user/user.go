package user

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"database/sql"
	"sync"
	"time"
)

type UserManager struct {
	util.DefaultModule
	managersI.IModule
	users             map[int]*objs.User            // 存储所有在线玩家，key为userId
	usersByOpenId     map[int]map[string]*objs.User // 存储所有在线玩家，key为serverId->openId
	usersBasicInfoMap map[int]*modelGame.UserBasicInfo
	usersMu           sync.Mutex
	hasRoleName       map[string]string // 存用户已有昵称
}

func NewUserManager(m managersI.IModule) *UserManager {
	userManager := &UserManager{
		users:             make(map[int]*objs.User),
		usersByOpenId:     make(map[int]map[string]*objs.User),
		usersBasicInfoMap: make(map[int]*modelGame.UserBasicInfo),
	}
	userManager.IModule = m
	return userManager
}

func (this *UserManager) Init() error {
	// 加载本服所有玩家
	allUsers, err := modelGame.GetUserModel().LoadAllUsers()
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if err == sql.ErrNoRows {
		return nil
	}
	this.hasRoleName = make(map[string]string, 0)
	userIds := make([]int, len(allUsers))
	for k, uInfo := range allUsers {
		userIds[k] = uInfo.Id
		this.usersBasicInfoMap[uInfo.Id] = uInfo

		if uInfo.NickName == "" {
			continue
		}
		this.hasRoleName[uInfo.NickName] = uInfo.NickName
	}

	if len(userIds) > 0 {
		//加载武将信息
		allHeros, err1 := modelGame.GetHeroModel().GetHerosDisplayByUserId(userIds)
		if err1 != nil && err1 != sql.ErrNoRows {
			return err1
		}
		for _, v := range allHeros {
			if this.usersBasicInfoMap != nil {
				if this.usersBasicInfoMap[v.UserId].HeroDisplay == nil {
					this.usersBasicInfoMap[v.UserId].HeroDisplay = make(map[int]*modelGame.HeroDisplay)
				}
				this.usersBasicInfoMap[v.UserId].HeroDisplay[v.Index] = v
				if v.Index == constUser.USER_HERO_MAIN_INDEX {
					this.usersBasicInfoMap[v.UserId].Level = v.ExpLvl
				}
			}
		}

	}

	logger.Info("加载全服玩家，共有玩家：%v", len(allUsers))
	return nil
}

func (this *UserManager) LoadUser(openId string, channelId int, clientIp string, serverId int, origin string, deviceId string) (*objs.User, error) {

	if user := this.GetUserByOpenId(openId, serverId); user != nil { //如果用户已存在，则踢下线

		logger.Info("user is already online，kick off openId=%v", openId)
		err := this.KickUserWithMsg(user, gamedb.ERRLOGGEDINOTHERDEVICE.Message)
		if err != nil {
			logger.Error("踢出玩家异常：%v", err)
		}
		user.OffLineWg.Wait()
	}

	user := objs.NewUser()
	user.DeviceId = deviceId
	user.Ip = clientIp
	user.Origin = origin
	var err error
	logger.Info("start load user openId=%v,serverIndex=%v", openId, serverId)
	user.User, err = modelGame.GetUserModel().GetByOpenId(openId, serverId)
	if err != nil && err != sql.ErrNoRows {
		logger.Info("loginResult=%v", err)
		return nil, err
	}
	if err == sql.ErrNoRows {
		err = this.newUser(user, openId, channelId, serverId)
	} else {
		err = this.getHero(user)
	}
	if err != nil {
		return nil, err
	}

	//加载账号信息，判断是否被封禁
	if !this.Test() {
		account, erra := modelCross.GetAccountModel().GetByOpenId(openId)
		if erra != nil {
			return nil, erra
		}
		user.AccountInfo = account
	}

	if ban, _ := this.CheckBan(user, constUser.BAN_TYPE_LOGIIN); ban {
		return nil, gamedb.ERROPENIDLOCKED
	}

	user.ChannelId = channelId
	user.LastUpdateTime = time.Now()
	user.OfflineAwardMark = false
	taskConf := gamedb.GetTaskConditionCfg(user.MainLineTask.TaskId)
	if taskConf != nil {
		if taskConf.ConditionType == pb.CONDITION_NPC_CHAT {
			user.MainLineTask.Process = 1
		}
	}
	this.AddUser(user)

	//玩家登录背包数据处理
	this.GetBag().Online(user)
	//玩家登入仓库背包数据处理
	this.GetBag().WareHouseOnline(user)
	this.GetStageManager().Online(user)
	this.GetEquip().Online(user)
	this.GetRein().OnLine(user)
	this.GetMaterialStage().OnLine(user)
	this.GetFieldBoss().Online(user)
	this.GetArena().Online(user)
	this.GetVipManager().Online(user)
	this.GetGift().Online(user)

	this.GetOnline().Online(user)
	this.GetMining().Online(user)
	this.GetExpStage().OnLine(user)
	this.GetDarkPalace().Online(user)
	this.GetPersonBoss().Online(user)
	this.GetVipBoss().OnLine(user)
	this.GetShop().OnLine(user)
	this.GetTalent().Online(user)
	this.GetFriend().Online(user)
	this.GetDailyTask().InfoInit(user)
	this.GetDailyTask().OnLine(user)
	this.GetMonthCard().Online(user)
	this.GetAchievement().Online(user)
	this.GetDailyPack().Online(user)
	this.GetWarOrder().Online(user)
	this.GetCardActivity().Rest(user, true)
	this.GetTreasure().Reset(user, true)
	this.GetSign().Online(user)
	this.GetCompetitve().OnLine(user)
	this.GetFieldFight().OnLine(user)
	this.GetPet().Online(user)
	this.GetDailyRank().OnlineCheck(user)
	this.GetSpendRebates().Online(user)
	this.GetAncientBoss().Online(user)
	this.GetRecharge().ContRechargeReset(user, false)
	this.GetLottery().OnlineCheck(user)
	this.GetTrialTask().OnlineCheck(user)
	this.GetTitle().Online(user)
	this.GetDaBao().Online(user)
	this.GetApplets().OnlineAddPhysicalPower(user)
	this.GetHellBoss().Online(user)
	this.GetLabel().Online(user)
	this.GetElf().Online(user)
	this.GetWing().Online(user)
	this.GetPrivilegeModule().Online(user)
	this.GetFirstDrop().Reset(user, common.GetResetTime(time.Now()))
	//更新每日状态
	this.updateDailyState(user)
	//主线任务处理
	//m.Task.CheckMainLineTask(user)
	//if user.Lvl == 0 {
	//	return user, nil
	//}
	this.GetTlog().PlayerLogin(user) //处理玩家老挂机时间
	this.UpdateUserCombatForIsLogin(user, -1, true, false)
	this.SyncUser(user, constUser.USER_STATUS_ONLINE)
	logger.Info("LoadUser OK. userId=%v nickName=%v serverId:%v", user.Id, user.NickName, user.ServerId)

	this.BroadcastAll(&pb.UserOnlineNtf{UserId: int32(user.Id), OnlineTime: time.Now().Unix()})
	rmodel.Friend.DelFriendInfo(user.Id)
	//进入服务器日志
	kyEvent.UserEnterServer(user)
	//记录登录日志
	if len(user.Heros) > 0 {
		kyEvent.UserLogin(user)
	}
	return user, nil
}

func (this *UserManager) getHero(user *objs.User) error {
	//获取武将信息
	heros, err := modelGame.GetHeroModel().GetHerosByUserId(user.Id)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	if heros != nil && len(heros) > 0 {
		for _, v := range heros {
			user.Heros[v.Index] = objs.NewHero(v)
		}
	}
	return nil
}

func (this *UserManager) newUser(user *objs.User, openId string, channelId, serverId int) error {

	logger.Info("Create new user openId=%v accountId=%v", openId)
	serverIndex := this.GetSystem().GetServerIndex(serverId)
	u := modelGame.NewUser(openId, channelId, serverId, serverIndex)
	user.User = u
	//初始化新玩家任务
	this.GetTask().InitNewUserTask(user)
	err := modelGame.GetUserModel().Create(u)
	if err != nil {
		return err
	}

	return nil
}
