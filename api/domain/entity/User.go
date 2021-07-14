package entity

import (
	"time"
)

//TODO: ADD UNIQUE CONSTRAINTS USERNAME
type User struct {
	ID              uint64    `json:"id"`
	Username        string    `json:"username"  `
	FirstName       string    `json:"first_name"  `
	LastName        string    `json:"last_name"  `
	Password        []byte    `json:"password"  `
	Intro           string    `json:"intro"  `
	About           string    `json:"about"  `
	Accomplishments string    `json:"accomplishments"  `
	AdditionalInfo  string    `json:"additional-info"  `
	JoinDate        time.Time `json:"join-date"  `
	BirthDate       time.Time `json:"birth-date"  `
}
