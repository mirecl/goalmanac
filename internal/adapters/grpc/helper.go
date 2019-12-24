package grpc

import (
	"fmt"
	"runtime"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// F ...
func F() string {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Sprintf("%s:%d", frame.File, frame.Line)
}

// Преобразовать время + продолжительность
func getTime(start *timestamp.Timestamp, duration string) (*time.Time, *time.Time, error) {
	// Устанавлтваем время (location)
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return nil, nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// Конвертируем время начала события
	t, err := ptypes.Timestamp(start)
	if err != nil {
		return nil, nil, status.Error(codes.InvalidArgument, err.Error())
	}
	t = t.In(loc)
	startTime := time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0, loc)

	// Определяем время окончания события
	timeEvent, err := time.ParseDuration(duration)
	if err != nil {
		return nil, nil, status.Error(codes.InvalidArgument, err.Error())
	}
	endTime := startTime.Add(timeEvent)
	return &startTime, &endTime, nil
}
