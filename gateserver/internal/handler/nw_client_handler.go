package handler

import (
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	//pbserver.Register(pb.CmdCopyReadyAckId, HandleCopyReadyAck)

	//if main.GATE_TEST {
	//	pbclient.RegisterClient(pb.CmdEnterGameReqId, TestHandleEnterGameReq) // 替换认证函数供Test状态使用
	//}
}

// TestHandleEnterGameReq 用于压力测试的模拟用户登录过程
func TestHandleEnterGameReq(conn nw.Conn, msgFrame *pb.MessageFrame) bool {
	//req, session := msgFrame.Body.(*pb.EnterGameReq), conn.GetSession().(*manager.ClientSession)
	//_ = req
	//// 目前是随机分配一个gs
	//gsSession := main.m.gsManager.GetRandomSession()
	//if !gsSession.IsConnected() {
	//	logger.Error("HandleEnterGameReq:Call gsSession error gsSession didnt connect")
	//	session.SendMessageToClient(msgFrame.TransId, builder.BuildErrorAck(errSystemError))
	//	return false
	//}
	//session.gsSession = gsSession
	//err := session.SendUserEnterToGS(false)
	//if err != nil {
	//	logger.Error("HandleEnterGameReq:gsSession.sendEnterToGs:error:%v", err)
	//	session.SendMessageToClient(msgFrame.TransId, builder.BuildErrorAck(errSystemError))
	//	return false
	//}
	return false // 出于安全考虑，不再转发给gs
}

///////////////////////////////////////
// SERVER SIDE MESSAGE
///////////////////////////////////////

func HandleCopyReadyAck(gateConn nw.Conn, clientSession nw.Session, msgFrame *pb.MessageFrame) bool {
	//msg := msgFrame.Body.(*pb.CopyReadyAck)
	//logger.Info("server send EnterGameAck: copyId=%v fightId=%v crossGroup=%v dynamicCfsId=%v", msg.StageId, msg.FightId, msg.CrossGroup, msg.DynamicCfsId)
	//// GateServer保存该玩家的跨服战serverId
	//clientSession.(*manager.ClientSession).BindFightSession(msg.FightId, msg.CrossGroup, int(msg.DynamicCfsId))
	return true
}
