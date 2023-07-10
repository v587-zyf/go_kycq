package manager

import (
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"cqserver/robots/conf"
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"
)

type ClientSessionMessage struct {
	util.AsyncMessage
	Session  *ClientSession
	MsgFrame *pb.MessageFrame
}

func (this *ClientSessionMessage) Handle() {
	this.Session.HandleMessage(this.MsgFrame)
}

type ClientSession struct {
	Id           uint32
	Conn         nw.Conn
	done         chan struct{}
	Robot        *Robot
	recvMsgCount int64
	longData     map[uint16][]byte
}

func NewClientSession(conn nw.Conn) *ClientSession {
	s := &ClientSession{
		Id:       idAllocator.Get(),
		Conn:     conn,
		done:     make(chan struct{}),
		longData: make(map[uint16][]byte),
	}
	return s
}

func (this *ClientSession) GetId() uint32 {
	return this.Id
}

func (this *ClientSession) GetConn() nw.Conn {
	return this.Conn
}

func (this *ClientSession) OnOpen(conn nw.Conn) {
	if this.Robot.status == conf.STATUS_NONE {
		this.Robot.status = conf.STATUS_CONNECTED
	}
}

func (this *ClientSession) OnClose(conn nw.Conn) {
	logger.Info("ClientSession::OnClose:%s", this.Robot.openId)
	this.Robot.status = conf.STATUS_OFF
	//if this.Robot.status > conf.STATUS_CONNECTED2 {
	//	this.Robot.status = conf.STATUS_NONE
	//}
}

func Unmarshals(clientsession *ClientSession, data0 []byte) ([]*pb.MessageFrame, error) {
	data := data0
	if len(data) < pb.HeaderSize {
		return nil, errors.New("packet has a wrong header")
	}
	var out []*pb.MessageFrame
	//var perLen int
	//totalLen := len(data)
	for {
		frame := &pb.MessageFrame{}
		frame.CmdId = binary.BigEndian.Uint16(data[0:2])
		frame.TransId = binary.BigEndian.Uint32(data[2:6])
		frame.Len = binary.BigEndian.Uint32(data[6:10])

		if uint32(len(data)-pb.HeaderSize) != frame.Len {
			if msgPrototype := pb.GetMsgPrototype(frame.CmdId); msgPrototype != nil {
				body := reflect.New(msgPrototype).Interface().(nw.ProtoMessage)
				if err := body.Unmarshal(data[pb.HeaderSize:frame.Len]); err != nil {
					return nil, err
				}
				frame.Body = body
				out = append(out, frame)

			} else {
				return nil, fmt.Errorf("Unmarshl error, cmdId: %d, data length: %d", frame.CmdId, len(data))
			}
		} else {
			clientsession.longData[frame.CmdId] = append(clientsession.longData[frame.CmdId], data[pb.HeaderSize:]...)
			if uint32(len(clientsession.longData[frame.CmdId])) == frame.Len {
				if msgPrototype := pb.GetMsgPrototype(frame.CmdId); msgPrototype != nil {
					body := reflect.New(msgPrototype).Interface().(nw.ProtoMessage)
					if err := body.Unmarshal(clientsession.longData[frame.CmdId]); err != nil {
						return nil, err
					}
					frame.Body = body
					out = append(out, frame)
					delete(clientsession.longData, frame.CmdId)

				} else {
					return nil, fmt.Errorf("Unmarshl error, cmdId: %d, data length: %d", frame.CmdId, len(data))
				}
			}
		}

		break
	}

	return out, nil

}

func (this *ClientSession) OnRecv(conn nw.Conn, data []byte) {

	msgFrames, err := Unmarshals(this, data)
	if err != nil {
		return
	}
	if len(this.Robot.Messages)+len(msgFrames) >= cap(this.Robot.Messages) { //不让堵塞
		return
	}
	for _, msgFrame := range msgFrames {
		this.Robot.Messages <- &ClientSessionMessage{Session: this, MsgFrame: msgFrame}
	}
}

func (this *ClientSession) HandleMessage(msgFrame *pb.MessageFrame) {
	handler := pb.GetHandler(msgFrame.CmdId)
	if handler == nil {
		//logger.Debug("no handler found, cmdId: %d", msgFrame.CmdId)
		return
	}

	_, _, err := handler(this.Conn, msgFrame.Body)
	if err != nil {
		logger.Error("nw_client_session:HandleMessage:err:%v", err)
	}
	return
}

func (this *ClientSession) SendMessage(transId uint32, msg nw.ProtoMessage) error {
	rb, err := pb.Marshal(pb.GetCmdIdFromType(msg), transId, msg)
	if err != nil {
		return err
	}
	_, err = this.Conn.Write(rb)
	return err
}

func (this *ClientSession) SendMessageToClient(transId uint32, msg nw.ProtoMessage) error {
	rb, err := pb.Marshal(pb.GetCmdIdFromType(msg), transId, msg)
	if err != nil {
		return err
	}
	_, err = this.Conn.Write(rb)
	return err
}

func (this *ClientSession) HandleErrorAck(msgFrame *pb.MessageFrame) {
	ack := msgFrame.Body.(*pb.ErrorAck)
	logger.Warn("server error: %s", ack.Message)
}
