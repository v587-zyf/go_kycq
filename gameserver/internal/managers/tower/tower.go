package tower

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constMail"
	"cqserver/gamelibs/publicCon/constRank"
	"cqserver/gamelibs/rmodel"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"strconv"
	"time"
)

type Tower struct {
	util.DefaultModule
	managersI.IModule
}

func NewTower(module managersI.IModule) *Tower {
	p := &Tower{IModule: module}
	return p
}

/**
 *  @Description: 试炼塔领取每日奖励
 *  @param user
 *  @param op
 *  @return error
 */
func (this *Tower) DayAward(user *objs.User, op *ophelper.OpBagHelperDefault) error {
	today := common.GetResetTime(time.Now())
	if today == user.Tower.DayAwardState {
		return gamedb.ERRAWARDGET
	}

	userTower := user.Tower
	layer := userTower.TowerLv
	if layer != 1 {
		layer -= 1
	}
	towerConf := gamedb.GetTowerReward(layer)
	if towerConf == nil {
		return gamedb.ERRUNKNOW
	}
	userTower.DayAwardState = today

	user.Dirty = true
	this.GetBag().AddItems(user, towerConf, op)
	return nil
}

/**
 *  @Description: 试炼塔抽奖
 *  @param user
 *  @param op
 *  @return int
 *  @return error
 */
func (this *Tower) Lottery(user *objs.User, op *ophelper.OpBagHelperDefault) (int, error) {
	userTower := user.Tower

	lotteryNumCfg := gamedb.GetConf().LotteryChance
	lotteryNum := (userTower.TowerLv - 1) / lotteryNumCfg
	if userTower.LotteryNum >= lotteryNum {
		return 0, gamedb.ERRNOTENOUGHTIMES
	}

	towerLotteryConf := gamedb.GetTowerLotteryCircleTowerLotteryCircleCfg(userTower.LotteryId)
	if towerLotteryConf == nil {
		return 0, gamedb.ERRPARAM
	}
	this.GetBag().AddItems(user, towerLotteryConf.ItemId, op)

	userTower.LotteryNum++
	userTower.LotteryId++
	user.Dirty = true
	return 0, nil
}

/**
 *  SendRankReward
 *  @Description: 试炼塔排行奖励
 *  @receiver this
**/
func (this *Tower) SendRankReward() {
	week := common.GetYearWeek(time.Now())
	if rmodel.Rank.GetRankReward(pb.RANKTYPE_TOWER) == week {
		return
	}
	rankSlice := this.GetRank().LoadRank(pb.RANKTYPE_TOWER, constRank.MAX)
	rank := 0
	for i := 0; i < len(rankSlice); i += 2 {
		userId := rankSlice[i]
		rank++
		this.DispatchEvent(userId, rank, func(userId int, user *objs.User, data interface{}) {
			rankIndex := data.(int)
			reward := gamedb.GetTowerRankReward(rankIndex)
			if reward != nil && len(reward) > 0 {
				items := make([]*model.Item, 0)
				for _, info := range reward {
					items = append(items, &model.Item{ItemId: info.ItemId, Count: info.Count})
				}
				this.GetMail().SendSystemMail(userId, constMail.MAILTYPE_TOWER_RANK, []string{strconv.Itoa(rankIndex)}, items, 0)
			}
		})
	}
	rmodel.Rank.SetRankReward(pb.RANKTYPE_TOWER, week)
}
