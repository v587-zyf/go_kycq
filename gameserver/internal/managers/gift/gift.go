package gift

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"math/rand"
	"strconv"
	"time"
)

const (
	RANDOM = 1
	CHOOSE = 2
	All    = 3
)

type GiftManager struct {
	util.DefaultModule
	managersI.IModule
}

func NewGiftManager(module managersI.IModule) *GiftManager {
	f := &GiftManager{IModule: module}
	return f
}

func (this *GiftManager) Online(user *objs.User) {
	this.LimitedMerge(user)
}

func (this *GiftManager) OpenGift(user *objs.User, types, giftItemId, num int, chooseItemId []int, op *ophelper.OpBagHelperDefault) error {
	if types < 0 || giftItemId < 0 || num <= 0 {
		return gamedb.ERRPARAM
	}

	logger.Debug("OpenGift  types:%v, giftItemId:%v, chooseItemId:%v  num:%v", types, giftItemId, chooseItemId, num)
	itemCfg := gamedb.GetItemBaseCfg(giftItemId)
	giftCfg := gamedb.GetGiftGiftCfg(giftItemId)
	if itemCfg == nil || giftCfg == nil {
		return gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg("item" + strconv.Itoa(giftItemId))
	}

	enough, _ := this.GetBag().HasEnough(user, giftItemId, num)
	if !enough {
		return gamedb.ERRNOTENOUGHGOODS
	}

	_ = this.GetBag().Remove(user, op, giftItemId, num)
	allItems := make(map[int]int)
	infos := make(gamedb.ItemInfos, 0)
	if types == RANDOM {
		if giftCfg.Type != RANDOM {
			return gamedb.ERRTYPEERR
		}
		for i := 1; i <= num; i++ {
			rand.Seed(time.Now().UnixNano() + int64(i))
			itemInfo := this.ItemInfo(giftCfg.Reward)
			if itemInfo != nil {
				allItems[itemInfo.ItemId] += itemInfo.Count
			}
		}
		logger.Debug("RANDOM allItems:%v", allItems)
		for itemId, count := range allItems {
			infos = append(infos, &gamedb.ItemInfo{ItemId: itemId, Count: count})
		}
	} else if types == CHOOSE {
		if giftCfg.Type != CHOOSE {
			return gamedb.ERRTYPEERR
		}
		if len(chooseItemId) != giftCfg.Choose {
			logger.Error("len(chooseItemId):%v != giftCfg.Choose:%v   giftCfgId:%v", len(chooseItemId), giftCfg.Choose, giftCfg.Id)
			return gamedb.ERRTYPEERR
		}
		for _, itemInfo := range giftCfg.Reward {
			if len(itemInfo) != 2 {
				continue
			}
			for _, itemId := range chooseItemId {
				if itemInfo[0] == itemId {
					infos = append(infos, &gamedb.ItemInfo{ItemId: itemInfo[0], Count: itemInfo[1] * num})
				}
			}
		}
	} else if types == All {
		if giftCfg.Type != All {
			return gamedb.ERRTYPEERR
		}
		if len(giftCfg.Reward) > 0 {
			for _, info := range giftCfg.Reward {
				infos = append(infos, &gamedb.ItemInfo{ItemId: info[0], Count: info[1] * num})
			}
		}
	}
	logger.Debug("userId:%v  infos:%v", user.Id, infos)
	for _, info := range infos {
		baseCfg := gamedb.GetItemBaseCfg(info.ItemId)
		if baseCfg == nil {
			continue
		}
		if baseCfg.CountLimit == 1 {
			for i := 0; i < info.Count; i++ {
				this.GetBag().AddItem(user, op, info.ItemId, 1)
			}
		} else {
			this.GetBag().AddItem(user, op, info.ItemId, info.Count)
		}
	}
	return nil
}

func (this *GiftManager) ItemInfo(randomItems gamedb.IntSlice2) *gamedb.ItemInfo {
	totalRate := 0
	//得到总权重
	for _, one := range randomItems {
		if len(one) < 3 || one[2] <= 0 {
			continue
		}
		totalRate += one[2]
	}
	if totalRate == 0 {
		return nil
	}
	random := rand.Intn(totalRate) + 1
	currentRandom := 0
	for _, one := range randomItems {
		if len(one) < 3 || one[0] <= 0 || one[1] <= 0 {
			continue
		}
		currentRandom += one[2]
		if currentRandom >= random {
			return &gamedb.ItemInfo{
				ItemId: one[0],
				Count:  one[1],
			}
		}
	}
	return nil
}
