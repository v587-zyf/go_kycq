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

type WorldAuctionModel struct {
	dbmodel.CommonModel
}

type WorldAuctionModels struct {
	models   map[string]*WorldAuctionModel
	newModel *WorldAuctionModel
}

var (
	worldAuctionModels = &WorldAuctionModels{
		models:   make(map[string]*WorldAuctionModel),
		newModel: &WorldAuctionModel{},
	}
	worldAuctionFields    = GetAllFieldsAsString(modelGame.AuctionItem{})
	worldAuctionTableName = (&modelGame.AuctionItem{}).TableName()
)

func dbWorldAuctionInit() {

	logger.Info("GuildModels 初始化")
	for k, _ := range base.Conf.DbConfigs {
		k1 := strings.Split(k, "_")
		if len(k1) == 2 && k1[0] == DB_SERVER {
			if worldAuctionModels.models[k] == nil {
				worldAuctionModels.models[k] = &WorldAuctionModel{}
			}
			dbmodel.Register(k, worldAuctionModels.models[k], func(dbMap *gorp.DbMap) {
				dbMap.AddTableWithName(modelGame.AuctionItem{}, worldAuctionTableName).SetKeys(true, "id")
			})
		}
	}
	dbmodel.Register(NEW_SERVER, worldAuctionModels.newModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(modelGame.AuctionItem{}, worldAuctionTableName).SetKeys(true, "id")
		orm.RegisterModelForAlias(NEW_SERVER, new(modelGame.AuctionItem))
	})
}

func GetWorldAuctionModel() *WorldAuctionModels {
	return worldAuctionModels
}

//获取有效玩家数据
func (this *WorldAuctionModels) GetDatas(dbKey string) ([]modelGame.AuctionItem, error) {
	var data []modelGame.AuctionItem
	sqlStr := fmt.Sprintf("select %s from %s where 1", worldAuctionFields, worldAuctionTableName)
	_, err := this.models[dbKey].DbMap().Select(&data, sqlStr)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//插入玩家数据
func (this *WorldAuctionModels) InsertNewData(auction *modelGame.AuctionItem) error {
	return this.newModel.DbMap().Insert(auction)
}

func (this *WorldAuctionModels) Clean() {
	this.newModel.DbMap().Exec("delete from " + worldAuctionTableName + " where 1")
}
