package bodyTemplates

import "time"

// TODO: END DATE NULL HOW????
type UserBackgroundHistoryCreateRequestBody struct {
	StartDate   time.Time `json:"start_date"  binding:"required"`
	EndDate     time.Time `json:"end_date"  binding:"required"`
	Category    string    `json:"category"  binding:"required"`
	Title       string    `json:"title"  binding:"required"`
	Description string    `json:"description"  binding:"required"`
	Location    string    `json:"location"  binding:"required"`
}

type UserBackgroundHistoryUpdateRequestBody struct {
	ID          uint64    `json:"id" binding:"required"`
	StartDate   time.Time `json:"start_date"  binding:"required"`
	EndDate     time.Time `json:"end_date"  binding:"required"`
	Category    string    `json:"category"  binding:"required"`
	Title       string    `json:"title"  binding:"required"`
	Description string    `json:"description"  binding:"required"`
	Location    string    `json:"location"  binding:"required"`
}

type UserBackgroundHistoryDeleteRequestBody struct {
	ID uint64 `json:"id" binding:"required"`
}

type UserBackgroundHistoryGetRequestBody struct {
	UserId uint64 `json:"user_id" binding:"required"`
}
