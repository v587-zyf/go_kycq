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

type PaoMaDeng struct {
	Id            int            `db:"id"`
	GmId          int            `db:"gmId" orm:"comment(平台用Id)"`
	Types         int            `db:"types" orm:"comment(0:准点播放 1:循环播放 2:间隔播放)"`
	CycleTimes    int            `db:"cycleTimes" orm:"comment(循环播放次数)"`
	IntervalTimes int            `db:"intervalTimes" orm:"comment(间隔播放时间)"`
	Content       string         `db:"content" orm:"type(text);comment(跑马灯内容)"`
	StartTime     time.Time      `db:"startTime" orm:"comment(开始时间)"`
	EndTime       time.Time      `db:"endTime" orm:"comment(结束时间)"`
	ChannelIds    model.IntSlice `db:"channelIds"  orm:"size(500);comment(渠道id)"`
	ServerIds     model.IntSlice `db:"serverIds"  orm:"size(500);comment(区服id)"`
}

type PaoMaDengModel struct {
	dbmodel.CommonModel
}

var (
	PaoMaDengModels = &PaoMaDengModel{}
	PaoMaDengFields = dbmodel.GetAllFieldsAsString(PaoMaDeng{})
)

func init() {
	dbmodel.Register(model.DB_ACCOUNT, PaoMaDengModels, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(PaoMaDeng{}, "paomadeng").SetKeys(true, "id")
	})
}

func GetPaoMaDengModel() *PaoMaDengModel {
	return PaoMaDengModels
}

func (this *PaoMaDengModel) Create(info *PaoMaDeng) error {
	return this.DbMap().Insert(info)
}

func (this *PaoMaDengModel) Update(info *PaoMaDeng) (int, error) {
	count, err := this.DbMap().Update(info)
	if err != nil {
		logger.Error("err", err)
	}
	return int(count), err
}

func (this *PaoMaDengModel) Del(gmId int) error {
	sql := fmt.Sprintf("delete from paomadeng where  gmId=%d", gmId)
	re, err := this.DbMap().Exec(sql)
	logger.Info("-------------------", re, &re)
	return err
}

func (this *PaoMaDengModel) GetPaomadengByGmId(gmId int) *PaoMaDeng {
	var paoMaDeng PaoMaDeng
	err := this.DbMap().SelectOne(&paoMaDeng, fmt.Sprintf("SELECT %s FROM paomadeng where gmId=%d", PaoMaDengFields, gmId))
	if err != nil && err != sql.ErrNoRows {
		logger.Error("获取跑马灯数据错误：%v", err)
		return nil
	}
	if err == sql.ErrNoRows {
		return nil
	}
	return &paoMaDeng

}

//types 1:登入公告 2.游戏公告
func (this *PaoMaDengModel) GetAllPaoMaDeng() ([]PaoMaDeng, error) {
	var paoMaDeng []PaoMaDeng
	_, err := this.DbMap().Select(&paoMaDeng, fmt.Sprintf("SELECT %s FROM paomadeng WHERE NOW() BETWEEN startTime AND endTime ORDER BY id DESC ", PaoMaDengFields))
	if err != nil {
		return nil, err
	}
	return paoMaDeng, nil
}
