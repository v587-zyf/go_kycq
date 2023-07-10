package base

import (
	"cqserver/golibs/logger"
)

type Team struct {
	teamId int
	actors map[int]Actor
}

func NewTeam(id int) *Team {
	return &Team{
		teamId: id,
		actors: make(map[int]Actor, 5),
	}
}

func (this *Team) AddActor(actor Actor) {
	this.actors[actor.GetUserId()] = actor
	logger.Info("team.go:AddActor:%v,len:%d teamId:%d", actor, len(this.actors), this.teamId)
}

func (this *Team) GetActor(userId int) Actor {
	return this.actors[userId]
}

func (this *Team) RemoveActor(objId int) {
	if actor, ok := this.actors[objId]; ok {
		delete(this.actors, objId)
		logger.Info("team.go:removeActor:%v,len:%d teamId:%d", actor, len(this.actors), this.teamId)
	}
}

func (this *Team) Range(f func(actor Actor) bool) bool {
	for _, actor := range this.actors {
		if f(actor) {
			return true
		}
	}
	return false
}

func (this *Team) IsAllDead() bool {
	for _, actor := range this.actors {
		if actor.GetProp().HpNow() > 0 {
			return false
		}
	}
	return true
}

func (this *Team) GetActorNum() int {
	return len(this.actors)
}

func (this *Team) GetTeamId() int {
	return this.teamId
}

func (this *Team) GetTeamUser() map[int]Actor {
	return this.actors
}
