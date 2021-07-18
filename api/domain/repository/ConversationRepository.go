package repository

import (
	"github.com/Atelier-Developers/alethia/domain/entity/Conversation"
)

type ConversationRepository interface {
	GetUserConversation(conversationId uint64, userId uint64) (Conversation.UserConversation, error)
	GetConversationCorrespondents(conversationId uint64) (uint64, uint64, error)
	ConversationExists(userId1 uint64, userId2 uint64) (bool, error)
	DoesConversationBelongToUser(conversationId uint64, userId uint64) (bool, error)
	SaveConversation(userId1 uint64, userId2 uint64) error
	UpdateUserConversation(userConversation Conversation.UserConversation) error
	DeleteConversation(conversation Conversation.Conversation) error
	GetUserConversations(userId uint64) ([]Conversation.UserConversation, error)
	//GetUserConversationsByUsername(userId uint64, username string) ([]Conversation.UserConversation)
}
