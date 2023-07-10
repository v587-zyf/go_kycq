// referenced from go-errors: https://github.com/go-errors

package errex

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime"
)

var MaxStackDepth = 20

type errEx interface {
	error
	BindExternalErr(err error) errEx
	ErrorStack() string
}

type errorObj struct {
	Err         error
	ExternalErr error
	stack       []uintptr
	frames      []StackFrame
}

func New(e interface{}) errEx {
	var err error
	switch e := e.(type) {
	case errEx:
		return e
	case error:
		err = e
	default:
		err = fmt.Errorf("%v", e)
	}
	stack := make([]uintptr, MaxStackDepth)
	length := runtime.Callers(2, stack[:])
	return &errorObj{
		Err:   err,
		stack: stack[:length],
	}
}

func (this *errorObj) Error() string {
	return this.Err.Error()
}

func (this *errorObj) BindExternalErr(err error) errEx {
	this.ExternalErr = err
	return this
}

func (this *errorObj) Stack() []byte {
	buf := bytes.Buffer{}
	for _, frame := range this.StackFrames() {
		buf.WriteString(frame.String())
	}
	return buf.Bytes()
}

func (this *errorObj) ErrorStack() string {
	return this.TypeName() + " " + this.Error() + "\n" + string(this.Stack())
}

func (this *errorObj) StackFrames() []StackFrame {
	if this.frames == nil {
		this.frames = make([]StackFrame, len(this.stack))

		for i, pc := range this.stack {
			this.frames[i] = NewStackFrame(pc)
		}
	}

	return this.frames
}

func (this *errorObj) TypeName() string {
	return reflect.TypeOf(this.Err).String()
}
