package interfaces

import (
	"github.com/Atelier-Developers/alethia/domain/entity/Post"
	"github.com/Atelier-Developers/alethia/domain/entity/notification"
	"github.com/Atelier-Developers/alethia/domain/repository"
	"github.com/Atelier-Developers/alethia/interfaces/bodyTemplates"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

type PostHandler struct {
	postRepository         repository.PostRepository
	notificationRepository repository.NotificationRepository
}

func NewPostHandler(postRepository repository.PostRepository, notificationRepository repository.NotificationRepository) PostHandler {
	return PostHandler{
		postRepository:         postRepository,
		notificationRepository: notificationRepository,
	}
}

func (postHandler *PostHandler) SavePost(c *gin.Context) {
	var postCreateRequestBody bodyTemplates.PostCreateRequestBody
	if err := c.ShouldBindJSON(&postCreateRequestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	userId, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	post := Post.Post{
		IsFeatured:  false,
		Description: postCreateRequestBody.Description,
		Created:     time.Now(),
		RepostId:    postCreateRequestBody.RepostId,
		PosterId:    userId.(uint64),
	}

	err := postHandler.postRepository.SavePost(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, nil)
}

func (postHandler *PostHandler) LikePost(c *gin.Context) {
	var postLikeRequestBody bodyTemplates.PostLikeRequestBody
	if err := c.ShouldBindJSON(&postLikeRequestBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	userId, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	var likedPost Post.Post
	err := postHandler.postRepository.GetPostById(postLikeRequestBody.PostId, &likedPost)
	if err != nil {
		c.JSON(http.StatusConflict, err)
		return
	}

	post := Post.Post{
		Id: postLikeRequestBody.PostId,
	}

	err = postHandler.postRepository.LikePost(&post, userId.(uint64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	n := notification.LikePost{
		UserId: userId.(uint64),
		PostId: likedPost.Id,
		Notification: notification.Notification{
			ReceiverId: likedPost.PosterId,
		},
	}
	err = postHandler.notificationRepository.CreateLikePostNotification(&n)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, nil)
}

func (postHandler *PostHandler) GetPostById(c *gin.Context) {
	_, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	postId, err := strconv.ParseInt(c.Param("post_id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	var post Post.Post
	err = postHandler.postRepository.GetPostById(uint64(postId), &post)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, post)
}

func (postHandler *PostHandler) GetPostLikesById(c *gin.Context) {
	_, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	postId, err := strconv.ParseInt(c.Param("post_id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	likes, err := postHandler.postRepository.GetPostLikes(uint64(postId))
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, likes)
}

func (postHandler *PostHandler) GetPostCommentsById(c *gin.Context) {
	_, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	postId, err := strconv.ParseInt(c.Param("post_id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	comments, err := postHandler.postRepository.GetPostComments(uint64(postId))
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, comments)
}

func (postHandler *PostHandler) GetPostRepostsById(c *gin.Context) {
	_, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	postId, err := strconv.ParseInt(c.Param("post_id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	reposts, err := postHandler.postRepository.GetPostReposts(uint64(postId))
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, reposts)
}
