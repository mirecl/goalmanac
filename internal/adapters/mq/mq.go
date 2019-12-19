package mq

import (
	"fmt"

	"github.com/mirecl/goalmanac/internal/adapters"
	"github.com/mirecl/goalmanac/internal/domain/interfaces"
	"github.com/streadway/amqp"
)

// APIServerMQ - структура MQ
type APIServerMQ struct {
	Logger  interfaces.MQLogger
	Storage interfaces.MQStorage
	Config  *adapters.Config
}

// getConnect - подключение к MQ
func getConnect(cfg *adapters.Config) (*amqp.Connection, error) {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		cfg.MQ.User,
		cfg.MQ.Password,
		cfg.MQ.Host,
		cfg.MQ.Port,
	)
	mq, err := amqp.Dial(dsn)
	if err != nil {
		return nil, err
	}
	return mq, nil
}

// getChannel - подключение к каналу
func getChannel(mq *amqp.Connection) (*amqp.Channel, error) {
	ch, err := mq.Channel()
	if err != nil {
		return nil, err
	}
	return ch, nil
}

// getQueue - Создание очереди
func getQueue(ch *amqp.Channel) (*amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		"event", // name
		true,    // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		return nil, err
	}
	return &q, nil
}
