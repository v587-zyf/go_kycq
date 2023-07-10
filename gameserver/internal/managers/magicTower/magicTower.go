package magicTower

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gamelibs/publicCon/constMail"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"strconv"
)

func NewMagicTowerManager(module managersI.IModule) *MagicTowerManager {
	return &MagicTowerManager{IModule: module}
}

type MagicTowerManager struct {
	util.DefaultModule
	managersI.IModule
}

/**
 *  @Description: 领取九层魔塔玩家信息（是否领奖 当前积分）
 *  @param user
 *  @param op
 **/
func (this *MagicTowerManager) MagicTowerGetUserInfo(user *objs.User) (int, int, error) {

	stageConf := gamedb.GetStageStageCfg(user.FightStageId)
	if stageConf == nil || stageConf.Type != constFight.FIGHT_TYPE_MAGIC_TOWER {
		return 0, 0, gamedb.ERRFIGHTID
	}
	return this.GetFight().MagicTowerGetUserInfo(user)
}

/**
 *  @Description: 领取九层魔塔层奖励
 *  @param user
 *  @param op
 **/
func (this *MagicTowerManager) MagicTowerlayerAward(user *objs.User, op *ophelper.OpBagHelperDefault) error {

	stageConf := gamedb.GetStageStageCfg(user.FightStageId)
	if stageConf == nil || stageConf.Type != constFight.FIGHT_TYPE_MAGIC_TOWER {
		return gamedb.ERRFIGHTID
	}
	return this.GetFight().MagicTowerlayerAward(user, op)
}

func (this *MagicTowerManager) EnterMagicTower(user *objs.User) (int, error) {

	var score int
	var err error
	stageConf := gamedb.GetStageStageCfg(user.FightStageId)
	if stageConf == nil || stageConf.Type != constFight.FIGHT_TYPE_MAGIC_TOWER {
		score, err = this.GetFight().EnterMagicTowerByLayer(user, constFight.MAGIC_TOWER_ENTER_TYPE_FIRST)
	} else {
		score, err = this.GetFight().EnterMagicTowerByLayer(user, constFight.MAGIC_TOWER_ENTER_TYPE_NEXT)
	}
	return score, err
}

func (this *MagicTowerManager) EndMagicTowerNtf(userRank []*pbserver.ShabakeRankScore) {
	logger.Info("九重魔塔结算  userRank:%v", userRank)

	if userRank == nil {
		return
	}
	//个人奖励处理
	for index, userInfo := range userRank {
		myRank := index + 1
		rewards := gamedb.GetMagicTowerRewardByRank(myRank)
		if rewards == nil {
			logger.Error("GetShabakePerRewardByRank nil ranks:%v", rewards)
			continue
		}

		// 发送邮件
		if userInfo.Score > 0 {
			args := []string{strconv.Itoa(myRank)}
			err := this.GetMail().SendSystemMailWithItemInfos(int(userInfo.Id), constMail.MAILTYPE_MAGICTOWER_RANK, args, rewards)
			if err != nil {
				logger.Error("ShabakeFightEndAck send rank reward error err is %v", err)
				continue
			}
			// 给在线玩家推送消息
			if user := this.GetUserManager().GetUser(int(userInfo.Id)); user != nil {
				this.GetUserManager().SendMessage(user, &pb.MagicTowerEndNtf{
					Rank: int32(myRank),
				}, true)
			}
		} else {
			this.DispatchEvent(int(userInfo.Id), nil, func(userId int, user *objs.User, data interface{}) {
				if user != nil {
					this.GetFight().ClientEnterPublicCopy(user, constFight.FIGHT_TYPE_MAIN_CITY_STAGE, 1)
				}
			})
		}
	}
	return
}
