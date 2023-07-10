package modelGame

import (
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constAuction"
	"cqserver/golibs/dbmodel"
	"fmt"
	"github.com/astaxie/beego/orm"
	"gopkg.in/gorp.v1"
)

// 玩家信息
type GuildAuctionItem struct {
	Id              int            `db:"id"`
	ItemId          int            `db:"itemId"`
	ItemCount       int            `db:"itemCount"`
	AuctionTime     int64          `db:"auctionTime" orm:"comment(上架时间)"`
	AuctionDuration int            `db:"auctionDuration" orm:"comment(持续时间)"`
	LastBidPlayerId int            `db:"lastBidPlayerId" orm:"comment(上一次竞拍的玩家)"`
	LastBidPrice    int            `db:"lastBidPrice" orm:"comment(上一次竞拍的价格)"`
	NowBidPlayerId  int            `db:"nowBidPlayerId"orm:"comment(当前竞拍的价格)"`
	NowBidPrice     int            `db:"nowBidPrice" orm:"comment(当前竞拍的价格)"`
	AllBidUsers     model.IntSlice `db:"allBidUsers" orm:"comment(所有竞拍的玩家)"`
	Status          int            `db:"status" orm:"comment(竞拍状态 1:竞拍中 2:已售出 3:流拍)"`
	AuctionGuild    int            `db:"auctionGuild"orm:"comment(可以竞拍的门派id)"`
	ExpireTime      int            `db:"expireTime" orm:"comment(过期时间)"` //数据删除时间，物品售出或者流拍后数据保留15天
	AuctionType     int            `db:"auctionType" orm:"comment(所属拍卖行类型,1:世界拍卖行 2:门派拍卖行)"`
	DropState       int            `db:"dropState" orm:"comment(哪个模块掉落的物品,1:世界拍卖行 2:门派拍卖行)"`
	CanGetRedAward  model.IntSlice `db:"canGetRedAward" orm:"comment(所有分红的玩家)"`
}

type GuildAuctionItemModel struct {
	dbmodel.CommonModel
}

var (
	guildAuctionItemModel  = &GuildAuctionItemModel{}
	guildAuctionItemFields = dbmodel.GetAllFieldsAsString(GuildAuctionItem{})
)

func init() {

	dbmodel.Register(model.DB_SERVER, guildAuctionItemModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(GuildAuctionItem{}, "guild_auction_item").SetKeys(false, "id")
		orm.RegisterModelForAlias(model.DB_SERVER, new(GuildAuctionItem))
	})
}

func GetGuildAuctionItemModel() *GuildAuctionItemModel {
	return guildAuctionItemModel
}

func (this *GuildAuctionItem) TableName() string {
	return "guild_auction_item"
}

func (this *GuildAuctionItemModel) GetAuctionItemsByGuildId(guildId, artifactType int) ([]*GuildAuctionItem, error) {
	sql := fmt.Sprintf("select %s from guild_auction_item where status = %v and guildId = %v and artifactType = %v",
		guildAuctionItemFields, constAuction.OnAuction, guildId, artifactType)
	var allItems []*GuildAuctionItem
	_, err := this.DbMap().Select(&allItems, sql)
	if err != nil {
		return nil, err
	}
	return allItems, nil
}

func (this *GuildAuctionItemModel) GetAllAuctionItems() ([]*GuildAuctionItem, error) {
	sql := fmt.Sprintf("select %s from guild_auction_item where status = %d", guildAuctionItemFields, constAuction.OnAuction)
	var allItems []*GuildAuctionItem
	_, err := this.DbMap().Select(&allItems, sql)
	if err != nil {
		return nil, err
	}
	return allItems, nil
}

func (this *GuildAuctionItemModel) DeleteExpiredItem(ts int64) error {
	sql := fmt.Sprintf("DELETE FROM guild_auction_item where expireTime > 0 AND expireTime <= %d", ts)
	_, err := this.DbMap().Exec(sql)
	return err
}

func (this *GuildAuctionItemModel) GetMaxId() (int, error) {
	sql := fmt.Sprintf("select MAX(id) from guild_auction_item")
	var maxId int
	err := this.DbMap().SelectOne(&maxId, sql)
	if err != nil {
		return -1, err
	}
	return maxId, nil
}

func (this *GuildAuctionItemModel) GetAuctionItemByAuctionId(id int) (*GuildAuctionItem, error) {
	sql := fmt.Sprintf("select %s from guild_auction_item where id = %d ", guildAuctionItemFields, id)
	var allItems *GuildAuctionItem
	err := this.DbMap().SelectOne(&allItems, sql)
	if err != nil {
		return nil, err
	}
	return allItems, nil
}
