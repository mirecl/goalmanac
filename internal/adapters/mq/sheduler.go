package mq

import (
	"context"
	"encoding/json"

	"github.com/mirecl/goalmanac/internal/adapters/http"
	"github.com/mirecl/goalmanac/internal/domain/entities"
)

// ServeSheduler ...
func (b *APIServerMQ) ServeSheduler() error {
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

	// Настройка для распределения нагрузки
	err = ch.Qos(1, 0, false)
	if err != nil {
		return err
	}

	// Создаем/Подключаемся к oчереди MQ
	q, err := getQueue(ch)
	if err != nil {
		return err
	}
	msgs, _ := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	b.Logger.Infof("%s", "Starting MQ Sheduler ...")
	ctx := context.Background()
	for d := range msgs {
		var event entities.Event
		// Преобразовываем сообщение
		err := json.Unmarshal(d.Body, &event)
		if err != nil {
			b.Logger.Errorf(http.F(), "%s", err)
			// Возвращаем в очередь при ошибке
			d.Reject(true)
			continue
		}
		// Отправляем сообщение
		b.Logger.Infof("Send Notify fo User %s, Title %s", event.User, event.Title)
		// Подтверждаем отправку сообщения
		d.Ack(false)
		// Записываем статус отправки сообщения
		err = b.Storage.ChangeStatusEventNotify(ctx, event.ID)
		if err != nil {
			b.Logger.Errorf(http.F(), "%s", err)
			continue
		}
	}
	return nil
}
