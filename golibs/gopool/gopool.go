package gopool

import "sync"

type Item interface {
	Execute()
}

type Worker func()

type Pool struct {
	items  chan Item
	once   sync.Once
	finish chan struct{}
}

func (this Worker) Execute() {
	this()
}

func NewPool(count int) *Pool {
	pool := &Pool{
		items:  make(chan Item, count*10),
		finish: make(chan struct{}),
	}
	for i := 0; i < count; i++ {
		go pool.run()
	}
	return pool
}

func (this *Pool) Add(item Item) {
	this.items <- item
}

func (this *Pool) run() {
	for {
		select {
		case item, ok := <-this.items:
			if !ok {
				close(this.finish)
				return
			}
			item.Execute()
		}
	}
}

func (this *Pool) Stop() {
	this.once.Do(func() {
		close(this.items)
	})
	<-this.finish
}
