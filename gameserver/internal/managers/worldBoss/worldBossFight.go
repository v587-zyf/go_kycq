package worldBoss

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constMail"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"strconv"
	"time"
)

func (this *WorldBoss) EnterWorldBossFight(user *objs.User, stageId int) error {

	worldBossConf := gamedb.GetWorldBossByStageId(stageId)

	// 判断是否是开启时间
	nowT := common.GetTimeSeconds(time.Now())
	openT := worldBossConf.OpenTime.GetSecondsFromZero()
	continueT := worldBossConf.OpenTime.GetSecondsFromZero() + worldBossConf.Continue*60
	if openT > nowT || nowT > continueT {
		return gamedb.ERRNOTTOTIME
	}

	//进入战斗
	this.GetFight().EnterResidentFightByStageId(user, stageId, 0)

	return nil
}

func (this *WorldBoss) WorldBossFightEndAck(ranks []int, lucker, stageId int) error {
	op := ophelper.NewOpBagHelperDefault(constBag.OpTypeWorldBossFight)
	// 循环榜单信息
	for k, uId := range ranks {
		rank := k + 1
		rankConf := gamedb.GetWorldRank(rank)
		// 发送邮件
		bags := makeMailBags(rankConf.Reward)
		args := []string{strconv.Itoa(rank)}
		err := this.GetMail().SendSystemMail(uId, constMail.WorldBoss_RankReward_Id, args, bags, 0)
		if err != nil {
			logger.Error("worldBossFightEndAck send rank reward error err is %v", err)
			continue
		}
		// 给在线玩家推送消息
		if user := this.GetUserManager().GetUser(uId); user != nil {
			this.GetUserManager().SendMessage(user, &pb.WorldBossFightResultNtf{
				StageId: int32(stageId),
				Rank:    int32(rank),
				Goods:   op.ToChangeItems(),
			}, false)
		}
	}
	// 幸运玩家
	luckConf := gamedb.GetWorldBossByStageId(stageId)
	// 发送邮件
	bags := makeMailBags(luckConf.Lucky)
	err := this.GetMail().SendSystemMail(lucker, constMail.WorldBoss_LuckReward_Id, nil, bags, 0)
	if err != nil {
		logger.Error("worldBossFightEndAck send lucker reward error err is %v", err)
		return err
	}

	this.BroadcastAll(this.GetWorldBossInfo())
	return nil
}
