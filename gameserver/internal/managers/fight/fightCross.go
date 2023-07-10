package fight

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pbserver"
)

/**
 *  @Description: 世界首领战斗结果
 *  @param endMsg
 */
func (this *Fight) crossWorldLeaderEnt(endMsg *pbserver.FSFightEndNtf) {

	fightResult := &pbserver.WorldLeaderFightEndNtf{}
	err := fightResult.Unmarshal(endMsg.CpData)
	if err != nil {
		logger.Error("接受战斗服世界首领结果解析异常,stageId:%v,err:%v", endMsg.StageId, err)
	}

	this.GetWorldLeader().EndWorldLeaderBossNtf(fightResult)
}

/**
 *  @Description:	世界首领排行 血量信息
 *  @param endMsg
 */
func (this *Fight) CrossWorldLeaderRankNtf(endMsg *pbserver.WorldLeaderFightRankNtf) {

	this.GetWorldLeader().WorldLeaderFightRankNtf(endMsg)
}

/**
 *  @Description: 进入跨服沙巴克战斗
 *  @param user
 *  @return error
 */
func (this *Fight) EnterShabakeCrossFight(user *objs.User) error {

	request := &pbserver.GSTOFSGetFightIdReq{
		StageId: int32(constFight.FIGHT_TYPE_SHABAKE_CROSS_STAGE),
	}
	replay := &pbserver.FSTOGSGetFightIdAck{}
	err := this.FSRpcCall(0, constFight.FIGHT_TYPE_SHABAKE_CROSS_STAGE, request, replay)
	if err != nil {
		logger.Error("沙巴克战斗获取Id异常,stageId：%v,err:%v", constFight.FIGHT_TYPE_SHABAKE_CROSS_STAGE, err)
		return gamedb.ERRFIGHTID
	}
	//玩家进入战斗
	err = this.GetFight().EnterFightByFightId(user, constFight.FIGHT_TYPE_SHABAKE_CROSS_STAGE, int(replay.FightId))
	if err != nil {
		return err
	}
	return nil
}

/**
 *  @Description: 跨服沙巴克战斗结果
 *  @param endMsg
 */
func (this *Fight) guildShabakeCrossEnd(endMsg *pbserver.FSFightEndNtf) {

	this.shabakeFightId = 0
	result := &pbserver.ShabakeCrossFightEndNtf{}
	err := result.Unmarshal(endMsg.CpData)
	if err != nil {
		logger.Error("解析沙巴克结果异常：%v", err)
		return
	}

	this.GetShaBaKeCross().CrossShaBakeFightEndNtf(result)
}