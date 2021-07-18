package interfaces

import (
	"github.com/Atelier-Developers/alethia/domain/entity/Conversation"
	"github.com/Atelier-Developers/alethia/domain/repository"
	"github.com/Atelier-Developers/alethia/interfaces/bodyTemplates"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type MessageHandler struct {
	conversationRepository repository.ConversationRepository
	messageRepository      repository.MessageRepository
}

func NewMessageHandler(conversationRepository repository.ConversationRepository, messageRepository repository.MessageRepository) MessageHandler {
	return MessageHandler{
		conversationRepository: conversationRepository,
		messageRepository:      messageRepository,
	}
}

func (messageHandler MessageHandler) AddMessage(c *gin.Context) {
	var userRequestBody bodyTemplates.MessageCreateRequestBody
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

	conversationBelongToUser, err := messageHandler.conversationRepository.DoesConversationBelongToUser(userRequestBody.ConversationId, userId.(uint64))

	if err != nil {
		log.Fatal(err)
	}

	if !conversationBelongToUser {
		log.Fatal("This is not this user's conversation.")
	}

	message := Conversation.Message{
		UserId:         userId.(uint64),
		ReplyId:        userRequestBody.ReplyId,
		ConversationId: userRequestBody.ConversationId,
		Body:           userRequestBody.MessageBody,
	}

	err = messageHandler.messageRepository.AddMessage(message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	id1, id2, err := messageHandler.conversationRepository.GetConversationCorrespondents(userRequestBody.ConversationId)

	var correspondentId uint64
	if id1 == userId.(uint64) {
		correspondentId = id2
	} else {
		correspondentId = id1
	}

	conversation, err := messageHandler.conversationRepository.GetUserConversation(userRequestBody.ConversationId, correspondentId)
	if err != nil {
		log.Fatal(err)
	}

	newConversation := Conversation.UserConversation{
		UserId:         conversation.UserId,
		Username:       conversation.Username,
		ConversationId: conversation.ConversationId,
		IsArchived:     conversation.IsArchived,
		IsDeleted:      conversation.IsDeleted,
		IsRead:         false,
	}

	err = messageHandler.conversationRepository.UpdateUserConversation(newConversation)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusCreated, nil)
}

func (messageHandler MessageHandler) GetMessage(c *gin.Context) {
	var userRequestBody bodyTemplates.MessageGetRequestBody
	if err := c.ShouldBindJSON(&userRequestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	_, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	message, err := messageHandler.messageRepository.GetMessage(userRequestBody.MessageId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	response := bodyTemplates.MessageGetResponseBody{
		Id:             message.Id,
		UserId:         message.UserId,
		ReplyId:        message.ReplyId,
		ConversationId: message.ConversationId,
		Body:           message.Body,
		Created:        message.Created,
	}

	c.JSON(http.StatusOK, response)
}

func (messageHandler MessageHandler) GetMessages(c *gin.Context) {
	userId, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	converastionId, err := strconv.ParseInt(c.Param("conversation_id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	conversationBelongToUser, err := messageHandler.conversationRepository.DoesConversationBelongToUser(uint64(converastionId), userId.(uint64))

	if err != nil {
		log.Fatal(err)
	}

	if !conversationBelongToUser {
		log.Fatal("This is not this user's conversation.")
	}

	messages, err := messageHandler.messageRepository.GetMessages(uint64(converastionId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, messages)
}
