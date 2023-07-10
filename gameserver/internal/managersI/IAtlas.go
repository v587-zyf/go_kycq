package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IAtlasManager interface {
	//图鉴激活
	AtlasActive(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.AtlasActiveAck) error
	//图鉴升星
	AtlasUpStar(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.AtlasUpStarAck) error
	//图鉴集合激活
	AtlasGatherActive(user *objs.User, id int, ack *pb.AtlasGatherActiveAck) error
	//图鉴集合升星
	AtlasGatherUpStar(user *objs.User, id int, ack *pb.AtlasGatherUpStarAck) error
	//图鉴穿戴
	Change(user *objs.User, heroIndex, id int, ack *pb.AtlasWearChangeAck) error
	//GM命令
	AtlasGm(user *objs.User, gmT string)
}
