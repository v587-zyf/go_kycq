package fight

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pbserver"
)

/**
 *  @Description: 地图npc事件
 *  @param user
 *  @param npcId	地图npc标识Id
 **/
func (this *Fight) FightNpcEventReq(user *objs.User, npcId int) error {

	stageConf := gamedb.GetStageStageCfg(user.FightStageId)
	if stageConf.Type != constFight.FIGHT_TYPE_SHABAKE_NEW{
		return gamedb.ERRUNKNOW
	}

	request := &pbserver.GsToFsFightNpcEventReq{
		UserId: int32(user.Id),
		NpcId:  int32(npcId),
	}
	reply := &pbserver.FsToGsFightNpcEventAck{}

	err := this.FSRpcCall(user.FightId, user.FightStageId, request, reply)
	if err != nil {
		return err
	}


	if stageConf.Type == constFight.FIGHT_TYPE_MAGIC_TOWER {
		this.GetMagicTower().EnterMagicTower(user)
	}
	return nil
}
