package treasureShop

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constShop"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"fmt"
	"time"
)

func NewTreasureShopManager(m managersI.IModule) *TreasureShop {
	return &TreasureShop{IModule: m}
}

type TreasureShop struct {
	util.DefaultModule
	managersI.IModule
}

/**
 *  @Description: 多宝阁加载
 *  @param user
 *  @return *pb.TreasureShopLoadAck
 */
func (this *TreasureShop) Load(user *objs.User) *pb.TreasureShopLoadAck {
	if this.checkOpen(user) {
		this.AutoRefreshShop(user, false)
	}
	userShopInfo := user.TreasureShop
	return &pb.TreasureShopLoadAck{
		BuyNum:      int32(userShopInfo.BuyNum),
		RefreshFree: userShopInfo.RefreshFree,
		RefreshTime: int64(userShopInfo.RefreshTime),
		Shop:        builder.BuildTreasureShop(userShopInfo.Shop),
		Car:         builder.BuildTreasureCar(userShopInfo.Car),
		EndTime:     this.getEndTime(user).Unix(),
	}
}

/**
 *  @Description: 多宝阁自动刷新商品
 *  @param user
 *  @param sendMsg
 */
func (this *TreasureShop) AutoRefreshShop(user *objs.User, sendMsg bool) {
	timeNow := time.Now()
	userShopInfo := user.TreasureShop
	if userShopInfo.RefreshTime > int(timeNow.Unix()) || !this.checkOpen(user) {
		return
	}
	shopMap := make(map[int]int)
	randShopItem := gamedb.RandTreasureShopItem()
	for id := range randShopItem {
		shopMap[id] = constShop.SHOP_ADD_NO
	}
	userShopInfo.Shop = shopMap
	m, _ := time.ParseDuration(fmt.Sprintf("%dm", gamedb.GetConf().TreasureTime))
	userShopInfo.RefreshTime = int(timeNow.Add(m).Unix())
	if sendMsg {
		this.GetUserManager().SendMessage(user, &pb.TreasureShopRefreshNtf{
			Shop:        builder.BuildTreasureShop(userShopInfo.Shop),
			RefreshTime: int64(userShopInfo.RefreshTime),
			Car:         builder.BuildTreasureCar(userShopInfo.Car),
			RefreshFree: userShopInfo.RefreshFree,
		}, true)
	}
	user.Dirty = true
}

/**
 *  @Description: 购物车物品变更
 *  @param user
 *  @param req
 *  @param ack
 *  @return error
 */
func (this *TreasureShop) CarChange(user *objs.User, shopId int, isAdd bool, ack *pb.TreasureShopCarChangeAck) error {
	if !this.checkOpen(user) {
		return gamedb.ERRNOTOPEN
	}
	userShopInfo := user.TreasureShop
	shopAddFlag, shopHasFlag := userShopInfo.Shop[shopId]
	if isAdd {
		if shopHasFlag {
			if shopAddFlag == constShop.SHOP_ADD_YES {
				return gamedb.ERRREPEATBUY
			} else {
				userShopInfo.Shop[shopId] = constShop.SHOP_ADD_YES
			}
		}
		userShopInfo.Car[shopId] += 1
	} else {
		carNum, ok := userShopInfo.Car[shopId]
		if !ok {
			return gamedb.ERRPARAM
		}
		carNum--
		if carNum < 1 {
			delete(userShopInfo.Car, shopId)
		} else {
			userShopInfo.Car[shopId] = carNum
		}
		if shopHasFlag {
			userShopInfo.Shop[shopId] = constShop.SHOP_ADD_NO
		}
	}
	user.Dirty = true
	ack.Car = builder.BuildTreasureCar(userShopInfo.Car)
	ack.ShopId = int32(shopId)
	ack.IsAdd = isAdd
	ack.Shop = builder.BuildTreasureShop(userShopInfo.Shop)
	return nil
}

/**
 *  @Description: 手动刷新
 *  @param user
 *  @param op
 *  @return error
 */
func (this *TreasureShop) RefreshShop(user *objs.User, op *ophelper.OpBagHelperDefault) error {
	if !this.checkOpen(user) {
		return gamedb.ERRNOTOPEN
	}
	if !user.TreasureShop.RefreshFree {
		cost := gamedb.GetConf().TreasureCost
		if err := this.GetBag().Remove(user, op, cost.ItemId, cost.Count); err != nil {
			return err
		}
	} else {
		user.TreasureShop.RefreshFree = false
	}
	user.TreasureShop.RefreshTime = 0
	this.AutoRefreshShop(user, true)
	return nil
}

/**
 *  @Description: 多宝阁购买
 *  @param user
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *TreasureShop) Buy(user *objs.User, shop []int32, op *ophelper.OpBagHelperDefault, ack *pb.TreasureShopBuyAck) error {
	if !this.checkOpen(user) {
		return gamedb.ERRNOTOPEN
	}
	userShopInfo := user.TreasureShop
	if len(userShopInfo.Car) < 1 || len(shop) < 1 || len(shop)%2 != 0 {
		return gamedb.ERRPARAM
	}
	if userShopInfo.BuyNum >= gamedb.GetConf().TreasureBuyTime {
		return gamedb.ERRBUYTIMESLIMIT
	}
	priceItem, priceCount := 0, 0
	addMap := make(map[int]int)
	delMap := make(map[int]int)
	for i := 0; i < len(shop); i += 2 {
		shopId := int(shop[i])
		buyNum := int(shop[i+1])
		if carNum, ok := userShopInfo.Car[shopId]; !ok || carNum < buyNum {
			return gamedb.ERRPARAM
		}
		shopCfg := gamedb.GetTreasureShopTreasureShopCfg(shopId)
		addMap[shopCfg.ItemId] += shopCfg.Count * buyNum
		priceItem = shopCfg.Price.ItemId
		priceCount += shopCfg.Price.Count * buyNum
		delMap[shopId] = buyNum
	}
	discount := gamedb.GetTreasureDiscount(priceCount)
	if err := this.GetBag().Remove(user, op, priceItem, priceCount-discount); err != nil {
		return err
	}
	addItems := make(gamedb.ItemInfos, 0)
	for itemId, count := range addMap {
		addItems = append(addItems, &gamedb.ItemInfo{
			ItemId: itemId,
			Count:  count,
		})
	}
	this.GetBag().AddItems(user, addItems, op)
	userShopInfo.BuyNum++
	for shopId, carNum := range delMap {
		n := userShopInfo.Car[shopId] - carNum
		if n < 1 {
			delete(userShopInfo.Car, shopId)
		} else {
			userShopInfo.Car[shopId] = n
		}
	}

	user.Dirty = true
	ack.BuyNum = int32(userShopInfo.BuyNum)
	ack.Car = builder.BuildTreasureCar(userShopInfo.Car)
	ack.Shop = builder.BuildTreasureShop(userShopInfo.Shop)
	ack.Goods = op.ToChangeItems()
	return nil
}

func (this *TreasureShop) checkOpen(user *objs.User) bool {
	flag := true
	condition := gamedb.GetFunctionFunctionCfg(pb.FUNCTIONID_TREASURE_SHOP).Condition
	for k, v := range condition {
		if _, check := this.GetCondition().CheckBySlice(user, -1, []int{k, v}); !check {
			flag = false
			break
		}
	}
	if time.Now().Unix() > this.getEndTime(user).Unix() {
		flag = false
	}
	return flag
}

func (this *TreasureShop) getEndTime(user *objs.User) time.Time {
	serverOpenTime := this.GetSystem().GetServerOpenTimeByServerId(user.ServerId)
	sustainTime := gamedb.GetConf().TreasureTime1
	addNum := gamedb.GetConf().TreasureTime2 - 1
	if addNum < 1 {
		addNum = 0
	}
	addDay := sustainTime[0] + addNum
	addHour, _ := time.ParseDuration(fmt.Sprintf(`%dh`, sustainTime[1]))
	addMinute, _ := time.ParseDuration(fmt.Sprintf(`%dm`, sustainTime[2]))
	endTime := serverOpenTime.AddDate(0, 0, addDay).Add(addHour).Add(addMinute)
	endTime = common.GetDate(endTime)
	return endTime
}
