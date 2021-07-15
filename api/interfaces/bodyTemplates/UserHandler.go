package bodyTemplates

import (
	"time"
)

type UserCreateRequestBody struct {
	Username  string	  `json:"username" binding:"required"`
	FirstName string     `json:"first_name" binding:"required"`
	LastName  string     `json:"last_name" binding:"required"`
	Password  string     `json:"password" binding:"required"`
	Intro	  string	 `json:"intro" binding:"required"`
	About	  string	 `json:"about" binding:"required"`
	Accomplishments	  string	 `json:"accomplishments" binding:"required"`
	AdditionalInfo	  string	 `json:"additional-info" binding:"required"`
	JoinDate		time.Time	  `json:"join-date"`
	BirthDate		time.Time	  `json:"birth-date" binding:"required"`
}

type UserUpdateRequestBody struct {
	Intro	  string	 `json:"intro"`
	About	  string	 `json:"about"`
	Accomplishments	  string	 `json:"accomplishments"`
	AdditionalInfo	  string	 `json:"additional-info"`
	BirthDate		time.Time	  `json:"birth-date"`
}

type UserLoginRequestBody struct {
	Username  string	  `json:"username" binding:"required"`
	Password  string     `json:"password" binding:"required"`
}

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
	ID          uint64    `json:"id" binding:"required"`
}



