package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdMagicTowerGetUserInfoReqId, HandlerMagicTowerGetUserInfoReq)
	pb.Register(pb.CmdMagicTowerlayerAwardReqId, HandlerMagicTowerlayerAwardReq)
}

func HandlerMagicTowerGetUserInfoReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User

	score, isGetAward, err := m.MagicTower.MagicTowerGetUserInfo(user)
	if err != nil {
		return nil, nil, err
	}
	ack := &pb.MagicTowerGetUserInfoAck{
		Score:      int32(score),
		IsGetAward: int32(isGetAward),
	}

	return ack, nil, nil
}

func HandlerMagicTowerlayerAwardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeMagicLayerAward)
	var err error
	err = m.MagicTower.MagicTowerlayerAward(user, op)
	if err != nil {
		return nil, nil, err
	}
	ack := &pb.MagicTowerlayerAwardAck{
		Goods: op.ToChangeItems(),
	}
	return ack, op, nil
}
