package fight

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/prop"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gameserver/internal/base"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbgt"
	"cqserver/protobuf/pbserver"
	"fmt"
	"runtime/debug"
	"sync"
	"time"
)

type Fight struct {
	util.DefaultModule
	managersI.IModule
	residentFight            map[int]uint32                       //常驻战斗记录 key为战斗类型加副本条件拼接，value为战斗Id
	fieldBossInfos           map[int]*pbserver.FsFieldBossInfoNtf //野外boss
	darkPalaceBossInfos      map[int]*pbserver.FsFieldBossInfoNtf //暗殿boss
	ancientBossInfos         map[int]*pbserver.FsFieldBossInfoNtf //远古首领
	hellBossInfos            map[int]*pbserver.FsFieldBossInfoNtf //炼狱首领
	guiardPillarFightEndTime map[int]int                          //龙柱守护结束时间
	userDieTime              map[int]int                          //玩家死亡时间
	shabakeFightId           int
	fightMu                  sync.RWMutex
}

func NewFight(module managersI.IModule) *Fight {
	f := &Fight{
		IModule:                  module,
		guiardPillarFightEndTime: make(map[int]int),
		userDieTime:              make(map[int]int),
	}
	return f
}

func (this *Fight) GetFieldBossInfos(stageId int) *pbserver.FsFieldBossInfoNtf {
	this.fightMu.RLock()
	defer func() {
		this.fightMu.RUnlock()
	}()
	return this.fieldBossInfos[stageId]
}

func (this *Fight) GetDarkPalaceBossInfos(stageId int) *pb.DarkPalaceBossNtf {
	this.fightMu.RLock()
	defer func() {
		this.fightMu.RUnlock()
	}()

	if info, ok := this.darkPalaceBossInfos[stageId]; ok {
		msg := &pb.DarkPalaceBossNtf{
			info.StageId, info.Hp, info.ReliveTime,
		}
		return msg
	}

	return nil
}

func (this *Fight) GetAncientBossInfos(stageId int) *pbserver.FsFieldBossInfoNtf {
	this.fightMu.RLock()
	defer func() {
		this.fightMu.RUnlock()
	}()
	return this.ancientBossInfos[stageId]
}

func (this *Fight) GetHellBossInfos(stageId int) *pbserver.FsFieldBossInfoNtf {
	this.fightMu.RLock()
	defer func() {
		this.fightMu.RUnlock()
	}()
	return this.hellBossInfos[stageId]
}

func (this *Fight) GetFieldBossUserDieInfos(userId int) int {
	this.fightMu.RLock()
	defer func() {
		this.fightMu.RUnlock()
	}()
	return this.userDieTime[userId]
}

func (this *Fight) SyncResidentFightId() {
	if !this.GetSystem().IsCross() {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			stackBytes := debug.Stack()
			logger.Error("panic when SyncResidentFightId:%v,%s", r, stackBytes)
		}
	}()
	var syncFunc = func() bool {
		request := &pbserver.GsToFsResidentFightReq{
			ServerId: int32(base.Conf.ServerId),
		}
		replay := &pbserver.FsResidentFightNtf{}
		err := this.RpcCallByFightServerId(this.GetSystem().GetCrossFightServerId(), request, replay)
		if err != nil {
			logger.Error("常驻战斗获取战斗Id异常,err:%v", err)
			return false
		}
		logger.Info("服务器请求跨服常驻战斗信息")
		this.fightMu.Lock()
		for k, v := range replay.ResidentFights {
			this.residentFight[int(k)] = v
		}

		for k, v := range replay.FieldBossFightInfo {
			stageId := int(k)
			stageConf := gamedb.GetStageStageCfg(stageId)
			switch stageConf.Type {
			case constFight.FIGHT_TYPE_HELL_BOSS:
				this.hellBossInfos[stageId] = v
			}
		}
		this.fightMu.Unlock()
		this.GetUserManager().CrossFightOpen()
		return true
	}

	tryNum := 0
	ticker := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-ticker.C:
			tryNum++
			if syncFunc() || tryNum > 5 {
				ticker.Stop()
				return
			}
		}
	}
}

/**
*  @Description: 获取战斗Id
*  @receiver this
*  @param stageId
*  @param ext
*  @return int
*  @return error
**/
func (this *Fight) getFightIdByStageId(stageId int, ext int) (int, error) {

	request := &pbserver.GSTOFSGetFightIdReq{
		StageId: int32(stageId),
		Ext:     int32(ext),
	}
	replay := &pbserver.FSTOGSGetFightIdAck{}
	err := this.FSRpcCall(0, stageId, request, replay)
	if err != nil {
		logger.Error("战斗获取Id异常,stageId：%v,err:%v", stageId, err)
		return 0, gamedb.ERRFIGHTID
	}
	return int(replay.FightId), nil
}

func (this *Fight) getResidentFightId(stageId int) uint32 {

	stageConf := gamedb.GetStageStageCfg(stageId)
	if stageConf.Type == constFight.FIGHT_TYPE_CROSS_WORLD_LEADER {
		request := &pbserver.GSTOFSGetFightIdReq{
			StageId: int32(stageId),
		}
		replay := &pbserver.FSTOGSGetFightIdAck{}
		err := this.FSRpcCall(0, stageId, request, replay)
		if err != nil {
			logger.Error("世界boss常驻战斗获取战斗Id异常,stageId：%v,err:%v", stageId, err)
			return 0
		}
		return uint32(replay.FightId)
	} else if stageConf.Type == constFight.FIGHT_TYPE_PUBLIC_DABAO_SINGLE {
		fightId, err := this.CreateFight(stageId, nil)
		if err != nil {
			return 0
		}
		return uint32(fightId)
	}

	this.fightMu.RLock()
	defer func() {
		this.fightMu.RUnlock()
	}()
	return this.residentFight[stageId]
}

func (this *Fight) EnterFightByFightId(user *objs.User, stageId, fightId int) error {

	return this.enterFight(user, stageId, fightId, false, 0)
}

func (this *Fight) enterFight(user *objs.User, stageId, fightId int, reliveToCity bool, helpToUseId int) error {

	//离开现有战斗
	err := this.LeaveFight(user, constFight.LEAVE_FIGHT_TYPE_NOMAL)
	if err != nil {
		return err
	}
	enterType := constFight.ENTER_FIGHT_TYPE_NOMAL
	if reliveToCity {
		enterType = constFight.ENTER_FIGHT_TYPE_RELIVE_TO_CITY
	}
	//创建玩家队伍
	teamId := this.GetTeamIndex(user, stageId)
	err = this.userHeroIntoFight(user, stageId, fightId, teamId, enterType, -1, helpToUseId)
	if err != nil {
		return err
	}
	//记录玩家战斗数据
	this.recordUserFightInfo(user, stageId, fightId, 0)
	return nil
}

func (this *Fight) NewHeroIntoFight(user *objs.User, heroIndex int) {

	if user.FightStageId == 0 || user.FightStageId == constFight.FIGHT_TYPE_MAIN_CITY_STAGE {
		return
	}
	teamIndex := this.GetTeamIndex(user, user.FightStageId)
	err := this.userHeroIntoFight(user, user.FightStageId, user.FightId, teamIndex, constFight.ENTER_FIGHT_TYPE_NEW_HERO, heroIndex, 0)
	if err != nil {
		logger.Error("新武将进入战斗异常：玩家：%v,err:%v", user.IdName(), err)
	}

}

func (this *Fight) userHeroIntoFight(user *objs.User, stageId, fightId int, teamId int, enterType int, heroIndex int, helpToUseId int) error {

	//vip特权
	if heroIndex == -1 {
		this.GetSkill().ClearTreasureCd(user, stageId)
	}

	if stageId == constFight.FIGHT_TYPE_MAIN_CITY_STAGE {
		heroIndex = constUser.USER_HERO_MAIN_INDEX
	}

	//创建战斗数据
	fightUser := this.createFightUserInfo(user, teamId, heroIndex, heroIndex == -1, stageId, helpToUseId)
	////主城非一号武将的时候，冒充一号武将
	//if stageId == constFight.FIGHT_TYPE_MAIN_CITY_STAGE && heroIndex != constUser.USER_HERO_MAIN_INDEX{
	//	fightUser.UserInfo.Heros[constUser.USER_HERO_MAIN_INDEX] = fightUser.UserInfo.Heros[int32(heroIndex)]
	//	fightUser.UserInfo.Heros[constUser.USER_HERO_MAIN_INDEX].Index = int32(constUser.USER_HERO_MAIN_INDEX)
	//	delete(fightUser.UserInfo.Heros,int32(heroIndex))
	//}
	fightUser.BirthArea = int32(this.GetTask().GetFightEnterArea(user, stageId))
	//进入新战斗
	requestMsg := &pbserver.FSEnterFightReq{
		FightUser: fightUser,
	}
	requestMsg.EnterType = int32(enterType)

	replyMsg := &pbserver.FSEnterFightAck{}
	err := this.FSRpcCall(fightId, stageId, requestMsg, replyMsg)
	//err := m.FSManager.RpcCallFightCenter(crossFightServerId, fightId, requestMsg, replyMsg)
	if err != nil {
		return err
	}
	//玩家进入战斗
	kyEvent.StageStart(user, stageId)
	return err
}

func (this *Fight) EnterFightByFightIdForUserRobot(userId int, fightId int, stageId int, teamId int) error {

	var fightUser *pbserver.User
	if userId > 0 {
		user := objs.NewUser()
		var err error
		user.User, err = modelGame.GetUserModel().GetByUserId(userId)
		if err != nil {
			return gamedb.ERRUNFOUNDUSER
		}
		//获取武将信息
		heros, err1 := modelGame.GetHeroModel().GetHerosByUserId(user.Id)
		if err1 != nil {
			return err1
		}
		for _, v := range heros {
			user.Heros[v.Index] = objs.NewHero(v)
		}
		this.GetUserManager().UpdateCombatRobot(user, -1)
		//创建战斗数据
		fightUser = this.createFightUserInfo(user, teamId, -1, true, stageId, 0)
		fightUser.UserType = constFight.FIGHT_USER_TYPE_PLAYER_data
	} else {
		var err error
		fightUser, err = this.robot(-userId, teamId)
		if err != nil || fightUser == nil {
			return err
		}
		fightUser.UserType = constFight.FIGHT_USER_TYPE_CONF
	}
	if fightUser == nil {
		return gamedb.ERRUNKNOW
	}

	//进入新战斗
	requestMsg := &pbserver.FSEnterFightReq{
		FightUser: fightUser,
	}

	replyMsg := &pbserver.FSEnterFightAck{}
	err := this.FSRpcCall(fightId, stageId, requestMsg, replyMsg)
	if err != nil {
		return err
	}
	return nil
}

func (this *Fight) robot(robotId int, teamId int) (*pbserver.User, error) {

	robotConf := gamedb.GetRobotRobotCfg(robotId)
	if robotConf == nil {
		return nil, gamedb.ERRSETTINGNOTFOUND
	}

	fightUser := &pbserver.User{}
	fightUser.LocatedServerId = uint32(base.Conf.ServerId)
	fightUser.TeamId = int32(teamId)
	fightUser.UserInfo = &pbserver.Actor{
		UserId:   uint32(robotConf.Id),
		NickName: robotConf.Name,
		Elf:      &pbserver.ElfInfo{},
	}

	fightUser.UserInfo.Heros = make(map[int32]*pbserver.ActorHero)
	for k, job := range robotConf.Job {

		heroIndex := k + 1
		actorHero := &pbserver.ActorHero{}
		fightUser.UserInfo.Heros[int32(heroIndex)] = actorHero
		actorHero.Index = int32(heroIndex)
		actorHero.Job = int32(job)
		actorHero.Sex = int32(robotConf.Gender[k])
		actorHero.Level = int32(robotConf.Level)
		actorHero.NickName = fmt.Sprintf("%s.%s", robotConf.Name, modelGame.GetHeroDefName(robotConf.Gender[k], job))
		actorHero.DisplayInfo = &pbserver.ActorDisplayInfo{
			ClothItemId:  int32(robotConf.Model2[k]),
			ClothType:    pb.DISPLAYTYPE_EQUIP,
			WeaponItemId: int32(robotConf.Model1[k]),
			WeaponType:   pb.DISPLAYTYPE_EQUIP,
		}
		//技能信息
		actorHero.Skills = make([]*pbserver.Skill, 0)
		for _, v := range robotConf.Skills[k] {
			skillId, skillLv := gamedb.GetSkillIdAndLv(v)
			actorHero.Skills = append(actorHero.Skills, &pbserver.Skill{
				Id:    int32(skillId),
				Level: int32(skillLv),
			})
		}
		//属性信息
		prop := prop.NewProp()
		if k == 0 {
			prop.Add(robotConf.Property1)
		} else if k == 1 {
			prop.Add(robotConf.Property2)
		} else {
			prop.Add(robotConf.Property3)
		}
		prop.Calc(job)
		actorHero.Prop = prop.ToFightActorProp()
	}
	return fightUser, nil
}

func (this *Fight) LeaveFight(user *objs.User, reason int) error {

	//判断玩家身上战斗Id,发送玩家离开战斗
	if user.FightId > 0 {
		leaveReq := &pbserver.FSLeaveFightReq{ActorSessionId: user.GateSessionId, Reason: uint32(reason)}
		if reason == constFight.LEAVE_FIGHT_TYPE_NOMAL {
			ack := &pbserver.FSLeaveFightAck{}
			err := this.FSRpcCall(user.FightId, user.FightStageId, leaveReq, ack)
			if err != nil {
				logger.Error("请求离开当前战斗异常：user:%v,战斗Id:%v", user.IdName(), user.FightId)
				return err
			}
		} else {
			this.FSSendMessage(user.FightId, user.FightStageId, leaveReq)
		}
		kyEvent.FightLeave(user, user.FightStageId, user.FightStartTime)
	}
	return nil
}

func (this *Fight) CreateFight(stageId int, cpData []byte) (int, error) {

	fightCreateMsg := &pbserver.FSCreateFightReq{}
	fightCreateMsg.StageId = int32(stageId)
	if len(cpData) > 0 {
		fightCreateMsg.CpData = cpData
	}

	repsMsg := &pbserver.FSCreateFightAck{}
	err := this.FSRpcCall(0, stageId, fightCreateMsg, repsMsg)
	//err := m.FSManager.RpcCallFightCenter(100000, 0, fightCreateMsg, repsMsg)
	if err != nil {
		logger.Error("创建战斗异常：stageId:%v,err:%v", stageId, err)
		return 0, err
	}
	if repsMsg.FightId <= 0 {
		return 0, gamedb.ERRUNKNOW
	}
	logger.Debug("创建战斗成功：stageId:%v,fightId:%v", stageId, repsMsg.FightId)
	return int(repsMsg.FightId), nil
}

func (this *Fight) EnterResidentFightByStageId(user *objs.User, stageId int, helpUserId int) error {

	fightId := this.getResidentFightId(stageId)
	if fightId <= 0 {
		return gamedb.ERRFIGHTID
	}

	if !this.GetFight().CheckInFightBefore(user, stageId) {
		return gamedb.ERRUSERINFIGHT
	}
	return this.enterFight(user, stageId, int(fightId), false, helpUserId)
}

func (this *Fight) RecordResidentFight(fightInfos *pbserver.FsResidentFightNtf, fromCross bool) {

	logger.Info("收到战斗服常驻战斗信息：%v", *fightInfos)
	this.fightMu.Lock()
	this.residentFight = make(map[int]uint32)
	this.fieldBossInfos = make(map[int]*pbserver.FsFieldBossInfoNtf)
	this.darkPalaceBossInfos = make(map[int]*pbserver.FsFieldBossInfoNtf)
	this.ancientBossInfos = make(map[int]*pbserver.FsFieldBossInfoNtf)
	this.hellBossInfos = make(map[int]*pbserver.FsFieldBossInfoNtf)
	for k, v := range fightInfos.ResidentFights {
		this.residentFight[int(k)] = v
	}

	for k, v := range fightInfos.FieldBossFightInfo {
		stageId := int(k)
		stageConf := gamedb.GetStageStageCfg(stageId)
		switch stageConf.Type {
		case constFight.FIGHT_TYPE_FIELDBOSS:
			this.fieldBossInfos[stageId] = v
		case constFight.FIGHT_TYPE_DARKPALACE_BOSS:
			this.darkPalaceBossInfos[stageId] = v
		case constFight.FIGHT_TYPE_ANCIENT_BOSS:
			this.ancientBossInfos[stageId] = v
		case constFight.FIGHT_TYPE_HELL_BOSS:
			if this.GetSystem().GetCrossFightServerId() > 0 && !fromCross {
				continue
			}
			this.hellBossInfos[stageId] = v
		}
	}
	this.fightMu.Unlock()
}

func (this *Fight) HandlerFieldBossInfoNtf(bossInfo *pbserver.FsFieldBossInfoNtf, isFromCross bool) {

	this.fightMu.Lock()
	stageId := int(bossInfo.StageId)
	stageConf := gamedb.GetStageStageCfg(stageId)
	switch stageConf.Type {
	case constFight.FIGHT_TYPE_DARKPALACE_BOSS:
		if this.darkPalaceBossInfos[stageId] != nil {
			this.darkPalaceBossInfos[int(bossInfo.StageId)] = bossInfo
			this.GetDarkPalace().SendDarkPalaceBossNtf(&pb.DarkPalaceBossNtf{
				StageId:    bossInfo.StageId,
				Blood:      bossInfo.Hp * 100,
				ReliveTime: bossInfo.ReliveTime,
			})
		} else {
			logger.Error("战斗服发来的暗殿boss信息，game服没有记录:%v", bossInfo.StageId)
		}
	case constFight.FIGHT_TYPE_FIELDBOSS:
		if this.fieldBossInfos[stageId] != nil {
			this.fieldBossInfos[int(bossInfo.StageId)] = bossInfo
			this.GetFieldBoss().SendFieldBossNtf(&pb.FieldBossNtf{
				StageId:    bossInfo.StageId,
				Blood:      bossInfo.Hp * 100,
				ReliveTime: bossInfo.ReliveTime,
			})
		} else {
			logger.Error("战斗服发来的野外boss信息，game服没有记录:%v", bossInfo.StageId)
		}
	case constFight.FIGHT_TYPE_ANCIENT_BOSS:
		if this.ancientBossInfos[stageId] != nil {
			this.ancientBossInfos[int(bossInfo.StageId)] = bossInfo
			this.GetAncientBoss().SendBossInfo(&pb.AncientBossNtf{
				StageId:    bossInfo.StageId,
				Blood:      bossInfo.Hp * 100,
				ReliveTime: bossInfo.ReliveTime,
				UserCount:  bossInfo.UserCount,
			})
		} else {
			logger.Error("战斗服发来的远古boss信息，game服没有记录:%v", bossInfo.StageId)
		}
	case constFight.FIGHT_TYPE_HELL_BOSS:
		if !this.GetSystem().IsCross() || (this.GetSystem().IsCross() && isFromCross) {
			if this.hellBossInfos[stageId] != nil {
				this.hellBossInfos[int(bossInfo.StageId)] = bossInfo
				var reliveTime int64 = 0
				nowTime := time.Now().Unix()
				if nowTime < bossInfo.ReliveTime {
					reliveTime = bossInfo.ReliveTime
				}
				this.BroadcastAll(&pb.HellBossNtf{
					StageId:    bossInfo.StageId,
					Blood:      bossInfo.Hp * 100,
					ReliveTime: reliveTime,
				})
			} else {
				logger.Error("战斗服发来的炼狱boss信息，game服没有记录:%v", bossInfo.StageId)
			}
		}
	}
	this.fightMu.Unlock()
}

func (this *Fight) HandlerFieldBossDieUserInfoNtf(fieldBossDieUserInfoNtf *pbserver.FsFieldBossDieUserInfoNtf) {

	this.fightMu.Lock()
	this.userDieTime[int(fieldBossDieUserInfoNtf.DieUserId)] = int(fieldBossDieUserInfoNtf.DieTime)
	this.fightMu.Unlock()
}

func (this *Fight) HandlerUserSkillUse(msg *pbserver.FsSkillUseNtf) {

	//userId := int(msg.UseId)
	//this.DispatchEvent(userId, msg, func(userId int, user *objs.User, data interface{}) {
	//	if user == nil {
	//		logger.Warn("接收到战斗服发送来的技能使用：%v，玩家不在线", *msg)
	//		return
	//	}
	//	msgData := data.(*pbserver.FsSkillUseNtf)
	//	//err := this.GetSkill().UpProficiency(user, int(msgData.HeroIndex), int(msgData.SkillId), msgData.CdStartTime, msgData.CdStopTime)
	//	//if err != nil {
	//	//	logger.Error("玩家使用技能，通知玩家异常：消息：%v,异常：%v", *msgData, err)
	//	//}
	//	if len(msgData.KillMonsterIds) > 0 {
	//		//通知任务系统，击杀了怪物
	//		for _, v := range msgData.KillMonsterIds {
	//			this.GetTask().UpdateTaskForKillMonster(user, int(v), 1)
	//		}
	//	}
	//	if msgData.KillUserNum > 0 {
	//		killUserNum := int(msgData.KillUserNum)
	//		this.GetWarOrder().WriteWarOrderTask(user, pb.WARORDERCONDITION_KILL_USER, []int{killUserNum})
	//		this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_KILL_USER, []int{killUserNum})
	//		this.GetUserManager().UpdateCombat(user, -1)
	//	}
	//})
}

func (this *Fight) ActorKillNtf(msg *pbserver.FsToGsActorKillNtf) {

	userId := int(msg.Killer)
	this.DispatchEvent(userId, msg, func(userId int, user *objs.User, data interface{}) {
		if user == nil {
			logger.Warn("接收到战斗服发送来的击杀消息：%v，玩家不在线", *msg)
			return
		}
		msgData := data.(*pbserver.FsToGsActorKillNtf)
		if !msgData.IsPlayer {
			//通知任务系统，击杀了怪物
			this.GetTask().UpdateTaskForKillMonster(user, int(msgData.BeKiller), 1)

		} else {
			this.GetWarOrder().WriteWarOrderTask(user, pb.WARORDERCONDITION_KILL_USER, []int{1})
			this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_KILL_USER, []int{1})
			this.GetUserManager().UpdateCombat(user, -1)
		}
	})
}

/**
 *  @Description:
 *  @param userId
 *  @param heroIndex
 */
func (this *Fight) ClearSkillCD(userId int, heroIndex int) {
	this.DispatchEvent(userId, nil, func(userId int, user *objs.User, data interface{}) {
		if user == nil {
			logger.Warn("接收到战斗服发送来的清理玩家技能Cd：%v，玩家不在线")
			return
		}

		this.GetUserManager().SendMessage(user, &pb.ClearSkillCdNtf{HeroIndex: int32(heroIndex)}, true)
	})
}

func (this *Fight) HandlerExpStageKillMonsterNtf(userId int) {

	this.DispatchEvent(userId, nil, func(userId int, user *objs.User, data interface{}) {
		if user == nil {
			logger.Warn("接收到战斗服发送来的经验副本击杀怪物：%v，玩家不在线")
			return
		}
		//this.GetExpStage().ExpStageDareNumNtf(user)
	})
}

func (this *Fight) ClientEnterPublicCopy(user *objs.User, stageId int, condition int) (nw.ProtoMessage, error) {

	publicCopyConf := gamedb.GetPublicCopyStageCfg(stageId)
	if publicCopyConf == nil {
		return nil, gamedb.ERRSETTINGNOTFOUND
	}

	costItemId := 0
	if publicCopyConf.ConditionType == pb.CONDITIONTYPE_ALL {

		if !this.GetCondition().CheckMultiByType(user, -1, publicCopyConf.Condition, publicCopyConf.ConditionType) {
			return nil, gamedb.ERRCONDITION
		}

	} else {

		if _, ok := publicCopyConf.Condition[condition]; !ok {
			return nil, gamedb.ERRCONDITION
		}
		if _, ok := this.GetCondition().Check(user, -1, condition, publicCopyConf.Condition[condition]); !ok {
			return nil, gamedb.ERRCONDITION
		}

		if condition == pb.CONDITION_COST_CHUAN_QI_BI {
			costItemId = pb.ITEMID_CHUANQI_BI
		} else if condition == pb.CONDITION_COST_INGOT {
			costItemId = pb.ITEMID_INGOT
		} else if condition == pb.CONDITION_COST_GOLD {
			costItemId = pb.ITEMID_GOLD
		}
	}

	if costItemId > 0 {
		op := ophelper.NewOpBagHelperDefault(constBag.OpTypeEnterPublicCopy)
		err := this.GetBag().Remove(user, op, costItemId, publicCopyConf.Condition[condition])
		if err != nil {
			return nil, err
		}
		this.GetUserManager().SendItemChangeNtf(user, op)
	}

	return this.EnterPublicCopy(user, stageId, false)
}

func (this *Fight) EnterPublicCopy(user *objs.User, stageId int, reliveToCity bool) (nw.ProtoMessage, error) {

	fightId := this.getResidentFightId(stageId)
	if fightId <= 0 {
		return nil, gamedb.ERRFIGHTID
	}

	err := this.enterFight(user, stageId, int(fightId), reliveToCity, 0)
	return nil, err
}

/**
 *  @Description: 记录玩家战斗信息 并推送gate
 *  @param user
 *  @param fightId
 */
func (this *Fight) recordUserFightInfo(user *objs.User, stageId int, fightId int, crossFightServerId int) {

	crossFightServerId = this.GetCrossFightServerId(stageId)
	//记录玩家离开关卡
	this.GetStageManager().LeaveStage(user)
	user.FightId = fightId
	user.FightStageId = stageId
	user.FightStartTime = time.Now()
	user.FightLessTimes = false
	req := &pbgt.UserFightInfoNtf{
		FightId:            int32(fightId),
		CrossFightServerId: int32(crossFightServerId),
	}
	//推送任务系统
	this.GetTask().UpdateTaskProcess(user, false, false)
	//推送gate,当前战斗Id
	this.GetUserManager().SendMessage(user, req, true)
}

/**
*  @Description: 进入战斗前检查
*  @receiver this
*  @param user
*  @param stageId
*  @return bool
**/
func (this *Fight) CheckInFightBefore(user *objs.User, stageId int) bool {

	stageConf := gamedb.GetStageStageCfg(stageId)
	if stageConf == nil {
		return false
	}
	if user.FightStageId > 0 && user.FightStageId == stageId {
		return false
	}
	return true
}

func (this *Fight) checkFightIsExistByFightId(fightId int, stageId int) bool {

	applyMsg := &pbserver.GSTOFSCheckFightReq{
		FightId: int32(fightId),
	}
	replayMsg := &pbserver.FSTOGSCheckFightAck{}
	err := this.FSRpcCall(-1, stageId, applyMsg, replayMsg)
	if err != nil {
		logger.Error("检查战斗是否存在异常：%v", err)
		return false
	}
	if replayMsg.FightId > 0 {
		return true
	}
	return false
}

func (this *Fight) Gm(user *objs.User, codes string) string {
	applyMsg := &pbserver.GsToFsGmReq{
		UserId: int32(user.Id),
		Cmd:    codes,
	}
	replayMsg := &pbserver.FsToGsGmAck{}
	err := this.FSRpcCall(user.FightId, user.FightStageId, applyMsg, replayMsg)
	if err != nil {
		logger.Error("检查战斗是否存在异常：%v", err)
		return "fail"
	}
	return replayMsg.Result
}

func (this *Fight) UseCutTreasure(user *objs.User) error {

	stageConf := gamedb.GetStageStageCfg(user.FightStageId)
	if stageConf == nil {
		return nil
	}
	mapTypeConf := gamedb.GetMaptypeGameCfg(stageConf.Type)
	if len(mapTypeConf.CanUseDDQG) == 0 || mapTypeConf.CanUseDDQG[0] == 0 {
		return gamedb.ERRSKILLCANNOTUSE
	}

	applyMsg := &pbserver.GsToFsUseCutTreasureReq{
		UserId:        int32(user.Id),
		CutTreasureLv: int32(user.CutTreasure),
	}
	replayMsg := &pbserver.FsToGsUseCutTreasureAck{}
	err := this.FSRpcCall(user.FightId, user.FightStageId, applyMsg, replayMsg)
	if err != nil {
		logger.Error("检查战斗是否存在异常：%v", err)
		return err
	}
	return nil

}
