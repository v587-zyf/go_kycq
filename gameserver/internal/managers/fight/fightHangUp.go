package fight

import (
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pbserver"
)

func (this *Fight) HangUpUserKillWave(msg *pbserver.HangUpKillWaveNtf) {

	userId := int(msg.UserId)
	//stageId := int(msg.StageId)
	this.DispatchEvent(userId, nil, func(userId int, user *objs.User, data interface{}) {

		if user == nil {
			return
		}
		op := ophelper.NewOpBagHelperDefault(constBag.OpTypeNormalStage)
		this.GetStageManager().StageFightEndReq(user, op)
	})
}

func (this *Fight) huangUpBossFightEnd(data []byte) {

	fightResult := &pbserver.HangUpBossFightEndNtf{}
	err := fightResult.Unmarshal(data)
	if err != nil {
		logger.Error("挂机boss战斗结果解析异常：%v", err)
		return
	}

	userId := int(fightResult.UserId)
	this.DispatchEvent(userId, fightResult, func(userId int, user *objs.User, data interface{}) {

		if user != nil {
			items := common.ConvertMapInt32ToInt(fightResult.Items)
			this.GetStageManager().StageBossKillResult(user, int(fightResult.StageId), int(fightResult.Result), items)
			this.GetFirstDrop().CheckIsFirstDrop(user, items)
			kyEvent.StageEnd(user, int(fightResult.StageId), int(fightResult.Result), user.FightStartTime, items)
		} else {
			logger.Error("战斗服发送来挂机boss战斗结果，玩家未找到：%v,结果：%v", userId, fightResult)
		}

	})
}
