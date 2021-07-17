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

type CommentWithLikeAndReplyCount struct {
	Comment
	LikeCount         uint64 `json:"like_count"`
	ReplyCount        uint64 `json:"reply_count"`
	IsLikedByThisUser bool   `json:"is_liked_by_this_user"`
}
