package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
	"time"
)

func BuildTitleList(user *objs.User, hasExpire bool) []*pb.Title {
	userTitle := user.Title
	timeNow := int(time.Now().Unix())
	pbSlice := make([]*pb.Title, 0)
	for id, info := range userTitle {
		if !hasExpire && info.EndTime != -1 && info.EndTime < timeNow {
			continue
		}
		pbSlice = append(pbSlice, BuildTitle(id, info))
	}
	return pbSlice
}

func BuildTitle(titleId int, unit *model.TitleUnit) *pb.Title {
	return &pb.Title{
		TitleId:   int32(titleId),
		StartTime: int64(unit.StartTime),
		EndTime:   int64(unit.EndTime),
		IsLook:    unit.IsLook,
	}
}
