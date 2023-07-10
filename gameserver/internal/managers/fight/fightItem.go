package fight

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gamelibs/publicCon/constFight"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/common"
	"cqserver/golibs/logger"
	"cqserver/golibs/nw"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"math"
	"time"
)

func (this *Fight) FsToGsAddItem(msg *pbserver.FSAddItemReq) (nw.ProtoMessage, error) {

	var err error
	var callErr error
	userId := int(msg.UserId)
	this.DispatchEvent(userId, msg, func(userId int, user *objs.User, data interface{}) {

		if user == nil {
			callErr = gamedb.ERRUNFOUNDUSER
			return
		}

		//添加道具
		op := ophelper.NewOpBagHelperDefault(int(msg.OpType))
		op.SetOpTypeSecond(user.FightStageId)
		dropItems := make(map[int]int)
		for k, v := range msg.Items {
			callErr = this.GetBag().Add(user, op, int(k), int(v))
			if callErr != nil {
				return
			}
			dropItems[int(k)] = int(v)
		}
		this.GetFirstDrop().CheckIsFirstDrop(user, dropItems)
		//推送添加道具
		this.GetUserManager().SendItemChangeNtf(user, op)
		//推送道具获取
		Ntf := &pb.FightItemsAddNtf{StageId: int32(user.FightStageId), Goods: op.ToChangeItems(), AddSource: 0}
		if op.GetOpType() == constBag.OpTypeMagicLayerTimeAward {
			Ntf.AddSource = 1
		}
		this.GetUserManager().SendMessage(user, Ntf, false)

	})
	if err != nil {
		return nil, err
	}
	if callErr != nil {
		return nil, callErr
	}

	ack := &pbserver.FSAddItemAck{
		IsSuccess: true,
	}
	return ack, nil
}

func (this *Fight) GsToFsPickUp(user *objs.User, objsIds []int32) (*pb.FightPickUpAck, *ophelper.OpBagHelperDefault, error) {

	ack := &pb.FightPickUpAck{
		Items: make(map[int32]*pb.ItemUnit),
	}

	stageConf := gamedb.GetStageStageCfg(user.FightStageId)
	if user.FightId <= 0 || stageConf == nil {
		logger.Error("玩家拾取物品，玩家不在战斗中 ,玩家：%v,战斗id：%v,战斗关卡：%v", user.IdName(), user.FightId, user.FightStageId)
		return nil, nil, gamedb.ERRFIGHTID
	}

	if len(objsIds) == 0 && !this.checkOneKeyPick(user) {
		logger.Error("玩家拾取物品，玩家申请一键拾取，权限不足：%v", user.IdName())
		return nil, nil, gamedb.ERRNOPOWER
	}

	if len(objsIds) == 0 {
		ack.IsOneKey = true
	}

	request := &pbserver.GsTOFsPickUpReq{
		UseId:  int32(user.Id),
		ObjIds: objsIds,
		IsPick: false,
	}
	reply := &pbserver.FsTOGsPickUpAck{}

	mapTypeConf := gamedb.GetMaptypeGameCfg(stageConf.Type)
	if mapTypeConf.BagFull == 1 {

		//预拾取
		err := this.FSRpcCall(user.FightId, user.FightStageId, request, reply)
		if err != nil {
			logger.Error("客户端请求拾取失败,玩家：%v,当前战斗：%v,拾取：%v，err:%v", user.IdName(), user.FightId, objsIds, err)
			return nil, nil, err
		}
		tempItems := make(gamedb.ItemInfos, len(reply.Items))
		k := 0
		for _, v := range reply.Items {
			tempItems[k] = &gamedb.ItemInfo{ItemId: int(v.ItemId), Count: int(v.ItemNum)}
			k++
		}

		if !this.GetBag().CheckHasEnoughPos(user, tempItems) {
			return nil, nil, gamedb.ERRBAGENOUGH
		}
	}

	//真实拾取
	request.IsPick = true
	reply = &pbserver.FsTOGsPickUpAck{}
	err1 := this.FSRpcCall(user.FightId, user.FightStageId, request, reply)
	if err1 != nil {
		logger.Error("客户端请求拾取失败,玩家：%v,当前战斗：%v,拾取：%v，err:%v", user.IdName(), user.FightId, objsIds, err1)
		return nil, nil, err1
	}

	op := ophelper.NewOpBagHelperDefault(constBag.OpTypePickUp)
	op.SetOpTypeSecond(user.FightStageId)
	hasRedPacketChange := false
	items := make(gamedb.ItemInfos, len(reply.Items))
	itemsCount := 0
	itemIds := make(map[int]int)
	for k, v := range reply.Items {
		ack.Items[k] = &pb.ItemUnit{ItemId: 0, Count: 0}
		ack.Items[k].ItemId = v.ItemId
		ack.Items[k].Count = int64(v.ItemNum)
		items[itemsCount] = &gamedb.ItemInfo{
			ItemId: int(v.ItemId),
			Count:  int(v.ItemNum),
		}
		itemsCount++
		itemIds[int(v.ItemId)] += 1
		//特殊处理  stage  传的是 monsterId
		this.GetAnnouncement().FightSendSystemChat(user, map[int]int{int(v.ItemId): int(v.ItemNum)}, int(v.MonsterId), pb.SCROLINGTYPE_KILL_MONSTER)
		itemConf := gamedb.GetItemBaseCfg(int(v.ItemId))
		if itemConf.Type == pb.ITEMTYPE_RED_PACKET {
			hasRedPacketChange = true
			user.RedPacketItem.PickNum += itemConf.EffectVal
			monsterConf := gamedb.GetMonsterMonsterCfg(int(v.MonsterId))
			if monsterConf != nil && len(monsterConf.DropSpecial) > 0 {
				user.RedPacketItem.PickInfo[monsterConf.DropSpecial[0]] += 1
			}
		}
	}
	isInmail := false
	if len(items) > 0 {
		isInmail = this.GetBag().AddItems(user, items, op)
	}

	this.GetFirstDrop().CheckIsFirstDrop(user, itemIds)
	ack.InMail = isInmail

	//战斗次数扣除，拾取物品时扣除
	this.fightTimeLess(user, stageConf.Type)

	if stageConf.Type == constFight.FIGHT_TYPE_DABAO {
		this.GetDaBao().SendSystemDropItem(user, user.FightStageId, reply.Items)
	}
	if hasRedPacketChange {
		this.FSSendMessage(user.FightId, user.FightStageId, &pbserver.GsToFsPickRedPacketInfo{
			UserId: int32(user.Id),
			RedPacket: &pbserver.ActorRedPacket{
				PickNum:   int32(user.RedPacketItem.PickNum),
				PickMax:   int32(gamedb.GetRedPacketDropMax(this.GetSystem().GetServerOpenDaysByServerId(user.ServerId))),
				PickInfos: common.ConvertMapIntToInt32(user.RedPacketItem.PickInfo),
			},
		})
		this.GetUserManager().SendMessage(user, &pb.UserRedPacketGetNumNtf{int32(user.RedPacketItem.PickNum)}, true)
	}

	return ack, op, nil
}

/**
 *  @Description: 战斗次数特殊扣除
 *  @param user
 *  @param stageType
 */
func (this *Fight) fightTimeLess(user *objs.User, stageType int) {
	if user.FightLessTimes {
		return
	}
	user.FightLessTimes = true
	stageId := user.FightStageId
	switch stageType {

	case constFight.FIGHT_TYPE_DARKPALACE_BOSS:
		user.DarkPalace.DareNum++
		this.GetUserManager().SendMessage(user, &pb.DarkPalaceDareNumNtf{
			DareNum: int32(user.DarkPalace.DareNum),
		}, true)
	case constFight.FIGHT_TYPE_PERSON_BOSS:
		hasFightNum := this.GetPersonBoss().GetBossKillNum(user.Id, stageId)
		this.GetPersonBoss().KillMonsterChangeDareNum(user, stageId, hasFightNum)
	case constFight.FIGHT_TYPE_VIPBOSS:
		this.GetVipBoss().KillMonsterChangeDareNum(user, stageId)
	case constFight.FIGHT_TYPE_TOWERBOSS:
		this.GetTower().KillMonsterChangeDareNum(user)
	case constFight.FIGHT_TYPE_STAGE_BOSS:
		this.GetStageManager().StageBossKill(user)
	case constFight.FIGHT_TYPE_HELL_BOSS:
		user.HellBoss.DareNum++
		this.GetUserManager().SendMessage(user, &pb.HellBossDareNumNtf{
			DareNum: int32(user.HellBoss.DareNum),
		}, true)
	}
	notWriteType := map[int]int{constFight.FIGHT_TYPE_STAGE: 0, constFight.FIGHT_TYPE_TOWERBOSS: 0, constFight.FIGHT_TYPE_STAGE_BOSS: 0}
	if _, ok := notWriteType[stageType]; !ok {
		this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_KILL_STAGE, []int{stageId, 1})
	}
	this.GetKillMonster().WriteKillMonster(user, stageId)
	user.Dirty = true
}

func (this *Fight) checkOneKeyPick(user *objs.User) bool {

	if this.GetVipManager().GetPrivilege(user, pb.VIPPRIVILEGE_QUICK_PICK) > 0 {
		return true
	}
	return false
}

/**
 *  @Description:	战斗服发来的扣除道具
 *  @param msg
 *  @return nw.ProtoMessage
 *  @return error
 */
func (this *Fight) UseItem(user *objs.User, itemId int) error {

	if itemId == pb.ITEMID_BACK_CITY {
		mainCityFightId := this.getResidentFightId(constFight.FIGHT_TYPE_MAIN_CITY_STAGE)
		if user.FightId > 0 {
			if user.FightId != int(mainCityFightId) {
				this.EnterPublicCopy(user, constFight.FIGHT_TYPE_MAIN_CITY_STAGE, false)
			} else {
				this.FSSendMessage(user.FightId, user.FightStageId, &pbserver.FsRandomDeliveryNtf{
					UserId: int32(user.Id),
					Rand:   false,
				})
			}
		} else {
			this.EnterPublicCopy(user, constFight.FIGHT_TYPE_MAIN_CITY_STAGE, false)
		}
	} else if itemId == pb.ITEMID_RANDOM_STONE {
		if user.FightId > 0 {
			this.FSSendMessage(user.FightId, user.FightStageId, &pbserver.FsRandomDeliveryNtf{
				UserId: int32(user.Id),
				Rand:   true,
			})
		}
	}
	itemConf := gamedb.GetItemBaseCfg(itemId)
	if itemConf.Type == pb.ITEMTYPE_POTION || itemConf.Type == pb.ITEMTYPE_HP_RECOVER || itemConf.Type == pb.ITEMTYPE_MP_RECOVER {
		if user.FightId > 0 {
			this.FSSendMessage(user.FightId, user.FightStageId, &pbserver.GsToFsUseItemNtf{
				UserId: int32(user.Id),
				ItemId: int32(itemId),
			})
		}
	}
	return nil
}

/**
 *  @Description: 鼓舞
 *  @param user
 *  @param op
 *  @return error
 */
func (this *Fight) CheerReq(user *objs.User, op *ophelper.OpBagHelperDefault) (int, int, error) {

	if user.FightId <= 0 {
		return 0, 0, gamedb.ERRFIGHTID
	}
	request := &pbserver.GsToFsGetCheerNumReq{
		UserId: int32(user.Id),
	}
	reply := &pbserver.FsToGsGetCheerNumAck{}

	err := this.FSRpcCall(user.FightId, user.FightStageId, request, reply)
	if err != nil {
		return 0, 0, err
	}

	if reply.CheerNum < 0 {
		return 0, 0, gamedb.ERRFIGHTID
	}
	stageConf := gamedb.GetStageStageCfg(user.FightStageId)
	costItemId, costNum, cheerLimit, cheerTotalLimit, err := this.getCheerSettingLimit(stageConf.Type)
	if err != nil {
		return 0, 0, err
	}
	if int(reply.CheerNum) >= cheerLimit {
		return int(reply.CheerNum), int(reply.GuildCheerNum), gamedb.ERRCHEERNUM
	}
	if int(reply.GuildCheerNum) >= cheerTotalLimit {
		return int(reply.GuildCheerNum), int(reply.GuildCheerNum), gamedb.ERRCHEERNUM
	}
	if costItemId > 0 && costNum > 0 {
		err = this.GetBag().Remove(user, op, costItemId, costNum)
		if err != nil {
			return int(reply.CheerNum), int(reply.GuildCheerNum), err
		}
	}

	cheerReq := &pbserver.GsToFsCheerReq{
		UserId: int32(user.Id),
	}
	err = this.FSSendMessage(user.FightId, user.FightStageId, cheerReq)
	if err != nil {
		return 0, 0, err
	}
	return int(reply.CheerNum + 1), int(reply.GuildCheerNum + 1), nil
}

func (this *Fight) getCheerSettingLimit(stageType int) (int, int, int, int, error) {

	costItemId := 0
	costNum := 0
	cherrNum := 0
	cheerTotalNum := math.MinInt32
	if stageType == constFight.FIGHT_TYPE_SHABAKE || stageType == constFight.FIGHT_TYPE_SHABAKE_NEW {
		costItemId, costNum, cherrNum = gamedb.GetConf().ShabakeBuff[0], gamedb.GetConf().ShabakeBuff[1], gamedb.GetConf().ShabakeBuff[2]
		if len(gamedb.GetConf().ShabakeBuff) > 4 {
			cheerTotalNum = gamedb.GetConf().ShabakeBuff[4]
		}
	} else if stageType == constFight.FIGHT_TYPE_CROSS_WORLD_LEADER {

		costItemId, costNum, cherrNum = gamedb.GetConf().WorldLeaderBuff[0], gamedb.GetConf().WorldLeaderBuff[1], gamedb.GetConf().WorldLeaderBuff[2]
		if len(gamedb.GetConf().WorldLeaderBuff) > 4 {
			cheerTotalNum = gamedb.GetConf().WorldLeaderBuff[4]
		}

	} else if stageType == constFight.FIGHT_TYPE_CROSS_SHABAKE {
		costItemId, costNum, cherrNum = gamedb.GetConf().KuafushabakeBuff[0], gamedb.GetConf().KuafushabakeBuff[1], gamedb.GetConf().KuafushabakeBuff[2]
		if len(gamedb.GetConf().KuafushabakeBuff) > 4 {
			cheerTotalNum = gamedb.GetConf().KuafushabakeBuff[4]
		}
	} else if stageType == constFight.FIGHT_TYPE_GUARDPILLAR {
		costItemId, costNum, cherrNum = gamedb.GetConf().GuardBuff[0], gamedb.GetConf().GuardBuff[1], gamedb.GetConf().GuardBuff[2]
		if len(gamedb.GetConf().GuardBuff) > 4 {
			cheerTotalNum = gamedb.GetConf().GuardBuff[4]
		}
	} else {
		return 0, 0, 0, 0, gamedb.ERRPARAM
	}
	return costItemId, costNum, cherrNum, cheerTotalNum, nil
}

/**
 *  @Description: 获取鼓舞使用次数
 *  @param user
 *  @return int
 *  @return error
 */
func (this *Fight) CheerGetUseNum(user *objs.User) (int, int, error) {
	if user.FightId <= 0 {
		return 0, 0, gamedb.ERRFIGHTID
	}
	request := &pbserver.GsToFsGetCheerNumReq{
		UserId: int32(user.Id),
	}
	reply := &pbserver.FsToGsGetCheerNumAck{}

	err := this.FSRpcCall(user.FightId, user.FightStageId, request, reply)
	if err != nil {
		return 0, 0, err
	}

	return int(reply.CheerNum), int(reply.GuildCheerNum), nil
}

/**
 *  @Description: 使用药水
 *  @param user
 *  @param op
 *  @return error
 */
func (this *Fight) UsePotion(user *objs.User, op *ophelper.OpBagHelperDefault) (int, error) {

	if user.FightId <= 0 {
		return 0, gamedb.ERRFIGHTID
	}
	request := &pbserver.GsToFsGetPotionCdReq{
		UserId: int32(user.Id),
	}
	reply := &pbserver.FsToGsGetPotionCdAck{}

	err := this.FSRpcCall(user.FightId, user.FightStageId, request, reply)
	if err != nil {
		return 0, err
	}

	if reply.UseTime < 0 {
		return 0, gamedb.ERRFIGHTID
	}

	now := time.Now().Unix()
	stageConf := gamedb.GetStageStageCfg(user.FightStageId)
	costItemId := 0
	costNum := 0
	cooldown := 0
	if stageConf.Type == constFight.FIGHT_TYPE_SHABAKE ||
		stageConf.Type == constFight.FIGHT_TYPE_SHABAKE_NEW ||
		stageConf.Type == constFight.FIGHT_TYPE_CROSS_WORLD_LEADER {

		if now-reply.UseTime > int64(gamedb.GetConf().ShabakePotion[4]) {
			costItemId = gamedb.GetConf().ShabakePotion[0]
			costNum = gamedb.GetConf().ShabakePotion[1]
			cooldown = gamedb.GetConf().ShabakePotion[4]
		} else {
			return 0, gamedb.ERRPOTIONCOOLDOWN
		}
	} else if stageConf.Type == constFight.FIGHT_TYPE_CROSS_SHABAKE {
		if now-reply.UseTime > int64(gamedb.GetConf().KuafushabakePotion[4]) {
			costItemId = gamedb.GetConf().KuafushabakePotion[0]
			costNum = gamedb.GetConf().KuafushabakePotion[1]
			cooldown = gamedb.GetConf().KuafushabakePotion[4]
		} else {
			return 0, gamedb.ERRPOTIONCOOLDOWN
		}
	} else {
		return 0, gamedb.ERRFIGHTTYPE

	}
	if costItemId > 0 && costNum > 0 {
		err = this.GetBag().Remove(user, op, costItemId, costNum)
		if err != nil {
			return 0, err
		}
	}

	usePotionReq := &pbserver.GsToFsUsePotionReq{
		UserId: int32(user.Id),
	}
	usePotionAck := &pbserver.FsToGsUsePotionAck{}
	err = this.FSRpcCall(user.FightId, user.FightStageId, usePotionReq, usePotionAck)
	if err != nil {
		return 0, err
	}
	return cooldown, nil
}

/**
 *  @Description: 使用药水CD
 *  @param user
 *  @return int
 *  @return error
 */
func (this *Fight) UsePotionCdReq(user *objs.User) (int, error) {
	if user.FightId <= 0 {
		return 0, gamedb.ERRFIGHTID
	}
	request := &pbserver.GsToFsGetPotionCdReq{
		UserId: int32(user.Id),
	}
	reply := &pbserver.FsToGsGetPotionCdAck{}

	err := this.FSRpcCall(user.FightId, user.FightStageId, request, reply)
	if err != nil {
		return 0, err
	}

	if reply.UseTime < 0 {
		return 0, gamedb.ERRFIGHTID
	}
	stageConf := gamedb.GetStageStageCfg(user.FightStageId)
	now := time.Now().Unix()
	cooldown := 0
	if stageConf.Type == constFight.FIGHT_TYPE_SHABAKE || stageConf.Type == constFight.FIGHT_TYPE_SHABAKE_NEW || stageConf.Type == constFight.FIGHT_TYPE_CROSS_WORLD_LEADER {

		if now-reply.UseTime >= int64(gamedb.GetConf().ShabakePotion[4]) {
			cooldown = 0
		} else {
			cooldown = gamedb.GetConf().ShabakePotion[4] - int(now-reply.UseTime)
		}
	} else {
		return 0, gamedb.ERRFIGHTID
	}
	return cooldown, err
}

/**
 *  @Description: 申请采集
 *  @param user
 *  @param ack
 */
func (this *Fight) CollectionReq(user *objs.User, objId int, ack *pb.FightCollectionAck) error {

	requestMsg := &pbserver.GsToFsCollectionReq{
		UserId: int32(user.Id),
		ObjId:  int32(objId),
	}
	replayMsg := &pbserver.FsToGsCollectionAck{}
	err := this.FSRpcCall(user.FightId, user.FightStageId, requestMsg, replayMsg)
	if err != nil {
		return err
	}
	ack.ObjId = int32(objId)
	ack.StartTime = replayMsg.StartTime
	ack.EndTime = replayMsg.EndTime
	return nil
}

/**
 *  @Description: 申请取消采集
 *  @param user
 *  @param ack
 */
func (this *Fight) CollectionCancelReq(user *objs.User, objId int, ack *pb.FightCollectionCancelAck) error {

	requestMsg := &pbserver.GsToFsCollectionCancelReq{
		UserId: int32(user.Id),
		ObjId:  int32(objId),
	}
	replayMsg := &pbserver.FsToGsCollectionCancelAck{}
	err := this.FSRpcCall(user.FightId, user.FightStageId, requestMsg, replayMsg)
	if err != nil {
		return err
	}
	return nil
}

/**
 *  @Description: 采集结束
 *  @param overMsg
 */
func (this *Fight) CollectionOverNtf(overMsg *pbserver.FsToGsCollectionNtf) {

	userId := int(overMsg.UserId)
	this.DispatchEvent(userId, overMsg, func(userId int, user *objs.User, data interface{}) {

		if user == nil {
			logger.Error("战斗服发送来采集结束，玩家未找到：%v,data:%v", userId, data)
			return
		}
		msg := data.(*pbserver.FsToGsCollectionNtf)
		op := ophelper.NewOpBagHelperDefault(constBag.OpTypeCollection)
		op.SetOpTypeSecond(user.FightStageId)
		for k, v := range msg.Items {

			this.GetBag().Add(user, op, int(k), int(v))
		}
		userNtf := &pb.FightCollectionNtf{
			Goods: op.ToChangeItems(),
		}
		this.GetUserManager().SendItemChangeNtf(user, op)
		this.GetUserManager().SendMessage(user, userNtf, true)
	})

}
