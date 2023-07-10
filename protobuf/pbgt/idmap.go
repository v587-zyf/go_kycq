// This is generated by git@gitlab.hd.com:yulong/genproto.git
// Do not modify here.

package pbgt
import (
    "reflect"
)

var msgPrototypes = make(map[uint16]reflect.Type)
var msgNames = make(map[uint16]string)
var msgLogLv = make(map[int]int)

func init() {
    msgPrototypes[CmdHandShakeReqId] = reflect.TypeOf((*HandShakeReq)(nil)).Elem()
    msgPrototypes[CmdHandShakeAckId] = reflect.TypeOf((*HandShakeAck)(nil)).Elem()
    msgPrototypes[CmdUserQuitRptId] = reflect.TypeOf((*UserQuitRpt)(nil)).Elem()
    msgPrototypes[CmdUserQuitNtfId] = reflect.TypeOf((*UserQuitNtf)(nil)).Elem()
    msgPrototypes[CmdRouteMessageId] = reflect.TypeOf((*RouteMessage)(nil)).Elem()
    msgPrototypes[CmdBroadcastNtfId] = reflect.TypeOf((*BroadcastNtf)(nil)).Elem()
    msgPrototypes[CmdUserFightInfoNtfId] = reflect.TypeOf((*UserFightInfoNtf)(nil)).Elem()
    msgPrototypes[CmdBroadcastByFSId] = reflect.TypeOf((*BroadcastByFS)(nil)).Elem()
    msgPrototypes[CmdServerPingReqId] = reflect.TypeOf((*ServerPingReq)(nil)).Elem()
    msgPrototypes[CmdServerPingAckId] = reflect.TypeOf((*ServerPingAck)(nil)).Elem()
    msgPrototypes[CmdGateMessageToFSId] = reflect.TypeOf((*GateMessageToFS)(nil)).Elem()
    msgPrototypes[CmdFSMessageToGateId] = reflect.TypeOf((*FSMessageToGate)(nil)).Elem()
    
    msgNames[CmdHandShakeReqId] = "HandShakeReq"
    msgNames[CmdHandShakeAckId] = "HandShakeAck"
    msgNames[CmdUserQuitRptId] = "UserQuitRpt"
    msgNames[CmdUserQuitNtfId] = "UserQuitNtf"
    msgNames[CmdRouteMessageId] = "RouteMessage"
    msgNames[CmdBroadcastNtfId] = "BroadcastNtf"
    msgNames[CmdUserFightInfoNtfId] = "UserFightInfoNtf"
    msgNames[CmdBroadcastByFSId] = "BroadcastByFS"
    msgNames[CmdServerPingReqId] = "ServerPingReq"
    msgNames[CmdServerPingAckId] = "ServerPingAck"
    msgNames[CmdGateMessageToFSId] = "GateMessageToFS"
    msgNames[CmdFSMessageToGateId] = "FSMessageToGate"
	
    msgLogLv[CmdHandShakeReqId] = 1
    msgLogLv[CmdHandShakeAckId] = 1
    msgLogLv[CmdUserQuitRptId] = 1
    msgLogLv[CmdUserQuitNtfId] = 1
    msgLogLv[CmdRouteMessageId] = 1
    msgLogLv[CmdBroadcastNtfId] = 1
    msgLogLv[CmdUserFightInfoNtfId] = 1
    msgLogLv[CmdBroadcastByFSId] = 1
    msgLogLv[CmdServerPingReqId] = 1
    msgLogLv[CmdServerPingAckId] = 1
    msgLogLv[CmdGateMessageToFSId] = 1
    msgLogLv[CmdFSMessageToGateId] = 1
}

func GetMsgPrototype(key uint16) reflect.Type {
    return msgPrototypes[key]
}

func GetMsgName(key uint16) string {
	return msgNames[key]
}

func GetMsgLogLv(key int) int {
	return msgLogLv[key]
}

const (
    CmdUnknownId uint16 = 0
    CmdHandShakeReqId = 11
    CmdHandShakeAckId = 12
    CmdUserQuitRptId = 15
    CmdUserQuitNtfId = 16
    CmdRouteMessageId = 17
    CmdBroadcastNtfId = 19
    CmdUserFightInfoNtfId = 20
    CmdBroadcastByFSId = 21
    CmdServerPingReqId = 23
    CmdServerPingAckId = 24
    CmdGateMessageToFSId = 25
    CmdFSMessageToGateId = 26
)

func GetCmdIdFromType(i interface{}) uint16 {
	switch i.(type) {
	case *HandShakeReq:
	     return 11
	case *HandShakeAck:
	     return 12
	case *UserQuitRpt:
	     return 15
	case *UserQuitNtf:
	     return 16
	case *RouteMessage:
	     return 17
	case *BroadcastNtf:
	     return 19
	case *UserFightInfoNtf:
	     return 20
	case *BroadcastByFS:
	     return 21
	case *ServerPingReq:
	     return 23
	case *ServerPingAck:
	     return 24
	case *GateMessageToFS:
	     return 25
	case *FSMessageToGate:
	     return 26
	default:
		return 0
	}
}
