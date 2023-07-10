package util

import "testing"

func TestGenerate(t *testing.T) {

	idAllcocator := NewUint32IdAllocator(1)
	var i uint32
	maxUint32 := ^uint32(0)
	for {
		i++
		id := idAllcocator.Get()
		if id < idAllcocator.initId || id > idAllcocator.maxId {
			t.Fatalf("id should bigger than initId,id:%d, initId:%d", id, idAllcocator.initId)
		}
		if i == maxUint32 {
			break
		}
	}

}
