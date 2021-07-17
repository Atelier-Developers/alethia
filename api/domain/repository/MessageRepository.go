package repository

import (
	"github.com/Atelier-Developers/alethia/domain/entity/Conversation"
)

type MessageRepository interface {
	AddMessage(message Conversation.Message) error
	GetMessages(conversationId uint64) ([]Conversation.Message, error)
	GetMessage(messageID uint64) (Conversation.Message, error)
}
