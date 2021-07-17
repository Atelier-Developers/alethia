package interfaces

import (
	"github.com/Atelier-Developers/alethia/domain/entity/Conversation"
	"github.com/Atelier-Developers/alethia/domain/repository"
	"github.com/Atelier-Developers/alethia/interfaces/bodyTemplates"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ConversationHandler struct {
	conversationRepository repository.ConversationRepository
}

func NewConversationHandler(conversationRepository repository.ConversationRepository) ConversationHandler {
	return ConversationHandler{
		conversationRepository: conversationRepository,
	}
}

func (conversationHandler ConversationHandler) AddConversation(c *gin.Context) {
	var userRequestBody bodyTemplates.ConversationCreateRequestBody
	if err := c.ShouldBindJSON(&userRequestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	userId, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	exists, err := conversationHandler.conversationRepository.ConversationExists(userId.(uint64), userRequestBody.UserId2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if exists {
		c.JSON(http.StatusConflict, err)
		return
	}

	err = conversationHandler.conversationRepository.SaveConversation(userId.(uint64), userRequestBody.UserId2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (conversationHandler ConversationHandler) UpdateUserConversation(c *gin.Context) {
	var userRequestBody bodyTemplates.ConversationUpdateRequestBody
	if err := c.ShouldBindJSON(&userRequestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	userId, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	conversationBelongToUser, err := conversationHandler.conversationRepository.DoesConversationBelongToUser(userRequestBody.ConversationId, userId.(uint64))

	if err != nil {
		log.Fatal(err)
	}

	if !conversationBelongToUser {
		log.Fatal("This is not this user's conversation.")
	}

	conversation := Conversation.UserConversation{
		UserId:         userId.(uint64),
		ConversationId: userRequestBody.ConversationId,
		IsArchived:     *userRequestBody.IsArchived,
		IsDeleted:      *userRequestBody.IsDeleted,
		IsRead:         *userRequestBody.IsRead,
	}

	err = conversationHandler.conversationRepository.UpdateUserConversation(conversation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, nil)

}

func (conversationHandler ConversationHandler) DeleteConversation(c *gin.Context) {
	var userRequestBody bodyTemplates.ConversationDeleteRequestBody
	if err := c.ShouldBindJSON(&userRequestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	userId, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	conversationBelongToUser, err := conversationHandler.conversationRepository.DoesConversationBelongToUser(userRequestBody.ConversationId, userId.(uint64))

	if err != nil {
		log.Fatal(err)
	}

	if !conversationBelongToUser {
		log.Fatal("This is not this user's conversation.")
	}

	conversation := Conversation.Conversation{Id: userRequestBody.ConversationId}

	err = conversationHandler.conversationRepository.DeleteConversation(conversation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, nil)

}

func (conversationHandler ConversationHandler) GetUserConversations(c *gin.Context) {
	userId, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}
	conversations, err := conversationHandler.conversationRepository.GetUserConversations(userId.(uint64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, conversations)
}
