package util

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"sync/atomic"
)

type Uint32IdAllocator struct {
	id     uint32
	initId uint32
	maxId  uint32
}

func NewUint32IdAllocator(seq int) *Uint32IdAllocator {
	initId := uint32(seq) << 28
	return &Uint32IdAllocator{id: initId, initId: initId, maxId: initId | 0xfffffff}
}

func (this *Uint32IdAllocator) Get() uint32 {
	id := atomic.AddUint32(&this.id, 1)
	if id == this.maxId {
		atomic.StoreUint32(&this.id, (this.initId + 1))
	}
	return id
}

func GenerateSessionId() (string, error) {
	k := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, k); err != nil {
		return "", nil
	}
	return hex.EncodeToString(k), nil
}
