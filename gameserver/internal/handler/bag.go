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
	pb.Register(pb.CmdBagInfoReqId, HandleagInfoReq)
	pb.Register(pb.CmdBagSpaceAddReqId, HandlerBagSpaceAddReq)
	pb.Register(pb.CmdBagSortReqId, HandlerBagSortReq)
	pb.Register(pb.CmdEquipLockReqId, HandlerEquipLockReq)
	pb.Register(pb.CmdEquipRecoverReqId, HandlerEquipRecoverReq) //回收
	pb.Register(pb.CmdItemUseReqId, HandlerItemUseReq)
	pb.Register(pb.CmdEquipDestroyReqId, HandlerEquipDestroyReq) //销毁
	//------仓库相关------
	pb.Register(pb.CmdWarehouseInfoReqId, HandlerWarehouseInfoReq)
	pb.Register(pb.CmdWareHouseSpaceAddReqId, HandlerWareHouseSpaceAddReq)
	pb.Register(pb.CmdWarehouseAddReqId, HandlerWarehouseAddReq)
	pb.Register(pb.CmdWarehouseShiftOutReqId, HandlerWarehouseShiftOutReq)
	pb.Register(pb.CmdWarehouseSortReqId, HandlerWarehouseBagSortReq)
}

func HandleagInfoReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	user := conn.GetSession().(*managers.ClientSession).User
	ack := builder.BuildBagInfoAck(user)

	return ack, nil, nil
}

//扩充背包
func HandlerBagSpaceAddReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeBagSpaceAdd)
	err := m.Bag.BagSpaceAdd(user, op)
	if err != nil {
		return nil, nil, err
	}
	ack := &pb.BagSpaceAddAck{
		BagMax: int32(user.BagInfo[constBag.BAG_TYPE_COMMON].MaxNum),
	}
	return ack, op, nil
}

func HandlerBagSortReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	user := conn.GetSession().(*managers.ClientSession).User
	err := m.Bag.BagSort(user, false)
	if err != nil {
		return nil, nil, err
	}
	ack := &pb.BagSortAck{}
	for _, v := range user.Bag {
		if v != nil {
			if v.ItemId > 0 {
				ack.Items = append(ack.Items, builder.BuildPbItem(user, v))
			}
		}
	}

	return ack, nil, nil

}

func HandlerEquipLockReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	req := p.(*pb.EquipLockReq)
	user := conn.GetSession().(*managers.ClientSession).User
	err := m.Bag.EquipLock(user, int(req.Pos))
	if err != nil {
		return nil, nil, err
	}
	return &pb.EquipLockAck{
		Pos:  req.Pos,
		Lock: user.EquipBag[user.Bag[constBag.BAG_TYPE_COMMON].EquipIndex].IsLock,
	}, nil, nil

}

func HandlerEquipRecoverReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.EquipRecoverReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeEquipRecover)
	err := m.Bag.EquipRecover(user, op, req.Positions)
	if err != nil {
		return nil, nil, err
	}
	return &pb.EquipRecoverAck{Goods: op.ToChangeItems()}, op, nil
}

func HandlerItemUseReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.ItemUseReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeItemUse)
	err := m.Bag.ItemUse(user, int(req.HeroIndex), int(req.ItemId), int(req.ItemNum), op)
	if err != nil {
		return nil, nil, err
	}
	goods := op.ToChangeItems()
	//以为回城石 随机石不再是道具，实际不扣除，这里强写
	if req.ItemId == pb.ITEMID_RANDOM_STONE || req.ItemId == pb.ITEMID_BACK_CITY {
		goods.Items = append(goods.Items, &pb.ItemUnit{ItemId: req.ItemId, Count: int64(req.ItemNum)})
	}
	return &pb.ItemUseAck{
		Goods: goods,
	}, op, nil
}

//--------------------仓库相关
func HandlerWarehouseInfoReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	user := conn.GetSession().(*managers.ClientSession).User
	ack := builder.BuildWarehouseBagInfoAck(user)

	return ack, nil, nil
}

func HandlerWareHouseSpaceAddReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeWareHouseSpaceAdd)
	err := m.Bag.WareHouseBagSpaceAdd(user, op)
	if err != nil {
		return nil, nil, err
	}
	return &pb.WareHouseSpaceAddAck{BagMax: int32(user.WarehouseBagInfo[constBag.WAREHOUSE_BAG_TYPE_COMMON].MaxNum)}, op, nil
}

//移动背包物品到仓库
func HandlerWarehouseAddReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.WarehouseAddReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeEquipMoveToWear)
	err := m.Bag.WareHouseBagAdd(user, op, req.Positions)
	if err != nil {
		return nil, nil, err
	}
	return &pb.WarehouseAddAck{Items: builder.BuildWarehouseBagInfoAck(user).Items}, op, nil
}

//移动仓库物品到背包
func HandlerWarehouseShiftOutReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.WarehouseShiftOutReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeEquipMoveToBag)
	err := m.Bag.WareHouseMoveToBag(user, op, req.Positions)
	if err != nil {
		return nil, nil, err
	}
	return &pb.WarehouseShiftOutAck{Items: builder.BuildWarehouseBagInfoAck(user).Items, Goods: op.ToChangeItems()}, op, nil
}

func HandlerWarehouseBagSortReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	user := conn.GetSession().(*managers.ClientSession).User
	err := m.Bag.BagSort(user, true)
	if err != nil {
		return nil, nil, err
	}
	ack := &pb.WarehouseSortAck{}
	for _, v := range user.WarehouseBag {
		if v != nil {
			if v.ItemId > 0 {
				ack.Items = append(ack.Items, builder.BuildPbItem(user, v))
			}
		}
	}
	return ack, nil, nil

}

//装备销毁
func HandlerEquipDestroyReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	req := p.(*pb.EquipDestroyReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeEquipDestroy)
	err := m.Bag.EquipDestroy(user, op, int(req.Positions), int(req.Count))
	if err != nil {
		return nil, nil, err
	}
	return &pb.EquipDestroyAck{}, op, nil
}
