package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IAncientBossManager interface {
	Online(user *objs.User)
	ResetAncientBoss(user *objs.User, date int)
	ResetAncientBossOwner()
	//加载首领列表
	Load(user *objs.User, area int, ack *pb.AncientBossLoadAck) error
	//购买战斗次数
	BuyNum(user *objs.User, use bool, buyNum int, op *ophelper.OpBagHelperDefault) error
	//进入战斗
	EnterAncientBossFight(user *objs.User, stageId int) error
	//战斗结果结算
	AncientBossFightResult(user *objs.User, winUserId, stageId int, items map[int]int)
	//获取最近归属列表
	GetOwnerList(user *objs.User, stageId int) []*pb.AncientBossOwnerInfo
	//推送怪物数据
	SendBossInfo(ancientBossInfo *pb.AncientBossNtf)
}
