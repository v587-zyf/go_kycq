package mining

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constMax"
	"cqserver/gamelibs/publicCon/constMining"
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/objs"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"runtime/debug"
	"time"
)

func (this *MiningManager) updateMiningRobotServices() {
	needAddTime := gamedb.GetConf().MiningRobot
	this.addMiningRobot()

	ticker := time.NewTicker(time.Second * time.Duration(needAddTime))
	defer func() {
		ticker.Stop()
		if r := recover(); r != nil {
			stackBytes := debug.Stack()
			logger.Error("panic when messagePump:%v,%s", r, stackBytes)
		}
	}()
	for {
		select {
		case <-ticker.C:
			this.addMiningRobot()
		}
	}
}

func (this *MiningManager) addMiningRobot() {
	this.MiningMu.Lock()
	defer this.MiningMu.Unlock()

	nowTime := time.Now()
	for robotId, mining := range this.RobotMap {
		if mining.ExpireTime.Unix() <= nowTime.Unix() {
			mining.ExpireTime = nowTime
			mining.ReceiveTime = nowTime
			mining.FindTime = nowTime
			mining.DeletedAt = nowTime
			delete(this.RobotMap, robotId)
			this.UpdateMiningDate(mining.MiningDb, true)
		}
	}
	miningNum := 0
	minerMaxLv := gamedb.GetMaxValById(0, constMax.MAX_MINING_LEVEL)
	for _, mining := range this.MiningMap {
		if mining.Miner < minerMaxLv && mining.ExpireTime.Unix() > nowTime.Unix() {
			miningNum++
		}
	}
	needAddNum := gamedb.GetConf().MiningRobotcondition - (miningNum + len(this.RobotMap))
	if needAddNum < 1 {
		return
	}
	robotCfgs := gamedb.GetRobotCfgs()
	okNum := 0
	for i := 1; i < 10000; i++ {
		robotId := common.RandNum(1, len(robotCfgs))
		if _, ok := this.RobotMap[-robotId]; ok {
			continue
		}
		robotCfg := robotCfgs[robotId]
		miner := robotCfgs[robotId].MiningLevel
		miningCfg := gamedb.GetMiningLvCfg(miner)
		addData := this.MakeMiningDbDate(-robotId, miner, nowTime, miningCfg.Time)
		addData.IsRobot = constMining.MINING_DATA_ROBOT_YES
		this.AddMiningData(addData, true)
		newMiningObj := objs.NewMining(addData)
		newMiningObj.Combat = this.calcRobotCombat(robotCfg)
		this.RobotMap[addData.UserId] = newMiningObj
		okNum++
		if okNum >= needAddNum {
			break
		}
	}
}

func (this *MiningManager) calcRobotCombat(robotCfg *gamedb.RobotRobotCfg) int {
	allCombat := 0
	for index, heroIndex := range constUser.USER_HERO_INDEX {
		property := make(map[int]int)
		switch heroIndex {
		case constUser.USER_HERO_MAIN_INDEX:
			property = robotCfg.Property1
		case constUser.USER_HERO_SECOND_INDEX:
			property = robotCfg.Property2
		case constUser.USER_HERO_THREE_INDEX:
			property = robotCfg.Property3
		}
		combat := builder.CalcCombat(robotCfg.Job[index], 0, -1, property, nil)
		allCombat += combat
	}
	return allCombat
}
