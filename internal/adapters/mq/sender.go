package mq

import (
	"context"
	"encoding/json"
	"os"
	"os/signal"
	"time"

	"github.com/mirecl/goalmanac/internal/adapters/http"
	"github.com/streadway/amqp"
)

// ServeSender ...
func (b *APIServerMQ) ServeSender() error {
	// Подключаемся к MQ
	mq, err := getConnect(b.Config)
	if err != nil {
		return err
	}
	defer mq.Close()
	// Подключаемся к каналу MQ
	ch, err := getChannel(mq)
	if err != nil {
		return err
	}
	defer ch.Close()

	// Создаем/Подключаемся к oчереди MQ
	q, err := getQueue(ch)
	if err != nil {
		return err
	}
	b.Logger.Infof("%s", "Starting MQ Sender ...")

	// Перехватываем сигналы завершения
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	dPolling, err := time.ParseDuration(b.Config.MQ.Polling)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(dPolling)
	ctx := context.Background()

L:
	for {
		select {
		case <-c: // Ждем завершения программы
			break L
		case <-ticker.C: // Каждые N секунд проверяем события для отправки
			// Выборка событий для отправки
			events, err := b.Storage.GetEventNotify(ctx, b.Config.MQ.Period)
			if err != nil {
				b.Logger.Errorf(http.F(), "%s", err)
				break
			}
			// Отправка событий в очередь
			for _, e := range events {
				payload, err := json.Marshal(e)
				if err != nil {
					b.Logger.Errorf(http.F(), "%s", err)
					continue
				}
				// Публикация в очередь
				err = ch.Publish(
					"",     // exchange
					q.Name, // routing key
					false,  // mandatory
					false,  // immediate
					amqp.Publishing{
						DeliveryMode: amqp.Persistent,
						ContentType:  "application/json",
						Body:         payload,
					})
				if err != nil {
					b.Logger.Errorf(http.F(), "%s", err)
				} else {
					b.Logger.Infof("Send event to MQ for UID %s", e.ID)
				}
			}
		}
	}
	b.Logger.Infof("%s", "MQ Sender shutting down...")
	return nil
}
