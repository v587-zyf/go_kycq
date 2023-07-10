package managersI

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IMaterialStage interface {
	OnLine(user *objs.User)
	ResetMaterialStage(user *objs.User, date int)
	EnterMaterialStageFight(user *objs.User, stageId int) error
	MaterialStageFightResultNtf(user *objs.User, isWin bool, stageId int) error
	MaterialStageSweep(user *objs.User, materialType int, op *ophelper.OpBagHelperDefault, ack *pb.MaterialStageSweepAck) error

	BuyNum(user *objs.User, mateType int) error
	BuyNumCheck(user *objs.User, mateType int) (int, gamedb.ItemInfos, error)
	MaterialBuyNum(user *objs.User, mateType int, use bool, op *ophelper.OpBagHelperDefault) error
}
