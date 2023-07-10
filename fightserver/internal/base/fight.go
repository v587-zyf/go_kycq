package base

import (
	"cqserver/fightserver/internal/scene"
	"cqserver/gamelibs/gamedb"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
)

type Message interface {
	Handle()
}

type Fight interface {
	GetContext() Fight // 获取真正的Fight实例对应的接口
	SetId(id uint32)
	GetId() uint32
	GetStageConf() *gamedb.StageStageCfg
	GetFightExtMark() int

	UserEnter(userEnterMsg *pbserver.FSEnterFightReq) error
	EnterSummon(leaderActor Actor, summonId int)
	EnterMuli(actor1 []Actor, point []*scene.Point, enterType int) error
	LeaveUser(userId int)
	Leave(actor Actor)
	Range(func(actor Actor) bool)
	GetPlayerNum() int
	GetOnlineActor(sessionId uint32) Actor // 获取在线的玩家角色
	GetActorByObjId(objId int) Actor
	GetSceneObj(objId int) scene.ISceneObj
	GetSceneAllObj() map[int]scene.ISceneObj
	GetScene() *scene.Scene
	KickActorByGate(serverId int)
	GetUserByUserId(userId int) map[int]Actor
	GetUserFitActor(userId int) Actor
	GetUserMainActor(userId int) Actor
	GetPetActor(userId int) Actor
	UpdateUserFigntInfo(userInfo *pbserver.Actor, heroIndex int)
	ChangeUserToHelper(userId, toHelpUserId int)
	PlayerFightNumChange(req *pbserver.GsToFsFightNumChangeReq)
	FightScoreLess(userId int, lessNum int) (int, error)
	CheckUserAllDie(actorUser Actor) bool
	CheckUserAllDieByHp(actorUser Actor) bool
	RandomDelivery(userId int, rand bool)
	GetBossAliveNum() int

	Start()
	Stop()
	DoLoop() chan struct{} // 执行fight的事件循环
	SendMessage(message Message)
	GetMessageChan() chan Message
	Begin()
	CanAttack() bool

	RunAI() // 执行AI
	UpdateFrame()
	OnEnd()
	OnRelive(actor Actor, reliveType int)
	CheckEnd() bool
	GetUserActors() map[int]Actor
	MonsterDrop(dropMonsterId int, dropX, dropY int, owner Actor, dropItems []*pbserver.ItemUnit)
	PickUp(userId int, objIds []int32,isPick bool) (map[int32]*pbserver.ItemUnitForPickUp, error)
	Collection(userId int, objs int) (int, error)
	CollectionCancel(userId int, objId int) error
	ResetCollection(objId int)
	UseFitReq(userId int, fitActor *pbserver.ActorFit) error
	FitCacelReq(userId int) error
	UpdatePet(userId int, pet *pbserver.ActorPet) error
	UpdateElf(userId int, elfInfo *pbserver.ElfInfo) error
	UpdateUserReqPacketInfo(userId int, redPacketInfo *pbserver.ActorRedPacket)
	UserRelive(userId, reliveType int) error
	GmReq(req *pbserver.GsToFsGmReq) string
	GetBossInfos() []*pb.FightBossInfoUnit

	OnEnterUser(userId int)
	OnActorEnter(atcor Actor)
	PostDamage(attacker, defender Actor, damage int)
	OnLeave(actor Actor)
	OnLeaveUser(userId int)
	OnDie(actor Actor, killer Actor)
	OnPickAll(lastPickObjId int)
	OnCollection(collections map[int]int)
	OnCheer(userId int)
	OnUsePotion(userId int)
	GetPowerRoll() string
	OnBossOwnerChange(monster Actor)
	NpcEventReq(userId, npcId int) error
}
