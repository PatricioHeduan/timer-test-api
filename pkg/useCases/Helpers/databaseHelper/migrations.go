package databaseHelper

import (
	"fmt"
	"timer-api/pkg/domain/timer"

	"gorm.io/gorm"
)

func migrate(db *gorm.DB) {
	// Migrate timer table
	err := db.AutoMigrate(&timer.Timer{})
	if err != nil {
		panic("Error migrating tables: " + err.Error())
	}

	fmt.Println("Migration run successfully")
}
