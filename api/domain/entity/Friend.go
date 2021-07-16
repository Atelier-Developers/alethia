package entity

import "time"

type Friend struct {
	UserId1   uint64    `json:"user_id_1"`
	Username1 string    `json:"username_1"`
	UserId2   uint64    `json:"user_id_2"`
	Username2 string    `json:"username_2"`
	Created   time.Time `json:"created"`
}
