package notification

type LikeComment struct {
	UserId uint64 `json:"user_id"`
	CommentId uint64 `json:"comment_id"`
	Notification
}
