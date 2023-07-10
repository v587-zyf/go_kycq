package activeUser

import (
	"cqserver/crosscenterserver/internal/managersI"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/publicCon/constActiveUser"
	"cqserver/gamelibs/rmodelCross"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"sync"
	"time"
)

type ActiveUser struct {
	util.DefaultModule
	managersI.IModule
	muActiveUser sync.RWMutex

	maxActiveDay         int                                       //可以记录的最大活跃天数
	allActiveUser        map[int]*modelCross.UserCrossInfo         //所有活跃玩家的信息
	allActiveUserByServe map[int]map[int]*modelCross.UserCrossInfo //map[serverId]map[userId]info
	dayActiveUserId      map[int]map[int]bool                      //map[day]map[userId]true	每日活跃的玩家id

}

func NewActiveUserManager(m managersI.IModule) *ActiveUser {
	return &ActiveUser{
		IModule:              m,
		allActiveUser:        make(map[int]*modelCross.UserCrossInfo),
		allActiveUserByServe: make(map[int]map[int]*modelCross.UserCrossInfo),
		dayActiveUserId:      make(map[int]map[int]bool),
	}
}

func (this *ActiveUser) Init() error {
	err := this.activeUserInit()
	if err != nil {
		return err
	}
	this.checkActiveUser()
	return nil
}

func (this *ActiveUser) checkActiveUser() {
	ticker := time.NewTicker(time.Second * 30)
	go func() {
		for {
			select {
			case <-ticker.C:
				this.updateActiveUsers()
			}
		}
	}()
}

func (this *ActiveUser) setActiveUsersInfo(info *modelCross.UserCrossInfo) {
	this.muActiveUser.Lock()
	defer this.muActiveUser.Unlock()
	this.allActiveUser[info.UserId] = info
}

func (this *ActiveUser) setActiveUsersInfoByServer(info *modelCross.UserCrossInfo) {
	this.muActiveUser.Lock()
	defer this.muActiveUser.Unlock()
	if this.allActiveUserByServe[info.ServerId] == nil {
		this.allActiveUserByServe[info.ServerId] = make(map[int]*modelCross.UserCrossInfo)
	}
	this.allActiveUserByServe[info.ServerId][info.UserId] = info
}

func (this *ActiveUser) setDayActiveUserId(info *modelCross.UserCrossInfo, activeDay int) {
	this.muActiveUser.Lock()
	defer this.muActiveUser.Unlock()
	for i := activeDay; i <= this.maxActiveDay; i++ {
		if this.dayActiveUserId[i] == nil {
			this.dayActiveUserId[i] = make(map[int]bool)
		}
		this.dayActiveUserId[i][info.UserId] = true
	}
}

func (this *ActiveUser) buildActiveUsersInfo(message *modelCross.UserCrossInfo) *constActiveUser.ActiveUsersInfo {

	info := &constActiveUser.ActiveUsersInfo{
		UserId:     message.UserId,
		UpdateTime: time.Now(),
		ServerId:   message.UserId,
		Nickname:   message.NickName,
		Avatar:     message.Avatar,
		Vip:        message.Vip,
		Combat:     message.Combat,
	}
	return info
}

func (this *ActiveUser) ActiveUsersAdd(message *modelCross.UserCrossInfo) {
	//activeUserInfo := this.buildActiveUsersInfo(message)

	this.setActiveUsersInfo(message)
	this.setActiveUsersInfoByServer(message)
	this.setDayActiveUserId(message, 1)
}

func (this *ActiveUser) GetUserIdsByActiveDay(day int) map[int]bool {
	this.muActiveUser.Lock()
	defer this.muActiveUser.Unlock()
	data := make(map[int]bool)
	if this.dayActiveUserId[day] == nil {
		return data
	}
	return this.dayActiveUserId[day]
}

func (this *ActiveUser) initActiveUserData() {
	this.muActiveUser.Lock()
	defer this.muActiveUser.Unlock()

	this.allActiveUser = make(map[int]*modelCross.UserCrossInfo)
}

func (this *ActiveUser) initActiveUserDataByServer() {
	this.muActiveUser.Lock()
	defer this.muActiveUser.Unlock()

	this.allActiveUserByServe = make(map[int]map[int]*modelCross.UserCrossInfo)
}

func (this *ActiveUser) initDayActiveUserId() {
	this.muActiveUser.Lock()
	defer this.muActiveUser.Unlock()

	this.dayActiveUserId = make(map[int]map[int]bool)
}

func (this *ActiveUser) activeUserInit() error {

	allUser, err := modelCross.GetUserCrossInfoModel().GetAllUsers()
	if err != nil {
		return err
	}
	this.maxActiveDay = rmodelCross.GetSystemSeting().GetSystemSettingConverInt(rmodelCross.SYSTEM_SETTING_MAX_ACTIVITY_DAY)
	logger.Info("updateActiveUsers|this.maxActiveDay = %d", this.maxActiveDay)

	for _, info := range allUser {
		activeDay := int(((time.Now().Unix() - info.UpdateTime.Unix()) / 86400) + 1)
		logger.Debug("活跃玩家整理，玩家:%v,serverId:%v", info.UserId, info.ServerId)
		if activeDay <= this.maxActiveDay {
			//activeInfo := this.buildActiveUsersInfo(&info)
			this.setActiveUsersInfo(info)
			this.setActiveUsersInfoByServer(info)
			this.setDayActiveUserId(info, activeDay)
		}
	}
	return nil
}

//定时清除非活跃玩家数据
func (this *ActiveUser) updateActiveUsers() {
	this.maxActiveDay = rmodelCross.GetSystemSeting().GetSystemSettingConverInt(rmodelCross.SYSTEM_SETTING_MAX_ACTIVITY_DAY)
	logger.Info("updateActiveUsers|this.maxActiveDay = %d", this.maxActiveDay)
	activeUsers := make(map[int]*modelCross.UserCrossInfo)
	this.muActiveUser.Lock()
	for userId, info := range this.allActiveUser {
		activeUsers[userId] = info
	}
	this.muActiveUser.Unlock()
	this.initActiveUserData()
	this.initActiveUserDataByServer()
	this.initDayActiveUserId()
	for _, info := range activeUsers {
		activeDay := int(((time.Now().Unix() - info.UpdateTime.Unix()) / 86400) + 1)
		logger.Debug("活跃玩家整理，玩家:%v,serverId:%v  activeDay:%v", info.UserId, info.ServerId, activeDay)
		if activeDay <= this.maxActiveDay {
			this.setActiveUsersInfo(info)
			this.setActiveUsersInfoByServer(info)
			this.setDayActiveUserId(info, activeDay)
		}
	}
}
