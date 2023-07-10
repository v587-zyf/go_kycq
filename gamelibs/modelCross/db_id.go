package modelCross

import (
	"cqserver/gamelibs/model"
	"cqserver/golibs/logger"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"sync"
	"time"

	"cqserver/golibs/dbmodel"

	"gopkg.in/gorp.v1"
)

type IdSeq struct {
	curId int
	maxId int
	Name  string
	mu    sync.Mutex
}

//next 得到下一个id
func (this *IdSeq) Next() (int, error) {
	this.mu.Lock()
	defer this.mu.Unlock()
	if this.curId > this.maxId || this.maxId == 0 {
		number, step, err := GetIdModel().GetIds(this.Name)
		if err != nil {
			return 0, errors.New(fmt.Sprintf("generate new id form %s failed: %s", this.Name, err.Error()))
		}
		this.curId = number * step
		this.maxId = this.curId + step - 1
	}

	returnId := this.curId
	this.curId++
	return int(returnId), nil
}

type Id struct {
	Id          int       `db:"id"`
	Number      int       `db:"number"`
	Name        string    `db:"name"`
	Step        int       `db:"step"`
	CreatedTime time.Time `db:"createdtime"`
	UpdateTime  time.Time `db:"modifiedtime"`
}

type IdModel struct {
	dbmodel.CommonModel
}

var (
	idModel  = &IdModel{}
	idFields = model.GetAllFieldsAsString(Id{})
)

func init() {
	dbmodel.Register(model.DB_ACCOUNT, idModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(Id{}, "ids").SetKeys(true, "id")
	})
}

/////////////////////////////////////////
// ids table
/////////////////////////////////////////

func GetIdModel() *IdModel {
	return idModel
}

func (this *IdModel) GetIds(name string) (number int, step int, err error) {
	//为了是存储过程能得到返回值，需要修改mysql driver中的packet.go文件
	//在writeAuthPacket函数中clientFlags增加 clientMultiStatements | clientMultiResults
	//这里不再用存储过程

	//预防invalid connection
	this.DbMap().Db.Ping()
	tx, err := this.DbMap().Db.Begin()
	if err != nil {
		return 0, 0, err
	}

	row := tx.QueryRow("select number, step from ids where name = ? for update", name)
	err = row.Scan(&number, &step)
	if err != nil {
		tx.Rollback()
		return 0, 0, err
	}
	number++
	_, err = tx.Exec("update ids set number = ? where name = ?", number, name)
	if err != nil {
		tx.Rollback()
		return 0, 0, err
	}
	tx.Commit()

	return
}

func (this *IdModel) Create(idSeq *Id) error {
	return this.DbMap().Insert(idSeq)
}

func (this *IdModel) CheckIdsCfg(ids []string) {
	logs.Info("CheckIdsCfg  ids:%v ", ids)
	for _, idName := range ids {
		num, _, _ := this.GetIds(idName)
		if num == 0 {
			logger.Info("idName:%v", idName)
			idInfo := &Id{Number: 1, Name: idName, Step: 1000, CreatedTime: time.Now(), UpdateTime: time.Now()}
			this.Create(idInfo)
		}
	}
}
