package killMonster

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/modelGame"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"sync"
	"time"
)

type KillMonster struct {
	util.DefaultModule
	managersI.IModule

	KillMonsterData map[int]*modelGame.KillMonsterDb
	Mu              sync.Mutex
	updateChan      chan *modelGame.KillMonsterDb
}

func NewKillMonsterManager(m managersI.IModule) *KillMonster {
	return &KillMonster{
		IModule:         m,
		KillMonsterData: make(map[int]*modelGame.KillMonsterDb),
	}
}

func (this *KillMonster) Init() error {
	killMonsterModel := modelGame.GetKillMonsterDbModel()
	allData, err := killMonsterModel.LoadAlLData()
	if err != nil {
		return err
	}
	for _, data := range allData {
		this.KillMonsterData[data.StageId] = data
	}

	this.updateChan = make(chan *modelGame.KillMonsterDb, 1000)
	go this.updateService()
	return nil
}

/**
 *  @Description: 记录击杀
 *  @param user
 *  @param stageId
 */
func (this *KillMonster) WriteKillMonster(user *objs.User, stageId int) {
	this.Mu.Lock()
	defer this.Mu.Unlock()
	monsterDb, ok := this.KillMonsterData[stageId]
	if ok {
		monsterDb.KillNumAll++
		this.updateChan <- monsterDb
	} else {
		newData := &modelGame.KillMonsterDb{
			StageId:         stageId,
			FirstKillUserId: user.Id,
			FirstKillTime:   time.Now(),
			KillNumAll:      1,
		}
		if err := modelGame.GetKillMonsterDbModel().Create(newData); err != nil {
			logger.Error("创建击杀怪物数据错误 stageId:%v err:%v", stageId, err)
			return
		}
		this.KillMonsterData[stageId] = newData
		monsterDb = newData

		this.BroadcastAll(&pb.KillMonsterUniKillNtf{StageId: int32(stageId), KlillUserId: int32(user.Id), KillUserName: user.NickName})
	}
	if perCfg := gamedb.GetKillMonsterPerByStageId(stageId); perCfg != nil {
		this.GetUserManager().SendMessage(user, &pb.KillMonsterPerKillNtf{StageId: int32(stageId)}, true)
	}
	this.BroadcastAll(&pb.KillMonsterMilKillNtf{StageId: int32(stageId), KillNum: int32(monsterDb.KillNumAll)})
}

func (this *KillMonster) updateService() {
	killMonsterModel := modelGame.GetKillMonsterDbModel()
	for {
		select {
		case db := <-this.updateChan:
			if err := killMonsterModel.Update(db); err != nil {
				logger.Error("更新击杀怪物数据错误 err:%v", err)
			}
		}
	}
}
