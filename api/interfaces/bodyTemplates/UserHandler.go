package bodyTemplates

import (
	"time"
)

type UserCreateRequestBody struct {
	Username        string    `json:"username" binding:"required"`
	FirstName       string    `json:"first_name" binding:"required"`
	LastName        string    `json:"last_name" binding:"required"`
	Password        string    `json:"password" binding:"required"`
	Intro           string    `json:"intro" binding:"required"`
	About           string    `json:"about" binding:"required"`
	Accomplishments string    `json:"accomplishments" binding:"required"`
	AdditionalInfo  string    `json:"additional_info" binding:"required"`
	JoinDate        time.Time `json:"join_date"`
	BirthDate       time.Time `json:"birth_date" binding:"required"`
}

type UserUpdateRequestBody struct {
	Intro           string    `json:"intro"`
	About           string    `json:"about"`
	Accomplishments string    `json:"accomplishments"`
	AdditionalInfo  string    `json:"additional_info"`
	BirthDate       time.Time `json:"birth_date"`
}

type UserLoginRequestBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserGetResponseBody struct {
	UserID			uint64    `json:"user_id"`
	Username        string    `json:"username"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Intro           string    `json:"intro"`
	About           string    `json:"about"`
	Accomplishments string    `json:"accomplishments"`
	AdditionalInfo  string    `json:"additional_info"`
	BirthDate       time.Time `json:"birth_date"`
	JoinDate        time.Time `json:"join_date"`
}

type UserGetRequestBody struct {
	Id uint64 `json:"id" binding:"required"`
}

type UserViewProfileRequestBody struct {
	Id uint64 `json:"id" binding:"required"`
}
