package buff

import (
	"cqserver/fightserver/internal/base"
	"cqserver/fightserver/internal/scene"
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/prop"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
)

func (this *BuffManager) createGroundBuff(buffConfT *gamedb.BuffBuffCfg, target base.Actor, skillId int) bool {

	if buffConfT.BuffType != pb.BUFFTYPE_GROUND_BUFF {
		return false
	}

	var point *scene.Point
	if target != nil {
		point = target.Point()
	}
	if point == nil {
		skillConf := gamedb.GetSkillLevelSkillCfg(skillId)
		if skillConf != nil {
			point = this.owner.GetScene().GetPointByDirAndMaxDis(this.owner.Point(), this.owner.GetDir(), skillConf.Range, false, nil)
		}
	}
	if point == nil {
		logger.Error("buffmanager 添加buff异常，玩家：%v,buffId:%v，释放坐标点为找到", this.owner.NickName(), buffConfT.Id)
		return true
	}

	areas := this.owner.GetScene().GetPointByRangeDis(point, buffConfT.Effect[1])
	if len(areas) <= 0 {
		logger.Error("buffmanager 添加buff异常，玩家：%v,buffId:%v，作用区域异常", this.owner.NickName(), point, buffConfT.Effect, buffConfT.Id)
		return true
	}
	groupBuff := &GroundBuff{}
	groupBuff.SceneBuff = scene.NewSceneBuff(buffConfT.Id,this.owner.GetUserId())
	groupBuff.DefaultBuff = NewDefaultBuff(buffConfT, this.owner, this.owner, groupBuff, 0)
	groupBuff.areas = areas
	err := groupBuff.EnterScene(this.owner.GetFight().GetScene(), point)
	if err != nil {
		logger.Error("buffmanager 添加buff异常,玩家：%v,buffId:%v，err:%v", this.owner.NickName(), buffConfT.Id, err)
	} else {
		logger.Error("buffmanager 添加buff,玩家：%v,buffId:%v", this.owner.NickName(), buffConfT.Id)
	}
	this.groupBuffs[groupBuff.GetObjId()] = groupBuff
	return true
}

type GroundBuff struct {
	*scene.SceneBuff
	*DefaultBuff
	lastRunTime int64
	areas       []*scene.Point //buff区域
}

//添加buff效果 单位时间改变
func (this *GroundBuff) OnAdd(buffHpChangeInfos *[]*pb.BuffHpChangeInfo, delBuffInfos *[]*pb.DelBuffInfo) {

}

func (this *GroundBuff) Run(buffHpChangeInfos *[]*pb.BuffHpChangeInfo, arg ...interface{}) {
	//流血的判断
	now := common.GetNowMillisecond()
	if now-this.lastRunTime < int64(this.buffT.Effect[constFight.BUFF_KEY_ZERO]) {
		return
	}
	this.lastRunTime = now
	targetNum := 0
	fight := this.owner.GetFight()

	propMaxId, propMinId := prop.GetAtkPropIdByJob(this.source.Job())
	baseValue := (this.source.GetProp().Get(propMaxId) + this.source.GetProp().Get(propMinId)) / 2
	hurt := int(float64(baseValue) * (float64(this.buffT.Effect[constFight.BUFF_KEY_THREE]) / 10000))

	for _, v := range this.areas {
		objs := v.GetAllObject()
		for _, obj := range objs {

			target := fight.GetActorByObjId(obj.GetObjId())
			if target == nil || !target.GetVisible() {
				continue
			}
			if !target.CanAttack() {
				continue
			}

			if !this.owner.IsEnemy(target) {
				continue
			}

			if wudi, _ := target.BuffHasType(pb.BUFFTYPE_INVINCIBLE, nil); wudi {
				continue
			}
			targetNum++
			//添加buff
			for i, l := constFight.BUFF_KEY_FOUR, len(this.buffT.Effect); i < l; i++ {
				target.AddBuff(this.buffT.Effect[i], this.owner, false)
			}

			if target.GetType() == pb.SCENEOBJTYPE_MONSTER {
				hurt = base.MonsterHurtLimit(this.owner,target,hurt)
			}

			//伤害
			if hurt > 0 {

				hurtEffect := &pb.BuffHpChangeInfo{}
				hurtEffect.Death = 0
				//记录玩家在战斗中
				target.SetInFightTime()
				realChangeHp, isDeath := target.ChangeHp(- hurt)
				hurtEffect.ChangeHp = int64(realChangeHp)
				hurtEffect.OwnerObjId = int32(target.GetObjId())
				hurtEffect.TotalHp = int64(target.GetProp().HpNow())
				hurtEffect.KillerId = int32(this.owner.GetObjId())
				nickName := this.owner.NickName()
				if leader, ok := this.owner.(base.ActorLeader); ok {
					nickName = leader.GetLeader().NickName()
				}
				hurtEffect.KillerName = nickName
				if isDeath {
					target.SetVisible(false)
					target.SetKiller(this.owner)
					hurtEffect.Death = 1
				}
				*buffHpChangeInfos = append(*buffHpChangeInfos, hurtEffect)
			}
			if targetNum >= this.buffT.Effect[constFight.BUFF_KEY_TWO] {
				return
			}
		}
	}
}
