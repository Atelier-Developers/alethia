package notification

type LikePost struct {
	UserId uint64 `json:"user_id"`
	PostId uint64 `json:"post_id"`
	Notification
}
