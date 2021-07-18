package bodyTemplates

import "time"

type MessageCreateRequestBody struct {
	ReplyId        uint64 `json:"reply_id"`
	ConversationId uint64 `json:"conversation_id" binding:"required"`
	MessageBody    string `json:"message_body" binding:"required"`
}

type MessagesGetRequestBody struct {
	ConversationId uint64 `json:"conversation_id" binding:"required"`
}

type MessageGetRequestBody struct {
	MessageId uint64 `json:"message_id" binding:"required"`
}

type MessageGetResponseBody struct {
	Id             uint64    `json:"id" binding:"required"`
	UserId         uint64    `json:"user_id"  binding:"required"`
	ReplyId        uint64    `json:"reply_id" binding:"required"`
	ConversationId uint64    `json:"conversation_id" binding:"required"`
	Body           string    `json:"body" binding:"required"`
	Created        time.Time `json:"created" binding:"required"`
}
