package announcement

import (
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/publicCon/constAnnouncement"
	"cqserver/gameserver/internal/base"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"time"
)

func (this *AnnouncementManager) PaoMaDengLoad() {
	ticker := time.NewTicker(time.Second * 60 * 5)
	for {
		select {
		case <-ticker.C:
			announcements, err := modelCross.GetPaoMaDengModel().GetAllPaoMaDeng()
			if err != nil {
				logger.Error("Load login announcements is error %s", err)
				continue
			}
			this.buildPaoMaDengInfo(announcements)
		}
	}
}

//立刻更新跑马灯
func (this *AnnouncementManager) UpPaoMaDengInfoNow() {

	announcements, err := modelCross.GetPaoMaDengModel().GetAllPaoMaDeng()
	if err != nil {
		logger.Error("Load login announcements is error %s", err)
		return
	}
	this.buildPaoMaDengInfo(announcements)
	return
}

func (this *AnnouncementManager) buildPaoMaDengInfo(infos []modelCross.PaoMaDeng) {

	ntf := &pb.PaoMaDengNtf{}
	channelId := 0
	allUsers := this.GetUserManager().GetAllOnlineUserInfo()
	for _, data := range allUsers {
		channelId = data.ChannelId
		break
	}
	ntf.PaoMaDengInfos = make([]*pb.PaoMaDengInfo, 0)
	for _, info := range infos {
		if info.Types == constAnnouncement.IntervalPlay {
			if int(time.Now().Unix()) < this.playTimesMark[info.Id] {
				continue
			}
			this.playTimesMark[info.Id] = int(time.Now().Unix()) + info.IntervalTimes*60 //下次可弹时间戳
		}

		if info.Types == constAnnouncement.CronPlay {
			if this.playTimesMark[info.Id] >= 1 {
				continue
			}
			this.playTimesMark[info.Id] += 1
		}

		serverMark := make(map[int]bool)
		channelMark := make(map[int]bool)
		for _, serverId := range info.ServerIds {
			serverMark[serverId] = true
		}
		for _, channel := range info.ChannelIds {
			channelMark[channel] = true
		}

		if len(info.ChannelIds) == 0 && len(info.ServerIds) == 0 {
			ntf.PaoMaDengInfos = append(ntf.PaoMaDengInfos, &pb.PaoMaDengInfo{Type: int32(info.Types), CycleTimes: int32(info.CycleTimes), Content: info.Content})
			continue
		}

		if len(info.ChannelIds) == 0 {
			if serverMark[base.Conf.ServerId] {
				ntf.PaoMaDengInfos = append(ntf.PaoMaDengInfos, &pb.PaoMaDengInfo{Type: int32(info.Types), CycleTimes: int32(info.CycleTimes), Content: info.Content})
			}
			continue
		}

		if len(info.ServerIds) == 0 {
			if channelMark[channelId] {
				ntf.PaoMaDengInfos = append(ntf.PaoMaDengInfos, &pb.PaoMaDengInfo{Type: int32(info.Types), CycleTimes: int32(info.CycleTimes), Content: info.Content})
			}
			continue
		}

		for _, channel := range info.ChannelIds {
			if channel == channelId && serverMark[base.Conf.ServerId] {
				ntf.PaoMaDengInfos = append(ntf.PaoMaDengInfos, &pb.PaoMaDengInfo{Type: int32(info.Types), CycleTimes: int32(info.CycleTimes), Content: info.Content})
			}
		}

	}

	this.BroadcastAll(ntf)
	return
}

func (this *AnnouncementManager) BroadCastFsKillInfo(infos string) {
	ntf := &pb.PaoMaDengNtf{}
	ntf.PaoMaDengInfos = make([]*pb.PaoMaDengInfo, 0)
	ntf.PaoMaDengInfos = append(ntf.PaoMaDengInfos, &pb.PaoMaDengInfo{Type: constAnnouncement.ActivityInfo, Content: infos})
	this.BroadcastAll(ntf)
}
