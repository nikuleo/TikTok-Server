package ossBucket

import (
	"TikTokServer/pkg/config"
	"TikTokServer/pkg/errorcode"
	"TikTokServer/pkg/tlog"
	"fmt"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var client *oss.Client

var (
	endpoint        string
	accessKeyID     string
	accessKeySecret string
	bucketName      string
	prefix          string
)

func OssInit() {
	var err error
	cfg := config.GetConfig("ossConfig")
	viper := cfg.Viper
	endpoint = viper.GetString("OSS.endpoint")
	accessKeyIDPath := viper.GetString("OSS.accessKeyIDPath")
	accessKeySecretPath := viper.GetString("OSS.accessKeySecretPath")

	dataID, err := os.ReadFile(accessKeyIDPath)
	if err != nil {
		tlog.Errorf("read key file failed: %v", err)
	}
	dataSecret, err := os.ReadFile(accessKeySecretPath)
	if err != nil {
		tlog.Errorf("read key file failed: %v", err)
	}

	accessKeyID = string(dataID[:len(dataID)-1])
	accessKeySecret = string(dataSecret[:len(dataSecret)-1])

	bucketName = viper.GetString("OSS.bucketName")
	prefix = viper.GetString("OSS.prefix")
	client, err = oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		errCode := errorcode.ErrHttpOssConnectFailed
		errCode.SetError(err)
		tlog.Errorf(errCode.Error(), err)
	}
	bucketInfo, _ := client.GetBucketInfo(bucketName)
	tlog.Infof("bucketInfo: %v", bucketInfo)
}

func UploadFileToOss(fileName string, filePath string, objPath string) (string, error) {
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		errCode := errorcode.ErrHttpOssConnectFailed
		errCode.SetError(err)
		tlog.Errorf(errCode.Error(), err)
		return "", errCode
	}
	objectName := objPath + fileName
	err = bucket.PutObjectFromFile(objectName, filePath, oss.Progress(&OssProgressListener{}))
	if err != nil {
		errCode := errorcode.ErrHttpOssBucketUploadFailed
		errCode.SetError(err)
		tlog.Errorf(errCode.Error(), err)
		return "", err
	}

	return prefix + objectName, nil
}

// 定义进度条监听器。
type OssProgressListener struct {
}

// 定义进度变更事件处理函数。
func (listener *OssProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		fmt.Printf("Transfer Started, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	case oss.TransferDataEvent:
		fmt.Printf("\rTransfer Data, ConsumedBytes: %d, TotalBytes %d, %d%%.",
			event.ConsumedBytes, event.TotalBytes, event.ConsumedBytes*100/event.TotalBytes)
	case oss.TransferCompletedEvent:
		fmt.Printf("\nTransfer Completed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	case oss.TransferFailedEvent:
		fmt.Printf("\nTransfer Failed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	default:
	}
}
