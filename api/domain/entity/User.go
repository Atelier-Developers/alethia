package entity

import (
	"database/sql"
	"time"
)

//TODO: ADD UNIQUE CONSTRAINTS USERNAME
//TODO: CHECK PASSWORD AND USERNAME LENGTH
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
	Location        string    `json:"location"`
}

type UserWithMutualConnection struct {
	User
	MutualConnection sql.NullInt64 `json:"mutual_connection"`
}

type UserWithMutualConnectionAndFriendshipStatus struct {
	UserWithMutualConnection
	IsFriendsWithThisUser uint64 `json:"is_friends_with_this_user"`
}
