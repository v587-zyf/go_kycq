package base

const (
	FIGHT_EVENT_TYPE_PICK_ALL  = 1 //全部拾取
	FIGHT_EVENT_TYPE_ACTOR_DIE = 2 //角色死亡
)

type FightEvent struct {
	EventType int
	Data      interface{}
}

type ActorDieEvent struct {
	DieActor Actor
	Killer   Actor
}
