package Conversation

import "time"

type Message struct {
	Id             uint64    `json:"id"`
	UserId         uint64    `json:"user_id"`
	Username       string    `json:"username"`
	ReplyId        uint64    `json:"reply_id"`
	ConversationId uint64    `json:"conversation_id"`
	Body           string    `json:"body"`
	Created        time.Time `json:"created"`
}
