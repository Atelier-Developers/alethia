package interfaces

import (
	"github.com/Atelier-Developers/alethia/domain/repository"
	"github.com/Atelier-Developers/alethia/interfaces/bodyTemplates"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type UserLanguageHandler struct {
	languageRepository repository.LanguageRepository
}

func NewUserLanguageHandler(languageRepository repository.LanguageRepository) UserLanguageHandler {
	return UserLanguageHandler{
		languageRepository: languageRepository,
	}
}

func (userLanguageHandler *UserLanguageHandler) GetLanguages(c *gin.Context) {
	languages, err := userLanguageHandler.languageRepository.GetLanguages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, languages)
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

	exists, err := userLanguageHandler.languageRepository.UserLanguageExists(userId.(uint64), userRequestBody.LanguageId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if exists {
		c.JSON(http.StatusConflict, err)
		return
	}

	err = userLanguageHandler.languageRepository.SaveUserLanguage(userId.(uint64), userRequestBody.LanguageId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (userLanguageHandler *UserLanguageHandler) DeleteUserLanguage(c *gin.Context) {
	var userRequestBody bodyTemplates.UserLanguageDeleteRequestBody
	languageId, err := strconv.ParseInt(c.Param("language_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, nil)
		return
	}
	userRequestBody.LanguageId = uint64(languageId)

	userId, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	err = userLanguageHandler.languageRepository.DeleteUserLanguage(userId.(uint64), userRequestBody.LanguageId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (userLanguageHandler *UserLanguageHandler) GetUserLanguages(c *gin.Context) {
	_, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	uId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	languages, err := userLanguageHandler.languageRepository.GetUserLanguages(uint64(uId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, languages)
}
