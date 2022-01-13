package service

import (
	"axiangcoding/antonstar/api-system/internal/app/entity"
	"axiangcoding/antonstar/api-system/pkg/logging"
	"axiangcoding/antonstar/api-system/pkg/mq"
	"encoding/json"
	"github.com/streadway/amqp"
)

func SendMessage(body entity.MQBody) error {
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
		logging.Errorf("Send message to queue error: %s", err)
	}
	return nil
}

func SendMessages(bodies ...entity.MQBody) error {
	for _, body := range bodies {
		err := SendMessage(body)
		if err != nil {
			return err
		}
	}
	return nil
}
