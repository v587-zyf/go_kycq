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

type GuildModel struct {
	dbmodel.CommonModel
}

type GuildModels struct {
	models   map[string]*GuildModel
	newModel *GuildModel
}

var (
	guildModels = &GuildModels{
		models:   make(map[string]*GuildModel),
		newModel: &GuildModel{},
	}
	guildFields = GetAllFieldsAsString(modelGame.Guild{})
	tableName   = (&modelGame.Guild{}).TableName()
)

func dbGuildInit() {

	logger.Info("GuildModels 初始化")
	for k, _ := range base.Conf.DbConfigs {
		k1 := strings.Split(k, "_")
		if len(k1) == 2 && k1[0] == DB_SERVER {
			if guildModels.models[k] == nil {
				guildModels.models[k] = &GuildModel{}
			}
			dbmodel.Register(k, guildModels.models[k], func(dbMap *gorp.DbMap) {
				dbMap.AddTableWithName(modelGame.Guild{}, tableName).SetKeys(true, "id")
			})
		}
	}
	dbmodel.Register(NEW_SERVER, guildModels.newModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(modelGame.Guild{}, tableName).SetKeys(true, "id")
		orm.RegisterModelForAlias(NEW_SERVER, new(modelGame.Guild))
	})
}

func GetGuildModel() *GuildModels {
	return guildModels
}

//获取有效玩家数据
func (this *GuildModels) GetDatas(dbKey string) ([]modelGame.Guild, error) {
	var data []modelGame.Guild
	sqlStr := fmt.Sprintf("select %s from %s where 1", guildFields, tableName)
	_, err := this.models[dbKey].DbMap().Select(&data, sqlStr)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//插入玩家数据
func (this *GuildModels) InsertNewData(hero *modelGame.Guild) error {
	return this.newModel.DbMap().Insert(hero)
}

func (this *GuildModels) Clean() {
	this.newModel.DbMap().Exec("delete from " + tableName + " where 1")
}
