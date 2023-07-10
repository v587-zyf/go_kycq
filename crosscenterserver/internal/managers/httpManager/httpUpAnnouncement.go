package httpManager

import (
	"cqserver/gamelibs/modelCross"
	"cqserver/gamelibs/ptsdk"
	"cqserver/gamelibs/publicCon/constAnnouncement"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pbserver"
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

var announcementMu sync.Mutex

/**
*  @Description:更新公告
*  @param w
*  @param r
**/
func httpDelAnnouncement(w http.ResponseWriter, r *http.Request) {
	defer func() {
		announcementMu.Unlock()
	}()
	announcementMu.Lock()

	announcementGmId := ptsdk.GetSdk().DelAnnouncement(w, r)
	if announcementGmId <= 0 {
		return
	}

	data := modelCross.GetAnnouncementModel().GetAnnouncementByGmId(announcementGmId)
	if data != nil {
		err := modelCross.GetAnnouncementModel().Del(announcementGmId)
		if err != nil {
			logger.Error("httpDelAnnouncement  del 公告 err:%v   data:%v", err, announcementGmId)
			ptsdk.GetSdk().HttpWriteReturnInfo(w, 400, "数据库操作失败", nil)
			return
		}
		m.GetGsServers().SendAllServerMessage(&pbserver.UpAnnouncementNowReq{})
	} else {
		err1 := modelCross.GetPaoMaDengModel().Del(announcementGmId)
		if err1 != nil {
			logger.Error("httpDelAnnouncement  del 跑马灯 err:%v   data:%v", err1, announcementGmId)
			ptsdk.GetSdk().HttpWriteReturnInfo(w, 400, "数据库操作失败", nil)
			return
		}
		m.GetGsServers().SendAllServerMessage(&pbserver.UpPaoMaDengNowReq{})
	}

	keys := common.NewStringSlice(modelCross.CLIENT_TO_LOGIN, "_")[0]
	portInfo, err := modelCross.GetServerPortInfoModel().GetLoginServerPortInfos(keys)
	if err == nil && portInfo != nil && len(portInfo) > 0 {
		for _, v := range portInfo {
			r, err := http.PostForm(fmt.Sprintf("http://%s:%d/upAnnouncement", v.Host, v.Port), url.Values{})
			logger.Info("通知登录更新公告：%v", r, err)
		}
	}

	ptsdk.GetSdk().HttpWriteReturnInfo(w, 200, "success", nil)
}

/**
*  @Description:设置公告
*  @param w
*  @param r
**/
func httpApplyAnnouncement(w http.ResponseWriter, r *http.Request) {

	defer func() {
		announcementMu.Unlock()
	}()
	announcementMu.Lock()

	announcement, paomadeng := ptsdk.GetSdk().ApplyAnnouncement(w, r)
	if announcement == nil && paomadeng == nil {
		return
	}
	//设置公告
	httpSetAnnouncement(w, announcement)
	//设置跑马灯
	httpSetPaoMaDeng(w, paomadeng)

}

func httpSetAnnouncement(w http.ResponseWriter, announcement *modelCross.Announcement) {
	if announcement == nil {
		return
	}
	oldAnnouncement := modelCross.GetAnnouncementModel().GetAnnouncementByGmId(announcement.GmId)
	if oldAnnouncement != nil {
		logger.Info("更新公告：%v", oldAnnouncement.Id, oldAnnouncement.GmId)
		announcement.Id = oldAnnouncement.Id
		_, err := modelCross.GetAnnouncementModel().Update(announcement)
		if err != nil {
			logger.Error("GetAnnouncementModel  insert err:%v   data:%v", err, announcement)
			ptsdk.GetSdk().HttpWriteReturnInfo(w, 400, "数据库操作失败", nil)
			return
		}
	} else {
		err := modelCross.GetAnnouncementModel().Create(announcement)
		if err != nil {
			logger.Error("GetAnnouncementModel  insert err:%v   data:%v", err, announcement)
			ptsdk.GetSdk().HttpWriteReturnInfo(w, 400, "数据库操作失败", nil)
			return
		}
	}

	ptsdk.GetSdk().HttpWriteReturnInfo(w, 200, "success", nil)
	m.GetGsServers().SendAllServerMessage(&pbserver.UpAnnouncementNowReq{})
	keys := common.NewStringSlice(modelCross.CLIENT_TO_LOGIN, "_")[0]
	portInfo, err := modelCross.GetServerPortInfoModel().GetLoginServerPortInfos(keys)
	if err == nil && portInfo != nil && len(portInfo) > 0 {
		for _, v := range portInfo {
			r, err := http.PostForm(fmt.Sprintf("http://%s:%d/upAnnouncement", v.Host, v.Port), url.Values{})
			logger.Info("通知登录更新公告：%v", r, err)
		}
	}
}

/**
*  @Description:设置跑马灯
*  @param w
*  @param r
**/
func httpSetPaoMaDeng(w http.ResponseWriter, pamadeng *modelCross.PaoMaDeng) {

	if pamadeng == nil {
		return
	}
	//设置跑马灯类型
	types := constAnnouncement.CronPlay
	if pamadeng.IntervalTimes > 0 {
		types = constAnnouncement.IntervalPlay
	} else {
		types = constAnnouncement.CyclePlay
	}
	pamadeng.Types = types

	oldData := modelCross.GetPaoMaDengModel().GetPaomadengByGmId(pamadeng.GmId)
	if oldData != nil {
		logger.Info("更新跑马灯：%v", oldData.Id, oldData.GmId)
		pamadeng.Id = oldData.Id
		_, err := modelCross.GetPaoMaDengModel().Update(pamadeng)
		if err != nil {
			logger.Error("GetAnnouncementModel  insert err:%v   data:%v", err, pamadeng)
			ptsdk.GetSdk().HttpWriteReturnInfo(w, 400, "数据库操作失败", nil)
			return
		}
	} else {
		err := modelCross.GetPaoMaDengModel().Create(pamadeng)
		if err != nil {
			logger.Error("GetAnnouncementModel  insert err:%v   data:%v", err, pamadeng)
			ptsdk.GetSdk().HttpWriteReturnInfo(w, 400, "数据库操作失败", nil)
			return
		}
	}

	ptsdk.GetSdk().HttpWriteReturnInfo(w, 200, "success", nil)
	m.GetGsServers().SendAllServerMessage(&pbserver.UpPaoMaDengNowReq{})
	return
}
