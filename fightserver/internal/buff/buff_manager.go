package buff

import (
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/common"
	"time"

	"cqserver/fightserver/internal/base"
	"cqserver/gamelibs/errex"
	"cqserver/gamelibs/gamedb"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"sync"
)

type BuffManager struct {
	buffs      sync.Map
	owner      base.Actor
	idx        int32 //buff序号
	groupBuffs map[int]*GroundBuff
}

func NewBuffManager(actor base.Actor) *BuffManager {
	return &BuffManager{
		owner:      actor,
		idx:        0,
		groupBuffs: make(map[int]*GroundBuff, 0),
	}
}

type BuffPriority struct {
	buffPriority int   //优先级
	buffIdx      int32 //序号
	endTime      int64
	buff         base.Buff
	isFlyUp      bool
}

const (
	MaxBuffIdx       = 1000000
	BuffEffectOwn    = 0
	BuffEffectTarget = 1
)

func (this *BuffManager) createBuff(buffT *gamedb.BuffBuffCfg, actor, sourceActor base.Actor, arg ...int) (base.Buff, error) {
	buffIdx := this.IncrBuffIdx()
	switch buffT.BuffType {
	case pb.BUFFTYPE_PARALYSIS, pb.BUFFTYPE_FREEZE, pb.BUFFTYPE_FLAME, pb.BUFFTYPE_INVINCIBLE, pb.BUFFTYPE_DEC_HURT, pb.BUFFTYPE_DEC_HURT_FIXED,
		pb.BUFFTYPE_MARK, pb.BUFFTYPE_BE_FATAL_REOVER_HP, pb.BUFFTYPE_ATTACK_ADD_BY_HP, pb.BUFFTYPE_SILENT, pb.BUFFTYPE_TREAT_BAN, pb.BUFFTYPE_TREAT_ADD,
		pb.BUFFTYPE_TREAT_LESS, pb.BUFFTYPE_ADD_HURT_BY_FIRE, pb.BUFFTYPE_WEAK, pb.BUFFTYPE_TREAT_ADD_BY_HP, pb.BUFFTYPE_HURT_TO_TREAT, pb.BUFFTYPE_HOLD,
		pb.BUFFTYPE_DIZZINESS, pb.BUFFTYPE_CONFUSION, pb.BUFFTYPE_SKILL_HURT_ADD, pb.BUFFTYPE_SKILL_CD, pb.BUFFTYPE_FIT_LIMIT, pb.BUFFTYPE_ADD_HURT, pb.BUFFTYPE_IMMUNE,
		pb.BUFFTYPE_BACK_HURT, pb.BUFFTYPE_TRIGGER_PROP_MUST, pb.BUFFTYPE_MP_TO_HP, pb.BUFFTYPE_SKILL_INVALID, pb.BUFFTYPE_DES_RECOVERY_HP, pb.BUFFTYPE_PETRIFIED_FOR_MONSTER:
		//状态类buff
		return NewStatusBuff(actor, sourceActor, buffT, buffIdx), nil
	case pb.BUFFTYPE_RELIVE:
		return NewStatusReliveBuff(actor, sourceActor, buffT, buffIdx), nil
	case pb.BUFFTYPE_POISONING, pb.BUFFTYPE_POISONING_PRO, pb.BUFFTYPE_BLEEDING_PRO, pb.BUFFTYPE_BURNING_PRO, pb.BUFFTYPE_ADD_HP_BY_FIXED,
		pb.BUFFTYPE_ADD_MP_BY_FIXED, pb.BUFFTYPE_ADD_HP_MP_FIXED, pb.BUFFTYPE_GET_OTHER_BUFF, pb.BUFFTYPE_ADD_HP_BY_SKILL_FIXED, pb.BUFFTYPE_ADD_HP_BY_SKILL_HURT_PRO,
		pb.BUFFTYPE_ADD_HP_BY_SKILL_PRO, pb.BUFFTYPE_POISONING_BY_ATK, pb.BUFFTYPE_BURNING_BY_ATK, pb.BUFFTYPE_BLEEDING_BY_ATK, pb.BUFFTYPE_RECOVERY_HP, pb.BUFFTYPE_POISONING_BY_ATK_FOR_MONSTER:
		//间隔效果类buff
		return NewIntervalPropBuff(actor, sourceActor, buffT, buffIdx, arg...), nil
	case pb.BUFFTYPE_ADD_PROP_BY_PRO, pb.BUFFTYPE_DEC_PROP_BY_PRO, pb.BUFFTYPE_ADD_PROP_BY_FIXED, pb.BUFFTYPE_DEC_PROP_BY_FIXED, pb.BUFFTYPE_ADD_PROP_BY_ATK,
		pb.BUFFTYPE_LESS_PROP_BY_ATK, pb.BUFFTYPE_CUT_SKILL:
		//属性buff
		return NewOnePropsBuff(actor, sourceActor, buffT, buffIdx), nil
	case pb.BUFFTYPE_ASPD_ADD_PRO:
		return NewOnePropsAspdBuff(actor, sourceActor, buffT, buffIdx), nil
	default:
		return nil, errex.NewErrorItem("buff type error", buffT.Id, buffIdx)
	}
}

/**
 *  @Description: 获取相同类型buff
 *  @param buffConfT	将要添加的buff的表配置
 *  @return base.Buff
 */
func (this *BuffManager) sameTypeBuff(buffConfT *gamedb.BuffBuffCfg) (bool, base.Buff) {
	//同类叠加规则，互不影响，直接添加
	if buffConfT.BuffRule == constFight.BUFF_RULE_NO_AFFECT {
		return true, nil
	}

	var sameTypebuff base.Buff
	var sameTypeBuffNum int
	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		buffConf := buff.GetBuffT()
		if buffConf != nil && buffConfT.BuffType == buff.GetBuffT().BuffType && buffConfT.Group == buffConf.Group {
			sameTypeBuffNum += 1
			if buffConfT.Remove != constFight.BUFF_CAN_REMOVE {
				return true
			}
			if sameTypebuff == nil {
				sameTypebuff = buff
			} else {
				if buffConfT.BuffRule == constFight.BUFF_RULE_REPLACE {
					if sameTypebuff.GetStartTime() > buff.GetStartTime() {
						sameTypebuff = buff
					}
				} else if buffConfT.BuffRule == constFight.BUFF_RULE_REPLACE_LOW_LV {
					if sameTypebuff.GetBuffT().BuffValue > buffConf.BuffValue {
						sameTypebuff = buff
					} else if sameTypebuff.GetBuffT().BuffValue == buffConfT.BuffValue {
						if sameTypebuff.GetStartTime() > buff.GetStartTime() {
							sameTypebuff = buff
						}
					}
				}
			}
		}
		return true
	})

	if sameTypebuff == nil {
		if sameTypeBuffNum >= buffConfT.Layer {
			logger.Debug("buff manager 相同类型组已达最大层数，不能再添加了,玩家：%v,当前组层数：%v", this.owner.NickName(), sameTypeBuffNum)
			return false, nil
		}
		return true, nil
	}

	////无限buff, 同buffId 不重复添加
	//if buffConfT.Time == -1 && sameTypebuff.GetBuffT().Id == buffConfT.Id {
	//	logger.Debug("-------------存在无限相同buff", buffConfT.Id)
	//	return false, nil
	//}

	if buffConfT.BuffRule == constFight.BUFF_RULE_REPLACE {
		//同类叠加规则，后上的BUFF顶替前面的BUFF；遇到多层buff，后上buff顶替掉第一层buff
		if sameTypeBuffNum >= buffConfT.Layer {
			return true, sameTypebuff
		}
	} else if buffConfT.BuffRule == constFight.BUFF_RULE_NO_REPLACE {
		//同类叠加规则，不能顶替，不能添加新buff
		logger.Debug("存在相同类型buff:%v", sameTypebuff.GetBuffT().Id)
		return false, nil
		//} else if buffConfT.BuffRule == constFight.BUFF_RULE_NO_AFFECT {
		//
		//	if sameTypeBuffNum >= buffConfT.Layer {
		//		logger.Debug("已达最大层数,%v,最大：%v，当前：%v", buffConfT.Id, sameTypeBuffNum, buffConfT.Layer)
		//		return false, nil
		//	}
	} else {
		//同类叠加规则，高优先级顶替低优先级
		if sameTypeBuffNum >= buffConfT.Layer {
			if buffConfT.BuffValue > sameTypebuff.GetBuffT().BuffValue {
				return true, sameTypebuff
			}
			logger.Debug("存在相同类型buff:%v,优先级:%v，低于已有buff：%v", sameTypebuff.GetBuffT().Id, buffConfT.BuffValue, sameTypebuff.GetBuffT().BuffValue)
			return false, sameTypebuff
		}
	}
	return true, nil
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
func (this *BuffManager) AddNewBuff(buffId int, sourceActor base.Actor, attackTarget base.Actor, isInit bool, arg ...int) (int, error) {
	buffConfT := gamedb.GetBuffBuffCfg(buffId)
	if buffConfT == nil {
		logger.Error("buffmanager 添加buff异常，玩家：%v,buffId：%v,未找到buff配置 ", this.owner.NickName(), buffId)
		return 0, errex.NewErrorItem("textError.NoBuff :%v", buffId)
	}

	//地面buff
	isGroundBuff := false
	if len(arg) > 0 {
		isGroundBuff = this.createGroundBuff(buffConfT, attackTarget, arg[0])
	}
	if isGroundBuff {
		return 0, nil
	}

	//免疫buff
	if this.isImmune(buffConfT, this.owner) {
		return 0, nil
	}
	//瞬间buff 回血 回蓝处理
	isMonment := this.momentBuff(buffConfT, this.owner, attackTarget)
	if isMonment {
		return 0, nil
	}
	//判断是否无敌buff
	if this.IsCanWudi() && buffConfT.Debuff == 1 {
		logger.Debug("buff 无敌状态添加buff失败", this.owner.GetObjId())
		return 0, nil
	}
	//随机概率
	isAdd := common.RandByTenShousand(buffConfT.Probability)
	if !isAdd {
		return 0, nil
	}
	canAdd, delBuff := this.sameTypeBuff(buffConfT)
	if !canAdd {
		return 0, nil
	}
	//创建buff
	newBuff, err := this.createBuff(buffConfT, this.owner, sourceActor, arg...)
	if err != nil {
		return 0, err
	}
	if buffConfT.BuffType == pb.BUFFTYPE_ADD_HP_BY_SKILL_HURT_PRO {
		newBuff.(*IntervalPropBuff).setAttackTarget(attackTarget)
	}
	//删除顶替的buff
	delBuffInfos := make([]*pb.DelBuffInfo, 0)
	if delBuff != nil {
		this.BuildDelBuff(this.owner, delBuff.GetBuffIdx(), &delBuffInfos)
		this.RemoveBuffByIdx(int(delBuff.GetBuffIdx()))
	}
	//无敌buff添加 移除所有debuff
	this.delDelBuffByAddWudiBuff(buffConfT.BuffType, &delBuffInfos)

	//添加buff,计算添加效果ntf
	buffHpChangeInfos := make([]*pb.BuffHpChangeInfo, 0)
	this.buffs.Store(int(newBuff.GetBuffIdx()), newBuff)
	newBuff.OnAdd(&buffHpChangeInfos, &delBuffInfos)

	ntf := this.BuildAddBuff(newBuff)
	ntf.DelBuffInfos = delBuffInfos
	if !isInit {
		this.owner.NotifyNearby(this.owner, ntf, nil)
	}

	hpChange := 0
	for _, buffHpChange := range buffHpChangeInfos {
		hpChange += int(buffHpChange.ChangeHp)
	}

	logger.Info("buff_manager,玩家:%v 添加buff：%v-%v,bufftype:%v,来源：%v", this.owner.NickName(), buffConfT.Id, buffConfT.Id, buffConfT.BuffType, sourceActor.NickName())
	return hpChange, nil
}

func (this *BuffManager) delDelBuffByAddWudiBuff(addBuffType int, delBuffInfos *[]*pb.DelBuffInfo) {

	if addBuffType != pb.BUFFTYPE_INVINCIBLE {
		return
	}
	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		//删除过期buff
		if buff.GetBuffT().Remove != constFight.BUFF_CAN_REMOVE {
			return true
		}
		if buff.GetBuffT().Debuff == constFight.BUFF_IS_DEBUFF_TRUE {
			this.BuildDelBuff(this.owner, buff.GetBuffIdx(), delBuffInfos)
			this.RemoveBuffByIdx(int(buff.GetBuffIdx()))
		}
		return true
	})
}

/**
 *  @Description:		添加buff
 *  @param buffId		buffId
 *  @param sourceActor  来源
 *  @return int			血量变化
 *  @return error		异常
 */
func (this *BuffManager) AddBuff(buffId int, sourceActor base.Actor, isInit bool, arg ...int) (int, error) {
	return this.AddNewBuff(buffId, sourceActor, nil, isInit, arg...)
}

func (this *BuffManager) LeaveScene() {
	if len(this.groupBuffs) > 0 {
		for _, v := range this.groupBuffs {
			v.LeaveScene()
		}
	}
}

func (this *BuffManager) Run() {

	now := time.Now().UnixNano() / int64(time.Millisecond)

	//地面buff
	for k, v := range this.groupBuffs {
		groundBuffHpChangeInfos := make([]*pb.BuffHpChangeInfo, 0)
		if v.IsExpire(now) {
			v.LeaveScene()
			delete(this.groupBuffs,k)
		} else {
			v.Run(&groundBuffHpChangeInfos, nil)
		}
		if len(groundBuffHpChangeInfos) > 0 {
			ntfHp := &pb.BuffHpChangeNtf{}
			ntfHp.BuffHpChangeInfos = groundBuffHpChangeInfos
			this.owner.NotifyNearby(v, ntfHp, nil)
		}
	}

	delBuffInfos := make([]*pb.DelBuffInfo, 0)
	buffHpChangeInfos := make([]*pb.BuffHpChangeInfo, 0)

	if this.owner.GetProp().HpNow() > 0 {

		hurtDec, hurtAdd, recoveHpDec := this.GetBuffHurtEffect()
		iswudi := this.IsCanWudi()
		this.buffs.Range(func(k, v interface{}) bool {
			buff := v.(base.Buff)
			//删除过期buff
			if buff.IsExpire(now) {
				this.BuildDelBuff(this.owner, buff.GetBuffIdx(), &delBuffInfos)
				this.RemoveBuffByIdx(int(buff.GetBuffIdx()))
				return true
			}
			buff.Run(&buffHpChangeInfos, hurtDec-hurtAdd, recoveHpDec,iswudi)
			return true
		})
	}

	if len(delBuffInfos) > 0 && this.owner.GetProp().HpNow() > 0 {
		ntf := &pb.BuffChangeNtf{}
		ntf.DelBuffInfos = delBuffInfos
		if this.owner != nil {
			this.owner.NotifyNearby(this.owner, ntf, nil)
		}
	}

	if len(buffHpChangeInfos) > 0 {
		ntfHp := &pb.BuffHpChangeNtf{}
		ntfHp.BuffHpChangeInfos = buffHpChangeInfos
		if this.owner != nil {
			this.owner.NotifyNearby(this.owner, ntfHp, nil)
		}
	}
}

func (this *BuffManager) RemoveBuffByIdx(buffIdx int) {
	if v, ok := this.buffs.Load(buffIdx); ok {
		buff := v.(base.Buff)
		buff.OnRemove()
		this.buffs.Delete(buffIdx)
		logger.Debug("buff_manager,玩家:%v,移除buff:%v,id:%v,来源：%v", this.owner.NickName(), buff.GetBuffT().Id, buff.GetBuffIdx(), buff.GetSource().NickName())
	}
}

func (this *BuffManager) IsCanMove() bool {
	canMove := true
	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		if buff.GetBuffT() == nil {
			return true
		}
		selfType := buff.GetBuffT().BuffType
		//麻痹 冰冻 定身 无法移动
		if selfType == pb.BUFFTYPE_PARALYSIS || selfType == pb.BUFFTYPE_FREEZE ||
			selfType == pb.BUFFTYPE_HOLD || selfType == pb.BUFFTYPE_DIZZINESS || selfType == pb.BUFFTYPE_CONFUSION || selfType == pb.BUFFTYPE_PETRIFIED_FOR_MONSTER {
			canMove = false
			logger.Debug("移动限制 玩家：%v,buffId:%v,buffType:%v", this.owner.NickName(), buff.GetBuffT().Id, buff.GetType())
			return false
		}
		return true
	})
	return canMove
}

func (this *BuffManager) IsCanUseSkiLL(skillId int) (bool, bool) {
	isCanUse := true
	skillEffect := true
	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		if buff.GetBuffT() == nil {
			return true
		}
		selfType := buff.GetBuffT().BuffType
		if selfType == pb.BUFFTYPE_PARALYSIS || selfType == pb.BUFFTYPE_FREEZE || selfType == pb.BUFFTYPE_SILENT ||
			selfType == pb.BUFFTYPE_DIZZINESS || selfType == pb.BUFFTYPE_CONFUSION || selfType == pb.BUFFTYPE_PETRIFIED_FOR_MONSTER {
			isCanUse = false
			skillEffect = false
			return false
		} else if selfType == pb.BUFFTYPE_SKILL_INVALID {
			for _, v := range buff.GetBuffT().Effect {
				if v == skillId {
					skillEffect = false
					return false
				}
			}
		}
		return true
	})
	return isCanUse, skillEffect
}

func (this *BuffManager) IsCanWudi() bool {
	canWudi := false
	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		if buff.GetBuffT() == nil {
			return true
		}
		selfType := buff.GetBuffT().BuffType
		if selfType == pb.BUFFTYPE_INVINCIBLE {
			canWudi = true
			return false
		}
		return true
	})
	return canWudi
}

func (this *BuffManager) IsCanTreat() bool {
	canTreat := true
	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		if buff.GetBuffT() == nil {
			return true
		}
		selfType := buff.GetBuffT().BuffType
		if selfType == pb.BUFFTYPE_TREAT_BAN {
			canTreat = false
			return false
		}
		return true
	})
	return canTreat
}

func (this *BuffManager) GetTreatEffect() int {
	effect := 0
	effectLess := 0
	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		if buff.GetBuffT() == nil {
			return true
		}
		selfType := buff.GetBuffT().BuffType
		if selfType == pb.BUFFTYPE_TREAT_ADD {
			effect += buff.GetBuffT().Effect[0]
		} else if selfType == pb.BUFFTYPE_TREAT_LESS {
			effectLess -= buff.GetBuffT().Effect[0]
		} else if selfType == pb.BUFFTYPE_TREAT_ADD_BY_HP {
			a := 0
			b := 0
			ok := false
			if a, ok = buff.GetBuffT().Effect[constFight.BUFF_KEY_ZERO]; !ok {
				logger.Error("buff 配置异常：%v,缺少效果参数:%v", buff.GetBuffT().Id, buff.GetBuffT().Effect)
				return true
			}
			if b, ok = buff.GetBuffT().Effect[constFight.BUFF_KEY_ONE]; !ok {
				logger.Error("buff 配置异常：%v,缺少效果参数:%v", buff.GetBuffT().Id, buff.GetBuffT().Effect)
				return true
			}

			effect += int(float64(this.owner.GetProp().HpNow())/float64(this.owner.GetProp().Get(pb.PROPERTY_HP))*10000/float64(a)) * b
		}
		return true
	})
	logger.Debug("获取buff治疗类：战斗Id:%v，玩家：%v,增益效果：%v,减益效果:%v", this.owner.GetFight().GetId(), this.owner.NickName(), effect, effectLess)
	return effect + effectLess
}

/**
 *  @Description: 获取指定buff类型的效果
 *  @param buffType
 *  @return int
 */
func (this *BuffManager) GetBuffEffectByBuffType(buffType int) int {

	effect := 0
	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		if buff.GetBuffT() == nil {
			return true
		}
		selfType := buff.GetBuffT().BuffType
		if selfType == buffType {
			effect += buff.GetBuffT().Effect[constFight.BUFF_KEY_ZERO]
		}
		return true
	})
	return effect
}

func (this *BuffManager) IncrBuffIdx() int32 {
	if this.idx >= MaxBuffIdx {
		this.idx = 0
	}
	this.idx++
	return this.idx
}

func (this *BuffManager) BuildAddBuff(buff base.Buff) *pb.BuffChangeNtf {
	sourceId := 0
	sourceUserId := 0
	if buff.GetSource() != nil {
		sourceId = buff.GetSource().GetObjId()
		sourceUserId = buff.GetSource().GetUserId()
	}
	return &pb.BuffChangeNtf{
		Buff: &pb.BuffInfo{
			SourceObjId:  int32(sourceId),
			SourceUserId: int32(sourceUserId),
			OwnerObjId:   int32(buff.GetOwenr().GetObjId()),
			OwnerUserId:  int32(buff.GetOwenr().GetUserId()),
			Idx:          buff.GetBuffIdx(),
			BuffId:       int32(buff.GetBuffT().Id),
			TotalTime:    buff.GetEndTime()},
	}
}

func (this *BuffManager) BuildDelBuff(actor base.Actor, idx int32, delBuffInfos *[]*pb.DelBuffInfo) {
	newDelBuff := &pb.DelBuffInfo{OwnerObjId: int32(actor.GetObjId()), Idx: idx}
	*delBuffInfos = append(*delBuffInfos, newDelBuff)
}

func (this *BuffManager) ClearALlBuff() {
	var sm sync.Map
	this.buffs = sm
	this.idx = 0
}

func (this *BuffManager) IsSpecialPriority(buffType int) bool {
	if buffType == pb.BUFFTYPE_INVINCIBLE {
		return true
	}
	return false
}

func (this *BuffManager) isImmune(buffConfT *gamedb.BuffBuffCfg, actor base.Actor) bool {

	if actor.GetType() == pb.SCENEOBJTYPE_MONSTER {

		if monster, ok := actor.(base.ActorMonster); ok {
			monsterT := monster.GetMonsterT()
			if len(monsterT.ImmuneBuff) > 0 {
				if monsterT.ImmuneBuff[0] == -1 {
					//免疫所有buff
					return true
				}
				for _, v := range monsterT.ImmuneBuff {
					if buffConfT.BuffType == v {
						return true
					}
				}
			}

		}
	} else {
		//64:对怪攻击系数中毒	65:对怪石化 仅仅对怪物生效
		if buffConfT.BuffType == pb.BUFFTYPE_POISONING_BY_ATK_FOR_MONSTER || buffConfT.BuffType == pb.BUFFTYPE_PETRIFIED_FOR_MONSTER {
			return true
		}
	}
	isImmune := false
	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		if buff.GetType() == pb.BUFFTYPE_IMMUNE {
			if len(buff.GetBuffT().Effect) > 0 {
				for _, v := range buff.GetBuffT().Effect {
					if v == buffConfT.BuffType {
						isImmune = true
						return false
					}
				}
			} else {
				if buffConfT.Debuff == constFight.BUFF_IS_DEBUFF_TRUE {
					isImmune = true
					return false
				}
			}
		}
		return true
	})

	return isImmune
}

func (this *BuffManager) GetAllBuffsPb() []*pb.BuffInfo {
	buffs := make([]*pb.BuffInfo, 0)
	this.buffs.Range(func(k, v interface{}) bool {
		idx := k.(int)
		buff := v.(base.Buff)
		sourceObjId := 0
		sourceUserId := 0
		if buff.GetSource() != nil {
			sourceObjId = buff.GetSource().GetObjId()
			sourceUserId = buff.GetSource().GetUserId()
		}
		buffs = append(buffs, &pb.BuffInfo{
			SourceObjId:  int32(sourceObjId),
			SourceUserId: int32(sourceUserId),
			OwnerObjId:   int32(buff.GetOwenr().GetObjId()),
			OwnerUserId:  int32(buff.GetOwenr().GetUserId()),
			Idx:          int32(idx),
			BuffId:       int32(buff.GetBuffT().Id),
			TotalTime:    buff.GetEndTime(),
		})
		return true
	})
	return buffs
}

/**
 *  @Description: 技能伤害加成
 *  @return int
 */
func (this *BuffManager) GetBuffSkillHurtAdd() int {
	skillHurtAdd := 0
	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		if buff.GetBuffT().BuffType == pb.BUFFTYPE_FLAME {
			skillHurtAdd += buff.GetBuffT().Effect[constFight.BUFF_KEY_ZERO]
			logger.Debug("获取buff类型:%v，buffId:%v,效果：%v,获取当前类型总效果值：%v", pb.BUFFTYPE_FLAME, buff.GetBuffT().Id, buff.GetBuffT().Effect, skillHurtAdd)
		}
		return true
	})
	return skillHurtAdd
}

/**
 *  @Description: 技能伤害加成,指定技能Id
 *  @return int
 */
func (this *BuffManager) GetBuffSkillHurtAddBySkillId(skillId int) int {
	skillHurtAdd := 0
	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		if buff.GetBuffT().BuffType == pb.BUFFTYPE_SKILL_HURT_ADD {
			if buff.GetBuffT().Effect[0] == skillId {
				skillHurtAdd += buff.GetBuffT().Effect[1]
				logger.Debug("获取buff类型:%v，buffId:%v,效果：%v", pb.BUFFTYPE_SKILL_HURT_ADD, buff.GetBuffT().Id, buff.GetBuffT().Effect)
			}
		}
		return true
	})
	return skillHurtAdd
}

func (this *BuffManager) BuffSkillInCd() bool {
	skillInCd := true
	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		if buff.GetBuffT().BuffType == pb.BUFFTYPE_SKILL_CD {
			if common.RandByTenShousand(buff.GetBuffT().Effect[0]) {
				skillInCd = false
				logger.Debug("获取buff类型:%v，buffId:%v,技能不进入cd，buff效果:%v", pb.BUFFTYPE_SKILL_CD, buff.GetBuffT().Id, buff.GetBuffT().Effect)
			}
			return false
		}
		return true
	})
	return skillInCd
}

/**
 *  @Description: 获取降低最终伤害比例
 *  @return int
 */
func (this *BuffManager) GetBuffFinalHurtDec() int {
	hurtDec, hurtAdd, _ := this.GetBuffHurtEffect()
	return hurtDec - hurtAdd
}

func (this *BuffManager) GetBuffHurtEffect() (int, int, int) {

	hurtDec := 0
	hurtAdd := 0
	recoverDec := 0
	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		if buff.GetBuffT().BuffType == pb.BUFFTYPE_DEC_HURT {
			hurtDec += buff.GetBuffT().Effect[constFight.BUFF_KEY_ZERO]
			logger.Debug("获取buff类型:%v，buffId:%v,效果：%v，总伤害修正：%v", pb.BUFFTYPE_DEC_HURT, buff.GetBuffT().Id, buff.GetBuffT().Effect, hurtDec)
		} else if buff.GetType() == pb.BUFFTYPE_ADD_HURT {
			hurtAdd += buff.GetBuffT().Effect[constFight.BUFF_KEY_ZERO]
			logger.Debug("获取buff类型:%v，buffId:%v,效果：%v，总伤害修正：%v", pb.BUFFTYPE_ADD_HURT, buff.GetBuffT().Id, buff.GetBuffT().Effect, hurtAdd)
		} else if buff.GetType() == pb.BUFFTYPE_DES_RECOVERY_HP {
			recoverDec += buff.GetBuffT().Effect[constFight.BUFF_KEY_ZERO]
			logger.Debug("获取buff类型:%v，buffId:%v,效果：%v，总回血修正：%v", pb.BUFFTYPE_DES_RECOVERY_HP, buff.GetBuffT().Id, buff.GetBuffT().Effect, recoverDec)
		}
		return true
	})
	return hurtDec, hurtAdd, recoverDec
}

func (this *BuffManager) GetBuffFinalHurtAdd() int {
	hurtAdd := 0
	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		if buff.GetBuffT().BuffType == pb.BUFFTYPE_ATTACK_ADD_BY_HP {

			a := 0
			b := 0
			ok := false
			if a, ok = buff.GetBuffT().Effect[constFight.BUFF_KEY_ZERO]; !ok {
				logger.Error("buff 配置异常：%v,缺少效果参数:%v", buff.GetBuffT().Id, buff.GetBuffT().Effect)
				return true
			}
			if b, ok = buff.GetBuffT().Effect[constFight.BUFF_KEY_ONE]; !ok {
				logger.Error("buff 配置异常：%v,缺少效果参数:%v", buff.GetBuffT().Id, buff.GetBuffT().Effect)
				return true
			}

			lowHp := this.owner.GetProp().Get(pb.PROPERTY_HP) - this.owner.GetProp().HpNow()
			maxHp := this.owner.GetProp().Get(pb.PROPERTY_HP)
			hurtAdd += int(float64(lowHp)/float64(maxHp)*10000/float64(a)) * b

			logger.Debug("获取buff类型:%v，buffId:%v,效果：%v，总伤害增加：%v", pb.BUFFTYPE_ATTACK_ADD_BY_HP, buff.GetBuffT().Id, buff.GetBuffT().Effect, lowHp, maxHp, hurtAdd)
		}
		return true
	})
	return hurtAdd
}

/**
 *  @Description: 获取吸收伤害后最终伤害
 *  @return int
 */
func (this *BuffManager) GetBuffFinalHurtDecFix(hurt int) int {

	allChangeMp := 0
	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		if buff.GetBuffT().BuffType == pb.BUFFTYPE_DEC_HURT_FIXED {
			hurt = buff.(*StatusBuff).decHurtFix(hurt)
			if hurt <= 0 {
				return false
			}
		} else if buff.GetBuffT().BuffType == pb.BUFFTYPE_MP_TO_HP {
			nowMp := this.owner.GetProp().MpNow()
			if nowMp > 0 {
				totalHurt := hurt
				hurt = common.MaxIntGet(totalHurt-int(float64(nowMp*buff.GetBuffT().Effect[0])/base.BASE_RATE), 0)
				decMp := common.MinIntGet(int(float64(totalHurt)/base.BASE_RATE), nowMp)
				this.owner.GetProp().SetMpNow(nowMp - decMp)
				allChangeMp += decMp
			}
		}
		return true
	})
	if allChangeMp > 0 {
		//推送血量变化
		MPChangeNtf := &pb.SceneObjMpNtf{
			ObjId:    int32(this.owner.GetObjId()),
			Mp:       int64(this.owner.GetProp().MpNow()),
			ChangeMp: int64(-allChangeMp),
			TotalMp:  int64(this.owner.GetProp().Get(pb.PROPERTY_MP)),
		}
		if u, ok := this.owner.(base.ActorUser); ok {
			u.SendMessage(MPChangeNtf)
		}
	}

	return hurt
}

//复制增益buff
func (this *BuffManager) CloneGainBuff(source base.Actor) {

	now := common.GetNowMillisecond()
	allBuff := source.GetAllBuffsPb()
	for _, v := range allBuff {

		buffT := gamedb.GetBuffBuffCfg(int(v.BuffId))
		if buffT.Debuff != 1 && now-v.TotalTime > 1000 {
			this.AddBuff(int(v.BuffId), this.owner, true)
		}
	}
}

//删除负面buff
func (this *BuffManager) DelDeBuff() {
	delBuffNtf := &pb.BuffDelNtf{
		DelBuffInfos: make([]*pb.DelBuffInfo, 0),
	}
	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		if buff.GetBuffT().Debuff == constFight.BUFF_IS_DEBUFF_TRUE || buff.GetBuffT().Resurrection == 1 {
			this.BuildDelBuff(this.owner, buff.GetBuffIdx(), &delBuffNtf.DelBuffInfos)
			this.RemoveBuffByIdx(int(buff.GetBuffIdx()))
		}
		return true
	})
	//推送客户端buff移除
	if len(delBuffNtf.DelBuffInfos) > 0 {
		this.owner.NotifyNearby(this.owner, delBuffNtf, nil)
	}
}

func (this *BuffManager) DelGoodBuff(layer int) {
	delBuffNtf := &pb.BuffDelNtf{
		DelBuffInfos: make([]*pb.DelBuffInfo, 0),
	}
	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		if buff.GetBuffT().Debuff == 2 {
			if buff.GetBuffT().Remove != constFight.BUFF_CAN_REMOVE {
				return true
			}
			this.BuildDelBuff(this.owner, buff.GetBuffIdx(), &delBuffNtf.DelBuffInfos)
			this.RemoveBuffByIdx(int(buff.GetBuffIdx()))
			if layer > 0 {
				layer -= 1
				if layer <= 0 {
					return false
				}
			}
		}
		return true
	})
	//推送客户端buff移除
	if len(delBuffNtf.DelBuffInfos) > 0 {
		this.owner.NotifyNearby(this.owner, delBuffNtf, nil)
	}
}

func (this *BuffManager) AspdBuffAddValue(isClear bool) {

	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		if buff.GetBuffT().BuffType == pb.BUFFTYPE_ASPD_ADD_PRO {
			aspdBuff := buff.(*OnePropsAspdBuff)
			aspdBuff.AddAttackSpeed(isClear)
			return false
		}
		return true
	})
}

func (this *BuffManager) GetBuffFireSkillHurtAdd() int {
	hurtAdd := 0
	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		if buff.GetBuffT().BuffType == pb.BUFFTYPE_ADD_HURT_BY_FIRE {
			hurtAdd += buff.GetBuffT().Effect[0]
			logger.Debug("获取buff类型:%v，buffId:%v,效果：%v，总伤害增加：%v", pb.BUFFTYPE_ADD_HURT_BY_FIRE, buff.GetBuffT().Id, buff.GetBuffT().Effect, hurtAdd)
		}
		return true
	})
	return hurtAdd
}

func (this *BuffManager) BuffHasType(buffType int, buffIds []int) (bool, int) {
	has := false
	hasBuffId := 0
	buffIdsMap := make(map[int]bool)
	if buffIds != nil && len(buffIds) > 0 {
		for _, v := range buffIds {
			buffIdsMap[v] = true
		}
	}
	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		if buffType == -1 || buff.GetBuffT().BuffType == buffType {
			if len(buffIdsMap) == 0 || (len(buffIdsMap) > 0 && buffIdsMap[buff.GetBuffT().Id]) {
				has = true
				hasBuffId = buff.GetBuffT().Id
				return false
			}
		}
		return true
	})
	return has, hasBuffId
}

func (this *BuffManager) BuffRemoveByBuffId(buffId int) {
	delBuffNtf := &pb.BuffDelNtf{
		DelBuffInfos: make([]*pb.DelBuffInfo, 0),
	}
	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		if buff.GetBuffT().Id == buffId {
			this.BuildDelBuff(this.owner, buff.GetBuffIdx(), &delBuffNtf.DelBuffInfos)
			this.RemoveBuffByIdx(int(buff.GetBuffIdx()))
			return false
		}
		return true
	})
	//推送客户端buff移除
	if len(delBuffNtf.DelBuffInfos) > 0 {
		this.owner.NotifyNearby(this.owner, delBuffNtf, nil)
	}
}

func (this *BuffManager) BuffRemove(buffLayer int, buffType []int) {

	buffTypeMap := make(map[int]bool)
	for _, v := range buffType {
		buffTypeMap[v] = true
	}
	delBuffNtf := &pb.BuffDelNtf{
		DelBuffInfos: make([]*pb.DelBuffInfo, 0),
	}
	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		//攻速buff特殊处理
		if buff.GetBuffT().BuffType == pb.BUFFTYPE_ASPD_ADD_PRO && buffTypeMap[buff.GetBuffT().BuffType] {
			aspdBuff := buff.(*OnePropsAspdBuff)
			aspdBuff.LessAttackSpeed(buffLayer)
			return false
		}
		//移除指定类型buff
		if buff.GetBuffT().Remove == constFight.BUFF_CAN_REMOVE && buffTypeMap[buff.GetBuffT().BuffType] {
			this.BuildDelBuff(this.owner, buff.GetBuffIdx(), &delBuffNtf.DelBuffInfos)
			this.RemoveBuffByIdx(int(buff.GetBuffIdx()))
			if buffLayer > 0 {
				buffLayer -= 1
				if buffLayer <= 0 {
					return false
				}
			}
		}
		return true
	})
	//推送客户端buff移除
	if len(delBuffNtf.DelBuffInfos) > 0 {
		//推送客户端 攻速变化
		this.owner.NotifyNearby(this.owner, delBuffNtf, nil)
	}
}

func (this *BuffManager) BuffFatalRecoveHp() {

	changeHp := 0
	max := this.owner.GetProp().Get(pb.PROPERTY_HP)

	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		if buff.GetBuffT().BuffType == pb.BUFFTYPE_BE_FATAL_REOVER_HP {
			changeHp = int(float64(max) * (float64(buff.GetBuffT().Effect[0]) / base.BASE_RATE))
		}
		return true
	})
	if changeHp > 0 {
		realChangeHp, _ := this.owner.ChangeHp(changeHp)
		buffHpChangeInfos := &pb.BuffHpChangeInfo{}
		buffHpChangeInfos.Death = 0
		buffHpChangeInfos.ChangeHp = int64(realChangeHp)
		buffHpChangeInfos.Idx = 0
		buffHpChangeInfos.OwnerObjId = int32(this.owner.GetObjId())
		buffHpChangeInfos.TotalHp = int64(this.owner.GetProp().HpNow())

		ntfHp := &pb.BuffHpChangeNtf{}
		ntfHp.BuffHpChangeInfos = []*pb.BuffHpChangeInfo{buffHpChangeInfos}
		if this.owner != nil {
			this.owner.NotifyNearby(this.owner, ntfHp, nil)
		}
	}
}

func (this *BuffManager) ReflexHurt() bool {
	isfeflex := false
	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		if buff.GetBuffT().BuffType == pb.BUFFTYPE_BACK_HURT {
			isfeflex = common.RandByTenShousand(buff.GetBuffT().Effect[constFight.BUFF_KEY_ZERO])
			return false
		}
		return true
	})
	return isfeflex
}

func (this *BuffManager) ClearFitBuff() {
	delBuffInfos := make([]*pb.DelBuffInfo, 0)
	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		//删除过期buff
		if buff.GetBuffT().FitPurge == 1 {
			this.BuildDelBuff(this.owner, buff.GetBuffIdx(), &delBuffInfos)
			this.RemoveBuffByIdx(int(buff.GetBuffIdx()))
			return true
		}
		return true
	})
	//TODO 暂时先不通知客户端，合体进入的时候，客户端可能会先移除
	//if len(delBuffInfos) > 0 && this.owner.GetProp().HpNow() > 0 {
	//	ntf := &pb.BuffChangeNtf{}
	//	ntf.DelBuffInfos = delBuffInfos
	//	if this.owner != nil {
	//		this.owner.NotifyNearby(this.owner, ntf, nil)
	//	}
	//}
}

func (this *BuffManager) ReliveBuffCheck() bool {
	isRelive := false
	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		//删除过期buff
		if buff.GetBuffT().BuffType == pb.BUFFTYPE_RELIVE {
			isRelive = buff.(*StatusReliveBuff).Relive()
			return false
		}
		return true
	})
	return isRelive
}

func (this *BuffManager) CheckTriggerPropMust(attacker base.Actor, propId int) bool {

	trigger := false
	hasBuff := make(map[int]bool)
	checkHasBuff := make(map[int]bool)
	this.buffs.Range(func(k, v interface{}) bool {
		buff := v.(base.Buff)
		if buff.GetSource() != nil && buff.GetSource().GetObjId() == attacker.GetObjId() {
			hasBuff[buff.GetBuffT().BuffType] = true
			if buff.GetType() == pb.BUFFTYPE_TRIGGER_PROP_MUST {
				if buff.GetBuffT().Effect[constFight.BUFF_KEY_ONE] == propId {
					if buff.GetBuffT().Effect[constFight.BUFF_KEY_ZERO] == -1 {
						trigger = true
						logger.Debug("buffmanager 玩家：%v,属性触发检查,拥有buff类型：%v", this.owner.NickName(), buff.GetBuffT().Effect)
						return false
					}
					checkHasBuff[buff.GetBuffT().Effect[constFight.BUFF_KEY_ZERO]] = true
				}
			}
		}
		return true
	})
	if trigger {
		return trigger
	}
	for k, _ := range checkHasBuff {
		if hasBuff[k] {
			logger.Debug("buffmanager 玩家：%v,属性触发检查,拥有buff类型：%v", this.owner.NickName(), k)
			return true
		}
	}
	return false
}
