package timerRepository

import (
	"timer-api/pkg/domain/response"
	"timer-api/pkg/domain/timer"
	"timer-api/pkg/useCases/Helpers/databaseHelper"
)

// Create inserts a new timer record into the DB
func Create(t *timer.Timer) response.Status {
	db := databaseHelper.Db
	result := db.Omit("id").Create(t)
	if result.Error != nil {
		return response.StatusInternalServerError
	}
	return response.StatusCreated
}

// GetLastTimer retrieves the last created timer
func GetLastTimer() (timer.Timer, response.Status) {
	var t timer.Timer
	db := databaseHelper.Db
	result := db.Order("id desc").First(&t)
	if result.Error != nil {
		return t, response.StatusInternalServerError
	}
	return t, response.StatusOk
}

// DeleteByID performs a hard delete of a timer by id
func DeleteByID(id int) response.Status {
	db := databaseHelper.Db
	result := db.Delete(&timer.Timer{}, id)
	if result.Error != nil {
		return response.StatusInternalServerError
	}
	if result.RowsAffected == 0 {
		return response.StatusNotFound
	}
	return response.StatusOk
}
