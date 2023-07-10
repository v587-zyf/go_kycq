package ai

import (
	"cqserver/fightserver/internal/actorPkg"
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/scene"
	"cqserver/gamelibs/fsm"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"time"
)

type MonsterFightState struct {
	*FightState
	ownerFightTime int64
}

func NewMonsterFightState(aifsm *fsm.FSM, actor base.Actor) *MonsterFightState {
	state := &MonsterFightState{}
	state.FightState = NewFightState(aifsm, actor, state)
	return state
}

func (this *MonsterFightState) getFightStageTarget() base.Actor {

	//优先攻击 归属
	fight := this.actor.GetFight()
	monster := this.actor.(*actorPkg.MonsterActor)
	//守卫龙柱，怪物不攻击归属，只攻击龙柱
	var target base.Actor
	if fight.GetStageConf().Type != constFight.FIGHT_TYPE_GUARDPILLAR {
		if monster.Owner() > 0 {
			ownerTarget := this.actor.GetFight().GetUserFitActor(monster.Owner())
			if ownerTarget != nil {
				target = ownerTarget
			} else {
				actors := this.actor.GetFight().GetUserByUserId(monster.Owner())
				if actors != nil {
					for _, v := range actors {
						if v.GetProp().HpNow() > 0 {
							target = v
							break
						}
					}
				}
			}
		}
	}

	if target != nil {
		dis := scene.DistanceByPoint(target.Point(), this.actor.Point())
		if dis > monster.MonsterT.Toattackarea {
			if this.ownerFightTime > 0 && time.Now().Unix()-this.ownerFightTime > int64(gamedb.GetConf().MonsterRangeTime) {
				monster.SetOwner(0)
			} else {
				if this.isCanMove() {
					monsterRandWalk(this.actor)
					this.fsm.SetSleepTime(400)
				}
			}
			return nil
		} else {
			this.ownerFightTime = time.Now().Unix()
			return target
		}
	}

	target = this.FightState.getFightStageTarget()
	//if target != nil {
	//	monster.SetOwner(target.GetObjId())
	//}

	return target
}
