package pbclient

import (
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"

	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

const HeaderSize = 10

type ClientHandler func(conn nw.Conn, msgFrame *pb.MessageFrame) bool
type ServerHandler func(conn nw.Conn, session nw.Session, msgFrame *pb.MessageFrame) bool

type clientHandlerItem struct {
	cmdId   uint16
	handler ClientHandler
}

type serverHandlerItem struct {
	cmdId   uint16
	handler ServerHandler
}

var clientHandlers = make(map[uint16]*clientHandlerItem)
var serverHandlers = make(map[uint16]*serverHandlerItem)

func RegisterClient(cmdId uint16, handler ClientHandler) {
	clientHandlers[cmdId] = &clientHandlerItem{
		cmdId:   cmdId,
		handler: handler,
	}
}

func RegisterServer(cmdId uint16, handler ServerHandler) {
	serverHandlers[cmdId] = &serverHandlerItem{
		cmdId:   cmdId,
		handler: handler,
	}
}

func GetClientHandler(cmdId uint16) ClientHandler {
	if h, ok := clientHandlers[cmdId]; ok {
		return h.handler
	}
	return nil
}

func GetServerHandler(cmdId uint16) ServerHandler {
	if h, ok := serverHandlers[cmdId]; ok {
		return h.handler
	}
	return nil
}

func UnmarshalClient(data []byte) (*pb.MessageFrame, error) {
	if len(data) < HeaderSize {
		return nil, fmt.Errorf("UnmarshalClient :packet has a wrong header:%v", data)
	}

	msgFrame := &pb.MessageFrame{}
	msgFrame.CmdId = binary.BigEndian.Uint16(data[0:2])
	msgFrame.TransId = binary.BigEndian.Uint32(data[2:6])
	msgFrame.Len = binary.BigEndian.Uint32(data[6:10])

	if msgPrototype := pb.GetMsgPrototype(msgFrame.CmdId); msgPrototype != nil {
		body := reflect.New(msgPrototype).Interface().(nw.ProtoMessage)
		if err := body.Unmarshal(data[HeaderSize:]); err != nil {
			return nil, err
		}
		msgFrame.Body = body
		return msgFrame, nil
	}
	return nil, errors.New("cmd not found")
}

func UnmarshalServer(data []byte) (*pb.MessageFrame, error) {
	if len(data) < HeaderSize {
		return nil, fmt.Errorf("UnmarshalServer :packet has a wrong header:%v", data)
	}

	msgFrame := &pb.MessageFrame{}
	msgFrame.CmdId = binary.BigEndian.Uint16(data[0:2])
	msgFrame.TransId = binary.BigEndian.Uint32(data[2:6])
	msgFrame.Len = binary.BigEndian.Uint32(data[6:10])

	if msgPrototype := pb.GetMsgPrototype(msgFrame.CmdId); msgPrototype != nil {
		body := reflect.New(msgPrototype).Interface().(nw.ProtoMessage)
		if err := body.Unmarshal(data[HeaderSize:]); err != nil {
			return nil, err
		}
		msgFrame.Body = body
		return msgFrame, nil
	}
	return nil, errors.New("cmd not found")
}

func GetCmdId(data []byte) uint16 {
	return binary.BigEndian.Uint16(data[0:2])
}

var Marshal = pb.Marshal
var Split = pb.Split
