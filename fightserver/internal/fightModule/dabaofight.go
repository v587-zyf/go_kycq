package fightModule

import (
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/gamelibs/gamedb"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pbserver"
)

type DaBaoFight struct {
	*DefaultFight
}

func NewDaBaoFight(stageId int) (*DaBaoFight, error) {
	var err error
	fight := &DaBaoFight{}
	fight.DefaultFight, err = NewDefaultFight(stageId, fight)
	if err != nil {
		return nil, err
	}
	fight.SetLifeTime(-1)
	fight.InitMonster()
	fight.Start()
	return fight, nil
}

//func (this *DaBaoFight) OnDie(actor base.Actor, killer base.Actor) {
//	if actor.GetType() == pb.SCENEOBJTYPE_MONSTER {
//		monster := actor.(*actorPkg.MonsterActor)
//		this.sendKillMonster(monster.MonsterT)
//	} else {
//		//allDie := this.CheckUserAllDie(actor)
//		//if allDie {
//		//	this.OnEnd()
//		//}
//	}
//}

func (this *DaBaoFight) MonsterDrop(dropMonsterId int, dropX, dropY int, owner base.Actor, dropItems []*pbserver.ItemUnit) {
	if user, ok := owner.(base.ActorPlayer); ok {
		userDabaoEnergy := user.GetPlayer().DaBaoEnergy()
		monsterT := gamedb.GetMonsterMonsterCfg(dropMonsterId)
		if monsterT == nil {
			logger.Error("获取怪物配置异常：%v", dropMonsterId)
			return
		}

		if userDabaoEnergy < monsterT.DaBaoEnergy {
			logger.Warn("玩家体力小于怪物掉落体力，不掉落，%v,体力：%v,怪物id:%v，体力需求：%v", owner.NickName(), userDabaoEnergy, dropMonsterId, monsterT.DaBaoEnergy)
			return
		}
		//物品掉落
		this.DefaultFight.MonsterDrop(dropMonsterId, dropX, dropY, owner, dropItems)
		//扣除玩家体力
		this.sendKillMonster(monsterT)
	} else {
		logger.Error("归属玩家数据异常：%v", owner)
	}
}

func (this *DaBaoFight) OnLeaveUser(userId int) {
	this.Stop()
}

func (this *DaBaoFight) OnEnd() {
	user := this.getUser()

	msg := &pbserver.FSFightEndNtf{
		FightType: int32(this.StageConf.Type),
		StageId:   int32(this.StageConf.Id),
	}

	//resultMsg := &pbserver.DaBaoResult{
	//	UserId: int32(user.GetUserId()),
	//	Items:  make(map[int32]int32),
	//}
	//allItem, err := this.PickUp(user.GetUserId(), []int32{})
	//if err != nil {
	//	logger.Error("打宝秘境结束 拾取物品错误 用户:%v err:%v", user.GetUserId(), err)
	//	return
	//}
	//for _, itemInfo := range allItem {
	//	resultMsg.Items[itemInfo.ItemId] = itemInfo.ItemNum
	//}
	//rb, _ := resultMsg.Marshal()
	//msg.CpData = rb
	logger.Info("发送game打宝秘境结果,hostId:%v 结果：%v", user.HostId(), *msg)
	net.GetGsConn().SendMessageToGs(uint32(user.HostId()), msg)
	this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_CLOSING, 10)
}

func (this *DaBaoFight) sendKillMonster(monster *gamedb.MonsterMonsterCfg) {
	//this.SyncEnergy(this.Energy - monster.DaBaoEnergy)
	user := this.getUser()
	msg := &pbserver.DaBaoKillMonsterNtf{
		UserId:    int32(user.GetUserId()),
		MonsterId: int32(monster.Monsterid),
		Energy:    int32(monster.DaBaoEnergy),
	}
	net.GetGsConn().SendMessageToGs(uint32(user.HostId()), msg)
}

func (this *DaBaoFight) getUser() base.Actor {
	var owner base.Actor
	actors := this.GetUserActors()
	if len(actors) > 0 {
		for _, v := range actors {
			if v.HostId() > 0 {
				owner = v
				break
			}
		}
	}
	logger.Debug("actors:%v owner:%v", actors, owner)
	return owner
}
