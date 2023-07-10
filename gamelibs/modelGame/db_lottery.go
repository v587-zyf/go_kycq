package modelGame

import (
	"cqserver/gamelibs/model"
	"cqserver/golibs/dbmodel"
	"fmt"
	"github.com/astaxie/beego/orm"
	"gopkg.in/gorp.v1"
	"time"
)

type Lottery struct {
	Id          int `db:"id" orm:"pk;auto"`
	Day         int `db:"day"`
	UserId      int `db:"userId" orm:"comment(玩家id)"`
	AwardNumber int `db:"awardNumber" orm:"comment(奖号)"`
	Share       int `db:"share" orm:"share(购买了多少份)"`
}

func (this *Lottery) TableName() string {
	return "lottery"
}

type LotteryModel struct {
	dbmodel.CommonModel
}

var (
	lotteryModel  = &LotteryModel{}
	lotteryFields = model.GetAllFieldsAsString(Lottery{})
)

func init() {
	dbmodel.Register(model.DB_SERVER, lotteryModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(Lottery{}, "lottery").SetKeys(true, "Id")
		orm.RegisterModelForAlias(model.DB_SERVER, new(Lottery))
	})
}

func GetLotteryModel() *LotteryModel {
	return lotteryModel
}

func (this *LotteryModel) GetAllLotteryInfos() ([]*Lottery, error) {
	var lotterys []*Lottery
	_, err := this.DbMap().Select(&lotterys, fmt.Sprintf("SELECT %s FROM lottery where day = %d", lotteryFields, time.Now().Day()))
	if err != nil {
		return nil, err
	}
	return lotterys, nil
}

func (this *LotteryModel) DeleteAllItem() error {
	sql := fmt.Sprintf("delete from lottery where 1=1")
	_, err := this.DbMap().Exec(sql)
	return err
}
