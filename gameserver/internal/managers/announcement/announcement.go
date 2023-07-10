package announcement

import (
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/publicCon/constAnnouncement"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"sync"
	"time"
)

type AnnouncementManager struct {
	util.DefaultModule
	managersI.IModule
	announcements []modelCross.Announcement
	playTimesMark map[int]int
	mutex         sync.RWMutex
}

func NewAnnouncementManager(m managersI.IModule) *AnnouncementManager {
	return &AnnouncementManager{
		IModule:       m,
		announcements: make([]modelCross.Announcement, 0),
		playTimesMark: make(map[int]int),
	}
}

func (this *AnnouncementManager) Init() error {
	this.UpAnnouncementInfos()
	go this.Load()
	go this.PaoMaDengLoad()
	return nil
}

func (this *AnnouncementManager) Load() {
	ticker := time.NewTicker(time.Second * 60 * 5)
	for {
		select {
		case <-ticker.C:
			announcements, err := modelCross.GetAnnouncementModel().GetAllLoginAnnouncement(constAnnouncement.Enter)
			if err != nil {
				logger.Error("Load login announcements is error %s", err)
				continue
			}
			this.mutex.Lock()
			this.announcements = announcements
			this.mutex.Unlock()
		}
	}
}

//强制更新 公告
func (this *AnnouncementManager) UpAnnouncementInfos() {
	announcements, err := modelCross.GetAnnouncementModel().GetAllLoginAnnouncement(constAnnouncement.Enter)
	logger.Info("UpAnnouncementInfos announcements:%v", announcements)
	if err != nil {
		logger.Error("Load login announcements is error %s", err)
		return
	}
	this.mutex.Lock()
	this.announcements = announcements
	this.mutex.Unlock()
}

func (this *AnnouncementManager) GetAnnouncement(user *objs.User) []*pb.AnnouncementInfo {
	announcementData := make([]*pb.AnnouncementInfo, 0)
	for _, info := range this.announcements {
		serverMark := make(map[int]bool)
		channelMark := make(map[int]bool)
		for _, serverId := range info.ServerIds {
			serverMark[serverId] = true
		}
		for _, channel := range info.ChannelIds {
			channelMark[channel] = true
		}

		if len(info.ChannelIds) == 0 && len(info.ServerIds) == 0 {
			announcementData = append(announcementData, &pb.AnnouncementInfo{Announcement: info.Announcement, Title: info.Title, Id: int32(info.Id)})
			continue
		}

		if len(info.ChannelIds) == 0 {
			if serverMark[user.ServerId] {
				announcementData = append(announcementData, &pb.AnnouncementInfo{Announcement: info.Announcement, Title: info.Title, Id: int32(info.Id)})
			}
			continue
		}

		if len(info.ServerIds) == 0 {
			if channelMark[user.ChannelId] {
				announcementData = append(announcementData, &pb.AnnouncementInfo{Announcement: info.Announcement, Title: info.Title, Id: int32(info.Id)})
			}
			continue
		}

		for _, channel := range info.ChannelIds {
			if channel == user.ChannelId && serverMark[user.ServerId] {
				announcementData = append(announcementData, &pb.AnnouncementInfo{Announcement: info.Announcement, Title: info.Title, Id: int32(info.Id)})
			}
		}
	}
	if len(announcementData) > 0 {
		return announcementData[:1]
	}
	return announcementData
}
