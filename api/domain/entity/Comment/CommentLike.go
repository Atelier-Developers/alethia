package Comment

import "time"

type CommentLike struct {
	CommentId uint64    `json:"comment_id"`
	UserId    uint64    `json:"user_id"`
	Created   time.Time `json:"created"`
}
