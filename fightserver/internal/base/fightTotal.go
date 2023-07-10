package base

import (
	"cqserver/protobuf/pb"
)

type IFightTotal interface {
	TotalDamage(attacker, defender Actor, hurt int)
	FightTotalInto(userId int)
}

type FightTotalData struct {
	damage   int //伤害输出
	beDamage int //承受伤害
	treat    int //治疗
}

//
//  @Description: 战斗统计类
//
type FightTotal struct {
	users map[int]*FightTotalData
}

func NewFightTotal() *FightTotal {
	fightTotal := &FightTotal{
		users: make(map[int]*FightTotalData),
	}
	return fightTotal
}

func (fd *FightTotal) FightTotalInto(userId int) {
	if userId != 0 {
		fd.users[userId] = &FightTotalData{}
	}
}

func (fd *FightTotal) TotalDamage(attacker, defender Actor, hurt int) {

	if attacker != nil && (attacker.GetType() == pb.SCENEOBJTYPE_USER ||
		attacker.GetType() == pb.SCENEOBJTYPE_FIT || attacker.GetType() == pb.SCENEOBJTYPE_PET || attacker.GetType() == pb.SCENEOBJTYPE_SUMMON) {
		userFightTotal := fd.getUserFightTotalData(attacker.GetUserId())
		userFightTotal.damage += hurt
	}

	if defender != nil && (attacker.GetType() == pb.SCENEOBJTYPE_USER || attacker.GetType() == pb.SCENEOBJTYPE_FIT || defender.GetType() == pb.SCENEOBJTYPE_SUMMON) {
		userFightTotal := fd.getUserFightTotalData(defender.GetUserId())
		userFightTotal.beDamage += hurt
	}
}

func (this *FightTotal) getUserFightTotalData(userId int) *FightTotalData {

	if _, ok := this.users[userId]; !ok {
		this.users[userId] = &FightTotalData{}
	}
	return this.users[userId]
}
