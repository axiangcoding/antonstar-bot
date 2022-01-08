package service

import (
	"axiangcoding/antonstar/api-system/pkg/logging"
	"axiangcoding/antonstar/api-system/pkg/mq"
	"github.com/streadway/amqp"
)

func SendMessage() {
	channel := mq.GetChannel()
	body := "Hello World!"
	err := channel.Publish(
		"",        // exchange
		"crawler", // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		logging.Error(err)
	}
}
