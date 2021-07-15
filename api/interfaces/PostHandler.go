package interfaces

import (
	"github.com/Atelier-Developers/alethia/domain/entity"
	"github.com/Atelier-Developers/alethia/domain/repository"
	"github.com/Atelier-Developers/alethia/interfaces/bodyTemplates"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type PostHandler struct {
	postRepository repository.PostRepository
}

func NewPostHandler(postRepository repository.PostRepository) PostHandler {
	return PostHandler{
		postRepository: postRepository,
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

	post := entity.Post{
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

	var tmp entity.Post
	err := postHandler.postRepository.GetPostById(postLikeRequestBody.PostId, &tmp)
	if err != nil {
		c.JSON(http.StatusConflict, err)
		return
	}

	post := entity.Post{
		Id: postLikeRequestBody.PostId,
	}

	err = postHandler.postRepository.LikePost(&post, userId.(uint64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nil)
}
