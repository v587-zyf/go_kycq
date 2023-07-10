package managersI

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IAuctionManager interface {
	//LoadInfo
	ProcessAuctionInfoReq(user *objs.User, auctionType int, ack *pb.AuctionInfoNtf) error

	//上架物品
	UpItemToAuction(user *objs.User, req *pb.AuctionPutawayItemReq, ack *pb.AuctionPutawayItemNtf, op *ophelper.OpBagHelperDefault) error

	ProcessBidInfo(user *objs.User, auctionId, auctionType int, ack *pb.BidInfoNtf) error

	//竞价
	ProcessBid(user *objs.User, req *pb.BidReq, ack *pb.BidNtf, op *ophelper.OpBagHelperDefault) error

	//定时3秒检查拍卖行
	CheckAuctionItemTask()

	//掉落物品进门派拍卖行
	DropItemToGuildAuction(guildId, itemId, count, dropState int, canGetRedReward model.IntSlice)

	//计算分红数量
	CalculateFenHonNum(guildId, nowBidPrice, canGetRedNum int) int

	//获取自己竞拍的物品信息
	MyBidNtf(user *objs.User, ack *pb.MyBidNtf)

	//我上架的物品信息
	MyPutAwayItemInfo(user *objs.User, ack *pb.MyPutAwayItemInfoAck)

	//我的竞拍信息
	GetUserBidInfos(user *objs.User, ack *pb.MyBidInfoItemAck)

	//拍卖行红点处理
	BroadcastAuctionRedPointNtf(guildId, types, isBright int)

	SendAuctionRedPointNtf(userId, types, isBright int)

	//兑换金锭
	ConversionGoldIngot(user *objs.User, num int, op *ophelper.OpBagHelperDefault, ack *pb.ConversionGoldIngotAck) error
}
