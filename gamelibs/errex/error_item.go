package errex

import (
	"cqserver/golibs/logger"
	"cqserver/golibs/util"
	"cqserver/protobuf/pb"
	"cqserver/protobuf/pbserver"
	"encoding/json"
	"fmt"
)

type ErrorItem struct {
	*util.ErrorT
}

func NewErrorItem(format string, args ...interface{}) *ErrorItem {
	return Create(-999, fmt.Sprintf(format, args...))
}

func Create(code int, message string) *ErrorItem {
	e := &ErrorItem{
		ErrorT: &util.ErrorT{},
	}
	e.Code = code
	e.Message = message
	return e
}

func (this *ErrorItem) Error() string {
	return this.Message
}

func (this *ErrorItem) ErrCode() int {
	return this.Code
}

func (this *ErrorItem) CloneWithMsg(format string, args ...interface{}) *ErrorItem {
	return Create(this.Code, fmt.Sprintf(format, args...))
}

func (this *ErrorItem) SprintfErrMsg(args ...interface{}) *ErrorItem {

	this.Message = fmt.Sprintf(this.Message, args...)
	return this
}

//返回异常
func (this *ErrorItem) GetErrorByJson() string {

	rt, _ := json.Marshal(this)
	return string(rt)
}

func GetErrorMessage(err error) string {
	ei, ok := err.(*ErrorItem)
	if !ok {
		return "unkonw error"
	}
	return ei.Message
}

func BuildClientErrorAck(err error) *pb.ErrorAck {
	ei, ok := err.(*ErrorItem)
	if !ok {
		logger.Error("未知异常错误：%v", err)
		return &pb.ErrorAck{
			Code:    -1,
			Message: "unkonw error",
		}
	}
	return &pb.ErrorAck{
		Code:    int32(ei.Code),
		Message: ei.Message,
	}
}

func BuildServerErrorAck(err error) *pbserver.ErrorAck {
	ei, ok := err.(*ErrorItem)
	if !ok {
		logger.Error("未知异常错误：%v", err)
		return &pbserver.ErrorAck{
			Code:    -1,
			Message: "unkonw error",
		}
	}
	return &pbserver.ErrorAck{
		Code:    int32(ei.Code),
		Message: ei.Message,
	}
}
