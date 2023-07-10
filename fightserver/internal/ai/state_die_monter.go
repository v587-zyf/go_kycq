package ai

import (
	"cqserver/fightserver/internal/actorPkg"
	"cqserver/fightserver/internal/base"
	"cqserver/gamelibs/fsm"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pbserver"
	"strconv"
	"time"
)

type MonsterDieState struct {
	*DieState
}

func NewMonsterDieState(aifsm *fsm.FSM, actor base.Actor) *MonsterDieState {
	return &MonsterDieState{
		DieState: &DieState{
			actor: actor,
			fsm:   aifsm,
		},
	}
}

func (this *MonsterDieState) Enter() {

	fight := this.actor.GetFight()
	monsterActor := this.actor.(*actorPkg.MonsterActor)
	//清空路径
	this.actor.SetPathSlice(nil)
	this.actor.ClearAllTheat()
	this.actor.DelAllBuff()
	var owner base.Actor
	if monsterActor.MonsterT.DropType == constFight.MONSTER_DROP_FOR_KILLER {
		if this.actor.Killer() != nil {
			owner = fight.GetUserMainActor(this.actor.Killer().GetUserId())
		}
	} else {
		owner = fight.GetUserMainActor(monsterActor.Owner())
	}
	if monsterActor.MonsterT.DropId > 0 && owner != nil {
		if player, ok := owner.(base.ActorPlayer); ok {
			userDropInfo := player.GetPlayer().RedPacketInfo()
			isFrist := player.GetPlayer().StageFightNum() == 0
			dropItems, _, err := gamedb.GetMonsterDrop(monsterActor.MonsterT.Monsterid, int(userDropInfo.PickNum), int(userDropInfo.PickMax),
				common.ConvertMapInt32ToInt(userDropInfo.PickInfos), isFrist, fight.GetStageConf().Type)

			if err != nil || len(dropItems) == 0 {
				logger.Error("推送怪物掉落到game异常,怪物id:%v,玩家：%v，err:%v ", monsterActor.MonsterT.Monsterid, owner.GetUserId(), err)
			} else {
				//推送战斗
				itemStr := "【"
				for _, item := range dropItems {
					itemStr += strconv.Itoa(item.ItemId) + "-" + strconv.Itoa(item.Count) + ";"
				}
				itemStr += "】"
				logger.Debug("战斗结束,怪物死亡,掉落，战斗stage:%v,归属：%v,怪物：%v,掉落数量：:%v,掉落物品：%v", fight.GetStageConf().Id, owner.NickName(), this.actor.NickName(), len(dropItems), itemStr)
				drop := make([]*pbserver.ItemUnit, len(dropItems))
				for k, v := range dropItems {
					drop[k] = &pbserver.ItemUnit{
						ItemId:  int32(v.ItemId),
						ItemNum: int32(v.Count),
					}
				}
				fight.MonsterDrop(monsterActor.MonsterT.Monsterid, this.actor.Point().X(), this.actor.Point().Y(), owner, drop)
			}
		}

	}

	stageConf := fight.GetStageConf()

	if _, ok := stageConf.Monster_num[monsterActor.GetBirthAreaIndex()]; ok {
		//怪物按区域复活，此处不复活
		this.fsm.Pause()
	} else if monsterActor.MonsterT.ReliveDelay > 0 {
		//延迟多久复活
		this.fsm.SetSleepTime(monsterActor.MonsterT.ReliveDelay)
		this.actor.SetReliveTime(common.GetNowMillisecond() + int64(monsterActor.MonsterT.ReliveDelay))
	} else if len(monsterActor.MonsterT.Refresh) > 0 {

		//固定时间点复活
		reliveTimes := common.GetMilliSecondsByString(monsterActor.MonsterT.Refresh)
		dayZeroMilliSecond := common.GetTimeSeconds(time.Now()) * 1000
		sleepTime := 0
		reliveTime := 0
		for _, v := range reliveTimes {
			if dayZeroMilliSecond < v {
				sleepTime = v - dayZeroMilliSecond
				reliveTime = v
			}
		}
		if sleepTime == 0 && reliveTime == 0 {
			sleepTime = reliveTimes[0] + 86400000 - dayZeroMilliSecond
			reliveTime = reliveTimes[0] + 86400000
		}
		this.fsm.SetSleepTime(sleepTime)
		this.actor.SetReliveTime(int64(reliveTime))

	} else {
		//fight.Leave(this.actor)
		this.fsm.Pause()
	}

	killNtfToGs(this.actor.Killer(),this.actor)
	this.actor.OnDie()
	this.actor.GetFight().OnDie(this.actor, this.actor.Killer())
}

func (this *MonsterDieState) Execute() {

	if this.actor.ReliveTime() <= 0 {
		return
	}
	if common.GetNowMillisecond() > this.actor.ReliveTime() {
		//复活
		monsterActor := this.actor.(*actorPkg.MonsterActor)
		monsterActor.SetOwner(0)
		this.actor.Relive(monsterActor.MonsterT.ReliveAddrType, constFight.RELIVE_TYPE_NOMAL)
	}
}
