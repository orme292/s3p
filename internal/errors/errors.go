package errors

import (
	"fmt"
	"strings"

	z "github.com/rs/zerolog"
)

type AppError struct {
	module string
	Msg    string
	Sev    z.Level
	Err    error
}

func NewError(module, msg string, sev z.Level) *AppError {
	return newError(module, msg, sev, nil)
}

func NewErrorWithBase(e *AppError, base error) *AppError {
	return &AppError{
		module: e.module,
		Msg:    e.Msg,
		Sev:    e.Sev,
		Err:    base,
	}
}

func NewErrorNilCheck(e *AppError, err error) *AppError {
	if err == nil {
		return e
	}
	return NewErrorWithBase(e, err)
}

func newError(module string, msg string, sev z.Level, err error) *AppError {
	return &AppError{
		module: module,
		Msg:    msg,
		Sev:    sev,
		Err:    err,
	}
}

func (e *AppError) Error() string {
	var errStr string
	if e.Err != nil {
		errStr = fmt.Sprintf(" (%s)", e.Err.Error())
	}
	return fmt.Sprintf("%s: %s [%s]", strings.ToUpper(e.Sev.String()), e.Msg, errStr)
}

func (e *AppError) Debug() string {
	var errStr string
	if e.Err != nil {
		errStr = fmt.Sprintf("\n\tOriginal Error: %s", e.Err.Error())
	}
	return fmt.Sprintf("%s level error\n\tModule %q\n\tApplication Error: %s%s",
		strings.ToUpper(e.Sev.String()), strings.ToUpper(e.module), e.Msg, errStr)
}

func (e *AppError) Halts() bool {
	return e.Sev >= z.ErrorLevel
}

func (e *AppError) Unwrap() error {
	return e.Err
}
