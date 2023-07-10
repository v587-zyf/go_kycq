package handler

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/managers"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
)

func init() {
	pb.Register(pb.CmdTaskDoneReqId, HandlerTaskDoneReq)
	pb.Register(pb.CmdTaskNpcStateReqId, HandlerTaskNpcReq)
	pb.Register(pb.CmdSetTaskInfoReqId, HandlerSetTaskInfoReq)
}

func HandlerTaskDoneReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	user := conn.GetSession().(*managers.ClientSession).User
	//req := p.(*pb.TaskDoneReq)
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeTask)
	err := m.Task.TaskDone(user, op)
	if err != nil {
		return nil, nil, err
	}
	return &pb.TaskDoneAck{
		Goods: op.ToChangeItems(),
	}, op, nil
}

func HandlerTaskNpcReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	user := conn.GetSession().(*managers.ClientSession).User

	taskConf := gamedb.GetTaskConditionCfg(user.MainLineTask.TaskId)
	if taskConf.ConditionType == pb.CONDITION_NPC_CHAT {
		if user.StageId == taskConf.ConditionValue[0] || user.FightStageId == taskConf.ConditionValue[0] {
			user.MainLineTask.Process = 1
		}
	}

	if taskConf.ConditionType == pb.CONDITION_KILL_UNKNOWN_BOSS {
		user.MainLineTask.Process = 1
	}
	user.Dirty = true
	return &pb.TaskInfoNtf{
		TaskId:      int32(user.MainLineTask.TaskId),
		Process:     int32(user.MainLineTask.Process),
		MarkProcess: int32(user.MainLineTask.MarkProcess),
	}, nil, nil
}

func HandlerSetTaskInfoReq(conn nw.Conn, p interface{}) (nw.ProtoMessage, pb.OpGoodsHelper, error) {

	user := conn.GetSession().(*managers.ClientSession).User
	req := p.(*pb.SetTaskInfoReq)

	if user.MainLineTask.TaskId == int(req.TaskId) {
		user.MainLineTask.MarkProcess = int(req.Process)
	}
	user.Dirty = true
	return &pb.SetTaskInfoAck{
		TaskId:  int32(user.MainLineTask.TaskId),
		Process: int32(user.MainLineTask.MarkProcess),
	}, nil, nil
}
