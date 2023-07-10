package system

import (
	"cqserver/gamelibs/modelCross"
	"cqserver/gameserver/internal/managersI"
	"cqserver/protobuf/pb"
	"errors"
	"sync"
	"time"

	"cqserver/gameserver/internal/base"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"runtime/debug"
)

type SystemManager struct {
	util.DefaultModule
	managersI.IModule
	serverInfos      map[int]*modelCross.ServerInfo
	crossServerInfos map[int]modelCross.ServerInfo
	lock             sync.Mutex
	closeFuncIds     []int
}

func NewSystemManager(module managersI.IModule) *SystemManager {
	return &SystemManager{
		serverInfos:      make(map[int]*modelCross.ServerInfo),
		crossServerInfos: make(map[int]modelCross.ServerInfo),
		IModule:          module,
	}
}

func (this *SystemManager) Init() error {
	serverInfos, err := modelCross.GetServerInfoModel().GetServerInfoByMergerServerId(base.Conf.ServerId)
	if err != nil {
		return err
	}
	for _, v := range serverInfos {
		this.serverInfos[v.ServerId] = v
	}

	if this.serverInfos[base.Conf.ServerId] == nil {
		return errors.New("获取服务器信息错误")
	}

	if this.serverInfos[base.Conf.ServerId].CrossFsId > 0 {
		datas, _ := modelCross.GetServerInfoModel().GetAllServerIdsByCrossFsIds(this.serverInfos[base.Conf.ServerId].CrossFsId)
		if datas != nil {
			for _, info := range datas {
				this.crossServerInfos[info.ServerId] = info
			}
		}
	}
	this.UpdateFuncState(false)
	go this.reloadServerInfo()

	logger.Info("GsPort=%v", this.serverInfos[base.Conf.ServerId].GsPort)
	return nil
}

// 重新加载服务器信息
func (this *SystemManager) reloadServerInfo() {
	tickerReload := time.NewTicker(time.Second * 10)

	for {
		select {
		case <-tickerReload.C:
			this.reloadServerTicker()
		}
	}
}

func (this *SystemManager) reloadServerTicker() {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("reloadServerInfo panic: %v, %s", err, debug.Stack())
		}
	}()

	serverInfo, err := modelCross.GetServerInfoModel().GetServerInfoByServerId(base.Conf.ServerId)
	if err != nil {
		logger.Error("reloadServerTicker DB Error: %v", err)
		return
	}

	if serverInfo.OpenTime.Unix() != this.GetMainServerOpenTime().Unix() {
		if serverInfo.OpenTime.Unix() < time.Now().Unix() {
			modelCross.GetServerInfoModel().UpdateServerOpenTime(base.Conf.ServerId, this.GetMainServerOpenTime())
			logger.Error("开服时间修改失效，只能大于当前系统时间: %v -> %v", this.GetMainServerOpenTime(), serverInfo.OpenTime)
			return
		}

		logger.Info("开服时间发生变化: %v -> %v", this.GetMainServerOpenTime(), serverInfo.OpenTime)
		this.SetServerOpenTimeByServerId(base.Conf.ServerId, serverInfo.OpenTime)
	}
	//修改跨服组
	if serverInfo.CrossFsId != this.serverInfos[base.Conf.ServerId].CrossFsId {
		this.serverInfos[base.Conf.ServerId].CrossFsId = serverInfo.CrossFsId
		this.fightCrossChangeConnect()
	}

	if this.serverInfos[base.Conf.ServerId].CrossFsId > 0 {
		datas, _ := modelCross.GetServerInfoModel().GetAllServerIdsByCrossFsIds(this.serverInfos[base.Conf.ServerId].CrossFsId)
		if datas != nil {
			for _, info := range datas {
				this.crossServerInfos[info.ServerId] = info
			}
		}
	}

}

//根据serverId获取开服天数
func (this *SystemManager) GetServerOpenDaysByServerId(serverId int) int {
	this.lock.Lock()
	defer this.lock.Unlock()
	serverInfo := this.getServerInfo(serverId)
	if serverInfo == nil {
		return -1
	}
	return common.GetTheDays(this.serverInfos[serverId].OpenTime)
}

//根据serverId获取和服天数
func (this *SystemManager) GetServerMergeDayByServerId(serverId int) int {
	this.lock.Lock()
	defer this.lock.Unlock()
	serverInfo := this.getServerInfo(serverId)
	if serverInfo == nil || serverInfo.MergeServerId < 1 {
		return -1
	}
	if serverInfo.MergeTime.Year()/2000 < 1 {
		return 0
	}
	return common.GetTheDays(this.serverInfos[serverId].MergeTime)
}

//获取服务器信息
func (this *SystemManager) getServerInfo(serverId int) *modelCross.ServerInfo {
	serverInfo, ok := this.serverInfos[serverId]
	if !ok {
		var err error
		serverInfo, err = modelCross.GetServerInfoModel().GetServerInfoByServerId(serverId)
		if err != nil {
			logger.Error("GetServerOpenDaysByServerId DB Error: %v serverId:%v", err, serverId)
			return nil
		}
		this.serverInfos[serverId] = serverInfo
	}
	return serverInfo
}

//获取合服后主服务器 开服天数
func (this *SystemManager) GetMergerServerOpenDaysByServerId(serverId int) int {
	this.lock.Lock()
	defer this.lock.Unlock()
	if _, ok := this.serverInfos[serverId]; !ok {
		serverInfo, err := modelCross.GetServerInfoModel().GetServerInfoByServerId(serverId)
		if err != nil {
			logger.Error("GetServerOpenDaysByServerId DB Error: %v serverId:%v", err, serverId)
			return -1
		}
		this.serverInfos[serverId] = serverInfo
	}
	mergerServerId := this.serverInfos[serverId].MergeServerId
	if mergerServerId == serverId {
		return common.GetTheDays(this.serverInfos[serverId].OpenTime)
	}

	if _, ok := this.serverInfos[mergerServerId]; !ok {
		serverInfo, err := modelCross.GetServerInfoModel().GetServerInfoByServerId(mergerServerId)
		if err != nil {
			logger.Error("GetServerOpenDaysByServerId DB Error: %v serverId:%v", err, mergerServerId)
			return -1
		}
		this.serverInfos[mergerServerId] = serverInfo
	}
	return common.GetTheDays(this.serverInfos[mergerServerId].OpenTime)
}

func (this *SystemManager) GetServerName(serverId int) string {
	this.lock.Lock()
	defer this.lock.Unlock()
	serverInfo := this.getServerInfo(serverId)
	if serverInfo == nil {
		return ""
	}
	return serverInfo.Name
}

//根据serverId获取开服天数//
//  GetServerOpenDaysByServerIdExcursionTime
//  @Description: 取当前时间的偏移 减少 reduceHour 小时
func (this *SystemManager) GetServerOpenDaysByServerIdByExcursionTime(serverId int, reduceHour time.Duration) int {
	this.lock.Lock()
	defer this.lock.Unlock()
	if _, ok := this.serverInfos[serverId]; !ok {
		serverInfo, err := modelCross.GetServerInfoModel().GetServerInfoByServerId(serverId)
		if err != nil {
			logger.Error("GetServerOpenDaysByServerId DB Error: %v serverId:%v", err, serverId)
			return -1
		}
		this.serverInfos[serverId] = serverInfo
	}
	openDay := common.GetTheDaysReduceHour(this.serverInfos[serverId].OpenTime, reduceHour)
	if openDay < 0 {
		//开服第一天特殊处理
		return 1
	}
	return openDay
}

//取主服务器的开服天数
func (this *SystemManager) GetRandomServerOpenDays() int {
	return common.GetTheDays(this.GetMainServerOpenTime())
}

// 根据玩家Id获取开服天数,如果出错则根据this.GetOpenDays()返回一个值。
func (this *SystemManager) GetServerOpenDaysByUserId(userId int) int {
	userBInfo := this.GetUserManager().GetUserBasicInfo(userId)
	if userBInfo != nil {
		return this.GetServerOpenDaysByServerId(userBInfo.ServerId)
	} else {
		logger.Error("get user: %d brief info error, user's brief info is nil, return random open days", userId)
		return this.GetRandomServerOpenDays()
	}
}

//根据serverId获取开服时间
func (this *SystemManager) GetServerOpenTimeByServerId(serverId int) time.Time {
	this.lock.Lock()
	defer this.lock.Unlock()
	if serverInfo, ok := this.serverInfos[serverId]; ok {
		return serverInfo.OpenTime
	}
	serverInfo, err := modelCross.GetServerInfoModel().GetServerInfoByServerId(serverId)
	if err != nil {
		logger.Error("GetServerOpenTimeByServerId DB Error: %v serverId:%v", err, serverId)
		return time.Now()
	}
	this.serverInfos[serverId] = serverInfo
	return this.serverInfos[serverId].OpenTime
}

//取主服务器的开服时间
func (this *SystemManager) GetMainServerOpenTime() time.Time {
	return this.GetServerOpenTimeByServerId(base.Conf.ServerId)
}

//根据serverId修改开服时间
func (this *SystemManager) SetServerOpenTimeByServerId(serverId int, openTime time.Time) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.serverInfos[serverId].OpenTime = openTime
}

/**
 * 获得本服在指定天数的时候是开服第几天
 */
func (this *SystemManager) GetOpenDaysByDesDate(desDate time.Time) int {
	serverOpenTime := this.GetMainServerOpenTime()
	serverOpenZeroTime := common.ZeroTimeOfDay(serverOpenTime).Unix()
	nowZeroTime := common.ZeroTimeOfDay(desDate).Unix()
	openDays := (nowZeroTime-serverOpenZeroTime)/(24*60*60) + 1
	return int(openDays)
}

func (this *SystemManager) GetPrefix() string {
	if this.serverInfos[base.Conf.ServerId].Prefix == "" {
		return "s"
	}
	return this.serverInfos[base.Conf.ServerId].Prefix
}

func (this *SystemManager) GetCrossFightServerId() int {
	return this.serverInfos[base.Conf.ServerId].CrossFsId
}

func (this *SystemManager) IsCross() bool {
	return this.serverInfos[base.Conf.ServerId].CrossFsId > 0
}

func (this *SystemManager) IsMerge() bool {
	return this.serverInfos[base.Conf.ServerId].MergeServerId > 0
}

func (this *SystemManager) ServerIdInLocalServer(serverId int) bool {
	return this.serverInfos[serverId] != nil
}

func (this *SystemManager) GetServerIndex(serverId int) int {

	if this.serverInfos[serverId] == nil {
		serverId = base.Conf.ServerId
	}
	return this.serverInfos[serverId].ServerIndex
}

func (this *SystemManager) GetServerIndexCrossFsId(serverId int) int {

	if this.serverInfos[serverId] == nil {
		serverId = base.Conf.ServerId
	}
	return this.serverInfos[serverId].CrossFsId
}

func (this *SystemManager) GetServerInfoByServerId(serverId int) *modelCross.ServerInfo {

	if this.serverInfos[serverId] == nil {
		serverId = base.Conf.ServerId
	}
	return this.serverInfos[serverId]
}
func (this *SystemManager) fightCrossChangeConnect() {

	//同步跨服战斗服常驻战斗（炼狱首领）
	go this.GetFight().SyncResidentFightId()
	//// 通知原跨服战斗服踢人下线
	//notifyReq := &pbserver.NodifyPlayerToLeaveFightReq{ServerId: int32(base.Conf.ServerId)}
	//notifyResp := &pbserver.NodifyPlayerToLeaveFightAck{}
	//err := m.FsManager.RpcCall(pbserver.CmdNodifyPlayerToLeaveFightReqId, notifyReq, notifyResp)
	//if err != nil {
	//	logger.Error("ModifyCrossGroup RpcCall err=%v", err)
	//}
	//logger.Info("ModifyCrossGroup serverId=%v kickPlayers=%v", base.Conf.ServerId, notifyResp.PlayerCount)
}

func (this *SystemManager) UpdateFuncState(sendClient bool) {
	all := modelCross.GetFuncCloseDbModel().Getall()
	this.closeFuncIds = make([]int, 0)
	if all != nil {
		for _, v := range all {
			if len(v.ServerIds) > 0 {
				for _, sid := range v.ServerIds {
					if sid == base.Conf.ServerId {
						this.closeFuncIds = append(this.closeFuncIds, v.FuncId)
						break
					}
				}
			} else {
				this.closeFuncIds = append(this.closeFuncIds, v.FuncId)
			}
		}
	}
	logger.Info("接收ccs发送来，更新模块开关信息：%v", this.closeFuncIds)
	if sendClient {
		ntf := &pb.FuncStateCloseNtf{
			CloseFuncId: common.ConvertIntSlice2Int32Slice(this.closeFuncIds),
		}
		this.BroadcastAll(ntf)
	}
}

func (this *SystemManager) GetFuncState() []int {
	return this.closeFuncIds
}

func (this *SystemManager) GetServerMergerIdAndMergerTime(serverId int) (int, time.Time) {
	this.lock.Lock()
	defer this.lock.Unlock()
	if _, ok := this.serverInfos[serverId]; !ok {
		serverInfo, err := modelCross.GetServerInfoModel().GetServerInfoByServerId(serverId)
		if err != nil {
			logger.Error("GetServerOpenDaysByServerId DB Error: %v serverId:%v", err, serverId)
			return -1, time.Time{}
		}
		this.serverInfos[serverId] = serverInfo
	}
	return this.serverInfos[serverId].MergeServerId, this.serverInfos[serverId].MergeTime
}

func (this *SystemManager) GetCrossServerBriefUserInfo() map[int32]*pb.BriefServerInfo {
	this.lock.Lock()
	defer this.lock.Unlock()
	data := make(map[int32]*pb.BriefServerInfo)
	if this.crossServerInfos == nil || len(this.crossServerInfos) <= 0 {
		return data
	}
	for _, info := range this.crossServerInfos {
		data[int32(info.ServerId)] = &pb.BriefServerInfo{ServerId: int32(info.ServerId), ServerName: info.Name, CrossFsId: int32(info.CrossFsId)}
	}
	return data
}
