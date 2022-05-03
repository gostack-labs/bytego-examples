package xresult

import (
	"errors"

	"github.com/gostack-labs/bytego"
)

//internalError
type internalError struct {
	err error
}

func (e *internalError) Error() string {
	return e.err.Error()
}

func InternalError(err error) error {
	return &internalError{err: err}
}

//errorResult
type errorResult struct {
	Code  int    `json:"code,omitempty"`
	Msg   string `json:"msg,omitempty"`
	cause error
}

func (r *errorResult) ErrCode() int {
	return r.Code
}

func (r *errorResult) Error() string {
	return r.Msg
}

func (r *errorResult) Cause() error {
	err := r.cause
	for err != nil {
		u := errors.Unwrap(err)
		if u == nil {
			return err
		}
		err = u
	}
	return err
}

func Fail(code int, msg string) *errorResult {
	return &errorResult{
		Code: code,
		Msg:  msg,
	}
}

func (r *errorResult) JSON(httpCode int, c *bytego.Ctx) error {
	return c.JSON(httpCode, r)
}

func Error(code int, err error) *errorResult {
	return &errorResult{
		Code:  code,
		Msg:   err.Error(),
		cause: err,
	}
}

//result
type successResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func (r *successResult) JSON(c *bytego.Ctx) error {
	return c.JSON(200, r)
}

func Success(data interface{}) *successResult {
	return &successResult{
		Code: 0,
		Data: data,
	}
}
