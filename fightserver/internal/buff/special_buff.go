package buff

import (
	"cqserver/fightserver/internal/base"
	"cqserver/gamelibs/gamedb"
	"cqserver/protobuf/pb"
)

type SpecialBuff struct {
	*DefaultBuff
}

const (
	Negative = 1
)

func NewSpecialBuff(act base.Actor, sourceActor base.Actor, buffT *gamedb.BuffBuffCfg, idx int32) *SpecialBuff {
	specialBuff := &SpecialBuff{}
	specialBuff.DefaultBuff = NewDefaultBuff(buffT, sourceActor, act, specialBuff, idx)
	return specialBuff
}

//添加buff效果
func (this *SpecialBuff) OnAdd(buffHpChangeInfos *[]*pb.BuffHpChangeInfo, delBuffInfos *[]*pb.DelBuffInfo) {

}
