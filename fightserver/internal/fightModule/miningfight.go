package fightModule

import (
	"cqserver/fightserver/internal/actorPkg"
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/net"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
)

type MiningFight struct {
	*DefaultFight
	fightOwner    base.Actor
	miningId      int  //抢夺/夺回矿Id
	defenderUseId int  //防守方id
	IsRetake      bool //是否夺回战
	inFight       bool
}

func NewMiningFight(stageId int, cpData []byte) (*MiningFight, error) {
	var err error
	f := &MiningFight{}
	f.DefaultFight, err = NewDefaultFight(stageId, f)
	if err != nil {
		return nil, err
	}
	f.SetLifeTime(-1)
	f.InitMonster()
	f.Start()

	return f, nil
}

func (this *MiningFight) OnDie(actor base.Actor, killer base.Actor) {
	allDie := this.CheckUserAllDie(actor)
	if allDie {
		result := pb.RESULTFLAG_SUCCESS
		if actor.GetUserId() == this.fightOwner.GetUserId() {
			result = pb.RESULTFLAG_FAIL
		}
		actors := this.GetUserActors()
		for _, v := range actors {
			if v.TeamIndex() == constFight.FIGHT_TEAM_ZERO {
				v.GetFSM().Pause()
			}
		}
		this.sendFightResult(result)
	}
}

func (this *MiningFight) OnActorEnter(actor base.Actor) {

	this.DefaultFight.OnActorEnter(actor)
	if actor.SessionId() != 0 {
		this.fightOwner = actor
	}
	if (actor.GetType() == pb.SCENEOBJTYPE_SUMMON || actor.GetType() == pb.SCENEOBJTYPE_PET) && actor.GetUserId() != this.fightOwner.GetUserId() {
		actor.GetFSM().Recover()
	}
}

func (this *MiningFight) OnLeaveUser(userId int) {
	if this.inFight {
		this.sendFightResult(pb.RESULTFLAG_FAIL)
		//logger.Info("发送game挖矿战斗结果,玩家离开 服务器：%v，结果：%v", userId)
	}
	if userId == this.fightOwner.GetUserId() {

		this.Stop()
	}
}

func (this *MiningFight) NewFightInfo(info *pbserver.MiningNewFightInfoReq) bool {

	if this.inFight {
		return false
	}

	this.IsRetake = info.IsRetake
	this.inFight = true
	this.miningId = int(info.MiningId)
	this.defenderUseId = int(info.MiningUserId)
	return true
}

func (this *MiningFight) resetFight() {
	this.IsRetake = false
	this.inFight = false
	this.miningId = 0
	this.defenderUseId = 0

	fitActor := this.GetUserFitActor(this.fightOwner.GetUserId())
	if fitActor != nil {
		fitActor.(*actorPkg.FitActor).ResetHeroHpFix()
		if fitActor.GetProp().HpNow() < fitActor.GetProp().Get(pb.PROPERTY_HP) {
			fitActor.GetProp().SetHpNow(fitActor.GetProp().Get(pb.PROPERTY_HP))
			HPChangeNtf := &pb.SceneObjHpNtf{
				ObjId:   int32(fitActor.GetObjId()),
				Hp:      int64(fitActor.GetProp().HpNow()),
				TotalHp: int64(fitActor.GetProp().Get(pb.PROPERTY_HP)),
			}
			this.GetScene().NotifyAll(HPChangeNtf)
		}
	} else {
		userActors := this.GetUserByUserId(this.fightOwner.GetUserId())
		for _, v := range userActors {
			if v.IsDeath() {
				v.Relive(constFight.RELIVE_ADDR_TYPE_SITU, constFight.RELIVE_TYPE_NOMAL)
			} else {
				if v.GetProp().HpNow() < v.GetProp().Get(pb.PROPERTY_HP) {
					v.GetProp().SetHpNow(v.GetProp().Get(pb.PROPERTY_HP))
					HPChangeNtf := &pb.SceneObjHpNtf{
						ObjId:   int32(v.GetObjId()),
						Hp:      int64(v.GetProp().HpNow()),
						TotalHp: int64(v.GetProp().Get(pb.PROPERTY_HP)),
					}
					this.GetScene().NotifyAll(HPChangeNtf)
				}
			}
		}
		petActor := this.GetPetActor(this.fightOwner.GetUserId())
		if petActor != nil {
			point := this.getPetEnterScenePoint(this.fightOwner.GetUserId(), petActor)
			petActor.SetVisible(true)
			petActor.MoveTo(point, pb.MOVETYPE_WALK, true, false)
			appearNtf := petActor.BuildAppearMessage()
			this.GetScene().NotifyNearby(petActor, appearNtf, nil)
		}
	}

	actors := this.GetUserActors()
	for _, v := range actors {
		if v.TeamIndex() == constFight.FIGHT_TEAM_ZERO {
			this.LeaveUser(v.GetUserId())
			break
		}
	}
}

func (this *MiningFight) sendFightResult(result int) {

	if !this.inFight {
		return
	}

	msg := &pbserver.FSFightEndNtf{
		FightType: int32(this.StageConf.Type),
		StageId:   int32(this.StageConf.Id),
	}

	resultMsg := &pbserver.MiningFightResultNtf{
		MiningUserId: int32(this.defenderUseId),
		MiningId:     int32(this.miningId),
		Result:       int32(result),
		IsRetake:     this.IsRetake,
	}

	if this.fightOwner != nil {
		resultMsg.UserId = int32(this.fightOwner.GetUserId())
	}

	rb, _ := resultMsg.Marshal()
	msg.CpData = rb
	logger.Info("发送game挖矿掠夺、夺回战斗结果,服务器：%v，结果：%v,resultMsg:%v", this.fightOwner.HostId(), *msg, resultMsg)
	net.GetGsConn().SendMessageToGs(uint32(this.fightOwner.HostId()), msg)
	this.resetFight()
}
