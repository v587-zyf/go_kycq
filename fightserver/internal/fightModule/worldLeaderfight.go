package fightModule

import (
	"cqserver/fightserver/internal/actorPkg"
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

/*世界首领*/
type WorldLeaderFight struct {
	*DefaultFight
	*FightWorldLeaderDamageRank
	*base.FightCheerByGuild
	*base.FightUsePotion
	boss              *actorPkg.MonsterActor
	worldLeaderConf   *gamedb.WorldLeaderConfCfg
	fightStatus       int //0初始化，1 即将开始 2开始 3 结束
	lastDecBossHpTime time.Time
	lastAttacker      int //最后攻击玩家Id
}

func NewWorldLeaderFight(stageId int) (*WorldLeaderFight, error) {
	var err error
	f := &WorldLeaderFight{
		FightUsePotion:    base.NewFightUsePotion(),
		FightCheerByGuild: base.NewFightCheerByGuild(),
	}
	f.DefaultFight, err = NewDefaultFight(stageId, f)
	if err != nil {
		return nil, err
	}
	f.FightWorldLeaderDamageRank = NewFightWorldLeaderDamageRank(f)
	f.SetLifeTime(-1)
	f.InitMonster()
	worldLeaderConf := gamedb.GetWorldLeaderByStageId(stageId)
	if worldLeaderConf == nil {
		logger.Info("世界首领配置获取失败")
		return nil, gamedb.ERRSETTINGNOTFOUND
	}
	f.worldLeaderConf = worldLeaderConf
	//只有一个怪物boss
	for _, v := range f.monsterActors {
		f.boss = v.(*actorPkg.MonsterActor)
	}
	f.fightStatus = constFight.WOWLD_BOSS_STATUS_INIT
	f.Start()
	return f, nil
}

func (this *WorldLeaderFight) PostDamage(attacker, defender base.Actor, damage int) {
	if defender.GetType() == base.ActorTypeMonster && attacker.GetUserId() != 0 {
		this.lastAttacker = attacker.GetUserId()
	}
}

func (this *WorldLeaderFight) OnActorEnter(actor base.Actor) {

	this.DefaultFight.OnActorEnter(actor)
	this.FightCheerByGuild.OnGuildActorEnter(actor)

}

func (this *WorldLeaderFight) OnEnterUser(userId int) {

	mainActor := this.GetUserMainActor(userId)
	this.FightDamageRankSetDamage(mainActor, 0)
	this.FightCheerByGuild.FightCheerUserInto(userId)
}

func (this *WorldLeaderFight) OnEnd() {

	this.boss.GetFSM().Stop()
	rank := this.FightWorldLeaderRankInfos(true)
	if len(rank.Ranks) > 0 {

		msg := &pbserver.FSFightEndNtf{
			FightType: int32(this.StageConf.Type),
			StageId:   int32(this.StageConf.Id),
		}

		resultNtf := &pbserver.WorldLeaderFightEndNtf{
			StageId:      int32(this.StageConf.Id),
			LastAttacker: int32(this.lastAttacker),
			Ranks:        rank.Ranks,
		}
		for _, v := range resultNtf.Ranks {
			if this.guildDamageUsers[int(v.GuildId)] == nil {
				v.Users = make([]int32, 0)
				continue
			}
			for userId := range this.guildDamageUsers[int(v.GuildId)] {
				v.Users = append(v.Users, int32(userId))
			}
		}
		rb, _ := resultNtf.Marshal()
		msg.CpData = rb
		logger.Info("推送世界首领战斗结果：%v", msg)
		net.GetGsConn().SendMessage(msg)
	}
	this.FightWorldLeaderDamageRank.FightDamageRankReset()
}

func (this *WorldLeaderFight) checkWorldBossStatus() {

	dayZeroTime := common.GetTimeSeconds(time.Now())
	startTime := this.worldLeaderConf.Time[0].GetSecondsFromZero()
	stopTime := this.worldLeaderConf.Time[1].GetSecondsFromZero()
	switch this.fightStatus {

	case constFight.WOWLD_BOSS_STATUS_INIT:
		if dayZeroTime >= startTime && dayZeroTime < stopTime {
			//开始
			this.fightStatus = constFight.WOWLD_BOSS_STATUS_RUN
			this.FightCheerByGuild = base.NewFightCheerByGuild()
			this.FightUsePotion = base.NewFightUsePotion()
			this.boss.GetProp().SetHpNow(this.boss.GetProp().Get(pb.PROPERTY_HP))
			this.boss.GetFSM().Recover()
			this.announcement()
		}
	case constFight.WOWLD_BOSS_STATUS_RUN:
		if dayZeroTime > stopTime {
			this.fightStatus = constFight.WOWLD_BOSS_STATUS_INIT
		}
	}
}

func (this *WorldLeaderFight) announcement() {

	ntf := &pbserver.WorldBossStatusNtf{
		StageId: int32(this.StageConf.Id),
		Status:  int32(this.fightStatus),
	}
	net.GetGsConn().SendMessage(ntf)
}

func (this *WorldLeaderFight) UpdateFrame() {

	this.checkWorldBossStatus()
	//单位时间内有伤害，扣血
	if this.fightStatus == constFight.WOWLD_BOSS_STATUS_RUN && time.Now().Sub(this.lastDecBossHpTime) > time.Second {
		this.lastDecBossHpTime = time.Now()
		dayZeroTime := common.GetTimeSeconds(time.Now())
		startTime := this.worldLeaderConf.Time[0].GetSecondsFromZero()
		stopTime := this.worldLeaderConf.Time[1].GetSecondsFromZero()
		continueTime := stopTime - startTime
		scale := float64(stopTime-dayZeroTime) / float64(continueTime)
		hasPlayer := false
		if this.GetPlayerNum() > 0 {
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
			hasPlayer = true
		}

		if dayZeroTime >= stopTime || this.boss.GetProp().HpNow() <= 0 {
			this.OnEnd()
		} else {
			if hasPlayer {
				//推送所有game战斗排行榜
				rank := this.FightWorldLeaderRankInfos(true)
				rank.StageId = int32(this.StageConf.Id)
				rank.BossHp = int32(scale * 100)
				net.GetGsConn().SendMessage(rank)
			}
		}
	}
}

func (this *WorldLeaderFight) OnCheer(userId int) {

	userMainActor := this.GetUserMainActor(userId)
	if userMainActor == nil {
		logger.Error("玩家发送来鼓舞，鼓舞玩家信息未找到：%v", userId)
		return
	}
	guildId := userMainActor.(base.ActorUser).GuildId()
	buffId := gamedb.GetConf().WorldLeaderBuff[3]
	this.FightCheerByGuild.GuildCheer(this, guildId, buffId)
	actors := this.GetUserActors()
	for _, v := range actors {
		if u, ok := v.(base.ActorUser); ok {
			if u.GuildId() == guildId {
				v.AddBuff(gamedb.GetConf().WorldLeaderBuff[3], userMainActor, false)
			}
		}
	}
}

func (this *WorldLeaderFight) OnUsePotion(userId int) {

	userActors := this.GetUserByUserId(userId)
	if userActors == nil {
		logger.Error("玩家发送来药水使用，玩家信息未找到：%v", userId)
		return
	}

	for _, v := range userActors {
		//推送血量变化
		changeHp, _ := v.ChangeHp(int(float64(gamedb.GetConf().ShabakePotion[2]) / 100 * float64(v.GetProp().Get(pb.PROPERTY_HP))))
		HPChangeNtf := &pb.SceneObjHpNtf{
			ObjId:    int32(v.GetObjId()),
			Hp:       int64(v.GetProp().HpNow()),
			ChangeHp: int64(changeHp),
			TotalHp:  int64(v.GetProp().Get(pb.PROPERTY_HP)),
		}
		v.NotifyNearby(v, HPChangeNtf, nil)
	}
}

//
//  @Description: 战斗伤害排行
// 获取排行版时进行排序，可根据参数强制排序，默认间隔1000毫秒
//
type GuildRankInfo struct {
	GuildId   int
	GuildName string
	ServerId  int
}
type FightWorldLeaderDamageRank struct {
	f                base.Fight
	guild            map[int]*GuildRankInfo
	damageRank       *common.SortedMap
	guildDamageUsers map[int]map[int]bool
	lastSortTime     time.Time
	sortRankInterval int64 //排名间隔时间（毫秒）
}

func NewFightWorldLeaderDamageRank(f base.Fight) *FightWorldLeaderDamageRank {
	fightRank := &FightWorldLeaderDamageRank{
		f:     f,
		guild: make(map[int]*GuildRankInfo),
		damageRank: &common.SortedMap{
			M: make(map[int]int64),
		},
		guildDamageUsers: make(map[int]map[int]bool),
		sortRankInterval: 1000,
	}
	return fightRank
}

func (fd *FightWorldLeaderDamageRank) FightDamageRankReset() {
	fd.guild = make(map[int]*GuildRankInfo)
	fd.damageRank = &common.SortedMap{
		M: make(map[int]int64),
	}
	fd.guildDamageUsers = make(map[int]map[int]bool)
}

func (fd *FightWorldLeaderDamageRank) FightDamageRankSetRankInterval(interval int) {
	fd.sortRankInterval = int64(interval)
}

func (fd *FightWorldLeaderDamageRank) FightDamageRankSetDamage(actor base.Actor, value int64) {
	mainActor := fd.f.GetUserMainActor(actor.GetUserId())
	if mainActor == nil {
		return
	}
	guildId := 0
	guildName := ""
	serverId := mainActor.HostId()
	if user, ok := mainActor.(base.ActorUser); ok {
		guildId = user.GuildId()
		guildName = user.GuildName()
	} else {
		return
	}

	//玩家第一造成伤害或者玩家退出游戏 重新进入游戏
	if _, ok := fd.guild[guildId]; !ok {
		fd.guild[guildId] = &GuildRankInfo{
			guildId,
			guildName,
			serverId,
		}
	}
	if fd.guildDamageUsers[guildId] == nil {
		fd.guildDamageUsers[guildId] = make(map[int]bool)
	}
	fd.guildDamageUsers[guildId][actor.GetUserId()] = true
	fd.damageRank.M[guildId] += value
}

func (fd *FightWorldLeaderDamageRank) FightDamageRankGetRank(isForce bool) []int {
	if isForce || fd.lastSortTime.IsZero() || (time.Now().Sub(fd.lastSortTime).Milliseconds()) > fd.sortRankInterval {

		return fd.damageRank.Sort()
	}
	return fd.damageRank.GetNowRank()
}

func (fd *FightWorldLeaderDamageRank) FightDamageRankGetRankInfos(actor base.Actor) nw.ProtoMessage {
	guildId := 0
	if actor != nil {
		guildId = actor.(base.ActorUser).GuildId()
	}
	rank := fd.FightDamageRankGetRank(false)
	ack := &pb.FightHurtRankAck{
		Ranks: make([]*pb.FightRankUnit, len(rank)),
	}
	for k, v := range rank {
		ack.Ranks[k] = &pb.FightRankUnit{
			Rank:       int32(k + 1),
			Name:       fd.guild[v].GuildName,
			Score:      fd.damageRank.M[v],
			ServerName: fightManager.GetServerName(fd.guild[v].ServerId),
		}
		if v == guildId {
			ack.MyRank = ack.Ranks[k]
		}
	}
	return ack
}

func (fd *FightWorldLeaderDamageRank) FightWorldLeaderRankInfos(rankSortForce bool) *pbserver.WorldLeaderFightRankNtf {

	rank := fd.FightDamageRankGetRank(rankSortForce)
	ack := &pbserver.WorldLeaderFightRankNtf{
		Ranks: make([]*pbserver.WorldLeaderRankUnit, len(rank)),
	}
	for k, v := range rank {
		ack.Ranks[k] = &pbserver.WorldLeaderRankUnit{
			Rank:      int32(k + 1),
			GuildId:   int32(v),
			GuildName: fd.guild[v].GuildName,
			Score:     fd.damageRank.M[v],
			ServerId:  int32(fd.guild[v].ServerId),
		}
	}
	return ack

}
