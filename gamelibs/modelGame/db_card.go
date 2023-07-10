package modelGame

import (
	"cqserver/gamelibs/model"
	"cqserver/golibs/dbmodel"
	"fmt"
	"github.com/astaxie/beego/orm"
	"gopkg.in/gorp.v1"
)

type Card struct {
	Id          int    `db:"id" orm:"pk;auto"`
	UserId      int    `db:"userId"orm:"comment(抽卡玩家id)"`
	ItemId      int    `db:"itemId" orm:"comment(抽中物品id)"`
	Count       int    `db:"count" orm:"comment(数量)"`
	ItemQuality int    `db:"itemQuality" orm:"comment(物品品质)"`
	Season      int    `db:"season" orm:"comment(周期)"`
	DrawType    int    `db:"drawType" orm:"comment(抽卡类型:1-单抽 10-十连抽)"`
	DrawTime    int    `db:"drawTime" orm:"comment(抽卡时间)"` //抽卡时间
	NickName    string `db:"nickName" orm:"comment(抽卡玩家昵称)"`
	ExpireTime  int64  `db:"expireTime" orm:"comment(过期时间)"` //数据删除时间
}

func (this *Card) TableName() string {
	return "card"
}

type CardModel struct {
	dbmodel.CommonModel
}

var (
	cardModel  = &CardModel{}
	cardFields = model.GetAllFieldsAsString(Card{})
)

func init() {
	dbmodel.Register(model.DB_SERVER, cardModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(Card{}, "card").SetKeys(true, "Id")
		orm.RegisterModelForAlias(model.DB_SERVER, new(Card))
	})
}

func GetCardModel() *CardModel {
	return cardModel
}

func (this *CardModel) GetAllDrawCardInfos(season, limit, before, after int) ([]*Card, error) {
	var cards []*Card
	_, err := this.DbMap().Select(&cards, fmt.Sprintf("SELECT %s FROM card where season = %v and itemQuality >= %v and itemQuality <= %v ORDER BY drawTime DESC LIMIT %v", cardFields, season, before, after, limit))
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func (this *CardModel) GetMyDrawCardInfos(season, userId, limit int) ([]*Card, error) {
	var cards []*Card
	_, err := this.DbMap().Select(&cards, fmt.Sprintf("SELECT %s FROM card WHERE season = %v and userId = %v  ORDER BY drawTime DESC LIMIT %v", cardFields, season, userId, limit))
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func (this *CardModel) DeleteExpiredItem(ts int64) error {
	sql := fmt.Sprintf("delete from card where expireTime > 0 AND expireTime <= %d", ts)
	_, err := this.DbMap().Exec(sql)
	return err
}

func (this *CardModel) DeleteAllItem() error {
	sql := fmt.Sprintf("delete from card where 1=1")
	_, err := this.DbMap().Exec(sql)
	return err
}
