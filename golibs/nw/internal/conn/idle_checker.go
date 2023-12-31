package conn

import (
	"sync"
	"time"

	"cqserver/golibs/nw"
)

type IdleChecker struct {
	InitConnChan   chan nw.Conn
	ActiveConnChan chan nw.Conn
	CloseConnChan  chan nw.Conn
	duration       time.Duration
	conns          map[nw.Conn]bool // connected but not authed Conn
	done           chan struct{}
	wg             sync.WaitGroup
}

func NewIdleChecker(duration time.Duration) *IdleChecker {
	idleChecker := &IdleChecker{
		InitConnChan:   make(chan nw.Conn, 50),
		ActiveConnChan: make(chan nw.Conn, 50),
		CloseConnChan:  make(chan nw.Conn, 50),
		duration:       duration,
		conns:          make(map[nw.Conn]bool),
		done:           make(chan struct{}),
	}
	idleChecker.wg.Add(1)
	go func() {
		idleChecker.run()
		idleChecker.wg.Done()
	}()
	return idleChecker
}

func (this *IdleChecker) run() {
	var duration = this.duration
	var ticker = time.NewTicker(3 * time.Second)
	defer func() {
		ticker.Stop()
	}()

	for {
		select {
		case conn := <-this.InitConnChan:
			this.conns[conn] = true
		case conn := <-this.ActiveConnChan:
			if _, ok := this.conns[conn]; ok {
				delete(this.conns, conn)
			}
		case conn := <-this.CloseConnChan:
			if _, ok := this.conns[conn]; ok {
				delete(this.conns, conn)
			}
		case <-ticker.C: // check idle time
			now := time.Now()
			for conn := range this.conns {
				if now.Sub(conn.GetConnTime()) > duration {
					delete(this.conns, conn)
					conn.Close()
				}
			}
		case <-this.done:
			return
		}
	}
}

func (this *IdleChecker) Stop() {
	close(this.done)
	this.wg.Wait()
}
