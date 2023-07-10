package worldLeader

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"encoding/json"
	"sync"
	"time"
)

type WorldLeader struct {
	util.DefaultModule
	managersI.IModule
	nowStageId       int //当前首领id
	stageHp          map[int]int
	stageRank        map[int][]*pbserver.WorldLeaderRankUnit
	guildEnterNumber map[int]map[int]map[int]bool
	Mu               sync.RWMutex
}

func NewWorldLeaderBoss(module managersI.IModule) *WorldLeader {
	w := &WorldLeader{IModule: module,
		stageHp:          make(map[int]int, 0),
		guildEnterNumber: make(map[int]map[int]map[int]bool, 0),
		stageRank:        make(map[int][]*pbserver.WorldLeaderRankUnit, 0),
	}
	return w
}

func (this *WorldLeader) Init() error {

	this.nowStageId = this.GetNowStageId()

	hp := rmodel.WorldLeader.GetWorldLeaderHpInfo(this.nowStageId)
	this.stageHp[this.nowStageId] = hp

	cfg := gamedb.GetWorldLeaderCfs()
	for _, data := range cfg {

		guildNum := rmodel.WorldLeader.GetWorldLeaderEnterGuilds(data.StageId)
		if guildNum == nil || len(guildNum) <= 0 {
			continue
		}
		for guildId := range guildNum {
			if this.guildEnterNumber[data.StageId] == nil {
				this.guildEnterNumber[data.StageId] = make(map[int]map[int]bool, 0)
			}
			if this.guildEnterNumber[data.StageId][guildId] == nil {
				this.guildEnterNumber[data.StageId][guildId] = make(map[int]bool)
			}

			guildNumbers := rmodel.WorldLeader.GetWorldLeaderEnterGuildNumber(data.StageId, guildId)
			if guildNumbers == nil || len(guildNumbers) <= 0 {
				continue
			}
			for _, userId := range guildNumbers {
				this.guildEnterNumber[data.StageId][guildId][userId] = true
			}
		}

		data1 := rmodel.WorldLeader.GetWorldLeaderRankInfo(data.StageId)
		info := make([]*pbserver.WorldLeaderRankUnit, 0)
		_ = json.Unmarshal([]byte(data1), &info)
		this.stageRank[data.StageId] = info
	}

	return nil
}

func (this *WorldLeader) LoadWorldLeader(user *objs.User, ack *pb.LoadWorldLeaderAck) error {
	if this.GetSystem().GetServerIndexCrossFsId(user.ServerId) <= 0 {
		logger.Error("世界首领活动未开启  this.GetSystem().GetServerIndexCrossFsId(user.ServerId:%v):%v", user.ServerId, this.GetSystem().GetServerIndexCrossFsId(user.ServerId))
		return nil
	}
	ack.NowStageId = int32(this.GetNowStageId())
	this.nowStageId = int(ack.NowStageId)
	ack.BossHp = int32(this.getStageHp(this.nowStageId))
	ack.GuildJoinNum = int32(this.getGuildEnterNumber(this.nowStageId, user.GuildData.NowGuildId))

	logger.Debug("nowStageId:%v  guildId:%v  num:%v  BossHp:%v", this.nowStageId, user.GuildData.NowGuildId, ack.GuildJoinNum, ack.BossHp)
	ack.WorldLeaderInfoByStage = make(map[int32]*pb.WorldLeaderInfo)
	data := gamedb.GetWorldLeaderCfs()
	logger.Debug("NowGuildId:%v", user.GuildData.NowGuildId)
	for _, info := range data {
		ack.WorldLeaderInfoByStage[int32(info.StageId)] = this.buildInfoByStage(info.StageId, user.GuildData.NowGuildId)
	}

	return nil
}

func (this *WorldLeader) buildInfoByStage(stageId, guildId int) *pb.WorldLeaderInfo {
	data := &pb.WorldLeaderInfo{}
	data.GuildJoinNum = int32(this.getGuildEnterNumber(stageId, guildId))
	return data
}

func (this *WorldLeader) WorldLeaderEnter(user *objs.User, stageId int, ack *pb.WorldLeaderEnterAck) error {
	err := this.GetCondition().CheckFunctionOpen(user, pb.FUNCTIONID_WORLD_LEADER)
	if err != nil {
		return err
	}
	if stageId <= 0 {
		return gamedb.ERRPARAM
	}

	if user.GuildData.NowGuildId <= 0 {
		return gamedb.ERRHAVENOGUILD
	}

	this.nowStageId = this.GetNowStageId()
	logger.Debug("WorldLeaderEnter userId:%v nowStageId:%v stageId:%v", user.Id, this.nowStageId, stageId)
	if stageId != this.nowStageId {
		if stageId > this.nowStageId {
			return gamedb.ERRACTIVITYNOTOPEN
		}
		if stageId < this.nowStageId {
			return gamedb.ERRACTIVITYCLOSE
		}
	}

	err = this.GetFight().EnterResidentFightByStageId(user, stageId, 0)
	if err != nil {
		logger.Error("进入 世界首领 失败  user:%v, stageId:%v err:%v", err, user, stageId, err)
		return err
	}

	this.setGuildEnterNumber(stageId, user.GuildData.NowGuildId, user.Id)

	ack.EnterState = true
	return nil
}

func (this *WorldLeader) setGuildEnterNumber(stageId, guildId, userId int) {
	defer this.Mu.Unlock()
	this.Mu.Lock()
	if this.guildEnterNumber[stageId] == nil {
		this.guildEnterNumber[stageId] = make(map[int]map[int]bool, 0)
	}
	if this.guildEnterNumber[stageId][guildId] == nil {
		this.guildEnterNumber[stageId][guildId] = make(map[int]bool)
	}
	if this.guildEnterNumber[stageId][guildId][userId] == false {
		rmodel.WorldLeader.SetWorldLeaderEnterGuild(stageId, guildId)
		rmodel.WorldLeader.SetWorldLeaderEnterGuildNumber(stageId, guildId, userId)
		this.guildEnterNumber[stageId][guildId][userId] = true
	}
}

func (this *WorldLeader) getGuildEnterNumber(stageId, guildId int) int {
	defer this.Mu.RUnlock()
	this.Mu.RLock()
	if this.guildEnterNumber[stageId] == nil {
		return 0
	}
	if this.guildEnterNumber[stageId][guildId] == nil {
		return 0
	}
	return len(this.guildEnterNumber[stageId][guildId])
}

func (this *WorldLeader) WorldLeaderRankInfo(user *objs.User, stageId int, ack *pb.GetWorldLeaderRankInfoAck) error {
	if stageId <= 0 {
		return gamedb.ERRPARAM
	}
	ack.BossHp = int32(this.getStageHp(stageId))
	ack.Ranks = this.buildStageRankInfo(stageId)
	ack.StageId = int32(stageId)
	return nil
}

func (this *WorldLeader) setStageRankInfos(stageId int, data []*pbserver.WorldLeaderRankUnit) {
	defer this.Mu.Unlock()
	this.Mu.Lock()
	if this.stageRank[stageId] == nil {
		this.stageRank[stageId] = make([]*pbserver.WorldLeaderRankUnit, 0)
	}
	this.stageRank[stageId] = data
}

func (this *WorldLeader) setStageHp(stageId, hp int) {
	defer this.Mu.Unlock()
	this.Mu.Lock()
	this.stageHp[stageId] = hp
}

func (this *WorldLeader) getStageHp(stageId int) int {
	defer this.Mu.RUnlock()
	this.Mu.RLock()

	return this.stageHp[stageId]
}

func (this *WorldLeader) getStageRankInfos(stageId int) []*pbserver.WorldLeaderRankUnit {
	defer this.Mu.RUnlock()
	this.Mu.RLock()
	return this.stageRank[stageId]
}

// 推送在线玩家boss开启
func (this *WorldLeader) SendClientWorldLeaderStart(stageId int) {
	//if this.GetSystem().GetServerIndexCrossFsId(base.Conf.ServerId) <= 0 {
	//	return
	//}
	cfgs := gamedb.GetWorldLeaderCfs()
	for _, data := range cfgs {
		if data.StageId == stageId {
			this.nowStageId = stageId
			this.setStageHp(stageId,100)
			this.BroadcastAll(&pb.WorldLeaderStartNtf{StageId: int32(this.nowStageId)})
			this.GetAnnouncement().SendSystemChat(nil, pb.SCROLINGTYPE_WORLD_LEADER_OPEN, -1, stageId)
			logger.Info("世界首领战斗开启 stage:%v", this.nowStageId)
			return
		}
	}
}

// 推送在线玩家门派boss结束
func (this *WorldLeader) sendClientFightEnd(stageId, userId, guildId, lastKillUser, rank int32) {
	ntf := &pb.WorldLeaderEndRewardNtf{}
	ntf.StageId = stageId
	userInfo := this.GetUserManager().GetUser(int(userId))
	if userInfo == nil {
		logger.Error("userInfo nil  userId:%v", userId)
		return
	}

	if lastKillUser == userId {
		ntf.Owner = this.GetUserManager().BuilderBrieUserInfo(int(userId))
	}
	ntf.Rank = rank
	logger.Debug("sendClientFightEnd stageId:%v, userId:%v, guildId:%v, lastKillUser:%v, rank:%v", stageId, userId, guildId, lastKillUser, rank)
	_ = this.GetUserManager().SendMessage(userInfo, ntf, true)
	this.GetCondition().RecordCondition(userInfo, pb.CONDITION_ALL_KILL_STAGE, []int{int(stageId), 1})
	this.GetCondition().RecordCondition(userInfo, pb.CONDITION_ALL_KILL_WORLDLEADER, []int{})
}

func (this *WorldLeader) buildStageRankInfo(stageId int) []*pb.WorldLeaderRankUnit {
	rankData := make([]*pb.WorldLeaderRankUnit, 0)
	data := this.getStageRankInfos(stageId)
	if data == nil {
		return rankData
	}
	for _, info := range data {
		rankData = append(rankData, &pb.WorldLeaderRankUnit{
			Rank:       info.Rank,
			GuildId:    info.GuildId,
			GuildName:  info.GuildName,
			Score:      info.Score,
			ServerId:   info.ServerId,
			ServerName: this.GetSystem().GetServerName(int(info.ServerId)),
		})
	}
	return rankData

}

func (this *WorldLeader) GetNowStageId() int {
	cfg := gamedb.GetWorldLeaderCfs()
	for _, data := range cfg {
		if data == nil {
			continue
		}
		if len(data.Time) < 2 {
			logger.Error("GetWorldLeaderCfs data.Id:%v  配置错误", data.Id)
			continue
		}
		before := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), data.Time[0].Hour, data.Time[0].Minute, data.Time[0].Second, 0, time.Local)
		after := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), data.Time[1].Hour, data.Time[1].Minute, data.Time[1].Second, 0, time.Local)
		if data.Id == 1 {
			logger.Debug("time.Now().Unix():%v < before.Unix():%v", time.Now().Unix(), before.Unix())
			if time.Now().Unix() < before.Unix() {
				return 0
			}
		}

		if data.Id == len(cfg) {
			if time.Now().Unix() > after.Unix() {
				return 0
			}
		}

		logger.Debug("id:%v before.Unix():%v  after.Unix():%v", data.Id, before.Unix(), after.Unix())
		if time.Now().Unix() >= before.Unix() && time.Now().Unix() < after.Unix() {
			return data.StageId
		}
	}
	return 0
}

func (this *WorldLeader) Reset() {
	defer this.Mu.Unlock()
	this.Mu.Lock()
	this.nowStageId = 0
	this.stageHp = make(map[int]int, 0)
	this.guildEnterNumber = make(map[int]map[int]map[int]bool, 0)
	this.stageRank = make(map[int][]*pbserver.WorldLeaderRankUnit, 0)
}
