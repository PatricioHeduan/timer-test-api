package timer

import "time"

type Timer struct {
	Id        uint      `gorm:"primaryKey" json:"id"`
	Seconds   int       `json:"seconds"`
	StartedAt time.Time `json:"startedAt"`
}
