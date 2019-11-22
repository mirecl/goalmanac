package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mirecl/goalmanac/internal/adapters/db"
	"github.com/mirecl/goalmanac/internal/domain/usecases"
)

func main() {
	memdb, _ := db.NewMemEventStorage()
	uses := &usecases.EventUsecases{EventStorage: memdb}
	start := time.Now()
	end := time.Now()
	var err error
	for i := 0; i < 20000; i++ {
		err = uses.AddEvent(context.Background(), "Grazhdankov", "Golang", "Tutorial and  big test", &start, &end)
	}
	cnt, _ := uses.GetCountEvent(context.Background())
	fmt.Println(*cnt, err)
}
