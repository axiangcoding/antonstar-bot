package service

import (
	"axiangcoding/antonstar/api-system/logging"
	"axiangcoding/antonstar/api-system/mq"
	"encoding/json"
	"github.com/streadway/amqp"
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
		err := SendMessage(body)
		if err != nil {
			return err
		}
	}
	return nil
}
