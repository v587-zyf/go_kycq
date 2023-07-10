package pbserver

import (
	"bytes"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"reflect"
)

const HeaderSize = 20
const MaxPacketSize = 1 * 1024 * 1024 // 允许的最大封包长度，不能超过网络接收缓冲区的大小
const ChecksumMark = 0xAA55AA55       // 校验码，用于判断是否是一个合法的包头

var ChecksumMarkArray = []byte{0xAA, 0x55, 0xAA, 0x55}
var log = logger.Get("default", true)

type MessageFrame struct {
	Len      uint32
	CmdId    uint16
	TransId  uint32 // rpc请求中的序列id
	RpcFlag  uint16 // 是否是rpc请求
	Reserved uint32
	Body     nw.ProtoMessage
}

type HandlerFunc func(conn nw.Conn, msgFrame *MessageFrame) (nw.ProtoMessage, error)

type handlerItem struct {
	cmdId   uint16
	handler HandlerFunc
}

// pbHandler manage protobuf handlers
type pbHandler struct {
	handlers map[uint16]*handlerItem
}

var Handler *pbHandler = &pbHandler{handlers: make(map[uint16]*handlerItem)}

func (this *pbHandler) Register(cmdId uint16, handler HandlerFunc) {
	this.handlers[cmdId] = &handlerItem{
		cmdId:   cmdId,
		handler: handler,
	}
}

func (this *pbHandler) HasHandler(cmdId uint16) bool {
	_, ok := this.handlers[cmdId]
	return ok
}

func (this *pbHandler) Handle(conn nw.Conn, msg interface{}) {
	msgFrame := msg.(*MessageFrame)
	h := this.handlers[msgFrame.CmdId]
	fmt.Println("收到消息", msgFrame.CmdId)
	ack, err := h.handler(conn, msgFrame)
	_ = ack
	_ = err
}

func (this *pbHandler) Unmarshal(data []byte) (interface{}, error) {
	return Unmarshal(data)
}

func Register(cmdId uint16, handler HandlerFunc) {
	Handler.Register(cmdId, handler)
}

func HasHandler(cmdId uint16) bool {
	return Handler.HasHandler(cmdId)
}

func GetHandler(cmdId uint16) HandlerFunc {
	return Handler.GetHandler(cmdId)
}

func Handle(conn nw.Conn, msg interface{}) {
	Handler.Handle(conn, msg)
}

func (this *pbHandler) GetHandler(cmdId uint16) HandlerFunc {
	if h, ok := this.handlers[cmdId]; ok {
		return h.handler
	}
	return nil
}

func Marshal(cmdId uint16, transId uint32, isRpcReq bool, msg nw.ProtoMessage) ([]byte, error) {
	size := msg.Size()
	data := make([]byte, HeaderSize+size)
	n, err := msg.MarshalTo(data[HeaderSize:])
	if err != nil {
		return nil, err
	}
	rpcFlag := 0
	if isRpcReq {
		rpcFlag = 1
	}
	dataLen := HeaderSize + n
	binary.BigEndian.PutUint32(data[0:4], uint32(ChecksumMark))
	binary.BigEndian.PutUint32(data[4:8], uint32(dataLen))
	binary.BigEndian.PutUint16(data[8:10], uint16(cmdId))
	binary.BigEndian.PutUint32(data[10:14], uint32(transId))
	binary.BigEndian.PutUint16(data[14:16], uint16(rpcFlag))
	binary.BigEndian.PutUint32(data[16:20], uint32(0))
	return data[:dataLen], nil
}

func Unmarshal(data []byte) (*MessageFrame, error) {
	if len(data) < HeaderSize {
		return nil, errors.New("packet has a wrong header")
	}
	frame := &MessageFrame{}
	frame.CmdId = binary.BigEndian.Uint16(data[8:10])
	frame.TransId = binary.BigEndian.Uint32(data[10:14])
	frame.RpcFlag = binary.BigEndian.Uint16(data[14:16])
	if msgPrototype, ok := msgPrototypes[frame.CmdId]; ok {
		body := reflect.New(msgPrototype).Interface().(nw.ProtoMessage)
		if err := body.Unmarshal(data[HeaderSize:]); err != nil {
			return nil, err
		}
		frame.Body = body
		return frame, nil
	}
	return nil, errors.New("cmd not found " + string(frame.CmdId))
}

func unmarshalHeader(data []byte) (*MessageFrame, error) {
	if len(data) < HeaderSize {
		return nil, errors.New("packet has a wrong header")
	}
	frame := &MessageFrame{}
	frame.CmdId = binary.BigEndian.Uint16(data[8:10])
	frame.TransId = binary.BigEndian.Uint32(data[10:14])
	frame.RpcFlag = binary.BigEndian.Uint16(data[14:16])
	return frame, nil
}

func unmarshalBody(data []byte, cmdId uint16) (nw.ProtoMessage, error) {
	if msgPrototype, ok := msgPrototypes[cmdId]; ok {
		body := reflect.New(msgPrototype).Interface().(nw.ProtoMessage)
		if err := body.Unmarshal(data[HeaderSize:]); err != nil {
			return nil, err
		}
		return body, nil
	}
	return nil, errors.New("cmd not found")
}

func UnmarshalMsgByCmdId(data []byte, cmdId uint16) (nw.ProtoMessage, error) {
	if msgPrototype, ok := msgPrototypes[cmdId]; ok {
		body := reflect.New(msgPrototype).Interface().(nw.ProtoMessage)
		if err := body.Unmarshal(data); err != nil {
			return nil, err
		}
		return body, nil
	}
	return nil, errors.New("cmd not found")
}

func TraceMsg(l *logs.BeeLogger, data []byte) {
	var cmdId = binary.BigEndian.Uint16(data[8:10])
	var rpcFlag = binary.BigEndian.Uint16(data[14:16])
	l.Debug("recved %s, rpcFlag: %d, data len: %d", msgNames[cmdId], rpcFlag, len(data))
}

func Split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	dataLen := len(data)
	if dataLen < HeaderSize {
		return 0, nil, nil
	}
	mark := binary.BigEndian.Uint32(data[0:4])
	if mark != ChecksumMark {
		index := bytes.Index(data, ChecksumMarkArray)
		if index <= 0 {
			// 没有招到ChecksumMark，丢弃所有数据
			log.Error("Split: wrong packet received, dataLen: %d", dataLen)
			return dataLen, nil, nil
		} else {
			// 招到了ChecksumMark，丢弃在这之前的数据
			log.Error("Split: wrong packet received, dataLen: %d", index)
			return index, nil, nil
		}
	}
	n := int(binary.BigEndian.Uint32(data[4:8]))
	if n <= 0 || n > MaxPacketSize {
		// 检测到长度不合法的包，丢弃4字节（ChecksumMark的长度)，以寻找下一个ChecksumMark
		log.Error("Split: wrong packet received, length in header: %d", n)
		return 4, nil, nil
	}
	if dataLen < n {
		return 0, nil, nil
	}
	return n, data[0:n], nil
}
