package builder

import (
	"cqserver/gamelibs/errex"
	"cqserver/protobuf/pb"
)

func BuildError(errorItem *errex.ErrorItem) *pb.ErrorAck {
	return &pb.ErrorAck{Code: int32(errorItem.Code), Message: errorItem.Message}
}

//func BuildTaskInfo(task model.MainLineTask, IsDetail bool) *pb.TaskInfoNtf {
//	taskT := gamedb.GetDb().GetTask(task.TaskId)
//	if taskT == nil {
//		return &pb.TaskInfoNtf{}
//	}
//	taskInfo := &pb.TaskInfoNtf{TaskId: int32(task.TaskId), Value: int32(task.Process), Status: int32(task.Award), Guided: task.Guided}
//	if IsDetail {
//		//taskInfo.Detail = taskT.ToProtoBuf()
//	}
//	return taskInfo
//}



//func BuildDisplayNtf(user *objs.User) *pb.DisplayNtf {
//	c, w, wing, horse := user.GetDisplayIdsArtifact()
//	return &pb.DisplayNtf{ClothDisplayId: int32(c), WeaponDisplayId: int32(w), WingDisplayId: int32(wing), HorseDisplayId: int32(horse)}
//}



