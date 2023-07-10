package managersI

import (
	"cqserver/gamelibs/gamedb"
	"cqserver/gameserver/internal/objs"
)

type IConditionManager interface {
	/**
	 *  @Description: 条件检查
	 *  @param user
	 *  @param condition
	 *  @param conditionValue
	 *  @return int		当前值
	 *  @return bool    是否通过
	 */
	Check(user *objs.User, heroIndex, condition int, conditionValue int) (int, bool)

	CheckConditionType(id int) bool

	/**
	 *  @Description: 		多条件检查 全满足
	 *  @param user
	 *  @param conditions
	 *  @return bool		是否通过
	 */
	CheckMulti(user *objs.User, heroIndex int, conditions map[int]int) bool

	/**
	 *  @Description: 多条件检查
	 *  @param user
	 *  @param heroIndex
	 *  @param conditions
	 *  @param conditionType	(pb.CONDITIONTYPE_ALL 全满足 pb.CONDITIONTYPE_JUST_ONE 满足之一)
	 *  @return bool
	 */
	CheckMultiByType(user *objs.User, heroIndex int, conditions map[int]int, conditionType int) bool

	CheckItemIsCanBeUse(user *objs.User, itemId int) (bool, *gamedb.EquipEquipCfg)
	/**
	 *  @Description: 记录condition数据
	 *  @param user
	 *  @param id		pb中condition类型
	 *  @param value	值
	 */
	RecordCondition(user *objs.User, id int, value []int)
	GetConditionData(user *objs.User, id int, value int) int

	//  CheckFunctionOpen
	//  @Description: 判断活动是否开启
	//  @param user
	//  @param moduleType  function 表 id
	//  @return error
	//
	CheckFunctionOpen(user *objs.User, moduleType int) error

	CheckBySlice(user *objs.User, heroIndex int, condition []int) ([]int, bool)
	CheckMultiBySlice2(user *objs.User, heroIndex int, conditions [][]int) bool
	CheckMultiByMap(user *objs.User, heroIndex int, conditions map[int]int) bool
}
