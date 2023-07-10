package managersI

import (
	"cqserver/gamelibs/modelGame"
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pbserver"
)

type IRecharge interface {
	Online(user *objs.User)
	RechargeReset(user *objs.User)
	ApplyPay(user *objs.User, rechargeId int32, payNum int) (string, error, bool)
	TestPay(user *objs.User, rechargeId int) error
	/**
    *  @Description: 申请充值
    *  @param user
    *  @param payNum
    *  @param payType
    *  @param typeId
    *  @param fromBg 是否后台申请
    *  @return string
    *  @return error
    *  @return bool
    *  @return *modelGame.OrderDb
    **/
	Pay(user *objs.User, payNum, payType, typeId int,fromBg bool) (string, error, bool, *modelGame.OrderDb)

	/**
	 *  @Description: 跨服中心通知充值结果
	 *  @param req
	 *  @return error
	 **/
	NotifyBuy(req *pbserver.RechageCcsToGsReq) error

	//连续充值
	ContRechargeReset(user *objs.User, reset bool)
	ContRechargeWrite(user *objs.User, buyNum int)
	ContRechargeReceive(user *objs.User, contRechargeId int, op *ophelper.OpBagHelperDefault) error
}
