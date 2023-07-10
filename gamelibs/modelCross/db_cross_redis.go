package modelCross

import (
	"cqserver/gamelibs/model"
	"cqserver/golibs/dbmodel"
	"fmt"
	"gopkg.in/gorp.v1"
)

type CrossRedisInfo struct {
	Id       int    `db:"id"`
	Network  string `db:"network"`
	Address  string `db:"address"`
	Password string `db:"password"`
	Db       int    `db:"db"`
}

type CrossRedisModel struct {
	dbmodel.CommonModel
}

var (
	crossRedisModel  = &CrossRedisModel{}
	crossRedisFileds = model.GetAllFieldsAsString(CrossRedisInfo{})
)

func init() {

	dbmodel.Register(model.DB_ACCOUNT, crossRedisModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(CrossRedisInfo{}, "cross_redis").SetKeys(true, "id")
	})
}


func GetCrossRedisModel()*CrossRedisModel {
	return crossRedisModel
}

func (this *CrossRedisModel) GetCrossRedisConfs() ([]CrossRedisInfo, error) {
	var info []CrossRedisInfo
	_, err := this.DbMap().Select(&info, fmt.Sprintf("select id,network,address,password,db from cross_redis where 1"))
	if err != nil {
		return nil, err
	}
	return info, nil
}
