package net

import (
	"cqserver/golibs/nw"
	"cqserver/protobuf/pbgt"
	"cqserver/protobuf/pbserver"
	"errors"
	"sync"
)

var (
	TransIdToServerId  = sync.Map{} //key:transid val:serverId
	TransIdToGsTransId = sync.Map{} //key:transid val:GsTransId
)

//获取从战斗中心转发的gate的消息
func GetFccRouteFromGateMsgFrame(msgFrame *pbgt.MessageFrame) (*pbgt.MessageFrame, int32, error) {
	tmpmsg := msgFrame.Body.(*pbgt.GateMessageToFS)
	resmsg, err := pbgt.Unmarshal(tmpmsg.GetMsg())
	if err != nil {
		return nil, -1, err
	}

	return resmsg, tmpmsg.ServerId, nil
}

//获取从战斗中心转发给gate的消息
func GetFccRouteToGateMessageBytes(serverId uint32, data []byte) ([]byte, error) {

	resp := &pbgt.FSMessageToGate{
		CrossServerId: 0,
		ServerId:      int32(serverId),
		Msg:           data,
	}
	rb, err := pbgt.Marshal(pbgt.CmdFSMessageToGateId, 0, 0, resp)
	if err != nil {
		return nil, err
	}
	return rb, err
}

//gs recv
func GetFssRouteFromGsMsgFrame(msgFrame *pbserver.MessageFrame) (*pbserver.MessageFrame, error) {
	tmpmsg := msgFrame.Body.(*pbserver.GSMessageToFS)
	resmsg, err := pbserver.Unmarshal(tmpmsg.GetMsg())
	if err != nil {
		return nil, err
	}

	transId := msgFrame.TransId //center的transId
	msgFrame = resmsg
	msgFrame.TransId = transId

	TransIdToServerId.Store(transId, tmpmsg.ServerId)
	TransIdToGsTransId.Store(transId, tmpmsg.GsTransId)
	return msgFrame, nil
}

//这里是rpc调用返回时使用
func GetFccRouteToGsMessageBytes(transId uint32, msg nw.ProtoMessage) ([]byte, error) {

	serverId, ok := TransIdToServerId.Load(transId)
	gsTransId, _ := TransIdToGsTransId.Load(transId)
	if transId > 0 && ok {
		gId := gsTransId.(int32)
		rb, err := pbserver.Marshal(pbserver.GetCmdIdFromType(msg), uint32(gId), false, msg)
		if err != nil {
			return nil, err
		}

		resp := &pbserver.FSMessageToGS{
			GsTransId: gId,
			ServerId:  serverId.(int32),
			Msg:       rb,
		}

		rb2, err2 := pbserver.Marshal(pbserver.GetCmdIdFromType(resp), transId, false, resp)
		if err2 != nil {
			return nil, err2
		}

		return rb2, nil
	}
	return nil, errors.New("DYNAMIC server transId error")

}

func GetGateNewMsgBytesNormal(serverId int, data []byte) ([]byte, error) {
	resp := &pbgt.FSMessageToGate{
		CrossServerId: 0,
		ServerId:      int32(serverId),
		Msg:           data,
	}
	rb, err := pbgt.Marshal(pbgt.CmdFSMessageToGateId, 0, 0, resp)
	if err != nil {
		return nil, err
	}

	return rb, nil
}
