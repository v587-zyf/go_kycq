package fight

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pbserver"
)

/**
 *  @Description: 进入九层魔塔下一层
 *  @param endMsg
 */
func (this *Fight) EnterMagicTowerByLayer(user *objs.User, t int) (int, error) {

	magicTowerConf, stageId := this.getMagicTowerConf(user, t)
	if magicTowerConf == nil {
		return 0, gamedb.ERRUNKNOW
	}
	score := 0
	if t == constFight.MAGIC_TOWER_ENTER_TYPE_NEXT {
		request := &pbserver.GsToFsFightScoreLessReq{
			UserId:  int32(user.Id),
			LessNum: int32(magicTowerConf.MarkConsume),
		}
		replay := &pbserver.FsToGsFightScoreLessAck{}
		err := this.FSRpcCall(user.FightId, magicTowerConf.StageId1, request, replay)
		if err != nil {
			logger.Error("九层魔塔积分扣除异常,stageId：%v,err:%v", magicTowerConf.StageId1, err)
			return 0, gamedb.MARKNOTENOUGH
		}
		score = int(replay.Score)
	}

	err := this.EnterMagicTowerFight(user, stageId)
	if err != nil {
		logger.Error("进入九层魔塔战斗异常 ,玩家：%v,stageId :%v,err:%v", user.Id, stageId, err)
	}

	return score, err
}

/**
 *  @Description: 进入九层魔塔下一层
 *  @param user
 *  @return error
 */
func (this *Fight) EnterMagicTowerFight(user *objs.User, stageId int) error {

	request := &pbserver.GSTOFSGetFightIdReq{
		StageId: int32(stageId),
	}
	replay := &pbserver.FSTOGSGetFightIdAck{}
	err := this.FSRpcCall(0, stageId, request, replay)
	if err != nil {
		logger.Error("九层魔塔战斗获取Id异常,stageId：%v,err:%v", stageId, err)
		return gamedb.ERRFIGHTID
	}
	//玩家进入战斗
	err = this.GetFight().EnterFightByFightId(user, stageId, int(replay.FightId))
	if err != nil {
		return err
	}
	return nil
}

/**
*  @Description: 九层魔塔 战斗结算
*  @receiver this
*  @param endMsg
**/
func (this *Fight) magicTowerFightResult(endMsg *pbserver.FSFightEndNtf) {

	fightResult := &pbserver.MagicTowerFightEnd{}
	err := fightResult.Unmarshal(endMsg.CpData)
	if err != nil {
		logger.Error("解析战斗服发送来九层魔塔战斗结果异常", err)
	}
	this.GetMagicTower().EndMagicTowerNtf(fightResult.UserRank)
}

func (this *Fight) getMagicTowerConf(user *objs.User, t int) (*gamedb.MagicTowerMagicTowerCfg, int) {
	var nowMagicTowerConf *gamedb.MagicTowerMagicTowerCfg
	var stageId int
	if t == constFight.MAGIC_TOWER_ENTER_TYPE_NEXT {
		//下一层
		gamedb.RangMagicTowerMagicTowerCfgs(func(conf *gamedb.MagicTowerMagicTowerCfg) bool {
			if conf.StageId1 == user.FightStageId {
				if nextConf := gamedb.GetMagicTowerMagicTowerCfg(conf.Id + 1); nextConf != nil {
					nowMagicTowerConf = conf
					stageId = nextConf.StageId1
					return false
				}
			}
			return true
		})

	} else if t == constFight.MAGIC_TOWER_ENTER_TYPE_RELIVE {
		//复活
		gamedb.RangMagicTowerMagicTowerCfgs(func(conf *gamedb.MagicTowerMagicTowerCfg) bool {
			if conf.StageId1 == user.FightStageId {
				nowMagicTowerConf = conf
				stageId = conf.StageId2
				return false

			}
			return true
		})

	} else {
		//进入第一层
		gamedb.RangMagicTowerMagicTowerCfgs(func(conf *gamedb.MagicTowerMagicTowerCfg) bool {
			if conf.Id == 1 {
				nowMagicTowerConf = conf
				stageId = conf.StageId1
				return false
			}
			return true
		})
	}
	return nowMagicTowerConf, stageId
}

/**
 *  @Description: 领取九层魔塔玩家信息（是否领奖 当前积分）
 *  @param user
 *  @param op
 **/
func (this *Fight) MagicTowerGetUserInfo(user *objs.User) (int, int, error) {

	request := &pbserver.MagicTowerGetUserInfoReq{
		UserId:     int32(user.Id),
		IsGetAward: false,
	}
	replay := &pbserver.MagicTowerGetUserInfoAck{}
	err := this.FSRpcCall(user.FightId, user.FightStageId, request, replay)
	if err != nil {
		logger.Error("九层魔塔战斗获取Id异常,stageId：%v,err:%v", user.FightStageId, err)
		return 0, 0, err
	}
	isGetAward := 0
	if replay.IsGetAward {
		isGetAward = 1
	}
	return int(replay.Score), isGetAward, nil
}

/**
 *  @Description: 领取九层魔塔层奖励
 *  @param user
 *  @param op
 **/
func (this *Fight) MagicTowerlayerAward(user *objs.User, op *ophelper.OpBagHelperDefault) error {

	request := &pbserver.MagicTowerGetUserInfoReq{
		UserId:     int32(user.Id),
		IsGetAward: true,
	}
	replay := &pbserver.MagicTowerGetUserInfoAck{}
	err := this.FSRpcCall(user.FightId, user.FightStageId, request, replay)
	if err != nil {
		logger.Error("九层魔塔战斗获取Id异常,stageId：%v,err:%v", user.FightStageId, err)
		return err
	}
	if !replay.CanGetAward {

		return gamedb.ERRAWARDGET
	}
	gamedb.RangMagicTowerMagicTowerCfgs(func(conf *gamedb.MagicTowerMagicTowerCfg) bool {
		if conf.StageId1 == user.FightStageId {
			this.GetBag().AddItems(user, conf.Rewards, op)
			return false
		}
		return true
	})
	return nil
}
