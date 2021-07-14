package bodyTemplates

import "time"

type UserRequestBody struct {
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

type UserLoginRequestBody struct {
	Username  string	  `json:"username" binding:"required"`
	Password  string     `json:"password" binding:"required"`
}


