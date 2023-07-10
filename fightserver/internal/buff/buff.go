package buff

import (
	"cqserver/fightserver/internal/base"
	"cqserver/gamelibs/gamedb"
	"cqserver/golibs/common"
	"cqserver/protobuf/pb"
)

type DefaultBuff struct {
	buffT   *gamedb.BuffBuffCfg
	context base.Buff

	source base.Actor //来源
	owner  base.Actor //buff持有者

	startTime int64 //buff生效时间毫秒
	endTime   int64 //结束时间毫秒
	idx       int32 //buff序号
}

func NewDefaultBuff(buffT *gamedb.BuffBuffCfg, source, owner base.Actor, context base.Buff, idx int32) *DefaultBuff {
	now := common.GetNowMillisecond()
	endTime := now + int64(buffT.Time)
	if buffT.Time < 0 {
		endTime = now + 30*86400000
	}

	return &DefaultBuff{
		buffT:     buffT,
		context:   context,
		source:    source,
		owner:     owner,
		startTime: now,
		endTime:   endTime,
		idx:       idx,
	}
}

func (this *DefaultBuff) GetSource() base.Actor {
	return this.source
}

func (this *DefaultBuff) GetOwenr() base.Actor {
	return this.owner
}

func (this *DefaultBuff) GetType() int {
	return this.buffT.BuffType
}

func (this *DefaultBuff) GetStartTime() int64 {
	return this.startTime
}

func (this *DefaultBuff) GetEndTime() int64 {
	return this.endTime
}

func (this *DefaultBuff) GetBuffIdx() int32 {
	return this.idx
}

func (this *DefaultBuff) GetBuffT() *gamedb.BuffBuffCfg {
	return this.buffT
}

func (this *DefaultBuff) OnAdd(buffHpChangeInfos *[]*pb.BuffHpChangeInfo, delBuffInfos *[]*pb.BuffInfo) {

}

func (this *DefaultBuff) OnRemove() {

}

func (this *DefaultBuff) IsExpire(now int64) bool {

	if this.endTime < now {
		return true
	}
	return false
}

func (this *DefaultBuff) Run(buffHpChangeInfos *[]*pb.BuffHpChangeInfo,arg ...interface{}) {

}
