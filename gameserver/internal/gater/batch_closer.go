package gater

import (
	"sync"

	"cqserver/golibs/gopool"
	"cqserver/golibs/nw"
)

// 由于每个连接关闭时可能会执行sql保存操作，为了防止服务器关闭时，大量sql操作导致阻塞s（目前的sql driver存在此问题），提供goroutine池逐步关闭连接

var closerPool *gopool.Pool
var closerMu sync.Mutex

func getCloserPool() *gopool.Pool {
	closerMu.Lock()
	defer closerMu.Unlock()
	if closerPool == nil {
		closerPool = gopool.NewPool(20)
	}
	return closerPool
}

func getCloserWorker(conn nw.Conn, wg *sync.WaitGroup) gopool.Worker {
	return gopool.Worker(func() {
		conn.Close()
		conn.Wait()
		wg.Done()
	})
}
