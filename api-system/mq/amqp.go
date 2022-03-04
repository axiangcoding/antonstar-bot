package mq

import (
	"axiangcoding/antonstar/api-system/logging"
	"axiangcoding/antonstar/api-system/settings"
	"github.com/streadway/amqp"
)

const (
	SourceGaijin       = "gaijin"
	SourceThunderskill = "thunder_skill"
)

type CrawBody struct {
	// 查询的ID
	QueryID string `json:"query_id"`
	// 调用的目标爬虫
	Target []string `json:"target"`
	// 查询玩家的昵称
	Nickname string `json:"nickname"`
}

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
		logging.Fatalf("Failed to declare a mq: %s", err)
	}
}

func initMQ() *amqp.Channel {
	conn, err := amqp.Dial(settings.Config.MQ.Source)
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
