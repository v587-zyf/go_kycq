package handler

import (
	"cqserver/gameserver/internal/managers"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdChatSendReqId, HandlerChatSendReq)
}

func HandlerChatSendReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ChatSendReq)
	user := conn.GetSession().(*managers.ClientSession).User

	_, err := m.Chat.ChatSendReq(user, req)
	if err != nil {
		return nil, nil, err
	}

	return nil, nil, nil
}
