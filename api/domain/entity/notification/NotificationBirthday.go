package notification

type Birthday struct {
	UserId uint64 `json:"user_id"`
	Notification
}
