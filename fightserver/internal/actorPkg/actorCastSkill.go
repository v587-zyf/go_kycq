package actorPkg

import (
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/scene"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/prop"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"math"
	"time"
)

/**
*  @Description: 						普通被动技能触发
*  @param attacker						释放者
*  @param passiveSkills					释放者被动技能
*  @param passiveSkillConditionTimes	释放者被动技能条件记录
*  @param passiveType					被动技能触发类型
*  @param target						目标
*  @param skill							前置技能
**/
func TriggerPassiveSkill(attacker base.Actor, passiveSkills []*base.Skill, passiveSkillConditionTimes map[int]int, passiveType int, target base.Actor, skill *base.Skill) {
	preSkillId := 0
	if skill != nil {
		preSkillId = skill.Skillid
		if passiveType == constFight.SKILL_PASSIVE_CONDITION_ATK {
			if skill.Skillid == constFight.SKILL_CUT_ZHAN ||
				skill.Skillid == constFight.SKILL_CUT_FA ||
				skill.Skillid == constFight.SKILL_CUT_DAO {
				return
			}
			//攻击 类型被动技能触发，只能是普通攻击 主动技能触发，其他不触发
			if skill.Type != pb.SKILLTYPE_ORDINARY && skill.Type != pb.SKILLTYPE_ACTIVE {
				return
			}
		}
	}

	for _, v := range passiveSkills {
		if v.LevelT.Passiveconditions1 == passiveType {

			if !triggerPassiveSkillCheckByActorInFight(attacker, v) {
				continue
			}

			if v.LevelT.Times > 0 {

				nowVal := passiveSkillConditionTimes[passiveType]
				if passiveType == constFight.SKILL_PASSIVE_DAMAGE {
					nowVal = attacker.DamageTotal()
				} else if passiveType == constFight.SKILL_PASSIVE_SAME_SECOND {
					nowVal = int(time.Now().Unix())
				}

				preVal := v.PassiveSkillConditionUse()
				if preVal+v.LevelT.Times > nowVal {
					continue
				}
				v.SetPassiveSkillConditionUse(nowVal)
			}

			//检查血量条件
			isPassiveSkill := v.PassiveSkillCheck(attacker, target, preSkillId)
			logger.Debug("被动技能触发检查,玩家：%v，技能：%v,条件检查是否通过：%v", attacker.NickName(), v.Skillid, isPassiveSkill)
			if !isPassiveSkill {
				continue
			}
			combatPass := passiveSkillCheckByCombat(v, attacker, target)
			if !combatPass {
				continue
			}
			passiveSkillCheckPass(attacker, v, target)
		}
	}
}

/**
*  @Description: 被动技能战斗力检查
*  @param passiveSkill
*  @param attack
*  @param target
*  @return bool
**/
func passiveSkillCheckByCombat(passiveSkill *base.Skill, attack, target base.Actor) bool {

	_, conditionOk1 := passiveSkill.LevelT.Passiveconditions2[constFight.SKILL_PASSIVE_COMBAT_HIGH]
	_, conditionOk2 := passiveSkill.LevelT.Passiveconditions2[constFight.SKILL_PASSIVE_COMBAT_LOW]
	if !conditionOk1 && !conditionOk2 {
		return true
	}

	if attack == nil || target == nil {
		return false
	}

	if attack.GetType() != pb.SCENEOBJTYPE_USER || target.GetType() != pb.SCENEOBJTYPE_USER {
		return false
	}

	attackPlayer, ok1 := attack.(base.ActorPlayer)
	if !ok1 {
		return false
	}
	defendPlayer, ok2 := target.(base.ActorPlayer)
	if !ok2 {
		return false
	}

	attackPlayerCombat := attackPlayer.GetPlayer().UserCombat()
	defendPlayerCombat := defendPlayer.GetPlayer().UserCombat()

	ratio := int(math.Abs(float64(defendPlayerCombat-attackPlayerCombat)) / float64(attackPlayerCombat) * common.TenThousand)
	result := false

	if value, conditionOk1 := passiveSkill.LevelT.Passiveconditions2[constFight.SKILL_PASSIVE_COMBAT_HIGH]; conditionOk1 {

		if defendPlayerCombat > attackPlayerCombat && ratio > value {
			result = true
		}
	} else if value, conditionOk1 := passiveSkill.LevelT.Passiveconditions2[constFight.SKILL_PASSIVE_COMBAT_LOW]; conditionOk1 {
		if defendPlayerCombat < attackPlayerCombat && ratio > value {
			result = true
		}
	}
	logger.Debug("被动技能战斗条件判断,技能：%v,是否通过：%v,自身：%v,自身战力：%v,敌方：%v,敌方战斗力：%v,条件：%v", passiveSkill.LevelT.Skillid, result, attack.NickName(), attackPlayerCombat, target.NickName(), defendPlayerCombat, passiveSkill.LevelT.Passiveconditions2)
	return result
}

func triggerPassiveSkillCheckByActorInFight(actor base.Actor, skill *base.Skill) bool {
	now := time.Now().UnixNano()
	if now-actor.InFightTimeLast() > int64(gamedb.GetConf().ResetFightStage)*1000000 {
		return false
	}
	if (now-actor.InFightTime())/1000000 < int64(skill.CastDuration) {
		return false
	}
	logger.Debug("被动技能触发时间检查，玩家：%v,技能：%v,进入战斗时间：%v,最后攻击/被攻击时间：%v,技能触发时间：%v", actor.NickName(), skill.Skillid, actor.InFightTime(), actor.InFightTimeLast(), skill.CastDuration)
	return true
}

/**
*  @Description: 			血量变化被动技能触发
*  @param attacker			释放者
*  @param passiveSkills		释放者被动技能
*  @param passiveType		触发被动技能类型
*  @param target			目标
*  @param oldHp				之前血量
*  @param newHp				当前血量
**/
func TriggerPassiveSkillByHpChange(attacker base.Actor, passiveSkills []*base.Skill, passiveType int, target base.Actor, oldHp, newHp int) {

	for _, v := range passiveSkills {
		if v.LevelT.Passiveconditions1 == passiveType {

			if !triggerPassiveSkillCheckByActorInFight(attacker, v) {
				continue
			}

			//检查血量条件
			if _, ok := v.LevelT.Passiveconditions2[constFight.SKILL_PASSIVE_HP_CHECK_LESS]; !ok {
				continue
			}

			canUse := v.CanUse(attacker.GetProp().MpNow())
			if canUse != nil {
				continue
			}

			lessX := v.LevelT.Passiveconditions2[constFight.SKILL_PASSIVE_HP_CHECK_LESS]
			lessXHp := int(float64(attacker.GetProp().Get(pb.PROPERTY_HP)) * float64(lessX) / base.BASE_RATE)
			old := oldHp / lessXHp
			new := newHp / lessXHp
			logger.Debug("被动技能触发检查,玩家：%v，技能：%v,血量变化值：%v-%v，总血量：%v，配置：%v,比例变化：%v-%v", attacker.NickName(), v.Skillid, oldHp, newHp, attacker.GetProp().Get(pb.PROPERTY_HP), lessX, old, new)
			if old == new {
				continue
			}
			passiveSkillCheckPass(attacker, v, target)
		}
	}
}

/**
*  @Description: 被动技能条件检查通过 触发释放
*  @param attacker	释放者
*  @param skill		被动技能
*  @param target	目标
*  @return bool
**/
func passiveSkillCheckPass(attacker base.Actor, skill *base.Skill, target base.Actor) bool {
	if !common.RandByTenShousand(skill.LevelT.Perskill) {
		return false
	}
	if skill.Cast == 1 {
		attacker.ReadyCastPassiveSkill(skill.Skillid, target)
	} else {
		base.BuffEffect(attacker, attacker, target, skill, constFight.BUFF_TARGET_SELF, constFight.BUFF_BY_SKILL_AFTER)
		base.BuffEffect(attacker, target, nil, skill, constFight.BUFF_TARGET_ENMY, constFight.BUFF_BY_SKILL_AFTER)
		//技能标记使用，记录cd
		attacker.UseSkill(skill, false)
	}
	return true
}

/**
*  @Description: 		获取技能释放点
*  @param attacker		技能释放者
*  @param skill			技能
*  @param dir			朝向
*  @param targetIds		目标
*  @param isElf			释放精灵技能
*  @return *scene.Point
*  @return error
**/
func getSkillCastPoint(attacker base.Actor, skill *base.Skill, dir int, targetIds []int, isElf bool) (*scene.Point, error) {

	if skill.SkillSkillCfg.Skillpoint == constFight.BUFF_TARGET_SELF {
		return attacker.Point(), nil
	}

	var target base.Actor
	if len(targetIds) > 0 {
		target = attacker.GetFight().GetActorByObjId(targetIds[0])
	}

	var dis = skill.LevelT.Distance
	if target != nil && target.CanAttack() {
		dis = scene.DistanceByPoint(attacker.Point(), target.Point())
		if dis > skill.LevelT.Distance {
			return nil, gamedb.ERRDISTANCE
		}
		return target.Point(), nil
	}

	return nil, gamedb.ERRGETSKILLCASTTARGET

}

/**
*  @Description: 		释放技能
*  @param attacker		释放者
*  @param skill			技能
*  @param dir			朝向
*  @param targetIds		目标
*  @param isElf			是否精灵技能
*  @return *pb.AttackEffectNtf
*  @return error
**/
func CastSkill(attacker base.Actor, skill *base.Skill, dir int, targetIds []int, isElf bool) (*pb.AttackEffectNtf, error) {

	if attacker.GetScene() == nil {
		return nil, gamedb.ERRUSERINFIGHT
	}

	//玩家当前在安全区
	if attacker.Point().IsSafe() || attacker.GetFight().GetStageConf().Type == constFight.FIGHT_TYPE_MAIN_CITY {
		return nil, gamedb.ERRINSAFEAREA
	}
	//判断战斗状态是否战斗中
	if !attacker.GetFight().CanAttack() {
		return base.SkillAttackEffect(attacker, skill, dir, make([]*pb.HurtEffect, 0), gamedb.ERRFIGHTID.Code), nil
	}
	//记录玩家在战斗中
	attacker.SetInFightTime()
	//召唤类技能
	if skill.LevelT.Summonid > 0 {
		return SummonSkillCast(attacker, skill, dir)
	}

	//获取技能释放点
	castSkillPoint, err := getSkillCastPoint(attacker, skill, dir, targetIds, isElf)
	if err != nil {
		logger.Error("释放技能：%v,未获取到技能释放点,目标：%v,err:%v", skill.Skillid, targetIds, err)
		return nil, err
	}
	//获取技能攻击区域
	offset := skill.GetAttackAreaOffset(dir)
	points, err := attacker.GetScene().GetSceneAreaByPointOffset(castSkillPoint, offset)
	if err != nil {
		logger.Error("获取攻击区域异常：%v", err)
		return nil, gamedb.ERRGETSKILLCASTAREA
	}

	//设置朝向 判断精灵技能攻击
	if !isElf {
		attacker.SetDir(dir)
	}

	if skill.SkillEffectType == pb.SKILLEFFECTTYPE_TREAT {

		return Treat(attacker, skill, dir, targetIds, points, isElf)
	}

	//对打击区域进行距离排序
	l := len(points)
	if len(points) > 1 {
		for i := 0; i < l; i++ {
			for j := i + 1; j < l; j++ {
				dis1 := scene.DistanceByPoint(attacker.Point(), points[i])
				dis2 := scene.DistanceByPoint(attacker.Point(), points[j])
				if dis1 > dis2 {
					points[i], points[j] = points[j], points[i]
				}
			}
		}
	}

	if skill.Skillid == constFight.SKILL_YEMAN_ID {
		return yeManChongZhuangSkill(attacker, skill, dir, targetIds, points)
	} else if skill.Skillid == constFight.SKILL_KANGJU_ID {
		return kangJuHuoHuanSkill(attacker, skill, dir, targetIds, points)
	} else {
		return normalHurtSkill(attacker, skill, dir, targetIds, points, isElf)
	}
}

/**
*  @Description: 普通攻击技能释放
*  @param attacker		释放者
*  @param skill			技能
*  @param dir			朝向
*  @param targetIds		目标
*  @param points		攻击区域
*  @param isElf			释放精灵技能
*  @return *pb.AttackEffectNtf
*  @return error
**/
func normalHurtSkill(attacker base.Actor, skill *base.Skill, dir int, targetIds []int, points []*scene.Point, isElf bool) (*pb.AttackEffectNtf, error) {

	//获取攻击目标
	var targets []base.Actor
	//范围查找目标
	targets = getTargetByPoint(attacker, skill, points, targetIds)
	if len(targets) <= 0 {
		logger.Warn("未找到攻击目标，技能释放空了", skill.Skillid, len(points), targetIds)
		//if !skill.LevelT.Untarget {
		return nil, gamedb.ERRGETSKILLCASTTARGET
		//}
	}

	//技能标记使用，记录cd
	attacker.UseSkill(skill, isElf)

	//计算技能伤害 推送
	hurts := make([]*pb.HurtEffect, 0)
	//攻击方buff添加
	selfBuffHurt := base.BuffEffect(attacker, attacker, nil, skill, constFight.BUFF_TARGET_SELF, constFight.BUFF_BY_SKILL_PRE)
	addHurtEffect(&hurts, selfBuffHurt)

	var castPoint *scene.Point
	var moveToPoint *scene.Point

	if skill.Skillid == constFight.SKILL_CHONGFENGZHAN_ID {

		skipPoint := attacker.Point()
		targetPoint := targets[0].Point()
		castPoint = attacker.Point()
		dis := scene.DistanceByPoint(castPoint, targetPoint)
		if dis > skill.LevelT.Range {
			skipPoint = attacker.GetScene().GetPointByDirAndMaxDis(attacker.Point(), dir, skill.LevelT.Range, true, attacker)
		}
		moveToPoint = skipPoint
		attacker.MoveTo(skipPoint, pb.MOVETYPE_WALK, true, false)
	}

	for k, target := range targets {

		//记录玩家在战斗中
		target.SetInFightTime()

		if wudi, _ := target.BuffHasType(pb.BUFFTYPE_INVINCIBLE, nil); wudi {
			h := base.CreateHurtEffect(target.GetObjId(), target.GetProp().HpNow(), 0, false, false, false, pb.HURTTYPE_NORMAL, 0, 0, 0, 0)
			h.IsWuDi = true
			addHurtEffect(&hurts, h)
			continue
		}

		//开始玩家是否命中/闪避 记录是否闪避
		isDodge := base.CalcIsDodge(attacker, target, skill)
		if isDodge {
			h := base.CreateHurtEffect(target.GetObjId(), target.GetProp().HpNow(), 0, isDodge, false, false, pb.HURTTYPE_NORMAL, 0, 0, 0, 0)
			addHurtEffect(&hurts, h)
			continue
		}

		//天赋触发
		talentEffectCal(attacker, target, skill)
		//添加技能buff
		buffHurtEffect := base.BuffEffect(attacker, target, nil, skill, constFight.BUFF_TARGET_ENMY, constFight.BUFF_BY_SKILL_PRE)
		//攻击方 攻击被动触发
		attacker.TriggerPassiveSkill(constFight.SKILL_PASSIVE_CONDITION_ATK, target, skill)
		//攻击敌方伤害
		if target.GetProp().HpNow() > 0 && skill.Target == pb.SKILLTARGETTYPE_EMEMY {
			h := base.Attack(attacker, target, skill, k, isElf)
			//记录攻击伤害到消息里面,一定要先记录攻击伤害 再记录技能前buff伤害
			addHurtEffect(&hurts, h)
			//添加到仇恨列表
			addToThreat(attacker, target, int(-h.ChangHp))
		}
		//将技能伤害计算前buff伤害记录到伤害消息里面
		addHurtEffect(&hurts, buffHurtEffect)
		if target.GetProp().HpNow() > 0 {
			//添加技能buff
			buffHurtEffect = base.BuffEffect(attacker, target, nil, skill, constFight.BUFF_TARGET_ENMY, constFight.BUFF_BY_SKILL_AFTER)
			addHurtEffect(&hurts, buffHurtEffect)
		}
		target.TriggerPassiveSkill(constFight.SKILL_PASSIVE_CONDITION_BE_ATK, attacker, nil)

	}
	//攻击方buff添加
	var attackTarget base.Actor
	if len(targets) > 0 {
		attackTarget = targets[0]
	}
	selfBuffHurt = base.BuffEffect(attacker, attacker, attackTarget, skill, constFight.BUFF_TARGET_SELF, constFight.BUFF_BY_SKILL_AFTER)
	addHurtEffect(&hurts, selfBuffHurt)
	//攻速buff影响
	attacker.ApsdBuff(targets)
	return base.NotifyAttackEffect(attacker, skill, dir, hurts, castPoint, moveToPoint, isElf), nil
}

/**
*  @Description: 天赋触发
*  @param attacker	攻击者
*  @param defender	防守者
*  @param skill 技能
**/
func talentEffectCal(attacker, defender base.Actor, skill *base.Skill) {

	talentEffect := skill.GetTalentEffect()
	if talentEffect == nil || len(talentEffect) == 0 {
		return
	}
	logger.Debug("当前技能课触发技能天赋效果，技能Id:%v,天赋效果：%v", skill.Skillid, talentEffect)
	for _, v := range talentEffect {
		effectConf := gamedb.GetEffectEffectCfg(v)
		if effectConf == nil {
			logger.Error("技能释放，获取效果配置一场：%v", v)
			continue
		}
		if len(effectConf.Attribute) > 0 {
			for propId, propV := range effectConf.Attribute {
				attacker.GetProp().SkillTempPropChange(propId, propV)
			}
		}
		if len(effectConf.Buffid) > 0 {
			for _, buffId := range effectConf.Buffid {
				buffConfT := gamedb.GetBuffBuffCfg(buffId)
				if buffConfT == nil {
					logger.Error("技能配置buff:%v，buff配置表中为找到相应buff", buffId)
					continue
				}

				if buffConfT.Target == constFight.BUFF_TARGET_SELF {
					attacker.AddNewBuff(buffId, attacker, nil, false)

				} else {
					if defender != nil {
						defender.AddNewBuff(buffId, attacker, nil, false)
					}
				}
			}
		}
	}
}

/**
*  @Description: 治疗技能释放
*  @param attacker		释放者
*  @param skill			技能
*  @param dir			朝向
*  @param targetIds		目标
*  @param points		释放区域
*  @param isElf			释放精灵技能
*  @return *pb.AttackEffectNtf
*  @return error
**/
func Treat(attacker base.Actor, skill *base.Skill, dir int, targetIds []int, points []*scene.Point, isElf bool) (*pb.AttackEffectNtf, error) {

	//技能标记使用，记录cd
	attacker.UseSkill(skill, isElf)

	//计算技能伤害 推送
	hurtEffects := make([]*pb.HurtEffect, 0)
	//攻击方buff添加
	selfBuffHurt := base.BuffEffect(attacker, attacker, nil, skill, constFight.BUFF_TARGET_SELF, constFight.BUFF_BY_SKILL_PRE)
	addHurtEffect(&hurtEffects, selfBuffHurt)

	treatTargets := make(map[int]int)
	if len(targetIds) > 0 {
		targetActor := attacker.GetFight().GetActorByObjId(targetIds[0])
		if targetActor != nil && targetActor.GetProp().HpNow() > 0 && targetActor.IsCanTreat() {
			treatTargets[targetActor.GetObjId()] = targetActor.GetProp().HpNow()
		}
	}
	for _, v := range points {
		targets := v.GetAllObject()
		for _, v := range targets {
			actor := v.GetContext().(base.Actor)
			if actor.GetUserId() != attacker.GetUserId() {
				continue
			}
			if actor.GetProp().HpNow() <= 0 || actor.GetProp().HpNow() >= actor.GetProp().Get(pb.PROPERTY_HP) {
				continue
			}
			if !actor.IsCanTreat() {
				continue
			}
			if _, ok := treatTargets[actor.GetObjId()]; !ok {
				treatTargets[actor.GetObjId()] = actor.GetProp().HpNow()
			}
		}
	}

	targetNum := len(treatTargets)
	for i := 0; i < targetNum; i++ {

		min := math.MaxInt64
		minObjId := 0
		for objId, v := range treatTargets {
			if v < min {
				minObjId = objId
			}
		}
		//移除已找到的目标
		delete(treatTargets, minObjId)
		actor := attacker.GetFight().GetActorByObjId(minObjId)
		if actor == nil {
			continue
		}
		//添加技能buff
		buffHurtEffect := base.BuffEffect(attacker, actor, nil, skill, constFight.BUFF_TARGET_ENMY, constFight.BUFF_BY_SKILL_PRE)

		attMaxPropId, attMinPropId := prop.GetAtkPropIdByJob(attacker.Job())
		attackPropValue := (attacker.GetProp().Get(attMaxPropId) + attacker.GetProp().Get(attMinPropId)) / 2
		skillAtk := 0.0
		if i >= len(skill.LevelT.Atk) {
			skillAtk = skill.LevelT.Atk[len(skill.LevelT.Atk)-1]
		} else {
			skillAtk = skill.LevelT.Atk[i]
		}
		buffEffect := actor.GetTreatEffect()
		changeHp, _ := actor.ChangeHp(int(float64(attackPropValue) * skillAtk * (1 + float64(buffEffect)/base.BASE_RATE)))

		logger.Debug("技能结束，治疗技能：%v,buff加成：%v,最终治疗效果：%v", skill.Skillid, buffEffect, changeHp)

		hurtEffect := base.CreateHurtEffect(minObjId, actor.GetProp().HpNow(), -changeHp, false, false, false, pb.HURTTYPE_NORMAL, 0, 0, 0, 0)
		hurtEffects = append(hurtEffects, hurtEffect)

		//将技能前buff记录到伤害消息里面
		addHurtEffect(&hurtEffects, buffHurtEffect)

		//添加技能后置buff
		buffHurtEffect = base.BuffEffect(attacker, actor, nil, skill, constFight.BUFF_TARGET_ENMY, constFight.BUFF_BY_SKILL_AFTER)
		addHurtEffect(&hurtEffects, buffHurtEffect)

		//达到目标治疗数量，跳出
		if i+1 >= skill.LevelT.TargetRoleNum {
			break
		}
	}

	//攻击方buff添加
	selfBuffHurt = base.BuffEffect(attacker, attacker, nil, skill, constFight.BUFF_TARGET_SELF, constFight.BUFF_BY_SKILL_AFTER)
	addHurtEffect(&hurtEffects, selfBuffHurt)

	attackEffectNtf := base.NotifyAttackEffect(attacker, skill, dir, hurtEffects, nil, nil, false)
	attackEffectNtf.IsElf = isElf
	return attackEffectNtf, nil
}

func yeManChongZhuangSkill(attacker base.Actor, skill *base.Skill, dir int, targetIds []int, points []*scene.Point) (*pb.AttackEffectNtf, error) {

	var attackEffect *pb.AttackEffectNtf
	////技能标记使用，记录cd
	//this.UseSkill(skill, false)
	//
	//attackPoint := this.Point()
	//
	//var target base.Actor
	////第一个碰到的对象
	//index := -1
	//for k, point := range points {
	//	index = k
	//	if point.IsBlock() {
	//		break
	//	}
	//	if point.Equal(attackPoint) {
	//		continue
	//	}
	//	objs := point.GetAllObject()
	//	if len(objs) > 0 {
	//		for _, v := range objs {
	//			targetActor := v.GetContext().(base.Actor)
	//			if targetActor.CanAttack() {
	//				target = targetActor
	//				break
	//			}
	//		}
	//	}
	//
	//	if target != nil {
	//		break //跳出来
	//	}
	//}
	//
	//var casterEndPoint *scene.Point
	//var targetEndPoint *scene.Point
	//
	//if target == nil || !this.context.IsEnemy(target) { //没找到目标
	//	//寻找玩家最后可以站立的点
	//	if index > -1 {
	//		casterEndPoint = points[index]
	//	}
	//	target = nil //找不到要把目标置空
	//} else {
	//
	//	finalIndex := index //格子索引
	//	//判断是否可冲撞、抗拒
	//	if this.canYeManChongZhuang(target) {
	//		//可以撞，寻找能撞的最后位置
	//		for i := index + 1; i < len(points); i++ {
	//
	//			point := points[i]
	//			if point.IsBlock() {
	//				break
	//			}
	//			if len(point.GetAllObject()) > 0 {
	//				break
	//			}
	//			finalIndex = i
	//		}
	//	}
	//
	//	if finalIndex > 0 {
	//		casterEndPoint = points[finalIndex-1]
	//	}
	//	targetEndPoint = points[finalIndex]
	//}
	//var castPoint *scene.Point
	//if casterEndPoint != nil && casterEndPoint != this.Point() {
	//	castPoint = this.Point()
	//	this.MoveTo(casterEndPoint, pb.MOVETYPE_WALK, true, false)
	//}
	//
	//if target != nil && targetEndPoint != target.Point() {
	//	if targetEndPoint != nil {
	//		target.MoveTo(targetEndPoint, pb.MOVETYPE_WALK, true, false)
	//	}
	//}
	//
	////计算技能伤害 推送
	//hurts := make([]*pb.HurtEffect, 0)
	//
	//
	////攻速buff判断
	//this.context.GetAllBuffsPb()
	//
	//selfBuffHurt := base.BuffEffect(this.context, this.context, nil, skill, constFight.BUFF_TARGET_SELF, constFight.BUFF_BY_SKILL_PRE)
	//addHurtEffect(&hurts, selfBuffHurt)
	//
	//if target != nil {
	//	buffHurtEffect := base.BuffEffect(this.context, target, nil, skill, constFight.BUFF_TARGET_ENMY, constFight.BUFF_BY_SKILL_PRE)
	//	if target.GetProp().HpNow() > 0 {
	//		h := base.Attack(this.context, target, skill, 0, false)
	//		//添加到仇恨列表
	//		this.addToThreat(this.context, target, int(-h.ChangHp))
	//		if targetEndPoint != nil {
	//			h.MoveToPoint = targetEndPoint.ToPbPoint()
	//		}
	//		addHurtEffect(&hurts, h)
	//	}
	//	if targetEndPoint != nil && buffHurtEffect != nil {
	//		buffHurtEffect.MoveToPoint = targetEndPoint.ToPbPoint()
	//	}
	//	addHurtEffect(&hurts, buffHurtEffect)
	//
	//	if target.GetProp().HpNow() > 0 {
	//		//添加技能buff
	//		buffHurtEffect = base.BuffEffect(this.context, target, nil, skill, constFight.BUFF_TARGET_ENMY, constFight.BUFF_BY_SKILL_AFTER)
	//		if targetEndPoint != nil && buffHurtEffect != nil {
	//			buffHurtEffect.MoveToPoint = targetEndPoint.ToPbPoint()
	//		}
	//		addHurtEffect(&hurts, buffHurtEffect)
	//	}
	//}
	////技能后buff
	//selfBuffHurt = base.BuffEffect(this.context, this.context, nil, skill, constFight.BUFF_TARGET_SELF, constFight.BUFF_BY_SKILL_AFTER)
	//addHurtEffect(&hurts, selfBuffHurt)
	//
	//attackEffect = base.NotifyAttackEffect(this, skill, dir, hurts, castPoint, casterEndPoint, false)
	//
	////攻速buff影响
	//if target != nil {
	//	this.apsdBuff([]base.Actor{target})
	//}

	return attackEffect, nil
}

func canYeManChongZhuang(target base.Actor) bool {

	return true
}

func kangJuHuoHuanSkill(attacker base.Actor, skill *base.Skill, dir int, targetIds []int, points []*scene.Point) (*pb.AttackEffectNtf, error) {

	var attackEffect *pb.AttackEffectNtf
	////技能标记使用，记录cd
	//this.UseSkill(skill, false)
	//
	//attackPoint := this.Point()
	//
	////计算技能伤害 推送
	//hurts := make([]*pb.HurtEffect, 0)
	//
	//for _, point := range points {
	//	if point.IsBlock() {
	//		continue
	//	}
	//	if point.Equal(attackPoint) {
	//		continue
	//	}
	//	objs := point.GetAllObject()
	//	if len(objs) > 0 {
	//		for _, v := range objs {
	//
	//			t := v.GetContext().(base.Actor)
	//			if !this.context.IsEnemy(t) {
	//				continue
	//			}
	//
	//			targetDir := scene.GetFaceDirByPoint(this.Point(), point)
	//			targetPoint := point
	//			for i := 0; i < constFight.SKILL_KANGJU_MOVE_DIS; i++ {
	//				tempPoint := targetPoint.GetNewNearPointByDir(targetDir)
	//				if tempPoint == nil || tempPoint.IsBlock() {
	//					break
	//				}
	//				if len(tempPoint.GetAllObject()) > 0 {
	//					break
	//				}
	//				targetPoint = tempPoint
	//			}
	//			if !point.Equal(targetPoint) {
	//				t.MoveTo(targetPoint, pb.MOVETYPE_WALK, true, false)
	//			}
	//
	//			h := base.Attack(this.context, t, skill, 0, false)
	//			this.addToThreat(this.context, t, int(-h.ChangHp))
	//			h.MoveToPoint = targetPoint.ToPbPoint()
	//
	//			hurts = append(hurts, h)
	//			if t.GetProp().HpNow() > 0 {
	//				//添加技能buff
	//				base.BuffEffect(this.context, t, nil, skill, constFight.BUFF_TARGET_ENMY, constFight.BUFF_BY_SKILL_AFTER)
	//			}
	//		}
	//	}
	//}
	//base.BuffEffect(this.context, this.context, nil, skill, constFight.BUFF_TARGET_SELF, constFight.BUFF_BY_SKILL_AFTER)
	//attackEffect = base.NotifyAttackEffect(this, skill, dir, hurts, nil, nil, false)

	return attackEffect, nil
}

/**
*  @Description: 召唤类技能
*  @param attacker				释放者
*  @param skill					技能
*  @param dir					朝向
*  @return *pb.AttackEffectNtf
*  @return error
**/
func SummonSkillCast(attacker base.Actor, skill *base.Skill, dir int) (*pb.AttackEffectNtf, error) {

	//技能标记使用，记录cd
	attacker.UseSkill(skill, false)

	player := attacker.(base.ActorPlayer).GetPlayer()
	actorSummons := player.SummonActors()
	summonConf := gamedb.GetSummonConfCfg(skill.LevelT.Summonid)

	if len(actorSummons) > 0 {
		groupMap := make(map[int]bool)
		if len(summonConf.Group) > 0 {
			for _, v := range summonConf.Group {
				groupMap[v] = true
			}
		}
		sameSummonNum := 0
		var sameSummonObj *SummonActor

		for _, v := range actorSummons {
			summonActor := v.(*SummonActor)
			if groupMap[summonActor.summonId] {
				attacker.GetFight().Leave(v)
			} else if summonActor.summonId == summonConf.Id {
				sameSummonNum += 1
				if sameSummonObj == nil || sameSummonObj.createTime > summonActor.createTime {
					sameSummonObj = summonActor
				}
			}
		}
		if sameSummonNum >= summonConf.Max && sameSummonObj != nil {
			attacker.GetFight().Leave(sameSummonObj)
		}
	}
	attacker.GetFight().EnterSummon(attacker, summonConf.Id)
	attackEffect := base.NotifyAttackEffect(attacker, skill, dir, []*pb.HurtEffect{}, nil, nil, false)
	return attackEffect, nil
}

/**
*  @Description: 整理技能伤害消息
*  @param hurts
*  @param hurtEffect
**/
func addHurtEffect(hurts *[]*pb.HurtEffect, hurtEffect *pb.HurtEffect) {
	if hurtEffect == nil {
		return
	}
	has := false
	for k, v := range *hurts {
		if v.ObjId == hurtEffect.GetObjId() {
			(*hurts)[k].Hp += hurtEffect.Hp
			if hurtEffect.IsDeath {
				(*hurts)[k].IsDeath = true
			}
			has = true
		}
	}
	//如果伤害消息中没有，则添加
	if !has {
		*hurts = append(*hurts, hurtEffect)
	}
}

/**
*  @Description: 获取攻击目标
*  @param attacker		攻击者
*  @param skill			技能
*  @param points		技能释放区域
*  @param targerIds		初始目标Id
*  @return []base.Actor
**/
func getTargetByPoint(attacker base.Actor, skill *base.Skill, points []*scene.Point, targerIds []int) []base.Actor {
	fight := attacker.GetFight()
	targets := make([]base.Actor, 0)
	targetType := -1
	//计算攻击区域内所有战斗单元的距离
	players := make(map[int]int)
	if len(targerIds) > 0 {
		targerId := targerIds[0]
		target := fight.GetActorByObjId(targerId)
		if target != nil && target.CanAttack() {
			inAttackArea := false
			for _, v := range points {
				if v.Equal(target.Point()) {
					inAttackArea = true
				}
			}
			if inAttackArea && target.CanAttack() {
				//判断技能攻击方
				if skill.Target == pb.SKILLTARGETTYPE_EMEMY && attacker.IsEnemy(target) {
					targets = append(targets, target)
					targetType = target.GetType()
					players[target.GetUserId()] = 1
				} else if skill.Target == pb.SKILLTARGETTYPE_FRIEND && attacker.IsFriend(target) {
					targets = append(targets, target)
					targetType = target.GetType()
					players[target.GetUserId()] = 1
				}
			}
		}
	}

	if len(targets) > 0 {
		if checkSkillTargetMaxNum(skill, targetType, players, len(targets)) {
			return targets
		}
	}

	for _, v := range points {
		objs := v.GetAllObject()
		for _, obj := range objs {

			target := fight.GetActorByObjId(obj.GetObjId())
			//logger.Debug("玩家：%v,技能:%v 释放 查找目标，坐标：%v,目标：%v-%v", this.context.NickName(), skill.Skillid, v.ToString(), target.NickName(), target.GetObjId())
			if target == nil || !target.GetVisible() {
				continue
			}
			if !target.CanAttack() {
				continue
			}
			if len(targerIds) > 0 && targerIds[0] == target.GetObjId() {
				continue
			}

			if targetType != -1 && !targetSameType(targetType, target.GetType()) {
				continue
			}

			//判断技能攻击方
			if (skill.Target == pb.SKILLTARGETTYPE_EMEMY && attacker.IsEnemy(target)) || (skill.Target == pb.SKILLTARGETTYPE_FRIEND && attacker.IsFriend(target)) {
				if targetType == -1 {
					targetType = target.GetType()
				}
				//目标是怪物判断
				if targetType == pb.SCENEOBJTYPE_MONSTER {
					targets = append(targets, target)
				} else {

					//玩家数量是否已满足需求
					if len(players) < skill.LevelT.Num_max {
						if players[target.GetUserId()] < skill.LevelT.TargetRoleNum {
							players[target.GetUserId()] += 1
							targets = append(targets, target)
						}
					} else {
						//玩家角色数量是否满足攻击需求
						if num, ok := players[target.GetUserId()]; ok {
							if num < skill.LevelT.TargetRoleNum {
								players[target.GetUserId()] += 1
								targets = append(targets, target)
							}
						}
					}
				}
				//检查技能攻击最大目标数量
				if checkSkillTargetMaxNum(skill, targetType, players, len(targets)) {
					return targets
				}
			}
		}
	}

	return targets
}

/**
*  @Description: 检查是否相同目标类型（技能攻击要么攻击怪物，要么攻击玩家，不能同时攻击）
*  @param targetType
*  @param actorType
*  @return bool
**/
func targetSameType(targetType int, actorType int) bool {

	if targetType == pb.SCENEOBJTYPE_MONSTER {
		return targetType == actorType
	} else {
		return actorType != pb.SCENEOBJTYPE_MONSTER
	}
}

/**
*  @Description: 检查技能攻击上限
*  @param skill			技能
*  @param targetType	目标类型
*  @param playerTargetNum	玩家数量
*  @param monsterTargetNum	怪物数量
*  @return bool
**/
func checkSkillTargetMaxNum(skill *base.Skill, targetType int, playerTargetNum map[int]int, monsterTargetNum int) bool {

	if targetType == pb.SCENEOBJTYPE_MONSTER {

		if skill.LevelT.MonsterNum > monsterTargetNum {
			return false
		}
	} else {

		if len(playerTargetNum) < skill.LevelT.Num_max {
			return false
		} else {
			for _, v := range playerTargetNum {
				if v < skill.LevelT.TargetRoleNum {
					return false
				}
			}
		}
	}
	return true
}

/**
 *  @Description: 添加到仇恨值列表 目前只有攻击方为玩家 被攻击方为怪物才会添加
 *  @param attacker
 *  @param defender
 *  @param hurt
 */
func addToThreat(attacker, defender base.Actor, hurt int) {
	if defender.GetType() == base.ActorTypeMonster {
		if attacker.GetType() == pb.SCENEOBJTYPE_USER || attacker.GetType() == pb.SCENEOBJTYPE_FIT || attacker.GetType() == pb.SCENEOBJTYPE_SUMMON {
			defender.AddTheat(attacker.GetObjId(), hurt)
		} else if attacker.GetType() == pb.SCENEOBJTYPE_PET {
			defender.AddTheat(attacker.(base.ActorLeader).GetLeader().GetObjId(), hurt)
		}
	}
}
