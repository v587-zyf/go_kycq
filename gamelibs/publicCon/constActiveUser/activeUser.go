package constActiveUser

import "time"

/**
 * 活跃玩家数据
 */
type ActiveUsersInfo struct {
	UserId     int       //玩家userID
	UpdateTime time.Time //更新数据的时候，判断是否活跃
	ServerId   int       //玩家所在serverID
	Nickname   string    //玩家名字
	Avatar     string    //玩家头像
	Vip        int       //玩家VIP等级
	Combat     int       //玩家战斗力
}


