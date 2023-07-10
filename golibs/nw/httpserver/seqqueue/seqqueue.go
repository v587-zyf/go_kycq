package seqqueue

import (
	"sync"
	"time"

	"cqserver/golibs/nw/internal/common"
)

// 任何发送给SeqQueue执行的消息需实现MessageInterface接口
type MessageInterface interface {
	Process()
	Wait()
}

type DefaultMessage struct {
	Done chan struct{}
}

func (this *DefaultMessage) Wait() {
	<-this.Done
}

type Option struct {
	NonactiveDuration time.Duration // 多长时间不活跃停止队列
	CallbackInterval  time.Duration // 定时回调
	CheckInterval     time.Duration // 定时器

	OnCallback func(q *SeqQueue, callbackType CallbackType)
}

type CallbackType int

const (
	CallbackTypeRegular    CallbackType = iota + 1 // 常规的定时回调
	CallbackTypeExpire                             // 过期回调
	CallbackTypeDone                               // Queue结束回调
	CallbackTypeServerDone                         // 模块退出回调
	CallbackTypeReActivite                         // session销毁前，从后台激活
)

type SeqQueue struct {
	LastActiveTime   int64
	LastLeaveTime    int64 //用户下线离开时间。等于0时代表在线，大于0时代表超过10秒无api调用，做离线处理
	LastCallbackTime int64
	userId           int
	messages         chan MessageInterface
	done             chan struct{}
	quitted          chan struct{}
	stopOnce         sync.Once
	UserData         interface{}
}

var (
	serverDone    chan struct{} // 代表整个模块关闭
	waitGroup     common.WaitGroupWrapper
	DefaultOption *Option
)

func init() {
	serverDone = make(chan struct{})
	DefaultOption = &Option{
		NonactiveDuration: 120 * time.Minute,
		CallbackInterval:  2 * time.Minute,
		CheckInterval:     100 * time.Millisecond,
		OnCallback:        func(q *SeqQueue, callbackType CallbackType) {},
	}
}

func New(userData interface{}, userId int) *SeqQueue {
	now := time.Now().Unix()
	seqQueue := &SeqQueue{
		userId:           userId,
		messages:         make(chan MessageInterface, 5),
		done:             make(chan struct{}),
		quitted:          make(chan struct{}),
		UserData:         userData,
		LastActiveTime:   now,
		LastCallbackTime: now,
	}
	waitGroup.Wrap(func() { seqQueue.messageLoop() })
	return seqQueue
}

func (this *SeqQueue) messageLoop() {
	var NONACTIVE_DURATION = int64(DefaultOption.NonactiveDuration / time.Second) //以秒为单位
	var CALLBACK_INTERVAL = int64(DefaultOption.CallbackInterval / time.Second)   //以秒为单位
	// CALLBACK_INTERVAL = 30                                                       //TODO
	var ticker = time.NewTicker(DefaultOption.CheckInterval)
	var now int64
	for {
		select {
		case <-serverDone:
			DefaultOption.OnCallback(this, CallbackTypeServerDone)
			goto exit
		case <-this.done:
			DefaultOption.OnCallback(this, CallbackTypeDone)
			goto exit
		case <-ticker.C:
			now = time.Now().Unix()
			if now-this.LastActiveTime >= NONACTIVE_DURATION+2 {
				DefaultOption.OnCallback(this, CallbackTypeExpire)
				goto exit
			} else if now-this.LastCallbackTime >= CALLBACK_INTERVAL {
				DefaultOption.OnCallback(this, CallbackTypeRegular)
				this.LastCallbackTime = now
			}
		case message := <-this.messages:
			message.Process()
		}
	}
exit:
	ticker.Stop()
	close(this.quitted)
}

func (this *SeqQueue) SendMessage(msg MessageInterface) {
	this.messages <- msg
	msg.Wait()
}

// 由外部主动调用，关闭事件处理循环
func (this *SeqQueue) Stop(wait bool) {
	this.stopOnce.Do(func() {
		close(this.done)
	})
	if wait {
		<-this.quitted
	}
}

func (this *SeqQueue) SetActiveTime(activeTime time.Time) {
	this.LastActiveTime = activeTime.Unix()
	if this.LastLeaveTime > 0 { //离线后又重新上线了（从后台切换回来）
		DefaultOption.OnCallback(this, CallbackTypeReActivite)
	}
}

// 停止所有的SeqQueue
func Stop() {
	close(serverDone)
	waitGroup.Wait()
}
