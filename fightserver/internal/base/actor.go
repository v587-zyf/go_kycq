package base

import (
	"cqserver/fightserver/internal/scene"
	"cqserver/gamelibs/fsm"
	"cqserver/gamelibs/gamedb"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
)

type AICreator func(actor Actor) *fsm.FSM

type ActorType int

const (
	ActorTypeUser    = pb.SCENEOBJTYPE_USER
	ActorTypeMonster = pb.SCENEOBJTYPE_MONSTER
)

var MoveInterval int64

type Actor interface {
	scene.ISceneObj
	GetProp() *ActorProp
	SetFight(fight Fight)
	RunAI()
	GetNextPathPoint() (*scene.Point, bool) //获取下一个可行走的路径
	SetPathSlice(pathSlice []*scene.Point)  //设置路径
	GetFSM() *fsm.FSM
	GetFight() Fight
	HasTheadTarget() bool        //是否有仇恨目标
	AddTheat(id, value int)      //增加仇恨值
	GetTargetByFirstTheat() int  //获取第一个攻击目标
	ClearTheatTarget(objdId int) //移除一个仇恨目标
	ClearAllTheat()              //清除所有仇恨
	GetReadySkill(target Actor) *Skill
	SetSkills(skills []*Skill)
	GetSkill(id int) *Skill
	UseSkill(skill *Skill, isElf bool)
	GetUserId() int
	NickName() string
	Job() int
	Sex() int
	Official() int
	CanMove() bool
	CanAttack() bool
	ClearSkillCD()
	CastSkill(skill *Skill, dir int, targetIds []int, isElf bool) (*pb.AttackEffectNtf, error)
	ReadyCastPassiveSkill(skillId int, target Actor)
	CastPassiveSkill()
	TriggerPassiveSkill(passiveType int, target Actor, skill *Skill)
	TriggerPassiveSkillByHpChange(passiveType int, target Actor, oldHp, newHp int)
	IsEnemy(target Actor) bool
	IsFriend(target Actor) bool
	AddOwner(attacker Actor, force bool)
	TeamIndex() int             //设置阵营
	SetTeamIndex(teamIndex int) //设置阵营
	BirthPoint() *scene.Point
	SetBirthPoint(point *scene.Point)
	Relive(reliveAddr int, reliveType int)
	OnDie() bool
	IsDeath() bool
	SetIsDeath(isDeath bool, deathReason int)
	DeathReason() int
	SetKiller(killer Actor)
	Killer() Actor
	DamageTotal() int
	/**
	 *  @Description: 血量改变
	 *  @param changeHp		改变值（扣血为负值）
	 *  @return realChange	实际变化值
	 *  @return isDeath		是否死亡
	 */
	ChangeHp(changeHp int) (realChange int, isDeath bool)
	//伤害统计
	CalDamage(damage int)
	InFightTime() int64
	InFightTimeLast() int64
	SetInFightTime()
	ReliveTime() int64
	SetReliveTime(reliveTime int64)

	GetAllBuffsPb() []*pb.BuffInfo
	GetBuffSkillHurtAdd() int
	GetBuffSkillHurtAddBySkillId(skillId int) int
	GetBuffFinalHurtDec() int
	GetBuffFinalHurtAdd() int
	GetBuffFireSkillHurtAdd() int
	GetBuffFinalHurtDecFix(hurt int) int
	DelDeBuff()
	DelGoodBuff(layer int)
	DelAllBuff()
	BuffRemove(buffLayer int, buffType []int)
	BuffRemoveByBuffId(buffId int)
	AddBuff(buffId int, sourceActor Actor, isInit bool, arg ...int) (int, error)
	/**
	 *  @Description: 			添加一个buff
	 *  @param buffId			buffId
	 *  @param sourceActor		来源
	 *  @param attackTarget		当前攻击目标
	 *  @param isInit			是否初始化
	 *  @param arg				其他参数
	 *  @return int				返回血量变化
	 *  @return error			返回异常
	 */
	AddNewBuff(buffId int, sourceActor Actor, attackTarget Actor, isInit bool, arg ...int) (int, error)
	AspdBuffAddValue(isClear bool)
	BuffHasType(buffType int, buffId []int) (bool, int)
	GetBuffEffectByBuffType(buffType int) int
	BuffFatalRecoveHp()
	IsCanTreat() bool
	GetTreatEffect() int
	//是否反弹伤害
	ReflexHurt() bool
	ClearFitBuff()
	CheckTriggerPropMust(attacker Actor, propId int) bool
	ApsdBuff(targets []Actor)
	NotifyNearby(obj scene.ISceneObj, msg nw.ProtoMessage, excludeSession map[uint32]bool)
}

type ActorUser interface {
	SetFightModel(fightModel int)
	GetHeroIndex() int
	GuildId() int
	GuildName() string
	SendMessage(msg nw.ProtoMessage)
	SetCollectionId(collectionId int)
	ResetCollectionStatus()
	GetElfAttack() float64
	UseCutTreasure(cutTreasureLv int) error
	UpdateElf(elf *pbserver.ElfInfo) error
	SetRelive()
}

type ActorMonster interface {
	GetMonsterT() *gamedb.MonsterMonsterCfg
	Owner() int
}

type ActorLeader interface {
	GetLeader() Actor
}

type ActorPlayer interface {
	GetPlayer() *PlayerActor
}
