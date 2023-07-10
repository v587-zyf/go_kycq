package base

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/protobuf/pb"
	"time"
)

type IFightCheer interface {
	FightCheerUserInto(userId int)
	/**
	 *  @Description: 获取鼓舞次数
	 *  @param userId
	 *  @param guildId
	 *  @return int 玩家鼓舞次数
	 *  @return int 门派鼓舞次数
	 */
	GetCheerNum(userId, guildId int) (int, int)
	FightUseCheer(fight Fight, userId int)
}

//
//  @Description: 战斗鼓舞类
//
type FightCheer struct {
	context    IFightCheer
	cheerUsers map[int]int //玩家 鼓舞次数
}

func NewFightCheer(context IFightCheer) *FightCheer {
	f := &FightCheer{
		context:    context,
		cheerUsers: make(map[int]int),
	}
	return f
}

func (fd *FightCheer) FightCheerUserInto(userId int) {
	if userId != 0 {
		if _, ok := fd.cheerUsers[userId]; !ok {
			fd.cheerUsers[userId] = 0
		}
	}
}

func (fd *FightCheer) GetCheerNum(userId, guildId int) (int, int) {

	return fd.context.GetCheerNum(userId, guildId)
}

func (fd *FightCheer) FightUseCheer(fight Fight, userId int) {

	fd.cheerUsers[userId] += 1
	fight.OnCheer(userId)
}

type CheerBuffUnit struct {
	BuffId     int
	CreateTime int64
}

type FightCheerByGuild struct {
	*FightCheer
	CheerBuff map[int][]*CheerBuffUnit //鼓舞buff(guildId->buff)
	cheerNum  map[int]int
}

func NewFightCheerByGuild() *FightCheerByGuild {
	f := &FightCheerByGuild{
		CheerBuff: make(map[int][]*CheerBuffUnit),
		cheerNum:  make(map[int]int),
	}
	f.FightCheer = NewFightCheer(f)
	return f
}

/**
*  @Description: 直接鼓舞
*  @receiver fd
*  @param fight
*  @param guildId
*  @param buffId
**/
func (fd *FightCheerByGuild) GuildCheer(fight Fight, guildId, buffId int) {

	fd.GuildCheerBuff(guildId, buffId)
	fd.cheerNum[guildId] += 1
	ntf := &pb.FightCheerNumChangeNtf{
		GuildId:       int32(guildId),
		GuildCheerNum: int32(fd.cheerNum[guildId]),
	}
	fight.GetScene().NotifyAll(ntf)
}

/**
*  @Description: 其他途径buff
*  @receiver fd
*  @param fight
*  @param guildId
*  @param buffId
**/
func (fd *FightCheerByGuild) GuildCheerBuff(guildId, buffId int) {

	if _, ok := fd.CheerBuff[guildId]; !ok {
		fd.CheerBuff[guildId] = make([]*CheerBuffUnit, 0)
	}
	fd.CheerBuff[guildId] = append(fd.CheerBuff[guildId], &CheerBuffUnit{buffId, time.Now().Unix()})

}

func (fd *FightCheerByGuild) OnGuildActorEnter(actor Actor) {

	if actor.GetType() == pb.SCENEOBJTYPE_USER {
		guildId := actor.(ActorUser).GuildId()
		if buffs, ok := fd.CheerBuff[guildId]; ok {
			now := time.Now().Unix()
			for _, v := range buffs {
				buffConf := gamedb.GetBuffBuffCfg(v.BuffId)
				if v.CreateTime+int64(buffConf.Time/1000) > now {
					actor.AddNewBuff(v.BuffId, actor, nil, false)
				}
			}
		}
	}
}
func (fd *FightCheerByGuild) GetCheerNum(userId, guildId int) (int, int) {
	usercheerNum := fd.cheerUsers[userId]
	guildCheerNum := 0
	if _, ok := fd.cheerNum[guildId]; ok {
		guildCheerNum = fd.cheerNum[guildId]
	}
	return usercheerNum, guildCheerNum
}
