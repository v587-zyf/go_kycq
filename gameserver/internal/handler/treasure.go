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
	pb.Register(pb.CmdSetTreasurePopUpStateReqId, HandleSetTreasurePopUpStateReq)
	pb.Register(pb.CmdChooseTreasureAwardReqId, HandledChooseTreasureAwardReq)
	pb.Register(pb.CmdBuyTreasureItemReqId, HandleBuyTreasureItemReq)
	pb.Register(pb.CmdTreasureApplyGetReqId, HandleTreasureApplyGetReq)
	pb.Register(pb.CmdTreasureInfosReqId, HandleTreasureInfosReq)
	pb.Register(pb.CmdGetTreasureIntegralAwardReqId, GetTreasureIntegralAwardReq)
	pb.Register(pb.CmdTreasureDrawInfoReqId, HandleTreasureDrawInfosReq)
}

//设置弹框提示状态
func HandleSetTreasurePopUpStateReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.SetTreasurePopUpStateReq)
	user := conn.GetSession().(*managers.ClientSession).User

	ack := &pb.SetTreasurePopUpStateAck{}
	if req.State != 0 {
		if req.State != 1 {
			return nil, nil, gamedb.ERRPARAM
		}
	}
	m.Treasure.SetPopUp(user, int(req.State), ack)
	return ack, nil, nil
}

/*选择转盘物品*/
func HandledChooseTreasureAwardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ChooseTreasureAwardReq)
	user := conn.GetSession().(*managers.ClientSession).User

	ack := &pb.ChooseTreasureAwardAck{}
	err := m.Treasure.ChooseTreasureItem(user, int(req.Type), req.Index, int(req.IsReplace), int(req.ReplaceIndex), ack)
	if err != nil {
		return nil, nil, err
	}
	return ack, nil, nil
}

//购买寻龙令
func HandleBuyTreasureItemReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	opHelper := ophelper.NewOpBagHelperDefault(constBag.OpTreasureBuyXunLongLinCost)

	ack := &pb.BuyTreasureItemAck{}
	err := m.Treasure.BuyTreasureItem(user, ack, opHelper)
	if err != nil {
		return nil, nil, err
	}
	return ack, opHelper, nil
}

//开始抽奖
func HandleTreasureApplyGetReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	opHelper := ophelper.NewOpBagHelperDefault(constBag.OpTreasureApplyGet)

	ack := &pb.TreasureApplyGetAck{}
	err := m.Treasure.ApplyGet(user, ack, opHelper)
	if err != nil {
		return nil, nil, err
	}
	return ack, nil, nil
}

//Load
func HandleTreasureInfosReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User

	ack := &pb.TreasureInfosAck{}
	err := m.Treasure.Load(user, ack)
	if err != nil {
		return nil, nil, err
	}
	return ack, nil, nil
}

//积分奖励
func GetTreasureIntegralAwardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User

	req := p.(*pb.GetTreasureIntegralAwardReq)
	opHelper := ophelper.NewOpBagHelperDefault(constBag.OpTreasureIntegralAward)

	ack := &pb.GetTreasureIntegralAwardAck{}
	err := m.Treasure.GetTreasureIntegralAward(user, int(req.Id), ack, opHelper)
	if err != nil {
		return nil, nil, err
	}
	return ack, opHelper, nil
}

func HandleTreasureDrawInfosReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.TreasureDrawInfoAck{}
	m.Treasure.DrawLoad(user, ack)
	return ack, nil, nil
}
