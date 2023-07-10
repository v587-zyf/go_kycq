package modelGame

import (
	"cqserver/gamelibs/model"
	"cqserver/golibs/dbmodel"
	"fmt"
	"github.com/astaxie/beego/orm"
	"gopkg.in/gorp.v1"
)

type AuctionBid struct {
	Id             int   `db:"id"`
	UserId         int   `db:"userId" orm:"comment(竞拍的玩家id)"`
	AuctionId      int   `db:"auctionId" orm:"comment(门派or世界拍卖行对应表id)"`
	AuctionType    int   `db:"auctionType" orm:"comment(拍卖行类型 1:世界拍卖行 2:门派拍卖行)" `
	ItemId         int   `db:"itemId" orm:"comment(物品id)"`
	ItemCount      int   `db:"itemCount" orm:"comment(物品数量)"`
	FirstBidTime   int64 `db:"firstBidTime" orm:"comment(第一次竞拍的时间)"`
	FinallyBidTime int64 `db:"finallyBidTime" orm:"comment(最后竞拍的时间)"`
	Status         int   `db:"status" orm:"comment(竞拍状态 1:竞拍中 2:已售出 3:流拍)"`
	FinalBidUserId int   `db:"finalBidUserId" orm:"comment(最后竞拍的玩家id)"`
	ExpireTime     int   `db:"expireTime" orm:"comment(数据删除时间，物品售出或者流拍后数据保留7天)"`
}

type AuctionBidModel struct {
	dbmodel.CommonModel
}

var (
	auctionBidModel       = &AuctionBidModel{}
	auctionBidModelFields = dbmodel.GetAllFieldsAsString(AuctionBid{})
)

func init() {
	dbmodel.Register(model.DB_SERVER, auctionBidModel, func(dbMap *gorp.DbMap) {
		dbMap.AddTableWithName(AuctionBid{}, "auction_bid").SetKeys(false, "id")
		orm.RegisterModelForAlias(model.DB_SERVER, new(AuctionBid))
	})
}

func GetAuctionBidModel() *AuctionBidModel {
	return auctionBidModel
}

func (this *AuctionBid) TableName() string {
	return "auction_bid"
}

func (this *AuctionBidModel) GetUserBidItems(userId int) ([]*AuctionBid, error) {
	var briefInfos []*AuctionBid

	sqlStr := fmt.Sprintf("SELECT %s FROM auction_bid WHERE userId = %v", auctionBidModelFields, userId)

	_, err := this.DbMap().Select(&briefInfos, sqlStr)
	if err != nil {
		return nil, err
	}

	return briefInfos, nil
}

func (this *AuctionBidModel) GetUserBidIdByUserIdAndAuctionId(userId, auctionId, auctionType int) (*AuctionBid, error) {
	var briefInfos *AuctionBid
	sqlStr := fmt.Sprintf("SELECT %s FROM auction_bid WHERE  userId = %v AND auctionId = %v  AND auctionType = %v", auctionBidModelFields, userId, auctionId, auctionType)
	err := this.DbMap().SelectOne(&briefInfos, sqlStr)
	if err != nil {
		return nil, err
	}

	return briefInfos, nil
}

func (this *AuctionBidModel) GetUserBidIdByUserIdsAndAuctionId(auctionId, auctionType int) ([]*AuctionBid, error) {
	var briefInfos []*AuctionBid
	sqlStr := fmt.Sprintf("SELECT %s FROM auction_bid WHERE   auctionId = %v  AND auctionType = %v ", auctionBidModelFields, auctionId, auctionType)
	_, err := this.DbMap().Select(&briefInfos, sqlStr)
	if err != nil {
		return nil, err
	}
	return briefInfos, nil
}

func (this *AuctionBidModel) GetAllUserBidItems() ([]*AuctionBid, error) {
	var briefInfos []*AuctionBid

	sqlStr := fmt.Sprintf("SELECT %s FROM auction_bid WHERE 1 = 1", auctionBidModelFields)

	_, err := this.DbMap().Select(&briefInfos, sqlStr)
	if err != nil {
		return nil, err
	}

	return briefInfos, nil
}

func (this *AuctionBidModel) DeleteByAuctionId(auctionId int) error {
	sqlStr := fmt.Sprintf("DELETE FROM auction_bid WHERE auctionId = %v", auctionId)
	_, err := this.DbMap().Exec(sqlStr)
	return err
}

func (this *AuctionBidModel) GetAllUserBidItemsByUserId(userId int) ([]*AuctionBid, error) {
	var briefInfos []*AuctionBid

	sqlStr := fmt.Sprintf("SELECT %s FROM auction_bid WHERE userId = %v ORDER BY ExpireTime  limit 30", auctionBidModelFields, userId)

	_, err := this.DbMap().Select(&briefInfos, sqlStr)
	if err != nil {
		return nil, err
	}

	return briefInfos, nil
}

func (this *AuctionBidModel) UpdateStatus(state, auctionId, auctionType int) error {
	_, err := this.DbMap().Exec("update auction_bid set status = ? where auctionId = ? and auctionType = ?", state, auctionId, auctionType)
	return err
}
