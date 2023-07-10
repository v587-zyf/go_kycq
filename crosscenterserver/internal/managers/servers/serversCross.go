package servers

import (
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/publicCon/constServer"
	"cqserver/gamelibs/rmodelCross"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"sort"
)

// 返回Gameserver服务器列表
func (this *ServerListManager) CrossMatch() {

	logger.Info("自动跨服开始")
	this.sMu.Lock()
	defer func() {
		if err := recover(); err != nil {
			logger.Error("crossMatch Panic Error. %T", err)
		}
		logger.Info("自动跨服结束")
		this.sMu.Unlock()
	}()

	matchServers := make([]*modelCross.ServerInfo, 0)
	for _, v := range this.serversInfoMap {

		canCross := this.checkCanCross(v)
		if !canCross {
			continue
		}
		matchServers = append(matchServers, v)
	}
	logger.Info("自动跨服服务器数量：%v", len(matchServers))

	//根据开服时间排序
	sort.Slice(matchServers, func(i, j int) bool {
		if matchServers[i].OpenTime.Unix() > matchServers[j].OpenTime.Unix() {
			return false
		}
		return true
	})

	//跨服
	this.doCross(matchServers)
}

/**
 *  @Description: 检查是否参与跨服
 *  @param gs
 *  @return bool
 */
func (this *ServerListManager) checkCanCross(gs *modelCross.ServerInfo) bool {

	//体验 提审服服不进入跨服
	if gs.IsTrialVersion == constServer.SERVER_TRIAL {
		return false
	}
	if gs.CrossFsId > 0 {
		return false
	}
	openDayLimit := rmodelCross.GetSystemSeting().GetSystemSettingConverInt(rmodelCross.SYSTEM_SETTING_CROSS_OPEN_DAY)
	activeUserLimit := rmodelCross.GetSystemSeting().GetSystemSettingConverInt(rmodelCross.SYSTEM_SETTING_CROSS_OPEN_ACTIVE_PLAYER)
	activeDay := rmodelCross.GetSystemSeting().GetSystemSettingConverInt(rmodelCross.SYSTEM_SETTING_CROSS_ACTIVITY_USER_DAY)
	rechargeLimit := rmodelCross.GetSystemSeting().GetSystemSettingConverInt(rmodelCross.SYSTEM_SETTING_CROSS_OPEN_SERVER_DAY_RECHARGE)

	if openDayLimit < 0 {
		return true
	}

	if activeUserLimit < 0 {
		return true
	}

	if rechargeLimit < 0 {
		return true
	}

	serveropenDay := common.GetTheDays(gs.OpenTime)
	if openDayLimit > -1 && serveropenDay > openDayLimit {
		return true
	}

	allActiveUses := this.GetActiveUser().GetUserIdsByActiveDay(activeDay)
	if len(allActiveUses) < activeUserLimit {
		return true
	}

	if rmodelCross.GetUserCrossInfoRmodle().GetDayRechargeNumByServerId(gs.ServerId) < rechargeLimit {
		return true
	}

	return false
}

func (this *ServerListManager) doCross(servers []*modelCross.ServerInfo) {

	crossServerIds := make(map[int]bool)
	crossGroup := make([][]int, 0)

	//循环第一参与跨服
	this.doCrossLoop(servers, true, crossServerIds, &crossGroup)
	//循环非第一次参与跨服
	this.doCrossLoop(servers, false, crossServerIds, &crossGroup)

	//分配跨服战斗服
	this.allotCrossFightServers(crossGroup)
}

func (this *ServerListManager) doCrossLoop(servers []*modelCross.ServerInfo, firstCross bool, crossServerIds map[int]bool, crossGroup *[][]int) {
	serverLen := len(servers)
	for k, v := range servers {

		if firstCross && v.CrossFirst == 1 {
			continue
		}
		if crossServerIds[v.ServerId] {
			continue
		}
		if k == serverLen {
			break
		}

		nowCross := make([]int, 0)
		nowCross = append(nowCross, v.ServerId)
		for i := k + 1; i < serverLen; i++ {

			if firstCross && servers[i].CrossFirst == 1 {
				continue
			}
			if crossServerIds[servers[i].ServerId] {
				continue
			}
			logger.Debug("自动跨服，服务器：%v 成功匹配到服务器：%v", v.ServerId, servers[i].ServerId)
			nowCross = append(nowCross, servers[i].ServerId)
			if len(nowCross) >= rmodelCross.GetSystemSeting().GetSystemSettingConverInt(rmodelCross.SYSTEM_SETTING_CROSS_MAX_SERVER) {
				break
			}
		}
		if len(nowCross) >= rmodelCross.GetSystemSeting().GetSystemSettingConverInt(rmodelCross.SYSTEM_SETTING_CROSS_MIN_SERVER) {
			*crossGroup = append(*crossGroup, nowCross)
			for _, v := range nowCross {
				crossServerIds[v] = true
			}
		}
	}
}

func (this *ServerListManager) allotCrossFightServers(crossServer [][]int) {

	logger.Info("开始分配跨服战斗服:%v", crossServer)
	//查找所有服务器
	crossFithtServerIds := make(map[int]bool)
	for _, v := range this.crossFightServer {
		crossFithtServerIds[v.Id] = true
	}

	for _, v := range this.serversInfoMap {
		if v.CrossFsId > 0 {
			delete(crossFithtServerIds, v.CrossFsId)
		}
	}

	//分配跨服战斗服Id
	for _, v := range crossServer {
		if len(crossFithtServerIds) == 0 {
			logger.Error("跨服战斗服不足")
			break
		}
		fsId := 0
		for fsId1 := range crossFithtServerIds {
			fsId = fsId1
			break
		}
		serversStr := common.JoinIntSlice(v, ",")
		modelCross.GetServerInfoModel().UpdateCrossGroups(serversStr, fsId)
	}
}
