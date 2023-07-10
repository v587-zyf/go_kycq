package manager

import (
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/publicCon/constAnnouncement"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"sync"
	"time"
)

type AnnouncementManager struct {
	util.DefaultModule
	announcements []modelCross.Announcement
	mutex         sync.RWMutex
}

func NewAnnouncementManager() *AnnouncementManager {
	return &AnnouncementManager{
		announcements: make([]modelCross.Announcement, 0),
	}
}

func (this *AnnouncementManager) Init() error {
	this.UpAnnouncementInfos()
	go this.Load()
	return nil
}

func (this *AnnouncementManager) Load() {
	ticker := time.NewTicker(time.Second * 60 * 5)
	for {
		select {
		case <-ticker.C:
			announcements, err := modelCross.GetAnnouncementModel().GetAllLoginAnnouncement(constAnnouncement.Login)
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

func (this *AnnouncementManager) UpAnnouncementInfos() {
	announcements, err := modelCross.GetAnnouncementModel().GetAllLoginAnnouncement(constAnnouncement.Login)
	logger.Info("UpAnnouncementInfos announcements:%v", announcements)
	if err != nil {
		logger.Error("Load login announcements is error %s", err)
		return
	}
	this.mutex.Lock()
	this.announcements = announcements
	this.mutex.Unlock()
}

func (this *AnnouncementManager) GetAnnouncementInfos() []*AnnouncementInfo {
	this.mutex.RLock()
	datas := make([]*AnnouncementInfo, 0)

	for _, announcement := range this.announcements {
		datas = append(datas, &AnnouncementInfo{Id: announcement.Id, Title: announcement.Title, Announcement: announcement.Announcement})
	}
	this.mutex.RUnlock()
	return datas
}
