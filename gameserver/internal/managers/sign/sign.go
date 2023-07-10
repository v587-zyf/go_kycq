package sign

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gamelibs/model"
	"cqserver/gameserver/internal/builder"
	"cqserver/gameserver/internal/managersI"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"strconv"
)

func NewSignManager(module managersI.IModule) *SignManager {
	p := &SignManager{IModule: module}
	return p
}

type SignManager struct {
	util.DefaultModule
	managersI.IModule
}

func (this *SignManager) Online(user *objs.User) {
	if user.Sign == nil {
		user.Sign = &model.Sign{SignDay: make(model.IntKv), Cumulative: make(model.IntKv)}
	}
	this.ResetSign(user, false)
}

func (this *SignManager) ResetSign(user *objs.User, reset bool) {
	openDays := this.GetSystem().GetServerOpenDaysByServerId(user.ServerId)
	userSign := user.Sign
	if userSign.ContinuitySign/10000 != openDays-1 && userSign.ContinuitySign/10000 != openDays {
		userSign.ContinuitySign = 0
	}

	signCircle := openDays / 30
	if userSign.ResetTime != signCircle && openDays%30 != 0 {
		user.Sign = &model.Sign{
			ResetTime:      signCircle,
			SignDay:        make(model.IntKv),
			Cumulative:     make(model.IntKv),
			ContinuitySign: userSign.ContinuitySign,
		}
		if reset {
			this.GetUserManager().SendMessage(user, &pb.SignResetNtf{SignInfo: builder.BuildSignInfo(user)}, true)
		}
	}
}

/**
 *  @Description: 签到
 *  @param user
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *SignManager) Sign(user *objs.User, op *ophelper.OpBagHelperDefault, ack *pb.SignAck) error {
	userSign := user.Sign
	openDay := this.changeServerOpenDay(user)

	if _, ok := userSign.SignDay[openDay]; ok {
		return gamedb.ERRSIGNREPEAT
	}
	if err := this.UpdateSign(user, op, userSign, openDay); err != nil {
		return err
	}

	ack.SignInfo = builder.BuildSignInfo(user)
	return nil
}

/**
 *  @Description: 补签
 *  @param user
 *  @param repairDay	补签哪一天
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *SignManager) Repair(user *objs.User, repairDay int, op *ophelper.OpBagHelperDefault, ack *pb.SignRepairAck) error {
	userSign := user.Sign
	if _, ok := userSign.SignDay[repairDay]; ok {
		return gamedb.ERRSIGNREPEAT
	}
	openDay := this.changeServerOpenDay(user)
	if repairDay > openDay {
		return gamedb.ERRPARAM
	}

	reward := gamedb.GetConf().RepairSign
	if err := this.GetBag().Remove(user, op, reward.ItemId, reward.Count); err != nil {
		return err
	}
	if err := this.UpdateSign(user, op, userSign, repairDay); err != nil {
		return err
	}

	ack.SignInfo = builder.BuildSignInfo(user)
	return nil
}

/**
 *  @Description: 累计奖励
 *  @param user
 *  @param cumulativeDay	奖励天数
 *  @param op
 *  @param ack
 *  @return error
 */
func (this *SignManager) Cumulative(user *objs.User, cumulativeDay int, op *ophelper.OpBagHelperDefault, ack *pb.CumulativeSignAck) error {
	userSign := user.Sign

	if userSign.Count < cumulativeDay {
		return gamedb.ERRNOTENOUGHSIGN
	}
	if userSign.Cumulative[cumulativeDay] != 0 {
		return gamedb.ERRREPEATRECEIVE
	}

	cfg := gamedb.GetCumulativeByDay(cumulativeDay)
	if cfg == nil {
		logger.Error("cumulativeSign cfg error day is %v", cumulativeDay)
		return gamedb.ERRPARAM
	}
	if err := this.GetBag().Add(user, op, cfg.Reward.ItemId, cfg.Reward.Count); err != nil {
		return err
	}
	userSign.Cumulative[cumulativeDay] = cumulativeDay
	user.Dirty = true

	ack.SignInfo = builder.BuildSignInfo(user)
	return nil
}

// 更新签到数据
func (this *SignManager) UpdateSign(user *objs.User, op *ophelper.OpBagHelperDefault, userSign *model.Sign, day int) error {
	signCfg := gamedb.GetSignSignCfg(day)
	if signCfg == nil {
		logger.Error("sign cfg error day is %v", day)
		return gamedb.ERRPARAM
	}
	if err := this.GetBag().Add(user, op, signCfg.Reward.ItemId, signCfg.Reward.Count); err != nil {
		return err
	}

	userSign.Count++
	userSign.SignDay[day] = 0
	user.Dirty = true

	conditionData := this.GetCondition().GetConditionData(user, pb.CONDITION_CONTINUOUS_SIGN, 0)
	openDay := this.GetSystem().GetServerOpenDaysByServerId(user.ServerId)
	signDay := 0
	if conditionData != 0 {
		signDay, _ = strconv.Atoi(strconv.Itoa(conditionData)[4:])
	}
	userSign.ContinuitySign = openDay*10000 + signDay
	this.GetCondition().RecordCondition(user, pb.CONDITION_ALL_SIGN, []int{1})
	return nil
}

func (this *SignManager) changeServerOpenDay(user *objs.User) int {
	day := this.GetSystem().GetServerOpenDaysByServerId(user.ServerId)
	day %= 30
	if day == 0 {
		day = 30
	}
	return day
}
