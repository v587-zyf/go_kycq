package PreviewFunction

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

func NewPreviewFunctionManager(m managersI.IModule) *PreviewFunctionManager {
	return &PreviewFunctionManager{
		IModule: m,
	}
}

type PreviewFunctionManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *PreviewFunctionManager) Load(user *objs.User, ack *pb.PreviewFunctionLoadAck) {
	ack.HaveBuyIds, ack.HavePointIds = this.buildHaveBugIds(user)
	return
}

func (this *PreviewFunctionManager) GetReward(user *objs.User, id int, op *ophelper.OpBagHelperDefault, ack *pb.PreviewFunctionGetAck) error {

	if id <= 0 {
		return gamedb.ERRPARAM
	}

	for _, ids := range user.PreviewFunction {
		if ids == id {
			return gamedb.ERRREPEATBUY
		}
	}

	cfg := gamedb.GetPreFunctionPreFunctionCfg(id)
	if cfg == nil {
		logger.Error("GetPreFunctionPreFunctionCfg nil id:%v", id)
		return gamedb.ERRPARAM
	}

	funCfg := gamedb.GetFunctionFunctionCfg(cfg.Condition)
	if funCfg == nil {
		logger.Error("GetFunctionFunctionCfg  nil cfg.Condition:%v", cfg.Condition)
		return gamedb.ERRPARAM
	}

	if check := this.GetCondition().CheckMulti(user, -1, funCfg.Condition); !check {
		return gamedb.ERRCONDITION
	}
	if cfg.Price.ItemId > 0 && cfg.Price.Count > 0 {
		enough, _ := this.GetBag().HasEnough(user, cfg.Price.ItemId, cfg.Price.Count)
		if !enough {
			return gamedb.ERRNOTENOUGHGOODS
		}
		_ = this.GetBag().Remove(user, op, cfg.Price.ItemId, cfg.Price.Count)
	}

	this.GetBag().AddItems(user, cfg.Item, op)
	user.PreviewFunction = append(user.PreviewFunction, id)
	user.Dirty = true
	ack.HaveBuyIds, ack.HavePointIds = this.buildHaveBugIds(user)
	ack.Id = int32(id)
	return nil
}

func (this *PreviewFunctionManager) SetPointId(user *objs.User, pointId int, ack *pb.PreviewFunctionPointAck) error {

	cfg := gamedb.GetPreFunctionPreFunctionCfg(pointId)
	if cfg == nil {
		return gamedb.ERRPARAM
	}

	user.PreviewFunctionPoint = append(user.PreviewFunctionPoint, pointId)
	_, ack.HavePointIds = this.buildHaveBugIds(user)
	return nil
}

func (this *PreviewFunctionManager) buildHaveBugIds(user *objs.User) ([]int32, []int32) {
	haveBuyIds := make([]int32, 0)
	for _, id := range user.PreviewFunction {
		haveBuyIds = append(haveBuyIds, int32(id))
	}
	havePointIds := make([]int32, 0)
	for _, id := range user.PreviewFunctionPoint {
		havePointIds = append(havePointIds, int32(id))
	}
	return haveBuyIds, havePointIds
}
