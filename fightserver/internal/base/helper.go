package base

import (
	"cqserver/fightserver/internal/net"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pbserver"
)

func AddItems(user Actor, items map[int32]int32, opType int32) error {
	rpc := &pbserver.FSAddItemReq{
		UserId: int32(user.GetUserId()),
		Items:  items,
		OpType: opType,
	}
	//ack := &pbserver.FSAddItemAck{}
	err := net.GetGsConn().SendMessage(rpc)
	//err := net.GetGsConn().CallGs(uint32(user.HostId()), rpc, ack)
	if err != nil {
		logger.Error(" 推送game添加物品异常,玩家：%v-%v,物品：%v,来源：%v", user.GetUserId(), user.NickName(), items, opType)
		return err
	}
	//if !ack.IsSuccess {
	//	return gamedb.ERRITEMADDFAIL
	//}
	return nil
}
