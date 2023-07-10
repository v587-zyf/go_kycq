package buff

import (
	"cqserver/fightserver/internal/base"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/prop"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
)

type OnePropsBuff struct {
	*DefaultBuff
	effect map[int]int //propId->effectValue
}

func NewOnePropsBuff(act base.Actor, sourceActor base.Actor, buffT *gamedb.BuffBuffCfg, idx int32) *OnePropsBuff {
	onePropsBuff := &OnePropsBuff{
		effect: make(map[int]int),
	}
	onePropsBuff.DefaultBuff = NewDefaultBuff(buffT, sourceActor, act, onePropsBuff, idx)
	return onePropsBuff
}

//添加buff效果(改变属性的)
func (this *OnePropsBuff) OnAdd(buffHpChangeInfos *[]*pb.BuffHpChangeInfo, delBuffInfos *[]*pb.DelBuffInfo) {

	if this.buffT.BuffType == pb.BUFFTYPE_ADD_PROP_BY_ATK || this.buffT.BuffType == pb.BUFFTYPE_LESS_PROP_BY_ATK {
		this.changeAtkByAttack()
		return
	}

	aspdChange := false
	for k, v := range this.buffT.Effect {
		baseValue := this.owner.GetProp().GetActorBaseProp(k)

		changeValue := 0
		if this.buffT.BuffType == pb.BUFFTYPE_ADD_PROP_BY_PRO {
			changeValue = int(float64(baseValue) * (float64(v) / 10000))
		} else if this.buffT.BuffType == pb.BUFFTYPE_DEC_PROP_BY_PRO {
			changeValue = -1 * int(float64(baseValue)*(float64(v)/10000))
		} else if this.buffT.BuffType == pb.BUFFTYPE_ADD_PROP_BY_FIXED {
			changeValue = v
		} else if this.buffT.BuffType == pb.BUFFTYPE_DEC_PROP_BY_FIXED {
			changeValue = -1 * v
		} else if this.buffT.BuffType == pb.BUFFTYPE_CUT_SKILL {
			changeValue = v
		}
		if k == pb.PROPERTY_ATT_SPEED {
			aspdChange = true
		}

		logger.Debug("buff 属性增加,玩家：%v，buffId:%v，增加/减少属性：%v 改变值：%v,基础值：%v", this.owner.NickName(), this.buffT.Id, k, changeValue, baseValue)

		if k == pb.PROPERTY_ATT_ALL {
			this.changeAttack(changeValue)
		} else if k == pb.PROPERTY_DEF_ALL {
			this.effect[pb.PROPERTY_DEF_MAX] = changeValue
			this.effect[pb.PROPERTY_DEF_MIN] = changeValue
			this.owner.GetProp().BuffPropChange(pb.PROPERTY_DEF_MAX, changeValue, false)
			this.owner.GetProp().BuffPropChange(pb.PROPERTY_DEF_MIN, changeValue, false)

		} else if k == pb.PROPERTY_ADF_ALL {
			this.effect[pb.PROPERTY_ADF_MAX] = changeValue
			this.effect[pb.PROPERTY_ADF_MIN] = changeValue
			this.owner.GetProp().BuffPropChange(pb.PROPERTY_ADF_MAX, changeValue, false)
			this.owner.GetProp().BuffPropChange(pb.PROPERTY_ADF_MIN, changeValue, false)

		} else {
			this.effect[k] = changeValue
			//增加buff效果
			this.owner.GetProp().BuffPropChange(k, changeValue, false)
		}

		if k == pb.PROPERTY_HP {
			if this.GetOwenr().GetProp().HpNow() > this.GetOwenr().GetProp().Get(pb.PROPERTY_HP) {
				this.GetOwenr().GetProp().SetHpNow(this.GetOwenr().GetProp().Get(pb.PROPERTY_HP))
			}
			//推送血量变化
			HPChangeNtf := &pb.SceneObjHpNtf{
				ObjId:    int32(this.GetOwenr().GetObjId()),
				Hp:       int64(this.GetOwenr().GetProp().HpNow()),
				ChangeHp: int64(0),
				TotalHp:  int64(this.GetOwenr().GetProp().Get(pb.PROPERTY_HP)),
			}
			this.GetOwenr().NotifyNearby(this.GetOwenr(), HPChangeNtf, nil)
		} else if k == pb.PROPERTY_MP {
		}
	}

	if aspdChange {
		ntf := &pb.BuffPropChangeNtf{
			ObjId:  int32(this.owner.GetObjId()),
			PropId: pb.PROPERTY_ATT_SPEED,
			Total:  int64(this.owner.GetProp().Get(pb.PROPERTY_ATT_SPEED)),
		}
		//推送客户端 攻速变化
		if u, ok := this.owner.(base.ActorUser); ok {
			u.SendMessage(ntf)
		}
		logger.Debug("buff 攻速改变，玩家：%v,当前攻速：%v", this.owner.NickName(), this.owner.GetProp().Get(pb.PROPERTY_ATT_SPEED))
	}
}

func (this *OnePropsBuff) changeAtkByAttack() {

	if this.owner != nil {
		propMaxId, propMinId := prop.GetAtkPropIdByJob(this.owner.Job())
		baseValue := (this.owner.GetProp().GetActorBaseProp(propMaxId) + this.owner.GetProp().GetActorBaseProp(propMinId)) / 2
		changeValue := -1 * int(float64(baseValue)*(float64(this.buffT.Effect[constFight.BUFF_KEY_ZERO])/10000))
		this.changeAttack(changeValue)
	}
}

/**
 *  @Description: 更新攻击值
 *  @param changeValue
 */
func (this *OnePropsBuff) changeAttack(changeValue int) {
	propMaxId, propMinId := prop.GetAtkPropIdByJob(this.owner.Job())
	this.effect[propMaxId] = changeValue
	this.effect[propMinId] = changeValue
	this.owner.GetProp().BuffPropChange(propMaxId, changeValue, false)
	this.owner.GetProp().BuffPropChange(propMinId, changeValue, false)
}

//取消buff效果(改变属性的)
func (this *OnePropsBuff) OnRemove() {
	actorProp := this.owner.GetProp()
	aspdChange := false
	for k, v := range this.effect {
		actorProp.BuffPropChange(k, -v, true)
		if k == pb.PROPERTY_ATT_SPEED {
			aspdChange = true
		}
	}
	if aspdChange {
		ntf := &pb.BuffPropChangeNtf{
			ObjId:  int32(this.owner.GetObjId()),
			PropId: pb.PROPERTY_ATT_SPEED,
			Total:  int64(this.owner.GetProp().Get(pb.PROPERTY_ATT_SPEED)),
		}
		//推送客户端 攻速变化
		if u, ok := this.owner.(base.ActorUser); ok {
			u.SendMessage(ntf)
		}
		logger.Debug("buff 攻速改变，玩家：%v,当前攻速：%v", this.owner.NickName(), this.owner.GetProp().Get(pb.PROPERTY_ATT_SPEED))
	}
	logger.Debug("buff 属性增加,删除 玩家：%v，buffId:%v,增加属性：%v", this.owner.NickName(), this.buffT.Id, this.buffT.Effect)
}
