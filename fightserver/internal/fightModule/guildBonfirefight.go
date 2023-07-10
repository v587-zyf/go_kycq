package fightModule

import (
	"cqserver/fightserver/internal/net"
	"cqserver/gamelibs/gamedb"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pbserver"
	"time"
)

type GuildBonfireFight struct {
	*DefaultFight
	userExps      map[int]int //获得经验数量
	lastCountTime int64
	guildId       int
}

func NewGuildBonfireFight(stageId int,guildId int) (*GuildBonfireFight, error) {
	var err error
	f := &GuildBonfireFight{
		userExps:      make(map[int]int),
		lastCountTime: time.Now().Unix(),
		guildId: guildId,
	}
	f.DefaultFight, err = NewDefaultFight(stageId, f)
	if err != nil {
		return nil, err
	}
	f.InitMonster()

	activityCfg := gamedb.GetGuildActivityGuildActivityCfg(1)
	lifeTime := activityCfg.CloseTime.GetSecondsFromZero() - common.GetTimeSeconds(time.Now())
	f.SetLifeTime(int64(lifeTime))
	f.Start()
	return f, nil
}

func (this *GuildBonfireFight) GetFightExtMark() int {
	return this.guildId
}

func (this *GuildBonfireFight) UpdateFrame() {

	now := time.Now().Unix()
	if now-this.lastCountTime >= int64(gamedb.GetConf().BonefireTime) {
		this.lastCountTime = now
		objsMap := make(map[int]bool)
		mapConf := gamedb.GetMapMapCfg(this.StageConf.Mapid)
		for _, v := range mapConf.Special {
			objs := this.Scene.GetSceneRectObjs(v)
			for _, v := range objs {
				obj := this.GetActorByObjId(v)
				if obj != nil && obj.GetUserId() > 0 {
					isDie := this.CheckUserAllDie(obj)
					if !isDie {
						objsMap[obj.GetUserId()] = true
					}
				}
			}
		}
		//计算经验 增加经验 暂时game计算了
		ntf := &pbserver.GuildbonfireExpAddNtf{
			GuildId: int32(this.guildId),
			UserIds: make([]int32, 0),
		}
		for k, _ := range objsMap {
			ntf.UserIds = append(ntf.UserIds, int32(k))
		}
		net.GetGsConn().SendMessage(ntf)
	}
}

func (this *GuildBonfireFight) OnEnd() {
	//经验副本里面肯定只有一个人
	msg := &pbserver.FSFightEndNtf{
		FightType: int32(this.StageConf.Type),
		StageId:   int32(this.StageConf.Id),
		UseTime:   int32(time.Now().Unix() - this.createTime),
	}
	userIds := this.GetPlayerUserids()
	if userIds != nil {
		msg.Winners = common.ConvertIntSlice2Int32Slice(userIds)
	}
	logger.Info("发送game公会篝火战斗结果,服务器：%v，结果：%v", *msg)
	net.GetGsConn().SendMessage(msg)
	this.SetFightStatusAndNextStatusTime(FIGHT_STATUS_CLOSING, 15)
}
