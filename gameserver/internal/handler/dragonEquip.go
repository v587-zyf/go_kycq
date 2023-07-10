package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdDragonEquipUpLvReqId, HandlerDragonEquipUpLvReq)
}

func HandlerDragonEquipUpLvReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.DragonEquipUpLvReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeDragonEquip)

	var err error
	ack := &pb.DragonEquipUpLvAck{}

	err = m.DragonEquip.UpLv(user, int(req.HeroIndex), int(req.Id), op, ack)
	if err != nil {
		return nil, nil, err
	}
	logger.Debug("HandlerDragonEquipUpLvReq ack is %v", ack)

	return ack, op, nil
}
