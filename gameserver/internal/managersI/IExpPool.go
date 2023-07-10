package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IExpPoolManager interface {
	Load(user *objs.User, ack *pb.ExpPoolLoadAck) error

	Upgrade(user *objs.User, op *ophelper.OpBagHelperDefault, ack *pb.ExpPoolUpGradeAck, index, times int) error

	//获取当前开服天数世界等级  and  最大存储上限
	GetExpWorldLvlAndLvLimit(user *objs.User) (int, int, int, float64)

	GetHeroMaxLv(user *objs.User) int

	AutoUpExpLv(user *objs.User)

	ExpPillUse(user *objs.User, itemId int, itemNum int, op *ophelper.OpBagHelperDefault) error
}
