package online

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constMail"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"time"
)

type Online struct {
	util.DefaultModule
	managersI.IModule
}

func NewOnline(module managersI.IModule) *Online {
	p := &Online{IModule: module}
	return p
}

//玩家上线 记录玩家登录时间
func (this *Online) Online(user *objs.User) {
	user.LoginTime = time.Now()
	this.ResetOnline(user)
}

func (this *Online) ResetOnline(user *objs.User) {
	user.OnlineTime = time.Now()
	today := this.GetSystem().GetServerOpenDaysByServerId(user.ServerId)
	if user.OnlineAward == nil || today != user.OnlineAward.Day {
		if user.OnlineAward != nil {
			cfgs := gamedb.GetRewardOnlineCfgs()
			onlineTime := user.OnlineAward.OnlineTime + int(time.Now().Sub(user.OnlineTime).Seconds())
			addMap := make(map[int]int)
			for id, cfg := range cfgs {
				if cfg.Time > onlineTime {
					continue
				}
				isReward := false
				for _, v := range user.OnlineAward.GetAwardIds {
					if v == id {
						isReward = true
						break
					}
				}
				if isReward {
					continue
				}
				for _, reward := range cfg.Rewards {
					addMap[reward.ItemId] += reward.Count
				}
			}
			if len(addMap) > 0 {
				bags := make([]*model.Item, 0)
				for itemId, count := range addMap {
					bags = append(bags, &model.Item{
						ItemId: itemId,
						Count:  count,
					})
				}
				err := this.GetMail().SendSystemMail(user.Id, constMail.ONLINE_ID, []string{}, bags, 0)
				if err != nil {
					logger.Error("online sendMail err:%v", err)
				}
			}
		}
		user.OnlineAward = &model.OnlineAward{
			Day:         today,
			GetAwardIds: make([]int, 0),
		}
	}
}

//玩家下线，记录玩家今天在线时间
func (this *Online) OffLine(user *objs.User, isOffLine bool) {
	timeNow := time.Now()
	onlineTime := int(timeNow.Sub(user.LastUpdateTime).Seconds())
	if isOffLine {
		user.OfflineTime = timeNow
	}
	if onlineTime > 30 {
		user.OnlineAward.OnlineTime += onlineTime
	}
	this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_ONLINE, []int{onlineTime})
	user.Dirty = true
}

//发放奖励
func (this *Online) GetOnlineAward(user *objs.User, awardId int, op *ophelper.OpBagHelperDefault) error {
	for _, v := range user.OnlineAward.GetAwardIds {
		if v == awardId {
			return gamedb.ERRAWARDGET
		}
	}
	conf := gamedb.GetRewardsOnlineAwardCfg(awardId)
	if conf == nil {
		return gamedb.ERRSETTINGNOTFOUND
	}
	onlineTime := user.OnlineAward.OnlineTime + int(time.Now().Sub(user.OnlineTime).Seconds())
	if conf.Time > onlineTime {
		return gamedb.ERRCONDITION
	}

	user.OnlineAward.GetAwardIds = append(user.OnlineAward.GetAwardIds, awardId)
	user.Dirty = true
	this.GetBag().AddItems(user, conf.Rewards, op)
	return nil
}
