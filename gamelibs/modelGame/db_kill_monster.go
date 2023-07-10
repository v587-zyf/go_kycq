package modelGame

import (
	"cqserver/gamelibs/model"
	"cqserver/golibs/dbmodel"
	"fmt"
	"github.com/astaxie/beego/orm"
	"gopkg.in/gorp.v1"
	"time"
)

type KillMonsterDb struct {
	dbmodel.DbTable
	Id              int                   `db:"id" orm:"pk;auto"`
	StageId         int                   `db:"stageId" orm:"comment(首领id)"`
	FirstKillUserId int                   `db:"firstKillUserId" orm:"comment(首次击杀玩家id)"`
	FirstKillTime   time.Time             `db:"firstKillTime" orm:"comment(首次击杀时间)"`
	KillNumAll      int                   `db:"killNumAll" orm:"type(int64);comment(总击杀数量)"`
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
	dbmodel.Register(model.DB_SERVER, killMonsterModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(KillMonsterDb{}, killMonsterDb.TableName()).SetKeys(true, "id")
		orm.RegisterModelForAlias(model.DB_SERVER, new(KillMonsterDb))
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

//根据stageId查询
func (this *KillMonsterModel) LoadDataByStageId(stageId int) (*KillMonsterDb, error) {
	sql := fmt.Sprintf(`select %s from %s where stageId=%d`, killMonsterFields, killMonsterDb.TableName(), stageId)
	var data KillMonsterDb
	err := this.DbMap().SelectOne(&data, sql)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
