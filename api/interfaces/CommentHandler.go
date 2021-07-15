package interfaces

import (
	"github.com/Atelier-Developers/alethia/domain/entity"
	"github.com/Atelier-Developers/alethia/domain/entity/notification"
	"github.com/Atelier-Developers/alethia/domain/repository"
	"github.com/Atelier-Developers/alethia/interfaces/bodyTemplates"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type CommentHandler struct {
	commentRepository      repository.CommentRepository
	notificationRepository repository.NotificationRepository
}

func NewCommentHandler(commentRepository repository.CommentRepository, notificationRepository repository.NotificationRepository, ) CommentHandler {
	return CommentHandler{
		commentRepository:      commentRepository,
		notificationRepository: notificationRepository,
	}
}

func (commentHandler *CommentHandler) SaveComment(c *gin.Context) {
	var commentCreateRequestBody bodyTemplates.CommentCreateRequestBody
	if err := c.ShouldBindJSON(&commentCreateRequestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	userId, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	comment := entity.Comment{
		Text:        commentCreateRequestBody.Text,
		CommenterId: userId.(uint64),
		PostId:      commentCreateRequestBody.PostId,
	}

	err := commentHandler.commentRepository.SaveComment(&comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	nc := notification.Comment{
		CommentId: comment.Id,
		Notification: notification.Notification{
			ReceiverId: userId.(uint64),
		},
	}
	err = commentHandler.notificationRepository.CreateCommentNotification(&nc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func (commentHandler *CommentHandler) LikeComment(c *gin.Context) {
	var commentLikeRequestBody bodyTemplates.CommentLikeRequestBody
	if err := c.ShouldBindJSON(&commentLikeRequestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	userId, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	var tmp entity.Comment
	err := commentHandler.commentRepository.GetCommentByID(commentLikeRequestBody.CommentId, &tmp)
	if err != nil {
		c.JSON(http.StatusConflict, err)
		return
	}

	comment := entity.Comment{
		Id: commentLikeRequestBody.CommentId,
	}

	err = commentHandler.commentRepository.LikeComment(&comment, userId.(uint64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (commentHandler *CommentHandler) ReplyComment(c *gin.Context) {
	var commentReplyRequestBody bodyTemplates.CommentReplyRequestBody
	if err := c.ShouldBindJSON(&commentReplyRequestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	userId, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	var tmp entity.Comment
	err := commentHandler.commentRepository.GetCommentByID(commentReplyRequestBody.RepliedCommentId, &tmp)
	if err != nil {
		c.JSON(http.StatusConflict, err)
		return
	}

	comment := entity.Comment{
		Text:        commentReplyRequestBody.Text,
		CommenterId: userId.(uint64),
		PostId:      commentReplyRequestBody.PostId,
	}
	err = commentHandler.commentRepository.SaveComment(&comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	err = commentHandler.commentRepository.ReplyComment(&comment, commentReplyRequestBody.RepliedCommentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}
