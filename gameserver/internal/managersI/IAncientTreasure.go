package managersI

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IAncientTreasureManager interface {

	//激活
	Active(user *objs.User, id int, ack *pb.AncientTreasuresActivateAck, op *ophelper.OpBagHelperDefault) error

	//注灵
	ZhuLin(user *objs.User, id int, ack *pb.AncientTreasuresZhuLinAck, op *ophelper.OpBagHelperDefault) error

	//升星
	UpStar(user *objs.User, id int, ack *pb.AncientTreasuresUpStarAck, op *ophelper.OpBagHelperDefault) error

	//觉醒
	JueXin(user *objs.User, treasureId, chooseIndex int, items []int32, ack *pb.AncientTreasuresJueXingAck, op *ophelper.OpBagHelperDefault) error

	//重置
	Reset(user *objs.User, id int, ack *pb.AncientTreasuresResertAck, op *ophelper.OpBagHelperDefault) error

	//获取
	GetAncientTreasureTaoZ(user *objs.User) map[int]int

	//获取条件属性
	GetConditionProp(user *objs.User, types int, props map[int]int, datas gamedb.IntSlice2) map[int]int

	//获取远古宝物信息
	BuildAncientTreasureInfo(user *objs.User) map[int32]*pb.AncientTreasureInfo

	GetAncientTreasureConditionValue(user *objs.User, ack *pb.AncientTreasuresCondotionInfosAck)
}
