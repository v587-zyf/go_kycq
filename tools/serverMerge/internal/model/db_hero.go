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

type HeroModel struct {
	dbmodel.CommonModel
}

type HeroModels struct {
	models   map[string]*HeroModel
	newModel *HeroModel
}

var (
	heroModels = &HeroModels{
		models:   make(map[string]*HeroModel),
		newModel: &HeroModel{},
	}
	heroFields = GetAllFieldsAsString(modelGame.Hero{})
	heroTableName = (&modelGame.Hero{}).TableName()
)

func dbHeroInit() {
	logger.Info("heroModels 初始化")
	for k, _ := range base.Conf.DbConfigs {
		k1 := strings.Split(k, "_")
		if len(k1) == 2 && k1[0] == DB_SERVER {
			if heroModels.models[k] == nil {
				heroModels.models[k] = &HeroModel{}
			}
			dbmodel.Register(k, heroModels.models[k], func(dbMap *gorp.DbMap) {
				dbMap.AddTableWithName(modelGame.Hero{}, heroTableName).SetKeys(true, "id")
			})
		}
	}
	dbmodel.Register(NEW_SERVER, heroModels.newModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(modelGame.Hero{}, heroTableName).SetKeys(true, "id")
		orm.RegisterModelForAlias(NEW_SERVER, new(modelGame.Hero))
	})
}

func GetHeroModel() *HeroModels {
	return heroModels
}

//获取有效玩家数据
func (this *HeroModels) GetDatas(dbKey string, userIds []int) ([]modelGame.Hero, error) {
	ids := ""
	for _, v := range userIds {
		ids += strconv.Itoa(v) + ","
	}
	ids = ids[:len(ids)-1]
	var data []modelGame.Hero
	sqlStr := fmt.Sprintf("select %s from hero where userId in (%s)", heroFields, ids)
	_, err := this.models[dbKey].DbMap().Select(&data, sqlStr)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//插入玩家数据
func (this *HeroModels) InsertNewData(hero *modelGame.Hero) error {
	return this.newModel.DbMap().Insert(hero)
}

func (this *HeroModels) Clean() {
	this.newModel.DbMap().Exec("delete from hero where 1")
}
