package base

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"math"
	"strconv"
)

const BASE_RATE = 10000.0

func randNum(min, max int) int {

	if max < min {
		max = min
	}
	return common.RandNum(min, max)
}

/**
*target 防守者
*skill 释放的技能
*targetskill 防守者技能
*isSkillClash 是否对撞
 */
func Attack(attacker, defender Actor, skill *Skill, targetNo int, isElf bool) *pb.HurtEffect {

	//玩家采集重置
	if defender.GetType() == pb.SCENEOBJTYPE_USER {
		defender.(ActorUser).ResetCollectionStatus()
	}

	var hurtEffect *pb.HurtEffect
	if attacker.GetType() == pb.SCENEOBJTYPE_PET || isElf {
		hurtEffect = petAttack(attacker, defender, skill, isElf)
	} else {
		hurtToTreat, multiHurt, seckill := skillAttackEffectPre(attacker, defender, skill)

		if seckill {
			return seckillCal(attacker, defender, skill)
		}

		//基础伤害值
		attackAtk, hurtType := calcBaseDamage(attacker, defender, skill, targetNo, multiHurt)
		if attackAtk == 0 {
			return CreateHurtEffect(defender.GetObjId(), defender.GetProp().HpNow(), 0, false, false, false, pb.HURTTYPE_NORMAL, 0, 0, 0, 0)
		}

		//判断伤害计算方式
		defenderDef := calcDefense(attacker, defender, skill)

		if attacker.GetType() == pb.SCENEOBJTYPE_MONSTER || attacker.GetType() == pb.SCENEOBJTYPE_SUMMON {
			hurtEffect = monsterAttack(attacker, defender, attackAtk, defenderDef, skill)
		} else if attacker.GetType() == pb.SCENEOBJTYPE_USER || attacker.GetType() == pb.SCENEOBJTYPE_FIT {
			hurtEffect = playerAttack(attacker, defender, attackAtk, defenderDef, skill, hurtType, hurtToTreat)
		} else {
			logger.Error("攻击者类型异常")
			return nil
		}
	}
	//重置技能临时属性
	attacker.GetProp().SkillTempPropReset()
	defender.GetProp().SkillTempPropReset()
	return hurtEffect
}

/**
 *  @Description: 战宠伤害
 *  @param attacker
 *  @param defender
 *  @param skill
 *  @return *pb.HurtEffect
 */
func petAttack(attacker, defender Actor, skill *Skill, isElf bool) *pb.HurtEffect {

	attack := 0.0
	if isElf {
		attack = attacker.(ActorUser).GetElfAttack()
	} else {
		attack = float64(attacker.GetProp().Get(pb.PROPERTY_ATT_PETS)) * (1 + float64(attacker.GetProp().Get(pb.PROPERTY_ATT_PETS_RATE))/BASE_RATE)
	}

	hurt := int(attack * skill.LevelT.Atk[0])
	//护盾伤害吸收
	hurt = defender.GetBuffFinalHurtDecFix(hurt)
	//怪物伤害限制
	hurt = MonsterHurtLimit(attacker, defender, hurt)

	totalHurt, isDeath, isRelive := changeTargetHp(attacker, defender, hurt)

	//记录归属
	defender.AddOwner(attacker, false)

	//写入伤害排行版
	writeIntoRank(attacker.GetFight(), attacker, totalHurt)

	logger.Debug("战斗伤害计算，战宠/精灵部分，是否精灵：%v,攻击方：%v,攻击力：%v，伤害：%v,最终伤害：%v", isElf, attacker.NickName(), attack, hurt, totalHurt)
	return CreateHurtEffect(defender.GetObjId(), defender.GetProp().HpNow(), totalHurt, false, isDeath, isRelive, pb.HURTTYPE_NORMAL, hurt, 0, 0, 0)
}

/**
 *  @Description: 玩家攻击伤害计算逻辑
 *  @param attacker
 *  @param defender
 *  @param attackAtk
 *  @param defenderDef
 *  @param skill
 *  @return *pb.HurtEffect
 */
func playerAttack(attacker, defender Actor, attackAtk, defenderDef int, skill *Skill, hurtType int, hurtToTreat int) *pb.HurtEffect {
	fightC := attacker.GetFight()
	fightCalcTypeRatio := gamedb.GetConf().FightCalcTypeRatio
	if defenderDef > 0 {
		fightCalcTypeRatio = float64(attackAtk) / float64(defenderDef)
	}
	hurt := 0
	if fightCalcTypeRatio >= gamedb.GetConf().FightCalcTypeRatio {
		//hurtType = pb.HURTTYPE_OFFICIAL
		//hurt = pofangCalcDamamge(attacker, defender, attackAtk, defenderDef)
		hurt = attackAtk - defenderDef
	} else {
		hurt = int(float64(attackAtk) * gamedb.GetConf().FightCalcRatio / gamedb.GetConf().FightCalcTypeRatio * math.Pow(2, fightCalcTypeRatio/gamedb.GetConf().FightCalcTypeRatio-1))
	}

	//格挡判断
	unBlock := 0
	if common.RandByTenShousand(defender.GetProp().Get(pb.PROPERTY_BLOCK_RATIO)) {
		unBlock = int(float64(hurt) * float64(defender.GetProp().Get(pb.PROPERTY_UN_BLOCK_RATIO)) / BASE_RATE)
		hurt = hurt - unBlock
		if hurt < 0 {
			hurt = 0
			logger.Error("战斗伤害异常，格挡伤害计算异常：%v", defender.GetProp().Get(pb.PROPERTY_UN_BLOCK_RATIO))
		}
		logger.Debug("战斗伤害计算，攻击方：%v,防守方：%v，技能：%v，触发格挡，格挡减少：%v", attacker.NickName(), defender.NickName(), skill.Skillid, unBlock)
	}

	//计算暴击
	criteMultipile := calcCritAdd(attacker, defender)
	hurt = int(float64(hurt) * criteMultipile)
	if criteMultipile > 1 {
		hurtType = pb.HURTTYPE_CRIT
		attacker.TriggerPassiveSkill(constFight.SKILL_PASSIVE_CONDITION_CRIT, defender, skill)
	}

	//致命一击
	fatalHurtAdd := calcFatalAdd(attacker, defender)
	hurt = int(float64(hurt) * fatalHurtAdd)
	if fatalHurtAdd > 1 {
		hurtType = pb.HURTTYPE_FATAL_HURT
		attacker.TriggerPassiveSkill(constFight.SKILL_PASSIVE_CONDITION_FATAL, defender, skill)
		defender.TriggerPassiveSkill(constFight.SKILL_PASSIVE_CONDITION_BE_FATAL, attacker, nil)
	}

	//职业伤害增加
	jobHurtAdd := jobHurtAddCalc(attacker, defender)
	//最终伤害增加减少
	finalHurtAdd := math.Max(1+float64(attacker.GetProp().Get(pb.PROPERTY_ADD_HURT)-defender.GetProp().Get(pb.PROPERTY_RED_HURT))/BASE_RATE, 0)
	//buff伤害增加减少
	hurtBuffAdd, hurtBuffDec := getDefengerHurtDec(attacker, defender)
	hurt = int(float64(hurt) * jobHurtAdd * finalHurtAdd * math.Max(1-hurtBuffDec, 0) * (1 + hurtBuffAdd))

	totalHurt := 0
	//角色类型切割伤害计算
	cutHurt, killHurt := calcActorCutHurt(attacker, defender)
	totalHurt = hurt + cutHurt + killHurt

	//护盾伤害吸收
	totalHurt = defender.GetBuffFinalHurtDecFix(totalHurt)
	if totalHurt < 0 {
		totalHurt = 0
		logger.Error("战斗伤害计算，伤害值为负值了：%", attacker.NickName(), defender.NickName(), skill.Skillid)
	}

	//计算额外伤害
	deadHurt := common.MaxIntGet(attacker.GetProp().Get(pb.PROPERTY_DEAD_HURT)-defender.GetProp().Get(pb.PROPERTY_UN_DEAD_HURT), 0)
	totalHurt += deadHurt

	//官职伤害
	officialHurt := 1.0
	pvpFightRatio := 1.0
	if defender.GetType() != pb.SCENEOBJTYPE_MONSTER {
		officialHurt = 1 + math.Min(math.Max(float64(attacker.Official()-defender.Official()), 0)/float64(gamedb.GetConf().OfficialHurtParameter[0]), float64(gamedb.GetConf().OfficialHurtParameter[1]))
		pvpFightRatio = float64(gamedb.GetConf().PVPfight) / BASE_RATE
	}
	totalHurt = int(float64(totalHurt) * officialHurt * pvpFightRatio)

	//战力压制计算
	powerRoll := powerRoll(attacker, defender)
	totalHurt = int(float64(totalHurt) * powerRoll)

	//怪物最大伤害修正
	totalHurt = MonsterHurtLimit(attacker, defender, totalHurt)

	//记录归属
	defender.AddOwner(attacker, false)
	//目标扣血
	isDeath := false
	isRelive := false
	//目标反弹伤害
	isReflexHurt := defender.ReflexHurt()
	backHurt := 0
	if !isReflexHurt {
		totalHurt, isDeath, isRelive = changeTargetHp(attacker, defender, totalHurt)
	} else {
		backHurt = totalHurt
		totalHurt = 0
	}
	if isDeath {
		attacker.TriggerPassiveSkill(constFight.SKILL_PASSIVE_KILL_TAREGET, nil, skill)
	}

	//反伤 吸血
	attackHurtEffect(fightC, attacker, defender, totalHurt, hurtToTreat, backHurt, isDeath)
	//写入伤害排行版
	writeIntoRank(fightC, attacker, totalHurt)
	logger.Debug("fightCalcTypeRatio:%v,暴击：%v,职业伤害加成：%v,致命一击伤害加成：%v,"+
		"最终伤害加成：%v 切割伤害：%v,斩杀伤害：%v,额外伤害：%v 最终伤害减少:%v,\n"+
		"官职伤害：%v,临时技能属性：[攻击方：%v,防守方：%v]"+
		"buff(27)伤害增加：%v,buff(6 57)伤害减少：%v,伤害反弹：%v",
		fightCalcTypeRatio, criteMultipile, jobHurtAdd, fatalHurtAdd,
		finalHurtAdd, cutHurt, killHurt, deadHurt, hurtBuffDec,
		officialHurt, attacker.GetProp().SkillTempPropToString(), defender.GetProp().SkillTempPropToString(),
		hurtBuffAdd, hurtBuffDec, isReflexHurt)
	logger.Debug("战斗伤害计算，攻击方：%v,防守方：%v，技能：%v,当前血量：%v,最高血量：%v,最终伤害：%v,死亡：%v",
		attacker.NickName(), defender.NickName(), skill.Skillid, defender.GetProp().HpNow(), defender.GetProp().Get(pb.PROPERTY_HP), totalHurt, isDeath)
	return CreateHurtEffect(defender.GetObjId(), defender.GetProp().HpNow(), totalHurt, false, isDeath, isRelive, hurtType, totalHurt, cutHurt, deadHurt, unBlock, killHurt)
}

func MonsterHurtLimit(attacker, defender Actor, hurt int) int {
	//怪物最大伤害修正
	if defender.GetType() == pb.SCENEOBJTYPE_MONSTER {
		defenderMonsterT := defender.(ActorMonster).GetMonsterT()
		if defenderMonsterT.Maxatt > 0 {
			maxHurt := float64(defender.GetProp().Get(pb.PROPERTY_HP)) * (float64(defenderMonsterT.Maxatt) / BASE_RATE)
			if attacker.GetType() == pb.SCENEOBJTYPE_FIT {
				maxHurt *= float64(gamedb.GetConf().FitMonster)
			}
			hurt = common.MinIntGet(hurt, int(maxHurt))
			logger.Debug("战斗伤害计算，怪物:%v，当前血量：%v,最大血量：%v,伤害:%v，最大伤害：%v", defenderMonsterT.Name, defender.GetProp().HpNow(), defender.GetProp().Get(pb.PROPERTY_HP), hurt, maxHurt)
		}
	}
	return hurt
}

/**
 *  @Description: 怪物伤害计算逻辑
 *  @param attacker
 *  @param defender
 *  @param attackAtk
 *  @param defenderDef
 *  @param skill
 *  @return *pb.HurtEffect
 */
func monsterAttack(attacker, defender Actor, attackAtk, defenderDef int, skill *Skill) *pb.HurtEffect {

	totalHurt := int(float64(attackAtk-defenderDef) * (1 - float64(defender.GetProp().Get(pb.PROPERTY_RED_HURT))/BASE_RATE))
	if totalHurt < 0 {
		totalHurt = 0
	}
	//计算伤害减少（含buff）
	_, hurtDec := getDefengerHurtDec(attacker, defender)
	//护盾伤害吸收
	totalHurt = defender.GetBuffFinalHurtDecFix(totalHurt)

	monsterHurtDec := 1.0
	//if attacker.GetType() == pb.SCENEOBJTYPE_MONSTER && (defender.GetType() == pb.SCENEOBJTYPE_USER || defender.GetType() == pb.SCENEOBJTYPE_FIT) {

	monsterHurtDec = float64(defender.GetProp().Get(pb.PROPERTY_RED_HURT_MON)) / BASE_RATE

	//}
	totalHurt = int(float64(totalHurt) * (1 - hurtDec) * math.Max(1-monsterHurtDec, 0))

	if attacker.GetType() == pb.SCENEOBJTYPE_SUMMON && defender.GetType() == pb.SCENEOBJTYPE_MONSTER {
		//怪物伤害限制
		totalHurt = MonsterHurtLimit(attacker, defender, totalHurt)
	}
	if totalHurt <= 0 {
		totalHurt = 0
	}

	isDeath := false
	isRelive := false
	totalHurt, isDeath, isRelive = changeTargetHp(attacker, defender, totalHurt)

	//记录归属
	defender.AddOwner(attacker, false)
	//写入伤害排行版
	writeIntoRank(attacker.GetFight(), attacker, totalHurt)
	logger.Debug("战斗伤害计算，攻击方怪物：%v,防守方：%v，技能：%v,fightCalcTypeRatio:%v,暴击：%v,切割伤害：%v,致命伤害：%v 最终伤害减少:%v 最终伤害：%v,是否死亡：%v,是否触发复活：%v,怪物伤害减少：%v",
		attacker.NickName(), defender.NickName(), skill.Skillid, 0, 0, 0, 0, 0, totalHurt, isDeath, isRelive, monsterHurtDec)
	return CreateHurtEffect(defender.GetObjId(), defender.GetProp().HpNow(), totalHurt, false, isDeath, isRelive, pb.HURTTYPE_NORMAL, totalHurt, 0, 0, 0)
}

//计算伤害减少（含buff）
func getDefengerHurtDec(attacker, defender Actor) (float64, float64) {
	//计算伤害减少（含buff）
	buffHurtDec := defender.GetBuffFinalHurtDec()
	hurtDec := float64(buffHurtDec) / BASE_RATE
	buffHurtAdd := attacker.GetBuffFinalHurtAdd()
	hurtAdd := float64(buffHurtAdd) / BASE_RATE
	return hurtAdd, hurtDec
}

/**
*  @Description:
*  @param objId
*  @param hpNow
*  @param damage
*  @param isDodge
*  @param isDeath
*  @param isRelive
*  @param hurtType
*  @param hurt 伤害[伤害值，切割伤害,额外伤害，格挡减少伤害，斩杀伤害]
*  @return *pb.HurtEffect
**/
func CreateHurtEffect(objId int, hpNow, changeHp int, isDodge, isDeath bool, isRelive bool, hurtType int, hurt ...int) *pb.HurtEffect {
	hurtEffect := &pb.HurtEffect{}
	hurtEffect.ObjId = int32(objId) //被攻击者房间ID
	hurtEffect.ChangHp = int64(-changeHp)
	hurtEffect.Hp = int64(hpNow)
	hurtEffect.IsDeath = isDeath
	hurtEffect.ReliveSelf = isRelive
	hurtEffect.IsDodge = isDodge
	hurtEffect.HurtType = int32(hurtType)
	hurtEffect.Hurt = 0
	if len(hurt) > 0 {
		hurtEffect.Hurt = int64(hurt[0])
	}
	hurtEffect.CutHurt = 0
	if len(hurt) > 1 {
		hurtEffect.CutHurt = int64(hurt[1])
	}
	hurtEffect.Deathblow = 0
	if len(hurt) > 2 {
		hurtEffect.Deathblow = int64(hurt[2])
	}
	hurtEffect.UnBlock = 0
	if len(hurt) > 3 {
		hurtEffect.UnBlock = int64(hurt[3])
	}
	hurtEffect.KillHurt = 0
	if len(hurt) > 4 {
		hurtEffect.KillHurt = int64(hurt[4])
	}
	return hurtEffect
}

func changeTargetHp(caster, target Actor, damage int) (int, bool, bool) {

	fightC := caster.GetFight()
	var decHp int
	var isDeath bool
	var isRelive bool
	if target.GetType() == ActorTypeMonster && (fightC.GetStageConf().Type == constFight.FIGHT_TYPE_WORLDBOSS || fightC.GetStageConf().Type == constFight.FIGHT_TYPE_CROSS_WORLD_LEADER) {

		//世界boss boss假扣血
		decHp = damage
	} else {
		realChange := 0
		realChange, isDeath = target.ChangeHp(-damage)
		decHp = -realChange
	}
	caster.CalDamage(decHp)
	fightC.PostDamage(caster, target, decHp)
	//如果目标死亡
	if isDeath {
		target.SetVisible(false)
		target.SetKiller(caster)
		//isRelive = target.OnDie()
		//killer := caster
		//fightC.ActorDieEvent(target, killer)
	}
	return decHp, isDeath, isRelive
}

//玩家是否闪避
func CalcIsDodge(attacker, defender Actor, skill *Skill) bool {

	attackerHit := attacker.GetProp().Get(pb.PROPERTY_HIT)
	defenderMiss := defender.GetProp().Get(pb.PROPERTY_MISS)

	if defenderMiss == 0 {
		return false
	}

	//必命中
	hitMust := defender.CheckTriggerPropMust(attacker, pb.PROPERTY_HIT)
	if hitMust {
		return false
	}

	pow := (float64(attackerHit+gamedb.GetConf().FightHitFix)/float64(defenderMiss+gamedb.GetConf().FightHitFix) - 1)
	if pow > 1 {
		//防止数据溢出
		pow = 1
	}

	hitRate := gamedb.GetConf().FightHitBase * math.Pow(2, pow)
	logger.Debug("战斗伤害计算，攻击方：%v,防守方：%v，技能：%v,命中率计算,命中：%v，闪避：%v,命中率：%v", attacker.NickName(), defender.NickName(), skill.Skillid, attackerHit, defenderMiss, hitRate)
	if common.RandByTenShousand(int(hitRate * BASE_RATE)) {
		return false
	}
	return true
}

func seckillCal(attacker, defender Actor, skill *Skill) *pb.HurtEffect {

	totalHurt := defender.GetProp().HpNow()
	isDeath := true
	isRelive := false
	totalHurt, isDeath, isRelive = changeTargetHp(attacker, defender, totalHurt)

	if isDeath {
		attacker.TriggerPassiveSkill(constFight.SKILL_PASSIVE_KILL_TAREGET, nil, skill)
	}

	//写入伤害排行版
	writeIntoRank(attacker.GetFight(), attacker, totalHurt)
	return CreateHurtEffect(defender.GetObjId(), defender.GetProp().HpNow(), totalHurt, false, isDeath, isRelive, 0, 0, 0, 0, 0)
}

func calcBaseDamage(attack, defender Actor, skill *Skill, targetNo int, multiHurt int) (int, int) {

	luck := attack.GetProp().Get(pb.PROPERTY_LUCKY)
	luckFix := gamedb.GetLuckyLuckyCfg(luck).AddAtk
	job := attack.Job()
	minAttPropId := 0
	maxAttPropId := 0
	//战士职业 普通攻击取物理攻击
	if job == pb.JOB_ZHANSHI || skill.Type == pb.SKILLTYPE_ORDINARY {
		minAttPropId = pb.PROPERTY_PATT_MIN
		maxAttPropId = pb.PROPERTY_PATT_MAX
	} else if job == pb.JOB_FASHI {
		minAttPropId = pb.PROPERTY_MATT_MIN
		maxAttPropId = pb.PROPERTY_MATT_MAX
	} else {
		minAttPropId = pb.PROPERTY_TATT_MIN
		maxAttPropId = pb.PROPERTY_TATT_MAX
	}
	//计算攻击力
	minAtt := attack.GetProp().Get(minAttPropId)
	maxAtt := attack.GetProp().Get(maxAttPropId)
	minAtt = minAtt + int(float64(maxAtt-minAtt)*luckFix)
	if maxAtt < minAtt {
		maxAtt = minAtt
	}
	//随机攻击力
	randAtk := randNum(minAtt, maxAtt)

	hurtType := pb.HURTTYPE_NORMAL
	//狂暴一击计算
	rageHurt := 0
	if attack.GetType() == pb.SCENEOBJTYPE_USER || attack.GetType() == pb.SCENEOBJTYPE_FIT {
		rageRatio := attack.GetProp().Get(pb.PROPERTY_RAGE_RATIO) - defender.GetProp().Get(pb.PROPERTY_UN_RAGE_RATIO)
		rageRatio = common.MaxIntGet(rageRatio, 0)
		if common.RandByTenShousand(rageRatio) {
			rageHurt = common.MaxIntGet(attack.GetProp().Get(pb.PROPERTY_RAGE_HURT)-defender.GetProp().Get(pb.PROPERTY_UN_RAGE_HURT), 0)
			hurtType = pb.HURTTYPE_RAGE
			//狂暴一击触发被动
			attack.TriggerPassiveSkill(constFight.SKILL_PASSIVE_CONDITION_RAGE, defender, skill)
		}
	}
	randAtk += rageHurt

	//技能攻击目标
	atkRatio := 0.0
	atkLen := len(skill.LevelT.Atk)
	if atkLen > 0 {
		if targetNo >= atkLen {
			atkRatio = skill.LevelT.Atk[atkLen-1]
		} else {
			atkRatio = skill.LevelT.Atk[targetNo]
		}
	}
	//buff指定技能伤害增加
	skillBuffHurtAddBySkill := float64(attack.GetBuffSkillHurtAddBySkillId(skill.Skillid)) / BASE_RATE
	//buff技能伤害增加
	skillBuffHurtAdd := float64(attack.GetBuffSkillHurtAdd()) / BASE_RATE
	//火焰buff
	skillBuffFireHurtAdd := 0.0
	for _, v := range gamedb.GetConf().FireSkills {
		if v == skill.Skillid {
			skillBuffFireHurtAdd = float64(attack.GetBuffFireSkillHurtAdd()) / BASE_RATE
			break
		}
	}

	//多倍伤害
	multiHurtFloat := 1.0
	if float64(multiHurt) > 0 {
		multiHurtFloat = float64(multiHurt) / BASE_RATE
	}

	//属性技能伤害增加
	skillHurtAdd := float64(attack.GetProp().Get(pb.PROPERTY_ADD_SKILL)-defender.GetProp().Get(pb.PROPERTY_RED_SKILL)) / BASE_RATE
	//伤害系数计算
	atkRatio = math.Max(atkRatio*(1+skillBuffHurtAddBySkill+skillBuffHurtAdd+skillBuffFireHurtAdd)*multiHurtFloat+skillHurtAdd, 0)
	//最终伤害值
	finalAtk := int((float64(randAtk) * atkRatio) + float64(skill.LevelT.Atk2))
	//合体修正
	fitAttFix := 1.0
	if attack.GetType() == pb.SCENEOBJTYPE_FIT && defender.GetType() != pb.SCENEOBJTYPE_MONSTER {
		fitAttFix = float64(gamedb.GetConf().FitPVP) / BASE_RATE
	}
	finalAtk = int(fitAttFix * float64(finalAtk))
	logger.Debug("战斗伤害计算 基础伤害值,目标：%v 最小攻击：%v，最大攻击：%v,狂暴一击：%v,攻击值：%v,攻击加成：[技能百分比加成：%v,skillBuffHurtAddBySkill:%v,skillBuffHurtAdd:%v, skillHurtAdd:%v,火焰系buff加成：%v,多倍：%v],合体修正：%v,最终攻击：%v",
		targetNo+1, minAtt, maxAtt, rageHurt, randAtk, atkRatio, skillBuffHurtAddBySkill, skillBuffHurtAdd, skillHurtAdd, skillBuffFireHurtAdd, multiHurtFloat, fitAttFix, finalAtk)
	return finalAtk, hurtType
}

func calcDefense(attacker, defender Actor, skill *Skill) int {

	job := attacker.Job()
	minAttPropId := 0
	maxAttPropId := 0

	if job == pb.JOB_ZHANSHI || skill.Type == pb.SKILLTYPE_ORDINARY {
		minAttPropId = pb.PROPERTY_DEF_MIN
		maxAttPropId = pb.PROPERTY_DEF_MAX
	} else {
		minAttPropId = pb.PROPERTY_ADF_MIN
		maxAttPropId = pb.PROPERTY_ADF_MAX
	}

	minDef := defender.GetProp().Get(minAttPropId)
	maxDef := defender.GetProp().Get(maxAttPropId)
	luck := defender.GetProp().Get(pb.PROPERTY_LUCKY)
	luckFix := gamedb.GetLuckyLuckyCfg(luck).AddAtk
	minDef = minDef + int(float64(maxDef-minDef)*luckFix)
	def := randNum(minDef, maxDef)
	ingoreDef := float64(attacker.GetProp().Get(pb.PROPERTY_IGNORE_DEF)) / BASE_RATE
	finalDef := int(float64(def) * (1 - ingoreDef))
	if finalDef <= 0 {
		finalDef = 0
	}
	logger.Debug("战斗伤害计算 防守方防御计算职业：%v,最小防御：%v，最大防御：%v,随机防御值：%v,无视防御：%v,最终防御：%v", job, minDef, maxDef, def, ingoreDef, finalDef)
	return finalDef
}

func jobHurtAddCalc(attacker, defender Actor) float64 {

	defenderJob := defender.Job()
	attackerJob := attacker.Job()
	attackerJobHurtAdd := 0
	defenderJobHurtDec := 0
	if defenderJob == pb.JOB_ZHANSHI {
		attackerJobHurtAdd = attacker.GetProp().Get(pb.PROPERTY_ADD_W)
	} else if defenderJob == pb.JOB_FASHI {
		attackerJobHurtAdd = attacker.GetProp().Get(pb.PROPERTY_ADD_M)
	} else {
		attackerJobHurtAdd = attacker.GetProp().Get(pb.PROPERTY_ADD_T)
	}
	if attackerJob == pb.JOB_ZHANSHI {
		defenderJobHurtDec = defender.GetProp().Get(pb.PROPERTY_RED_W)
	} else if attackerJob == pb.JOB_FASHI {
		defenderJobHurtDec = defender.GetProp().Get(pb.PROPERTY_RED_M)
	} else {
		defenderJobHurtDec = defender.GetProp().Get(pb.PROPERTY_RED_T)
	}
	jobHurt := math.Max(1+float64(attackerJobHurtAdd-defenderJobHurtDec+attacker.GetProp().Get(pb.PROPERTY_ADD_PLAYER)-defender.GetProp().Get(pb.PROPERTY_RED_PLAYER))/BASE_RATE, 0)
	logger.Debug("战斗伤害计算 破防状态计算,攻击方职业伤害增加：%v,防守方职业伤害减少：%v,最终伤害增加：%v", attackerJobHurtAdd, defenderJobHurtDec, jobHurt)
	return jobHurt
}

/**
 *  @Description: 计算暴击
 *  @param attacker
 *  @param defender
 *  @return float64 返回暴击增加
 */
func calcCritAdd(attacker, defender Actor) float64 {

	//必暴击检查
	isCrit := defender.CheckTriggerPropMust(attacker, pb.PROPERTY_CRIT)
	critRate := 0.0
	if !isCrit {

		critRate = float64(attacker.GetProp().Get(pb.PROPERTY_CRIT)-defender.GetProp().Get(pb.PROPERTY_UN_CRIT)) / BASE_RATE
		if critRate <= 0 {
			return 1
		}
		if critRate < gamedb.GetConf().FightCritRateLimit[0] {
			critRate = gamedb.GetConf().FightCritRateLimit[0]
		}
		if critRate > gamedb.GetConf().FightCritRateLimit[1] {
			critRate = gamedb.GetConf().FightCritRateLimit[1]
		}
		isCrit = common.RandByTenShousand(int(critRate * BASE_RATE))
	}

	crit := math.Max(gamedb.GetConf().FightCritBase+float64(attacker.GetProp().Get(pb.PROPERTY_ADD_CRIT_HIT)-defender.GetProp().Get(pb.PROPERTY_RED_CRIT_HIT))/BASE_RATE, 1)
	logger.Debug("战斗伤害计算，攻击方：%v,防守方：%v，critRate:%v,是否触发暴击：%v,crit:%v", attacker.NickName(), defender.NickName(), critRate, isCrit, crit)
	if !isCrit {
		return 1
	}

	return crit

}

/**
 *  @Description: 计算致命一击
 *  @param attacker
 *  @param defender
 *  @return float64 返回致命一击伤害增加
 */
func calcFatalAdd(attacker, defender Actor) float64 {

	fatalRate := common.MaxIntGet(attacker.GetProp().Get(pb.PROPERTY_FATAL_RATIO)-defender.GetProp().Get(pb.PROPERTY_UN_FATAL_RATIO), 0)
	if !common.RandByTenShousand(fatalRate) {
		return 1
	}

	fatalHurtAdd := math.Max(gamedb.GetConf().FightDeathBase+float64(attacker.GetProp().Get(pb.PROPERTY_ADD_FATAL)-defender.GetProp().Get(pb.PROPERTY_RED_FATAL))/BASE_RATE, 1)
	if fatalHurtAdd > 1 {
		//防守方受致命攻击回血
		defender.BuffFatalRecoveHp()
	}
	logger.Debug("战斗伤害计算，攻击方：%v,防守方：%v，fatalRate:%v,fatalHurtAdd:%v", attacker.NickName(), defender.NickName(), fatalRate, fatalHurtAdd)
	return fatalHurtAdd

}

func powerRoll(attacker, defender Actor) float64 {

	powerRoll := 1.0
	if attacker.GetType() != pb.SCENEOBJTYPE_USER {
		return powerRoll
	}

	attackUserCombat := attacker.(ActorPlayer).GetPlayer().UserCombat()
	defenderCombat := 0
	f := attacker.GetFight()
	isPvp := false
	if defender.GetType() == pb.SCENEOBJTYPE_MONSTER {

		defenderCombat, _ = strconv.Atoi(f.GetPowerRoll())

	} else if defender.GetType() == pb.SCENEOBJTYPE_USER {

		defenderCombat = defender.(ActorPlayer).GetPlayer().UserCombat()
		isPvp = true
	}
	if defenderCombat > 0 {

		powerRoll = 1 - float64(gamedb.GetPowerRoll(common.RoundFloat(float64(attackUserCombat)/float64(defenderCombat), 2), isPvp))/BASE_RATE
	}
	logger.Debug("战斗伤害计算，攻击方：%v,防守方：%v，攻击者战力：%v,防守方战斗力：%v,最终碾压值：%v", attacker.NickName(), defender.NickName(), attackUserCombat, defenderCombat, powerRoll)
	return powerRoll
}

func calcActorCutHurt(attacker, defender Actor) (int, int) {

	if defender.GetType() == pb.SCENEOBJTYPE_USER || defender.GetType() == pb.SCENEOBJTYPE_FIT {
		return 0, 0
	}
	cutHurt := defender.GetProp().Get(pb.PROPERTY_HP)*attacker.GetProp().Get(pb.PROPERTY_CUT_RATIO)/BASE_RATE + attacker.GetProp().Get(pb.PROPERTY_CUT)
	killHurt := 0
	//斩杀判断
	rate := attacker.GetProp().Get(pb.PROPERTY_ADD_KILL) - defender.GetProp().Get(pb.PROPERTY_RED_KILL)
	if rate > 0 && common.HitRateTenThousand(rate) {
		killHurt = int(float64(defender.GetProp().Get(pb.PROPERTY_HP))*(float64(attacker.GetProp().Get(pb.PROPERTY_KILL_RATE))/BASE_RATE)) + attacker.GetProp().Get(pb.PROPERTY_KILL)
	}
	return cutHurt, killHurt
}

func attackHurtEffect(fight Fight, attacker, defender Actor, totalHurt int, hurtToTreat int, backHurt int, isDeath bool) {

	//先计算反伤，反伤死了，就不在计算吸血了
	now := attacker.GetProp().HpNow()
	changeHp := 0
	suckNum := 0
	suckFromMonster := 0
	skillEffectHurtToTheat := 0
	buffEffectHurtToTheat := 0
	reflexHurt := 0
	attackIsDeath := false
	if !isDeath {
		//反伤计算
		reflexHurt = int(float64(totalHurt)*float64(defender.GetProp().Get(pb.PROPERTY_REFLEX_RATIO))/BASE_RATE) + defender.GetProp().Get(pb.PROPERTY_REFLEX) + backHurt
	}

	if now > reflexHurt {
		//技能效果 伤害转为治疗处理
		if hurtToTreat > 0 {
			skillEffectHurtToTheat = int(float64(totalHurt) * float64(hurtToTreat) / BASE_RATE)
		}
		//buff效果
		buffEffect := attacker.GetBuffEffectByBuffType(pb.BUFFTYPE_HURT_TO_TREAT)
		if buffEffect > 0 {
			buffEffectHurtToTheat = int(float64(totalHurt) * float64(buffEffect) / BASE_RATE)
			userNum := 1
			fitActor := fight.GetUserFitActor(attacker.GetUserId())
			if fitActor == nil {
				heros := fight.GetUserByUserId(attacker.GetUserId())
				if heros != nil && len(heros) > 0 {
					userNum = len(heros)
				}
				buffEffectHurtToTheat = buffEffectHurtToTheat / userNum
				if buffEffectHurtToTheat > 0 {
					for _, v := range heros {
						if v.GetObjId() != attacker.GetObjId() {
							temChangeHp, _ := v.ChangeHp(buffEffectHurtToTheat)
							HPChangeNtf := &pb.SceneObjHpNtf{
								ObjId:    int32(v.GetObjId()),
								Hp:       int64(v.GetProp().HpNow()),
								ChangeHp: int64(temChangeHp),
								TotalHp:  int64(v.GetProp().Get(pb.PROPERTY_HP)),
							}
							v.NotifyNearby(v, HPChangeNtf, nil)
						}
					}
				}
			}
		}

		//吸血计算
		suckNumRate := float64(attacker.GetProp().Get(pb.PROPERTY_SUCK_RATIO)-defender.GetProp().Get(pb.PROPERTY_UN_SUCK)) / BASE_RATE
		if suckNumRate < 0 {
			suckNumRate = 0
		}
		suckNum = int(float64(totalHurt)*(suckNumRate)) + attacker.GetProp().Get(pb.PROPERTY_SUCK)
		if defender.GetType() == pb.SCENEOBJTYPE_MONSTER {
			suckFromMonster = int(float64(totalHurt) * (float64(attacker.GetProp().Get(pb.PROPERTY_SUCK_RATIO_MON)) / BASE_RATE))
		}
		changeHp = suckNum + suckFromMonster + skillEffectHurtToTheat + buffEffectHurtToTheat - reflexHurt
	} else {
		reflexHurt = now
		changeHp = -now
		attackIsDeath = true
	}

	//血量有变化，推送客户端
	if changeHp != 0 {
		changeHp, attackIsDeath = attacker.ChangeHp(changeHp)
		HPChangeNtf := &pb.SceneObjHpNtf{
			ObjId:    int32(attacker.GetObjId()),
			Hp:       int64(attacker.GetProp().HpNow()),
			ChangeHp: int64(changeHp),
			TotalHp:  int64(attacker.GetProp().Get(pb.PROPERTY_HP)),
		}
		if attackIsDeath {
			HPChangeNtf.KillerId = int32(defender.GetObjId())
			HPChangeNtf.KillerName = defender.NickName()
		}
		attacker.NotifyNearby(attacker, HPChangeNtf, nil)
	}

	fight.PostDamage(defender, attacker, reflexHurt)
	//如果目标死亡
	if attackIsDeath {
		attacker.SetVisible(false)
		attacker.SetKiller(defender)
	}
	logger.Debug("战斗伤害计算，攻击方：%v,防守方：%v，反伤 吸血计算，反伤：%v,吸血：%v,对怪物吸血：%v，技能效果伤害转换血量：%v,buff效果伤害转血量：%v,实际改变血量：%v",
		attacker.NickName(), defender.NickName(), reflexHurt, suckNum, suckFromMonster, skillEffectHurtToTheat, buffEffectHurtToTheat, changeHp)
}

/**
 *  @Description:写入排行版
 *  @param f
 *  @param attacker
 *  @param damage
 */
func writeIntoRank(f Fight, attacker Actor, damage int) {
	if rFight, ok := f.(IFightDamageRank); ok {
		if attacker.GetType() == pb.SCENEOBJTYPE_USER || attacker.GetType() == pb.SCENEOBJTYPE_PET || attacker.GetType() == pb.SCENEOBJTYPE_FIT || attacker.GetType() == pb.SCENEOBJTYPE_SUMMON {
			mainActor := f.GetUserMainActor(attacker.GetUserId())
			if mainActor != nil {
				rFight.FightDamageRankSetDamage(mainActor, int64(damage))
			}
		}
	}
}
