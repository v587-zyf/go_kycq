package fight

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
)

func (this *Fight) EnterPaodian(user *objs.User, stageId int) error {

	fightId, err := this.getFightIdByStageId(stageId, 0)
	if err != nil {
		return nil
	}

	if fightId <= 0 {
		//创建战斗
		fightId, err = this.GetFight().CreateFight(stageId, nil)
		if err != nil {
			return err
		}
	}

	//玩家进入战斗
	err = this.GetFight().EnterFightByFightId(user, stageId, fightId)
	if err != nil {
		return err
	}
	return nil
}

/**
 *  @Description: 泡点Pk 玩家增加奖励
 *  @param endMsg
 */
func (this *Fight) PaodianGoodsAdd(endMsg *pbserver.PaodianGoodsAddNtf) {

	paodianRewardConf := gamedb.GetPaodianConfByStageId(int(endMsg.StageId))
	for userId, times := range endMsg.UserIds {
		this.DispatchEvent(int(userId), int(times), func(userId int, user *objs.User, data interface{}) {
			if user == nil {
				logger.Error("泡点获取玩家数据异常：%v，配置：%v，奖励倍数：%v", userId, paodianRewardConf.Id, times)
				return
			}
			t := data.(int)
			this.GetPaoDian().PaoDianRewardNtf(user, paodianRewardConf.Id, t)
		})
	}
}

/**
 *  @Description: 泡点Pk结束
 *  @param endMsg
 */
func (this *Fight) paodianFightEnd(endMsg *pbserver.FSFightEndNtf) {
	ntf := &pb.PaodianFightEnd{
		StageId: endMsg.StageId,
	}
	this.BroadcastAll(ntf)
}
