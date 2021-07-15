package interfaces

import (
	"github.com/Atelier-Developers/alethia/domain/repository"
	"github.com/Atelier-Developers/alethia/infrastructure/auth"
	"github.com/Atelier-Developers/alethia/interfaces/bodyTemplates"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserLanguageHandler struct {
	languageRepository repository.LanguageRepository
	authInterface      auth.AuthInterface
	tokenInterface     auth.TokenInterface
}

func NewUserLanguageHandler(languageRepository repository.LanguageRepository, authInterface auth.AuthInterface, tokenInterface auth.TokenInterface) UserLanguageHandler {
	return UserLanguageHandler{
		languageRepository: languageRepository,
		authInterface:      authInterface,
		tokenInterface:     tokenInterface,
	}
}

func (userLanguageHandler *UserLanguageHandler) AddUserLanguage(c *gin.Context) {
	var userRequestBody bodyTemplates.UserLanguageCreateRequestBody
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

	err := userLanguageHandler.languageRepository.SaveUserLanguage(userId.(uint64), userRequestBody.LanguageId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (userLanguageHandler *UserLanguageHandler) DeleteUserLanguage(c *gin.Context) {
	var userRequestBody bodyTemplates.UserLanguageDeleteRequestBody
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

	err := userLanguageHandler.languageRepository.DeleteUserLanguage(userId.(uint64), userRequestBody.LanguageId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (userLanguageHandler *UserLanguageHandler) GetUserLanguages(c *gin.Context) {
	userId, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	languages, err := userLanguageHandler.languageRepository.GetUserLanguages(userId.(uint64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, languages)
}