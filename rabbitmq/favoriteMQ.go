package rabbitmq

import (
	"TikTokServer/model"
	"TikTokServer/pkg/tlog"
	"strconv"
	"strings"

	"github.com/streadway/amqp"
)

var (
	FavChannel    *amqp.Channel
	DisFavChannel *amqp.Channel
)

func InitFavMQ() {
	var err error
	FavChannel, err = MQConn.Channel()
	if err != nil {
		tlog.Error(err.Error())
		panic(err)
	}

	DisFavChannel, err = MQConn.Channel()
	if err != nil {
		tlog.Error(err.Error())
		panic(err)
	}
}

func SendFavoriteMsg(msg string) error {
	q, err := FavChannel.QueueDeclare(
		"fav_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = FavChannel.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func SendDisFavoriteMsg(msg string) error {
	q, err := DisFavChannel.QueueDeclare(
		"disfav_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	err = DisFavChannel.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)

	if err != nil {
		return err
	}
	return nil
}

func FavoriteConsumer() {
	q, err := FavChannel.QueueDeclare(
		"fav_queue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		tlog.Error(err.Error())
		panic(err)
	}

	msg, err := FavChannel.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		tlog.Error(err.Error())
		panic(err)
	}

	var forever chan struct{}

	go func() {
		for d := range msg {
			// tlog.Infof("receive message: %v", string(d.Body))
			// time.Sleep(time.Second * 1)
			// tlog.Infof("message %v done", string(d.Body))
			params := strings.Split(string(d.Body), " ")
			userID, _ := strconv.ParseInt(params[0], 10, 64)
			videoID, _ := strconv.ParseInt(params[1], 10, 64)

			// 如果入库失败重试两次
			for i := 0; i <= Retry; i++ {

				err := model.Favorite(userID, videoID)
				if err != nil {
					tlog.Error(err.Error())
				} else {
					break
				}

			}
			d.Ack(false) // 不管最后是否入库成功，都要 Ack，否则队列会积压消息
		}
	}()

	<-forever

}

func DisFavoriteConsumer() {
	q, err := DisFavChannel.QueueDeclare(
		"disfav_queue",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		tlog.Error(err.Error())
		panic(err)
	}

	msg, err := DisFavChannel.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		tlog.Error(err.Error())
		panic(err)
	}

	var forever chan struct{}

	go func() {
		for d := range msg {
			// tlog.Infof("receive message: %v", string(d.Body))
			// time.Sleep(time.Second * 1)
			// tlog.Infof("message %v done", string(d.Body))
			params := strings.Split(string(d.Body), " ")
			userID, _ := strconv.ParseInt(params[0], 10, 64)
			videoID, _ := strconv.ParseInt(params[1], 10, 64)

			// 如果入库失败重试两次
			for i := 0; i <= Retry; i++ {

				err := model.DisFavorite(userID, videoID)
				if err != nil {
					tlog.Error(err.Error())
				} else {
					break
				}

			}
			d.Ack(false)
		}
	}()

	<-forever

}
