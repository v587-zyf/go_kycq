package constBag

const (
	BAG_TYPE_COMMON = 1 //普通背包

	BAG_ADD_TYPE_INIT   = 1 //背包格子增加 初始背包
	BAG_ADD_TYPE_VIP    = 2 //背包格子增加 vip
	BAG_ADD_TYPE_ONLINE = 3 //背包格子增加 在线时长
	BAG_ADD_TYPE_ITEM   = 4 //背包格子增加 道具

	//--------仓库相关
	WAREHOUSE_BAG_TYPE_COMMON = 2 //仓库背包

	WAREHOUSE_BAG_ADD_TYPE_INIT = 5 //初始仓库背包  bag表 type= 5
	WAREHOUSE_BAG_ADD_TYPE_ITEM = 6 //仓库格子增加 道具
)

const (
	Unbind = 0 //绑定标识
	Bind   = 1 //非绑定标识
)

const (
	COMPOSE_LUCKY_NOT = 0
	COMPOSE_LUCKY_CAN = 1
)
