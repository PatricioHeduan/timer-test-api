package timerhandler

import (
	"time"
	"timer-api/internal/data/infrastructure/timerRepository"
	"timer-api/pkg/domain/response"
	"timer-api/pkg/domain/timer"
)

type Handler struct{}

func New() Handler {
	return Handler{}
}

// CreateTimer starts/creates a new timer
func (h Handler) CreateTimer(t timer.Timer) response.Status {
	expireAt := time.Now().Add(time.Duration(t.Seconds) * time.Second)
	t.ExpireAt = expireAt
	return timerRepository.Create(&t)
}

// GetTimer returns a timer by id
func (h Handler) GetLastTimer() (timer.Timer, response.Status) {
	return timerRepository.GetLastTimer()
}
