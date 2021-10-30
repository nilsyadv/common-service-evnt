package error

import (
	"fmt"
	"runtime"
	"strconv"
)

type IUnexpected interface {
	Error() string
	GetErrorCode() string
	GetUnexpected() string
	GetCause() error
}

type UnexpectedError struct {
	cause      error
	stackTrace string
	errCode    string
}

func CreateUnexpectedError(errCode string, err error) UnexpectedError {
	const depth = 20
	var ptrs [depth]uintptr
	n := runtime.Callers(2, ptrs[:])
	ptrSlice := ptrs[0:n]
	stack := ""
	for _, pc := range ptrSlice {
		stackFunc := runtime.FuncForPC(pc)
		_, line := stackFunc.FileLine(pc)
		stack = stack + stackFunc.Name() + ":" + strconv.Itoa(line) + "\n"
	}
	return UnexpectedError{errCode: errCode, cause: err, stackTrace: stack}
}

func (st UnexpectedError) GetErrorCode() string {
	return st.errCode
}

func (st UnexpectedError) GetUnexpected() string {
	return st.stackTrace
}

func (st UnexpectedError) GetCause() error {
	return st.cause
}

// Error returns the error string
func (e UnexpectedError) Error() string {
	return fmt.Sprintf("[%v:%v]", e.errCode, e.cause)
}
