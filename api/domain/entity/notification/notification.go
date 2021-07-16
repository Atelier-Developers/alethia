package notification

import "time"

type Notification struct {
	Id         uint64    `json:"id"`
	ReceiverId uint64    `json:"receiver_id"`
	Created    time.Time `json:"created"`
	IsSeen     bool      `json:"is_seen"`
}
