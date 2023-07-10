package buff

import (
	"cqserver/fightserver/internal/base"
	"cqserver/gamelibs/gamedb"
	"cqserver/golibs/logger"
	_ "cqserver/golibs/logger"
	"cqserver/protobuf/pb"
)

type StatusBuff struct {
	*DefaultBuff
	decHurt int //伤害减少
}

func NewStatusBuff(act base.Actor, sourceActor base.Actor, buffT *gamedb.BuffBuffCfg, idx int32) *StatusBuff {
	statusBuff := &StatusBuff{}
	statusBuff.DefaultBuff = NewDefaultBuff(buffT, sourceActor, act, statusBuff, idx)
	return statusBuff
}

func (this *StatusBuff) OnAdd(buffHpChangeInfos *[]*pb.BuffHpChangeInfo, delBuffInfos *[]*pb.DelBuffInfo) {

	this.decHurtAdd()
}

func (this *StatusBuff) decHurtAdd() {

	if this.buffT.BuffType == pb.BUFFTYPE_DEC_HURT_FIXED {

		if this.buffT.Effect[0] >= 0 {
			this.decHurt += int(float64(this.buffT.Effect[0]) / 10000 * float64(this.owner.GetProp().Get(pb.PROPERTY_HP)))
			this.decHurt += this.buffT.Effect[1]
		}
	}
}

func (this *StatusBuff) IsExpire(now int64) bool {
	if this.endTime < now {
		return true
	}

	if this.buffT.BuffType == pb.BUFFTYPE_DEC_HURT_FIXED {
		if this.buffT.Effect[0] >= 0 && this.decHurt <= 0 {
			return true
		}
	}
	return false
}

func (this *StatusBuff) decHurtFix(hurt int) int {

	oldDecHurt := this.decHurt
	newHurt := 0
	if this.buffT.Effect[0] >= 0 {
		if hurt > this.decHurt {

			newHurt = hurt - this.decHurt
			this.decHurt = 0

		} else if hurt < this.decHurt {

			this.decHurt -= hurt
			newHurt = 0

		} else {
			this.decHurt = 0
			newHurt = 0
		}
	}
	logger.Debug("获取伤害类型：%v，伤害：%v,可抵挡伤害:%v，抵挡伤害后最终伤害：%v，剩余可抵挡伤害：%v", pb.BUFFTYPE_DEC_HURT_FIXED, hurt, oldDecHurt, newHurt, this.decHurt)
	return newHurt
}
