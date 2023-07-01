package response

import (
	message "TikTokServer/idl/gen"
	"TikTokServer/pkg/errorcode"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	he := errorcode.HttpSuccess
	dbErr := errorcode.ErrHttpDatabase
	dbErr.SetError(errors.New("db error"))
	Success(ctx, nil, nil)
	Success(ctx, he, &message.DouyinUserResponse{
		StatusCode: int32(he.ErrCode),
		StatusMsg:  "",
		User:       nil,
	})
	Success(ctx, dbErr, nil)
}
