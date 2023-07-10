package base

import (
	"sync/atomic"
)

// MinFightId以内的为保留id，用于maincity
// fightId的个位数为serverSeq，同时支持的FightServer数为9
const (
	MinFightId uint32 = 1000
)

type IdAllocator struct {
	serverSeq uint32
	idSeq     uint32
}

func NewIdAllocator(serverSeq int) *IdAllocator {
	return &IdAllocator{
		serverSeq: uint32(serverSeq),
		idSeq:     MinFightId + uint32(serverSeq),
	}
}

func (this *IdAllocator) GenerateFightId() uint32 {
	id := atomic.AddUint32(&this.idSeq, 10)
	if id < MinFightId {
		id = atomic.AddUint32(&this.idSeq, MinFightId)
	}
	return id
}

func (this *IdAllocator) GenerateMainCityId(kingdom int, lineNo int) uint32 {
	if lineNo < 1 || lineNo >= 100 {
		panic("maincity lineno must be between 1 and 99")
	}
	return uint32(kingdom*100 + lineNo*10 + int(this.serverSeq))
}
