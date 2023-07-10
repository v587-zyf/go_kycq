package pbserver

import (
	"cqserver/golibs/nw"
	"cqserver/golibs/nw/rpc"
	"cqserver/golibs/util"
)

func RpcMarshaler(context rpc.Context, req nw.ProtoMessage) ([]byte, error) {
	return Marshal(GetCmdIdFromType(req), context.GetTransId(), true, req)
}

func RpcUnmarshaler(contextFinder rpc.ContextFinder, data []byte) (interface{}, rpc.Context, error) {
	msgFrame, err := unmarshalHeader(data)
	if err != nil {
		return nil, nil, err
	}
	if msgFrame.TransId == 0 || msgFrame.RpcFlag == 1 { // 不是rpc的返回包, RpcFlag==1代表请求而不是返回
		body, err := unmarshalBody(data, msgFrame.CmdId)
		if err != nil {
			return nil, nil, err
		}
		msgFrame.Body = body
		return msgFrame, nil, nil
	}
	context := contextFinder.FindContext(msgFrame.TransId)
	if context == nil { // maybe cancelled
		return nil, nil, nil
	}

	if msgFrame.CmdId != CmdErrorAckId {
		if err := context.GetResult().Unmarshal(data[HeaderSize:]); err != nil {
			return nil, nil, err
		}
	} else {
		var ei = new(ErrorAck)
		if err := ei.Unmarshal(data[HeaderSize:]); err != nil {
			return nil, nil, err
		}
		err := &util.ErrorT{int(ei.Code), ei.Message}
		context.SetError(err)
		return nil, context, nil
	}
	return msgFrame, context, nil
}
