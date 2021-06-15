package bootstrap

import (
	"github.com/koines/code-cadets-2021/lecture_4/03_bet_acceptance_api/cmd/config"
	"github.com/streadway/amqp"
)

// RabbitMq bootstraps the rabbit mq connection.
func RabbitMq() *amqp.Channel {
	conn, err := amqp.Dial(config.Cfg.Rabbit.ConnectionQueue)
	if err != nil {
		panic(err)
	}

	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	return channel
}
