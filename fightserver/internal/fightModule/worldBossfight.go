package fightModule

import (
	"cqserver/fightserver/internal/actorPkg"
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"math/rand"
	"strconv"
	"time"
)

/*世界boss*/
type WorldBossFight struct {
	*DefaultFight
	*base.FightDamageRank
	boss              *actorPkg.MonsterActor
	worlbBossConf     *gamedb.WorldBossWorldBossCfg
	fightStatus       int //0初始化，1 即将开始 2开始 3 结束
	lastDecBossHpTime time.Time
}

func NewWorldBossFight(stageId int) (*WorldBossFight, error) {
	var err error
	worldBossFight := &WorldBossFight{}
	worldBossFight.DefaultFight, err = NewDefaultFight(stageId, worldBossFight)
	if err != nil {
		return nil, err
	}
	worldBossFight.FightDamageRank = base.NewFightDamageRank()
	worldBossFight.SetLifeTime(-1)
	worldBossFight.InitMonster()
	worldBossConf := gamedb.GetWorldBossConfByStageId(stageId)
	if worldBossConf == nil {
		logger.Info("世界boss配置获取失败")
		return nil, gamedb.ERRSETTINGNOTFOUND
	}
	worldBossFight.worlbBossConf = worldBossConf
	//只有一个怪物boss
	for _, v := range worldBossFight.monsterActors {
		worldBossFight.boss = v.(*actorPkg.MonsterActor)
	}
	worldBossFight.fightStatus = constFight.WOWLD_BOSS_STATUS_INIT
	worldBossFight.Start()
	return worldBossFight, nil
}

func (this *WorldBossFight) OnDie(actor base.Actor, killer base.Actor) {
	if actor.GetType() == base.ActorTypeMonster {
		this.OnEnd()
	}
}

func (this *WorldBossFight) OnEnd() {

	rank := this.FightDamageRankGetRank(true)
	if len(rank) > 0 {

		luckUserId := rank[rand.Intn(len(rank))]
		msg := &pbserver.FSFightEndNtf{
			FightType: int32(this.StageConf.Type),
			StageId:   int32(this.StageConf.Id),
			Winners:   common.ConvertIntSlice2Int32Slice(rank),
			CpData:    []byte(strconv.Itoa(luckUserId)),
		}
		net.GetGsConn().SendMessage(msg)
	}
	this.FightDamageRank.FightDamageRankReset()
}

func (this *WorldBossFight) checkWorldBossStatus() {

	dayZeroTime := common.GetTimeSeconds(time.Now())
	prepareTime := this.worlbBossConf.PrepareTime.GetSecondsFromZero()
	startTime := this.worlbBossConf.OpenTime.GetSecondsFromZero()
	stopTime := this.worlbBossConf.OpenTime.GetSecondsFromZero() + this.worlbBossConf.Continue*60
	switch this.fightStatus {

	case constFight.WOWLD_BOSS_STATUS_INIT:
		if dayZeroTime > prepareTime && dayZeroTime < startTime {
			//即将开始 推送即将开始公告
			this.fightStatus = constFight.WOWLD_BOSS_STATUS_READY
			this.announcement(constFight.WOWLD_BOSS_STATUS_READY)
		} else if dayZeroTime >= startTime && dayZeroTime < stopTime {
			//开始
			this.fightStatus = constFight.WOWLD_BOSS_STATUS_RUN
			this.announcement(constFight.WOWLD_BOSS_STATUS_RUN)
		}
	case constFight.WOWLD_BOSS_STATUS_READY:
		if dayZeroTime >= startTime && dayZeroTime < stopTime {
			//开始
			this.fightStatus = constFight.WOWLD_BOSS_STATUS_RUN
			this.announcement(constFight.WOWLD_BOSS_STATUS_RUN)
		}
	case constFight.WOWLD_BOSS_STATUS_RUN:
		if dayZeroTime > stopTime {
			this.fightStatus = constFight.WOWLD_BOSS_STATUS_INIT
		}
	}
}

func (this *WorldBossFight) announcement(status int) {

	ntf := &pbserver.WorldBossStatusNtf{
		StageId: int32(this.StageConf.Id),
		Status:  int32(constFight.WOWLD_BOSS_STATUS_READY),
	}
	net.GetGsConn().SendMessage(ntf)
}

func (this *WorldBossFight) UpdateFrame() {

	//单位时间内有伤害，扣血
	if this.fightStatus == constFight.WOWLD_BOSS_STATUS_RUN && time.Now().Sub(this.lastDecBossHpTime) > time.Second {
		this.lastDecBossHpTime = time.Now()
		dayZeroTime := common.GetTimeSeconds(time.Now())
		startTime := this.worlbBossConf.OpenTime.GetSecondsFromZero()
		continueTime := this.worlbBossConf.Continue * 60
		scale := float64(continueTime-(dayZeroTime-startTime)) / float64(continueTime)
		if scale <= 0 {
			scale = 0
		}
		this.boss.GetProp().SetHpNow(int(float64(this.boss.GetProp().Get(pb.PROPERTY_HP)) * scale))
		//通知客户端
		ntf := &pb.SceneObjHpNtf{
			ObjId: int32(this.boss.GetObjId()),
			Hp:    int64(this.boss.GetProp().HpNow()),
		}
		this.GetScene().NotifyAll(ntf)
		if this.boss.GetProp().HpNow() <= 0 {
			this.OnEnd()
		}
	}
	this.checkWorldBossStatus()
}
