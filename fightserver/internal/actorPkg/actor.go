package actorPkg

import (
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/buff"
	"cqserver/fightserver/internal/scene"
	"cqserver/gamelibs/fsm"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"time"
)

type threatActor struct {
	actorObjId int //角色战斗Id
	threat     int //仇恨值
}

type PassiveSkillReady struct {
	skillId int
	target  base.Actor
}

type DefaultActor struct {
	*scene.DefaultSceneObj
	context         base.Actor //角色实体
	userId          uint32
	nickName        string
	avatar          string
	job             int
	sex             int
	displayInfo     *pbserver.ActorDisplayInfo
	prop            *base.ActorProp //属性
	birthPoint      *scene.Point    //出生点
	official        int             //官职
	isDeath         bool
	deathReason     int
	killer          base.Actor
	lastMoveTime    int64
	moveFastTimes   int
	damageTotal     int
	inFightTime     int64 //进入战斗时间 纳秒（玩家释放技能或被攻击 buff掉血也算，反之 一定时间game表配置则为退出战斗状态）
	inFightTimeLast int64 //最近一次进入战斗时间 纳秒（玩家释放技能或被攻击 buff掉血也算，反之 一定时间game表配置则为退出战斗状态）
	reliveTime      int64

	threatMap                  []*threatActor     //仇恨列表
	target                     base.Actor         //攻击目标
	skills                     []*base.Skill      //技能
	passiveSkill               []*base.Skill      //技能
	useSkillIndex              int                //上次使用技能索引
	useSkillTime               time.Time          //上次使用技能时间
	teamIndex                  int                //阵营
	passiveSkillReady          *PassiveSkillReady //触发准备释放的被动技能
	passiveSkillConditionTimes map[int]int        //被动技能条件计数

	fsm         *fsm.FSM
	fight       base.Fight
	pathSlice   []*scene.Point
	buffManager *buff.BuffManager
}

func (this *DefaultActor) ReliveTime() int64 {
	return this.reliveTime
}

func (this *DefaultActor) SetReliveTime(reliveTime int64) {
	this.reliveTime = reliveTime
}

func (this *DefaultActor) InFightTimeLast() int64 {
	return this.inFightTimeLast
}

func (this *DefaultActor) InFightTime() int64 {
	return this.inFightTime
}

func (this *DefaultActor) SetInFightTime() {
	now := time.Now().UnixNano()
	if now-this.inFightTimeLast > int64(gamedb.GetConf().ResetFightStage)*1000000 {
		this.inFightTime = now
	}
	this.inFightTimeLast = now
}

func (this *DefaultActor) DamageTotal() int {
	return this.damageTotal
}

func (this *DefaultActor) Killer() base.Actor {
	return this.killer
}

func (this *DefaultActor) SetKiller(killer base.Actor) {
	logger.Debug("当前角色：%v,击杀者：", this.NickName(), killer)
	this.killer = killer
}

func (this *DefaultActor) IsDeath() bool {
	return this.isDeath
}

func (this *DefaultActor) SetIsDeath(isDeath bool, deathReason int) {
	this.isDeath = isDeath
	this.deathReason = deathReason
}

func (this *DefaultActor) DeathReason() int {
	return this.deathReason
}

func (this *DefaultActor) Official() int {
	return this.official
}

func (this *DefaultActor) SetOfficial(official int) {
	this.official = official
}

func (this *DefaultActor) Sex() int {
	return this.sex
}

func (this *DefaultActor) TeamIndex() int {
	return this.teamIndex
}

func (this *DefaultActor) SetTeamIndex(teamIndex int) {
	this.teamIndex = teamIndex
}

func (this *DefaultActor) SetBirthPoint(point *scene.Point) {
	if this.birthPoint == nil {
		this.birthPoint = point
	}
}

func (this *DefaultActor) BirthPoint() *scene.Point {
	return this.birthPoint
}

func NewDefaultActor(typ int, name string, avatar string, job int32, dispalyInfo *pbserver.ActorDisplayInfo, fsm *fsm.FSM, context base.Actor) *DefaultActor {
	return &DefaultActor{
		DefaultSceneObj:            scene.NewDefaultSceneObj(context, typ),
		context:                    context,
		fsm:                        fsm,
		skills:                     make([]*base.Skill, 0),
		passiveSkill:               make([]*base.Skill, 0),
		useSkillIndex:              -1,
		threatMap:                  make([]*threatActor, 0),
		target:                     nil,
		buffManager:                buff.NewBuffManager(context),
		prop:                       base.NewActorProp(),
		displayInfo:                dispalyInfo,
		avatar:                     avatar,
		nickName:                   name,
		job:                        int(job),
		isDeath:                    false,
		passiveSkillConditionTimes: make(map[int]int),
	}
}

func (this *DefaultActor) GetNextPathPoint() (*scene.Point, bool) {
	lenght := len(this.pathSlice)
	if this.pathSlice == nil || lenght <= 0 {
		return nil, false
	}
	point := this.pathSlice[lenght-1]
	this.pathSlice = this.pathSlice[:lenght-1]
	isLast := len(this.pathSlice) == 0
	return point, isLast
}

func (this *DefaultActor) PathSlice() []*scene.Point {
	return this.pathSlice
}

func (this *DefaultActor) SetPathSlice(pathSlice []*scene.Point) {
	this.pathSlice = pathSlice
}

func (this *DefaultActor) CanMove() bool {

	now := time.Now().UnixNano()
	if this.lastMoveTime > 0 && now-this.lastMoveTime < base.MoveInterval {
		this.moveFastTimes += 1
		logger.Debug("移动过快 玩家：%v,上次移动时间：%v,当前时间：%v,限制移动间隔：%v,移动过快次数：%v", this.nickName, this.lastMoveTime, now, base.MoveInterval, this.moveFastTimes)
		if this.moveFastTimes >= constFight.MOVE_FAST_LIMIT {
			return false
		}
	} else {
		this.moveFastTimes = 0
	}

	canMove := this.buffManager.IsCanMove()
	return canMove
}

func (this *DefaultActor) MoveTo(point *scene.Point, moveType int, moveForce, sendClient bool) error {
	this.lastMoveTime = time.Now().UnixNano()
	return this.DefaultSceneObj.MoveTo(point, moveType, moveForce, sendClient)
}

func (this *DefaultActor) CanAttack() bool {

	if this.IsSceneObj() {
		return false
	}

	if this.GetType() == pb.SCENEOBJTYPE_PET {
		return false
	}

	if this.GetProp().HpNow() <= 0 {
		return false
	}

	//if this.buffManager.IsCanWudi() {
	//	return false
	//}
	return true
}

func (this *DefaultActor) IsCanTreat() bool {
	return this.buffManager.IsCanTreat()
}

func (this *DefaultActor) GetTreatEffect() int {
	return this.buffManager.GetTreatEffect()
}

func (this *DefaultActor) DisplayInfo() *pbserver.ActorDisplayInfo {
	return this.displayInfo
}

func (this *DefaultActor) SetDisplayInfo(displayInfo *pbserver.ActorDisplayInfo) {
	this.displayInfo = displayInfo
}

func (this *DefaultActor) Avatar() string {
	return this.avatar
}

func (this *DefaultActor) NickName() string {
	return this.nickName
}

func (this *DefaultActor) Job() int {
	return this.job
}

func (this *DefaultActor) GetProp() *base.ActorProp {
	return this.prop
}

func (this *DefaultActor) SetFight(fight base.Fight) {
	this.fight = fight
}

func (this *DefaultActor) Relive(reliveAddr int, reliveType int) {
	//复活
	this.reliveTime = 0
	point := this.Point()
	if reliveAddr == constFight.RELIVE_ADDR_TYPE_BIRTH {
		point = this.BirthPoint()
	}
	this.killer = nil
	this.GetProp().SetHpNow(this.GetProp().Get(pb.PROPERTY_HP))
	this.GetProp().SetMpNow(this.GetProp().Get(pb.PROPERTY_MP))
	//删除debuf 跟部分buff
	this.DelDeBuff()
	this.SetVisible(true)
	this.SetIsDeath(false, 0)
	this.GetScene().ReliveSceneObj(this.context, point, reliveType)
	this.GetFight().OnRelive(this.context, reliveType)
	this.context.TriggerPassiveSkill(constFight.SKILL_PASSIVE_RELIVE, nil, nil)
	this.GetFSM().Event(base.StateTriggerActive)
}

func (this *DefaultActor) ChangeHp(changeHp int) (realChange int, isDeath bool) {

	oldHp := this.GetProp().HpNow()
	if changeHp < 0 {
		realChange = this.GetProp().DecHP(changeHp)
	} else {
		realChange = this.GetProp().AddHP(changeHp)
	}
	if this.GetProp().HpNow() <= 0 {
		isDeath = true
	}
	newHp := this.GetProp().HpNow()

	this.context.TriggerPassiveSkill(constFight.SKILL_PASSIVE_CONDITION_NORMAL, nil, nil)
	if newHp < oldHp {
		this.TriggerPassiveSkillByHpChange(constFight.SKILL_PASSIVE_CONDITION_NORMAL, nil, newHp, oldHp)
	}
	return
}

func (this *DefaultActor) CalDamage(damage int) {
	this.damageTotal += damage
	this.context.TriggerPassiveSkill(constFight.SKILL_PASSIVE_DAMAGE, nil, nil)
}

func (this *DefaultActor) OnDie() bool {
	this.isDeath = true
	return false
}

func (this *DefaultActor) LeaveScene() {
	this.buffManager.LeaveScene()
	this.DefaultSceneObj.LeaveScene()
}

func (this *DefaultActor) RunAI() {
	if this.fight == nil || this.fsm == nil {
		return
	}
	if this.buffManager == nil {
		return
	}

	this.fsm.Run()
	this.context.TriggerPassiveSkill(constFight.SKILL_PASSIVE_SAME_SECOND, nil, nil)
	this.buffManager.Run()

	if this.target != nil && this.target.GetProp().HpNow() <= 0 {
		this.LockTarget(nil)
	}
}

func (this *DefaultActor) GetFSM() *fsm.FSM {
	return this.fsm
}

func (this *DefaultActor) GetFight() base.Fight {
	if this.fight != nil {
		return this.fight.GetContext()
	}
	return nil
}

func (this *DefaultActor) HasTheadTarget() bool {
	return len(this.threatMap) > 0
}

//增加仇恨值
func (this *DefaultActor) AddTheat(id, value int) {

	for _, v := range this.threatMap {
		if v.actorObjId == id {
			v.threat += value
			return
		}
	}

	this.threatMap = append(this.threatMap, &threatActor{id, value})
}

//仇恨值最高的目标Id
func (this *DefaultActor) GetTargetByMaxTheat() int {

	maxObjId := 0
	maxThreat := 0
	for _, v := range this.threatMap {
		if maxObjId == 0 || v.threat > maxThreat {
			maxObjId = v.actorObjId
		}
	}
	return maxObjId
}

//仇恨值最高的目标Id
func (this *DefaultActor) GetTargetByFirstTheat() int {
	if len(this.threatMap) > 0 {
		return this.threatMap[0].actorObjId
	}
	return 0
}

func (this *DefaultActor) LockTarget(target base.Actor) {
	this.target = target
}

//移除一个仇恨值目标
func (this *DefaultActor) ClearTheatTarget(objId int) {
	if this.target != nil && this.target.GetObjId() == objId {
		this.target = nil
	}
	for k, v := range this.threatMap {
		if v.actorObjId == objId {
			this.threatMap = append(this.threatMap[:k], this.threatMap[k+1:]...)
			break
		}
	}
}

//移除一个仇恨值目标
func (this *DefaultActor) ClearAllTheat() {
	this.target = nil
	this.threatMap = make([]*threatActor, 0)
}

func (this *DefaultActor) IsEnemy(target base.Actor) bool {
	return false
}
func (this *DefaultActor) IsFriend(target base.Actor) bool {
	return false
}

func (this *DefaultActor) ClearSkillCD() {
	for _, skill := range this.GetSkills() {
		skill.SetNextAttackTime(0)
	}
}

func (this *DefaultActor) GetReadySkill(target base.Actor) *base.Skill {

	//dir := scene.GetFaceDirByPoint(this.Point(), target.Point())
	//offsetX, offsetY := target.Point().X()-this.Point().X(), target.Point().Y()-this.Point().Y()
	dis := scene.DistanceByPoint(this.Point(), target.Point())
	skillCount := len(this.skills)
	var skill *base.Skill
	for i := 1; i <= skillCount; i++ {
		this.useSkillIndex += 1
		this.useSkillIndex = this.useSkillIndex % skillCount
		tempSkill := this.skills[this.useSkillIndex]
		if useErr, _ := this.CanUseSkill(tempSkill); useErr != nil {
			continue
		}
		if tempSkill.LevelT.Distance < dis {
			continue
		}
		//if !tempSkill.InAttackArea(dir, offsetX, offsetY) {
		//	//logger.Debug("----------------------",dir,offsetX,offsetX)
		//	continue
		//}
		skill = tempSkill
		//如果不是普攻，则返回找到的技能，如果普攻继续找
		if tempSkill.Type != pb.SKILLTYPE_ORDINARY {
			break
		}
	}
	return skill
}

func (this *DefaultActor) CanUseSkill(skill *base.Skill) (error, bool) {
	attackSpeed := this.GetProp().Get(pb.PROPERTY_ATT_SPEED)
	attackInterval := int64(gamedb.GetAttIntervalByAttSpeed(attackSpeed)) - 20
	if skill.Aspd && !this.useSkillTime.IsZero() && time.Now().Sub(this.useSkillTime).Milliseconds() < attackInterval {
		return gamedb.ERRATTACKSPEED, false
	}
	canUse, skillEffect := this.buffManager.IsCanUseSkiLL(skill.Skillid)
	if !canUse {
		return gamedb.ERRSKILLCASTBYBUFF, false
	}
	err := skill.CanUse(this.GetProp().MpNow())
	return err, skillEffect
}

func (this *DefaultActor) UseSkill(skill *base.Skill, isElf bool) {
	if isElf {
		skill.Use()
	} else {
		if skill.Type != pb.SKILLTYPE_PASSIVE && skill.Type != pb.SKILLTYPE_PASSIVE2 {
			this.useSkillTime = time.Now()
		}
		if this.buffManager.BuffSkillInCd() {
			skill.Use()
		}
		//扣除蓝量
		this.GetProp().SetMpNow(this.GetProp().MpNow() - skill.LevelT.MP)
	}
}

/**
 *  @Description: 获取技能释放点
 *  @param skill
 *  @param dir
 *  @param targetIds
 *  @return *scene.Point
 */
func (this *DefaultActor) getSkillCastPoint(skill *base.Skill, dir int, targetIds []int, isElf bool) (*scene.Point, error) {

	var target base.Actor
	if len(targetIds) > 0 {
		target = this.GetFight().GetActorByObjId(targetIds[0])
	}
	//判断精灵技能攻击
	if !isElf {
		this.SetDir(dir)
	}

	var dis = skill.LevelT.Distance
	if target != nil {
		dis = scene.DistanceByPoint(this.Point(), target.Point())
		if dis > skill.LevelT.Distance {
			return nil, gamedb.ERRDISTANCE
		}
	}

	if skill.SkillSkillCfg.Skillpoint == constFight.BUFF_TARGET_SELF {
		return this.Point(), nil
	} else {
		if target == nil {
			return nil, gamedb.ERRGETSKILLCASTTARGET
		} else {
			return target.Point(), nil
		}
	}

	//offsetX, offsetY := scene.GetDirOffset(dir)
	//return this.GetFight().GetScene().GetPointByXY(this.Point().X()+offsetX*dis, this.Point().Y()+offsetY*dis)
}

func (this *DefaultActor) TriggerPassiveSkillByHpChange(passiveType int, target base.Actor, oldHp, newHp int) {

	TriggerPassiveSkillByHpChange(this.context, this.passiveSkill, passiveType, target, oldHp, newHp)
}

func (this *DefaultActor) TriggerPassiveSkill(passiveType int, target base.Actor, skill *base.Skill, ) {

	if passiveType == constFight.SKILL_PASSIVE_CONDITION_ATK {
		//攻击 类型被动技能触发，只能是普通攻击 主动技能触发(切割技能不触发)，其他不触发
		if skill.Skillid == constFight.SKILL_CUT_ZHAN ||
			skill.Skillid == constFight.SKILL_CUT_FA ||
			skill.Skillid == constFight.SKILL_CUT_DAO {
			return
		}
		if skill.Type == pb.SKILLTYPE_ORDINARY || skill.Type == pb.SKILLTYPE_ACTIVE {
			this.passiveSkillConditionTimes[passiveType] ++
		} else {
			return
		}

	} else {
		this.passiveSkillConditionTimes[passiveType] ++
	}
	TriggerPassiveSkill(this.context, this.passiveSkill, this.passiveSkillConditionTimes, passiveType, target, skill)
}

func (this *DefaultActor) ReadyCastPassiveSkill(skillId int, target base.Actor) {
	if this.passiveSkillReady == nil {
		this.passiveSkillReady = &PassiveSkillReady{
			skillId: skillId,
			target:  target,
		}
	} else {
		logger.Debug("一次触发了2个被动技能，本次被忽略了,玩家：%v,skillId:%v", this.context.NickName(), skillId)
	}
}

func (this *DefaultActor) CastPassiveSkill() {

	if this.passiveSkillReady == nil {
		return
	}
	skillId := this.passiveSkillReady.skillId
	target := this.passiveSkillReady.target
	this.passiveSkillReady = nil
	skill := this.context.GetSkill(skillId)
	if skill == nil {
		logger.Error("释放被动技能异常，技能未找到,玩家：%v，技能：%v", this.context.NickName(), skillId)
		return
	}

	if err := skill.CanUse(this.GetProp().MpNow()); err != nil {
		logger.Error("释放被动技能异常，技能释放条件不足,玩家：%v，技能：%v,err:%v", this.context.NickName(), skillId, err)
		return
	}
	targetId := make([]int, 0)
	if target != nil {
		targetId = append(targetId, target.GetObjId())
	}
	_, err := CastSkill(this.context, skill, this.GetDir(), targetId, false)
	if err != nil {
		logger.Error("释放被动技能异常,玩家：%v，技能：%v,err:%v", this.context.NickName(), skillId, err)
	}
}

func (this *DefaultActor) CastSkill(skill *base.Skill, dir int, targetIds []int, isElf bool) (*pb.AttackEffectNtf, error) {
	return CastSkill(this.context, skill, dir, targetIds, isElf)
}

//攻速buff影响
func (this *DefaultActor) ApsdBuff(targets []base.Actor) {
	//锁定目标
	if len(targets) <= 0 {
		return
	}
	isChangeTarget := true
	if this.target != nil {
		for _, v := range targets {
			if v.GetObjId() == this.target.GetObjId() {
				isChangeTarget = false
			}
		}
	}
	if isChangeTarget {
		this.LockTarget(targets[0])
	}
	this.AspdBuffAddValue(isChangeTarget)
}

func (this *DefaultActor) AddOwner(attacker base.Actor, force bool) {
}

func (this *DefaultActor) GetSkills() []*base.Skill {
	return this.skills
}

func (this *DefaultActor) SetSkills(skills []*base.Skill) {

	useSkills := make([]*base.Skill, 0)
	passiveSkills := make([]*base.Skill, 0)

	for _, v := range skills {
		if v.SkillSkillCfg.Type == pb.SKILLTYPE_PASSIVE || v.SkillSkillCfg.Type == pb.SKILLTYPE_PASSIVE2 {
			passiveSkills = append(passiveSkills, v)
		} else {
			useSkills = append(useSkills, v)
		}
	}
	this.passiveSkill = passiveSkills
	this.skills = useSkills
}

func (this *DefaultActor) GetSkill(id int) *base.Skill {
	for _, skill := range this.skills {
		if skill.Skillid != id {
			continue
		}
		return skill
	}
	for _, skill := range this.passiveSkill {
		if skill.Skillid == id {
			return skill
		}
	}
	logger.Error("UserActor: skill nil")
	return nil
}

func (this *DefaultActor) GetUserId() int {
	if this.GetType() == pb.SCENEOBJTYPE_USER {
		return this.context.(*UserActor).GetUserId()
	}
	return 0
}

func (this *DefaultActor) GetAllBuffsPb() []*pb.BuffInfo {
	return this.buffManager.GetAllBuffsPb()
}

func (this *DefaultActor) GetBuffSkillHurtAdd() int {
	return this.buffManager.GetBuffSkillHurtAdd()
}

func (this *DefaultActor) GetBuffSkillHurtAddBySkillId(skillId int) int {
	return this.buffManager.GetBuffSkillHurtAddBySkillId(skillId)
}

func (this *DefaultActor) GetBuffFinalHurtDec() int {
	return this.buffManager.GetBuffFinalHurtDec()
}
func (this *DefaultActor) GetBuffFinalHurtAdd() int {
	return this.buffManager.GetBuffFinalHurtAdd()
}

func (this *DefaultActor) GetBuffFinalHurtDecFix(hurt int) int {
	return this.buffManager.GetBuffFinalHurtDecFix(hurt)
}

func (this *DefaultActor) DelAllBuff() {
	this.buffManager.ClearALlBuff()
}

//删除负面buff
func (this *DefaultActor) DelDeBuff() {
	this.buffManager.DelDeBuff()
}
func (this *DefaultActor) DelGoodBuff(layer int) {
	this.buffManager.DelGoodBuff(layer)
}

func (this *DefaultActor) AddBuff(buffId int, sourceActor base.Actor, isInit bool, arg ...int) (int, error) {
	hpChange, err := this.buffManager.AddBuff(buffId, sourceActor, isInit, arg...)
	return hpChange, err
}

/**
 *  @Description: 			添加一个buff
 *  @param buffId			buffId
 *  @param sourceActor		来源
 *  @param target			当前攻击目标
 *  @param isInit			是否初始化
 *  @param arg				其他参数
 *  @return int				返回血量变化
 *  @return error			返回异常
 */
func (this *DefaultActor) AddNewBuff(buffId int, sourceActor base.Actor, target base.Actor, isInit bool, arg ...int) (int, error) {
	hpChange, err := this.buffManager.AddNewBuff(buffId, sourceActor, target, isInit, arg...)
	return hpChange, err
}

func (this *DefaultActor) AspdBuffAddValue(isClear bool) {
	this.buffManager.AspdBuffAddValue(isClear)
}

func (this *DefaultActor) GetBuffFireSkillHurtAdd() int {
	return this.buffManager.GetBuffFireSkillHurtAdd()
}

func (this *DefaultActor) BuffHasType(buffType int, buffIds []int) (bool, int) {
	return this.buffManager.BuffHasType(buffType, buffIds)
}

func (this *DefaultActor) GetBuffEffectByBuffType(buffType int) int {
	return this.buffManager.GetBuffEffectByBuffType(buffType)
}

func (this *DefaultActor) BuffRemove(buffLayer int, buffType []int) {
	this.buffManager.BuffRemove(buffLayer, buffType)
}

func (this *DefaultActor) BuffRemoveByBuffId(buffId int) {
	this.buffManager.BuffRemoveByBuffId(buffId)
}

func (this *DefaultActor) BuffFatalRecoveHp() {
	this.buffManager.BuffFatalRecoveHp()
}

func (this *DefaultActor) ReflexHurt() bool {
	return this.buffManager.ReflexHurt()
}

func (this *DefaultActor) ClearFitBuff() {
	this.buffManager.ClearFitBuff()
}

func (this *DefaultActor) CheckTriggerPropMust(attacker base.Actor, propId int) bool {
	return this.buffManager.CheckTriggerPropMust(attacker, propId)
}

func (this *DefaultActor) NotifyNearby(obj scene.ISceneObj, msg nw.ProtoMessage, excludeSession map[uint32]bool) {
	if this.GetScene() != nil {
		this.GetScene().NotifyNearby(obj, msg, excludeSession)
	}
}
