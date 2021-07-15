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

type UserBackgroundHistoryHandler struct {
	backgroundHistoryRepository repository.BackgroundHistoryRepository
	authInterface               auth.AuthInterface
	tokenInterface              auth.TokenInterface
}

func NewUserBackgroundHistoryHandler(backgroundHistoryRepository repository.BackgroundHistoryRepository, authInterface auth.AuthInterface, tokenInterface auth.TokenInterface) UserBackgroundHistoryHandler {
	return UserBackgroundHistoryHandler{
		backgroundHistoryRepository: backgroundHistoryRepository,
		authInterface:               authInterface,
		tokenInterface:              tokenInterface,
	}
}

func (userBackgroundHistoryHandler *UserBackgroundHistoryHandler) AddBackgroundHistory(c *gin.Context) {
	var userRequestBody bodyTemplates.UserBackgroundHistoryCreateRequestBody
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

	backgroundHistory := entity.BackgroundHistory{
		UserID:      userId.(uint64),
		StartDate:   userRequestBody.StartDate,
		EndDate:     userRequestBody.EndDate,
		Category:    userRequestBody.Category,
		Title:       userRequestBody.Title,
		Description: userRequestBody.Description,
		Location:    userRequestBody.Location,
	}

	err := userBackgroundHistoryHandler.backgroundHistoryRepository.SaveBackgroundHistory(&backgroundHistory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	//TODO: CREATE NOTIF TOO?
	c.JSON(http.StatusCreated, nil)
}

func (userBackgroundHistoryHandler *UserBackgroundHistoryHandler) EditBackgroundHistory(c *gin.Context) {
	var userRequestBody bodyTemplates.UserBackgroundHistoryUpdateRequestBody
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

	backgroundHistory := entity.BackgroundHistory{
		ID:          userRequestBody.ID,
		StartDate:   userRequestBody.StartDate,
		EndDate:     userRequestBody.EndDate,
		Category:    userRequestBody.Category,
		Title:       userRequestBody.Title,
		Description: userRequestBody.Description,
		Location:    userRequestBody.Location,
	}

	err := userBackgroundHistoryHandler.backgroundHistoryRepository.UpdateBackgroundHistory(&backgroundHistory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// TODO CREATE NOTIF TOO? (ONLY EDIT END DATE NULL BACKGROUNDS + NOTIF CHANGE WORK ON END DATE NULL + GET ENDING DATE FOR ENDING BACKGROUND HISTORY)
	c.JSON(http.StatusCreated, nil)
}

func (userBackgroundHistoryHandler *UserBackgroundHistoryHandler) DeleteBackgroundHistory(c *gin.Context) {
	var userRequestBody bodyTemplates.UserBackgroundHistoryDeleteRequestBody
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

	backgroundHistory := entity.BackgroundHistory{
		ID: userRequestBody.ID,
	}

	err := userBackgroundHistoryHandler.backgroundHistoryRepository.DeleteBackgroundHistory(&backgroundHistory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, nil)
}

func (userBackgroundHistoryHandler *UserBackgroundHistoryHandler) GetUserBackgroundHistories(c *gin.Context) {
	userId, exists := c.Get("user_id")

	if !exists {
		log.Fatal("User Id does not exist!")
	}

	backgrounds, err := userBackgroundHistoryHandler.backgroundHistoryRepository.GetUserBackgroundHistory(userId.(uint64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, backgrounds)
}

