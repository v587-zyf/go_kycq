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

type GuildAuctionModel struct {
	dbmodel.CommonModel
}

type GuildAuctionModels struct {
	models   map[string]*GuildAuctionModel
	newModel *GuildAuctionModel
}

var (
	guildAuctionModels = &GuildAuctionModels{
		models:   make(map[string]*GuildAuctionModel),
		newModel: &GuildAuctionModel{},
	}
	guildAuctionFields    = GetAllFieldsAsString(modelGame.GuildAuctionItem{})
	guildAuctionTableName = (&modelGame.GuildAuctionItem{}).TableName()
)

func dbGuildAuctionInit() {

	logger.Info("GuildModels 初始化")
	for k, _ := range base.Conf.DbConfigs {
		k1 := strings.Split(k, "_")
		if len(k1) == 2 && k1[0] == DB_SERVER {
			if guildAuctionModels.models[k] == nil {
				guildAuctionModels.models[k] = &GuildAuctionModel{}
			}
			dbmodel.Register(k, guildAuctionModels.models[k], func(dbMap *gorp.DbMap) {
				dbMap.AddTableWithName(modelGame.GuildAuctionItem{}, guildAuctionTableName).SetKeys(true, "id")
			})
		}
	}
	dbmodel.Register(NEW_SERVER, guildAuctionModels.newModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(modelGame.GuildAuctionItem{}, guildAuctionTableName).SetKeys(true, "id")
		orm.RegisterModelForAlias(NEW_SERVER, new(modelGame.GuildAuctionItem))
	})
}

func GetGuildAuctionModel() *GuildAuctionModels {
	return guildAuctionModels
}

//获取有效玩家数据
func (this *GuildAuctionModels) GetDatas(dbKey string) ([]modelGame.GuildAuctionItem, error) {
	var data []modelGame.GuildAuctionItem
	sqlStr := fmt.Sprintf("select %s from %s where 1", guildAuctionFields, guildAuctionTableName)
	_, err := this.models[dbKey].DbMap().Select(&data, sqlStr)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//插入玩家数据
func (this *GuildAuctionModels) InsertNewData(auction *modelGame.GuildAuctionItem) error {
	return this.newModel.DbMap().Insert(auction)
}

func (this *GuildAuctionModels) Clean() {
	this.newModel.DbMap().Exec("delete from " + guildAuctionTableName + " where 1")
}
