package fightModule

import (
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/gamelibs/gamedb"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"fmt"
	"time"
)

type ExpFight struct {
	*DefaultFight
	*base.FightReady
	nowWave             int //当前第几波
	nowWaveLess         int //当前波剩余数量
	fightResult         bool
	killMonsterNum      int //击杀怪物数量
	exp                 int //获得经验数量
	killMonsterSendUser bool
	fightOwner          base.Actor
}

func NewExpFight(stageId int) (*ExpFight, error) {
	var err error
	expFight := &ExpFight{
		nowWave:             0,
		killMonsterSendUser: false,
	}
	expFight.DefaultFight, err = NewDefaultFight(stageId, expFight)
	if err != nil {
		return nil, err
	}
	//expFight.InitMonsterByWave(expFight.nowWave)
	expFight.FightReady = base.NewFightReady()
	expFight.Start()

	return expFight, nil
}

func (this *ExpFight) OnDie(actor base.Actor, killer base.Actor) {

	if actor.GetType() == base.ActorTypeMonster {

		this.killMonsterNum++
		this.nowWaveLess--
		if !this.killMonsterSendUser {

			msg := &pbserver.ExpStageKillMonsterNtf{
				UserId: int32(this.fightOwner.GetUserId()),
			}
			net.GetGsConn().SendMessageToGs(uint32(this.fightOwner.HostId()), msg)
			this.killMonsterSendUser = true
		}
		ntf := &pb.ExpStageKillInfoNtf{
			KillMonsterNum: int32(this.killMonsterNum),
			GetExp:         int32(this.exp),
		}
		if u, ok := this.fightOwner.(base.ActorUser); ok {
			u.SendMessage(ntf)
		}
		logger.Debug("当前第几波：%v,剩余怪物数量：%v,击杀：%v,玩家：%v,限制：%v", this.nowWave, this.nowWaveLess, this.killMonsterNum, actor.GetObjId(), gamedb.GetConf().ExpStageRefreshMonster)
		if this.nowWaveLess <= gamedb.GetConf().ExpStageRefreshMonster {

			if this.nowWaveLess == 0 && this.nowWave == len(this.StageConf.Monster_group)-1 {

				this.fightResult = true
				this.OnEnd()
			} else if this.nowWave < len(this.StageConf.Monster_group)-1 {
				this.nowWave += 1

				//for _, monster := range this.monsterActors {
				//	if monster.GetObjId() != actor.GetObjId() && monster.GetProp().HpNow() <= 0 {
				//		this.Leave(monster)
				//	}
				//}
				this.InitMonsterByWave()
			}
		}

	} else {
		allDie := this.CheckUserAllDie(actor)
		if allDie {
			this.fightResult = false
			this.OnEnd()
		}
	}
}

func (this *ExpFight) MonsterDrop(dropMonsterId int, dropX, dropY int, owner base.Actor, dropItems []*pbserver.ItemUnit) {

	//this.DefaultFight.MonsterDrop(dropMonsterId, dropX, dropY, owner, dropItems)
	for _, v := range dropItems {
		if v.ItemId == pb.ITEMID_EXP {
			this.exp += int(v.ItemNum)
		}
	}
}

func (this *ExpFight) OnActorEnter(actor base.Actor) {
	this.DefaultFight.OnActorEnter(actor)
	if actor.GetType() == pb.SCENEOBJTYPE_USER && actor.HostId() > 0 {
		this.fightOwner = actor
	}
}

func (this *ExpFight) OnLeaveUser(userId int) {
	this.Stop()
}

func (this *ExpFight) OnEnd() {
	if this.status != FIGHT_STATUS_RUNNING {
		return
	}
	//经验副本里面肯定只有一个人
	msg := &pbserver.FSFightEndNtf{
		FightType: int32(this.StageConf.Type),
		StageId:   int32(this.StageConf.Id),
		UseTime:   int32(time.Now().Unix() - this.createTime),
	}
	if this.fightResult {
		msg.Winners = []int32{int32(this.fightOwner.GetUserId())}
	} else {
		msg.Losers = []int32{int32(this.fightOwner.GetUserId())}
	}
	msg.CpData = []byte(fmt.Sprintf("%d,%d", this.killMonsterNum, this.exp))
	logger.Info("发送game经验副本战斗结果,服务器：%v，结果：%v", this.fightOwner.HostId(), *msg)
	net.GetGsConn().SendMessageToGs(uint32(this.fightOwner.HostId()), msg)
	this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_CLOSING, 15)
}

func (this *ExpFight) Begin() {

	this.InitMonsterByWave()
	//重置战斗开始时间
	this.createTime = time.Now().Unix()
}

func (this *ExpFight) InitMonsterByWave() {
	monsterNum, _ := this.DefaultFight.InitMonsterByWave(this.nowWave)
	this.nowWaveLess += monsterNum
}
