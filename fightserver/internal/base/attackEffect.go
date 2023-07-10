package base

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
)

//区分技能前buff 技能后buff
func BuffEffect(source, owner Actor, attackTarget Actor, skill *Skill, buffTarget int, skillBuffUse int) *pb.HurtEffect {
	hurtTargetHp := 0

	buffs := skill.LevelT.Buff
	if skillBuffUse == constFight.BUFF_BY_SKILL_PRE {
		buffs = skill.LevelT.Prebuff
	}

	for _, v := range buffs {
		buffConfT := gamedb.GetBuffBuffCfg(v)
		if buffConfT == nil {
			logger.Error("技能配置buff:%v，buff配置表中为找到相应buff", v)
			continue
		}

		if buffConfT.Target == buffTarget {
			hpChange, _ := owner.AddNewBuff(v, source, attackTarget, false, skill.LevelT.Skillid)
			hurtTargetHp += hpChange
		}
	}

	//has := false
	//for k, v := range *hurt {
	//	if int(v.ObjId) == owner.GetObjId() {
	//		(*hurt)[k].Hp += int64(hurtOwnHp)
	//		if owner.GetProp().HpNow() <= 0 {
	//			(*hurt)[k].IsDeath = true
	//		}
	//		has = true
	//	}
	//}
	//如果伤害消息中没有，则添加
	if hurtTargetHp != 0 {
		hurtEffect := &pb.HurtEffect{ObjId: int32(owner.GetObjId()), Hp: int64(owner.GetProp().HpNow()),ChangHp: int64(hurtTargetHp)}
		if owner.GetProp().HpNow() <= 0 {
			hurtEffect.IsDeath = true
		}
		return hurtEffect
	}
	return nil
}
