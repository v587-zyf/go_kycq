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

type MiningModel struct {
	dbmodel.CommonModel
}

type MiningModels struct {
	models   map[string]*MiningModel
	newModel *MiningModel
}

var (
	miningModels = &MiningModels{
		models:   make(map[string]*MiningModel),
		newModel: &MiningModel{},
	}
	miningFields    = GetAllFieldsAsString(modelGame.MiningDb{})
	miningTableName = (&modelGame.MiningDb{}).TableName()
)

func dbMiningInit() {

	logger.Info("miningModels 初始化")
	for k, _ := range base.Conf.DbConfigs {
		k1 := strings.Split(k, "_")
		if len(k1) == 2 && k1[0] == DB_SERVER {
			if miningModels.models[k] == nil {
				miningModels.models[k] = &MiningModel{}
			}
			dbmodel.Register(k, miningModels.models[k], func(dbMap *gorp.DbMap) {
				dbMap.AddTableWithName(modelGame.MiningDb{}, miningTableName).SetKeys(true, "id")
			})
		}
	}
	dbmodel.Register(NEW_SERVER, miningModels.newModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(modelGame.MiningDb{}, miningTableName).SetKeys(true, "id")
		orm.RegisterModelForAlias(NEW_SERVER, new(modelGame.MiningDb))
	})
}

func GetMiningModel() *MiningModels {
	return miningModels
}

//获取有效玩家数据
func (this *MiningModels) GetDatas(dbKey string, userIds []int) ([]modelGame.MiningDb, error) {

	ids := ""
	for _, v := range userIds {
		ids += strconv.Itoa(v) + ","
	}
	ids = ids[:len(ids)-1]

	var data []modelGame.MiningDb
	sqlStr := fmt.Sprintf("select %s from %s where userId in (%s) ", miningFields, miningTableName, ids)
	_, err := this.models[dbKey].DbMap().Select(&data, sqlStr)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//插入玩家数据
func (this *MiningModels) InsertNewData(data *modelGame.MiningDb) error {
	return this.newModel.DbMap().Insert(data)
}

func (this *MiningModels) Clean() {
	this.newModel.DbMap().Exec("delete from " + miningTableName + " where 1")
}
