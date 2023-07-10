package fightModule

import (
	"cqserver/fightserver/internal/actorPkg"
	"cqserver/fightserver/internal/base"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/protobuf/pb"
)

/*公共打宝*/
type PublicDabaoFight struct {
	*DefaultFight
	areaRelive map[int][]int
}

func NewPublicDabaoFight(stageId int) (*PublicDabaoFight, error) {
	var err error
	publicDabaoFight := &PublicDabaoFight{}
	publicDabaoFight.DefaultFight, err = NewDefaultFight(stageId, publicDabaoFight)
	if err != nil {
		return nil, err
	}
	publicDabaoFight.SetLifeTime(-1)
	err = publicDabaoFight.InitMonster()
	if err != nil {
		return nil, err
	}
	publicDabaoFight.areaReliveInit()
	publicDabaoFight.Start()
	return publicDabaoFight, nil
}

func (this *PublicDabaoFight) areaReliveInit() {
	if len(this.StageConf.Monster_num) > 0 {

		this.areaRelive = make(map[int][]int)
		for _, v := range this.monsterActors {

			monster := v.(*actorPkg.MonsterActor)
			areaIndex := monster.GetBirthAreaIndex()
			if _, ok := this.StageConf.Monster_num[areaIndex]; !ok {
				continue
			}
			if _, ok := this.areaRelive[areaIndex]; !ok {
				this.areaRelive[areaIndex] = make([]int, 0)
			}
			this.areaRelive[areaIndex] = append(this.areaRelive[areaIndex], v.GetObjId())
		}
	}
}

func (this *PublicDabaoFight) OnDie(actor base.Actor, killer base.Actor) {

	if actor.GetType() == pb.SCENEOBJTYPE_MONSTER {

		monster := actor.(*actorPkg.MonsterActor)

		this.updateBossFamilyInfo(actor)

		areaIndex := monster.GetBirthAreaIndex()
		if num, ok := this.StageConf.Monster_num[areaIndex]; ok {

			aliveNum := 0
			for _, v := range this.areaRelive[areaIndex] {
				monsterActor := this.GetActorByObjId(v)
				if monsterActor.GetProp().HpNow() > 0 {
					aliveNum += 1
				}
			}
			if aliveNum < num {
				for _, v := range this.areaRelive[areaIndex] {
					actorObj := this.GetActorByObjId(v)
					if actorObj.GetProp().HpNow() <= 0 {

						monsterActor := actorObj.(*actorPkg.MonsterActor)
						monsterActor.SetOwner(0)
						actorObj.Relive(monsterActor.MonsterT.ReliveAddrType, constFight.RELIVE_TYPE_NOMAL)
					}
				}
			}
		}
	} else if actor.GetType() == pb.SCENEOBJTYPE_USER {
		if actor.(*actorPkg.UserActor).GetPlayer().CheckUserDie() {
			this.BossOwnerChange(actor, killer)
		}
	}
}

func (this *PublicDabaoFight) BossOwnerChange(dieActor, killer base.Actor) {

	for _, v := range this.monsterActors {
		if m, ok := v.(base.ActorMonster); ok {
			if m.Owner() == dieActor.GetUserId() {
				v.AddOwner(killer, true)
			}
		}
	}
}

func (this *PublicDabaoFight) OnEnd() {

}

func (this *PublicDabaoFight) OnRelive(actor base.Actor, reliveType int) {

	if actor.GetType() == base.ActorTypeMonster {

		//this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_RUNNING, 0)
		//this.sendFieldBossInfo()
	}
	this.updateBossFamilyInfo(actor)
}

func (this *PublicDabaoFight) sendFieldBossInfo() {

}

func (this *PublicDabaoFight) UpdateFrame() {

}

func (this *PublicDabaoFight) updateBossFamilyInfo(actor base.Actor) {
	if actor.GetType() != pb.SCENEOBJTYPE_MONSTER {
		return
	}
	monster := actor.(*actorPkg.MonsterActor)
	if monster.MonsterT.Type == constFight.MONSTER_TYPE_BOSS {
		bossFamilyConf := gamedb.GetBossFamilyBossFamilyCfg(this.GetStageConf().Id)
		if bossFamilyConf != nil {
			GetFightMgr().UpdateBossFamilInfo(this)
		}
	}
}

func (this *PublicDabaoFight) OnLeaveUser(userId int) {
	if this.GetStageConf().Type == constFight.FIGHT_TYPE_PUBLIC_DABAO_SINGLE {
		this.Stop()
	}
}
