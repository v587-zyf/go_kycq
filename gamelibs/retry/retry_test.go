package retry

import (
	"fmt"
	// "sync"
	"testing"
)

type testInfo struct {
	id    int
	count int
}

func (this *testInfo) GetId() int {
	return this.id
}

func (this *testInfo) Execut() bool {
	if this.id == 99 {
		fmt.Printf("id:%d count:%d \n", this.id, this.count)
	}

	this.count++
	return false
}

func Test_Retry(t *testing.T) {
	a := NewRetryManager()
	a.Start()
	fmt.Println("begin")
	// var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		b := &testInfo{id: i, count: 0}
		a.AddService(b, []int64{1, 3, 5})
	}
	c := make(chan struct{})
	select {
	case <-c:
		return
	}
}
