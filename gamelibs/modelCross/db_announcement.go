package modelCross

import (
	"cqserver/gamelibs/model"
	"cqserver/golibs/dbmodel"
	"cqserver/golibs/logger"
	"database/sql"
	"fmt"
	"gopkg.in/gorp.v1"
	"time"
)

type Announcement struct {
	Id           int            `db:"id"`
	GmId         int            `db:"gmId" orm:"comment(平台用Id)"`
	Type         int            `db:"type" orm:"comment(1:登入公告 2.游戏公告)"`
	Title        string         `db:"title" orm:"size(300);comment(公告名称)"`
	StartTime    time.Time      `db:"startTime" orm:"comment(开始时间)"`
	EndTime      time.Time      `db:"endTime" orm:"comment(结束时间)"`
	ChannelIds   model.IntSlice `db:"channelIds"  orm:"size(500);comment(渠道id)"`
	ServerIds    model.IntSlice `db:"serverIds"  orm:"size(500);comment(区服id)"`
	Announcement string         `db:"announcement" orm:"type(text);comment(公告内容)"`
}

type AnnouncementModel struct {
	dbmodel.CommonModel
}

var (
	loginAnnouncementModel  = &AnnouncementModel{}
	loginAnnouncementFields = dbmodel.GetAllFieldsAsString(Announcement{})
)

func init() {
	dbmodel.Register(model.DB_ACCOUNT, loginAnnouncementModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(Announcement{}, "announcement").SetKeys(true, "id")
	})
}

func GetAnnouncementModel() *AnnouncementModel {
	return loginAnnouncementModel
}

func (this *AnnouncementModel) Create(info *Announcement) error {
	return this.DbMap().Insert(info)
}

func (this *AnnouncementModel) Update(info *Announcement) (int, error) {
	count, err := this.DbMap().Update(info)
	if err != nil {
		logger.Error("err", err)
	}
	return int(count), err
}

func (this *AnnouncementModel) Del(gmId int) error {
	sql := fmt.Sprintf("delete from announcement where  gmId=%d", gmId)
	_, err := this.DbMap().Exec(sql)
	return err
}

func (this *AnnouncementModel) GetAnnouncementByGmId(gmId int) *Announcement {
	var announcement Announcement
	err := this.DbMap().SelectOne(&announcement, fmt.Sprintf("SELECT %s FROM announcement where gmId=%d", loginAnnouncementFields, gmId))
	if err != nil && err != sql.ErrNoRows {
		logger.Error("获取公告数据错误：%v", err)
		return nil
	}
	if err == sql.ErrNoRows {
		return nil
	}
	return &announcement

}

//types 1:登入公告 2.游戏公告
func (this *AnnouncementModel) GetAllLoginAnnouncement(types int) ([]Announcement, error) {
	var announcement []Announcement
	_, err := this.DbMap().Select(&announcement, fmt.Sprintf("SELECT %s FROM announcement WHERE NOW() BETWEEN startTime AND endTime  and type = %v  ORDER BY id DESC limit 5 ", loginAnnouncementFields, types))
	if err != nil {
		return nil, err
	}
	return announcement, nil
}
