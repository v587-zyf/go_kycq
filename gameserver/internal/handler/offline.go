package handler

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdOfflineAwardLoadReqId, OfflineTimeLoadReq)
	pb.Register(pb.CmdOfflineAwardGetReqId, OfflineAwardGetReq)
}

func OfflineTimeLoadReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	ack := &pb.OfflineAwardLoadAck{0,0}
	//user := conn.GetSession().(*managers.ClientSession).User
	//if user.OfflineTime.Unix() <= 0 {
	//	return ack, nil, nil
	//}
	//offlineTime := time.Now().Unix() - user.OfflineTime.Unix()
	//ack.OfflineTimes = offlineTime
	//if offlineTime < 60 {
	//	ack.GetExpNum = 0
	//	return ack, nil, nil
	//}
	//minutes := offlineTime / 60
	//stageCfg := gamedb.GetHookMapHookMapCfg(user.StageId)
	//
	//count := minutes * int64(stageCfg.Name)
	//monthCardPrivilege := m.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_HANGUP_EXP)
	//if monthCardPrivilege != 0 {
	//	count = int64(common.CalcTenThousand(monthCardPrivilege, int(count)))
	//}
	//ack.GetExpNum = count
	return ack, nil, nil
}

func OfflineAwardGetReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	//user := conn.GetSession().(*managers.ClientSession).User
	ack := &pb.OfflineAwardGetAck{}
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeOfflineReward)
	//err := m.Offline.GetAward(user, ack, op)
	//if err != nil {
	//	return nil, nil, err
	//}
	return ack, op, nil
}
