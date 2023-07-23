package rabbitmq

import (
	"TikTokServer/pkg/config"
	"TikTokServer/pkg/tlog"
	"fmt"

	"github.com/streadway/amqp"
)

var (
	MQConn *amqp.Connection
)

func InitRabbitMQ() {
	var err error
	cfg := config.GetConfig("rabbitmqConfig").Viper
	address := cfg.GetString("rabbitmq.address")
	port := cfg.GetString("rabbitmq.port")
	user := cfg.GetString("rabbitmq.user")
	password := cfg.GetString("rabbitmq.password")

	amqpURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, address, port)
	MQConn, err = amqp.Dial(amqpURL)
	if err != nil {
		tlog.Error(err.Error())
	}

	InitFavMQ()

	RunConsumer()

	tlog.Info("RabbitMQ init success.")
}

func RunConsumer() {
	go FavoriteConsumer()
	go DisFavoriteConsumer()
}

func MQClose() {
	FavChannel.Close()
	DisFavChannel.Close()
	MQConn.Close()
	tlog.Info("RabbitMQ close.")
}
