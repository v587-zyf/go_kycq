package builder

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

func BuildLabel(user *objs.User) *pb.Label {
	userLabel := user.Label
	return &pb.Label{
		LabelId:   int32(userLabel.Id),
		Job:       int32(userLabel.Job),
		Transfer:  int32(userLabel.Transfer),
		DayReward: userLabel.DayReward,
	}
}
