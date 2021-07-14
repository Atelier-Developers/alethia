package interfaces

import (
	"github.com/Atelier-Developers/alethia/domain/entity"
	"github.com/Atelier-Developers/alethia/domain/repository"
	"github.com/Atelier-Developers/alethia/interfaces/bodyTemplates"
	"github.com/gin-gonic/gin"
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
	//TODO Check validity repost ID

	post := entity.Post{
		IsFeatured:  postCreateRequestBody.IsFeatured,
		Description: postCreateRequestBody.Description,
		Created:     time.Now(),
		RepostId:    postCreateRequestBody.RepostId,
		PosterId:    1,
	}

	err := postHandler.postRepository.SavePost(&post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, nil)
}
