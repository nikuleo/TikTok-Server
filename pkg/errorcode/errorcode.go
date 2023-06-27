package errorcode

import "fmt"

type ErrCode struct {
	Code int
	Msg  string
}

type HttpError struct {
	HttpCode int
	ErrCode
}

func (e ErrCode) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Code, e.Msg)
}

func NewErrorCode(code int, msg string) ErrCode {
	return ErrCode{
		Code: code,
		Msg:  msg,
	}
}

func NewHttpErr(code, httpCode int, msg string) HttpError {
	return HttpError{
		HttpCode: httpCode,
		ErrCode: ErrCode{
			Code: code,
			Msg:  msg,
		},
	}
}

func (e ErrCode) WithMsg(msg string) ErrCode {
	e.Msg = msg
	return e
}
