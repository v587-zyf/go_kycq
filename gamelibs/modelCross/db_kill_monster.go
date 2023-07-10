package modelCross

import (
	"cqserver/gamelibs/model"
	"cqserver/golibs/dbmodel"
	"fmt"
	"gopkg.in/gorp.v1"
	"time"
)

type KillMonsterDb struct {
	dbmodel.DbTable
	ServerId        int       `db:"serverId" orm:"comment(服务器id)"`
	StageId         int       `db:"stageId" orm:"comment(首领id)"`
	FirstKillUserId int       `db:"firstKillUserId" orm:"comment(首次击杀玩家id)"`
	FirstKillTime   time.Time `db:"firstKillTime" orm:"comment(首次击杀时间)"`
	KillNumAll      int       `db:"killNumAll" orm:"type(int64);comment(总击杀数量)"`
}

func (this *KillMonsterDb) TableName() string {
	return "kill_monster"
}

type KillMonsterModel struct {
	dbmodel.CommonModel
}

var (
	killMonsterDb     = &KillMonsterDb{}
	killMonsterModel  = &KillMonsterModel{}
	killMonsterFields = model.GetAllFieldsAsString(KillMonsterDb{})
)

func init() {
	dbmodel.Register(model.DB_ACCOUNT, killMonsterModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(KillMonsterDb{}, killMonsterDb.TableName()).SetKeys(false, "serverId", "stageId")
	})
}

func GetKillMonsterDbModel() *KillMonsterModel {
	return killMonsterModel
}

func (this *KillMonsterModel) Create(data *KillMonsterDb) error {
	return this.DbMap().Insert(data)
}

func (this *KillMonsterModel) Update(data *KillMonsterDb) error {
	_, err := this.DbMap().Update(data)
	return err
}

//获取所有数据
func (this *KillMonsterModel) LoadAlLData() ([]*KillMonsterDb, error) {
	sql := fmt.Sprintf("select %s from %s where 1", killMonsterFields, killMonsterDb.TableName())
	var data []*KillMonsterDb
	_, err := this.DbMap().Select(&data, sql)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//根据服务器id加载首领击杀信息
func (this *KillMonsterModel) LoadAllDataByServerId(serverId int) ([]*KillMonsterDb, error) {
	sql := fmt.Sprintf("select %s from %s where serverId=%d", killMonsterFields, killMonsterDb.TableName(), serverId)
	var data []*KillMonsterDb
	_, err := this.DbMap().Select(&data, sql)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//根据serverId, stageId查询
func (this *KillMonsterModel) LoadDataByStageId(serverId, stageId int) (*KillMonsterDb, error) {
	sql := fmt.Sprintf(`select %s from %s where serverId=%d and stageId=%d`, killMonsterFields, killMonsterDb.TableName(), serverId, stageId)
	var data KillMonsterDb
	err := this.DbMap().SelectOne(&data, sql)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
