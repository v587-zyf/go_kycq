package modelCross

import (
	"cqserver/gamelibs/model"
	"cqserver/golibs/dbmodel"
	"cqserver/golibs/logger"
	"database/sql"
	"errors"
	"fmt"
	"gopkg.in/gorp.v1"
)

type WhiteListDb struct {
	Id      int    `db:"id"`
	GMId    int    `db:"gmId" orm:"size(200);comment(gm平台Id)"`
	Valtype int    `db:"valtype" orm:"size(50);comment(白名单类型（1：ip,2:账号）)"`
	Value   string `db:"value" orm:"size(200);comment(白名单))"`
}

type WhiteListDbModel struct {
	dbmodel.CommonModel
}

var (
	whiteListDbModel       = &WhiteListDbModel{}
	whiteListDbModelFields = dbmodel.GetAllFieldsAsString(WhiteListDb{})
)

func init() {
	dbmodel.Register(model.DB_ACCOUNT, whiteListDbModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(WhiteListDb{}, "white_list").SetKeys(true, "id")
	})
}

func GetWhiteListDbModel() *WhiteListDbModel {
	return whiteListDbModel
}

func (this *WhiteListDbModel) Create(info *WhiteListDb) error {
	return this.DbMap().Insert(info)
}

func (this *WhiteListDbModel) Del(gmId int, whiteVal string) error {
	sql := "delete from white_list where 1"
	if gmId > 0 {
		sql += fmt.Sprintf(" and gmId=%d", gmId)
	} else if len(whiteVal) > 0 {
		sql += fmt.Sprintf(" and Value='%s'", whiteVal)
	} else {
		logger.Error("删除白名单数据异常")
		return errors.New("参数错误")
	}
	_, err := this.DbMap().Exec(sql)
	return err
}

func (this *WhiteListDbModel) Getall() []*WhiteListDb {
	var all []*WhiteListDb
	_, err := this.DbMap().Select(&all, "select * from white_list where 1")
	if err != nil && err != sql.ErrNoRows {

		logger.Error("获取功能关闭配置错误：%v", err)
	}
	return all
}
