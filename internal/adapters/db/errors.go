package db

import "errors"

var (
	// ErrNotFoundDelete - не найдены данные
	ErrNotFoundDelete = errors.New("No Data for Delete")
	// ErrNotFoundUpdate - не найдены данные
	ErrNotFoundUpdate = errors.New("No Data for Update")
)
