package entity

import "time"

type Message struct {
	Id             uint64    `json:"id"`
	ReplyId        uint64    `json:"reply_id"`
	ConversationId uint64    `json:"conversation_id"`
	Body           string    `json:"body"`
	Created        time.Time `json:"created"`
	IsRead         bool      `json:"is_read"`
}
