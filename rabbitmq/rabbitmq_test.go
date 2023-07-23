package rabbitmq

import (
	"TikTokServer/pkg/tlog"
	"fmt"
	"testing"
)

func init() {
	tlog.InitLog()
	// InitRabbitMQ()
}

func TestInitRabbitMQ(t *testing.T) {
	// defer MQClose()
	InitRabbitMQ()
}

func TestSendFavoriteMsg(t *testing.T) {
	defer MQClose()
	InitRabbitMQ()

	RunConsumer()

	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("userID: %d, videoID: %d", i, i)
		err := SendFavoriteMsg(msg)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestFavroiteConsumer(t *testing.T) {
	// defer MQClose()
	InitRabbitMQ()
	FavoriteConsumer()
}

func TestRabbitMQClose(t *testing.T) {
	MQClose()
}
