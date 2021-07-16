package entity

import "time"

type Invite struct {
	UserId1 uint64    `json:"user_id_1"`
	UserId2 uint64    `json:"user_id_2"`
	Created time.Time `json:"created"`
	Body    string    `json:"body"`
}
