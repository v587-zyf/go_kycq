package builder

import (
	"time"

	"cqserver/protobuf/pb"
)

func BuildEventNtfWithSourceId(eventId int, args []string, sourceId int) *pb.EventNtf {
	return &pb.EventNtf{Id: int32(eventId), Ts: int32(time.Now().Unix()), Args: args, SourceId: int32(sourceId)}
}

func BuildEventNtf(eventId int, args []string) *pb.EventNtf {
	return BuildEventNtfWithSourceId(eventId, args, 0)
}
