package errors

import (
	"fmt"
	"runtime"
)

type withCode struct {
	code  int
	err   error
	cause error
	stack *stack
}

func WithCode(code int, format string, args ...any) error {
	return &withCode{
		err:   fmt.Errorf(format, args...),
		code:  code,
		stack: callers(),
	}
}

func WrapC(code int, err error, format string, args ...any) error {
	if err == nil {
		return nil
	}

	return &withCode{
		code:  code,
		err:   fmt.Errorf(format, args...),
		cause: err,
		stack: callers(),
	}
}

func (w *withCode) Error() string {
	return fmt.Sprintf("%v", w.err)
}

func (w *withCode) Cause() error {
	return w.cause
}

func (w *withCode) Unwrap() error {
	return w.cause
}

type stack []uintptr

func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[0:n]
	return &st
}

func Cause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}

		if cause.Cause() == nil {
			break
		}

		err = cause.Cause()
	}

	return err
}

func Wrap(err error, msg string) error {
	if err == nil {
		return nil
	}

	if e, ok := err.(*withCode); ok {
		return &withCode{
			err:   fmt.Errorf(msg),
			cause: err,
			code:  e.code,
			stack: callers(),
		}
	}

	err = &withMessage{
		err: err,
		msg: msg,
	}

	return &withStack{
		err:   err,
		stack: callers(),
	}
}

type withMessage struct {
	err error
	msg string
}

func (w *withMessage) Error() string {
	return fmt.Sprintf("%s: %v", w.msg, w.err)
}

type withStack struct {
	err   error
	stack *stack
}

func (w *withStack) Error() string {
	return ""
}
