package db

import (
	"context"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/stdlib" //db driver
	"github.com/jmoiron/sqlx"
	"github.com/mirecl/goalmanac/internal/adapters"
	"github.com/mirecl/goalmanac/internal/domain/entities"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

//MQEventStorage - инстанция БД для MQ
type MQEventStorage struct {
	db *sqlx.DB
}

//NewMQStorage - создаем инстанцию БД в памяти
func NewMQStorage(cfg *adapters.Config) (*MQEventStorage, error) {
	ctx := context.Background()
	// Создаем строку подключения
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Database,
	)
	// Подключаемся к БД
	db, err := sqlx.Open("pgx", dsn) // *sqlx.DB
	if err != nil {
		return nil, err
	}
	// Проверка подключения к БД
	if err = db.PingContext(ctx); err != nil {
		return nil, err
	}
	log.WithFields(log.Fields{"type": "db"}).Info("Connect to DB (MQ) - Good!")
	return &MQEventStorage{db: db}, nil
}

// GetEventNotify ...
func (m *MQEventStorage) GetEventNotify(ctx context.Context, period string) ([]*entities.Event, error) {
	// Итоговый массив данных
	eventArr := make([]*entities.Event, 0, 100)
	// Формируем SQL
	sql := `SELECT id, "user", title, body, starttime, endtime FROM almanac
			WHERE notify is null and (starttime BETWEEN :start and :end)`

	startTime := time.Now()
	// Определяем время окончания события
	timeEvent, err := time.ParseDuration(period)
	if err != nil {
		return nil, err
	}
	endTime := startTime.Add(timeEvent)

	// Выполняем запрос SQL
	rows, err := m.db.NamedQueryContext(context.Background(), sql, map[string]interface{}{
		"start": startTime,
		"end":   endTime,
	})
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Считываем данные
	for rows.Next() {
		var event entities.Event
		err := rows.StructScan(&event)
		if err != nil {
			return nil, err
		}
		eventArr = append(eventArr, &event)
	}
	return eventArr, nil
}

// ChangeStatusEventNotify ...
func (m *MQEventStorage) ChangeStatusEventNotify(ctx context.Context, id uuid.UUID) error {
	// Формируем SQL
	sql := `UPDATE almanac SET notify = :status WHERE id = :id`
	// Выполняем запрос SQL
	_, err := m.db.NamedExecContext(ctx, sql, map[string]interface{}{
		"status": "1",
		"id":     id,
	})
	return err
}
