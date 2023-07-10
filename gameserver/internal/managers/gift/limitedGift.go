package gift

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gamelibs/publicCon/constConstant"
	"cqserver/gameserver/internal/kyEvent"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/protobuf/pb"
	"time"
)

const (
	LIMITEDGIFT_DEF_GRADE = 1
	LIMITEDGIFT_DEF_LV    = 1
)

func (this *GiftManager) LimitedMerge(user *objs.User) {
	userLimitGift := user.LimitedGift
	if userLimitGift.MergeData == nil || len(userLimitGift.MergeData) < 1 {
		return
	}
	list, tLv := userLimitGift.List, userLimitGift.TLv
	nowTime := int(time.Now().Unix())
	for t, grade := range userLimitGift.MergeData {
		if lv, ok := tLv[t]; ok {
			limitedGiftCfg := gamedb.GetLimitedGiftByTypeAndLv(t, lv)
			if limitedGiftCfg == nil {
				continue
			}
			list[t][lv] = &model.LimitGiftUnit{
				Lv:        lv,
				Grade:     grade,
				StartTime: nowTime,
				EndTime:   nowTime + limitedGiftCfg.Time,
			}
		}
	}
	userLimitGift.MergeData = make(model.IntKv)
}

func (this *GiftManager) SendLimitedGift(user *objs.User) *pb.LimitedGiftNtf {
	ntf := &pb.LimitedGiftNtf{List: make([]*pb.LimitedGiftInfo, 0)}
	nowTime := int(time.Now().Unix())

	userLimitGift := user.LimitedGift
	list, isBuy, tLv := userLimitGift.List, userLimitGift.IsBuy, userLimitGift.TLv
	for _, moduleType := range pb.LIMITEDGIFTTYPE_ARRAY {
		moduleLv := tLv[moduleType]
		if len(list[moduleType]) <= 0 {
			continue
		}
		if !isBuy[moduleType] && list[moduleType][moduleLv].EndTime > nowTime {
			ntf.List = append(ntf.List, this.LimitedGiftInfo(moduleType, moduleLv, list[moduleType][moduleLv]))
		} else {
			nextLv := moduleLv + 1
			if limitGiftUnit, ok := list[moduleType][nextLv]; ok {
				limitedGiftCfg := gamedb.GetLimitedGiftByTypeAndLv(moduleType, nextLv)
				gradeStatus, grade := this.CheckLimitGiftGrade(userLimitGift.GradeStatus[moduleType], userLimitGift.IsBuy[moduleType], list[moduleType][moduleLv].Grade)

				limitGiftUnit.StartTime = nowTime
				limitGiftUnit.EndTime = nowTime + limitedGiftCfg.Time
				limitGiftUnit.Grade = grade

				userLimitGift.GradeStatus[moduleType] = gradeStatus
				userLimitGift.TLv[moduleType] = nextLv
				userLimitGift.IsBuy[moduleType] = false
				ntf.List = append(ntf.List, this.LimitedGiftInfo(moduleType, nextLv, limitGiftUnit))
			}
		}
	}
	user.Dirty = true
	return ntf
}

/**
 *  @Description: 限时礼包根据模块发送
 *  @param user
 *  @param t	模块
 */
func (this *GiftManager) LimitedGiftSend(user *objs.User) {
	nowTime := int(time.Now().Unix())
	userLimitGift := user.LimitedGift
	list, tLv := userLimitGift.List, userLimitGift.TLv
	for _, moduleType := range pb.LIMITEDGIFTTYPE_ARRAY {
		if _, ok := list[moduleType]; !ok {
			list[moduleType] = make(map[int]*model.LimitGiftUnit)
		}
		lv, ok := tLv[moduleType]
		if !ok {
			lv = LIMITEDGIFT_DEF_LV
		} else {
			lv++
		}
		for i := lv; i < constConstant.COMPUTE_TEN_THOUSAND; i++ {
			limitedCfg := gamedb.GetLimitedGiftByTypeAndLv(moduleType, i)
			if limitedCfg == nil {
				break
			}
			if _, ok := list[moduleType][i]; ok {
				continue
			}
			flag := false
			for heroIndex := range user.Heros {
				if check := this.GetCondition().CheckMulti(user, heroIndex, limitedCfg.Condition); check {
					flag = true
					break
				}
			}
			if !flag {
				break
			}
			list[moduleType][i] = &model.LimitGiftUnit{Lv: i}
		}
		if !ok && (list[moduleType][lv] != nil && list[moduleType][lv].EndTime == 0) {
			limitedGiftCfg := gamedb.GetLimitedGiftByTypeAndLv(moduleType, lv)
			if limitedGiftCfg != nil {
				var gradeStatus, grade int
				gradeStatus, grade = pb.LIMITEDGIFTGRADESTATUS_KEEP, LIMITEDGIFT_DEF_GRADE
				list[moduleType][lv] = &model.LimitGiftUnit{
					Lv:        lv,
					Grade:     grade,
					StartTime: nowTime,
					EndTime:   nowTime + limitedGiftCfg.Time,
				}
				userLimitGift.TLv[moduleType] = lv
				userLimitGift.GradeStatus[moduleType] = gradeStatus
			}
		}
	}
	limitedGiftNtf := this.SendLimitedGift(user)
	if len(limitedGiftNtf.List) > 0 {
		this.GetUserManager().SendMessage(user, limitedGiftNtf, true)
	}
}

/**
 *  @Description: 限时礼包购买
 *  @param user
 *  @param op
 *  @return error
 */
func (this *GiftManager) LimitedGiftBuy(user *objs.User, moduleType int, op *ophelper.OpBagHelperDefault) error {
	if _, ok := pb.LIMITEDGIFTTYPE_MAP[moduleType]; !ok {
		return gamedb.ERRPARAM
	}
	userLimitGift, nowLimitedGift, reward, payNum, payT, confId, _, err := this.GetNowLimitedGiftInfo(user, moduleType)
	if err != nil {
		return err
	}
	switch payT {
	case pb.LIMITEDGIFTBUYTYPE_INGOT:
		if err := this.GetBag().Remove(user, op, pb.ITEMID_INGOT, payNum); err != nil {
			return err
		}
	case pb.LIMITEDGIFTBUYTYPE_MONEY:
		return gamedb.ERRLIMITEDGIFTBUYTYPE
	}
	this.GetBag().AddItems(user, reward, op)
	nowLimitedGift.IsBuy = true
	userLimitGift.IsBuy[moduleType] = true
	user.Dirty = true
	kyEvent.UserGiftBuy(user, pb.MONEYPAYTYPE_LIMITED_GIFT, confId, pb.ITEMID_INGOT, payNum, payT)
	return nil
}

/**
 *  @Description: 限时礼包，校验购买类型，支付金额
 *  @param user
 *  @param typeId	limitedGift表id
 *  @param payNum	支付金额
 */
func (this *GiftManager) LimitedGiftCheckBuy(user *objs.User, typeId, payNum int) (string, error) {
	limitedGiftCfg := gamedb.GetLimitedGiftLimitedGiftCfg(typeId)
	if limitedGiftCfg == nil {
		return "", gamedb.ERRPARAM
	}
	_, nowLimitedGift, _, payMoney, payT, _, name, err := this.GetNowLimitedGiftInfo(user, limitedGiftCfg.Type)
	if err != nil {
		return "", err
	}
	if payT == pb.LIMITEDGIFTBUYTYPE_INGOT {
		return "", gamedb.ERRLIMITEDGIFTBUYTYPE
	}
	if payNum != payMoney {
		return "", gamedb.ERRBUYNUM
	}
	if nowLimitedGift.IsBuy {
		return "", gamedb.ERRREPEATBUY
	}
	return name, nil
}

/**
 *  @Description: 限时礼包支付后操作
 *  @param user
 *  @param typeId	limitedGift表id
 *  @param op
 */
func (this *GiftManager) LimitedGiftBuyOperation(user *objs.User, typeId int, op *ophelper.OpBagHelperDefault) {
	limitedGiftCfg := gamedb.GetLimitedGiftLimitedGiftCfg(typeId)
	if limitedGiftCfg == nil {
		logger.Error("限时礼包配置未找到 id:%v", typeId)
		return
	}

	userLimitGift, nowLimitedGift, reward, payNum, payT, _, _, _ := this.GetNowLimitedGiftInfo(user, limitedGiftCfg.Type)
	if nowLimitedGift.IsBuy && userLimitGift.IsBuy[limitedGiftCfg.Type] {
		return
	}
	this.GetBag().AddItems(user, reward, op)
	nowLimitedGift.IsBuy = true
	userLimitGift.IsBuy[limitedGiftCfg.Type] = true
	user.Dirty = true
	kyEvent.UserGiftBuy(user, pb.MONEYPAYTYPE_LIMITED_GIFT, typeId, 0, payNum, payT)
	this.GetUserManager().SendMessage(user, &pb.LimitedGiftBuyAck{Goods: op.ToChangeItems(), Type: int32(limitedGiftCfg.Type)}, true)
}

/**
 *  @Description: 获取限时礼包信息
 *  @param user
 *  @param moduleType		模块
 *  @return *model.LimitGift		用户限时礼包
 *  @return *model.LimitGiftUnit	限时礼包具体模块信息
 *  @return gamedb.ItemInfos		模块当前档次奖励信息
 *  @return int						支付金额
 *  @return int						支付类型
 *  @return error
 */
func (this *GiftManager) GetNowLimitedGiftInfo(user *objs.User, moduleType int) (*model.LimitGift, *model.LimitGiftUnit, gamedb.ItemInfos, int, int, int, string, error) {
	userLimitGift := user.LimitedGift
	nowTime := int(time.Now().Unix())

	list, tLv := userLimitGift.List, userLimitGift.TLv
	nowTlv := tLv[moduleType]
	nowLimitedGift := list[moduleType][nowTlv]
	if nowLimitedGift == nil || nowTlv == 0 {
		return nil, nil, nil, 0, 0, 0, "", gamedb.ERRPARAM
	}
	if nowLimitedGift.IsBuy || nowLimitedGift.EndTime < nowTime {
		return nil, nil, nil, 0, 0, 0, "", gamedb.ERRLIMITEDGIFTBUY
	}
	limitedGiftCfg := gamedb.GetLimitedGiftByTypeAndLv(moduleType, nowTlv)
	var reward gamedb.ItemInfos
	var payNum, payT int
	var name string
	switch nowLimitedGift.Grade {
	case pb.LIMITEDGIFTGRADE_ONE:
		reward = limitedGiftCfg.Reward1
		payNum = limitedGiftCfg.Consume1
		payT = limitedGiftCfg.Type1
		name = limitedGiftCfg.Display1
	case pb.LIMITEDGIFTGRADE_TWO:
		reward = limitedGiftCfg.Reward2
		payNum = limitedGiftCfg.Consume2
		payT = limitedGiftCfg.Type2
		name = limitedGiftCfg.Display2
	case pb.LIMITEDGIFTGRADE_THREE:
		reward = limitedGiftCfg.Reward3
		payNum = limitedGiftCfg.Consume3
		payT = limitedGiftCfg.Type3
		name = limitedGiftCfg.Display3
	case pb.LIMITEDGIFTGRADE_FOUR:
		reward = limitedGiftCfg.Reward4
		payNum = limitedGiftCfg.Consume4
		payT = limitedGiftCfg.Type4
		name = limitedGiftCfg.Display4
	case pb.LIMITEDGIFTGRADE_FIVE:
		reward = limitedGiftCfg.Reward5
		payNum = limitedGiftCfg.Consume5
		payT = limitedGiftCfg.Type5
		name = limitedGiftCfg.Display5
	}
	return userLimitGift, nowLimitedGift, reward, payNum, payT, limitedGiftCfg.Id, name, nil
}

func (this *GiftManager) CheckLimitGiftGrade(gradeState int, isBuy bool, nowLimitedGiftGrade int) (int, int) {
	grade, gradeStatus := 0, 0
	switch gradeState {
	case pb.LIMITEDGIFTGRADESTATUS_UP:
		if isBuy {
			gradeStatus = pb.LIMITEDGIFTGRADESTATUS_UP
			if _, ok := pb.LIMITEDGIFTGRADE_MAP[nowLimitedGiftGrade+1]; ok {
				grade = nowLimitedGiftGrade + 1
			} else {
				grade = nowLimitedGiftGrade
			}
		} else {
			gradeStatus = pb.LIMITEDGIFTGRADESTATUS_DOWN
			if _, ok := pb.LIMITEDGIFTGRADE_MAP[nowLimitedGiftGrade-1]; ok {
				grade = nowLimitedGiftGrade - 1
			} else {
				grade = nowLimitedGiftGrade
			}
		}
	case pb.LIMITEDGIFTGRADESTATUS_DOWN:
		if isBuy {
			gradeStatus = pb.LIMITEDGIFTGRADESTATUS_KEEP
			grade = nowLimitedGiftGrade
		} else {
			gradeStatus = pb.LIMITEDGIFTGRADESTATUS_UP
			if _, ok := pb.LIMITEDGIFTGRADE_MAP[nowLimitedGiftGrade+1]; ok {
				grade = nowLimitedGiftGrade + 1
			} else {
				grade = nowLimitedGiftGrade
			}
		}
	default:
		if isBuy {
			gradeStatus = pb.LIMITEDGIFTGRADESTATUS_UP
			if _, ok := pb.LIMITEDGIFTGRADE_MAP[nowLimitedGiftGrade+1]; ok {
				grade = nowLimitedGiftGrade + 1
			} else {
				grade = nowLimitedGiftGrade
			}
		} else {
			gradeStatus = pb.LIMITEDGIFTGRADESTATUS_DOWN
			if _, ok := pb.LIMITEDGIFTGRADE_MAP[nowLimitedGiftGrade-1]; ok {
				grade = nowLimitedGiftGrade - 1
			} else {
				grade = nowLimitedGiftGrade
			}
		}
	}
	return gradeStatus, grade
}

func (this *GiftManager) LimitedGiftInfo(t, tlv int, giftUnit *model.LimitGiftUnit) *pb.LimitedGiftInfo {
	return &pb.LimitedGiftInfo{
		Type:      int32(t),
		Lv:        int32(tlv),
		Grade:     int32(giftUnit.Grade),
		StartTime: int64(giftUnit.StartTime),
		EndTime:   int64(giftUnit.EndTime),
	}
}

/**
 *  @Description: 限时礼包合服，把合服所用时间加到当前礼包的结束时间
 *  @param user
 *  @param t
 */
func (this *GiftManager) LimitedGiftMerge(user *objs.User, t int) {
	userLimitedGift := user.LimitedGift
	for grade, lv := range userLimitedGift.TLv {
		userLimitedGift.List[grade][lv].EndTime += t
	}
}
