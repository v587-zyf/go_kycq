package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"time"
)

func init() {
	pb.Register(pb.CmdGetOnlineAwardInfoReqId, HandlerGetOnlineAwardInfoReq)
	pb.Register(pb.CmdGetOnlineAwardReqId, HandlerGetOnlineAwardReq)
}

//获取在线奖励信息
func HandlerGetOnlineAwardInfoReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	user := conn.GetSession().(*managers.ClientSession).User
	onlineTime := user.OnlineAward.OnlineTime + int(time.Now().Sub(user.LastUpdateTime).Seconds())

	ack := &pb.GetOnlineAwardInfoAck{
		OnlineTime: int32(onlineTime),
		GetAwardId: common.ConvertIntSlice2Int32Slice(user.OnlineAward.GetAwardIds),
	}
	return ack, nil, nil
}

//领取在线奖励
func HandlerGetOnlineAwardReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	req := p.(*pb.GetOnlineAwardReq)
	user := conn.GetSession().(*managers.ClientSession).User

	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeOnlineAward)

	err := m.Online.GetOnlineAward(user, int(req.AwardId), op)
	if err != nil {
		return nil, nil, err
	}

	ack := &pb.GetOnlineAwardAck{
		GetAwardId: common.ConvertIntSlice2Int32Slice(user.OnlineAward.GetAwardIds),
	}

	return ack, op, nil
}
