package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

type IAnnouncementManager interface {
	//获取当前 公告
	GetAnnouncement(user *objs.User) []*pb.AnnouncementInfo

	//立刻更新公告
	UpAnnouncementInfos()

	//立刻更新跑马灯
	UpPaoMaDengInfoNow()

	//广播系统消息
	SendSystemChat(user *objs.User, types, itemId, stageId int)

	//战斗掉落物品广播
	FightSendSystemChat(user *objs.User, items map[int]int, stageId, types int)

	//系统消息内容
	BuildContent(user *objs.User, types, itemId, stageId int) string

	BroadCastFsKillInfo(infos string)
}
