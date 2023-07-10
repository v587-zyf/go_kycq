package holyBeast

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constBag"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"fmt"
)

func NewHolyBeastManager(module managersI.IModule) *HolyBeastManager {
	return &HolyBeastManager{IModule: module}
}

type HolyBeastManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *HolyBeastManager) Load(user *objs.User, ack *pb.HolyBeastLoadInfoAck) {
	ack.HolyBeastInfos = this.buildInfo(user)
	return
}

//
//  Activate
//  @Description:圣兽激活
//  @param user
//  @param heroIndex:第几个hero
//  @param types:圣兽类型
//  @return error
//
func (this *HolyBeastManager) Activate(user *objs.User, heroIndex, types int, ack *pb.HolyBeastActivateAck, op *ophelper.OpBagHelperDefault) error {

	if user.Heros[heroIndex] == nil {
		return gamedb.ERRHOLYBEAST
	}

	if user.Heros[heroIndex].HolyBeastInfos[types].Star >= 0 {

		return gamedb.ERRHOLYBEAST1
	}

	if user.Heros[heroIndex].HolyBeastInfos[types] == nil {
		logger.Error("HolyBeastManager Activate types:%v  heroIndex:%v", types, heroIndex)
		return gamedb.ERRPARAM
	}
	cfg := gamedb.GetHolyBeastByTypesAndStar(types, 0)
	if cfg == nil {
		logger.Error("配置错误 types:%v star:%v", types, 0)
		return gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg("圣兽类型" + fmt.Sprintf("%v", types))
	}

	for _, data := range cfg.Active {
		if has, _ := this.GetBag().HasEnough(user, data.ItemId, data.Count); !has {
			return gamedb.ERRNOTENOUGHGOODS
		}
	}
	_ = this.GetBag().RemoveItemsInfos(user, op, cfg.Active)
	user.Heros[heroIndex].HolyBeastInfos[types].Star = 0
	user.Dirty = true
	ack.HolyBeastInfos = this.buildInfo(user)
	ack.HolyPoint = int32(user.Heros[heroIndex].HolyAllPoint)
	this.GetUserManager().UpdateCombat(user, heroIndex)
	return nil
}

//
//  UpStar
//  @Description: 圣兽升星
//  @receiver this
//  @param user
//  @param heroIndex
//  @param types 圣兽类型
//
func (this *HolyBeastManager) UpStar(user *objs.User, heroIndex, types int, ack *pb.HolyBeastUpStarAck, op *ophelper.OpBagHelperDefault) error {

	if user.Heros[heroIndex] == nil {
		return gamedb.ERRHOLYBEAST
	}

	if user.Heros[heroIndex].HolyBeastInfos[types] == nil {
		logger.Error("HolyBeastManager Activate types:%v  heroIndex:%v", types, heroIndex)
		return gamedb.ERRPARAM
	}
	if user.Heros[heroIndex].HolyBeastInfos[types].Star < 0 {

		return gamedb.ERRHOLYBEAST2
	}
	maxStar := gamedb.GetHolyBeastByTypesMaxStar(types)
	if user.Heros[heroIndex].HolyBeastInfos[types].Star == maxStar {
		return gamedb.ERRHOLYBEAST3
	}

	cfg := gamedb.GetHolyBeastByTypesAndStar(types, user.Heros[heroIndex].HolyBeastInfos[types].Star+1)
	if cfg == nil {
		logger.Error("配置错误 types:%v star:%v", types, user.Heros[heroIndex].HolyBeastInfos[types].Star+1)
		return gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg("圣兽类型" + fmt.Sprintf("%v", types))
	}

	if cfg.Effect != nil || cfg.SelectProperties != nil {
		return gamedb.ERRHOLYBEAST4
	}

	for _, data := range cfg.Active {
		ok, _ := this.GetBag().HasEnough(user, data.ItemId, data.Count)
		if !ok {
			return gamedb.ERRNOTENOUGHGOODS
		}
	}
	//if user.Heros[heroIndex].HolyAllPoint < cfg.ActiveNum {
	//	logger.Error("user.Heros[heroIndex:%v].HolyAllPoint:%v  data.Count:%v", heroIndex, user.Heros[heroIndex].HolyAllPoint, cfg.ActiveNum)
	//	return gamedb.ERRNOTENOUGHGOODS
	//}

	_ = this.GetBag().RemoveItemsInfos(user, op, cfg.Active)
	kyEvent.ShenShouUp(user, heroIndex, types, user.Heros[heroIndex].HolyBeastInfos[types].Star, cfg.Star)
	user.Heros[heroIndex].HolyBeastInfos[types].Star = cfg.Star

	//圣兽升星任务
	this.GetTask().AddTaskProcess(user, pb.CONDITION_UPGRADE_SHEN_SHOU, -1)
	user.Dirty = true
	ack.HolyBeastInfos = this.buildInfo(user)
	ack.HolyPoint = int32(user.Heros[heroIndex].HolyAllPoint)
	this.GetUserManager().UpdateCombat(user, heroIndex)
	this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_HERO_HOLY_BEAST_UP_STAR, []int{})

	return nil
}

//
//  ChooseProp
//  @Description: 特殊星  特殊属性选择
//  @receiver this
//  @param user
//  @param heroIndex
//  @param types
//  @param chooseIndex 选择物品的下标
//  @param ack
//
func (this *HolyBeastManager) ChooseProp(user *objs.User, heroIndex, types, chooseIndex int, ack *pb.HolyBeastChoosePropAck, op *ophelper.OpBagHelperDefault) error {

	if user.Heros[heroIndex] == nil {
		return gamedb.ERRHOLYBEAST
	}

	if user.Heros[heroIndex].HolyBeastInfos[types] == nil {
		logger.Error("HolyBeastManager Activate types:%v  heroIndex:%v", types, heroIndex)
		return gamedb.ERRPARAM
	}
	if user.Heros[heroIndex].HolyBeastInfos[types].Star < 0 {

		return gamedb.ERRHOLYBEAST2
	}

	maxStar := gamedb.GetHolyBeastByTypesMaxStar(types)
	if user.Heros[heroIndex].HolyBeastInfos[types].Star == maxStar {
		return gamedb.ERRHOLYBEAST3
	}

	cfg := gamedb.GetHolyBeastByTypesAndStar(types, user.Heros[heroIndex].HolyBeastInfos[types].Star+1)
	if cfg == nil {
		logger.Error("配置错误 types:%v star:%v", types, user.Heros[heroIndex].HolyBeastInfos[types].Star+1)
		return gamedb.ERRSETTINGNOTFOUND.SprintfErrMsg("圣兽类型" + fmt.Sprintf("%v", types))
	}
	if cfg.Effect == nil && cfg.SelectProperties == nil {
		logger.Error("接口请求错误 types:%v star:%v", types, user.Heros[heroIndex].HolyBeastInfos[types].Star+1)
		return gamedb.ERRPARAM
	}
	chooseId := 0
	if cfg.Effect != nil {
		if len(cfg.Effect) > 0 {
			if chooseIndex < 0 || chooseIndex > len(cfg.Effect)-1 {
				logger.Error("chooseIndex:%v", chooseIndex)
				return gamedb.ERRPARAM
			}
			chooseId = cfg.Effect[chooseIndex]
		}
	}
	if cfg.SelectProperties != nil {
		if len(cfg.SelectProperties) > 0 {
			if chooseIndex < 0 || chooseIndex > len(cfg.SelectProperties)-1 {
				logger.Error("chooseIndex:%v", chooseIndex)
				return gamedb.ERRPARAM
			}
			chooseId = cfg.SelectProperties[chooseIndex][0]
		}
	}

	for _, data := range cfg.Active {
		ok, _ := this.GetBag().HasEnough(user, data.ItemId, data.Count)
		if !ok {
			return gamedb.ERRNOTENOUGHGOODS
		}
	}

	_ = this.GetBag().RemoveItemsInfos(user, op, cfg.Active)

	//圣兽升星任务
	this.GetTask().AddTaskProcess(user, pb.CONDITION_UPGRADE_SHEN_SHOU, -1)

	kyEvent.ShenShouSkillUp(user, heroIndex, types, chooseId, user.Heros[heroIndex].HolyBeastInfos[types].Star, cfg.Star)
	user.Heros[heroIndex].HolyBeastInfos[types].Star = cfg.Star
	user.Heros[heroIndex].HolyBeastInfos[types].ChooseProp[cfg.Star] = chooseIndex
	user.Dirty = true
	ack.HolyBeastInfos = this.buildInfo(user)
	this.GetUserManager().UpdateCombat(user, heroIndex)

	return nil
}

//
//  Rest
//  @Description:圣灵点重置
//  @param heroIndex
//  @param types
//
func (this *HolyBeastManager) Rest(user *objs.User, heroIndex, types int, ack *pb.HolyBeastRestAck, op *ophelper.OpBagHelperDefault) error {

	logger.Info("圣灵点重置 userId:%v  heroIndex:%v  types:%v", user.Id, heroIndex, types)
	if user.Heros[heroIndex] == nil {
		return gamedb.ERRHOLYBEAST
	}

	if user.Heros[heroIndex].HolyBeastInfos[types] == nil {
		logger.Error("HolyBeastManager Activate types:%v  heroIndex:%v", types, heroIndex)
		return gamedb.ERRPARAM
	}
	if user.Heros[heroIndex].HolyBeastInfos[types].Star < 0 {

		return gamedb.ERRHOLYBEAST2
	}

	cfg1 := gamedb.GetConf().ResetHolyBeast
	if has, _ := this.GetBag().HasEnough(user, cfg1[types-1].ItemId, cfg1[types-1].Count); !has {
		return gamedb.ERRNOTENOUGHGOODS
	}
	op1 := ophelper.NewOpBagHelperDefault(constBag.OpTypeHolyBeastRestCost)
	_ = this.GetBag().Remove(user, op1, cfg1[types-1].ItemId, cfg1[types-1].Count)

	returnItems := gamedb.GetReturnHolyPoints(types, user.Heros[heroIndex].HolyBeastInfos[types].Star)
	if returnItems == nil || len(returnItems) <= 0 {
		logger.Error("heroIndex:%v types:%v  star:%v  returnItems:%v", heroIndex, types, user.Heros[heroIndex].HolyBeastInfos[types].Star, returnItems)
		return gamedb.ERRHOLYBEAST5
	}

	//user.Heros[heroIndex].HolyAllPoint += addNum
	this.GetBag().AddItems(user, returnItems, op)
	user.Heros[heroIndex].HolyBeastInfos[types].Star = 0
	user.Heros[heroIndex].HolyBeastInfos[types].ChooseProp = make(model.IntKv)
	user.Dirty = true
	ack.HolyPoint = int32(user.Heros[heroIndex].HolyAllPoint)
	ack.HolyBeastInfos = this.buildInfo(user)
	ack.Goods = op1.ToChangeItems()
	this.GetUserManager().UpdateCombat(user, heroIndex)

	return nil
}

func (this *HolyBeastManager) buildInfo(user *objs.User) map[int32]*pb.HolyBeastInfos {

	data := make(map[int32]*pb.HolyBeastInfos)
	for heroIndex, info := range user.Heros {
		if data[int32(heroIndex)] == nil {
			data[int32(heroIndex)] = &pb.HolyBeastInfos{HolyBeastInfo: make([]*pb.HolyBeastInfo, 0)}
		}
		data[int32(heroIndex)].HeroIndex = int32(heroIndex)
		data[int32(heroIndex)].AllPonts = int32(info.HolyAllPoint)
		for key, info1 := range info.HolyBeastInfos {
			choosePro := make(map[int32]int32)
			for k, v := range info1.ChooseProp {
				choosePro[int32(k)] = int32(v)
			}
			data[int32(heroIndex)].HolyBeastInfo = append(data[int32(heroIndex)].HolyBeastInfo, &pb.HolyBeastInfo{Type: int32(key), Star: int32(info1.Star), ChooseProperty: choosePro})
		}
	}
	return data
}
