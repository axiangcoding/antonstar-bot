package service

import (
	"encoding/json"
	"github.com/axiangcoding/ax-web/logging"
	"github.com/axiangcoding/ax-web/mq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func SendMessage(body mq.CrawBody) error {
	channel := mq.GetChannel()
	b, err := json.Marshal(body)
	if err != nil {
		logging.Errorf("Can't marshal json: %s", err)
		return err
	}
	err = channel.Publish(
		"",        // exchange
		"crawler", // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		})
	if err != nil {
		logging.Errorf("Send message to mq error: %s", err)
	}
	return nil
}

func SendMessages(bodies ...mq.CrawBody) error {
	for _, body := range bodies {
		if err := SendMessage(body); err != nil {
			return err
		}
	}
	return nil
}
