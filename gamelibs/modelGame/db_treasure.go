package modelGame

import (
	"cqserver/gamelibs/model"
	"cqserver/golibs/dbmodel"
	"fmt"
	"github.com/astaxie/beego/orm"
	"gopkg.in/gorp.v1"
)

type Treasure struct {
	Id          int    `db:"id" orm:"pk;auto"`
	UserId      int    `db:"userId"orm:"comment(寻龙探宝玩家id)"`
	ItemId      int    `db:"itemId" orm:"comment(抽中物品id)"`
	Count       int    `db:"count" orm:"comment(数量)"`
	ItemQuality int    `db:"itemQuality" orm:"comment(物品品质)"`
	Season      int    `db:"season" orm:"comment(周期)"`
	DrawTime    int    `db:"drawTime" orm:"comment(抽卡时间)"` //抽卡时间
	NickName    string `db:"nickName" orm:"comment(抽卡玩家昵称)"`
	ExpireTime  int64  `db:"expireTime" orm:"comment(过期时间)"` //数据删除时间
}

func (this *Treasure) TableName() string {
	return "treasure"
}

type TreasureModel struct {
	dbmodel.CommonModel
}

var (
	treasureModel  = &TreasureModel{}
	treasureFields = model.GetAllFieldsAsString(Treasure{})
)

func init() {
	dbmodel.Register(model.DB_SERVER, treasureModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(Treasure{}, "treasure").SetKeys(true, "Id")
		orm.RegisterModelForAlias(model.DB_SERVER, new(Treasure))
	})
}

func GetTreasureModel() *TreasureModel {
	return treasureModel
}

func (this *TreasureModel) GetAllTreasureInfos(season, limit, before, after int) ([]*Treasure, error) {
	var treasures []*Treasure
	_, err := this.DbMap().Select(&treasures, fmt.Sprintf("SELECT %s FROM treasure where season = %v and itemQuality >= %v and itemQuality <= %v ORDER BY drawTime DESC LIMIT %v", treasureFields, season, before, after, limit))
	if err != nil {
		return nil, err
	}
	return treasures, nil
}

func (this *TreasureModel) GetMyTreasureInfos(season, userId, limit int) ([]*Treasure, error) {
	var treasures []*Treasure
	_, err := this.DbMap().Select(&treasures, fmt.Sprintf("SELECT %s FROM treasure WHERE season = %v and userId = %v  ORDER BY drawTime DESC LIMIT %v", treasureFields, season, userId, limit))
	if err != nil {
		return nil, err
	}
	return treasures, nil
}

func (this *TreasureModel) DeleteExpiredItem(ts int64) error {
	sql := fmt.Sprintf("delete from treasure where expireTime > 0 AND expireTime <= %d", ts)
	_, err := this.DbMap().Exec(sql)
	return err
}

func (this *TreasureModel) DeleteAllItem() error {
	sql := fmt.Sprintf("delete from treasure where 1=1")
	_, err := this.DbMap().Exec(sql)
	return err
}
