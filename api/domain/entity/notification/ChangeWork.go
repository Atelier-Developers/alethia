package notification

type ChangeWork struct {
	UserHistoryId uint64 `json:"user_history_id"`
	Type          string `json:"type"`
	Notification
}
