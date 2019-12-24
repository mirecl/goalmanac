package db

import (
	"context"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/stdlib" //db driver
	"github.com/jinzhu/now"
	"github.com/jmoiron/sqlx"
	"github.com/mirecl/goalmanac/internal/adapters"
	"github.com/mirecl/goalmanac/internal/domain/entities"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

//SQLEventStorage - структура БД в памяти
type SQLEventStorage struct {
	db *sqlx.DB
}

//NewSQLStorage - создаем инстанцию БД в памяти
func NewSQLStorage(cfg *adapters.Config) (*SQLEventStorage, error) {
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
	log.WithFields(log.Fields{"type": "db"}).Info("Connect to DB - Good!")
	return &SQLEventStorage{db: db}, nil
}

//Save - сохраняем событие
func (s *SQLEventStorage) Save(ctx context.Context, e *entities.Event) error {
	// Формируем SQL
	sql := `INSERT INTO almanac( id, "user", title, body, starttime, endtime) 
	               VALUES (:id, :user, :title, :body, :start, :end)`
	// Выполняем запрос SQL
	_, err := s.db.NamedExecContext(ctx, sql, map[string]interface{}{
		"id":    e.ID,
		"user":  e.User,
		"title": e.Title,
		"body":  e.Body,
		"start": e.StartTime,
		"end":   e.EndTime,
	})
	return err
}

// Delete ...
func (s *SQLEventStorage) Delete(ctx context.Context, id uuid.UUID) error {
	// Формируем SQL
	sql := `DELETE FROM almanac WHERE id = :id`
	// Выполняем запрос SQL
	_, err := s.db.NamedExecContext(ctx, sql, map[string]interface{}{
		"id": id,
	})
	return err
}

// Update ...
func (s *SQLEventStorage) Update(ctx context.Context, e *entities.Event) error {
	// Формируем SQL
	sql := `UPDATE almanac 
			SET ("user", title, body, starttime, endtime) = (:user, :title, :body, :start, :end)
			WHERE id = :id`
	// Выполняем запрос SQL
	_, err := s.db.NamedExecContext(ctx, sql, map[string]interface{}{
		"id":    e.ID,
		"user":  e.User,
		"title": e.Title,
		"body":  e.Body,
		"start": e.StartTime,
		"end":   e.EndTime,
	})
	return err
}

// GetAll ...
func (s *SQLEventStorage) GetAll(ctx context.Context) ([]*entities.Event, error) {
	// Итоговый массив данных
	eventArr := make([]*entities.Event, 0, 100)
	// Формируем SQL
	sql := `SELECT id, "user", title, body, starttime, endtime FROM almanac`
	// Выполняем запрос SQL
	rows, err := s.db.NamedQueryContext(context.Background(), sql, map[string]interface{}{})
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

// GetForDay ...
func (s *SQLEventStorage) GetForDay(ctx context.Context, user string) ([]*entities.Event, error) {
	// Формируем SQL
	sql := `SELECT id, "user", title, body, starttime, endtime FROM almanac 
			WHERE "user" = :id and (starttime BETWEEN :start and :end)`
	// Определяем диапазоны
	from := now.BeginningOfDay()
	to := now.EndOfDay()
	// Выполняем запрос SQL
	events, err := s.getEventBetweenDays(ctx, user, sql, from, to)
	if err != nil {
		return nil, err
	}
	return events, nil
}

// GetForWeek ...
func (s *SQLEventStorage) GetForWeek(ctx context.Context, user string) ([]*entities.Event, error) {
	// Формируем SQL
	sql := `SELECT id, "user", title, body, starttime, endtime FROM almanac 
			WHERE "user" = :id and (starttime BETWEEN :start and :end)`
	// Определяем диапазоны
	from := now.BeginningOfWeek()
	to := now.EndOfWeek()
	// Выполняем запрос SQL
	events, err := s.getEventBetweenDays(ctx, user, sql, from, to)
	if err != nil {
		return nil, err
	}
	return events, nil
}

// GetForMonth ...
func (s *SQLEventStorage) GetForMonth(ctx context.Context, user string) ([]*entities.Event, error) {
	// Формируем SQL
	sql := `SELECT id, "user", title, body, starttime, endtime FROM almanac 
			WHERE "user" = :id and (starttime BETWEEN :start and :end)`
	// Определяем диапазоны
	from := now.BeginningOfMonth()
	to := now.EndOfMonth()
	// Выполняем запрос SQL
	events, err := s.getEventBetweenDays(ctx, user, sql, from, to)
	if err != nil {
		return nil, err
	}
	return events, nil

}

// getEventBetweenDays ...
func (s *SQLEventStorage) getEventBetweenDays(ctx context.Context, user, sql string, from, to time.Time) ([]*entities.Event, error) {
	// Итоговый массив данных
	eventArr := make([]*entities.Event, 0, 100)
	// Выполняем запрос SQL
	rows, err := s.db.NamedQueryContext(ctx, sql, map[string]interface{}{
		"id":    user,
		"start": from,
		"end":   to,
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
