package Conversation

type UserConversation struct {
	UserId         uint64 `json:"user_id"`
	Username       string `json:"username"`
	ConversationId uint64 `json:"conversation_id"`
	IsArchived     bool   `json:"is_archived"`
	IsDeleted      bool   `json:"is_deleted"`
	IsRead         bool   `json:"is_read"`
}
