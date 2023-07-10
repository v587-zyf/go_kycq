package fightModule

import (
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/gamelibs/gamedb"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
)

/*爬塔*/
type TowerFight struct {
	*DefaultFight
	fightResult int
	towerConf   *gamedb.TowerTowerCfg
}

func NewTowerFight(stageId int) (*TowerFight, error) {
	var err error
	towerFight := &TowerFight{
		towerConf: gamedb.GetTowerByStageId(stageId),
	}
	towerFight.DefaultFight, err = NewDefaultFight(stageId, towerFight)
	if err != nil {
		return nil, err
	}
	towerFight.InitMonster()
	towerFight.Start()
	return towerFight, nil
}

func (this *TowerFight) OnDie(actor base.Actor, killer base.Actor) {
	if actor.GetType() == pb.SCENEOBJTYPE_MONSTER {

		allDie := true
		for _, v := range this.monsterActors {
			if v.GetProp().HpNow() > 0 {
				allDie = false
				break
			}
		}
		if allDie {
			this.fightResult = pb.RESULTFLAG_SUCCESS
			this.SetLifeTime(-1)
		}

	} else {
		allDie := this.CheckUserAllDie(actor)
		if allDie {
			this.fightResult = pb.RESULTFLAG_FAIL
			this.OnEnd()
		}
	}
}

func (this *TowerFight) OnLeaveUser(userId int) {
	this.Stop()
}

func (this *TowerFight) OnPickAll(lastPickObjId int) {

	this.OnEnd()
}

func (this *TowerFight) OnEnd() {
	//战斗中，肯定是时间到了
	this.fightOver()

	//设置延时销毁
	this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_CLOSING, 15)

}

func (this *TowerFight) fightOver() {
	var fightOwner base.Actor
	actors := this.GetUserActors()
	if len(actors) > 0 {
		//爬塔boss里面肯定只有一个人
		for _, v := range actors {
			if v.HostId() > 0 {
				fightOwner = v
			}
		}
		msg := &pbserver.FSFightEndNtf{
			FightType: int32(this.StageConf.Type),
			StageId:   int32(this.StageConf.Id),
		}
		resultMsg := &pbserver.TowerFightResult{
			UserId: int32(fightOwner.GetUserId()),
			Result: int32(this.fightResult),
			Items:  make(map[int32]int32),
		}
		for k, v := range this.playerPickUp[fightOwner.GetUserId()] {
			resultMsg.Items[int32(k)] = int32(v)
		}
		rb, _ := resultMsg.Marshal()
		msg.CpData = rb

		net.GetGsConn().SendMessageToGs(uint32(fightOwner.HostId()), msg)
		this.playerPickUp = make(map[int]map[int]int)
	}
}

func (this *TowerFight) ContinueFight(stageId int) {

	//this.fightResult = pb.RESULTFLAG_FAIL
	//stageConf := gamedb.GetStageStageCfg(stageId)
	//if stageConf == nil {
	//	logger.Error("获取关卡配置错误", stageId)
	//	return
	//}
	//for _, v := range this.monsterActors {
	//	this.Leave(v)
	//}
	//this.StageConf = stageConf
	//this.InitMonster()
	//this.createTime = time.Now().Unix()
	//this.SetLifeTime(int64(this.GetStageConf().LifeTime))
	////更新玩家坐标
	//users := this.GetUserActors()
	//for _, v := range users {
	//	birthPoint, _ := this.GetUserBirthPoint(-1)
	//	v.MoveTo(birthPoint, pb.MOVETYPE_WALK, true, true)
	//}
	////取消延时销毁
	//this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_RUNNING, 0)
}

func (this *TowerFight) GetPowerRoll() string {
	return this.towerConf.Recommend
}
