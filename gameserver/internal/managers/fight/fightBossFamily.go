package fight

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
)

func (this *Fight) GetBossFamilyInfo(t int) (map[int32]int32, error) {

	request := &pbserver.BossFamilyBossInfoReq{
		BossFamilyType: int32(t),
	}

	reply := &pbserver.BossFamilyBossInfoAck{}

	err := this.FSRpcCall(0, 0, request, reply)

	if err != nil {
		return nil, err
	}

	return reply.BossFamilyInfo, nil

}

func (this *Fight) EnterBossFamily(user *objs.User,stageId int) error {

	conf := gamedb.GetBossFamilyBossFamilyCfg(stageId)
	if conf == nil {
		return gamedb.ERRSETTINGNOTFOUND
	}

	if !this.CheckInFightBefore(user,stageId){
		return gamedb.ERRUSERINFIGHT
	}

	if !this.GetCondition().CheckMultiByType(user, -1, conf.Condition,  pb.CONDITIONTYPE_ALL) {
		return gamedb.ERRCONDITION
	}

	fightId := this.getResidentFightId(stageId)
	if fightId <= 0 {
		return gamedb.ERRFIGHTID
	}

	return this.enterFight(user, stageId, int(fightId), false, 0)

}
