package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdPrivilegeBuyReqId, HandlerPrivilegeBuyReq)
}

func HandlerPrivilegeBuyReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.PrivilegeBuyReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypePrivilegeBuy)
	if err := m.Privilege.Buy(user, int(req.GetPrivilegeId()), op); err != nil {
		return nil, nil, err
	}
	return &pb.PrivilegeBuyAck{PrivilegeId: req.GetPrivilegeId(), Goods: op.ToChangeItems()}, op, nil
}