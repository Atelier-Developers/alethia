package repository

import (
	"github.com/Atelier-Developers/alethia/domain/entity/Conversation"
)

type MessageRepository interface {
	AddMessage(message Conversation.Message) error
	GetMessages(conversationId uint64) ([]Conversation.Message, error)
	GetMessage(messageID uint64) (Conversation.Message, error)
	//ConversationExists(userId1 uint64, userId2 uint64) (bool, error)
	//DoesConversationBelongToUser(conversationId uint64, userId uint64) (bool, error)
	//SaveConversation(userId1 uint64, userId2 uint64) error
	//UpdateUserConversation(userConversation Conversation.UserConversation) error
	//DeleteConversation(conversation Conversation.Conversation) error
	////GetConversationMessages(conversationId uint64) ([]entity.Message, error)
	//GetUserConversations(userId uint64) ([]Conversation.UserConversation, error)
}
