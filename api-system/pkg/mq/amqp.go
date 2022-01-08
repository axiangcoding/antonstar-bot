package mq

import (
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
	failOnError(err, "Failed to declare a queue")
}

func failOnError(err error, msg string) {
	if err != nil {
		logging.Fatalf("%s: %s", msg, err)
	}
}

func initMQ() *amqp.Channel {
	conn, err := amqp.Dial("amqp://localhost:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	//defer ch.Close()
	return ch
}

func GetChannel() *amqp.Channel {
	return channel
}
