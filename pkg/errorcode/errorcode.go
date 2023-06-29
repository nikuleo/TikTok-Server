package errorcode

import (
	"errors"
)

type HttpError struct {
	ErrCode    int    // 错误码
	HttpStatus int    // http 状态码, REST API 风格
	Msg        string // 暴露给客户端的信息
	Err        error  // 原始错误 log 与 错误类型判断信息
}

func (e HttpError) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return e.Msg
}

// 重写 Unwarp 方法，保证一致性
func (e HttpError) Unwarp() error {
	return e.Err
}

// 重写 Dig 方法
func (e HttpError) Dig() HttpError {
	var he HttpError
	if errors.As(e.Err, &he) {
		return he.Dig()
	}
	return e
}

func NewHttpErr(code, httpCode int, msg string, err error) HttpError {
	return HttpError{
		ErrCode:    code,
		HttpStatus: httpCode,
		Msg:        msg,
		Err:        err,
	}
}

func (e *HttpError) SetMsg(msg string) {
	e.Msg = msg
}

func (e *HttpError) SetError(err error) {
	e.Err = err
}
