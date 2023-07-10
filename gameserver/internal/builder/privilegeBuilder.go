package builder

import "cqserver/gameserver/internal/objs"

func BuildPrivilege(user *objs.User) []int32 {
	return user.Privilege.KeysInt32()
}
