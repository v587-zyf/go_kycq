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

type AuctionBidInfoModel struct {
	dbmodel.CommonModel
}

type AuctionBidInfoModels struct {
	models   map[string]*AuctionBidInfoModel
	newModel *AuctionBidInfoModel
}

var (
	auctionBidModels = &AuctionBidInfoModels{
		models:   make(map[string]*AuctionBidInfoModel),
		newModel: &AuctionBidInfoModel{},
	}
	auctionBidFields        = GetAllFieldsAsString(modelGame.AuctionBid{})
	auctionBidInfoTableName = (&modelGame.AuctionBid{}).TableName()
)

func dbAuctionBidInit() {

	logger.Info("GuildModels 初始化")
	for k, _ := range base.Conf.DbConfigs {
		k1 := strings.Split(k, "_")
		if len(k1) == 2 && k1[0] == DB_SERVER {
			if auctionBidModels.models[k] == nil {
				auctionBidModels.models[k] = &AuctionBidInfoModel{}
			}
			dbmodel.Register(k, auctionBidModels.models[k], func(dbMap *gorp.DbMap) {
				dbMap.AddTableWithName(modelGame.AuctionBid{}, auctionBidInfoTableName).SetKeys(true, "id")
			})
		}
	}
	dbmodel.Register(NEW_SERVER, auctionBidModels.newModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(modelGame.AuctionBid{}, auctionBidInfoTableName).SetKeys(true, "id")
		orm.RegisterModelForAlias(NEW_SERVER, new(modelGame.AuctionBid))
	})
}

func GetAuctionBidModel() *AuctionBidInfoModels {
	return auctionBidModels
}

//获取有效玩家数据
func (this *AuctionBidInfoModels) GetDatas(dbKey string) ([]modelGame.AuctionBid, error) {
	var data []modelGame.AuctionBid
	sqlStr := fmt.Sprintf("select %s from %s where 1", auctionBidFields, auctionBidInfoTableName)
	_, err := this.models[dbKey].DbMap().Select(&data, sqlStr)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//插入玩家数据
func (this *AuctionBidInfoModels) InsertNewData(hero *modelGame.AuctionBid) error {
	return this.newModel.DbMap().Insert(hero)
}

func (this *AuctionBidInfoModels) Clean() {
	this.newModel.DbMap().Exec("delete from " + auctionBidInfoTableName + " where 1")
}
