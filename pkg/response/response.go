package response

import (
	"TikTokServer/pkg/errorcode"
	"TikTokServer/pkg/tlog"
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

func Success(c *gin.Context, e errorcode.HttpError, data interface{}) {
	if data == nil {
		ginResponse(c, e.HttpCode, Response{e.ErrCode.Code, e.ErrCode.Msg})
		return
	} else {
		WrapHttpErr(c, e, data)
		ginResponse(c, e.HttpCode, data)
	}
}

func Fail(c *gin.Context, e errorcode.HttpError, data interface{}) {
	if data == nil {
		ginResponse(c, e.HttpCode, Response{e.ErrCode.Code, e.ErrCode.Msg})
		return
	} else {
		WrapHttpErr(c, e, data)
		ginResponse(c, e.HttpCode, data)
		// TODO: c.Abort()
	}

}

func WrapHttpErr(c *gin.Context, e errorcode.HttpError, data interface{}) {
	getValue := reflect.ValueOf(data)
	field := getValue.Elem().FieldByName("StatusMsg")
	if field.CanSet() {
		field.SetString(e.Msg)
	} else {
		tlog.Debug("cant set msg")
	}
	fieldCode := getValue.Elem().FieldByName("StatusCode")
	if fieldCode.CanSet() {
		fieldCode.SetInt(int64(e.ErrCode.Code))
	} else {
		tlog.Debug("cant set StatusCode")
	}
}
