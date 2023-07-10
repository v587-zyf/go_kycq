package modelCross

import (
	"cqserver/gamelibs/model"
	"cqserver/golibs/dbmodel"
	"cqserver/golibs/logger"
	"database/sql"
	"fmt"
	"gopkg.in/gorp.v1"
)

type FuncCloseDb struct {
	FuncId    int            `db:"funcId"`
	ServerIds model.IntSlice `db:"serverIds"`
}

type FuncCloseDbModel struct {
	dbmodel.CommonModel
}

var (
	funcCloseDbModel       = &FuncCloseDbModel{}
	funcCloseDbModelFields = dbmodel.GetAllFieldsAsString(FuncCloseDb{})
)

func init() {
	dbmodel.Register(model.DB_ACCOUNT, funcCloseDbModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(FuncCloseDb{}, "func_close").SetKeys(false, "funcId")
	})
}

func GetFuncCloseDbModel() *FuncCloseDbModel {
	return funcCloseDbModel
}

func (this *FuncCloseDbModel) Create(info *FuncCloseDb) error {
	return this.DbMap().Insert(info)
}

func (this *FuncCloseDbModel) Update(info *FuncCloseDb) (int, error) {
	count, err := this.DbMap().Update(info)
	if err != nil {
		fmt.Println("err", err)
	}
	return int(count), err
}

func (this *FuncCloseDbModel) Del(info *FuncCloseDb) {
	this.DbMap().Delete(info)
}

func (this *FuncCloseDbModel) Getall() []*FuncCloseDb {
	var all []*FuncCloseDb
	_, err := this.DbMap().Select(&all, "select * from func_close where 1")
	if err != nil && err != sql.ErrNoRows {

		logger.Error("获取功能关闭配置错误：%v", err)
	}
	return all
}
