package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IExpStageManager interface {
	OnLine(user *objs.User)
	ResetExpStage(user *objs.User, date int)
	EnterExpStageFight(user *objs.User, stageId int) error
	ExpStageFightResultNtf(user *objs.User, monsterNum, getExp, stageId int)
	Double(user *objs.User, op *ophelper.OpBagHelperDefault, stageId int, ack *pb.ExpStageDoubleAck) error
	ExpStageDareNumNtf(user *objs.User)

	ExpStageBuyNumCheck(user *objs.User) (int, error)
	ExpStageBuyNumNtf(user *objs.User) error
	ExpStageBuyNum(user *objs.User, use bool, op *ophelper.OpBagHelperDefault) error

	ExpStageSweep(user *objs.User, stageId int, op *ophelper.OpBagHelperDefault, ack *pb.ExpStageSweepAck) error
}
