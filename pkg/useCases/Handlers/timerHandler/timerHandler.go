package timerHandler

import (
	"time"
	"timer-api/internal/data/infrastructure/timerRepository"
	"timer-api/pkg/domain/response"
	"timer-api/pkg/domain/timer"
)

// CreateTimer starts/creates a new timer
func CreateTimer(t timer.Timer) (timer.Timer, response.Status) {
	t.StartedAt = time.Now()
	return t, timerRepository.Create(&t)
}

// GetTimer returns a timer by id
func GetLastTimer() (timer.Timer, response.Status) {
	return timerRepository.GetLastTimer()
}

// DeleteTimer deletes a timer by id
func DeleteTimer(id int) response.Status {
	return timerRepository.DeleteByID(id)
}
