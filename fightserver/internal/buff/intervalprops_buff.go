package buff

import (
	"cqserver/fightserver/internal/base"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/prop"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"math"
)

type IntervalPropBuff struct {
	*DefaultBuff
	lastRunTime  int64
	interval     int64
	arg          []int
	attackTarget base.Actor
}

func NewIntervalPropBuff(act base.Actor, sourceActor base.Actor, buffT *gamedb.BuffBuffCfg, idx int32, arg ...int) *IntervalPropBuff {
	intervalPropBuff := &IntervalPropBuff{}
	intervalPropBuff.DefaultBuff = NewDefaultBuff(buffT, sourceActor, act, intervalPropBuff, idx)
	intervalPropBuff.arg = arg
	intervalPropBuff.interval = 1000
	if buffT.BuffType == pb.BUFFTYPE_BLEEDING_BY_ATK || buffT.BuffType == pb.BUFFTYPE_RECOVERY_HP {
		if interval, ok := buffT.Effect[1]; ok {
			intervalPropBuff.interval = int64(interval)
		}
	} else if buffT.BuffType == pb.BUFFTYPE_GET_OTHER_BUFF {
		intervalPropBuff.interval = int64(buffT.Effect[0]) * 1000
	}

	return intervalPropBuff
}

//添加buff效果 单位时间改变
func (this *IntervalPropBuff) OnAdd(buffHpChangeInfos *[]*pb.BuffHpChangeInfo, delBuffInfos *[]*pb.DelBuffInfo) {

	//this.calcAddHpMp(buffHpChangeInfos)
	//this.calcLessHp(buffHpChangeInfos)
	//this.lastRunTime = common.GetNowMillisecond()
}

func (this *IntervalPropBuff) Run(buffHpChangeInfos *[]*pb.BuffHpChangeInfo, arg ...interface{}) {
	//流血的判断
	now := common.GetNowMillisecond()
	if now-this.lastRunTime < this.interval {
		return
	}
	if this.owner.GetProp().HpNow() <= 0 {
		return
	}
	this.lastRunTime = now
	this.calcAddHpMp(buffHpChangeInfos, arg[1].(int))
	iswudi := false
	if len(arg) > 2 {
		if wudi, ok := arg[2].(bool); ok {
			iswudi = wudi
		}
	}
	if !iswudi {
		this.calcLessHp(buffHpChangeInfos, arg[0].(int))
	}
	this.createOtherBuff(buffHpChangeInfos)
}

func (this *IntervalPropBuff) setAttackTarget(attackTarget base.Actor) {
	this.attackTarget = attackTarget
}

func (this *IntervalPropBuff) calcAddHpMp(buffHpChangeInfos *[]*pb.BuffHpChangeInfo, recoveHpDec int) {

	//血量变化值
	changeHp := 0
	changeMp := 0
	confValue := this.buffT.Effect[0]
	switch this.buffT.BuffType {
	case pb.BUFFTYPE_ADD_HP_BY_SKILL_PRO, pb.BUFFTYPE_RECOVERY_HP:
		changeHp = int(float64(this.owner.GetProp().Get(pb.PROPERTY_HP)) * float64(confValue) / 10000)
	case pb.BUFFTYPE_ADD_HP_BY_FIXED, pb.BUFFTYPE_ADD_HP_BY_SKILL_FIXED:
		changeHp = confValue
	case pb.BUFFTYPE_ADD_MP_BY_FIXED:
		changeMp = confValue
	case pb.BUFFTYPE_ADD_HP_BY_SKILL_HURT_PRO:
		if len(this.arg) > 0 {
			//skillLvConf := gamedb.GetSkillLevelSkillCfg(this.arg[0])
			//if skillLvConf != nil  {
			propMaxId, propMinId := prop.GetAtkPropIdByJob(this.owner.Job())
			att := common.RandNum(this.owner.GetProp().Get(propMinId), this.owner.GetProp().Get(propMaxId))
			scale := float64(this.buffT.Effect[0]) / 10000
			changeHp = int(float64(att) * scale)
			logger.Debug("buff 技能恢复伤害百分比生命，玩家：%v,攻击值：%v,技能：%v,最总回血：%v", this.owner.NickName(), att, this.arg[0], changeHp)
			//} else {
			//	logger.Error("获取技能伤害回血buff,获取技能配置错误，玩家：%v,buff:%v,技能Id：%v", this.owner.NickName(), this.buffT.Id, this.arg[0])
			//}
		}
	case pb.BUFFTYPE_ADD_HP_MP_FIXED:
		changeHp = this.buffT.Effect[pb.PROPERTY_HP]
		changeMp = this.buffT.Effect[pb.PROPERTY_MP]
	default:
		return
	}
	if recoveHpDec > 0 {
		changeHp = int(float64(changeHp) * math.Max(1-float64(recoveHpDec)/base.BASE_RATE, 0))
	}

	logger.Debug("buff血量增加,玩家：%v,buffId:%v,增加血量：%v,蓝量变化：%v,回血减少：%v", this.owner.NickName(), this.buffT.Id, changeHp, changeMp, recoveHpDec)
	if changeHp > 0 {
		realChangeHp, _ := this.owner.ChangeHp(changeHp)
		hurtEffect := &pb.BuffHpChangeInfo{}
		hurtEffect.Death = 0
		hurtEffect.ChangeHp = int64(realChangeHp)
		hurtEffect.Idx = this.GetBuffIdx()
		hurtEffect.OwnerObjId = int32(this.GetOwenr().GetObjId())
		hurtEffect.TotalHp = int64(this.GetOwenr().GetProp().HpNow())
		*buffHpChangeInfos = append(*buffHpChangeInfos, hurtEffect)
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

func (this *IntervalPropBuff) calcLessHp(buffHpChangeInfos *[]*pb.BuffHpChangeInfo, hurtDec int) {

	hurtDecFloat := float64(hurtDec) / base.BASE_RATE
	//血量变化值
	changeHp := 0
	confValue := this.buffT.Effect[0]
	switch this.buffT.BuffType {
	case pb.BUFFTYPE_POISONING_PRO, pb.BUFFTYPE_BLEEDING_PRO, pb.BUFFTYPE_BURNING_PRO:
		changeHp = int(float64(this.owner.GetProp().Get(pb.PROPERTY_HP)) * float64(confValue) / 10000)
	case pb.BUFFTYPE_POISONING:
		changeHp = confValue
	case pb.BUFFTYPE_POISONING_BY_ATK, pb.BUFFTYPE_BURNING_BY_ATK, pb.BUFFTYPE_BLEEDING_BY_ATK, pb.BUFFTYPE_POISONING_BY_ATK_FOR_MONSTER:
		if this.source != nil {
			var baseValue int
			if this.source.GetType() == pb.SCENEOBJTYPE_PET {
				baseValue = this.source.GetProp().Get(pb.PROPERTY_ATT_PETS)
			} else {
				propMaxId, propMinId := prop.GetAtkPropIdByJob(this.source.Job())
				baseValue = (this.source.GetProp().Get(propMaxId) + this.source.GetProp().Get(propMinId)) / 2
			}

			changeHp = int(float64(baseValue) * (float64(confValue) / 10000))
		}
	default:
		return
	}

	changeHp = int(float64(changeHp) * (1 - hurtDecFloat))

	if this.owner.GetType() == pb.SCENEOBJTYPE_MONSTER {
		changeHp = base.MonsterHurtLimit(this.source, this.owner, changeHp)
	}

	hurtEffect := &pb.BuffHpChangeInfo{}
	hurtEffect.Death = 0
	//记录玩家在战斗中
	this.owner.SetInFightTime()
	realChangeHp, isDeath := this.owner.ChangeHp(- changeHp)
	if isDeath {
		hurtEffect.Death = 1
		this.owner.SetVisible(false)
		this.owner.SetKiller(this.source)
		//this.owner.OnDie()
		//killer := this.source
		//this.owner.GetFight().ActorDieEvent(this.owner, killer)
	}
	hurtEffect.ChangeHp = int64(realChangeHp)
	hurtEffect.Idx = this.GetBuffIdx()
	hurtEffect.OwnerObjId = int32(this.GetOwenr().GetObjId())
	hurtEffect.TotalHp = int64(this.GetOwenr().GetProp().HpNow())
	if this.source != nil {
		nickName := this.source.NickName()
		if leader, ok := this.source.(base.ActorLeader); ok {
			nickName = leader.GetLeader().NickName()
		}
		hurtEffect.KillerId = int32(this.source.GetObjId())
		hurtEffect.KillerName = nickName
	}
	*buffHpChangeInfos = append(*buffHpChangeInfos, hurtEffect)
	logger.Debug("buff血量减少,玩家：%v,buffId:%v,减少血量：%v，buff伤害减少：%v,是否死亡:%v", this.owner.NickName(), this.buffT.Id, changeHp, hurtDecFloat, this.owner.GetProp().HpNow() <= 0)
}

func (this *IntervalPropBuff) createOtherBuff(buffHpChangeInfos *[]*pb.BuffHpChangeInfo) {

	if this.buffT.BuffType != pb.BUFFTYPE_GET_OTHER_BUFF {
		return
	}
	_, err := this.owner.AddBuff(this.buffT.Effect[1], this.owner, false)
	if err != nil {
		logger.Error("定时产生buff时异常，玩家：%v,异常：%v", this.owner.NickName(), err)
	}
	logger.Debug("buff定时生成其他buff,玩家：%v,当前buff:%v，buff效果：%v", this.owner.NickName(), this.buffT.Id, this.buffT.Effect)

}
