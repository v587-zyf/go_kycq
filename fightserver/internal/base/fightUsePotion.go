package base

import "time"

type IFightUsePotion interface {
	FightPotionUserInto(userId int)
	GetUserPotionCooldown(userId int) int64
	FightUsePotionFunc(fight Fight, userId int)
}

//
//  @Description: 战斗使用药水
//
type FightUsePotion struct {
	potionUsers map[int]int64 //玩家->上次使用时间
}

func NewFightUsePotion() *FightUsePotion {
	f := &FightUsePotion{
		potionUsers: make(map[int]int64),
	}
	return f
}

func (fd *FightUsePotion) FightPotionUserInto(userId int) {
	if userId != 0 {
		if _, ok := fd.potionUsers[userId]; !ok {
			fd.potionUsers[userId] = 0
		}
	}
}

func (fd *FightUsePotion) GetUserPotionCooldown(userId int) int64 {

	return fd.potionUsers[userId]
}

func (fd *FightUsePotion) FightUsePotionFunc(fight Fight, userId int) {

	fd.potionUsers[userId] = time.Now().Unix()
	fight.OnUsePotion(userId)
}
