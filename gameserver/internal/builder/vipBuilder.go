package builder

import (
	"cqserver/gamelibs/publicCon/constUser"
	"cqserver/gameserver/internal/objs"
)

func BuildVipGift(user *objs.User) []int32 {
	pbVipGift := make([]int32, 0)
	userVipGift := user.VipGift
	for lv := range userVipGift {
		pbVipGift = append(pbVipGift, int32(lv))
	}
	return pbVipGift
}

func BuildVipCustomer(user *objs.User) bool {
	flag := false
	if user.VipCustomer == constUser.VIP_CUSTOMER_LOOK_YES {
		flag = true
	}
	return flag
}