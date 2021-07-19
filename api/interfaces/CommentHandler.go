package interfaces

import (
	"github.com/Atelier-Developers/alethia/domain/entity/Comment"
	"github.com/Atelier-Developers/alethia/domain/entity/Post"
	"github.com/Atelier-Developers/alethia/domain/entity/notification"
	"github.com/Atelier-Developers/alethia/domain/repository"
	"github.com/Atelier-Developers/alethia/interfaces/bodyTemplates"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type CommentHandler struct {
	commentRepository      repository.CommentRepository
	notificationRepository repository.NotificationRepository
	postRepository         repository.PostRepository
}

func NewCommentHandler(commentRepository repository.CommentRepository, notificationRepository repository.NotificationRepository, postRepository repository.PostRepository) CommentHandler {
	return CommentHandler{
		commentRepository:      commentRepository,
		notificationRepository: notificationRepository,
		postRepository:         postRepository,
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

	comment := Comment.Comment{
		Text:        commentCreateRequestBody.Text,
		CommenterId: userId.(uint64),
		PostId:      commentCreateRequestBody.PostId,
	}

	err := commentHandler.commentRepository.SaveComment(&comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	var post Post.Post
	err = commentHandler.postRepository.GetPostById(comment.PostId, &post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	nc := notification.Comment{
		CommentId: comment.Id,
		Notification: notification.Notification{
			ReceiverId: post.PosterId,
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

	var comment Comment.Comment
	err := commentHandler.commentRepository.GetCommentByID(commentLikeRequestBody.CommentId, &comment)
	if err != nil {
		c.JSON(http.StatusConflict, err)
		return
	}

	err = commentHandler.commentRepository.LikeComment(&comment, userId.(uint64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	if userId.(uint64) != comment.CommenterId {
		n := notification.LikeComment{
			UserId:    userId.(uint64),
			CommentId: commentLikeRequestBody.CommentId,
			Notification: notification.Notification{
				ReceiverId: comment.CommenterId,
			},
		}
		err = commentHandler.notificationRepository.CreateLikeCommentNotification(&n)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	}
	c.JSON(http.StatusOK, nil)
}

func (commentHandler *CommentHandler) GetCommentLikesById(c *gin.Context) {
	_, exists := c.Get("user_id")
	if !exists {
		log.Fatal("User Id does not exist!")
	}

	commentId, err := strconv.ParseInt(c.Param("comment_id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	commentLikes, err := commentHandler.commentRepository.GetCommentLikes(uint64(commentId))
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, commentLikes)
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

	var repliedComment Comment.Comment
	err := commentHandler.commentRepository.GetCommentByID(commentReplyRequestBody.RepliedCommentId, &repliedComment)
	if err != nil {
		c.JSON(http.StatusConflict, err)
		return
	}

	comment := Comment.Comment{
		Text:        commentReplyRequestBody.Text,
		CommenterId: userId.(uint64),
		PostId:      commentReplyRequestBody.PostId,
	}
	err = commentHandler.commentRepository.SaveComment(&comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	id, err := commentHandler.commentRepository.ReplyComment(&comment, commentReplyRequestBody.RepliedCommentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if repliedComment.CommenterId != userId.(uint64) {
		n := notification.ReplyComment{
			CommentId: id,
			Notification: notification.Notification{
				ReceiverId: repliedComment.CommenterId,
			},
		}
		err = commentHandler.notificationRepository.CreateReplyCommentNotification(&n)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}
	}
	c.JSON(http.StatusOK, nil)
}
