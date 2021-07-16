package Post

import "time"

type Post struct {
	Id          uint64    `json:"id"`
	IsFeatured  bool      `json:"is-featured"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	RepostId    uint64    `json:"repost_id"`
	PosterId    uint64    `json:"poster_id"`
}
