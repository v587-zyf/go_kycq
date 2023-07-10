package worldBoss

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/managersI"
	"cqserver/golibs/common"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"time"
)

type WorldBoss struct {
	util.DefaultModule
	managersI.IModule
}

func NewWorldBoss(module managersI.IModule) *WorldBoss {
	w := &WorldBoss{IModule: module}
	return w
}

func makeMailBags(reward gamedb.ItemInfos) (bags []*model.Item) {
	for _, itemInfo := range reward {
		bags = append(bags, &model.Item{
			ItemId: itemInfo.ItemId,
			Count:  itemInfo.Count,
		})
	}
	return
}

func (this *WorldBoss) WorldBossInfoNtf() error {
	this.BroadcastAll(this.GetWorldBossInfo())
	return nil
}

func (this *WorldBoss) GetWorldBossInfo() *pb.WorldBossInfoNtf {
	worldBosses := gamedb.GetWorldBosses()
	var keyMap []int
	// 把所有key(key目前是stageId)存起来
	for k, _ := range worldBosses {
		keyMap = append(keyMap, k)
	}
	//logger.Debug("GetWorldBossInfo not sort keymap is %v", keyMap)
	// 冒泡排序，根据开放时间
	swap := true
	for i := 0; i < len(keyMap) && swap == true; i++ {
		swap = false
		for j := len(keyMap) - 1; j > i; j-- {
			if worldBosses[keyMap[j]].OpenTime.GetSecondsFromZero() < worldBosses[keyMap[j-1]].OpenTime.GetSecondsFromZero() {
				keyMap[j], keyMap[j-1] = keyMap[j-1], keyMap[j]
				swap = true
			}
		}
	}
	//logger.Debug("GetWorldBossInfo sort keymap is %v", keyMap)
	// 比对时间
	now := int(time.Now().Unix()) - common.GetZeroTimeUnixFrom1970()
	stageId := 0
	isNextDay := false
	for i := 0; i < len(keyMap); i++ {
		thisKey := keyMap[i]
		thisOpenTime := worldBosses[thisKey].OpenTime.GetSecondsFromZero()
		continueTime := thisOpenTime + worldBosses[thisKey].Continue*60
		// 还没达到开放时间或已经开启
		if now < continueTime {
			stageId = thisKey
			break
		}
	}
	if stageId == 0 {
		stageId = keyMap[0]
		isNextDay = true
	}
	//logger.Debug("stageId is %v", stageId)

	worldConf := gamedb.GetWorldBossByStageId(stageId)
	times := common.GetZeroTimeUnixFrom1970()
	if isNextDay {
		times += 24 * 60 * 60
	}
	openTime := worldConf.OpenTime.GetSecondsFromZero()
	prepareTime := worldConf.PrepareTime.GetSecondsFromZero()

	worldBossInfo := &pb.WorldBossInfoNtf{
		Id:          int32(worldConf.Id),
		PrepareTime: int32(times + prepareTime),
		OpenTime:    int32(times + openTime),
		CloseTime:   int32((times + openTime) + worldConf.Continue*60),
	}
	return worldBossInfo
}
