package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdJuexueUpLevelReqId, HandlerJuexueUpLevelReq) // 绝学升级
}

func HandlerJuexueUpLevelReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.JuexueUpLevelReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeJuexueUpLv)

	id := int(req.Id)

	err := m.Juexue.JuexueUpLevel(user, id, op)
	if err != nil {
		return nil, nil, err
	}
	ack := &pb.JuexueUpLevelAck{
		Juexue: builder.BuilderJuexue(user.Juexues[id]),
	}
	return ack, op, nil
}
