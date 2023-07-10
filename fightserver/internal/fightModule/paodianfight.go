package fightModule

import (
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/gamelibs/gamedb"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"time"
)

type PaodianFight struct {
	*DefaultFight
	userEnterTime     map[int]int64 //玩家进入时间
	lastCountTime     int64
	paodianRewardConf *gamedb.PaoDianRewardPaoDianRewardCfg
	topUser           map[int]bool
}

func NewPaodianFight(stageId int) (*PaodianFight, error) {
	var err error
	f := &PaodianFight{
		userEnterTime: make(map[int]int64),
		lastCountTime: time.Now().Unix(),
		topUser:       make(map[int]bool),
	}
	paodianReardConf := gamedb.GetPaodianConfByStageId(stageId)
	if paodianReardConf == nil {
		return nil, gamedb.ERRSETTINGNOTFOUND
	}
	f.paodianRewardConf = paodianReardConf
	f.DefaultFight, err = NewDefaultFight(stageId, f)
	if err != nil {
		return nil, err
	}

	paodianDailyConf := gamedb.GetDailyActivityDailyActivityCfg(pb.DAILYACTIVITYTYPE_PAODIAN)
	lifeTime := paodianDailyConf.CloseTime.GetSecondsFromZero() - common.GetTimeSeconds(time.Now())
	f.SetLifeTime(int64(lifeTime))
	f.Start()

	return f, nil
}

func (this *PaodianFight) GetTopUser() *pb.PaodianTopUserNtf {
	msg := &pb.PaodianTopUserNtf{
		UserIds: make([]int32, 0),
	}
	for k, _ := range this.topUser {
		msg.UserIds = append(msg.UserIds, int32(k))
	}
	return msg
}

func (this *PaodianFight) UpdateFrame() {

	now := time.Now().Unix()
	if now-this.lastCountTime >= int64(this.paodianRewardConf.Interval) {
		this.lastCountTime = now
		//计算经验 增加经验 暂时game计算了
		ntf := &pbserver.PaodianGoodsAddNtf{
			StageId: int32(this.StageConf.Id),
			UserIds: make(map[int32]int32),
		}
		for k, _ := range this.topUser {
			isDie := this.CheckUserDieByUserId(k)
			if !isDie {
				ntf.UserIds[int32(k)] = int32(this.paodianRewardConf.Times)
			}
		}
		for k, _ := range this.userEnterTime {
			if this.topUser[k] {
				continue
			}
			isDie := this.CheckUserDieByUserId(k)
			if !isDie {
				ntf.UserIds[int32(k)] = int32(1)
			}
		}
		//logger.Debug("当前泡点：%v,玩家经验：%v", this.GetStageConf().Id, ntf.UserIds)
		net.GetGsConn().SendMessage(ntf)
	}
}

func (this *PaodianFight) OnDie(actor, killer base.Actor) {

	if actor.GetType() == pb.SCENEOBJTYPE_USER || actor.GetType() == pb.SCENEOBJTYPE_FIT {
		allDie := this.CheckUserAllDie(actor)
		if !allDie {
			return
		}
		userId := actor.GetUserId()
		if !this.topUser[userId] {
			return
		}
		delete(this.topUser, userId)
		if killer != nil {

			attackerActor := this.GetActorByObjId(killer.GetObjId())
			if attackerActor != nil {
				attackerUserId := attackerActor.GetUserId()
				if !this.topUser[attackerUserId] {
					this.addTopUser(attackerUserId)
				}
			}
		}
		this.addTopUser(-1)
	}
}

func (this *PaodianFight) OnEnterUser(userId int) {
	this.userEnterTime[userId] = time.Now().Unix()
	this.addTopUser(userId)
	fightManager.PaodianUserEnter()
}

func (this *PaodianFight) OnLeaveUser(userId int) {
	delete(this.userEnterTime, userId)
	if this.topUser[userId] {
		delete(this.topUser, userId)
		this.addTopUser(-1)
	}
	logger.Debug("玩家离开战斗，战斗stageId :%v，玩家%v,地图玩家：%v", this.GetStageConf().Id, userId, this.userEnterTime)
	fightManager.PaodianUserEnter()
}

func (this *PaodianFight) addTopUser(userId int) {

	if len(this.topUser) >= this.paodianRewardConf.TopUserNum {
		return
	}
	if userId > 0 {
		this.topUser[userId] = true
	} else {
		early := 0
		for userId, t := range this.userEnterTime {
			if this.topUser[userId] {
				continue
			}
			if early == 0 || this.userEnterTime[early] < t {
				early = userId
			}
		}
		if early > 0 {
			this.topUser[early] = true
		}
	}
	ntf := this.GetTopUser()
	this.Scene.NotifyAll(ntf)
}

func (this *PaodianFight) OnEnd() {

	msg := &pbserver.FSFightEndNtf{
		FightType: int32(this.StageConf.Type),
		StageId:   int32(this.StageConf.Id),
		UseTime:   int32(time.Now().Unix() - this.createTime),
	}
	logger.Info("发送game泡点战斗结束,服务器：%v，结果：%v", *msg)
	net.GetGsConn().SendMessage(msg)
	this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_CLOSING, 15)
}
