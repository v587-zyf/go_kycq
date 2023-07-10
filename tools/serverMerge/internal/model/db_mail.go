package model

import (
	"cqserver/gamelibs/modelGame"
	"cqserver/golibs/dbmodel"
	"cqserver/golibs/logger"
	"cqserver/tools/serverMerge/internal/base"
	"fmt"
	"github.com/astaxie/beego/orm"
	"gopkg.in/gorp.v1"
	"strconv"
	"strings"
)

type MailModel struct {
	dbmodel.CommonModel
}

type MailModels struct {
	models   map[string]*MailModel
	newModel *MailModel
}

var (
	mailModels = &MailModels{
		models:   make(map[string]*MailModel),
		newModel: &MailModel{},
	}
	mailFields = GetAllFieldsAsString(modelGame.Mail{})
	mailTableName = (&modelGame.Mail{}).TableName()
)

func dbMailInit() {
	logger.Info("mail 初始化")
	for k, _ := range base.Conf.DbConfigs {
		k1 := strings.Split(k, "_")
		if len(k1) == 2 && k1[0] == DB_SERVER {
			if mailModels.models[k] == nil {
				mailModels.models[k] = &MailModel{}
			}
			dbmodel.Register(k, mailModels.models[k], func(dbMap *gorp.DbMap) {
				dbMap.AddTableWithName(modelGame.Mail{}, mailTableName).SetKeys(false, "Id")
			})
		}
	}
	dbmodel.Register(NEW_SERVER, mailModels.newModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(modelGame.Mail{}, mailTableName).SetKeys(true, "Id")
		orm.RegisterModelForAlias(NEW_SERVER, new(modelGame.Mail))
	})
}

func GetMailModel() *MailModels {
	return mailModels
}

func (this *MailModels) GetDatas(dbKey string, userIds []int) ([]modelGame.Mail, error) {
	ids := ""
	for _, v := range userIds {
		ids += strconv.Itoa(v) + ","
	}
	ids = ids[:len(ids)-1]
	var data []modelGame.Mail
	sqlStr := fmt.Sprintf("select %s from mail where userId in (%s) and redeemedAt='0000-00-00 00-00-00'", mailFields, ids)
	_, err := this.models[dbKey].DbMap().Select(&data, sqlStr)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (this *MailModels) InsertNewData(mailData *modelGame.Mail) error {
	return this.newModel.DbMap().Insert(mailData)
}

func (this *MailModels) Clean() {
	this.newModel.DbMap().Exec("delete from mail where 1")
}
