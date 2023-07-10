package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdFashionUpLevelReqId, HandlerFashionUpLevelReq)
	pb.Register(pb.CmdFashionWearReqId, HandlerFashionWearReq)
}

//获取指定类型排行榜
func HandlerFashionUpLevelReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	req := p.(*pb.FashionUpLevelReq)
	user := conn.GetSession().(*managers.ClientSession).User

	heroIndex := int(req.HeroIndex)
	fashionId := int(req.FashionId)
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeFashionUpLv)
	err := m.Fashion.FashionUpLevel(user, op, heroIndex, fashionId)
	if err != nil {
		return nil, nil, err
	}
	ack := &pb.FashionUpLevelAck{
		HeroIndex: req.HeroIndex,
		Fashion: &pb.Fashion{
			Id:    req.FashionId,
			Level: int32(user.Heros[heroIndex].Fashions[fashionId].Lv),
		},
	}
	return ack, op, nil
}

func HandlerFashionWearReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	req := p.(*pb.FashionWearReq)
	user := conn.GetSession().(*managers.ClientSession).User
	err := m.Fashion.FashionWear(user, int(req.HeroIndex), int(req.FashionId), req.IsWear)
	if err != nil {
		return nil, nil, err
	}
	ack := &pb.FashionWearAck{
		HeroIndex:     req.HeroIndex,
		WearFashionId: req.FashionId,
		IsWear:        req.IsWear,
	}
	return ack, nil, nil
}
