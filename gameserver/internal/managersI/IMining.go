package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
	"cqserver/protobuf/pb"
)

type IMiningManager interface {
	Online(user *objs.User)
	// 每日重置
	MiningReset(user *objs.User, date int)
	// 升级矿工
	UpMiner(user *objs.User, isMax bool, op *ophelper.OpBagHelperDefault, ack *pb.MiningUpMinerAck) error
	// 购买次数
	BuyNum(user *objs.User, op *ophelper.OpBagHelperDefault, ack *pb.MiningBuyNumAck) error
	// 开始挖矿
	StartWork(user *objs.User, ack *pb.MiningStartAck) error
	// 挖矿列表
	List(ack *pb.MiningListAck)
	// 领取奖励
	Draw(user *objs.User, op *ophelper.OpBagHelperDefault, ack *pb.MiningDrawAck) error
	// 获取是否被抢
	GetDrawStatus(user *objs.User) bool
	// 查看被掠夺信息
	GetRobInfo(user *objs.User, ack *pb.MiningDrawLoadAck) error
	// 进入矿洞
	In(user *objs.User)error
	// 掠夺
	Rob(user *objs.User, id int) error
	// 掠夺列表
	RobList(user *objs.User, ack *pb.MiningRobListAck)

	/**
	 *  @Description: 抢夺战战斗结果
	 *  @param user
	 *  @param id
	 *  @param isWin
	 */
	RobFightBack(user *objs.User, id int, isWin bool)

	/**
	 *  @Description: 夺回战 战斗结果
	 *  @param user
	 *  @param id
	 *  @param isWin
	 */
	RobBackFightBack(user *objs.User, id int, isWin bool)

	/**
	 *  @Description: 夺回奖励
	 *  @param user
	 *  @param id	挖矿记录id(表id)
	 *  @return error
	 */
	RobBack(user *objs.User, id int) error
}
