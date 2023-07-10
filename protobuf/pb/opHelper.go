package pb

import "cqserver/golibs/nw"

type OpGoodsHelper interface {
	SetOpType(t int) // for record
	GetOpType() int  // for record
	/**
	 *  @Description: 二级来源
	 *  @return int
	 **/
	OpTypeSecond() int
	/**
	 *  @Description: 二级来源
	 *  @param opTypeSecond
	 **/
	SetOpTypeSecond(opTypeSecond int)
	OnGoodsChange(goods interface{}, count int)
	ToGoodsChangeMessages() []nw.ProtoMessage
}
