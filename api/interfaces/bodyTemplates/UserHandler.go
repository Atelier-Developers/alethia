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
	Location        string    `json:"location"`
}

type UserUpdateRequestBody struct {
	Intro           string    `json:"intro"`
	About           string    `json:"about"`
	Accomplishments string    `json:"accomplishments"`
	AdditionalInfo  string    `json:"additional_info"`
	BirthDate       time.Time `json:"birth_date"`
	Location        string    `json:"location"`
}

type UserLoginRequestBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserGetResponseBody struct {
	UserID          uint64    `json:"user_id"`
	Username        string    `json:"username"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	Intro           string    `json:"intro"`
	About           string    `json:"about"`
	Accomplishments string    `json:"accomplishments"`
	AdditionalInfo  string    `json:"additional_info"`
	BirthDate       time.Time `json:"birth_date"`
	Location        string    `json:"location"`
	JoinDate        time.Time `json:"join_date"`
}

type UserGetMutualConnectionsResponseBody struct {
	UserID           uint64    `json:"user_id"`
	Username         string    `json:"username"`
	FirstName        string    `json:"first_name"`
	LastName         string    `json:"last_name"`
	Intro            string    `json:"intro"`
	About            string    `json:"about"`
	Accomplishments  string    `json:"accomplishments"`
	AdditionalInfo   string    `json:"additional_info"`
	BirthDate        time.Time `json:"birth_date"`
	Location         string    `json:"location"`
	JoinDate         time.Time `json:"join_date"`
	MutualConnection uint64    `json:"mutual_connection"`
}

type UsersGetByUsernameResponseBody struct {
	UserID                uint64    `json:"user_id"`
	Username              string    `json:"username"`
	FirstName             string    `json:"first_name"`
	LastName              string    `json:"last_name"`
	Intro                 string    `json:"intro"`
	About                 string    `json:"about"`
	Accomplishments       string    `json:"accomplishments"`
	AdditionalInfo        string    `json:"additional_info"`
	BirthDate             time.Time `json:"birth_date"`
	Location              string    `json:"location"`
	JoinDate              time.Time `json:"join_date"`
	MutualConnection      uint64    `json:"mutual_connection"`
	IsFriendsWithThisUser uint64    `json:"is_friends_with_this_user"`
}

type UserGetByIdResponseBody struct {
	UserID                      uint64    `json:"user_id"`
	Username                    string    `json:"username"`
	FirstName                   string    `json:"first_name"`
	LastName                    string    `json:"last_name"`
	Intro                       string    `json:"intro"`
	About                       string    `json:"about"`
	Accomplishments             string    `json:"accomplishments"`
	AdditionalInfo              string    `json:"additional_info"`
	BirthDate                   time.Time `json:"birth_date"`
	Location                    string    `json:"location"`
	JoinDate                    time.Time `json:"join_date"`
	MutualConnection            uint64    `json:"mutual_connection"`
	IsFriendsWithThisUser       uint64    `json:"is_friends_with_this_user"`
	IsInvitedByThisUser         bool      `json:"is_invited_by_this_user"`
	HasConversationWithThisUser bool      `json:"has_conversation_with_this_user"`
}

type UserGetRequestBody struct {
	Id uint64 `json:"id" binding:"required"`
}

type UsersGetByUsernameRequestBody struct {
	Username string `json:"username"`
	Location string `json:"location"`
	Language string `json:"language"`
	Company  string `json:"company"`
}

type UserViewProfileRequestBody struct {
	Id uint64 `json:"id" binding:"required"`
}
