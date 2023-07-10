package manager

import (
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/publicCon/constServer"
	"cqserver/gamelibs/rmodelCross"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"runtime/debug"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ServerListManager struct {
	util.DefaultModule
	serverMu         sync.RWMutex
	serverInfos      []*modelCross.ServerInfo // 常规服务器列表 - 用于发送给客户端
	newestServer     *modelCross.ServerInfo   //最新服
	trailServerInfo  *modelCross.ServerInfo   // 当前提审服信息
	trailVersion1    int                      // 提审服版本-高位
	trailVersion2    int                      // 提审服版本-中位
	trailVersion3    int                      // 提审服版本-低位
	trailVersion     string
	whiteListMu      sync.RWMutex
	whiteListIp      map[string]bool //白名单
	whiteListAccount map[string]bool //白名单
}

func NewServerListManager() *ServerListManager {
	return &ServerListManager{
		trailVersion1: 99,
		trailVersion2: 99,
		trailVersion3: 99,
	}
}

func (this *ServerListManager) Init() error {
	this.fetchServerList()
	this.updateWhiteList()
	go this.updateServerList()
	return nil
}

func (this *ServerListManager) Stop() {
}

func (this *ServerListManager) updateServerList() {
	var ticker = time.NewTicker(time.Second * 5)
	var whiteListTicker = time.NewTicker(time.Minute * 5)
	for {
		select {
		case <-ticker.C:
			this.fetchServerList()
			continue
		case <-whiteListTicker.C:
			this.updateWhiteList()
		}
	}
}

func (this *ServerListManager) fetchServerList() {
	defer func() {
		if r := recover(); r != nil {
			stackBytes := debug.Stack()
			logger.Error("panic fetchServerList:%v,%s", r, stackBytes)
		}
		this.serverMu.Unlock()
	}()

	this.serverMu.Lock()
	serverInfos, err := modelCross.GetServerInfoModel().GetServerInfos()
	if err != nil {
		logger.Error("fetchServerList:%v", err)
		return
	}
	now := time.Now()
	this.serverInfos = make([]*modelCross.ServerInfo, 0)
	this.newestServer = nil
	for _, serverInfo := range serverInfos {

		//开服时间未到
		if serverInfo.OpenTime.After(now) {
			continue
		}
		// 抓取提审服的版本号
		if serverInfo.IsTrialVersion == constServer.SERVER_TRIAL {
			//ts := strings.Split(serverInfo.Version, ".")
			//this.trailVersion1, _ = strconv.Atoi(strings.TrimSpace(ts[0]))
			//this.trailVersion2, _ = strconv.Atoi(strings.TrimSpace(ts[1]))
			//this.trailVersion3, _ = strconv.Atoi(strings.TrimSpace(ts[2]))
			this.trailServerInfo = serverInfo
		} else {
			this.serverInfos = append(this.serverInfos, serverInfo)
			if this.newestServer == nil || this.newestServer.OpenTime.Before(serverInfo.OpenTime) {
				this.newestServer = serverInfo
			}
		}
	}

	version, _ := modelCross.GetSystemSettingModel().GetSetting(rmodelCross.SYSTEM_SETTING_TRIAL_SERVER_VERSION)
	if len(version) > 0 {
		this.trailVersion = version
		ts := strings.Split(version, ".")
		this.trailVersion1, _ = strconv.Atoi(strings.TrimSpace(ts[0]))
		this.trailVersion2, _ = strconv.Atoi(strings.TrimSpace(ts[1]))
		this.trailVersion3, _ = strconv.Atoi(strings.TrimSpace(ts[2]))
	}
	logger.Info("版号服版本号=%v.%v.%v  正式服数=%v", this.trailVersion1, this.trailVersion2, this.trailVersion3, len(this.serverInfos))
}

// 根据最后登录服务器ID、版本号、渠道标识获取服务器列表
func (this *ServerListManager) GetServerList() []*modelCross.ServerInfo {

	if this.serverInfos == nil {
		this.fetchServerList()
	}

	this.serverMu.RLock()
	servers := make([]*modelCross.ServerInfo, 0)

	// 是否是提审版本
	//if this.isTrailVersion(version) {
	//	logger.Info("isTrailVersion v=%v appId=%v", version, appId)
	//	return []modelCross.ServerInfo{this.trailServerInfo}
	//} else {
	// 构建待选的服务器列表
	for _, serverInfo := range this.serverInfos {
		//if serverInfo.AppId == appId {
		servers = append(servers, serverInfo)
		//}
	}
	this.serverMu.RUnlock()
	return servers
	//}
}

// 是否是提审版本，即当前版本>=提审服版本
func (this *ServerListManager) isTrailVersion(version string) bool {
	defer func() {
		if r := recover(); r != nil {
			stackBytes := debug.Stack()
			logger.Error("panic isTrailVersion:%v,%s", r, stackBytes)
		}
	}()

	curS := strings.Split(version, ".")
	v1, _ := strconv.Atoi(strings.TrimSpace(curS[0]))
	v2, _ := strconv.Atoi(strings.TrimSpace(curS[1]))
	v3, _ := strconv.Atoi(strings.TrimSpace(curS[2]))
	if v1 > this.trailVersion1 {
		return true
	} else if v1 < this.trailVersion1 {
		return false
	}

	if v2 > this.trailVersion2 {
		return true
	} else if v2 < this.trailVersion2 {
		return false
	}

	if v3 > this.trailVersion3 {
		return true
	} else if v3 < this.trailVersion3 {
		return false
	}

	return false
}

func (this *ServerListManager) TrailVersion() string {
	return this.trailVersion
}

func (this *ServerListManager) GetServerInfoById(serverId int, appId string) *modelCross.ServerInfo {
	//serverInfos := this.GetServerList()
	this.serverMu.RLock()
	for _, serverInfo := range this.serverInfos {
		if serverInfo.ServerId == serverId && (serverInfo.AppId == appId || appId == "") {
			this.serverMu.RUnlock()
			return serverInfo
		}
	}
	this.serverMu.RUnlock()
	return nil
}

func (this *ServerListManager) GetNewestServer() *modelCross.ServerInfo {
	return this.newestServer
}

func (this *ServerListManager) GetTrialServer() *modelCross.ServerInfo {
	return this.trailServerInfo
}

func (this *ServerListManager) updateWhiteList() {
	all := modelCross.GetWhiteListDbModel().Getall()
	this.whiteListMu.Lock()
	this.whiteListIp = make(map[string]bool)
	this.whiteListAccount = make(map[string]bool)
	if len(all) > 0 {
		for _, v := range all {
			if v.Valtype == 1 {
				this.whiteListIp[v.Value] = true
			} else {
				this.whiteListAccount[v.Value] = true
			}
		}
	}
	this.whiteListMu.Unlock()
}

func (this *ServerListManager) InWhiteList(ip, openId string) bool {
	this.whiteListMu.RLock()
	in := false
	if _, ok := this.whiteListIp[ip]; ok {
		in = true
	}
	if _, ok := this.whiteListAccount[openId]; ok {
		in = true
	}
	this.whiteListMu.RUnlock()
	return in
}
