package base

import (
	"cqserver/fightserver/internal/scene"
	"cqserver/gamelibs/fsm"
	"cqserver/golibs/common"
	"cqserver/protobuf/pb"
	"math"
)

const (
	_                  fsm.Trigger = iota
	StateTriggerActive             //活动（随机走动 回家 ）
	StateTriggerFight              //战斗（释放技能，目标移动）
	StateTriggerDeath              //死亡
)

func BuildDefaltSceneObjMessage(inactor Actor) *pb.SceneObj {
	r := &pb.SceneObj{}
	r.ObjType = int32(inactor.GetType())
	r.Point = inactor.Point().ToPbPoint()
	r.Dir = int32(inactor.GetDir())
	r.ObjId = int32(inactor.GetObjId())

	r.Hp = int64(inactor.GetProp().HpNow())
	r.HpMax = int64(inactor.GetProp().Get(pb.PROPERTY_HP))
	r.Mp = int64(inactor.GetProp().MpNow())
	r.MpMax = int64(inactor.GetProp().Get(pb.PROPERTY_MP))
	//预留buff
	r.Buffs = inactor.GetAllBuffsPb()
	if inactor.GetType() != pb.SCENEOBJTYPE_MONSTER {
		if leader, ok := inactor.(ActorLeader); ok {
			r.ServerId = int32(leader.GetLeader().HostId())
		}
	}
	r.TeamId = int32(inactor.TeamIndex())
	return r
}

//常用算法1,万分比
//公式为:【  min(max（（(att-def）/（10000+def））*10000,minLimit),attLimit)  】
func commonRateAlgorithm(att, def int, minLimit int, maxLimit int) float64 {
	return math.Min(math.Max(float64(att-def)/float64(10000+def)*10000, float64(minLimit)), float64(maxLimit)) / 10000
}

func SkillAttackEffect(attacker Actor, skill *Skill, dir int, effects []*pb.HurtEffect, err int) *pb.AttackEffectNtf {

	nickName := attacker.NickName()
	if leader, ok := attacker.(ActorLeader); ok {
		nickName = leader.GetLeader().NickName()
	}
	attackEffectNtf := &pb.AttackEffectNtf{AttackerId: int32(attacker.GetObjId()), AttackerName: nickName, Hurts: make([]*pb.HurtEffect, 0)}
	if skill != nil {
		attackEffectNtf.SkillId = int32(skill.Skillid)
		attackEffectNtf.SkillLv = int32(skill.LevelT.Skillid)
		attackEffectNtf.SkillStartT = skill.GetSkillUseTime()
		attackEffectNtf.SkillStopT = skill.GetNextAttackTime()
	}
	attackEffectNtf.Point = attacker.Point().ToPbPoint()
	attackEffectNtf.Hurts = effects
	attackEffectNtf.Dir = int32(dir)
	attackEffectNtf.ServerTime = common.GetNowMillisecond()
	attackEffectNtf.Err = int32(err)
	attackEffectNtf.MpNow = int64(attacker.GetProp().MpNow())
	attackEffectNtf.HpNow = int64(attacker.GetProp().HpNow())
	return attackEffectNtf
}

func NotifyAttackEffect(attacker Actor, skill *Skill, dir int, effects []*pb.HurtEffect, castPoint, moveToPoint *scene.Point, isElf bool) *pb.AttackEffectNtf {

	attackEffectNtf := SkillAttackEffect(attacker, skill, dir, effects, 0)
	var excludeSession map[uint32]bool

	if attacker.GetUserId() > 0 && skill.Type != pb.SKILLTYPE_PASSIVE2 {
		sessionId := attacker.SessionId()
		if sessionId <= 0 {
			mainActor := attacker.GetFight().GetUserMainActor(attacker.GetUserId())
			sessionId = mainActor.SessionId()
		}
		if sessionId > 0 {
			excludeSession = map[uint32]bool{sessionId: true}
		}
	}
	if castPoint != nil {
		attackEffectNtf.Point = castPoint.ToPbPoint()
	}
	if moveToPoint != nil {
		attackEffectNtf.MoveToPoint = moveToPoint.ToPbPoint()
	}
	attackEffectNtf.IsElf = isElf
	attacker.NotifyNearby(attacker, attackEffectNtf, excludeSession)

	return attackEffectNtf
}
