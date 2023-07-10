package mining

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/modelGame"
	"cqserver/gamelibs/publicCon/constMail"
	"cqserver/gamelibs/publicCon/constMax"
	"cqserver/gamelibs/publicCon/constMining"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"database/sql"
	"sync"
	"time"
)

type MiningManager struct {
	util.DefaultModule
	managersI.IModule

	MiningMap map[int]*objs.Mining
	MiningMu  sync.Mutex
	RobotMap  map[int]*objs.Mining
}

func NewMiningManager(module managersI.IModule) *MiningManager {
	return &MiningManager{IModule: module}
}

func (this *MiningManager) Init() error {
	this.MiningMap = make(map[int]*objs.Mining)
	this.RobotMap = make(map[int]*objs.Mining)

	timeNow := time.Now()
	robotCfgs := gamedb.GetRobotCfgs()
	allMiningData, err := modelGame.GetMiningModel().GetMiningAll()
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	for _, mining := range allMiningData {
		this.MiningMu.Lock()
		if mining.IsRobot == constMining.MINING_DATA_ROBOT_NO {
			this.MiningMap[mining.UserId] = objs.NewMining(mining)
		} else if mining.IsRobot == constMining.MINING_DATA_ROBOT_YES {
			if mining.ExpireTime.Unix() <= timeNow.Unix() {
				mining.ExpireTime = timeNow
				mining.ReceiveTime = timeNow
				mining.FindTime = timeNow
				mining.DeletedAt = timeNow
				this.UpdateMiningDate(mining, true)
			} else {
				newMining := objs.NewMining(mining)
				newMining.Combat = this.calcRobotCombat(robotCfgs[-mining.UserId])
				this.RobotMap[mining.UserId] = newMining
			}
		}
		this.MiningMu.Unlock()
	}

	go this.updateMiningRobotServices()
	return nil
}

func (this *MiningManager) Online(user *objs.User) {
	date := common.GetResetTime(time.Now())
	this.MiningReset(user, date)
}

func (this *MiningManager) MiningReset(user *objs.User, date int) {
	if user.Mining.ResetTime != date {
		timeNow := time.Now()
		mining := this.MiningMap[user.Id]
		if mining != nil && mining.ExpireTime.Unix() <= time.Now().Unix() {
			getRewardMap := this.GetRewardMap(mining.MiningDb)
			bags := make([]*model.Item, 0)
			for itemId, count := range getRewardMap {
				bags = append(bags, &model.Item{
					ItemId: itemId,
					Count:  count,
				})
			}
			err := this.GetMail().SendSystemMail(user.Id, constMail.MINING_REWARD_ID, []string{}, bags, 0)
			if err != nil {
				logger.Error("MiningReset sendMail err:%v", err)
			}
			mining.ReceiveTime = timeNow
			this.MiningMu.Lock()
			this.DelMiningData(mining.MiningDb)
			this.MiningMu.Unlock()
			this.GetWarOrder().WriteWarOrderTask(user, pb.WARORDERCONDITION_MINING_NUM, []int{1})
		}
		user.Mining = &model.Mining{
			ResetTime: date,
		}
	}
	if user.Mining.WorkTime != 0 && this.MiningMap[user.Id] == nil {
		user.Mining.WorkTime = 0
		user.Mining.Miner = 0
	}
}

/**
 *  @Description: 升级矿工
 *  @param user
 *  @param isMax	是否一键升级
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *MiningManager) UpMiner(user *objs.User, isMax bool, op *ophelper.OpBagHelperDefault, ack *pb.MiningUpMinerAck) error {
	userMining := user.Mining

	maxLv := gamedb.GetMaxValById(0, constMax.MAX_MINING_LEVEL)
	if userMining.Miner >= maxLv {
		return gamedb.ERRLVENOUGH
	}
	lvCfg := gamedb.GetMiningLvCfg(userMining.Miner)
	consumes := lvCfg.Consume
	if isMax {
		consumes = lvCfg.ConsumeMax
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, consumes); err != nil {
		return err
	}

	if isMax {
		userMining.Miner = maxLv
	} else {
		isUp := false
		if userMining.Luck >= lvCfg.Lucky {
			userMining.Miner++
			userMining.Luck -= lvCfg.Lucky
		} else {
			isUp = common.RandByTenShousand(lvCfg.Probability)
			if !isUp {
				userMining.Luck++
			} else {
				userMining.Miner++
				userMining.Luck = 0
			}
		}
		ack.IsUp = isUp
	}
	user.Dirty = true

	ack.Luck = int32(userMining.Luck)
	ack.Miner = int32(userMining.Miner)
	return nil
}

/**
 *  @Description: 购买挖矿次数
 *  @param user
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *MiningManager) BuyNum(user *objs.User, op *ophelper.OpBagHelperDefault, ack *pb.MiningBuyNumAck) error {
	userMining := user.Mining
	maxBuyNum := gamedb.GetConf().MiningBuyMaxNum
	if userMining.BuyNum >= maxBuyNum+this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_MINING_BUYNUM) {
		return gamedb.ERRPURCHASECAPENOUGH
	}

	buyConsume := gamedb.GetConf().MiningBuy
	consumeMap := make(map[int]int)
	for _, itemInfo := range buyConsume {
		num, _ := this.GetBag().GetItemNum(user, itemInfo.ItemId)
		if num < itemInfo.Count {
			return gamedb.ERRNOTENOUGHGOODS
		}
		consumeMap[itemInfo.ItemId] = itemInfo.Count
	}
	for itemId, count := range consumeMap {
		this.GetBag().Remove(user, op, itemId, count)
	}
	userMining.BuyNum++
	user.Dirty = true

	ack.BuyNum = int32(userMining.BuyNum)
	return nil
}

/**
 *  @Description: 开始挖矿
 *  @param user
 *  @param ack
 *  @return error
 */
func (this *MiningManager) StartWork(user *objs.User, ack *pb.MiningStartAck) error {
	userMining := user.Mining
	lvCfg := gamedb.GetMiningLvCfg(userMining.Miner)
	if userMining.WorkTime != 0 {
		expireTime := userMining.WorkTime + lvCfg.Time
		if int64(expireTime) < time.Now().Unix() {
			return gamedb.ERRDRAW
		} else {
			return gamedb.ERRMINING
		}
	}
	maxNum := gamedb.GetConf().MiningWorkMaxNum + this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_MINING_FREENUM)
	hasNum := maxNum + userMining.BuyNum - userMining.WorkNum
	if hasNum <= 0 {
		return gamedb.ERRMININGWORKMAX
	}

	nowTime := time.Now()
	userMining.WorkTime = int(nowTime.Unix())
	userMining.WorkNum++
	user.Dirty = true

	this.MiningMu.Lock()
	this.AddMiningData(this.MakeMiningDbDate(user.Id, userMining.Miner, nowTime, lvCfg.Time), false)
	this.MiningMu.Unlock()
	kyEvent.Mining(user, userMining.Miner, userMining.WorkNum, lvCfg.Reward)

	ack.WorkNum = int32(userMining.WorkNum)
	ack.WorkTime = int64(userMining.WorkTime)
	this.GetDailyTask().CompletionOfTask(user, pb.DAILYTASKACTIVITYTYPE_KUANG_DONG_ZHENG_DUO, 1)
	this.GetCondition().RecordCondition(user, pb.CONDITION_GO_TO_WA_KUANG, []int{1})
	this.GetTask().AddTaskProcess(user, pb.CONDITION_GO_TO_WA_KUANG, -1)
	return nil
}

/**
 *  @Description: 挖矿列表
 *  @param ack
 */
func (this *MiningManager) List(ack *pb.MiningListAck) {
	pbList := make(map[int64]*pb.MiningListInfo, 0)
	minerMaxLv := gamedb.GetMaxValById(0, constMax.MAX_MINING_LEVEL)
	this.MiningMu.Lock()
	defer this.MiningMu.Unlock()
	for userId, mining := range this.MiningMap {
		userInfo := this.GetUserManager().GetUserBasicInfo(userId)
		if mining.ExpireTime.Unix() <= time.Now().Unix() || mining.Miner >= minerMaxLv || !mining.Rtime.IsZero() {
			continue
		}
		pbList[int64(mining.Id)] = &pb.MiningListInfo{
			Uid:    int32(userId),
			Name:   userInfo.NickName,
			Combat: int64(userInfo.Combat),
			Time:   mining.WorkTime.Unix(),
			Miner:  int32(mining.Miner),
			Id:     int64(mining.Id),
		}
	}
	for robotId, mining := range this.RobotMap {
		robotInfo := gamedb.GetRobotRobotCfg(-robotId)
		if mining.ExpireTime.Unix() <= time.Now().Unix() || mining.Miner >= minerMaxLv || !mining.Rtime.IsZero() {
			continue
		}
		pbList[int64(mining.Id)] = &pb.MiningListInfo{
			Uid:    int32(robotId),
			Name:   robotInfo.Name,
			Combat: int64(mining.Combat),
			Time:   mining.WorkTime.Unix(),
			Miner:  int32(mining.Miner),
			Id:     int64(mining.Id),
		}
	}
	ack.MiningList = pbList
}

/**
 *  @Description: 领取奖励
 *  @param user
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *MiningManager) Draw(user *objs.User, op *ophelper.OpBagHelperDefault, ack *pb.MiningDrawAck) error {
	miningInfo := this.MiningMap[user.Id]
	if miningInfo == nil {
		return gamedb.ERRMININGNOTREWARD
	}

	if time.Now().Unix() < miningInfo.ExpireTime.Unix() {
		return gamedb.ERRMINING
	}
	if !miningInfo.ReceiveTime.IsZero() {
		return gamedb.ERRMININGNOTREWARD
	}

	getRewardMap := this.GetRewardMap(miningInfo.MiningDb)
	items := make(gamedb.ItemInfos, 0)
	for itemId, count := range getRewardMap {
		items = append(items, &gamedb.ItemInfo{
			ItemId: itemId,
			Count:  count,
		})
	}
	this.GetBag().AddItems(user, items, op)
	userMining := user.Mining
	userMining.WorkTime = 0
	userMining.Miner = 0
	user.Dirty = true

	miningInfo.ReceiveTime = time.Now()
	miningInfo.DeletedAt = time.Now()
	this.MiningMu.Lock()
	this.DelMiningData(miningInfo.MiningDb)
	this.MiningMu.Unlock()

	ack.Goods = op.ToChangeItems()
	this.GetWarOrder().WriteWarOrderTask(user, pb.WARORDERCONDITION_MINING_NUM, []int{1})
	return nil
}

func (this *MiningManager) GetRewardMap(mining *modelGame.MiningDb) map[int]int {
	lvCfg := gamedb.GetMiningLvCfg(mining.Miner)
	getRewardMap := make(map[int]int)
	for itemId, count := range lvCfg.Reward {
		if !mining.Rtime.IsZero() {
			count -= lvCfg.Lose[itemId]
		}
		getRewardMap[itemId] = count
	}
	return getRewardMap
}

/**
 *  @Description: 是否被抢夺
 *  @param user
 *  @return bool
 */
func (this *MiningManager) GetDrawStatus(user *objs.User) bool {
	if this.MiningMap[user.Id] != nil && !this.MiningMap[user.Id].Rtime.IsZero() {
		return true
	}
	return false
}

/**
 *  @Description: 被掠夺信息
 *  @param user
 *  @param ack
 */
func (this *MiningManager) GetRobInfo(user *objs.User, ack *pb.MiningDrawLoadAck) error {
	this.MiningMu.Lock()
	defer this.MiningMu.Unlock()
	miningInfo := this.MiningMap[user.Id]
	if miningInfo == nil {
		return gamedb.ERRMININGOK
	}
	if !miningInfo.Rtime.IsZero() {
		ack.RobId = int64(miningInfo.Ruid)
		ack.RobName = this.GetUserManager().GetUserBasicInfo(miningInfo.Ruid).NickName
		ack.RobTime = miningInfo.Rtime.Unix()
		ack.RId = int64(miningInfo.Id)
	}
	ack.Status = this.GetDrawStatus(user)
	return nil
}

/**
 *  @Description: 被抢夺列表
 *  @param user
 *  @param ack
 */
func (this *MiningManager) RobList(user *objs.User, ack *pb.MiningRobListAck) {
	userId := user.Id
	//userInfo := this.GetUserManager().GetUserBasicInfo(userId)
	pbMineRob := make(map[int64]*pb.MiningRob)
	robList, err := modelGame.GetMiningModel().GetRobListByUserId(userId)
	if err != nil && err != sql.ErrNoRows {
		logger.Debug("RobList getList err:%v", err)
	}
	for _, mining := range robList {
		if mining.Rtime.IsZero() {
			continue
		}
		robUserInfo := this.GetUserManager().GetUserBasicInfo(mining.Ruid)
		pbMineRob[int64(mining.Id)] = &pb.MiningRob{
			Name:    robUserInfo.NickName,
			Combat:  int64(robUserInfo.Combat),
			Miner:   int32(mining.Miner),
			RobTime: mining.Rtime.Unix(),
			Id:      int64(mining.Id),
		}
	}
	ack.MineRob = pbMineRob
}

func (this *MiningManager) MakeMiningDbDate(UserId int, miner int, nowTime time.Time, addTime int) *modelGame.MiningDb {
	return &modelGame.MiningDb{
		UserId:     UserId,
		Miner:      miner,
		WorkTime:   nowTime,
		ExpireTime: time.Unix(nowTime.Unix()+int64(addTime), 0),
		CreatedAt:  nowTime,
		IsRobot:    constMining.MINING_DATA_ROBOT_NO,
	}
}

func (this *MiningManager) AddMiningData(mining *modelGame.MiningDb, onlyAddDb bool) error {
	//mining.Id, _ = modelGame.GetMiningModel().GetMiningId()
	err := modelGame.GetMiningModel().Create(mining)
	if err != nil {
		logger.Error("AddMiningData create err:%v", err)
		return err
	}
	if !onlyAddDb {
		this.MiningMap[mining.UserId] = objs.NewMining(mining)
	}
	logger.Debug("AddMiningData MiningMap:%v", this.MiningMap)
	return nil
}

func (this *MiningManager) UpdateMiningDate(mining *modelGame.MiningDb, onlyUpdateDb bool) error {
	err := modelGame.GetMiningModel().Update(mining)
	if err != nil {
		logger.Error("UpdateMiningDate create err:%v", err)
		return err
	}
	if !onlyUpdateDb {
		this.MiningMap[mining.UserId] = objs.NewMining(mining)
	}
	logger.Debug("UpdateMiningDate MiningMap:%v", this.MiningMap)
	return nil
}

func (this *MiningManager) DelMiningData(mining *modelGame.MiningDb) error {
	err := modelGame.GetMiningModel().Update(mining)
	if err != nil {
		logger.Error("DelMiningData create err:%v", err)
		return err
	}
	delete(this.MiningMap, mining.UserId)
	logger.Debug("DelMiningData MiningMap:%v", this.MiningMap)
	return nil
}
