package user

import (
	"cqserver/crosscenterserver/internal/managersI"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/rmodelCross"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pbserver"
	"sync"
	"time"
)

type UserManager struct {
	util.DefaultModule
	managersI.IModule
	cdb           chan *modelCross.UserCrossInfo
	muUser        sync.Mutex     // 同步锁，记录更新名字
	recordUserMap map[int]string // 记录的玩家
	userServerId  map[int]int    // 记录玩家服务器ID
}

func NewUserManager(m managersI.IModule) *UserManager {
	return &UserManager{
		IModule:       m,
		cdb:           make(chan *modelCross.UserCrossInfo, 200),
		recordUserMap: make(map[int]string),
		userServerId:  make(map[int]int),
	}
}

func (this *UserManager) Init() error {
	go this.updateToDB()
	return nil
}

func (this *UserManager) updateToDB() {
	//回溯错误
	defer func() {
		if err := recover(); err != nil {
			logger.Error("UserManager dbOperate Panic Error. %v", err)
		}
	}()
	for {
		select {
		case userInfo := <-this.cdb:
			//写入跨服缓存
			err := rmodelCross.GetUserCrossInfoRmodle().Set(userInfo)
			if err != nil {
				logger.Error("setUserInfo to redis error:%v", err)
			}

			//存库
			count, err := modelCross.GetUserCrossInfoModel().Update(userInfo)
			if err != nil {
				logger.Error("updateToDB Error. %v", err)
			}
			if count == 0 {
				err = modelCross.GetUserCrossInfoModel().Create(userInfo)
				if err != nil {
					logger.Error("insertToDB error:%v", err)
				}
			}
		}
	}
}

func (this *UserManager) GetUserServerId(userId int) int {

	costTimeStart := time.Now()
	defer func() {
		if costT := time.Now().Sub(costTimeStart); costT > 10*time.Millisecond {
			logger.Warn("crossGuildFightManager cost record GetUserServerId costTime:%v lenGuild:%v", costT)
		}
	}()

	this.muUser.Lock()
	defer this.muUser.Unlock()
	serverId := this.userServerId[userId]
	if serverId == 0 {
		userInfo := rmodelCross.GetUserCrossInfoRmodle().Get(userId)
		if userInfo != nil {
			serverId = userInfo.ServerId
			serverInfo := this.GetGsServers().GetServerInfo(serverId)
			if serverInfo != nil {
				serverId = serverInfo.MergeServerId
				this.userServerId[userId] = serverId
			} else {
				this.userServerId[userId] = serverId
			}
		}
	}
	return serverId
}

func (this *UserManager) UserInfoSync(message *pbserver.SyncUserInfoNtf) {
	logger.Info("User:%d info update", message.UserId)
	userInfo := &modelCross.UserCrossInfo{
		UserId:        int(message.UserId),
		ServerId:      int(message.ServerId),
		ServerIndex:   int(message.ServerIndex),
		ChannelId:     int(message.ChannelId),
		OpenId:        message.OpenId,
		NickName:      message.Nickname,
		Avatar:        message.Avatar,
		LoginTime:     time.Unix(int64(message.LoginTime), 0),
		CreateTime:    time.Unix(int64(message.CreateTime), 0),
		UpdateTime:    time.Now(),
		OffLineTime:   time.Unix(int64(message.OfflineTime), 0),
		Gold:          int(message.Gold),
		Ingot:         int(message.Ingot),
		Vip:           int(message.Vip),
		TaskId:        int(message.TaskId),
		Combat:        int(message.Combat),
		Recharge:      int(message.Recharge),
		RechargeToken: int(message.TokenRecharge),
		Exp:           int(message.Exp),
		Heros:         make(model.CrossHeros, len(message.Heros)),
	}
	if len(message.Heros) > 0 {
		for k, v := range message.Heros {
			userInfo.Heros[k] = model.CrossHero{
				HeroIndex: int(v.HeroIndex),
				Job:       int(v.Job),
				Sex:       int(v.Sex),
				Level:     int(v.Level),
				Combat:    int(v.Combat),
			}
		}
	}
	if message.LastRechargeTime > 0 {
		userInfo.LastRechargeTime = time.Unix(int64(message.LastRechargeTime), 0)
	}

	this.cdb <- userInfo

	//存储活跃玩家数据
	this.GetActiveUser().ActiveUsersAdd(userInfo)
}
