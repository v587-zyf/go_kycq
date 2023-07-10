package constAuction

import "time"

const (
	OnAuction = 1 // 竞拍中
	ItemSold  = 2 // 已售出
	PassIn    = 3 // 流拍

	WorldAuction = 1 //世界拍卖行
	GuildAuction = 2

	AuctionItemOk          = 1
	AuctionItemErrDb       = 2  // 操作数据库失败
	AuctionItemErrConf     = 3  // 没找到配置
	AuctionItemErrNoItem   = 4  // 背包中没有此物品
	AuctionItemErrCannot   = 5  // 不能拍卖
	AuctionItemNotInPeriod = 6  // 不在上架时间内
	AuctionMaxLimitDaily   = 7  //已达每日收入上限
	AuctionTogetherLimit   = 8  //最多同时上架xx个物品
	AuctionUpCountErr      = 9  //上架数量 <= 0
	AuctionSetPriceErr     = 10 //定价超过浮动上限
	AuctionOverCount       = 11 //超过上架数量上限

	AuctionSrcPlayer      = 1 // 玩家拍卖
	AuctionSrcGuildPassIn = 2 // 门派流拍到世界拍卖的
	AuctionSrcGuild       = 3 // 门派拍卖
	AuctionSrcSystem      = 4 // 系统拍卖

	DbChanSize = 50 // 数据库操作buffer channel size

	OpUpdate = 1
	OpInsert = 2

	CodeSuccess                     = 1 // 竞价成功
	CodeReachBuyNowPrice            = 2 // 一口价
	CodeItemSold                    = 3 // 道具已售出
	CodePriceBeyond                 = 4 // 竞价被超越
	CodeMiNoEnough                  = 5
	CodeUnknown                     = 6  // 配置不存在
	CodeNotInAuctionTime            = 7  // 不在竞拍时间内
	CodeBidPriceTheHighest          = 8  // 您当前的竞价最高
	CodeBidPriceLowerThanStartPrice = 9  // 竞价低于起拍价
	CodeCanNotBidMyselfItem         = 10 // 不能竞价自己上架的物品

	AuctionDataSaveDuration   = 15 * 86400      // 已售出物品数据保留时间(秒)
	AuctionDataUpdateDuartion = 5 * time.Minute // 拍卖数据库更新间隔
	BidSuccessAddDuration     = 120             // 竞价成功时增加的竞拍时间（秒）

	AuctionResultPassIn  = 0 // 流拍
	AuctionResultSuccess = 1 // 成功

	//公会拍卖行 掉落状态
	DropWorldLeader = -1 //世界首领掉落
	DropShaBake     = -2 //沙巴克掉落

)
