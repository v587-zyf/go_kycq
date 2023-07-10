package retry

/*用于需要定时多次重复执行的函数进行托管执行*/

import (
	"time"

	"cqserver/golibs/gopool"
)

type RetryItem interface {
	GetId() int
	Execute() bool
}

type retryItem struct {
	item       RetryItem
	startTime  int64
	updateTime int64
	timeCount  []int64
	count      int
	state      bool
}

func (this *retryItem) GetCount() int {
	return this.count
}

func (this *retryItem) IsDelete() bool {
	if this.count >= len(this.timeCount) {
		return true
	}
	return false
}

func (this *retryItem) GetTryTime() int64 {
	return this.updateTime + this.timeCount[this.count]
}

func (this *retryItem) UpdateTimeAndCount() bool {
	this.count++
	if this.count >= len(this.timeCount) {
		return true
	}
	this.updateTime = time.Now().Unix()
	return false
}

func (this *retryItem) Execute() {
	if this.item.Execute() || this.UpdateTimeAndCount() {
		this.count = len(this.timeCount)
		return
	}
	this.state = false
}

type RetryManager struct {
	itemChans      chan *retryItem
	deleteItemChan chan int
	done           chan struct{}
	itemMap        map[int]*retryItem
	workPool       *gopool.Pool
	workNum        int
}

func NewRetryManager(workNum int) *RetryManager {
	return &RetryManager{
		itemChans: make(chan *retryItem, 100),
		itemMap:   make(map[int]*retryItem),
		done:      make(chan struct{}),
		workNum:   workNum,
	}
}

func newRetryItem(item RetryItem, timeCount []int64) *retryItem {
	return &retryItem{
		item:       item,
		startTime:  time.Now().Unix(),
		updateTime: time.Now().Unix(),
		timeCount:  timeCount,
		count:      0,
		state:      false,
	}
}

var retryManager *RetryManager

func init() {
	retryManager = NewRetryManager(10)
	retryManager.Start()
}

//timeCount参数为执行次数的数组，单位为秒
func Add(item RetryItem, timeCount []int64) {
	retryManager.itemChans <- newRetryItem(item, timeCount)
}

func (this *RetryManager) Start() {
	go this.loop()
	this.workPool = gopool.NewPool(this.workNum)
}

func (this *RetryManager) addServiceToMap(item *retryItem) {
	if this.itemMap[item.item.GetId()] == nil {
		this.itemMap[item.item.GetId()] = item
	}
}

func (this *RetryManager) run() {
	for id, item := range this.itemMap {
		curTime := time.Now().Unix()
		if item.IsDelete() {
			delete(this.itemMap, id)
		} else if item.GetTryTime() <= curTime && !item.state {
			item.state = true
			this.workPool.Add(item)
		}
	}
}

func (this *RetryManager) loop() {
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			this.run()
		case item := <-this.itemChans:
			this.addServiceToMap(item)
		case <-this.done:
			return
		}
	}
}
