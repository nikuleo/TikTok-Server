package ossBucket

import (
	"TikTokServer/pkg/tlog"
	"fmt"
	"testing"
)

func init() {
	tlog.InitLog()
	OssInit()
}

func TestUploadFileToOss(t *testing.T) {
	ossPath, err := UploadCoverToOss("test.jpeg", "/home/niku/图片/test/test.jpeg")
	fmt.Println("ossPath:", ossPath, err)
}
