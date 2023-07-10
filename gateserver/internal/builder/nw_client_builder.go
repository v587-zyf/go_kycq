package builder

import "cqserver/protobuf/pb"
import "cqserver/gamelibs/errex"

func BuildErrorAck(ei *errex.ErrorItem) *pb.ErrorAck {
	return &pb.ErrorAck{
		Code:    int32(ei.Code),
		Message: ei.Message,
	}
}
