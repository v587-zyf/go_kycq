package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdPetActiveReqId, HandlerPetActiveReq)
	pb.Register(pb.CmdPetUpLvReqId, HandlerPetUpLvReq)
	pb.Register(pb.CmdPetUpGradeReqId, HandlerPetUpGradeReq)
	pb.Register(pb.CmdPetBreakReqId, HandlerPetBreakReq)
	pb.Register(pb.CmdPetChangeWearReqId, HandlerPetChangeWearReq)

	pb.Register(pb.CmdPetAppendageReqId, HandlerPetAppendageReq)
}

func HandlerPetActiveReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.PetActiveReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypePetActive)

	ack := &pb.PetActiveAck{}
	if err := m.Pet.Active(user, int(req.Id), op, ack); err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}

func HandlerPetUpLvReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.PetUpLvReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypePetUpLv)

	ack := &pb.PetUpLvAck{}
	if err := m.Pet.OneKeyUpLv(user, int(req.Id), int(req.ItemId), op, ack); err != nil {
	//if err := m.Pet.UpLv(user, int(req.Id), int(req.ItemId), int(req.ItemNum), op, ack); err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}

func HandlerPetUpGradeReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.PetUpGradeReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypePetUpGrade)

	ack := &pb.PetUpGradeAck{}
	if err := m.Pet.UpGrade(user, int(req.Id), op, ack); err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}

func HandlerPetBreakReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.PetBreakReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypePetBreak)

	ack := &pb.PetBreakAck{}
	if err := m.Pet.Break(user, int(req.Id), op, ack); err != nil {
		return nil, nil, err
	}
	return ack, op, nil
}

func HandlerPetChangeWearReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.PetChangeWearReq)
	user := conn.GetSession().(*managers.ClientSession).User

	ack := &pb.PetChangeWearAck{}
	if err := m.Pet.ChangeWear(user, int(req.Id), ack); err != nil {
		return nil, nil, err
	}
	return ack, nil, nil
}

func HandlerPetAppendageReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {
	req := p.(*pb.PetAppendageReq)
	user := conn.GetSession().(*managers.ClientSession).User
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypePetAppendage)

	if err := m.Pet.AppendageStrengthen(user, int(req.PetId), op); err != nil {
		return nil, nil, err
	}
	return &pb.PetAppendageAck{PetId: req.PetId, Lv: int32(user.PetAppendage[int(req.PetId)])}, op, nil
}
