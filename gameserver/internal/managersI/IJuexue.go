package managersI

import (
	"cqserver/gameserver/internal/objs"
	"cqserver/gameserver/internal/ophelper"
)

type IJuexueManager interface {

	/**
	 *  @Description: 绝学升级
	 *  @param user
	 *  @param id
	 *  @param op
	 *  @return error
	 */
	JuexueUpLevel(user *objs.User, id int, op *ophelper.OpBagHelperDefault) error
}
