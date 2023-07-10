package buff

import (
	"cqserver/fightserver/internal/base"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
)

func (this *BuffManager) momentBuff(buffConfT *gamedb.BuffBuffCfg, owner, attackTarget base.Actor) bool {

	if buffConfT.Time != 0 {
		return false
	}

	switch buffConfT.BuffType {
	case pb.BUFFTYPE_ADD_HP_BY_FIXED, pb.BUFFTYPE_ADD_MP_BY_FIXED, pb.BUFFTYPE_ADD_HP_MP_FIXED,pb.BUFFTYPE_ADD_HP_BY_SKILL_PRO,pb.BUFFTYPE_ADD_HP_BY_SKILL_FIXED:
		this.addHpMp(buffConfT, owner)
	case pb.BUFFTYPE_SUCK:
		this.suckBuff(buffConfT, owner, attackTarget)
	case pb.BUFFTYPE_CLEAR_SKILL_CD:
		this.clearSkillCd(buffConfT, owner)
	default:
		logger.Error("buff,玩家：%v,添加buff:%v,buffType:%v，buff时间为0，未实现",owner.NickName(),buffConfT.Id,buffConfT.BuffType)
	}
	return true
}

func (this *BuffManager) clearSkillCd(buffConfT *gamedb.BuffBuffCfg, owner base.Actor) {

	if !common.RandByTenShousand(buffConfT.Probability) {
		return
	}
	logger.Debug("触发被动技能buff刷新技能CD,玩家：%v", this.owner.NickName())
	owner.ClearSkillCD()
}

func (this *BuffManager) suckBuff(buffConfT *gamedb.BuffBuffCfg, owner base.Actor, attackTarget base.Actor) {

	if !common.RandByTenShousand(buffConfT.Probability) {
		return
	}
	addHp := int(float64(attackTarget.GetProp().Get(pb.PROPERTY_HP)) * (float64(buffConfT.Effect[0]) / base.BASE_RATE))
	realAddHp, _ := owner.ChangeHp(addHp)
	logger.Debug("触发被动技能buff吸血回复自身,玩家：%v,回血：%v", this.owner.NickName(), realAddHp)
	if realAddHp > 0 {
		//推送血量变化
		HPChangeNtf := &pb.SceneObjHpNtf{
			ObjId:    int32(owner.GetObjId()),
			Hp:       int64(owner.GetProp().HpNow()),
			ChangeHp: int64(realAddHp),
			TotalHp:  int64(owner.GetProp().Get(pb.PROPERTY_HP)),
		}
		owner.NotifyNearby(owner, HPChangeNtf, nil)
	}
}

func (this *BuffManager) addHpMp(buffConfT *gamedb.BuffBuffCfg, owner base.Actor) {

	changeHp := 0
	changeMp := 0
	confValue := buffConfT.Effect[constFight.BUFF_KEY_ZERO]
	switch buffConfT.BuffType {
	case pb.BUFFTYPE_ADD_HP_BY_SKILL_PRO:
		changeHp = int(float64(this.owner.GetProp().Get(pb.PROPERTY_HP)) * float64(confValue) / 10000)
	case pb.BUFFTYPE_ADD_HP_BY_FIXED, pb.BUFFTYPE_ADD_HP_BY_SKILL_FIXED:
		changeHp = confValue
	case pb.BUFFTYPE_ADD_MP_BY_FIXED:
		changeMp = confValue
	case pb.BUFFTYPE_ADD_HP_MP_FIXED:
		changeHp = buffConfT.Effect[pb.PROPERTY_HP]
		changeMp = buffConfT.Effect[pb.PROPERTY_MP]
	default:
		return
	}
	logger.Debug("buff 玩家:%v触发瞬回类型buff：%v，回复hp:%v,mp:%v",owner.NickName(),buffConfT.BuffType,buffConfT.Id,changeHp,changeMp)
	if changeHp > 0 {
		realChangeHp, _ := this.owner.ChangeHp(changeHp)
		HPChangeNtf := &pb.SceneObjHpNtf{
			ObjId:    int32(owner.GetObjId()),
			Hp:       int64(owner.GetProp().HpNow()),
			ChangeHp: int64(realChangeHp),
			TotalHp:  int64(owner.GetProp().Get(pb.PROPERTY_HP)),
		}
		owner.NotifyNearby(owner, HPChangeNtf, nil)
	}
	if changeMp > 0 {
		maxMp := this.owner.GetProp().Get(pb.PROPERTY_MP)
		nowMp := this.owner.GetProp().MpNow()
		if nowMp+changeMp > maxMp {
			changeMp = maxMp - nowMp
		}
		this.owner.GetProp().SetMpNow(nowMp + changeMp)
		MPChangeNtf := &pb.SceneObjMpNtf{
			ObjId:    int32(this.owner.GetObjId()),
			Mp:       int64(this.owner.GetProp().MpNow()),
			ChangeMp: int64(changeMp),
			TotalMp:  int64(maxMp),
		}
		if u, ok := this.owner.(base.ActorUser); ok {
			u.SendMessage(MPChangeNtf)
		}
	}
}
