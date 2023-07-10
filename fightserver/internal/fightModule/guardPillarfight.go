package fightModule

import (
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"time"
)

const FIGHT_INFO_BROAD_INTERVAL = 1000

type GuardPillarFight struct {
	*DefaultFight
	*base.FightCheerByGuild
	*base.FightDamageRank
	passWave        int //当前打完第几波
	nowWaveLess     int //当前波剩余数量
	nowWaveMonsters int //当前波剩余总数
	nextMonstTime   int64
	longzhuActor    base.Actor
	rankBoradTime   time.Time
	guildId         int
}

func NewGuardPillarFight(stageId int, guildId int) (*GuardPillarFight, error) {
	var err error
	f := &GuardPillarFight{
		FightCheerByGuild: base.NewFightCheerByGuild(),
		guildId:           guildId,
	}
	f.DefaultFight, err = NewDefaultFight(stageId, f)
	if err != nil {
		return nil, err
	}
	f.FightDamageRank = base.NewFightDamageRank()
	f.DefaultFight.InitMonsterByWave(0)
	//当前一定只有一个龙柱怪物
	for _, v := range f.monsterActors {
		//设置龙柱队伍为1，跟玩家同队
		v.SetTeamIndex(constFight.FIGHT_TEAM_ONE)
		f.longzhuActor = v
	}

	f.InitCollection()
	f.Start()
	f.nextMonstTime = time.Now().Unix() + int64(gamedb.GetConf().StartTime)
	logger.Info("创建守卫龙柱战斗成功")
	return f, nil
}

func (this *GuardPillarFight) GetFightExtMark() int {
	return this.guildId
}

func (this *GuardPillarFight) UpdateFrame() {

	if this.status != FIGHT_STATUS_RUNNING {
		return
	}
	if this.nextMonstTime <= 0 {
		return
	}
	if time.Now().Unix() > this.nextMonstTime {
		this.InitMonsterByWave()
		this.broadFightInfo()
	}
}

func (this *GuardPillarFight) PostDamage(attacker, defender base.Actor, damage int) {

	if attacker.GetType() != pb.SCENEOBJTYPE_MONSTER {
		if time.Now().Sub(this.rankBoradTime).Milliseconds() > FIGHT_INFO_BROAD_INTERVAL {
			this.broadFightInfo()
		}
	}
}

func (this *GuardPillarFight) InitMonsterByWave() {
	monsterNum, _ := this.DefaultFight.InitMonsterByWave(this.passWave + 1)
	this.nowWaveLess = monsterNum
	this.nowWaveMonsters = monsterNum
	this.nextMonstTime = 0
}

func (this *GuardPillarFight) OnDie(actor, killer base.Actor) {

	if actor.GetType() == pb.SCENEOBJTYPE_MONSTER {

		if actor.GetObjId() == this.longzhuActor.GetObjId() {
			//龙柱被打掉了，战斗结束
			this.OnEnd()
		} else {
			this.nowWaveLess--
			if this.nowWaveLess == 0 {
				this.passWave++
				if this.passWave+1 == len(this.StageConf.Monster_group) {
					this.OnEnd()
				} else {
					this.nextMonstTime = time.Now().Unix() + int64(gamedb.GetConf().NextTime)
					this.broadFightInfo()
				}
			}
			logger.Debug("龙柱守卫 怪物死亡，当前波数：%v,剩余数量：%v,下一波刷新时间：%v", this.passWave+1, this.nowWaveLess, this.nextMonstTime)
		}
	}
}

func (this *GuardPillarFight) OnActorEnter(actor base.Actor) {

	this.DefaultFight.OnActorEnter(actor)
	this.FightCheerByGuild.OnGuildActorEnter(actor)
}

func (this *GuardPillarFight) OnEnterUser(userId int) {

	this.FightCheerByGuild.FightCheerUserInto(userId)
	mainActor := this.GetUserMainActor(userId)
	this.FightDamageRankSetDamage(mainActor, 0)
}

func (this *GuardPillarFight) OnEnd() {

	msg := &pbserver.FSFightEndNtf{
		FightType: int32(this.StageConf.Type),
		StageId:   int32(this.StageConf.Id),
		UseTime:   int32(time.Now().Unix() - this.createTime),
	}

	rank := this.FightDamageRank.FightDamageRankGetRank(true)
	endMsg := &pbserver.GuardPillarFightEnd{
		Wave:  int32(this.passWave),
		Users: common.ConvertIntSlice2Int32Slice(rank),
	}
	rb, _ := endMsg.Marshal()
	msg.CpData = rb
	logger.Info("发送game龙柱战斗结束,服务器：%v，结果：%v", *msg)
	net.GetGsConn().SendMessage(msg)
	this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_CLOSING, 15)
}

func (this *GuardPillarFight) OnCheer(userId int) {

	userMainActor := this.GetUserMainActor(userId)
	if userMainActor == nil {
		logger.Error("玩家发送来鼓舞，鼓舞玩家信息未找到：%v", userId)
		return
	}
	guildId := userMainActor.(base.ActorUser).GuildId()
	buffId := gamedb.GetConf().GuardBuff[3]
	//记录门派鼓舞数据
	this.FightCheerByGuild.GuildCheer(this, guildId, buffId)
	actors := this.GetUserActors()
	for _, v := range actors {
		v.AddBuff(gamedb.GetConf().GuardBuff[3], userMainActor, false)
	}
}

func (this *GuardPillarFight) OnCollection(colllection map[int]int) {

}

func (this *GuardPillarFight) FightDamageRankGetRankInfos(actor base.Actor) nw.ProtoMessage {
	ntf := &pb.GuardPillarFightNtf{}
	ntf.Rank = this.FightDamageRank.FightDamageRankGetRankInfos(actor)
	ntf.Wave = int32(this.passWave)
	if this.nextMonstTime == 0 {
		ntf.Wave += 1
	}
	ntf.NextTime = int32(this.nextMonstTime)
	ntf.MonsterTotal = int32(this.nowWaveMonsters)
	ntf.Monsterless = int32(this.nowWaveLess)
	ntf.FightEndTime = int32(this.lifeTime - (time.Now().Unix() - this.createTime))
	return ntf
}

func (this *GuardPillarFight) broadFightInfo() {
	ntf := this.FightDamageRankGetRankInfos(nil)
	this.GetScene().NotifyAll(ntf)
}
