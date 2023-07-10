package buff

import (
	"cqserver/fightserver/internal/base"
	"cqserver/gamelibs/gamedb"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	_ "cqserver/golibs/logger"
	"time"
)

type StatusReliveBuff struct {
	*StatusBuff
	lastUseTime int64
}

func NewStatusReliveBuff(act base.Actor, sourceActor base.Actor, buffT *gamedb.BuffBuffCfg, idx int32) *StatusReliveBuff {
	statusBuff := &StatusReliveBuff{
		lastUseTime: 0,
	}
	statusBuff.StatusBuff = NewStatusBuff(act, sourceActor, buffT, idx)
	return statusBuff
}

func (this *StatusReliveBuff) Relive() bool {
	now := time.Now().Unix()
	if now-this.lastUseTime < int64(this.buffT.Effect[1]) {
		return false
	}
	rate := common.RandByTenShousand(this.buffT.Effect[0])
	logger.Debug("buff 时候复活检查触发，玩家：%v,buffId:%v，是否触发复活：%v", this.owner.NickName(), this.buffT.Id, rate)
	if rate {
		this.lastUseTime = now
		return true
	}
	return false
}
