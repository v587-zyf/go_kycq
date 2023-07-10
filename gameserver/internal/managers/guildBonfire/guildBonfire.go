package guildBonfire

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constConstant"
	"cqserver/gamelibs/publicCon/constGuild"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"math"
)

const (
	WOOD = 1
	GOLD = 2
)

func NewGuildBonfireManager(m managersI.IModule) *GuildBonfireManager {
	guild := &GuildBonfireManager{}
	guild.IModule = m
	return guild
}

type GuildBonfireManager struct {
	util.DefaultModule
	managersI.IModule
	opChan chan opMsg
}

type opMsg struct {
	stopUserId int
	allStop    bool
}

func (this *GuildBonfireManager) Init() error {
	this.opChan = make(chan opMsg, 50)
	return nil
}

func (this *GuildBonfireManager) LoadInfo(user *objs.User, ack *pb.GuildBonfireLoadAck) error {

	if user.GuildData.NowGuildId <= 0 {
		return gamedb.ERRHAVENOGUILD
	}

	state := this.JudgeGuildBonfireIsOpen(user)
	if !state {
		return gamedb.ERRGUILDBONFIREISNOTOPEN
	}

	guildInfo := this.GetGuild().GetGuildInfo(user.GuildData.NowGuildId)
	if guildInfo == nil {
		logger.Error("GetGuildInfo nil userId:%v  guildId:%v", user.Id, user.GuildData.NowGuildId)
		return gamedb.ERRHAVENOGUILD
	}
	ack.ExpAddPercent = this.GetAddPercent(user.GuildData.NowGuildId)
	ack.PeopleList = this.GetWoodPeople(user) //投放木材的玩家
	return nil
}

//
//  GuildAddExpPercent
//  @Description:投放木材或者花费元宝  增加经验百分比加成
//
func (this *GuildBonfireManager) GuildAddExpPercent(user *objs.User, op *ophelper.OpBagHelperDefault, consumptionType int, ack *pb.GuildBonfireAddExpAck) error {

	state := this.JudgeGuildBonfireIsOpen(user)
	if !state {
		return gamedb.ERRGUILDBONFIREISNOTOPEN
	}

	cfg := gamedb.GetGuildBonfireGuildBonfireCfg(consumptionType)
	if cfg == nil {
		return gamedb.ERRGUILDBONFIRETYPEERR
	}

	if user.GuildData.NowGuildId <= 0 {
		return gamedb.ERRHAVENOGUILD
	}

	guildInfo := this.GetGuild().GetGuildInfo(user.GuildData.NowGuildId)
	if guildInfo == nil {
		return gamedb.ERRNOGUILD
	}
	if guildInfo.DonateTimes[consumptionType] >= cfg.Times {
		return gamedb.ERRGUILDBONFIREOVERTIMES
	}

	enough, _ := this.GetBag().HasEnough(user, cfg.Item.ItemId, cfg.Item.Count)
	if !enough {
		return gamedb.ERRNOTENOUGHGOODS
	}

	err := this.GetBag().Remove(user, op, cfg.Item.ItemId, cfg.Item.Count)
	if err != nil {
		logger.Error("GuildAddExp remove err:%v", err)
		return err
	}

	guildInfo.DonateUsers = append(guildInfo.DonateUsers, user.Id, consumptionType)
	guildInfo.DonateTimes[consumptionType] += 1
	this.GetGuild().SetGuildInfo(guildInfo)
	ack.ExpAddPercent = this.GetAddPercent(user.GuildData.NowGuildId)
	ack.PeopleList = this.GetWoodPeople(user) //投放木材的玩家
	return nil
}
func (this *GuildBonfireManager) AddUserExp(userIds []int) {

	for _, userId := range userIds {
		userInfo := this.GetUserManager().GetUser(userId)
		if userInfo != nil {
			this.AddItemInfosByPercent(userInfo)

		}
	}

}

func (this *GuildBonfireManager) StopAddUserExp(userId int) {

	this.opChan <- opMsg{stopUserId: userId}

}

func (this *GuildBonfireManager) StopAllUserAdd() {
	this.opChan <- opMsg{allStop: true}
}

//
//  JudgeGuildBonfireIsOpen
//  @Description: 判断活动是否开启
//
func (this *GuildBonfireManager) JudgeGuildBonfireIsOpen(user *objs.User) bool {
	flag := true
	if err, _ := this.GetGuild().CheckActiveOpen(user, constGuild.GUILD_ACTIVITY_BONFIRE); err != nil {
		flag = false
	}
	return flag
}

func (this *GuildBonfireManager) GetWoodPeople(user *objs.User) []*pb.WoodPeople {
	info := make([]*pb.WoodPeople, 0)

	data := this.GetGuild().GetGuildInfo(user.GuildData.NowGuildId).DonateUsers
	if data != nil {
		for i, j := 0, len(data); i < j; i += 2 {
			userId := data[i]
			types := data[i+1]
			userInfo := this.GetUserManager().BuilderBrieUserInfo(int(userId))
			if userInfo == nil {
				continue
			}
			info = append(info, &pb.WoodPeople{
				NickName: userInfo.Name,
				Avatar:   userInfo.Avatar,
				Times:    int32(1),
				Types:    int32(types),
			})
		}
	}
	return info
}

//
//  GetAddPercent
//  @Description: 获取道具增加百分比
//
func (this *GuildBonfireManager) GetAddPercent(guildId int) float32 {
	addPercent := 0
	guildInfo := this.GetGuild().GetGuildInfo(guildId)
	if guildInfo == nil {
		return float32(addPercent)
	}
	for types, v := range gamedb.GetGuildBonfireGuildBonfire() {
		times := guildInfo.DonateTimes[types]
		addPercent += v.Promote * times
	}
	return float32(addPercent) / float32(constConstant.WAN_FEN_BI)
}

//
//  AddItemInfosByPercent
//  @Description: //玩家再篝火范围内增加Exp
//
func (this *GuildBonfireManager) AddItemInfosByPercent(user *objs.User) {
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeGuildBonfireAddExp)
	addItems := gamedb.GetConf().BonefireRewards
	addPer := this.GetAddPercent(user.GuildData.NowGuildId) + 1
	if addItems != nil {
		for _, item := range addItems {
			count := int(math.Ceil(float64(item.Count) * float64(addPer)))
			this.GetBag().Add(user, op, item.ItemId, count)
		}
	}
	this.GetUserManager().SendItemChangeNtf(user, op)
}

func (this *GuildBonfireManager) GuildBonfireIsOpenNtf(isOpen bool) {
	this.BroadcastAll(&pb.GuildBonfireOpenStateNtf{IsOpen: isOpen})
	return
}

//func (this *GuildBonfireManager) AddUserExp(userId int) {
//	user := this.GetUserManager().GetAllUserInfoIncludeOfflineUser(userId)
//	if user == nil {
//		logger.Error("AddUserExp user == nil  userId:%v", userId)
//		return
//	}
//	times := gamedb.GetConf().BonefireTime
//	ticker := time.NewTicker(time.Second * time.Duration(times))
//	for {
//		select {
//		case <-ticker.C:
//			logger.Debug("篝火 玩家:%v  每 %v秒 add exp", user.Id, times)
//			this.AddItemInfosByPercent(user)
//
//		case msg := <-this.opChan:
//			if user.Id == msg.stopUserId {
//				logger.Debug("篝火 玩家:%v  离开增加经验区", user.Id)
//				return
//			}
//			if msg.allStop {
//				logger.Debug("活动结束")
//				return
//			}
//
//		}
//	}
//
//}
