package servers

import (
	"cqserver/crosscenterserver/internal/managersI"
	"cqserver/gamelibs/modelCross"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/golibs/util"
	"cqserver/protobuf/pbserver"
	"errors"
	"runtime/debug"
	"sync"
	"time"
)

// 服务器列表管理
type ServerListManager struct {
	util.DefaultModule
	managersI.IModule
	sMu                       sync.RWMutex
	serversInfoMap            map[int]*modelCross.ServerInfo           //服务器列表，存储服务器连接信息
	crossFightServer          map[int]*modelCross.CrossFightServerInfo //跨服战斗服
	serversInfoMapByCrossFsId map[int]map[int]*modelCross.ServerInfo   //跨服组下服务器列表，存储服务器连接信息
}

func NewServerListManager(modueI managersI.IModule) *ServerListManager {
	return &ServerListManager{
		IModule:                   modueI,
		serversInfoMap:            make(map[int]*modelCross.ServerInfo, 0),
		crossFightServer:          make(map[int]*modelCross.CrossFightServerInfo),
		serversInfoMapByCrossFsId: make(map[int]map[int]*modelCross.ServerInfo),
	}
}

func (this *ServerListManager) Init() error {
	go this.reloadServerlist()
	return nil
}

// 返回Gameserver服务器列表
func (this *ServerListManager) GetServerList() []*modelCross.ServerInfo {
	this.sMu.RLock()
	defer this.sMu.RUnlock()

	serverList := make([]*modelCross.ServerInfo, 0)
	for _, sInfo := range this.serversInfoMap {
		if sInfo.IsClose == 1 {
			serverList = append(serverList, sInfo)
		}
	}
	return serverList
}

func (this *ServerListManager) GetServerInfo(serverId int) *modelCross.ServerInfo {
	this.sMu.RLock()
	defer this.sMu.RUnlock()
	return this.serversInfoMap[serverId]
}

func (this *ServerListManager) setServerInfo(serverId int, s *modelCross.ServerInfo) {
	this.sMu.Lock()
	defer this.sMu.Unlock()
	this.serversInfoMap[serverId] = s
	if s.CrossFsId > 0 {
		if this.serversInfoMapByCrossFsId[s.CrossFsId] == nil {
			this.serversInfoMapByCrossFsId[s.CrossFsId] = make(map[int]*modelCross.ServerInfo, 0)
		}
		this.serversInfoMapByCrossFsId[s.CrossFsId][s.ServerId] = s
	}
}

//获取跨服组下 服务器信息
func (this *ServerListManager) GetAllCrossGroupServerInfo() map[int]map[int]*modelCross.ServerInfo {
	this.sMu.RLock()
	defer this.sMu.RUnlock()
	return this.serversInfoMapByCrossFsId
}

//获取跨服组下 服务器信息
func (this *ServerListManager) GetCrossGroupServerInfoByCrossFsId(crossFsId int) map[int]*modelCross.ServerInfo {
	this.sMu.RLock()
	defer this.sMu.RUnlock()
	return this.serversInfoMapByCrossFsId[crossFsId]
}

func (this *ServerListManager) SendMessage(serverId int, msg nw.ProtoMessage) error {

	session := gsManager.GetGsSession(uint32(serverId))
	if session == nil {
		return errors.New("gs not fount")
	}
	return session.SendMessage(0, msg)
}

func (this *ServerListManager) SendAllServerMessage(msg nw.ProtoMessage) {

	gsManager.Range(func(id uint32, session nw.Session) bool {
		err := session.(*GsSession).SendMessage(0, msg)
		if err != nil {
			logger.Error("推送全服消息：%v,异常：%v", msg, err)
		}
		return false
	})
}

func (this *ServerListManager) CallMessage(serverId int, requestMsg nw.ProtoMessage, resultMsg nw.ProtoMessage) error {
	session := gsManager.GetGsSession(uint32(serverId))
	if session == nil {
		return errors.New("gs not fount")
	}
	return session.CallMessage(requestMsg, resultMsg)
}

// 重载服务器列表 间隔10秒
func (this *ServerListManager) reloadServerlist() {
	this.reloadServerListTicker()
	tickerReload := time.NewTicker(time.Second * 10)
	var running bool
	for {
		select {
		case <-tickerReload.C:
			if running {
				continue
			}
			running = true
			this.reloadServerListTicker()
			running = false
		}
	}
}

func (this *ServerListManager) UpServerListTicker() {
	this.reloadServerListTicker()
}

func (this *ServerListManager) reloadServerListTicker() {

	//回溯错误
	defer func() {
		if err := recover(); err != nil {
			logger.Error("reloadServerListTicker panic: %v, %s", err, debug.Stack())
		}
	}()

	//logger.Info("加载服务器完整列表")
	// 游戏服列表
	serverInfos, err := modelCross.GetServerInfoModel().GetServerInfos()
	if err != nil {
		logger.Error("GetServerListAll error : %v", err)
		return
	}
	logger.Debug("获取服务器数量：%v", len(serverInfos))
	for _, v := range serverInfos {
		if v.ServerId != v.MergeServerId {
			continue
		}
		oldServerInfo := this.GetServerInfo(v.ServerId)
		this.setServerInfo(v.ServerId, v)
		if oldServerInfo != nil && oldServerInfo.CrossFsId != v.CrossFsId {
			ntf := &pbserver.CCSToGsCrossFsIdChangeNtf{}
			ntf.CrossFsId = int32(v.CrossFsId)
			this.SendMessage(v.ServerId, ntf)
		}
	}

	// 跨服战斗服列表
	crossFightServerInfos, errC := modelCross.GetCrossFightServerInfoModel().GetAllCrossFightServerList()
	if errC != nil {
		logger.Error("GetCrossFightServerListAll error : %v", errC)
		return
	}

	for _, v := range crossFightServerInfos {
		this.crossFightServer[v.Id] = &v
	}
}

//通过serverId获取开服天数
func (this *ServerListManager) GetOpenDaysByServerId(serverId int) int {
	info := this.GetServerInfo(serverId)
	if info != nil {
		return common.GetTheDays(info.OpenTime)
	}

	return -1
}
