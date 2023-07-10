package rpc

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"cqserver/golibs/nw"
)

type callContext struct {
	transId  uint32
	result   nw.ProtoMessage
	userData interface{}
	callback AsyncCallBack
	err      error
	done     chan struct{}
}

type Context interface {
	GetTransId() uint32
	GetResult() nw.ProtoMessage
	GetUserData() interface{}
	SetError(err error)
}

type ContextFinder interface {
	FindContext(transId uint32) Context
}

type RpcWrapper struct {
	Conn         nw.Conn
	timeout      time.Duration
	marshal      Marshaler
	unmarshal    Unmarshaler
	pendingItems map[uint32]*callContext
	pendingMu    sync.Mutex
	idseq        uint32
}

type Marshaler func(context Context, req nw.ProtoMessage) ([]byte, error)
type Unmarshaler func(contextFinder ContextFinder, data []byte) (interface{}, Context, error)
type AsyncCallBack func(context Context)

var ErrRpcTimeOut = errors.New("rpcs call time out")

func NewRpcWrapper(conn nw.Conn, timeout time.Duration, marshal Marshaler, unmarshal Unmarshaler) *RpcWrapper {
	if timeout == 0 {
		timeout = 3 * time.Second
	}
	return &RpcWrapper{
		Conn:         conn,
		timeout:      timeout,
		marshal:      marshal,
		unmarshal:    unmarshal,
		pendingItems: make(map[uint32]*callContext),
	}
}

func (this *callContext) GetTransId() uint32 {
	return this.transId
}

func (this *callContext) GetResult() nw.ProtoMessage {
	return this.result
}

func (this *callContext) GetUserData() interface{} {
	return this.userData
}

func (this *callContext) SetError(err error) {
	this.err = err
}

func (this *RpcWrapper) getNextTransId() uint32 {
	transId := atomic.AddUint32(&this.idseq, 1)
	if transId == 0 {
		transId = atomic.AddUint32(&this.idseq, 1)
	}
	return transId
}

func (this *RpcWrapper) NewContext(resp nw.ProtoMessage, userData interface{}) Context {
	return &callContext{
		transId:  this.getNextTransId(),
		result:   resp,
		userData: userData,
	}
}

func (this *RpcWrapper) FindContext(transId uint32) Context {
	var context Context
	var ok bool
	this.pendingMu.Lock()
	context, ok = this.pendingItems[transId]
	this.pendingMu.Unlock()
	if !ok {
		return nil
	}
	return context
}

func (this *RpcWrapper) addContext(transId uint32, ctx *callContext) {
	this.pendingMu.Lock()
	this.pendingItems[transId] = ctx
	this.pendingMu.Unlock()
}

func (this *RpcWrapper) removeContext(transId uint32) {
	this.pendingMu.Lock()
	delete(this.pendingItems, transId)
	this.pendingMu.Unlock()
}

func (this *RpcWrapper) doCall(ctx *callContext, req nw.ProtoMessage, isSync bool) error {
	transId := ctx.GetTransId()
	this.addContext(transId, ctx)

	rb, err := this.marshal(ctx, req)
	if err != nil {
		return err
	}
	_, err = this.Conn.Write(rb)
	if err != nil {
		return err
	}
	if !isSync {
		return nil
	}
	select {
	case <-ctx.done:
		return ctx.err
	case <-time.After(this.timeout):
		this.removeContext(transId)
		return ErrRpcTimeOut
	}
}

func (this *RpcWrapper) Call(context Context, req nw.ProtoMessage) error {
	ctx := context.(*callContext)
	ctx.done = make(chan struct{})
	return this.doCall(ctx, req, true)
}

func (this *RpcWrapper) AsyncCall(context Context, req nw.ProtoMessage, callback AsyncCallBack) error {
	ctx := context.(*callContext)
	ctx.callback = callback
	return this.doCall(ctx, req, false)
}

func (this *RpcWrapper) HookRecv(data []byte) (interface{}, error) {
	msgFrame, context, err := this.unmarshal(this, data)
	if err != nil {
		return nil, err
	}
	if context == nil {
		return msgFrame, nil
	}
	ctx := context.(*callContext)
	this.removeContext(ctx.GetTransId())
	if ctx.done != nil {
		close(ctx.done)
	}
	if ctx.callback != nil {
		ctx.callback(ctx)
	}
	return nil, nil
}
