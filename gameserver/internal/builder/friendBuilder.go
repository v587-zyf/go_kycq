package builder

import (
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/common"
	"cqserver/protobuf/pb"
)

func BuildFriendMsgLog(msgLog model.MsgLogs, notShowTime int) []*pb.MsgLog {
	pbMsgLog := make([]*pb.MsgLog, 0)
	if msgLog != nil {
		for _, log := range msgLog {
			if log.Time <= notShowTime {
				continue
			}
			pbMsgLog = append(pbMsgLog, BuildFriendMsg(log))
		}
	}
	return pbMsgLog
}

func BuildFriendMsg(log *model.MsgLog) *pb.MsgLog {
	msg := common.UnicodeEmojiDecode(log.Msg)
	return &pb.MsgLog{
		Msg:  msg,
		Time: int64(log.Time),
		IsMy: log.IsMy,
	}
}

func BuildFriendHeroEquip(equips model.Equips) map[int32]*pb.EquipUnit {
	pbMap := make(map[int32]*pb.EquipUnit)
	for pos, equip := range equips {
		pbMap[int32(pos)] = BuildPbEquipUnit(equip)
	}
	return pbMap
}

func BuildFriendHeroPro(prop map[int]int) map[int32]int64 {
	pbMap := make(map[int32]int64)
	for pid, pVal := range prop {
		pbMap[int32(pid)] = int64(pVal)
	}
	return pbMap
}

func BuildHeroDisplay(display *model.Display) *pb.Display {
	return &pb.Display{
		ClothItemId:     int32(display.ClothItemId),
		ClothType:       int32(display.ClothType),
		WeaponItemId:    int32(display.WeaponItemId),
		WeaponType:      int32(display.WeaponType),
		WingId:          int32(display.WingId),
		MagicCircleLvId: int32(display.MagicCircleLvId),
		TitleId:         int32(display.TitleId),
		LabelId:         int32(display.LabelId),
		LabelJob:        int32(display.LabelJob),
	}
}

func BuildIsFriendApply(user *objs.User) bool {
	b := false
	if len(rmodel.Friend.GetFriendApply(user.Id)) > 0 {
		b = true
	}
	return b
}
