package pbgt

import (
	"bytes"
	"encoding/binary"
	"errors"
	fmt "fmt"
	"reflect"

	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

const HeaderSize = 18
const MaxPacketSize = 1 * 1024 * 1024 // 允许的最大封包长度，不能超过网络接收缓冲区的大小
const ChecksumMark = 0xAA55AA55       // 校验码，用于判断是否是一个合法的包头

var ChecksumMarkArray = []byte{0xAA, 0x55, 0xAA, 0x55}
var log = logger.Get("default", true)
type MessageFrame struct {
	Len       uint32
	CmdId     uint16
	SessionId uint32
	Reserved  uint32
	Body      interface{}
}

type HandlerFunc func(conn nw.Conn, msgFrame *MessageFrame)

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

func (this *pbHandler) GetHandler(cmdId uint16) HandlerFunc {
	if h, ok := this.handlers[cmdId]; ok {
		return h.handler
	}
	return nil
}

func (this *pbHandler) Unmarshal(data []byte) (*MessageFrame, error) {
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

func Marshal(cmdId uint16, sessionId uint32, reserved uint32, msg interface{}) ([]byte, error) {
	var data []byte
	var dataLen int
	if cmdId == CmdRouteMessageId { // 中转消息的body直接是[]byte
		if bytesMsg, ok := msg.([]byte); ok {
			dataLen = HeaderSize + len(bytesMsg)
			data = make([]byte, dataLen)
			copy(data[HeaderSize:], bytesMsg)
		} else {
			return nil, errors.New("need []byte to Marshal")
		}
	} else { // 其他消息的body为proto.Message
		if protoMsg, ok := msg.(nw.ProtoMessage); ok {
			size := protoMsg.Size()
			data = make([]byte, HeaderSize+size)
			n, err := protoMsg.MarshalTo(data[HeaderSize:])
			if err != nil {
				return nil, err
			}
			dataLen = HeaderSize + n
		} else {
			return nil, errors.New("need proto.Message to Marshal")
		}
	}

	binary.BigEndian.PutUint32(data[0:4], uint32(ChecksumMark))
	binary.BigEndian.PutUint32(data[4:8], uint32(dataLen))
	binary.BigEndian.PutUint16(data[8:10], uint16(cmdId))
	binary.BigEndian.PutUint32(data[10:14], uint32(sessionId))
	binary.BigEndian.PutUint32(data[14:18], uint32(reserved))
	return data, nil
}

func MarshalRouteMsg(sessionId uint32, reserved uint32, transId uint32, msg nw.ProtoMessage) ([]byte, error) {
	var headerSizeAll int = HeaderSize + pb.HeaderSize
	pbgtCmd := CmdRouteMessageId
	cmdId := pb.GetCmdIdFromType(msg)
	data := make([]byte, headerSizeAll+msg.Size())
	n, err := msg.MarshalTo(data[headerSizeAll:])
	if err != nil {
		return nil, err
	}
	dataLen := pb.HeaderSize + n
	binary.BigEndian.PutUint16(data[HeaderSize:HeaderSize+2], uint16(cmdId))
	binary.BigEndian.PutUint32(data[HeaderSize+2:HeaderSize+6], uint32(transId))
	binary.BigEndian.PutUint16(data[HeaderSize+6:HeaderSize+8], uint16(dataLen))
	dataLen = headerSizeAll + n
	binary.BigEndian.PutUint32(data[0:4], uint32(ChecksumMark))
	binary.BigEndian.PutUint32(data[4:8], uint32(dataLen))
	binary.BigEndian.PutUint16(data[8:10], uint16(pbgtCmd))
	binary.BigEndian.PutUint32(data[10:14], uint32(sessionId))
	binary.BigEndian.PutUint32(data[14:18], uint32(reserved))
	return data, nil
}

func Unmarshal(data []byte) (*MessageFrame, error) {
	if len(data) < HeaderSize {
		return nil, errors.New("packet has a wrong header")
	}

	frame := &MessageFrame{}
	frame.Len = binary.BigEndian.Uint32(data[4:8])
	frame.CmdId = binary.BigEndian.Uint16(data[8:10])
	frame.SessionId = binary.BigEndian.Uint32(data[10:14])
	frame.Reserved = binary.BigEndian.Uint32(data[14:18])
	if frame.CmdId == CmdRouteMessageId {
		frame.Body = data[HeaderSize:]
		return frame, nil
	} else if msgPrototype := GetMsgPrototype(frame.CmdId); msgPrototype != nil {
		body := reflect.New(msgPrototype).Interface().(nw.ProtoMessage)
		if err := body.Unmarshal(data[HeaderSize:]); err != nil {
			return nil, err
		}
		frame.Body = body
		return frame, nil
	}
	return nil, fmt.Errorf("Unmarshl error, cmdId: %d, data length: %d", frame.CmdId, len(data))
}

func GetCmdId(data []byte) uint16 {
	if len(data) < HeaderSize {
		return 0
	}
	return binary.BigEndian.Uint16(data[8:10])
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
