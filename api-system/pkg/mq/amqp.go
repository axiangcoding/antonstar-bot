package mq

import (
	"axiangcoding/antonstar/api-system/internal/app/conf"
	"axiangcoding/antonstar/api-system/pkg/logging"
	"github.com/streadway/amqp"
)

var channel *amqp.Channel

func Setup() {
	channel = initMQ()
	_, err := channel.QueueDeclare(
		"crawler", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		logging.Fatalf("Failed to declare a queue: %s", err)
	}
}

func initMQ() *amqp.Channel {
	conn, err := amqp.Dial(conf.Config.MQ.Source)
	if err != nil {
		logging.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	ch, err := conn.Channel()
	if err != nil {
		logging.Fatalf("Failed to open a channel: %s", err)
	}
	return ch
}

func GetChannel() *amqp.Channel {
	return channel
}
