package user

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

/**
 *  @Description: 跨服战斗开启
 */
func (this *UserManager) CrossFightOpen() {
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
			stageConf := gamedb.GetStageStageCfg(user.FightStageId)
			if stageConf.Type == constFight.FIGHT_TYPE_HELL_BOSS ||
				stageConf.Type == constFight.FIGHT_TYPE_CROSS_WORLD_LEADER {
				this.GetFight().ClientEnterPublicCopy(user, constFight.FIGHT_TYPE_MAIN_CITY_STAGE, pb.CONDITION_USER_LEVEL)
				this.SendMessage(user, &pb.CrossFightOpenNtf{StageId: int32(user.FightStageId)}, true)
			}
		})
	}
}
