package util

type Message interface {
	Handle()
	Wait()
	Done()
}

type SyncMessage struct {
	done chan struct{}
}

func NewSyncMessage() *SyncMessage {
	return &SyncMessage{
		done: make(chan struct{}),
	}
}

func (this *SyncMessage) Done() {
	close(this.done)
}

func (this *SyncMessage) Wait() {
	<-this.done
}

type AsyncMessage struct {
}

func NewAsyncMessage() *AsyncMessage {
	return new(AsyncMessage)
}

func (this AsyncMessage) Done() {
}

func (this AsyncMessage) Wait() {
}
