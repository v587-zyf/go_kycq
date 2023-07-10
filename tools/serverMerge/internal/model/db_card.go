package model

import (
	"cqserver/gamelibs/modelGame"
	"cqserver/golibs/dbmodel"
	"cqserver/golibs/logger"
	"cqserver/tools/serverMerge/internal/base"
	"fmt"
	"github.com/astaxie/beego/orm"
	"gopkg.in/gorp.v1"
	"strings"
)

type CardModel struct {
	dbmodel.CommonModel
}

type CardModels struct {
	models   map[string]*CardModel
	newModel *CardModel
}

var (
	cardModels = &CardModels{
		models:   make(map[string]*CardModel),
		newModel: &CardModel{},
	}
	cardFields    = GetAllFieldsAsString(modelGame.Card{})
	cardTableName = (&modelGame.Card{}).TableName()
)

func dbCardInit() {

	logger.Info("CardModels 初始化")
	for k, _ := range base.Conf.DbConfigs {
		k1 := strings.Split(k, "_")
		if len(k1) == 2 && k1[0] == DB_SERVER {
			if cardModels.models[k] == nil {
				cardModels.models[k] = &CardModel{}
			}
			dbmodel.Register(k, cardModels.models[k], func(dbMap *gorp.DbMap) {
				dbMap.AddTableWithName(modelGame.Card{}, cardTableName).SetKeys(true, "id")
			})
		}
	}
	dbmodel.Register(NEW_SERVER, cardModels.newModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(modelGame.Card{}, cardTableName).SetKeys(true, "id")
		orm.RegisterModelForAlias(NEW_SERVER, new(modelGame.Card))
	})
}

func GetCardModel() *CardModels {
	return cardModels
}

//获取有效玩家数据
func (this *CardModels) GetDatas(dbKey string) ([]modelGame.Card, error) {
	var data []modelGame.Card
	sqlStr := fmt.Sprintf("select %s from %s where 1", cardFields, cardTableName)
	_, err := this.models[dbKey].DbMap().Select(&data, sqlStr)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//插入玩家数据
func (this *CardModels) InsertNewData(card *modelGame.Card) error {
	return this.newModel.DbMap().Insert(card)
}

func (this *CardModels) Clean() {
	this.newModel.DbMap().Exec("delete from " + cardTableName + " where 1")
}
