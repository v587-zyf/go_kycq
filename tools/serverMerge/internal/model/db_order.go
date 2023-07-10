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

type OrderModel struct {
	dbmodel.CommonModel
}

type OrderModels struct {
	models   map[string]*OrderModel
	newModel *OrderModel
}

var (
	orderModels = &OrderModels{
		models:   make(map[string]*OrderModel),
		newModel: &OrderModel{},
	}
	orderFields    = GetAllFieldsAsString(modelGame.OrderDb{})
	orderTableName = (&modelGame.OrderDb{}).TableName()
)

func dbOrderInit() {

	logger.Info("OrderModel 初始化")
	for k, _ := range base.Conf.DbConfigs {
		k1 := strings.Split(k, "_")
		if len(k1) == 2 && k1[0] == DB_SERVER {
			if orderModels.models[k] == nil {
				orderModels.models[k] = &OrderModel{}
			}
			dbmodel.Register(k, orderModels.models[k], func(dbMap *gorp.DbMap) {
				dbMap.AddTableWithName(modelGame.OrderDb{}, orderTableName).SetKeys(true, "id")
			})
		}
	}
	dbmodel.Register(NEW_SERVER, orderModels.newModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(modelGame.OrderDb{}, orderTableName).SetKeys(true, "id")
		orm.RegisterModelForAlias(NEW_SERVER, new(modelGame.OrderDb))
	})
}

func GetOrderModel() *OrderModels {
	return orderModels
}

//获取有效玩家数据
func (this *OrderModels) GetOrderData(dbKey string, userIds []int) ([]modelGame.OrderDb, error) {

	ids := ""
	for _, v := range userIds {
		ids += strconv.Itoa(v) + ","
	}
	ids = ids[:len(ids)-1]

	var data []modelGame.OrderDb
	sqlStr := fmt.Sprintf("select %s from %s where userId in (%s) ", orderFields, orderTableName, ids)
	_, err := this.models[dbKey].DbMap().Select(&data, sqlStr)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//插入玩家数据
func (this *OrderModels) InsertNewData(data *modelGame.OrderDb) error {
	return this.newModel.DbMap().Insert(data)
}

func (this *OrderModels) Clean() {
	this.newModel.DbMap().Exec("delete from " + orderTableName + " where 1")
}

