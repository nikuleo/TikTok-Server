package response

import (
	"TikTokServer/pkg/errorcode"
	"TikTokServer/pkg/tlog"
	"errors"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Response struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func ginResponse(c *gin.Context, httpStatus int, data interface{}) {
	c.JSON(httpStatus, data)
}

func Success(c *gin.Context, err error, data interface{}) {

	var he errorcode.HttpError

	if err == nil {
		he = errorcode.HttpSuccess
	} else {
		errors.As(err, &he)
	}

	if data == nil {
		ginResponse(c, he.HttpStatus, Response{he.ErrCode, he.Msg})
		return
	} else {
		WrapHttpErr(c, he, data)
		ginResponse(c, he.HttpStatus, data)
	}
}

func Fail(c *gin.Context, err error, data interface{}) {

	var he errorcode.HttpError
	if errors.As(err, &he) {
	} else {
		he = errorcode.ErrHttpUnknown
		he.SetError(err)
	}

	if data == nil {
		ginResponse(c, he.HttpStatus, Response{he.ErrCode, he.Msg})
		return
	} else {
		WrapHttpErr(c, he, data)
		ginResponse(c, he.HttpStatus, data)
		// TODO: c.Abort()
	}

}

// 反射包装返回的 json 结构
func WrapHttpErr(c *gin.Context, e errorcode.HttpError, data interface{}) {
	getValue := reflect.ValueOf(data)
	field := getValue.Elem().FieldByName("StatusMsg")
	if field.CanSet() {
		field.SetString(e.Msg + e.Error())
		// protobuf 字段是指针类型，需要用 reflect.ValueOf(&e.Msg)
		// field.Set(reflect.ValueOf(&e.Msg))
	} else {
		tlog.Debug("cant set msg")
	}
	fieldCode := getValue.Elem().FieldByName("StatusCode")
	if fieldCode.CanSet() {
		fieldCode.SetInt(int64(e.ErrCode))
	} else {
		tlog.Debug("cant set StatusCode")
	}
}
