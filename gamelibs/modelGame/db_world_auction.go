package modelGame

import (
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constAuction"
	"cqserver/golibs/dbmodel"
	"cqserver/golibs/logger"
	"fmt"
	"github.com/astaxie/beego/orm"
	"gopkg.in/gorp.v1"
)

// 玩家信息
type AuctionItem struct {
	Id              int            `db:"id"`
	ItemId          int            `db:"itemId"`
	ItemCount       int            `db:"itemCount"`
	AuctionUserId   int            `db:"auctionUserId" orm:"comment(上架玩家id)"`
	AuctionTime     int            `db:"auctionTime"orm:"comment(上架时间)"`
	AuctionDuration int            `db:"auctionDuration" orm:"comment(持续时间)"`
	LastBidPlayerId int            `db:"lastBidPlayerId" orm:"comment(上一次竞拍的玩家)"`
	LastBidPrice    int            `db:"lastBidPrice" orm:"comment(上一次竞拍的价格)"`
	NowBidPlayerId  int            `db:"nowBidPlayerId" orm:"comment(当前竞拍的玩家)"`
	PutAwayPrice    int            `db:"putAwayPrice" orm:"comment(玩家上架价格)"`
	NowBidPrice     int            `db:"nowBidPrice" orm:"comment(当前竞拍的价格)"`
	AllBidUsers     model.IntSlice `db:"allBidUsers" orm:"comment(所有竞拍过的玩家)"`
	Status          int            `db:"status" orm:"comment(竞拍状态 1:竞拍中 2:已售出 3:流拍)"`
	AuctionSrc      int            `db:"auctionSrc" orm:"comment(玩家上架还是系统上架  1:是玩家上架)"`
	ExpireTime      int            `db:"expireTime" orm:"comment(数据删除时间，物品售出或者流拍后数据保留7天)"`
	AuctionType     int            `db:"auctionType" orm:"comment(所属拍卖行类型,1:世界拍卖行 2:门派拍卖行)"`
}

type AuctionItemModel struct {
	dbmodel.CommonModel
}

var (
	auctionItemModel  = &AuctionItemModel{}
	auctionItemFields = dbmodel.GetAllFieldsAsString(AuctionItem{})
)

func init() {
	dbmodel.Register(model.DB_SERVER, auctionItemModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(AuctionItem{}, "auction_item").SetKeys(false, "id")
		orm.RegisterModelForAlias(model.DB_SERVER, new(AuctionItem))
	})
}

func GetAuctionItemModel() *AuctionItemModel {
	return auctionItemModel
}

func (this *AuctionItem) TableName() string {
	return "auction_item"
}

func (this *AuctionItemModel) GetAllAuctionItemByType(artifactType int) ([]*AuctionItem, error) {
	sql := fmt.Sprintf("select %s from auction_item where status = %d and artifactType = %d ORDER BY auctionTime", auctionItemFields, constAuction.OnAuction, artifactType)
	var allItems []*AuctionItem
	_, err := this.DbMap().Select(&allItems, sql)
	if err != nil {
		return nil, err
	}
	return allItems, nil
}

func (this *AuctionItemModel) GetAuctionItem() ([]*AuctionItem, error) {
	sql := fmt.Sprintf("select %s from auction_item where status = %d ORDER BY auctionTime", auctionItemFields, constAuction.OnAuction)
	var allItems []*AuctionItem
	_, err := this.DbMap().Select(&allItems, sql)
	if err != nil {
		return nil, err
	}
	return allItems, nil
}

func (this *AuctionItemModel) DeleteExpiredItem(ts int64) error {
	sql := fmt.Sprintf("DELETE FROM auction_item where expireTime > 0 AND expireTime <= %d", ts)
	_, err := this.DbMap().Exec(sql)

	return err
}

func (this *AuctionItemModel) GetMaxId() (int, error) {
	sql := fmt.Sprintf("select MAX(id) from auction_item")
	var maxId int
	err := this.DbMap().SelectOne(&maxId, sql)
	if err != nil {
		return -1, err
	}
	return maxId, nil
}

func (this *AuctionItemModel) GetUserNotSellAuction(userId int) []int {
	sql := fmt.Sprintf("select itemId from auction_item where status = %d and auctionUserId = %d", constAuction.OnAuction, userId)
	var ids []int
	_, err := this.DbMap().Select(&ids, sql)
	if err != nil {
		logger.Error("GetUserNotSellAuction|error:%v, userId:%d", err, userId)
		return []int{}
	}
	return ids
}

func (this *AuctionItemModel) GetAuctionItemByAuctionUser(userId, limit int) ([]*AuctionItem, error) {
	sql := fmt.Sprintf("select %s from auction_item where  auctionUserId = %v and status != 1 ORDER BY auctionTime limit %v", auctionItemFields, userId, limit)
	var allItems []*AuctionItem
	_, err := this.DbMap().Select(&allItems, sql)
	if err != nil {
		return nil, err
	}
	return allItems, nil
}

func (this *AuctionItemModel) GetAuctionItemByAuctionUserInBid(userId int) ([]*AuctionItem, error) {
	sql := fmt.Sprintf("select %s from auction_item where  auctionUserId = %v and status = 1 ORDER BY auctionTime", auctionItemFields, userId)
	var allItems []*AuctionItem
	_, err := this.DbMap().Select(&allItems, sql)
	if err != nil {
		return nil, err
	}
	return allItems, nil
}

func (this *AuctionItemModel) GetAuctionItemByAuctionId(id int) (*AuctionItem, error) {
	sql := fmt.Sprintf("select %s from auction_item where id = %d ", auctionItemFields, id)
	var allItems *AuctionItem
	err := this.DbMap().SelectOne(&allItems, sql)
	if err != nil {
		return nil, err
	}
	return allItems, nil
}
