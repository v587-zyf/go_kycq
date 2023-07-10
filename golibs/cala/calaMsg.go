package cala

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"cqserver/golibs/util"
)

//统计打印消息内容

type CalaMsg struct {
	codeId int64  //累加编号 用于打印时 标记每个消息
	name   string //来源
}

//统计打印消息内容
func (this *CalaMsg) PrintMsg(cmdId uint16, datalen int) (int64, time.Time) {

	now := time.Now()
	//fmt.Println("PrintMsg name", this.name, "cmdId", cmdId, "codeId", this.codeId, "datalen", datalen, now)
	return this.codeId, now
}

func (this *CalaMsg) PrintEnd(cmdId uint16, codeId int64, t time.Time) {
	//fmt.Println("PrintEnd name", this.name, "cmdId", cmdId, "codeId", codeId, time.Since(t).Seconds())
}

type CalaManager struct {
	Calas map[string]*CalaMsg
	mu    sync.RWMutex
}

var Cala *CalaManager

func init() {
	Cala = &CalaManager{
		Calas: make(map[string]*CalaMsg),
	}
}

func RegisterCalaMsg(name string) {
	Cala.mu.Lock()
	defer Cala.mu.Unlock()
	if Cala.Calas[name] != nil {
		return
	}

	one := &CalaMsg{
		codeId: 0,
		name:   name,
	}
	Cala.Calas[name] = one

	go util.SafeRun(func() {
		tick := time.NewTicker(1 * time.Minute)
		var codeId int64 = 0
		for {
			select {
			case <-tick.C:
				num := one.codeId - codeId
				//消息超过每分钟300打印
				if num > 300 {
					fmt.Println("PrintEnd name cal",time.Now(), name, "msgNum ", num) //总的消息频率
				}
				codeId += num
			}
		}
	})
}

//移除消息统计
func RemoveCalMsg( name string ){

	Cala.mu.Lock()
	defer Cala.mu.Unlock()
	if Cala.Calas[name] != nil {
		delete(Cala.Calas,name)
	}
}

func (this *CalaManager) AddMsg(name string) *CalaMsg {
	this.mu.RLock()
	defer this.mu.RUnlock()
	one := this.Calas[name]
	if one != nil {
		atomic.AddInt64(&(one.codeId), 1)
	}
	return one
}
