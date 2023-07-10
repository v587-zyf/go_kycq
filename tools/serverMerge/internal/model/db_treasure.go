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

type TreasureModel struct {
	dbmodel.CommonModel
}

type TreasureModels struct {
	models   map[string]*TreasureModel
	newModel *TreasureModel
}

var (
	treasureModels = &TreasureModels{
		models:   make(map[string]*TreasureModel),
		newModel: &TreasureModel{},
	}
	treasureFields    = GetAllFieldsAsString(modelGame.Treasure{})
	treasureTableName = (&modelGame.Treasure{}).TableName()
)

func dbTreasureInit() {

	logger.Info("TreasureModels 初始化")
	for k, _ := range base.Conf.DbConfigs {
		k1 := strings.Split(k, "_")
		if len(k1) == 2 && k1[0] == DB_SERVER {
			if treasureModels.models[k] == nil {
				treasureModels.models[k] = &TreasureModel{}
			}
			dbmodel.Register(k, treasureModels.models[k], func(dbMap *gorp.DbMap) {
				dbMap.AddTableWithName(modelGame.Treasure{}, treasureTableName).SetKeys(true, "id")
			})
		}
	}
	dbmodel.Register(NEW_SERVER, treasureModels.newModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(modelGame.Treasure{}, treasureTableName).SetKeys(true, "id")
		orm.RegisterModelForAlias(NEW_SERVER, new(modelGame.Treasure))
	})
}

func GetTreasureModel() *TreasureModels {
	return treasureModels
}

//获取有效玩家数据
func (this *TreasureModels) GetDatas(dbKey string) ([]modelGame.Treasure, error) {
	var data []modelGame.Treasure
	sqlStr := fmt.Sprintf("select %s from %s where 1", treasureFields, treasureTableName)
	_, err := this.models[dbKey].DbMap().Select(&data, sqlStr)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//插入玩家数据
func (this *TreasureModels) InsertNewData(treasure *modelGame.Treasure) error {
	return this.newModel.DbMap().Insert(treasure)
}

func (this *TreasureModels) Clean() {
	this.newModel.DbMap().Exec("delete from " + treasureTableName + " where 1")
}
