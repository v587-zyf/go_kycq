package fight

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
)

func (this *Fight) DaBaoKillMonster(msg *pbserver.DaBaoKillMonsterNtf) {
	userId := int(msg.UserId)
	this.DispatchEvent(userId, nil, func(userId int, user *objs.User, data interface{}) {
		if user == nil {
			return
		}
		this.GetDaBao().SyncEnergy(user, -int(msg.Energy))
		this.GetCondition().RecordCondition(user, pb.CONDITION_DAO_BAO_MI_JIN_BOSS_KILL_NUMS, []int{1})
		this.GetUserManager().UpdateCombat(user, -1)
	})
}

func (this *Fight) daBaoFightResult(endMsg *pbserver.FSFightEndNtf) {
	//fightResult := &pbserver.DaBaoResult{}
	//err := fightResult.Unmarshal(endMsg.CpData)
	//if err != nil {
	//	logger.Error("解析战斗服发送来玩家打宝秘境结果异常", err)
	//}
	//
	//items := make([]*model.Item, 0)
	//if len(fightResult.GetItems()) > 0 {
	//	for itemId, itemCount := range fightResult.GetItems() {
	//		items = append(items, &model.Item{ItemId: int(itemId), Count: int(itemCount)})
	//	}
	//
	//	this.GetMail().SendSystemMail(int(fightResult.GetUserId()), constMail.MAILTYPE_GUARDPILLAR_ROUNDS, []string{}, items, 0)
	//	logger.Debug("打宝秘境结束 玩家未拾取物品:%v 已由邮件发送", fightResult.GetItems())
	//}
}
