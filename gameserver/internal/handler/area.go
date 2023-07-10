package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdAreaUpLvReqId, HandlerAreaUpLvReq)
}

func HandlerAreaUpLvReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.AreaUpLvReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeArea)

	err := m.Area.UpLv(user, int(req.GetHeroIndex()), int(req.GetId()), op)
	if err != nil {
		return nil, nil, err
	}

	return &pb.AreaUpLvAck{
		HeroIndex: req.HeroIndex,
		Id:        req.Id,
		Lv:        int32(user.Heros[int(req.HeroIndex)].Area[int(req.Id)]),
	}, op, nil
}
