package pb

import (
	"bytes"
	"cqserver/golibs/logger"
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"

	"cqserver/golibs/nw"
	"github.com/astaxie/beego/logs"
)

const HeaderSize = 10
const MaxPacketSize = 1024 * 10240

var errPacketWrongFormat = errors.New("packet has a wrong format or exceeds max size")

type MessageFrame struct {
	CmdId   uint16
	TransId uint32
	Len     uint32
	Body    nw.ProtoMessage
}

type HandlerFunc func(conn nw.Conn, p interface{}) (nw.ProtoMessage, OpGoodsHelper, error)

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

func (this *pbHandler) GetHandler(cmdId uint16) HandlerFunc {
	if h, ok := this.handlers[cmdId]; ok {
		return h.handler
	}
	return nil
}

func Marshal(cmdId uint16, transId uint32, msg nw.ProtoMessage) ([]byte, error) {
	size := msg.Size()
	data := make([]byte, HeaderSize+size)
	n, err := msg.MarshalTo(data[HeaderSize:])
	if err != nil {
		return nil, err
	}
	dataLen := HeaderSize + n
	binary.BigEndian.PutUint16(data[0:2], uint16(cmdId))
	binary.BigEndian.PutUint32(data[2:6], uint32(transId))
	binary.BigEndian.PutUint32(data[6:10], uint32(n))
	if dataLen <= MaxPacketSize {
		return data[:dataLen], nil
	} else {
		return marshalLongPacketNew(cmdId, transId, msg)
	}
}

func marshalLongPacketNew(cmdId uint16, transId uint32, msg nw.ProtoMessage) ([]byte, error) {

	data := make([]byte, msg.Size())
	dataLen, err := msg.MarshalTo(data)
	if err != nil {
		return nil, err
	}

	if cmdId == CmdEnterGameAckId {
		logger.Info("数据长度为size：%v,len:%v", msg.Size(), dataLen)
	}

	var index = 0
	var buf bytes.Buffer
	var nextIndex int

	for {
		isLast := false
		if index+MaxPacketSize <= dataLen {
			nextIndex = index + MaxPacketSize
		} else {
			nextIndex = dataLen
			isLast = true
		}
		var reserved = make([]byte, HeaderSize)
		binary.BigEndian.PutUint16(reserved[0:2], uint16(cmdId))
		binary.BigEndian.PutUint32(reserved[2:6], uint32(transId))
		binary.BigEndian.PutUint16(reserved[6:8], uint16(dataLen))
		buf.Write(reserved)
		buf.Write(data[index:nextIndex])
		index = nextIndex
		if isLast {
			break
		}
	}
	return buf.Bytes(), nil
}

//
//func marshalLongPacket(data []byte) ([]byte, error) {
//	var dataLen = len(data)
//	var index = 0
//	var buf bytes.Buffer
//	var reserved = make([]byte, HeaderSize)
//	var nextIndex int
//	for {
//		if index+MaxPacketSize-HeaderSize <= dataLen {
//			nextIndex = index + MaxPacketSize - HeaderSize
//		} else {
//			nextIndex = dataLen
//		}
//		isLast := nextIndex >= dataLen
//		binary.BigEndian.PutUint16(reserved[0:2], uint16(CmdComposeDataAckId))
//		endMark := 0
//		if isLast {
//			endMark = 1
//		}
//		binary.BigEndian.PutUint32(reserved[2:6], uint32(endMark))
//		binary.BigEndian.PutUint16(reserved[6:8], uint16(nextIndex-index+HeaderSize))
//		buf.Write(reserved)
//		buf.Write(data[index:nextIndex])
//		if isLast {
//			break
//		}
//		index = nextIndex
//	}
//	return buf.Bytes(), nil
//}

func Unmarshal(data []byte) (*MessageFrame, error) {
	if len(data) < HeaderSize {
		return nil, errors.New("packet has a wrong header")
	}
	frame := &MessageFrame{}
	frame.CmdId = binary.BigEndian.Uint16(data[0:2])
	frame.TransId = binary.BigEndian.Uint32(data[2:6])
	frame.Len = binary.BigEndian.Uint32(data[6:10])
	if msgPrototype := GetMsgPrototype(frame.CmdId); msgPrototype != nil {
		body := reflect.New(msgPrototype).Interface().(nw.ProtoMessage)
		if err := body.Unmarshal(data[HeaderSize:]); err != nil {
			return nil, err
		}
		frame.Body = body
		return frame, nil
	}
	return nil, fmt.Errorf("Unmarshl error, cmdId: %d, data length: %d", frame.CmdId, len(data))
}

func unmarshalHeader(data []byte) (*MessageFrame, error) {
	if len(data) < HeaderSize {
		return nil, errors.New("packet has a wrong header")
	}
	frame := &MessageFrame{}
	frame.CmdId = binary.BigEndian.Uint16(data[0:2])
	frame.TransId = binary.BigEndian.Uint32(data[2:6])
	frame.Len = binary.BigEndian.Uint32(data[6:10])
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

func GetCmdId(data []byte) uint16 {
	if len(data) < HeaderSize {
		return 0
	}
	return binary.BigEndian.Uint16(data[0:2])
}

func TraceMsg(l *logs.BeeLogger, data []byte) {
	var cmdId = binary.BigEndian.Uint16(data[0:2])
	l.Debug("recved %s, data len: %d", msgNames[cmdId], len(data))
}

func Split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	dataLen := len(data)
	if dataLen < HeaderSize {
		return 0, nil, nil
	}
	n := int(binary.BigEndian.Uint16(data[6:10]))
	if n == 0 || n > MaxPacketSize {
		return 0, nil, errPacketWrongFormat
	}
	if dataLen < n {
		return 0, nil, nil
	}
	return n, data[0:n], nil
}
