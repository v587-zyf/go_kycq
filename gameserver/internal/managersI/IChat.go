package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

type IChatManager interface {
	// 通用
	ChatSendReq(user *objs.User, chatSendReq *pb.ChatSendReq) (*pb.ChatSendAck, error)
	// 系统消息
	ChatSendSystemMsg(sysMsg string)
	/**
    *  @Description: 协助请求广播
    *  @param user
    *  @return error
    **/
	ChatForFightHelp(user *objs.User) error
}
