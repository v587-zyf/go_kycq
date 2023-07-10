package fight

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"time"
)

/**
 *  @Description:	战斗玩家申请复活
 *  @param user
 *  @return *pb.FightUserReliveAck
 *  @return *ophelper.OpBagHelperDefault
 */
func (this *Fight) GsToFsRelive(user *objs.User, safeRelive bool) (*pb.FightUserReliveAck, *ophelper.OpBagHelperDefault, error) {

	if user.FightId <= 0 {
		return nil, nil, gamedb.ERRFIGHTID
	}

	request := &pbserver.GsToFSCheckUserReliveReq{
		UserId: int32(user.Id),
	}
	reply := &pbserver.FsToGSCheckUserReliveAck{}

	err := this.FSRpcCall(user.FightId, user.FightStageId, request, reply)
	if err != nil {
		return nil, nil, err
	}

	if !reply.IsDie {
		return nil, nil, gamedb.ERRUSERLIVE
	}

	stageConf := gamedb.GetStageStageCfg(user.FightStageId)
	if stageConf == nil {
		return nil, nil, gamedb.ERRFIGHTID
	}

	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeFightRelive)
	reliveReq := &pbserver.GsToFSUserReliveReq{
		UserId: int32(user.Id),
	}
	replyReliveReq := &pbserver.FSToGsUserReliveAck{}
	if safeRelive {

		if stageConf.Type == constFight.FIGHT_TYPE_MAGIC_TOWER {
			_, err := this.EnterMagicTowerByLayer(user, constFight.MAGIC_TOWER_ENTER_TYPE_RELIVE)
			if err != nil {
				return nil, nil, err
			}
		} else {
			dieTime := this.GetFieldBossUserDieInfos(user.Id)
			if dieTime+10 > int(time.Now().Unix()) {
				return nil, nil, gamedb.ERRREFCD
			}
			reliveReq.ReliveType = constFight.RELIVE_ADDR_TYPE_BIRTH
			err = this.FSRpcCall(user.FightId, user.FightStageId, reliveReq, replyReliveReq)
			if err != nil {
				return nil, nil, err
			}
		}

	} else {
		reliveReq.ReliveType = constFight.RELIVE_ADDR_TYPE_SITU
		cost := this.getReliveCost(stageConf, int(reply.ReliveByIngotTimes))

		if cost > 0 {
			err = this.GetBag().Remove(user, op, pb.ITEMID_INGOT, cost)
			if err != nil {
				return nil, nil, err
			}
		}
		err1 := this.FSRpcCall(user.FightId, user.FightStageId, reliveReq, replyReliveReq)
		if err1 != nil {
			return nil, nil, err1
		}
	}

	return &pb.FightUserReliveAck{
		ReliveTimes:        replyReliveReq.ReliveTimes,
		ReliveByIngotTimes: replyReliveReq.ReliveByIngotTimes,
	}, op, nil
}

/**
 *  @Description:
 *  @param stageConf
 *  @param reliveTime 已经复活次数
 *  @return int
 */
func (this *Fight) getReliveCost(stageConf *gamedb.StageStageCfg, reliveTime int) int {

	reliveTime += 1
	if stageConf.Consume[reliveTime] > 0 {
		return stageConf.Consume[reliveTime]
	}
	max := 0
	for _, v := range stageConf.Consume {
		if v > max {
			max = v
		}
	}
	return max
}
