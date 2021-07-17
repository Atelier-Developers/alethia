package Post

import "time"

type Like struct {
	PostId   uint64    `json:"post_id"`
	UserId   uint64    `json:"user_id"`
	Created  time.Time `json:"created"`
	Username string    `json:"username"`
}
