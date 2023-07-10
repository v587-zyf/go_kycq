package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdOfficialUpLevelReqId, HandlerOfficialUpLevelReq)
}

func HandlerOfficialUpLevelReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	user := conn.GetSession().(*managers.ClientSession).User

	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeOfficialUpLv)
	err := m.Official.OfficialUpLevel(user, op)
	if err != nil {
		return nil, nil, err
	}
	ack := &pb.OfficialUpLevelAck{
		NewLv: int32(user.Official),
	}
	return ack, op, nil
}
