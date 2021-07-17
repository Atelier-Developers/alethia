package Comment

import (
	"database/sql"
	"time"
)

type Comment struct {
	Id                uint64        `json:"id"`
	Text              string        `json:"text"`
	Description       string        `json:"description"`
	Created           time.Time     `json:"created"`
	CommenterId       uint64        `json:"commenter_id"`
	CommenterUsername string        `json:"commenter_username"`
	PostId            uint64        `json:"post_id"`
	RepliedCommentId  sql.NullInt64 `json:"replied_comment_id"`
}
