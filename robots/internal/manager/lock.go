package manager

import (
	"cqserver/protobuf/pb"
	"flag"
)

var (
	getList = flag.Bool("getList", false, "specify config file")
)

func (this *Robot) lockTest() {
	//this.SendMessage(0, this.MakeMsg(pb.CmdGuildLoadInfoReqId, ""))
	if *getList {
		//拉取挖矿列表
		//this.SendMessage(0, this.MakeMsg(pb.CmdMiningListReqId, ""))
		//获取公会
		//this.SendMessage(0, this.MakeMsg(pb.CmdAllGuildInfosReqId, ""))
		//拉取世界首领信息
		//this.SendMessage(0, this.MakeMsg(pb.CmdLoadWorldLeaderReqId, ""))
		//拉取拍卖行
		//this.SendMessage(0, this.MakeMsg(pb.CmdAuctionInfoReqId, fmt.Sprintf(`%d`, constAuction.WorldAuction)))
		//抽卡

	} else {
		//写入挖矿
		//this.SendMessage(0, this.MakeMsg(pb.CmdMiningStartReqId, ""))
		//加入公会
		//this.SendMessage(0, this.MakeMsg(pb.CmdApplyJoinGuildReqId, "5000"))
		//写入世界首领
		//this.SendMessage(0, this.MakeMsg(pb.CmdWorldLeaderEnterReqId, "71001"))
		//写入拍卖行
		//this.SendMessage(0, this.MakeMsg(pb.CmdBidReqId, fmt.Sprintf(`%d,%d,%d`, constAuction.WorldAuction, 7, 0)))
	}
	for i := 0; i < 100; i++ {
		this.SendMessage(0, this.MakeMsg(pb.CmdCardActivityApplyGetReqId, "10"))
	}

}
