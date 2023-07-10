package title

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"time"
)

func NewTitleManager(m managersI.IModule) *Title {
	return &Title{
		IModule: m,
	}
}

type Title struct {
	util.DefaultModule
	managersI.IModule
}

func (this *Title) Online(user *objs.User) {
	userTitle := user.Title
	for _, info := range userTitle {
		if info.EndTime != -1 {
			this.CheckExpire(user)
			break
		}
	}
}

/**
 *  @Description: 称号激活
 *  @param user
 *  @param titleId	称号id
 *  @param op
 *  @return error
 */
func (this *Title) Active(user *objs.User, titleId int, op *ophelper.OpBagHelperDefault) error {
	titleCfg := gamedb.GetTitleTitleCfg(titleId)
	if titleCfg == nil || len(titleCfg.Item) > 1 {
		return gamedb.ERRPARAM
	}

	userTitle := user.Title
	timeNow := int(time.Now().Unix())
	title, ok := userTitle[titleId]
	if ok {
		if title.EndTime > timeNow || title.EndTime == -1 {
			return gamedb.ERRREPEATACTIVE
		}
	}
	if err := this.GetBag().RemoveItemsInfos(user, op, titleCfg.Item); err != nil {
		return err
	}

	endTime := -1
	if titleCfg.Time != -1 {
		endTime = timeNow + titleCfg.Time
		user.CheckTitleExpire = true
	}
	userTitle[titleId] = &model.TitleUnit{
		StartTime: timeNow,
		EndTime:   endTime,
		IsLook:    true,
	}
	user.Dirty = true
	this.GetUserManager().UpdateCombat(user, -1)
	return nil
}

/**
 *  @Description: 称号自动激活
 *  @param user
 */
func (this *Title) AutoActive(user *objs.User) {
	//userTitle := user.Title
	//timeNow := int(time.Now().Unix())
	//titleCfgs := gamedb.GetAutoActiveTitleCfgs()
	//pbSlice := make([]*pb.Title, 0)
	//for id, cfg := range titleCfgs {
	//	if _, ok := userTitle[id]; ok {
	//		continue
	//	}
	//	conditionFlag := true
	//	for _, condition := range cfg.Condition {
	//		if _, check := this.GetCondition().CheckBySlice(user, -1, condition); !check {
	//			conditionFlag = false
	//			break
	//		}
	//	}
	//	if !conditionFlag {
	//		continue
	//	}
	//	endTime := -1
	//	if cfg.Time != -1 {
	//		endTime = timeNow + cfg.Time
	//		user.CheckTitleExpire = true
	//	}
	//	userTitle[id] = &model.TitleUnit{
	//		StartTime: timeNow,
	//		EndTime:   endTime,
	//	}
	//	pbSlice = append(pbSlice, builder.BuildTitle(id, userTitle[id]))
	//}
	//if len(pbSlice) > 0 {
	//	this.GetUserManager().SendMessage(user, &pb.TitleAutoActiveNtf{TitleList: pbSlice}, true)
	//	this.GetUserManager().UpdateCombat(user, -1)
	//}
}

/**
 *  @Description: 称号校验是否过期
 *  @param user
 */
func (this *Title) CheckExpire(user *objs.User) {
	userTitle := user.Title
	userHeros := user.Heros
	timeNow := int(time.Now().Unix())
	pbSlice := make([]*pb.Title, 0)
	isUpdateDisplay := false
	checkExpireNum := 0
	for id, info := range userTitle {
		if info.IsExpire {
			continue
		}
		if info.EndTime != -1 {
			if info.EndTime < timeNow {
				pbSlice = append(pbSlice, builder.BuildTitle(id, info))
				for heroIndex, hero := range userHeros {
					if hero.Wear.TitleId != id {
						continue
					}
					hero.Wear.TitleId = 0
					isUpdateDisplay = true
					this.GetUserManager().SendMessage(user, &pb.TitleRemoveAck{HeroIndex: int32(heroIndex), TitleId: int32(id)}, true)
				}
				info.IsExpire = true
			} else {
				checkExpireNum++
			}
		}
	}
	if checkExpireNum < 1 {
		user.CheckTitleExpire = false
	} else {
		user.CheckTitleExpire = true
	}
	if len(pbSlice) > 0 {
		this.GetUserManager().SendMessage(user, &pb.TitleExpireNtf{TitleList: pbSlice}, true)
		if isUpdateDisplay {
			this.GetUserManager().SendDisplay(user)
		}
		this.GetUserManager().UpdateCombat(user, -1)
	}
	user.Dirty = true
}

/**
 *  @Description: 称号穿戴
 *  @param user
 *  @param heroIndex
 *  @param titleId
 *  @return error
 */
func (this *Title) Wear(user *objs.User, heroIndex, titleId int) error {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND
	}

	userTitle := user.Title
	title, ok := userTitle[titleId]
	if !ok {
		return gamedb.ERRNOTACTIVE
	}
	if title.EndTime != -1 && title.EndTime < int(time.Now().Unix()) {
		return gamedb.ERREXPIRE
	}
	for heroIndex, hero := range user.Heros {
		if hero.Wear.TitleId == titleId {
			hero.Wear.TitleId = 0
			this.GetUserManager().SendMessage(user, &pb.TitleRemoveAck{HeroIndex: int32(heroIndex), TitleId: int32(titleId)}, true)
		}
	}
	hero.Wear.TitleId = titleId
	user.Dirty = true

	this.GetUserManager().SendDisplay(user)
	return nil
}

/**
 *  @Description: 称号卸下
 *  @param user
 *  @param heroIndex
 *  @return error
 */
func (this *Title) Remove(user *objs.User, heroIndex int) (error, int) {
	hero := user.Heros[heroIndex]
	if hero == nil {
		return gamedb.ERRHERONOTFOUND, 0
	}

	wearTitleId := hero.Wear.TitleId
	hero.Wear.TitleId = 0
	user.Dirty = true

	this.GetUserManager().SendDisplay(user)
	return nil, wearTitleId
}

/**
 *  @Description: 称号查看
 *  @param user
 *  @param titleId
 *  @return error
 */
func (this *Title) Look(user *objs.User, titleId int) error {
	userTitle := user.Title
	title, ok := userTitle[titleId]
	if !ok {
		return gamedb.ERRNOTACTIVE
	}
	title.IsLook = true
	user.Dirty = true
	return nil
}
