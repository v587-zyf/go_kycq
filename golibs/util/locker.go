package util

import (
	"fmt"
	"runtime/debug"
	"sync"
	"time"
)

/*
由于Go不提供超时锁，所以自己实现了支持超时机制的互斥锁Locker和读写锁RWLocker。
为了方便供第三方程序使用，提供了根据Key获取超时互斥锁和超时读写锁的复合对象LockerUtil和RWLockerUtil。
为了在出现锁超时时方便查找问题，会记录上次成功获得锁时的堆栈信息；并且在本次获取锁失败时，同时返回上次成功时的堆栈信息和本次的堆栈信息。
*/

const (
	// 默认超时的毫秒数(1小时)
	con_Default_Timeout_Milliseconds = 60 * 60 * 1000

	// 写锁每次休眠的时间比读锁的更短，这样是因为写锁有更高的优先级，所以尝试的频率更大
	// 写锁每次休眠的毫秒数
	con_Lock_Sleep_Millisecond = 1

	// 读锁每次休眠的毫秒数
	con_RLock_Sleep_Millisecond = 2

	//默认解锁时间
	default_time_out = 1000
)

// 获取超时时间
func getTimeout(timeout int) int {
	if timeout > 0 {
		return timeout
	} else {
		return con_Default_Timeout_Milliseconds
	}
}

// 写锁对象
type Locker struct {
	timeout   int
	write     int // 使用int而不是bool值的原因，是为了与RWLocker中的read保持类型的一致；
	prevStack []byte
	mutex     sync.Mutex
	isDebug   bool
}

// 内部锁
// 返回值：
// 加锁是否成功
func (this *Locker) lock() bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	// 如果已经被锁定，则返回失败
	if this.write == 1 {
		return false
	}

	// 否则，将写锁数量设置为１，并返回成功
	this.write = 1

	// 记录Stack信息
	if this.isDebug {
		this.prevStack = debug.Stack()
	}

	return true
}

// 尝试加锁，如果在指定的时间内失败，则会返回失败；否则返回成功
// timeout:指定的毫秒数,timeout<=0则将会死等
// 返回值：
// 成功或失败
// 如果失败，返回上一次成功加锁时的堆栈信息
// 如果失败，返回当前的堆栈信息
func (this *Locker) Lock() (successful bool, prevStack string, currStack string) {

	// 遍历指定的次数（即指定的超时时间）
	for i := 0; i < this.timeout; i = i + con_Lock_Sleep_Millisecond {
		// 如果锁定成功，则返回成功
		if this.lock() {
			successful = true
			break
		}

		// 如果锁定失败，则休眠con_Lock_Sleep_Millisecond ms，然后再重试
		time.Sleep(con_Lock_Sleep_Millisecond * time.Millisecond)
	}

	// 如果时间结束仍然是失败，则返回上次成功的堆栈信息，以及当前的堆栈信息
	if successful == false {
		prevStack = string(this.prevStack)
		currStack = string(debug.Stack())
		fmt.Println(prevStack)
	}
	return
}

// 锁定（死等方式）
func (this *Locker) WaitLock() {
	successful, prevStack, currStack := this.Lock()
	if successful == false {
		fmt.Printf("Locker.WaitLock():{PrevStack:%s, currStack:%s}\n", prevStack, currStack)
	}
}

// 解锁
func (this *Locker) Unlock() {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.write = 0
}

func (this *Locker) SetDebug(debug bool) {
	this.isDebug = debug
	this.timeout = 5000
}

// 创建新的锁对象
func NewLocker(timeout int) *Locker {
	if timeout <= 0 {
		timeout = default_time_out
	}
	return &Locker{
		timeout: timeout,
	}
}

// 读写锁对象
type RWLocker struct {
	timeout   int
	read      int
	write     int // 使用int而不是bool值的原因，是为了与read保持类型的一致；
	prevStack []byte
	mutex     sync.Mutex
	isDebug   bool
}

// 尝试加写锁
// 返回值：加写锁是否成功
func (this *RWLocker) lock() bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	// 如果已经被锁定，则返回失败
	if this.write == 1 || this.read > 0 {
		return false
	}

	// 否则，将写锁数量设置为１，并返回成功
	this.write = 1

	// 记录Stack信息
	this.prevStack = debug.Stack()

	return true
}

// 写锁定
// timeout:超时毫秒数,timeout<=0则将会死等
// 返回值：
// 成功或失败
// 如果失败，返回上一次成功加锁时的堆栈信息
// 如果失败，返回当前的堆栈信息
func (this *RWLocker) Lock() (successful bool, prevStack string, currStack string) {
	//timeout = getTimeout(timeout)

	// 遍历指定的次数（即指定的超时时间）
	for i := 0; i < this.timeout; i = i + con_Lock_Sleep_Millisecond {
		// 如果锁定成功，则返回成功
		if this.lock() {
			successful = true
			break
		}

		// 如果锁定失败，则休眠con_Lock_Sleep_Millisecond ms，然后再重试
		time.Sleep(con_Lock_Sleep_Millisecond * time.Millisecond)
	}

	// 如果时间结束仍然是失败，则返回上次成功的堆栈信息，以及当前的堆栈信息
	if successful == false {
		prevStack = string(this.prevStack)
		currStack = string(debug.Stack())
	}

	return
}

// 写锁定(死等)
func (this *RWLocker) WaitLock() {
	successful, prevStack, currStack := this.Lock()
	if successful == false {
		fmt.Printf("RWLocker:WaitLock():{PrevStack:%s, currStack:%s}\n", prevStack, currStack)
	}
}

// 释放写锁
func (this *RWLocker) Unlock() {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	this.write = 0
}

// 尝试加读锁
// 返回值：加读锁是否成功
func (this *RWLocker) rlock() bool {
	this.mutex.Lock()
	defer this.mutex.Unlock()

	// 如果已经被锁定，则返回失败
	if this.write == 1 {
		return false
	}

	// 否则，将读锁数量加１，并返回成功
	this.read += 1

	// 记录Stack信息
	if this.isDebug {
		this.prevStack = debug.Stack()
	}

	return true
}

// 读锁定
// timeout:超时毫秒数,timeout<=0则将会死等
// 返回值：
// 成功或失败
// 如果失败，返回上一次成功加锁时的堆栈信息
// 如果失败，返回当前的堆栈信息
func (this *RWLocker) RLock() (successful bool, prevStack string, currStack string) {

	// 遍历指定的次数（即指定的超时时间）
	// 读锁比写锁优先级更低，所以每次休眠2ms，所以尝试的次数就是时间/2
	for i := 0; i < this.timeout; i = i + con_RLock_Sleep_Millisecond {
		// 如果锁定成功，则返回成功
		if this.rlock() {
			successful = true
			break
		}

		// 如果锁定失败，则休眠2ms，然后再重试
		time.Sleep(con_RLock_Sleep_Millisecond * time.Millisecond)
	}

	// 如果时间结束仍然是失败，则返回上次成功的堆栈信息，以及当前的堆栈信息
	if successful == false {
		prevStack = string(this.prevStack)
		currStack = string(debug.Stack())
	}

	return
}

// 读锁定(死等)
func (this *RWLocker) WaitRLock() {
	successful, prevStack, currStack := this.RLock()
	if successful == false {
		fmt.Printf("RWLocker:WaitRLock():{PrevStack:%s, currStack:%s}\n", prevStack, currStack)
	}
}

// 释放读锁
func (this *RWLocker) RUnlock() {
	this.mutex.Lock()
	defer this.mutex.Unlock()
	if this.read > 0 {
		this.read -= 1
	}
}

func (this *RWLocker) SetDebug(debug bool) {
	this.isDebug = debug
}

// 创建新的读写锁对象
func NewRWLocker(timeout int) *RWLocker {
	if timeout <= 0 {
		timeout = default_time_out
	}
	return &RWLocker{
		timeout: timeout,
	}
}
