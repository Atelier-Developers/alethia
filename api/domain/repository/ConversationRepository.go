package repository

import "github.com/Atelier-Developers/alethia/domain/entity"

type ConversationRepository interface {
	ConversationExists(userId1 uint64, userId2 uint64) (bool, error)
	SaveConversation(conversation *entity.Conversation) error
	UpdateConversation(conversation entity.Conversation) error
	DeleteConversation(conversation entity.Conversation) error
	GetConversation(userId1 uint64, userId2 uint64) (*entity.Conversation, error)
	GetUserConversations(userId uint64) ([]entity.Conversation, error)
}
