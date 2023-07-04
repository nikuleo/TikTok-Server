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
	ossPath, err := UploadFileToOss("test.jpeg", "/home/niku/图片/test/test.jpeg", "covers/")
	fmt.Println("ossPath:", ossPath, err)
}
