package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/protobuf/pb"
)

type IDailyActivity interface {
	/**
	 *  @Description: 日常活动入口
	 *  @param user
	 *  @param activityId	dailyActivity表id
	 *  @return error
	 */
	EnterDailyActivity(user *objs.User, activityId int,stageId int) error

	List(ack *pb.DailyActivityListAck)
}
