package buff

import (
	"cqserver/fightserver/internal/base"
	"cqserver/gamelibs/gamedb"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
)

type OnePropsAspdBuff struct {
	*OnePropsBuff
	attackSpeedAdd []int //propId->effectValue
}

func NewOnePropsAspdBuff(act base.Actor, sourceActor base.Actor, buffT *gamedb.BuffBuffCfg, idx int32) *OnePropsAspdBuff {
	OnePropsAspdBuff := &OnePropsAspdBuff{
		OnePropsBuff:   NewOnePropsBuff(act, sourceActor, buffT, idx),
		attackSpeedAdd: make([]int, 0),
	}
	return OnePropsAspdBuff
}

func (this *OnePropsAspdBuff) AddAttackSpeed(isClear bool) {

	if isClear {
		for _, v := range this.attackSpeedAdd {
			this.effect[pb.BUFFTYPE_ASPD_ADD_PRO] -= v
			this.owner.GetProp().BuffPropChange(pb.BUFFTYPE_ASPD_ADD_PRO, -v, true)
		}
		this.attackSpeedAdd = make([]int, 0)
	}

	if len(this.attackSpeedAdd) >= this.buffT.Layer {
		return
	}
	this.attackSpeedAdd = append(this.attackSpeedAdd, this.buffT.Effect[pb.PROPERTY_ATT_SPEED])
	this.effect[pb.BUFFTYPE_ASPD_ADD_PRO] += this.buffT.Effect[pb.PROPERTY_ATT_SPEED]
	this.owner.GetProp().BuffPropChange(pb.BUFFTYPE_ASPD_ADD_PRO, this.buffT.Effect[pb.PROPERTY_ATT_SPEED], true)
	ntf := &pb.BuffPropChangeNtf{
		ObjId:  int32(this.owner.GetObjId()),
		PropId: pb.PROPERTY_ATT_SPEED,
		Total:  int64(this.owner.GetProp().Get(pb.PROPERTY_ATT_SPEED)),
	}
	//推送客户端 攻速变化
	if u, ok := this.owner.(base.ActorUser); ok {
		u.SendMessage(ntf)
	}
	logger.Debug("buff 攻速增加,玩家：%v，增加：%v", this.owner.NickName(), this.attackSpeedAdd)
}

func (this *OnePropsAspdBuff) LessAttackSpeed(layer int) {

	l := len(this.attackSpeedAdd)
	if l > layer {
		l = layer
	}
	for i := 0; i < l; i++ {
		this.effect[pb.BUFFTYPE_ASPD_ADD_PRO] -= this.attackSpeedAdd[i]
		this.owner.GetProp().BuffPropChange(pb.BUFFTYPE_ASPD_ADD_PRO, -this.attackSpeedAdd[i], true)
	}
	this.attackSpeedAdd = this.attackSpeedAdd[l:]
	ntf := &pb.BuffPropChangeNtf{
		ObjId:  int32(this.owner.GetObjId()),
		PropId: pb.PROPERTY_ATT_SPEED,
		Total:  int64(this.owner.GetProp().Get(pb.PROPERTY_ATT_SPEED)),
	}
	//推送客户端 攻速变化
	if u, ok := this.owner.(base.ActorUser); ok {
		u.SendMessage(ntf)
	}
	logger.Debug("buff 攻速增加,去除增加的攻速,玩家：%v，增加：%v,删掉层数：%v", this.owner.NickName(), this.attackSpeedAdd, layer)
}
