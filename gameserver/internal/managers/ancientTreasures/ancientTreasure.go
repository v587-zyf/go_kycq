package ancientTreasures

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
)

type AncientTreasure struct {
	util.DefaultModule
	managersI.IModule
}

func NewAncientTreasure(module managersI.IModule) *AncientTreasure {
	return &AncientTreasure{IModule: module}
}

//  Active
//  @Description://远古宝物 激活
//  @receiver this
//  @param id 远古宝物id
//
func (this *AncientTreasure) Active(user *objs.User, id int, ack *pb.AncientTreasuresActivateAck, op *ophelper.OpBagHelperDefault) error {
	if user.AncientTreasure[id] != nil {
		return gamedb.ERRREPEATACTIVE
	}

	treasureCfg := gamedb.GetTreasureTreasureCfg(id)
	if treasureCfg == nil {
		return gamedb.ERRPARAM
	}

	has, _ := this.GetBag().HasEnoughItems(user, treasureCfg.Item)
	if !has {
		return gamedb.ERRNOTENOUGHGOODS
	}

	err := this.GetBag().RemoveItemsInfos(user, op, treasureCfg.Item)
	if err != nil {
		return err
	}

	user.AncientTreasure[id] = &model.AncientTreasures{
		Types: treasureCfg.Type,
	}
	this.GetUserManager().UpdateCombat(user, -1)
	user.Dirty = true
	ack.TreasureId = int32(id)
	return nil
}

func (this *AncientTreasure) upCheck(user *objs.User, id int) error {

	treasureCfg := gamedb.GetTreasureTreasureCfg(id)
	if treasureCfg == nil {
		return gamedb.ERRPARAM
	}

	if user.AncientTreasure[id] == nil {
		return gamedb.ERRNOTACTIVE
	}
	return nil
}

//
//  ZhuLin
//  @Description:
//  @receiver this 远古宝物 注灵
//  @param id 远古宝物id
//
func (this *AncientTreasure) ZhuLin(user *objs.User, id int, ack *pb.AncientTreasuresZhuLinAck, op *ophelper.OpBagHelperDefault) error {

	err := this.upCheck(user, id)
	if err != nil {
		return err
	}

	maxLv := gamedb.GetAncientTreasureZhuLinMaxLvById(id)
	if maxLv == user.AncientTreasure[id].ZhuLinLv {
		return gamedb.ERRLVENOUGH
	}

	zhuLinCfg := gamedb.GetAncientTreasureZhuLinLvById(id, user.AncientTreasure[id].ZhuLinLv)
	if zhuLinCfg == nil {
		return gamedb.ERRPARAM
	}

	has, _ := this.GetBag().HasEnoughItems(user, zhuLinCfg.Consume)
	if !has {
		return gamedb.ERRNOTENOUGHGOODS
	}

	err = this.GetBag().RemoveItemsInfos(user, op, zhuLinCfg.Consume)
	if err != nil {
		return err
	}

	user.AncientTreasure[id].ZhuLinLv++
	user.Dirty = true
	this.GetUserManager().UpdateCombat(user, -1)
	ack.TreasureId = int32(id)
	ack.ZhuLinLv = int32(user.AncientTreasure[id].ZhuLinLv)
	return nil
}

//
//  UpStar
//  @Description:
//  @receiver this  远古宝物 升星
//  @param id  远古宝物id
//
func (this *AncientTreasure) UpStar(user *objs.User, id int, ack *pb.AncientTreasuresUpStarAck, op *ophelper.OpBagHelperDefault) error {

	err := this.upCheck(user, id)
	if err != nil {
		return err
	}

	maxStar := gamedb.GetAncientTreasureStarMaxLvById(id)
	if maxStar == user.AncientTreasure[id].Star {
		return gamedb.ERRLVENOUGH
	}

	zhuLinCfg := gamedb.GetAncientTreasureStarById(id, user.AncientTreasure[id].Star)
	if zhuLinCfg == nil {
		return gamedb.ERRPARAM
	}

	has, _ := this.GetBag().HasEnoughItems(user, zhuLinCfg.Consume)
	if !has {
		return gamedb.ERRNOTENOUGHGOODS
	}

	err = this.GetBag().RemoveItemsInfos(user, op, zhuLinCfg.Consume)
	if err != nil {
		return err
	}
	user.AncientTreasure[id].Star++
	user.Dirty = true
	this.GetUserManager().UpdateCombat(user, -1)
	ack.TreasureId = int32(id)
	ack.StarLv = int32(user.AncientTreasure[id].Star)
	return nil
}

//
//  JueXin
//  @Description:
//  @receiver this   觉醒  只有满星才可以觉醒   觉醒后不可以重置 星数 和 等级
//  @param id  远古宝物id
//
func (this *AncientTreasure) JueXin(user *objs.User, treasureId, chooseIndex int, items []int32, ack *pb.AncientTreasuresJueXingAck, op *ophelper.OpBagHelperDefault) error {

	err := this.upCheck(user, treasureId)
	if err != nil {
		return err
	}

	if user.AncientTreasure[treasureId].JueXinLv >= 1 {
		return gamedb.ERRREPEATACTIVE
	}

	allItems, err := this.checkJueXinAllChooseItems(user, treasureId, items, chooseIndex)
	if err != nil {
		return err
	}

	has, _ := this.GetBag().HasEnoughItems(user, allItems)
	if !has {
		return gamedb.ERRNOTENOUGHGOODS
	}

	err = this.GetBag().RemoveItemsInfos(user, op, allItems)
	if err != nil {
		return err
	}

	user.AncientTreasure[treasureId].JueXinLv++
	user.Dirty = true
	this.GetUserManager().UpdateCombat(user, -1)
	ack.TreasureId = int32(treasureId)
	return nil
}

//觉醒 物品选择 后端校验
func (this *AncientTreasure) checkJueXinAllChooseItems(user *objs.User, treasureId int, items []int32, chooseIndex int) (gamedb.ItemInfos, error) {

	cfg := gamedb.GetAncientTreasureJueXinCfg(treasureId)
	if cfg == nil {
		return nil, gamedb.ERRPARAM
	}

	if chooseIndex < 0 || chooseIndex > len(cfg.Consume)-1 {
		return nil, gamedb.ERRPARAM
	}
	if len(items)%2 != 0 {
		logger.Error(" treasureId:%v items :%v", treasureId, items)
		return nil, gamedb.ERRPARAM
	}

	allItems := gamedb.ItemInfos{}
	chooseItemsQualityMap := make(map[int]int)
	for i, j := 0, len(items); i < j; i += 2 {
		itemId := int(items[i])
		count := int(items[i+1])
		base := gamedb.GetItemBaseCfg(itemId)
		if base == nil {
			continue
		}
		chooseItemsQualityMap[base.Quality] += count
		allItems = append(allItems, &gamedb.ItemInfo{ItemId: itemId, Count: count})
	}

	chooseItems := cfg.Consume[chooseIndex]
	chooseItemsMap := make(map[int]int)
	for i, j := 0, len(chooseItems); i < j; i += 2 {
		chooseItemsMap[chooseItems[i]] += chooseItems[i+1]
	}

	for quality, count := range chooseItemsMap {
		if chooseItemsQualityMap[quality] < count {
			logger.Error("远古宝物觉醒 激活消耗碎片 选择错误 userId:%v treasureId :%v, chooseIndex :%v, items:%v  chooseItemsQualityMap:%v  chooseItemsMap:%v", user.Id, treasureId, chooseIndex, items, chooseItemsQualityMap, chooseItemsMap)
			return nil, gamedb.ERRPARAM
		}
	}
	return allItems, nil
}

//
//  Reset
//  @Description: 重置 注灵等级 和 星数  激活图鉴的不可重置
//  @receiver this
//  @param id  远古宝物id
//
func (this *AncientTreasure) Reset(user *objs.User, id int, ack *pb.AncientTreasuresResertAck, op *ophelper.OpBagHelperDefault) error {

	err := this.upCheck(user, id)
	if err != nil {
		return err
	}

	has, _ := this.GetBag().HasEnoughItems(user, gamedb.GetConf().TreasureReset)
	if !has {
		return gamedb.ERRNOTENOUGHGOODS
	}
	this.GetBag().RemoveItemsInfos(user, op, gamedb.GetConf().TreasureReset)

	allReturnItems := gamedb.ItemInfos{}
	allReturnMap := make(map[int]int)
	//重置注灵
	nowLv := user.AncientTreasure[id].ZhuLinLv
	if nowLv > 0 {
		for i := nowLv - 1; i >= 0; i-- {
			cfg := gamedb.GetAncientTreasureZhuLinLvById(id, i)
			if cfg == nil {
				continue
			}
			for _, v := range cfg.Consume {
				allReturnMap[v.ItemId] += v.Count
			}
		}
	}
	//重置升星
	nowStar := user.AncientTreasure[id].Star
	if nowStar > 0 {
		for i := nowStar - 1; i >= 0; i-- {
			cfg := gamedb.GetAncientTreasureStarById(id, i)
			if cfg == nil {
				continue
			}
			for _, v := range cfg.Consume {
				allReturnMap[v.ItemId] += v.Count
			}
		}
	}

	//重置激活
	baseCfg := gamedb.GetTreasureTreasureCfg(id)
	if baseCfg.Item != nil {
		for _, v := range baseCfg.Item {
			allReturnMap[v.ItemId] += v.Count
		}
	}

	delete(user.AncientTreasure, id)

	for itemId, count := range allReturnMap {
		allReturnItems = append(allReturnItems, &gamedb.ItemInfo{ItemId: itemId, Count: count})
	}

	logger.Error("userId:%v  远古宝物重置 id:%v   重置可获得道具:%v", user.Id, id, allReturnMap)
	this.GetBag().AddItems(user, allReturnItems, op)
	this.GetUserManager().UpdateCombat(user, -1)
	ack.TreasureId = int32(id)
	return nil
}

// 获取远古宝物  达成的  K:套装id  v:达成的特殊加成
func (this *AncientTreasure) GetAncientTreasureTaoZ(user *objs.User) map[int]int {
	jueXinMap := make(map[int]int) //k:套装id v:可以加成的效果

	for id, data := range gamedb.GetAllAncientTreasureSuit() {
		allThreeStar := make([]int, 0)
		allJueXin := make([]int, 0)
		needContinue := false
		for _, tid := range data.TruesureId {
			if user.AncientTreasure[tid] == nil {
				needContinue = true
				break
			}
			if user.AncientTreasure[tid].Star >= 3 {
				allThreeStar = append(allThreeStar, tid)
			}

			if user.AncientTreasure[tid].JueXinLv >= 1 {
				allJueXin = append(allJueXin, tid)
			}

		}
		if needContinue {
			continue
		}
		//全都激活 加成
		jueXinMap[id] = 1

		//全部达成3星
		if len(allThreeStar) >= len(data.TruesureId) {
			jueXinMap[id] = 2
		}

		//全都激活觉醒效果
		if len(allJueXin) >= len(data.TruesureId) {
			jueXinMap[id] = 3
		}
	}
	return jueXinMap
}

func (this *AncientTreasure) GetConditionProp(user *objs.User, types int, props map[int]int, datas gamedb.IntSlice2) map[int]int {

	if types == 2 {
		//类型2 -- 特殊条件，增加属性   condition类型,条件数量 |加成属性id,属性数量,加成属性id,属性数量
		if len(datas) > 0 {
			if len(datas[0]) >= 3 {
				num, _ := this.GetCondition().Check(user, -1, datas[0][0], datas[0][1])
				multiple := num / datas[0][1]
				if multiple <= 0 {
					return props
				}
				if multiple > datas[0][2] {
					multiple = datas[0][2]
				}
				for index, data := range datas {
					if index == 0 {
						continue
					}
					for i, j := 0, len(data); i < j; i += 2 {
						props[data[i]] += data[i+1] * multiple
					}
				}
			}
		}
	}

	return props
}

func (this *AncientTreasure) BuildAncientTreasureInfo(user *objs.User) map[int32]*pb.AncientTreasureInfo {

	data := make(map[int32]*pb.AncientTreasureInfo)

	for tid, info := range user.AncientTreasure {
		data[int32(tid)] = &pb.AncientTreasureInfo{
			ZhuLinLv: int32(info.ZhuLinLv),
			StarLv:   int32(info.Star),
			JueXinLv: int32(info.JueXinLv),
			Types:    int32(info.Types),
		}

	}
	return data
}

func (this *AncientTreasure) GetAncientTreasureConditionValue(user *objs.User, ack *pb.AncientTreasuresCondotionInfosAck) {

	data := make(map[int32]int32)
	treasureType := gamedb.GetConf().TreasureType
	for _, types := range treasureType {
		num, _ := this.GetCondition().Check(user, -1, types, 10)
		data[int32(types)] = int32(num)
	}
	ack.AncientTreasureConditionInfos = data
	this.GetUserManager().UpdateCombat(user, -1)
	return
}
