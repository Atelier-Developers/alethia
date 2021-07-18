package entity

import (
	"database/sql"
	"time"
)

type BackgroundHistory struct {
	ID          uint64       `json:"id"`
	UserID      uint64       `json:"user_id"`
	StartDate   time.Time    `json:"start_date"  `
	EndDate     sql.NullTime `json:"end_date"  `
	Category    string       `json:"category"  `
	Title       string       `json:"title"  `
	Description string       `json:"description"  `
	Location    string       `json:"location"  `
}
