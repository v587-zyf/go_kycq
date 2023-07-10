package base

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
)

func skillAttackEffectPre(attacker, defender Actor, skill *Skill) (int, int,bool) {

	hurtToTreat := 0
	multiHurt := 0
	for _, v := range skill.LevelT.Effects {

		skillAttackEffectConf := gamedb.GetSkillAttackEffectSkillAttackEffectCfg(v)
		if skillAttackEffectConf == nil {
			logger.Error("获取技能攻击效果配置异常：%v", skill.Skillid)
			continue
		}

		//概率触发判断
		if !common.RandByTenShousand(skillAttackEffectConf.Probability) {
			continue
		}

		target := attacker
		if skillAttackEffectConf.Target == constFight.BUFF_TARGET_ENMY {
			target = defender
		}
		//前提条件判断
		if len(skillAttackEffectConf.BuffType) > 0 {
			buffIds := []int{}
			if len(skillAttackEffectConf.BuffType) > 1 {
				buffIds = skillAttackEffectConf.BuffType[1:]
			}
			has, buffId := target.BuffHasType(skillAttackEffectConf.BuffType[0], buffIds)
			if !has {
				continue
			}
			if skillAttackEffectConf.BuffLeave == 1 && buffId > 0 {
				target.BuffRemoveByBuffId(buffId)
			}
		}

		if skillAttackEffectConf.TargetType>0 && skillAttackEffectConf.TargetType != target.GetType(){
			continue
		}

		//概率触发判断
		if common.RandByTenShousand(skillAttackEffectConf.SeckillPro) {
			logger.Info("触发技能攻击效果 秒杀 攻击方：%v，防守方：%v,技能：%v,效果ID:%v",attacker.NickName(), defender.NickName(), skill.Skillid, skillAttackEffectConf.Id)
			return 0,0,true
		}

		logger.Debug("触发技能攻击效果 攻击方：%v，防守方：%v,技能：%v,效果ID:%v", attacker.NickName(), defender.NickName(), skill.Skillid, skillAttackEffectConf.Id)
		//buff添加
		skillAttackEffectBuffAdd(attacker, defender, skillAttackEffectConf)
		//buff移除
		skillAttackEffectBuffRemove(attacker, defender, skillAttackEffectConf)

		switch skillAttackEffectConf.Type {
		case constFight.SKILL_ATTACK_EFFECT_TYPE_1:
			skillAttackEffectPropChange(attacker, defender, skillAttackEffectConf)
			//buff
		case constFight.SKILL_ATTACK_EFFECT_TYPE_3:
			//buff
			hurtToTreat += skillAttackEffectConf.EnmyEffectValue[0]
		case constFight.SKILL_ATTACK_EFFECT_TYPE_5:
			//buff
			multiHurt += skillAttackEffectConf.SelfEffectValue[0]
		default:
			continue
		}
	}

	return hurtToTreat, multiHurt,false
}

func skillAttackEffectPropChange(attacker, defender Actor, skillAttackEffectConf *gamedb.SkillAttackEffectSkillAttackEffectCfg) {

	if len(skillAttackEffectConf.EnmyAddProp) > 0 && len(skillAttackEffectConf.EnmyAddProp) == len(skillAttackEffectConf.EnmyEffectValue) {
		for k, v := range skillAttackEffectConf.EnmyAddProp {
			defender.GetProp().SkillTempPropChange(v, skillAttackEffectConf.EnmyEffectValue[k])
		}
	}

	//添加属性效果
	if len(skillAttackEffectConf.SelfAddProp) > 0 && len(skillAttackEffectConf.SelfAddProp) == len(skillAttackEffectConf.SelfEffectValue) {
		for k, v := range skillAttackEffectConf.SelfAddProp {
			attacker.GetProp().SkillTempPropChange(v, skillAttackEffectConf.SelfEffectValue[k])
		}
	}
}

func skillAttackEffectBuffRemove(attacker, defender Actor, skillAttackEffectConf *gamedb.SkillAttackEffectSkillAttackEffectCfg) {

	if len(skillAttackEffectConf.EnmybuffRemoveNum) > 0 {
		if skillAttackEffectConf.EnmybuffRemoveNum[0] == -1 {
			defender.DelGoodBuff(skillAttackEffectConf.EnmyBuffRemoveLayer)
		} else {
			//buff
			defender.BuffRemove(skillAttackEffectConf.EnmyBuffRemoveLayer, skillAttackEffectConf.EnmybuffRemoveNum)
		}
	}

	if len(skillAttackEffectConf.SelfbuffRemoveNum) > 0 {
		if skillAttackEffectConf.SelfbuffRemoveNum[0] == -1 {
			attacker.DelDeBuff()
		} else {
			//buff
			attacker.BuffRemove(skillAttackEffectConf.SelfbuffRemoveLayer, skillAttackEffectConf.SelfbuffRemoveNum)
		}
	}
}

func skillAttackEffectBuffAdd(attacker, defender Actor, skillAttackEffectConf *gamedb.SkillAttackEffectSkillAttackEffectCfg) {

	if len(skillAttackEffectConf.SelfAddProp) > 0 {

		for _, v := range skillAttackEffectConf.SelfaddBuff {
			attacker.AddNewBuff(v, attacker, nil, false)
		}
	}

	if len(skillAttackEffectConf.EnmyAddProp) > 0 {

		for _, v := range skillAttackEffectConf.EnmyAddBuff {
			defender.AddNewBuff(v, attacker, nil, false)
		}
	}
}
