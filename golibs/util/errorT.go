package util

type IError interface {
	Error() string
	ErrCode() int
}

type ErrorT struct {
	Code    int
	Message string
}

func (this *ErrorT) Error() string {
	return this.Message
}

func (this *ErrorT) ErrCode() int {
	return this.Code
}
