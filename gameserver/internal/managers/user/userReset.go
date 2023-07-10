package user

import (
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/common"
	"cqserver/protobuf/pb"
	"time"
)

func (this *UserManager) TimingUpdate(t int) {
	this.usersMu.Lock()
	uIds := make([]int, len(this.users))
	i := 0
	for v := range this.users {
		uIds[i] = v
		i++
	}
	this.usersMu.Unlock()
	for _, v := range uIds {
		this.DispatchEvent(v, nil, func(userId int, user *objs.User, data interface{}) {
			switch t {
			case 0:
				this.userZeroTimeUpdate(user)
			case 5:
				this.moduleFun()
				this.userHour5Reset(user)
			case constUser.RESET_DAILYPACK:
				this.GetDailyPack().ResetDailyPack(user, true)
			case constUser.RESET_SPENDREBATE:
				this.GetSpendRebates().ResetSpendRebate(user, true)
			case constUser.RESET_FIRSTRECHARGE:
				this.GetUserManager().SendMessage(user, &pb.ResetNtf{Type: map[int32]int32{pb.RESETTYPE_FIRST_RECHARGE: pb.RESETTYPE_FIRST_RECHARGE}, NewDayTime: int32(time.Now().Unix())}, true)
			}
			user.Dirty = true
		})
	}
}

func (this *UserManager) userHour5Reset(user *objs.User) {
}

//每日5点定时模块任务
func (this *UserManager) moduleFun() {

}

func (this *UserManager) userZeroTimeUpdate(user *objs.User) {
	resetTime := common.GetResetTime(time.Now())
	ntfMap := make(map[int32]int32)

	//签到
	this.GetSign().ResetSign(user, true)
	ntfMap[pb.RESETTYPE_SIGN] = pb.RESETTYPE_SIGN
	//挖矿
	this.GetMining().MiningReset(user, resetTime)
	ntfMap[pb.RESETTYPE_MINING] = pb.RESETTYPE_MINING
	//经验副本
	this.GetExpStage().ResetExpStage(user, resetTime)
	ntfMap[pb.RESETTYPE_EXPSTAGE] = pb.RESETTYPE_EXPSTAGE
	//暗殿boss
	this.GetDarkPalace().ResetDarkPalace(user, resetTime)
	ntfMap[pb.RESETTYPE_DARKPALACE] = pb.RESETTYPE_DARKPALACE
	//野外boss
	this.GetFieldBoss().ResetFieldBoss(user, resetTime)
	ntfMap[pb.RESETTYPE_FIELDBOSS] = pb.RESETTYPE_FIELDBOSS
	//个人boss
	this.GetPersonBoss().ResetPersonBoss(user, resetTime)
	ntfMap[pb.RESETTYPE_PERSONBOSS] = pb.RESETTYPE_PERSONBOSS
	//vipBoss
	this.GetVipBoss().ResetVipBossFightNum(user, resetTime)
	ntfMap[pb.RESETTYPE_VIPBOSS] = pb.RESETTYPE_VIPBOSS
	//材料副本
	this.GetMaterialStage().ResetMaterialStage(user, resetTime)
	ntfMap[pb.RESETTYPE_MATERIALSTAGE] = pb.RESETTYPE_MATERIALSTAGE
	//商城
	this.GetShop().ResetShop(user, resetTime, true)
	ntfMap[pb.RESETTYPE_SHOP] = pb.RESETTYPE_SHOP
	//每日任务
	this.GetDailyTask().Reset(user, resetTime, true)
	ntfMap[pb.RESETTYPE_DAILY_TASK] = pb.RESETTYPE_DAILY_TASK
	//月卡
	this.GetMonthCard().ResetMonthCard(user, resetTime)
	ntfMap[pb.RESETTYPE_MONTH_CARD] = pb.RESETTYPE_MONTH_CARD
	//在线奖励
	this.GetOnline().ResetOnline(user)
	ntfMap[pb.RESETTYPE_ONLINE] = pb.RESETTYPE_ONLINE
	//战令
	if this.GetWarOrder().ResetWarOrder(user, true) {
		ntfMap[pb.RESETTYPE_WAR_ORDER] = pb.RESETTYPE_WAR_ORDER
	}
	//竞技场
	this.GetCompetitve().DayReset(user, true)
	ntfMap[pb.RESETTYPE_COMPETITVE] = pb.RESETTYPE_COMPETITVE
	//野战
	this.GetFieldFight().DayReset(user, true)
	ntfMap[pb.RESETTYPE_FIELDFIGHT] = pb.RESETTYPE_FIELDFIGHT
	//抽卡获得重置
	this.GetCardActivity().Rest(user, false)
	ntfMap[pb.RESETTYPE_CARD] = pb.RESETTYPE_CARD
	//寻龙探宝
	this.GetTreasure().Reset(user, false)
	ntfMap[pb.RESETTYPE_TREASURE] = pb.RESETTYPE_TREASURE
	//膜拜
	this.updateDailyState(user)
	ntfMap[pb.RESETTYPE_DAILY_RANK] = pb.RESETTYPE_DAILY_RANK
	//开服日期
	ntfMap[pb.RESETTYPE_OPENDAY] = int32(time.Now().Unix())
	//世界首领
	ntfMap[pb.RESETTYPE_WORLD_LEADER] = pb.RESETTYPE_WORLD_LEADER
	//远古首领
	this.GetAncientBoss().ResetAncientBoss(user, resetTime)
	ntfMap[pb.RESETTYPE_ANCIENTBOSS] = pb.RESETTYPE_ANCIENTBOSS
	//充值
	this.GetRecharge().RechargeReset(user)
	//连续充值
	this.GetRecharge().ContRechargeReset(user, true)
	ntfMap[pb.RESETTYPE_CONTRECHARGE] = pb.RESETTYPE_CONTRECHARGE
	//掉落红包
	if user.RedPacketItem.Day != resetTime {
		user.RedPacketItem = &model.RedPacketItem{
			Day:      resetTime,
			PickNum:  0,
			PickInfo: make(model.IntKv),
		}
		this.GetUserManager().SendMessage(user, &pb.UserRedPacketGetNumNtf{0}, true)
	}
	//多宝阁
	ntfMap[pb.RESETTYPE_TREASURE_SHOP] = pb.RESETTYPE_TREASURE_SHOP
	//摇彩
	this.GetLottery().UserReset(user)
	ntfMap[pb.RESETTYPE_LOTTERY] = pb.RESETTYPE_LOTTERY
	//试炼塔
	ntfMap[pb.RESETTYPE_TOWER] = pb.RESETTYPE_TOWER
	//炼狱首领
	this.GetHellBoss().ResetHellBoss(user, resetTime)
	ntfMap[pb.RESETTYPE_HELLBOSS] = pb.RESETTYPE_HELLBOSS
	//头衔
	this.GetLabel().ResetLabel(user, resetTime)
	ntfMap[pb.RESETTYPE_LABEL] = pb.RESETTYPE_LABEL
	//红包
	this.GetFirstDrop().Reset(user, resetTime)
	ntfMap[pb.RESETTYPE_RED_PACKET_USE_NUMS] = pb.RESETTYPE_RED_PACKET_USE_NUMS

	this.GetUserManager().SendMessage(user, &pb.ResetNtf{Type: ntfMap, NewDayTime: int32(time.Now().Unix())}, true)
}
