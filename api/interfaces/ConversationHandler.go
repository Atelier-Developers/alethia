package interfaces

import (
	"github.com/Atelier-Developers/alethia/domain/entity"
	"github.com/Atelier-Developers/alethia/domain/repository"
	"github.com/Atelier-Developers/alethia/infrastructure/auth"
	"github.com/Atelier-Developers/alethia/interfaces/bodyTemplates"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ConversationHandler struct {
	conversationRepository repository.ConversationRepository
	authInterface          auth.AuthInterface
	tokenInterface         auth.TokenInterface
}

func NewConversationHandler(conversationRepository repository.ConversationRepository, authInterface auth.AuthInterface, tokenInterface auth.TokenInterface) ConversationHandler {
	return ConversationHandler{
		conversationRepository: conversationRepository,
		authInterface:          authInterface,
		tokenInterface:         tokenInterface,
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

	conversation := entity.Conversation{UserId1: userId.(uint64), UserId2: userRequestBody.UserId2, IsArchived: userRequestBody.IsArchived, IsDeleted: userRequestBody.IsDeleted, IsRead: userRequestBody.IsRead}

	err = conversationHandler.conversationRepository.SaveConversation(&conversation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}
