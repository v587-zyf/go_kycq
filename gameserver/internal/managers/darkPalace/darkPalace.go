package darkPalace

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"time"
)

func NewAreaManager(m managersI.IModule) *DarkPalaceManager {
	return &DarkPalaceManager{IModule: m}
}

type DarkPalaceManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *DarkPalaceManager) Online(user *objs.User) {
	date := common.GetResetTime(time.Now())
	this.ResetDarkPalace(user, date)
}

func (this *DarkPalaceManager) ResetDarkPalace(user *objs.User, date int) {
	userDarkPalace := user.DarkPalace
	if userDarkPalace.ResetTime != date {
		userDarkPalace.ResetTime = date
		userDarkPalace.BuyNum = 0
		userDarkPalace.DareNum = 0
		userDarkPalace.HelpNum = 0
	}
}

/**
 *  @Description: 暗殿加载信息
 *  @param floor  层数
 *  @param ack
 */
func (this *DarkPalaceManager) Load(floor int, ack *pb.DarkPalaceLoadAck) {
	pbDarkPalace := make([]*pb.DarkPalaceBossNtf, 0)
	darkPalaceBossMap := gamedb.GetDarkPalaceBossByFloor(floor)
	for stageId := range darkPalaceBossMap {
		info := this.GetDarkPalaceBossInfo(stageId)
		if info != nil {
			pbDarkPalace = append(pbDarkPalace, info)
		} else {
			logger.Error("获取暗殿boss战斗信息异常", stageId)
		}
	}
	ack.DarkPalaceBoss = pbDarkPalace
}

/**
 *  @Description: 暗殿购买奖励次数
 *  @param user
 *  @param use	是否购买并使用
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *DarkPalaceManager) BuyNum(user *objs.User, use bool, buyNum int, op *ophelper.OpBagHelperDefault, ack *pb.DarkPalaceBuyNumAck) error {
	if buyNum < 1 {
		return gamedb.ERRPARAM
	}
	userDarkPalace := user.DarkPalace
	if userDarkPalace.BuyNum+buyNum > gamedb.GetConf().NewRewardMaxTimes+this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_DARKPALACE_BUYNUM) {
		return gamedb.ERRENOUGHTIMES
	}

	cost := gamedb.GetConf().DarkPalaceBuy
	consume := cost[0]
	hasNum, _ := this.GetBag().GetItemNum(user, consume.ItemId)
	if hasNum < 1 && use {
		consume = cost[1]
	}
	if err := this.GetBag().Remove(user, op, consume.ItemId, consume.Count*buyNum); err != nil {
		return err
	}
	user.DarkPalace.BuyNum += buyNum
	user.Dirty = true

	ack.BuyNum = int32(user.DarkPalace.BuyNum)

	this.GetFight().UpdateUserfightNum(user)

	return nil

}

/**
 *  @Description: 发送暗殿boss信息(复活或死亡)
 *  @param stageId
 *  @param floor	层数
 *  @param blood	血量
 *  @param reliveTime	复活时间
 */
func (this *DarkPalaceManager) SendDarkPalaceBossNtf(darkPalaceBossNtf *pb.DarkPalaceBossNtf) {
	this.BroadcastAll(darkPalaceBossNtf)
}

/**
 *  @Description: 获取怪物信息
 *  @param stageId	怪物id
 */
func (this *DarkPalaceManager) GetDarkPalaceBossInfo(stageId int) *pb.DarkPalaceBossNtf {
	return this.GetFight().GetDarkPalaceBossInfos(stageId)
}

/**
 *  @Description: 获取剩余次数
 *  @return int
 */
func (this *DarkPalaceManager) GetSurplusNum(user *objs.User) int {
	userDarkPalace := user.DarkPalace
	defFightNum := gamedb.GetConf().DarkPalaceChallengeTimes
	vipPrivilege := this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_DARKPALACE_FREENUM)
	return defFightNum + userDarkPalace.BuyNum + vipPrivilege - userDarkPalace.DareNum
}
