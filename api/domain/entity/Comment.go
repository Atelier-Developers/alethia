package entity

import "time"

type Comment struct {
	Id          uint64    `json:"id"`
	Text        bool      `json:"text"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
	CommenterId uint64    `json:"commenter_id"`
	PostId      uint64    `json:"post_id"`
}
