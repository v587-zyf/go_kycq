package handler

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdDaBaoEquipUpReqId, HandlerDaBaoEquipUpReq)
	pb.Register(pb.CmdEnterDaBaoMysteryReqId, HandlerDaBaoMysteryReq)
	pb.Register(pb.CmdDaBaoMysteryEnergyItemBuyReqId, HandlerDaBaoMysteryEnergyItemBuyReq)
	pb.Register(pb.CmdDaBaoMysteryEnergyAddReqId, HandlerDaBaoMysteryEnergyAddReq)
}

func HandlerDaBaoEquipUpReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.DaBaoEquipUpReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeDaBaoEquipUp)
	equipT := int(req.GetEquipType())
	if err := m.DaBao.UpEquip(user, equipT, op); err != nil {
		return nil, nil, err
	}
	m.UserManager.SendItemChangeNtf(user, op)
	return &pb.DaBaoEquipUpAck{EquipType: req.GetEquipType(), Lv: int32(user.DaBaoEquip[equipT])}, nil, nil
}

func HandlerDaBaoMysteryReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.EnterDaBaoMysteryReq)
	user := conn.GetSession().(*managers.ClientSession).User
	return nil, nil, m.DaBao.EnterMystery(user, int(req.GetStageId()))
}

func HandlerDaBaoMysteryEnergyItemBuyReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.DaBaoMysteryEnergyItemBuyReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeDaBaoMysteryEnergyItemBuy)

	shopId := int(req.GetShopId())
	if shopCfg := gamedb.GetShopTypeCfg(shopId); shopCfg != nil && shopCfg.Item.ItemId != gamedb.GetConf().DaBaoMysteryEnergyItem.ItemId {
		return nil, nil, gamedb.ERRPARAM
	}
	ack := &pb.ShopBuyAck{}
	if err := m.Shop.Buy(user, shopId, 1, op, ack); err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}

func HandlerDaBaoMysteryEnergyAddReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.DaBaoMysteryEnergyAddReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeDaBaoMysteryEnergyAdd)
	if err := m.DaBao.EnergyItemUse(user, int(req.GetItemId()), op); err != nil {
		return nil, nil, err
	}
	return &pb.DaBaoMysteryEnergyAddAck{ItemId: req.GetItemId()}, op, nil
}
