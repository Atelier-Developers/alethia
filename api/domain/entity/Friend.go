package entity

import "time"

type Friend struct {
	UserId1 uint64    `json:"user_id_1"`
	UserId2 uint64    `json:"user_id_2"`
	Created time.Time `json:"created"`
}
